package peer

import (
	"net"
	"time"
)

type CoreTcpSocketOption struct {
	readBufferSize int
	writeBufferSize int
	noDelay bool
	maxPacketSize int

	readTimeout time.Duration
	writeTimeout time.Duration
}

func (self *CoreTcpSocketOption) SetSocketBuffer(readBufferSize, writeBufferSize int, noDelay bool) {
	self.readBufferSize = readBufferSize
	self.writeBufferSize = writeBufferSize
	self.noDelay = noDelay
}

func (self *CoreTcpSocketOption) SetSocketDeadline(read, write time.Duration) {
	self.readTimeout = read
	self.writeTimeout = write
}

func (self *CoreTcpSocketOption) SetMaxPacketSize(maxSize int) {
	self.maxPacketSize = maxSize
}

func (self *CoreTcpSocketOption) MaxPacketSize() int {
	return self.maxPacketSize
}

func (self *CoreTcpSocketOption) ApplySocketOption(conn net.Conn) {
	if c, ok := conn.(*net.TCPConn); ok {
		if self.readBufferSize > 0 {
			c.SetReadBuffer(self.readBufferSize)
		}

		if self.writeBufferSize > 0 {
			c.SetWriteBuffer(self.writeBufferSize)
		}

		c.SetNoDelay(self.noDelay)
	}
}

func (self *CoreTcpSocketOption) ApplySocketReadTimeout(conn net.Conn, cb func()) {
	if self.readTimeout > 0 {
		conn.SetReadDeadline(time.Now().Add(self.readTimeout))
		cb()
		conn.SetReadDeadline(time.Time{})
	} else {
		cb()
	}
}

func (self *CoreTcpSocketOption) ApplySocketWriteTimeout(conn net.Conn, cb func()) {
	if self.writeTimeout > 0 {
		conn.SetWriteDeadline(time.Now().Add(self.writeTimeout))
		cb()
		conn.SetWriteDeadline(time.Time{})
	} else {
		cb()
	}
}