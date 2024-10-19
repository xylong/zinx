package znet

import "github.com/spf13/cast"

// Message 消息
type Message struct {
	Id     uint64 // 消息id
	Length uint64 // 消息长度
	Data   []byte // 消息内容
}

func (m *Message) GetMsgId() uint64 {
	return m.Id
}

func (m *Message) SetMsgId(id uint64) {
	m.Id = id
}

func (m *Message) GetMsgLen() uint64 {
	return cast.ToUint64(len(m.Data))
}

func (m *Message) SetMsgLen(length uint64) {
	m.Length = length
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}
