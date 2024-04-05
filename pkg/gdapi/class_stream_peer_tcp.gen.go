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

type ptrsForStreamPeerTCPList struct {
  fnBind gdc.MethodBindPtr
  fnConnectToHost gdc.MethodBindPtr
  fnPoll gdc.MethodBindPtr
  fnGetStatus gdc.MethodBindPtr
  fnGetConnectedHost gdc.MethodBindPtr
  fnGetConnectedPort gdc.MethodBindPtr
  fnGetLocalPort gdc.MethodBindPtr
  fnDisconnectFromHost gdc.MethodBindPtr
  fnSetNoDelay gdc.MethodBindPtr
}

var ptrsForStreamPeerTCP ptrsForStreamPeerTCPList

func initStreamPeerTCPPtrs(iface gdc.Interface) {

  className := StringNameFromStr("StreamPeerTCP")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("bind")
    defer methodName.Destroy()
    ptrsForStreamPeerTCP.fnBind = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3167955072))
  }
  {
    methodName := StringNameFromStr("connect_to_host")
    defer methodName.Destroy()
    ptrsForStreamPeerTCP.fnConnectToHost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 993915709))
  }
  {
    methodName := StringNameFromStr("poll")
    defer methodName.Destroy()
    ptrsForStreamPeerTCP.fnPoll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166280745))
  }
  {
    methodName := StringNameFromStr("get_status")
    defer methodName.Destroy()
    ptrsForStreamPeerTCP.fnGetStatus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 859471121))
  }
  {
    methodName := StringNameFromStr("get_connected_host")
    defer methodName.Destroy()
    ptrsForStreamPeerTCP.fnGetConnectedHost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("get_connected_port")
    defer methodName.Destroy()
    ptrsForStreamPeerTCP.fnGetConnectedPort = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_local_port")
    defer methodName.Destroy()
    ptrsForStreamPeerTCP.fnGetLocalPort = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("disconnect_from_host")
    defer methodName.Destroy()
    ptrsForStreamPeerTCP.fnDisconnectFromHost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_no_delay")
    defer methodName.Destroy()
    ptrsForStreamPeerTCP.fnSetNoDelay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
}

type StreamPeerTCP struct {
  StreamPeer
}

func (me *StreamPeerTCP) BaseClass() string {
  return "StreamPeerTCP"
}

func NewStreamPeerTCP() *StreamPeerTCP {
  str := StringNameFromStr("StreamPeerTCP") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &StreamPeerTCP{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *StreamPeerTCP) Bind(port int64, host String, ) Error {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port) , host.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&port)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeerTCP.fnBind), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *StreamPeerTCP) ConnectToHost(host String, port int64, ) Error {
  cargs := []gdc.ConstTypePtr{host.AsCTypePtr(), gdc.ConstTypePtr(&port) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&port)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeerTCP.fnConnectToHost), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *StreamPeerTCP) Poll() Error {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeerTCP.fnPoll), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *StreamPeerTCP) GetStatus() StreamPeerTCPStatus {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret StreamPeerTCPStatus

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeerTCP.fnGetStatus), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *StreamPeerTCP) GetConnectedHost() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeerTCP.fnGetConnectedHost), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *StreamPeerTCP) GetConnectedPort() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeerTCP.fnGetConnectedPort), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StreamPeerTCP) GetLocalPort() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeerTCP.fnGetLocalPort), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StreamPeerTCP) DisconnectFromHost()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeerTCP.fnDisconnectFromHost), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeerTCP) SetNoDelay(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeerTCP.fnSetNoDelay), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
