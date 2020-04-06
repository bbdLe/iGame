package codec

import (
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/err"
)

func DecodeMessage(msgId int, data []byte) (interface{}, *comm.MessageMeta, error) {
	meta := comm.MessageMetaByID(msgId)
	if meta == nil {
		return nil, nil, err.NewErrorContext("MsgId not exist", msgId)
	}

	val := meta.NewType()
	if err := meta.Codec.Decode(data, val); err != nil {
		return nil, meta, err
	}

	return val, meta, nil
}

func EncodeMessage(msg interface{}, ctx comm.ContextSet) ([]byte, *comm.MessageMeta, error) {
	meta := comm.MessageMetaByMsg(msg)
	if meta == nil {
		return nil, nil, err.NewErrorContext("Msg Type not exist", ctx)
	}

	data, err := meta.Codec.Encode(msg, ctx)
	if err != nil {
		return nil, meta, err
	}

	return data.([]byte), meta, nil
}