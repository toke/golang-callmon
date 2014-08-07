package main

import (
    "fmt"
    "./fritzbox"
)


func handleMessages(msgchan <-chan fritzbox.FbEvent){
  ev := <- msgchan
 
  fmt.Printf("Event: %s\n", ev)
  if ev.EventName == "CALL" {
    fmt.Printf("Event: %s\n", ev.Destination)
  } else if ev.EventName == "RING" {
    fmt.Printf("Event: %s\n", ev.Source)
  }

  fmt.Printf("! %s\n", ev)
}

func main() {
  c := new(fritzbox.CallmonHandler).Connect("192.168.92.1")
  if c.Connected {
    recv := make(chan fritzbox.FbEvent)
    go handleMessages(recv)
 
    // Inject a test message
    f := c.Parse("06.08.14 14:52:26;CALL;1;10;50000001;012344567;SIP1;")
    recv <- f

    c.Loop(recv)
  }
  c.Close()

  fmt.Println("NEVER EVER GONNA GIVE YOU UP")

}
