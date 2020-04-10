//go:generate protoc --gofast_out=. *.proto

package proto

import (
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/codec"
	"github.com/bbdLe/iGame/comm/util"
	"reflect"

	_ "github.com/bbdLe/iGame/comm/codec/gogopb"
)

func init() {
	comm.RegMessageMeta(&comm.MessageMeta{
		MsgId: int(util.StringHash("proto.EchoReq")),
		Type: reflect.TypeOf((*EchoReq)(nil)).Elem(),
		Codec: codec.MustGetCodec("gogopb"),
	})

	comm.RegMessageMeta(&comm.MessageMeta{
		MsgId: int(util.StringHash("proto.EchoRes")),
		Type: reflect.TypeOf((*EchoRes)(nil)).Elem(),
		Codec: codec.MustGetCodec("gogopb"),
	})
}