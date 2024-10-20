package ziface

// IDataPack 封包、拆包
type IDataPack interface {
	// GetHeadLength 获取包头长度
	GetHeadLength() uint32

	// Pack 封包
	Pack(IMessage) ([]byte, error)

	// Unpack 拆包
	Unpack([]byte) (IMessage, error)
}
