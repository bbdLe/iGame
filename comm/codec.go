package comm

// 编解码接口
type Codec interface {
	// 编码
	Encode(msgObj interface{}, ctx ContextSet) (data interface{}, err error)

	// 解码
	Decode(data interface{}, msgObj interface{}) error

	// 编码器名字
	Name() string

	// http
	MimeType() string
}
