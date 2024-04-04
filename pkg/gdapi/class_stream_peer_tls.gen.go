// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

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
  classNameV := StringNameFromStr("StreamPeerTLS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("poll")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeerTLS) AcceptStream(stream StreamPeer, server_options TLSOptions, ) Error {
  classNameV := StringNameFromStr("StreamPeerTLS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("accept_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4292689651) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{stream.AsCTypePtr(), server_options.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *StreamPeerTLS) ConnectToStream(stream StreamPeer, common_name String, client_options TLSOptions, ) Error {
  classNameV := StringNameFromStr("StreamPeerTLS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("connect_to_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 57169517) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{stream.AsCTypePtr(), common_name.AsCTypePtr(), client_options.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *StreamPeerTLS) GetStatus() StreamPeerTLSStatus {
  classNameV := StringNameFromStr("StreamPeerTLS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_status")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1128380576) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret StreamPeerTLSStatus

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *StreamPeerTLS) GetStream() StreamPeer {
  classNameV := StringNameFromStr("StreamPeerTLS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2741655269) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStreamPeer()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *StreamPeerTLS) DisconnectFromStream()  {
  classNameV := StringNameFromStr("StreamPeerTLS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("disconnect_from_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
