package websocket

import (
	"encoding/binary"
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/codec"
	"github.com/bbdLe/iGame/comm/util"
	"github.com/gorilla/websocket"
)

const (
	MsgIdLen = 2
)

type WebSocketTransmiiter struct {
}

func (WebSocketTransmiiter) OnRecvMessage(sess comm.Session) (interface{}, error) {
	conn := sess.Raw().(*websocket.Conn)
	msgType, data, err := conn.ReadMessage()
	if err != nil {
		return nil, err
	}

	if len(data) < MsgIdLen {
		return nil, util.ErrMsgIdPacket
	}

	switch msgType {
	case websocket.BinaryMessage:
		msgId := binary.LittleEndian.Uint16(data)

		msg, _, err := codec.DecodeMessage(int(msgId), data[MsgIdLen:])
		if err != nil {
			return nil, err
		}

		return msg, nil
	}

	return nil, nil
}

func (WebSocketTransmiiter) OnSendMessage(sess comm.Session, msg interface{}) (err error) {
	conn, ok := sess.Raw().(*websocket.Conn)
	if !ok || conn == nil {
		return nil
	}

	var (
		msgData []byte
		msgid int
	)

	switch m := msg.(type) {
	case *comm.RawPacket:
		msgData = m.Data
		msgid = m.MsgID
	default:
		data, meta, err := codec.EncodeMessage(msg, nil)
		if err != nil {
			return err
		}
		msgid = meta.MsgId
		msgData = data
	}

	buff := make([]byte, MsgIdLen + len(msgData))
	binary.LittleEndian.PutUint16(buff, uint16(msgid))
	copy(buff[MsgIdLen:], msgData)

	err = conn.WriteMessage(websocket.BinaryMessage, buff)
	if err != nil {
		return err
	}

	return
}