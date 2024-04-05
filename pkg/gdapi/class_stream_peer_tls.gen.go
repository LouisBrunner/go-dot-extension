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

type ptrsForStreamPeerTLSList struct {
  fnPoll gdc.MethodBindPtr
  fnAcceptStream gdc.MethodBindPtr
  fnConnectToStream gdc.MethodBindPtr
  fnGetStatus gdc.MethodBindPtr
  fnGetStream gdc.MethodBindPtr
  fnDisconnectFromStream gdc.MethodBindPtr
}

var ptrsForStreamPeerTLS ptrsForStreamPeerTLSList

func initStreamPeerTLSPtrs(iface gdc.Interface) {

  className := StringNameFromStr("StreamPeerTLS")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("poll")
    defer methodName.Destroy()
    ptrsForStreamPeerTLS.fnPoll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("accept_stream")
    defer methodName.Destroy()
    ptrsForStreamPeerTLS.fnAcceptStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4292689651))
  }
  {
    methodName := StringNameFromStr("connect_to_stream")
    defer methodName.Destroy()
    ptrsForStreamPeerTLS.fnConnectToStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 57169517))
  }
  {
    methodName := StringNameFromStr("get_status")
    defer methodName.Destroy()
    ptrsForStreamPeerTLS.fnGetStatus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1128380576))
  }
  {
    methodName := StringNameFromStr("get_stream")
    defer methodName.Destroy()
    ptrsForStreamPeerTLS.fnGetStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2741655269))
  }
  {
    methodName := StringNameFromStr("disconnect_from_stream")
    defer methodName.Destroy()
    ptrsForStreamPeerTLS.fnDisconnectFromStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
}

type StreamPeerTLS struct {
  StreamPeer
}

func (me *StreamPeerTLS) BaseClass() string {
  return "StreamPeerTLS"
}

func NewStreamPeerTLS() *StreamPeerTLS {
  str := StringNameFromStr("StreamPeerTLS") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &StreamPeerTLS{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type StreamPeerTLSStatus int
const (
  StreamPeerTLSStatusStatusDisconnected StreamPeerTLSStatus = 0
  StreamPeerTLSStatusStatusHandshaking StreamPeerTLSStatus = 1
  StreamPeerTLSStatusStatusConnected StreamPeerTLSStatus = 2
  StreamPeerTLSStatusStatusError StreamPeerTLSStatus = 3
  StreamPeerTLSStatusStatusErrorHostnameMismatch StreamPeerTLSStatus = 4
)

func (me *StreamPeerTLS) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *StreamPeerTLS) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *StreamPeerTLS) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *StreamPeerTLS) Poll()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeerTLS.fnPoll), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeerTLS) AcceptStream(stream StreamPeer, server_options TLSOptions, ) Error {
  cargs := []gdc.ConstTypePtr{stream.AsCTypePtr(), server_options.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeerTLS.fnAcceptStream), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *StreamPeerTLS) ConnectToStream(stream StreamPeer, common_name String, client_options TLSOptions, ) Error {
  cargs := []gdc.ConstTypePtr{stream.AsCTypePtr(), common_name.AsCTypePtr(), client_options.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeerTLS.fnConnectToStream), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *StreamPeerTLS) GetStatus() StreamPeerTLSStatus {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret StreamPeerTLSStatus

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeerTLS.fnGetStatus), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *StreamPeerTLS) GetStream() StreamPeer {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStreamPeer()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeerTLS.fnGetStream), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *StreamPeerTLS) DisconnectFromStream()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeerTLS.fnDisconnectFromStream), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
