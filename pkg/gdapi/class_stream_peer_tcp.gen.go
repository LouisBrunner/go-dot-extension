// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type StreamPeerTCP struct {
  obj gdc.ObjectPtr
}

func (me *StreamPeerTCP) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *StreamPeerTCP) BaseClass() string {
  return "StreamPeerTCP"
}

type StreamPeerTCPStatus int
const (
  StreamPeerTCPStatusStatusNone StreamPeerTCPStatus = 0
  StreamPeerTCPStatusStatusConnecting StreamPeerTCPStatus = 1
  StreamPeerTCPStatusStatusConnected StreamPeerTCPStatus = 2
  StreamPeerTCPStatusStatusError StreamPeerTCPStatus = 3
)

func  (me *StreamPeerTCP) Bind(port int, host String, ) { // TODO: return value
  // TODO: implement
}

func  (me *StreamPeerTCP) ConnectToHost(host String, port int, ) { // TODO: return value
  // TODO: implement
}

func  (me *StreamPeerTCP) Poll() { // TODO: return value
  // TODO: implement
}

func  (me *StreamPeerTCP) GetStatus() { // TODO: return value
  // TODO: implement
}

func  (me *StreamPeerTCP) GetConnectedHost() { // TODO: return value
  // TODO: implement
}

func  (me *StreamPeerTCP) GetConnectedPort() { // TODO: return value
  // TODO: implement
}

func  (me *StreamPeerTCP) GetLocalPort() { // TODO: return value
  // TODO: implement
}

func  (me *StreamPeerTCP) DisconnectFromHost() { // TODO: return value
  // TODO: implement
}

func  (me *StreamPeerTCP) SetNoDelay(enabled bool, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
