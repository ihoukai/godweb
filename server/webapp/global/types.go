package global

type Result struct {
	Code int         // 与parse errorcode 保持一致
	Msg  string      // 消息
	Data interface{} // 数据实体
}
