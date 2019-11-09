package common

import (
	"io"

	"github.com/vmihailenco/msgpack"
)

// Message struct is used to exchange the control messages between the peers
// in the control stream.
type Message struct {
	Command string `msgpack:"command"`
	Context string `msgpack:"context"`
}

// NewMessage is used to create a new message object
func NewMessage(command, context string) *Message {
	return &Message{
		Command: command,
		Context: context,
	}
}

// EncodeTo is used to write the msgpack notation of the message
// struct
func (m *Message) EncodeTo(w io.Writer) error {
	return msgpack.NewEncoder(w).Encode(m)
}

// DecodeFrom is used to decode the message into the message struct
func (m *Message) DecodeFrom(r io.Reader) (*Message, error) {
	return m, msgpack.NewDecoder(r).Decode(m)
}
