golang-callmon
==============

Fritzbox callmon tools

Basic implementation of a callmon library for Fritzbox [callmonitor][callmon]
feature.

This is my very first [go][golang] program so even if it works for my use case
it does not mean it follows best practices. But I try to improve it as I learn.

Use it by enabeling the call monitor feature on your Fritzbox by dailing
`#96*5*` install golang and `go run callmon-example/main.go` alternative append
the IP of your Fritzbox. It injects one Demo CALL in the output. Calling over
your Fritzbox or getting a call should print some output on your terminal.

The example code currently does not much but shows the library is working.

## Install

    go get github.com/toke/golang-callmon
    go install github.com/toke/golang-callmon/cmd/callmon-example

## Run Example

    callmon-example fritz.box

This example will also output some test data in the beginning. It is not
ment to do anything in production.


[callmon]: http://www.wehavemorefun.de/fritzbox/Callmonitor
[golang]: http://golang.org/
