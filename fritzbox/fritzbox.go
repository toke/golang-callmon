package fritzbox

import (
    "fmt"
    "time"
    "strings"
    "net"
    "bufio"
    "log"
)


const (
    FB_TIME_FMT = "02.01.06 15:04:05"
    CALL        = "CALL"
    RING        = "RING"
    DISCONNECT  = "DISCONNECT"
    CONNECT     = "CONNECT"
)

type CallmonHandler struct {
  conn      net.Conn
  Connected bool
  Host      string
  event     chan FbEvent
}

type FbEvent struct {
  Timestamp   time.Time
  EventName   string
  Id          string
  Internal    string   // InternalCallerId
  Source      string   // LocalCallerId
  Destination string   // RemoteCallerId
  Duration    string   // time.Duration
  Parameter []string
}

type Reason struct {
  Timestamp         time.Time
  EventName         string
  Id                int
  InternalCallerId  string
  LocalCallerId     string
  RemoteCallerId    string
  Duration          time.Duration
  Parameter       []string
}


func (e FbEvent) Notify (recv chan FbEvent) {
    recv <- e
}

func (e FbEvent) String () string {
  return fmt.Sprintf("Event %s: %s", e.Timestamp, e.EventName)
}



func (c CallmonHandler) Connect(host string, recv chan FbEvent) CallmonHandler {
  c.Host     = host
  c.event    = recv

  addr, err := net.ResolveTCPAddr("tcp", host + ":1012")
  if err != nil {
    log.Fatal(err)
  }

  conn, err := net.DialTCP("tcp", nil, addr)
  if err!= nil {
    log.Fatal(err)
    c.Connected = false
  } else {
    c.Connected = true
  }

  conn.SetKeepAlivePeriod(time.Duration(30) * time.Second)
  conn.SetKeepAlive(true)

  c.conn = conn
  return c
}

func (c CallmonHandler) Close() {
  if c.conn != nil {
    c.conn.Close()
  }
  c.Connected = false
}


func (c CallmonHandler) read_loop() {

}

func (c CallmonHandler) Loop() {
  connbuf := bufio.NewReader(c.conn)
  for {
    str, err := connbuf.ReadString('\n')
    if err!= nil {
      c.conn.Close()
      log.Println(err)
      break
    }
    if len(str)>0 {
      c.event <- c.Parse(str)
    }
  }
}


func (c CallmonHandler) Parse(line string) FbEvent {
  l := strings.Split(line, ";")
  time, _ := time.Parse(FB_TIME_FMT, l[0])

  e:= FbEvent{
    Timestamp: time,
    EventName: l[1],
    Id:        l[2],
  }

  switch e.EventName {
    case CALL:
      e.Internal    = l[3]
      e.Source      = l[4]
      e.Destination = l[5]
      e.Parameter   = l[6:]
    case RING:
      e.Source      = l[3]
      e.Destination = l[4]
      e.Parameter   = l[5:]
    case CONNECT:
      e.Destination = l[3]
    case DISCONNECT:
      e.Duration    = l[3]
  }

  return e
}
