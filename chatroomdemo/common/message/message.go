package common

const (
	LoginMesType    = "loginMes"
	LoginResMesType = "loginResMes"
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` // 消息的类型
}

// 定义两个消息类型
type LoginMes struct {
	UserId   int    `json:"userid"`
	UserPwd  string `json:"userpwd"`
	UserName string `json:"username"`
}

type LoginResMes struct {
	Code  int    `json:"code"`  // 返回状态码，500 未注册 200注册成功
	Error string `json:"error"` // 返回错误信息
}
