package fritzbox

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

const (
	MSG_TIME_FMT = "02.01.06 15:04:05"
)

/*const (
    CALL        = "CALL"
    RING        = "RING"
    DISCONNECT  = "DISCONNECT"
    CONNECT     = "CONNECT"
)
*/

// Message is the interface that all MQTT messages implement.
type Message interface {
	// Decode reads the message extended headers and payload from
	// r. Typically the values for hdr and packetRemaining will
	// be returned from Header.Decode.
	Decode(r io.Reader) error
}

const (
	MSG_CALL       = EventType("CALL")
	MSG_RING       = EventType("RING")
	MSG_CONNECT    = EventType("CONNECT")
	MSG_DISCONNECT = EventType("DISCONNECT")
)

type EventType string
type SubscriberNumber string

func (s EventType) String() string {
	return "[" + string(s) + "]"
}

func (s SubscriberNumber) String() string {
	return "tel:" + string(s)
}

type Header struct {
	Timestamp time.Time
	EventName EventType
	Id        int
}

func (hdr *Header) Decode(r io.Reader) (eventType EventType, err error) {

	*hdr = Header{
		//Timestamp: time.Parse(MSG_TIME_FMT,nil),
		EventName: EventType("nil"),
		//Id:        strconv.Atoi("0"),
	}
	return
}

type Extension struct {
	Connection string
}

func (hdr *Extension) Decode(r io.Reader) (err error) {
	return nil
}

type Event struct {
	Header
	InternalCallerId string
	LocalCallerId    SubscriberNumber
	RemoteCallerId   SubscriberNumber
	Duration         time.Duration
	Extension
}

type CallMessage struct {
	Header
	InternalCallerId string
	LocalCallerId    SubscriberNumber
	RemoteCallerId   SubscriberNumber
	Extension
}

func (msg *CallMessage) Decode(r io.Reader, hdr Header) (err error) {
	*msg = CallMessage{}
	return nil
}

type RingMessage struct {
	Header
	LocalCallerId  SubscriberNumber
	RemoteCallerId SubscriberNumber
	Extension
}

func (msg *RingMessage) Decode(r io.Reader, hdr Header) (err error) {
	*msg = RingMessage{}
	return nil
}

type ConnectMessage struct {
	Header
	LocalCallerId SubscriberNumber
	Extension
}

func (msg *ConnectMessage) Decode(r io.Reader, hdr Header) (err error) {
	*msg = ConnectMessage{}
	return nil
}

type DisconnectMessage struct {
	Header
	Duration time.Duration
	Extension
}

func (msg *DisconnectMessage) Decode(r io.Reader, hdr Header) (err error) {
	*msg = DisconnectMessage{}
	return nil
}

func EventFromString(line string) Event {

	l := strings.Split(line, ";")
	l = l[:len(l)-1] // remove last garbage element

	e := Event{}
	e.Timestamp, _ = time.Parse(MSG_TIME_FMT, l[0])
	e.EventName = EventType(l[1])
	e.Id, _ = strconv.Atoi(l[2])
	e.Connection = l[6]

	switch e.EventName {
	case CALL:
		e.InternalCallerId = l[3]
		e.LocalCallerId = SubscriberNumber(l[4])
		e.RemoteCallerId = SubscriberNumber(l[5])
		//e.Parameter           = l[6:]
	case RING:
		e.RemoteCallerId = SubscriberNumber(l[3])
		e.LocalCallerId = SubscriberNumber(l[4])
		//e.Parameter           = l[5:]
	case CONNECT:
		e.LocalCallerId = SubscriberNumber(l[3])
	case DISCONNECT:
		dur_int, _ := strconv.Atoi(l[3])
		e.Duration = time.Duration(dur_int)
	}
	return e
}

func MessageFromString(line string) (err error) {

	l := strings.Split(line, ";")
	l = l[:len(l)-1] // remove last garbage element

	h := Header{}
	h.Timestamp, _ = time.Parse(MSG_TIME_FMT, l[0])
	h.EventName = EventType(l[1])
	h.Id, _ = strconv.Atoi(l[2])

	switch h.EventName {
	case MSG_CALL:
		e := CallMessage{}
		e.Header = h
		e.InternalCallerId = l[3]
		e.LocalCallerId = SubscriberNumber(l[4])
		e.RemoteCallerId = SubscriberNumber(l[5])
		fmt.Println(e)
		//e.Parameter           = l[6:]
	case MSG_RING:
		e := RingMessage{}
		e.Header = h
		e.RemoteCallerId = SubscriberNumber(l[3])
		e.LocalCallerId = SubscriberNumber(l[4])
		//e.Parameter           = l[5:]
	case MSG_CONNECT:
		e := ConnectMessage{}
		e.Header = h
		e.LocalCallerId = SubscriberNumber(l[3])
	case MSG_DISCONNECT:
		e := DisconnectMessage{}
		e.Header = h
		dur_int, _ := strconv.Atoi(l[3])
		e.Duration = time.Duration(dur_int)
	}
	return nil
}

/*
func (e Event) String () string {
  return fmt.Sprintf("%s", e.Timestamp)
}
*/
