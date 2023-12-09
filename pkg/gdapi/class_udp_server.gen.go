// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type UDPServer struct {
  obj gdc.ObjectPtr
}

func (me *UDPServer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *UDPServer) BaseClass() string {
  return "UDPServer"
}



// Enums

func (me *UDPServer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *UDPServer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *UDPServer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *UDPServer) Listen(port int, bind_address String, )  {
  panic("TODO: implement")
}

func  (me *UDPServer) Poll()  {
  panic("TODO: implement")
}

func  (me *UDPServer) IsConnectionAvailable()  {
  panic("TODO: implement")
}

func  (me *UDPServer) GetLocalPort()  {
  panic("TODO: implement")
}

func  (me *UDPServer) IsListening()  {
  panic("TODO: implement")
}

func  (me *UDPServer) TakeConnection()  {
  panic("TODO: implement")
}

func  (me *UDPServer) Stop()  {
  panic("TODO: implement")
}

func  (me *UDPServer) SetMaxPendingConnections(max_pending_connections int, )  {
  panic("TODO: implement")
}

func  (me *UDPServer) GetMaxPendingConnections()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
