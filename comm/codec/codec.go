package codec

import "github.com/bbdLe/iGame/comm/peer"

// 编解码接口
type Codec interface {
	// 编码
	Encode(msgObj interface{}, context peer.ContextSet) (data interface{}, err error)

	// 解码
	Decode(data interface{}, msgObj interface{}) error

	// 编码器名字
	Name() string

	// http
	MimeType() string
}
