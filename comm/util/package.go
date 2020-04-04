package util

import (
	"encoding/binary"
	"errors"
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

func RecvLTVPacket(reader io.Reader, maxPacketSize uint16) (interface{}, error) {
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
	if pkgSize > maxPacketSize {
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
