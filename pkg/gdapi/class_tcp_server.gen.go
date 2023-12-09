// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TCPServer struct {
  obj gdc.ObjectPtr
}

func (me *TCPServer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TCPServer) BaseClass() string {
  return "TCPServer"
}



// Enums

func (me *TCPServer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TCPServer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *TCPServer) Listen(port int, bind_address String, )  {
  panic("TODO: implement")
}

func  (me *TCPServer) IsConnectionAvailable()  {
  panic("TODO: implement")
}

func  (me *TCPServer) IsListening()  {
  panic("TODO: implement")
}

func  (me *TCPServer) GetLocalPort()  {
  panic("TODO: implement")
}

func  (me *TCPServer) TakeConnection()  {
  panic("TODO: implement")
}

func  (me *TCPServer) Stop()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
