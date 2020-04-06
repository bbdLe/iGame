package util

import (
	"encoding/binary"
	"errors"
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/codec"
	"io"
)

var (
	ErrMaxPacket = errors.New("packet over size")
	ErrMinPacket = errors.New("packet short size")
	ErrMsgIdPacket = errors.New("packet msg id wrong")
)

const (
	bodySize = 2
	msgIdSize = 2
)

func RecvLTVPacket(reader io.Reader, maxPacketSize int) (interface{}, error) {
	var sizeBuf = make([]byte, bodySize)
	n, err := io.ReadFull(reader, sizeBuf)
	if err != nil {
		return nil, err
	}
	if n < bodySize {
		return nil, ErrMinPacket
	}

	pkgSize := binary.LittleEndian.Uint16(sizeBuf)
	// 包太大
	if int(pkgSize) > maxPacketSize {
		return nil, ErrMaxPacket
	}

	var pkgBuf = make([]byte, pkgSize)
	n, err = io.ReadFull(reader, pkgBuf)
	if err != nil {
		return nil, err
	}
	if n < msgIdSize {
		return nil, ErrMsgIdPacket
	}

	msgId := binary.LittleEndian.Uint16(pkgBuf)
	msgBuf := pkgBuf[msgIdSize:]

	msg, _, err := codec.DecodeMessage(int(msgId), msgBuf)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func SendLTVPacket(writer io.Writer, ctx comm.ContextSet, msg interface{})  error {
	var (
		msgData []byte
		msgID int
		meta *comm.MessageMeta
	)

	switch m := msg.(type) {
	case *comm.RawPacket:
		msgData = m.Data
		msgID = m.MsgID
	default:
		var err error
		msgData, meta, err = codec.EncodeMessage(msg, ctx)
		if err != nil {
			return err
		}

		msgID = meta.MsgId
	}

	pkg := make([]byte, bodySize + msgIdSize + len(msgData))

	// Length
	binary.LittleEndian.PutUint16(pkg, uint16(msgIdSize + len(msgData)))
	// Type
	binary.LittleEndian.PutUint16(pkg[bodySize:], uint16(msgID))
	// Data
	copy(pkg[bodySize + msgIdSize:], msgData)

	// Write
	err := WriteFull(writer, pkg)
	if err != nil {
		return err
	}

	return nil
}
