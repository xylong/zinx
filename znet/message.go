package znet

// Message 消息
type Message struct {
	Id     uint32 // 消息id
	Length uint32 // 消息长度
	Data   []byte // 消息内容
}

func (m *Message) GetMsgId() uint32 {
	return m.Id
}

func (m *Message) SetMsgId(id uint32) {
	m.Id = id
}

func (m *Message) GetMsgLen() uint32 {
	return m.Length
}

func (m *Message) SetMsgLen(length uint32) {
	m.Length = length
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}
