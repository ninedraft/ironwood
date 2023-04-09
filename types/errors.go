package types

import "fmt"

// Error is any error generated by the PacketConn. Note that other errors may still be returned, if e.g. HandleConn returns due to a network error. An Error may be wrapped to provide additional context.
type Error uint

const (
	ErrUndefined Error = iota
	ErrEncode
	ErrDecode
	ErrClosed
	ErrTimeout
	ErrBadMessage
	ErrEmptyMessage
	ErrOversizedMessage
	ErrUnrecognizedMessage
	ErrPeerNotFound
	ErrBadAddress
	ErrBadKey
)

func (e Error) Error() string {
	var s string
	switch e {
	case ErrUndefined:
		s = "Undefined"
	case ErrEncode:
		s = "Encode"
	case ErrDecode:
		s = "Decode"
	case ErrClosed:
		s = "Closed"
	case ErrTimeout:
		s = "Timeout"
	case ErrBadMessage:
		s = "BadMessage"
	case ErrEmptyMessage:
		s = "EmptyMessage"
	case ErrOversizedMessage:
		s = "OversizedMessage"
	case ErrUnrecognizedMessage:
		s = "UnrecognizedMessage"
	case ErrPeerNotFound:
		s = "PeerNotFound"
	case ErrBadAddress:
		s = "BadAddress"
	case ErrBadKey:
		s = "BadKey"
	default:
		s = fmt.Sprintf("Unrecognized error code: %d", e)
	}
	return s
}