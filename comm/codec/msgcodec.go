package codec

import (
	"github.com/bbdLe/iGame/comm/err"
	"github.com/bbdLe/iGame/comm/meta"
)

func DecodeMessage(msgId int, data []byte) (interface{}, *meta.MessageMeta, error) {
	meta := meta.MessageMetaByID(msgId)
	if meta == nil {
		return nil, nil, err.NewErrorContext("MsgId not exist", msgId)
	}

	val := meta.NewType()
	if err := meta.Codec.Decode(data, val); err != nil {
		return nil, meta, err
	}

	return val, meta, nil
}
