package codec

import "github.com/bbdLe/iGame/comm/err"

var (
	registedCodeccs []Codec
)

func RegisterCodec(c Codec) {
	if GetCodec(c.Name()) != nil {
		panic(err.NewError("duplicate codec :" + c.Name()))
	}

	registedCodeccs = append(registedCodeccs, c)
}

func GetCodec(name string) Codec {
	for _, c := range registedCodeccs {
		if c.Name() == name {
			return c
		}
	}

	return nil
}

func MustGetCodec(name string) Codec {
	c := GetCodec(name)
	if c == nil {
		panic(err.NewError("Get Codec fail :" + name))
	} else {
		return c
	}
}