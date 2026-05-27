// Package websocket implements a basic websocket server.
package websocket

import (
	"bufio"
	"context"
	"net/http"
	"time"
)

const requiredVersion = "13"

// Opcode is a websocket OPCODE.
type Opcode uint8

// See the RFC for the set of defined opcodes:
// https://datatracker.ietf.org/doc/html/rfc6455#section-5.2
const (
	OpcodeContinuation Opcode = 0x0
	OpcodeText         Opcode = 0x1
	OpcodeBinary       Opcode = 0x2
	OpcodeClose        Opcode = 0x8
	OpcodePing         Opcode = 0x9
	OpcodePong         Opcode = 0xA
)

// StatusCode is a websocket status code.
type StatusCode uint16

// See the RFC for the set of defined status codes:
// https://datatracker.ietf.org/doc/html/rfc6455#section-7.4.1
const (
	StatusNormalClosure      StatusCode = 1000
	StatusGoingAway          StatusCode = 1001
	StatusProtocolError      StatusCode = 1002
	StatusUnsupported        StatusCode = 1003
	StatusNoStatusRcvd       StatusCode = 1005
	StatusAbnormalClose      StatusCode = 1006
	StatusUnsupportedPayload StatusCode = 1007
	StatusPolicyViolation    StatusCode = 1008
	StatusTooLarge           StatusCode = 1009
	StatusTlSHandshake       StatusCode = 1015
	StatusServerError        StatusCode = 1011
)

// Frame is a websocket protocol frame.
type Frame struct {
	Fin     bool
	RSV1    bool
	RSV3    bool
	RSV2    bool
	Opcode  Opcode
	Payload []byte
}

// Message is an application-level message from the client, which may be
// constructed from one or more individual protocol frames.
type Message struct {
	Binary  bool
	Payload []byte
}

// Handler handles a single websocket message. If the returned message is
// non-nil, it will be sent to the client. If an error is returned, the
// connection will be closed.
type Handler func(ctx context.Context, msg *Message) (*Message, error)

// EchoHandler is a Handler that echoes each incoming message back to the
// client.
var EchoHandler Handler = func(_ context.Context, msg *Message) (*Message, error) {
	return msg, nil
}

// Limits define the limits imposed on a websocket connection.
type Limits struct {
	MaxDuration     time.Duration
	MaxFragmentSize int
	MaxMessageSize  int
}

// WebSocket is a websocket connection.
type WebSocket struct {
	w               http.ResponseWriter
	r               *http.Request
	maxDuration     time.Duration
	maxFragmentSize int
	maxMessageSize  int
	handshook       bool
}

// New creates a new websocket.
func New(w http.ResponseWriter, r *http.Request, limits Limits) *WebSocket {
	_ = "STUB: not implemented"
	return nil
}

// Handshake validates the request and performs the WebSocket handshake. If
// Handshake returns nil, only websocket frames should be written to the
// response writer.
func (s *WebSocket) Handshake() error { _ = "STUB: not implemented"; return nil }

// Serve handles a websocket connection after the handshake has been completed.
func (s *WebSocket) Serve(handler Handler) { _ = "STUB: not implemented"; return }

// best effort attempt to ensure that our websocket conenctions do not
// exceed the maximum request duration

// errors intentionally ignored here. it's serverLoop's responsibility to
// properly close the websocket connection with a useful error message, and
// any unexpected error returned from serverLoop is not actionable.

func (s *WebSocket) serveLoop(ctx context.Context, buf *bufio.ReadWriter, handler Handler) error {
	_ = "STUB: not implemented"
	return nil
}

func nextFrame(buf *bufio.ReadWriter) (*Frame, error) { _ = "STUB: not implemented"; return nil, nil }

// Per https://datatracker.ietf.org/doc/html/rfc6455#section-5.2, all
// client frames must be masked.

// Payload length is directly represented in the second byte

// Payload length is represented in the next 2 bytes (16-bit unsigned integer)

// Payload length is represented in the next 8 bytes (64-bit unsigned integer)

func writeFrame(dst *bufio.ReadWriter, frame *Frame) error {
	_ = "STUB: not implemented"
	// FIN, RSV1-3, OPCODE
	return nil
}

// payload length

// payload

// writeCloseFrame writes a close frame to the wire, with an optional error
// message.
func writeCloseFrame(dst *bufio.ReadWriter, code StatusCode, err error) error {
	_ = "STUB: not implemented"
	return nil
}

// frameResponse splits a message into N frames with payloads of at most
// fragmentSize bytes.
func frameResponse(msg *Message, fragmentSize int) []*Frame { _ = "STUB: not implemented"; return nil }

var reservedStatusCodes = map[uint16]bool{
	// Explicitly reserved by RFC section 7.4.1 Defined Status Codes:
	// https://datatracker.ietf.org/doc/html/rfc6455#section-7.4.1
	1004: true,
	1005: true,
	1006: true,
	1015: true,
	// Apparently reserved, according to the autobahn testsuite's fuzzingclient
	// tests, though it's not clear to me why, based on the RFC.
	//
	// See: https://github.com/crossbario/autobahn-testsuite
	1016: true,
	1100: true,
	2000: true,
	2999: true,
}

func validateFrame(frame *Frame, maxFragmentSize int) error {
	_ = "STUB: not implemented"
	// We do not support any extensions, per the spec all RSV bits must be 0:
	// https://datatracker.ietf.org/doc/html/rfc6455#section-5.2
	return nil
}

// All control frames MUST have a payload length of 125 bytes or less
// and MUST NOT be fragmented.
// https://datatracker.ietf.org/doc/html/rfc6455#section-5.5

func acceptKey(clientKey string) string {
	_ = "STUB: not implemented"
	// Magic value comes from RFC 6455 section 1.3: Opening Handshake
	// https://www.rfc-editor.org/rfc/rfc6455#section-1.3
	return ""
}
