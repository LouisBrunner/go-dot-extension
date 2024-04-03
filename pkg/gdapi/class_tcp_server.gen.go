// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TCPServer struct {
  RefCounted
}

func (me *TCPServer) BaseClass() string {
  return "TCPServer"
}

func NewTCPServer() *TCPServer {
  str := StringNameFromStr("TCPServer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &TCPServer{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *TCPServer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TCPServer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TCPServer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *TCPServer) Listen(port int64, bind_address String, ) Error {
  classNameV := StringNameFromStr("TCPServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("listen")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3167955072) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port), gdc.ConstTypePtr(bind_address.AsCTypePtr()), }
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TCPServer) IsConnectionAvailable() bool {
  classNameV := StringNameFromStr("TCPServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_connection_available")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TCPServer) IsListening() bool {
  classNameV := StringNameFromStr("TCPServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_listening")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TCPServer) GetLocalPort() int64 {
  classNameV := StringNameFromStr("TCPServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_local_port")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TCPServer) TakeConnection() StreamPeerTCP {
  classNameV := StringNameFromStr("TCPServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("take_connection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 30545006) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewStreamPeerTCP()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TCPServer) Stop()  {
  classNameV := StringNameFromStr("TCPServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("stop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
