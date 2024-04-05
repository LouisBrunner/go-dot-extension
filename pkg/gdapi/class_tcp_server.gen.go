// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForTCPServerList struct {
  fnListen gdc.MethodBindPtr
  fnIsConnectionAvailable gdc.MethodBindPtr
  fnIsListening gdc.MethodBindPtr
  fnGetLocalPort gdc.MethodBindPtr
  fnTakeConnection gdc.MethodBindPtr
  fnStop gdc.MethodBindPtr
}

var ptrsForTCPServer ptrsForTCPServerList

func initTCPServerPtrs(iface gdc.Interface) {

  className := StringNameFromStr("TCPServer")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("listen")
    defer methodName.Destroy()
    ptrsForTCPServer.fnListen = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3167955072))
  }
  {
    methodName := StringNameFromStr("is_connection_available")
    defer methodName.Destroy()
    ptrsForTCPServer.fnIsConnectionAvailable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("is_listening")
    defer methodName.Destroy()
    ptrsForTCPServer.fnIsListening = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_local_port")
    defer methodName.Destroy()
    ptrsForTCPServer.fnGetLocalPort = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("take_connection")
    defer methodName.Destroy()
    ptrsForTCPServer.fnTakeConnection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 30545006))
  }
  {
    methodName := StringNameFromStr("stop")
    defer methodName.Destroy()
    ptrsForTCPServer.fnStop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
}

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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port) , bind_address.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&port)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTCPServer.fnListen), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TCPServer) IsConnectionAvailable() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTCPServer.fnIsConnectionAvailable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TCPServer) IsListening() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTCPServer.fnIsListening), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TCPServer) GetLocalPort() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTCPServer.fnGetLocalPort), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TCPServer) TakeConnection() StreamPeerTCP {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStreamPeerTCP()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTCPServer.fnTakeConnection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TCPServer) Stop()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTCPServer.fnStop), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
