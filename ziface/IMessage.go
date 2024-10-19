package ziface

type IMessage interface {
	// GetMsgId 获取消息id
	GetMsgId() uint64
	// SetMsgId 设置消息id
	SetMsgId(uint64)

	// GetMsgLen 获取消息长度
	GetMsgLen() uint64
	// SetMsgLen 设置消息长度
	SetMsgLen(uint64)

	// GetData 获取消息内容
	GetData() []byte
	// SetData 设置消息内容
	SetData([]byte)
}
