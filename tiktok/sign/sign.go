package sign

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"sort"
	"tiktokShop/tiktok/common/config"
)

//生成签名

type Sign struct {
	config *config.Config
}

var newServer *Sign

func GetNewService(config *config.Config) *Sign {
	if newServer == nil {
		newServer = &Sign{
			config: config,
		}
	}
	return newServer
}

// 生成请求tiktok shop 的 sign 签名
func (s *Sign) GetSign(api string, contentType string, query map[string]string, body map[string]any) string {
	//获取请求参数所有KEY集合
	keys := make([]string, len(query))
	idx := 0
	for k := range query {
		if k != "sign" && k != "access_token" {
			keys[idx] = k
			idx++
		}
	}

	// 按字母顺序重新排列参数的键
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	//以{key}{value}格式连接所有参数
	input := ""
	for _, key := range keys {
		input = input + key + query[key]
	}

	// 附加请求路径
	input = api + input

	// 如果请求报头的Content-type不是multipart/form-data，则将body附加到结尾
	if contentType != "multipart/form-data" {
		var jsonStr []byte
		if body != nil {
			jsonStr, _ = json.Marshal(body)
		}
		input = input + string(jsonStr)
	}

	// 将步骤5中生成的字符串与App secret包在一起
	input = s.config.App.Secret + input + s.config.App.Secret

	return s.generateSHA256(input, s.config.App.Secret)
}

// 生成签名
func (s *Sign) generateSHA256(input, secret string) string {
	// 对摘要字节流进行十六进制编码，使用sha256生成带盐的符号(secret)
	h := hmac.New(sha256.New, []byte(secret))
	if _, err := h.Write([]byte(input)); err != nil {
		return ""
	}

	return hex.EncodeToString(h.Sum(nil))
}
