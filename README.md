golang-callmon
==============

Fritzbox callmon tools

Basic implementation of a callmon library for Fritzbox [callmonitor][callmon] feature.

This is my very first [go][golang] program so even if it works for my use case it does
not mean it follows best practices.

Use it by enabeling the call monitor feature on your Fritzbox by dailing `#96*5*`
install golang and `go run example.go` alternative append the IP of your Fritzbox.
It injects one Demo CALL in the output. Calling over your Fritzbox or getting a call
should print some output on your terminal.

The example code currently does not much but shows the library is working.

## Install

    git get github.com/toke/golang-callmon

## Run Example

    go run ${GOPATH}/src/github.com/toke/golang-callmon/example.go fritz.box



[callmon]: http://www.wehavemorefun.de/fritzbox/Callmonitor
[golang]: http://golang.org/
