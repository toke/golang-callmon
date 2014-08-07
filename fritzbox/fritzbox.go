package fritzbox

import (
    "fmt"
    "time"
    "strings"
    "net"
    "bufio"
    "log"
)


const FB_TIME_FMT = "02.01.06 15:04:05"

type CallmonHandler struct {
  conn      net.Conn
  cb        func(FbEvent)
  eventchan chan<- FbEvent
  Connected bool
  Host      string
}

type FbEvent struct {
  Timestamp   time.Time
  EventName   string
  Id          string
  Internal    string
  Source      string
  Destination string
  Duration    string
  Parameter []string
}

func (e FbEvent) Event() FbEvent {
  return e
}


func (e FbEvent) String () string {
  return fmt.Sprintf("Event %s: %s", e.Timestamp, e.EventName)
}



func (c CallmonHandler) Connect(host string) CallmonHandler {
  c.Host = host
  conn, err := net.DialTimeout("tcp", host + ":1012", time.Duration(30)*time.Second)
  if err!= nil {
    log.Fatal(err)
    c.Connected = false
  } else {
    c.Connected = true
  }

  c.conn = conn
  return c
}

func (c CallmonHandler) Close() {
  if c.conn != nil {
    c.conn.Close()
  }
  c.Connected = false
}


func (c CallmonHandler) Loop(recv chan FbEvent) {
  connbuf := bufio.NewReader(c.conn)

  for {
    str, err := connbuf.ReadString('\n')
    if len(str)>0 {
      recv <- c.Parse(str) 
    }
    if err!= nil {
      log.Println(err)
      break
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
    case "CALL":
      e.Internal    = l[3]
      e.Source      = l[4]
      e.Destination = l[5]
      e.Parameter   = l[6:]
    case "RING":
      e.Source      = l[3]
      e.Destination = l[4]
      e.Parameter   = l[5:]
    case "CONNECT":
      e.Destination = l[3]
    case "DISCONNECT":
      e.Duration    = l[3]
  }

  return e
}
