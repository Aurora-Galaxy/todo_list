package serializer

// Response 基础序列化器
 //响应的内容
type Response struct {
	Status int  `json:"status"`
	Data interface{} `json:"data"`
	Total int `json:"total"`
	Msg string `json:"msg"`
	Error string `json:"error"`
}

// TokenData 带token的Data结构
type TokenData struct{
	User interface{} `json:"user"`
	Token string  `json:"token"`
}