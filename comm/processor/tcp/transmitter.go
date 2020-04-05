package tcp

import (
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/bbdLe/iGame/comm/session"
	"github.com/bbdLe/iGame/comm/util"
	"io"
	"net"
)

type TCPMessageTransmitter struct {
}

type socketOpt interface {
	MaxPacketSize() uint16
	ApplySocketReadTimeout(conn net.Conn, cb func())
	ApplySocketWriteTimeout(conn net.Conn, cb func())
}

func (TCPMessageTransmitter) OnRecvMessage(sess session.Session) (msg interface{}, err error) {
	reader, ok := sess.Raw().(io.Reader)
	if !ok || reader == nil {
		return nil,nil
	}

	opt := sess.Peer().(socketOpt)
	if conn, ok := reader.(net.Conn); ok {
		opt.ApplySocketReadTimeout(conn, func() {
			msg, err = util.RecvLTVPacket(reader, opt.MaxPacketSize())
		})
	}

	return
}

func (TCPMessageTransmitter) OnSendMessage(sess session.Session, msg interface{}) (err error) {
	writer, ok := sess.Raw().(io.Writer)
	if !ok || writer == nil {
		return nil
	}

	opt := sess.Peer().(socketOpt)
	if conn, ok := writer.(net.Conn); ok {
		opt.ApplySocketWriteTimeout(conn, func() {
			err = util.SendLTVPacket(writer, sess.(peer.ContextSet), msg)
		})
	}

	return
}
