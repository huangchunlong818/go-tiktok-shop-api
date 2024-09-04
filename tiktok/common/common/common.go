package common

//共用tiktok配置
import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/config"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/sign"
	"strconv"
	"time"
)

type TiktokShopCommon struct {
	config *config.Config
}

var newServer *TiktokShopCommon

// getNewService 是一个私有函数，用于返回 tiktokShopAuths 实例
func GetNewService(config *config.Config) *TiktokShopCommon {
	if newServer == nil {
		newServer = &TiktokShopCommon{
			config: config,
		}
	}
	return newServer
}

// 通用结构体
type ComApiRsp struct {
	Code     int            `json:"code"`     //逻辑状态码
	Message  string         `json:"message"`  //错误信息
	Data     map[string]any `json:"data"`     //数据
	HttpCode int            `json:"httpCode"` //请求tiktok的HTTP状态码
}

// 通用错误体
type ComErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// 发送请求参数
type SendParams struct {
	Api     string            //请求API地址，带域名，全地址，不带URL参数
	Query   map[string]string //URL具体请求参数
	Body    map[string]any    //请求体具体请求参数
	Headers map[string]string //具体头部请求参数
	Method  string            //请求方式 get post put delete
	Files   map[string]string //文件上传
}

// 通用tiktok shop api 请求接口类型定义
type GetApiConfig struct {
	ContentType, Method, Api, FullApi, Token string
}

var (
	ErrCode     = 777 //请求tiktok接口发生错误的时候，自定义的错误码
	restyClient = resty.New()
)

// 通用结构体
type TiktokApiRsp struct {
	Code     int    `json:"code"`     //逻辑状态码
	Message  string `json:"message"`  //错误信息
	Data     []byte `json:"data"`     //数据
	HttpCode int    `json:"httpCode"` //请求tiktok的HTTP状态码
}

// 通用tiktok shop api 请求   reqs 接口基本信息，带token    query URL参数(不带app_key，sign，timestamp)   body 请求体参数
func (c *TiktokShopCommon) SendTiktokApi(ctx context.Context, reqs GetApiConfig, query map[string]string, body map[string]any, files map[string]string) TiktokApiRsp {
	result := TiktokApiRsp{}
	if reqs.Token == "" && reqs.FullApi == "" {
		result.Code = ErrCode
		result.Message = "Token和请求API不能为空"
		return result
	}

	//共用参数设置进query,主要是app_key，sign，timestamp
	if query == nil {
		query = make(map[string]string)
	}
	query["app_key"] = c.config.App.AppKey
	query["timestamp"] = strconv.FormatInt(time.Now().UTC().Unix(), 10)
	query["sign"] = sign.GetNewService(c.config).GetSign(reqs.Api, reqs.ContentType, query, body) //获取签名

	//设置头部参数 x-tts-access-token  和 Content-Type
	header := map[string]string{
		"Content-Type":       reqs.ContentType,
		"x-tts-access-token": reqs.Token,
	}

	//组装请求参数
	params := SendParams{
		Api:     reqs.FullApi,
		Method:  reqs.Method,
		Query:   query,
		Body:    body,
		Headers: header,
		Files:   files,
	}

	//请求接口
	return c.SendApi(ctx, params)
}

func (c *TiktokShopCommon) IsSuccess(r TiktokApiRsp) bool {
	return r.HttpCode > 199 && r.HttpCode < 300 && r.Code == 0
}

// 发起API请求 目前适用tiktok shop 所有API请求，跟auth 区分开
func (c *TiktokShopCommon) SendApi(ctx context.Context, params SendParams) TiktokApiRsp {
	//定义响应体
	var (
		res    ComApiRsp
		errRsp ComErrorResponse
		resp   *resty.Response
	)
	result := TiktokApiRsp{
		Code:     0,
		Message:  "",
		HttpCode: 200,
	}

	//请求tiktok
	var err error
	timeout := 10 * time.Second

	// 文件上传会比较耗时
	if len(params.Files) > 0 {
		timeout = 3 * time.Minute
	}

	restyClient.SetTimeout(timeout)

	tmpResty := restyClient.R().
		SetContext(ctx). //如果ctx.Done()通道关闭，则中断请求执行
		SetQueryParams(params.Query).
		SetBody(params.Body)

	// 文件上传
	if len(params.Files) > 0 {
		tmpResty = tmpResty.SetFiles(params.Files)
	}

	tmpResty = tmpResty.
		SetHeaders(params.Headers).
		SetResult(&res).
		SetError(&errRsp)

	switch params.Method {
	case "post":
		resp, err = tmpResty.Post(params.Api)
	case "get":
		resp, err = tmpResty.Get(params.Api)
	case "put":
		resp, err = tmpResty.Put(params.Api)
	case "delete":
		resp, err = tmpResty.Delete(params.Api)
	default:
		err = errors.New("请求方式错误")
	}
	if err != nil {
		result.Code = ErrCode
		result.Message = err.Error()
		return result
	}
	if resp.IsSuccess() {
		//这里也有可能是失败的
		if res.Code > 0 {
			result.Code = res.Code
			result.Message = res.Message
		} else {
			result.Data = resp.Body()
		}
	} else {
		result.HttpCode = resp.StatusCode()
		result.Code = errRsp.Code
		result.Message = errRsp.Message
	}

	return result
}
