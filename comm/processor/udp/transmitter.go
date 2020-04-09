package udp

import (
	"encoding/binary"
	"fmt"

	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/codec"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/peer/udp"
)

const (
	MTU = 1472
	PacketLen = 2
	MsgIDLen = 2
	HeaderLen = PacketLen + MsgIDLen
)

type UDPMessageTransmitter struct {
}

func (UDPMessageTransmitter) OnRecvMessage(sess comm.Session) (msg interface{}, err error) {
	data := sess.Raw().(udp.DataReader).ReadData()
	return recvPacket(data)
}

func (UDPMessageTransmitter) OnSendMessage(sess comm.Session, msg interface{}) (err error) {
	return sendPacket(sess.Raw().(udp.DataWriter), sess.(comm.ContextSet), msg)
}

func recvPacket(pkgData []byte) (msg interface{}, err error) {
	if len(pkgData) < PacketLen {
		return nil, nil
	}

	datasize := binary.LittleEndian.Uint16(pkgData)
	if int(datasize) != len(pkgData) || datasize > MTU {
		return nil, nil
	}

	msgid := binary.LittleEndian.Uint16(pkgData[PacketLen:])
	msgData := pkgData[HeaderLen:]

	msg, _, err = codec.DecodeMessage(int(msgid), msgData)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func sendPacket(writer udp.DataWriter, ctx comm.ContextSet, pkgData interface{}) error {
	msgData, meta, err := codec.EncodeMessage(pkgData, ctx)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("EncodeMessage fail : %v", err))
		return err
	}

	buff := make([]byte, HeaderLen + len(msgData))

	binary.LittleEndian.PutUint16(buff, uint16(HeaderLen + len(msgData)))
	binary.LittleEndian.PutUint16(buff[PacketLen:], uint16(meta.MsgId))
	copy(buff[HeaderLen:], msgData)

	writer.WriteData(buff)

	return nil
}
