package gogopb

import (
	"github.com/bbdLe/iGame/comm/codec"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/gogo/protobuf/proto"
)

type gogopbCodec struct {}

func (self *gogopbCodec) Name() string {
	return "gogopb"
}

func (self *gogopbCodec) MimeType() string {
	return "appplication/x-protobuf"
}

func (self *gogopbCodec) Encode(msgObj interface{}, ctx peer.ContextSet) (data interface{}, err error) {
	return proto.Marshal(msgObj.(proto.Message))
}

func (self *gogopbCodec) Decode(data interface{}, msgObj interface{}) error {
	return proto.Unmarshal(data.([]byte), msgObj.(proto.Message))
}

func init() {
	codec.RegisterCodec(new(gogopbCodec))
}