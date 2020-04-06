package sysmsg

import "fmt"

type CloseReason int32

const (
	CloseReason_IO CloseReason = iota
	CloseReason_Manual
)

func (self CloseReason) String() string {
	switch self {
	case CloseReason_IO:
		return "IO"
	case CloseReason_Manual:
		return "Manual"
	}

	return "Unknown"
}

type SessionClose struct {
	Reason CloseReason
}

func (self *SessionClose) String() string {
	return fmt.Sprintf("%+v", *self)
}

type SessionConnectError struct {}
type SessionConnected struct {}
type SessionAccepted struct {}

func (self *SessionConnectError) String() string { return fmt.Sprintf("%+v", *self)}
func (self *SessionConnected) String() string { return fmt.Sprintf("%+v", *self)}
func (self *SessionAccepted) String() string { return fmt.Sprintf("%+v", *self)}
