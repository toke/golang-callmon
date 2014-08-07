golang-callmon
==============

Fritzbox callmon tools

Basic implementation of a callmon library for Fritzbox [callmonitor][1] feature.

This is my very first go program so even if it works for my use case it does
not mean it follows best practices.

Us it by enabeling the call monitor feature on your Fritzbox by dailing `#96*5*`
install golang and `go run example.go` alternative append the IP of your Fritzbox.
It injects one Demo CALL in the output. Calling over your Fritzbox or getting a call
should print some output on your terminal.

The example code currently does not much but shows the library is working. 

[1]: http://www.wehavemorefun.de/fritzbox/Callmonitor


