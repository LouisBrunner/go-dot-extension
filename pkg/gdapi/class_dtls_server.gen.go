// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type DTLSServer struct {
  obj gdc.ObjectPtr
}

func (me *DTLSServer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *DTLSServer) BaseClass() string {
  return "DTLSServer"
}



// Enums

func (me *DTLSServer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *DTLSServer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *DTLSServer) Setup(server_options TLSOptions, )  {
  panic("TODO: implement")
}

func  (me *DTLSServer) TakeConnection(udp_peer PacketPeerUDP, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
