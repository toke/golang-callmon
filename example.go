package main

import (
    "os"
    "fmt"
    "encoding/json"
    "github.com/toke/golang-callmon/fritzbox"
)


func handleMessages(msgchan <-chan fritzbox.FbEvent){
  for {
      ev := <- msgchan

      jsonm,_ := json.Marshal(&ev)
      fmt.Printf("Some JSON: %s\n", jsonm)

      if ev.EventName == fritzbox.CALL {
        fmt.Printf("%s Event: %s->%s\n", ev.EventName, ev.Source, ev.Destination)
      } else if ev.EventName == fritzbox.RING {
        fmt.Printf("%s Event: %s->%s\n", ev.EventName, ev.Source, ev.Destination)
      } else {
          fmt.Printf("! %s\n", ev)
      }
  }
}

func mainloop(host string) {
  c := new(fritzbox.CallmonHandler).Connect(host)

  defer c.Close()

  if c.Connected {
    recv := make(chan fritzbox.FbEvent)
    go handleMessages(recv)
 
    // Inject a test message
    f := c.Parse("06.08.14 14:52:26;CALL;1;10;50000001;012344567;SIP1;\r\n")
    recv <- f

    c.Loop(recv)
  }
}

func main() {
  arg := os.Args

  host := "fritz.box"
  if (len(arg) > 1 && arg[1] != "") {
    host = arg[1]
  }

  mainloop(host)
  fmt.Println("NEVER EVER GONNA GIVE YOU UP")
}
