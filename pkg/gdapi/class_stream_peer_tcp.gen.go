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



// Enums

type StreamPeerTCPStatus int
const (
  StreamPeerTCPStatusStatusNone StreamPeerTCPStatus = 0
  StreamPeerTCPStatusStatusConnecting StreamPeerTCPStatus = 1
  StreamPeerTCPStatusStatusConnected StreamPeerTCPStatus = 2
  StreamPeerTCPStatusStatusError StreamPeerTCPStatus = 3
)

func (me *StreamPeerTCP) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *StreamPeerTCP) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *StreamPeerTCP) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *StreamPeerTCP) Bind(port int, host String, )  {
  panic("TODO: implement")
}

func  (me *StreamPeerTCP) ConnectToHost(host String, port int, )  {
  panic("TODO: implement")
}

func  (me *StreamPeerTCP) Poll()  {
  panic("TODO: implement")
}

func  (me *StreamPeerTCP) GetStatus()  {
  panic("TODO: implement")
}

func  (me *StreamPeerTCP) GetConnectedHost()  {
  panic("TODO: implement")
}

func  (me *StreamPeerTCP) GetConnectedPort()  {
  panic("TODO: implement")
}

func  (me *StreamPeerTCP) GetLocalPort()  {
  panic("TODO: implement")
}

func  (me *StreamPeerTCP) DisconnectFromHost()  {
  panic("TODO: implement")
}

func  (me *StreamPeerTCP) SetNoDelay(enabled bool, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
