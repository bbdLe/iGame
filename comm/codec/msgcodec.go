package codec

import (
	"github.com/bbdLe/iGame/comm/err"
	"github.com/bbdLe/iGame/comm/meta"
	"github.com/bbdLe/iGame/comm/peer"
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

func EncodeMessage(msg interface{}, ctx peer.ContextSet) ([]byte, *meta.MessageMeta, error) {
	meta := meta.MessageMetaByMsg(msg)
	if meta == nil {
		return nil, nil, err.NewErrorContext("Msg Type not exist", ctx)
	}

	data, err := meta.Codec.Encode(msg, ctx)
	if err != nil {
		return nil, meta, err
	}

	return data.([]byte), meta, nil
}