package sign

// 生成sign的请求参数
type SignParam struct {
	Param  map[string]string //请求参数
	Path   string            //请求路径
	Header string            //请求头中的  Content-type 值
	Body   []byte            //请求体
}
