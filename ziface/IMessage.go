package ziface

type IMessage interface {
	// GetMsgId 获取消息id
	GetMsgId() uint32
	// SetMsgId 设置消息id
	SetMsgId(uint322 uint32)

	// GetMsgLen 获取消息长度
	GetMsgLen() uint32
	// SetMsgLen 设置消息长度
	SetMsgLen(uint322 uint32)

	// GetData 获取消息内容
	GetData() []byte
	// SetData 设置消息内容
	SetData([]byte)
}
