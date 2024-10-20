package znet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"zinx/utils"
	"zinx/ziface"
)

// DataPack 数据包
type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (p *DataPack) GetHeadLength() uint32 {
	// length uint32(4子节) + id uint32(4子节)
	return 8
}

// Pack 封包
// length｜id｜data
func (p *DataPack) Pack(message ziface.IMessage) ([]byte, error) {
	var (
		err    error
		buffer = bytes.NewBuffer([]byte{})
	)

	// 将length写进缓冲区
	if err = binary.Write(buffer, binary.LittleEndian, message.GetMsgLen()); err != nil {
		return nil, err
	}

	// 将id写进缓冲区
	if err = binary.Write(buffer, binary.LittleEndian, message.GetMsgId()); err != nil {
		return nil, err
	}

	// 将data写进缓冲区
	if err = binary.Write(buffer, binary.LittleEndian, message.GetData()); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// Unpack 解包
// 将包的head信息读出来，再根据head信息中的data长度，再进行一次读取
func (p *DataPack) Unpack(data []byte) (ziface.IMessage, error) {
	var (
		err     error
		buffer  = bytes.NewReader(data)
		message = &Message{}
	)

	// 读取length
	if err = binary.Read(buffer, binary.LittleEndian, &message.Length); err != nil {
		return nil, err
	}

	// 读取id
	if err = binary.Read(buffer, binary.LittleEndian, &message.Id); err != nil {
		return nil, err
	}

	// 判断length是否超出最大限制
	if utils.GlobalObject.MaxPackageSize > 0 && utils.GlobalObject.MaxPackageSize < message.Length {
		return nil, fmt.Errorf("too larage")
	}

	return message, nil

}
