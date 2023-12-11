// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type StreamPeerTLS struct {
  obj gdc.ObjectPtr
}

func (me *StreamPeerTLS) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *StreamPeerTLS) BaseClass() string {
  return "StreamPeerTLS"
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
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StreamPeerTLS) AcceptStream(stream StreamPeer, server_options TLSOptions, ) Error {
  classNameV := StringNameFromStr("StreamPeerTLS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("accept_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4292689651) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(stream.AsCTypePtr()), gdc.ConstTypePtr(server_options.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StreamPeerTLS) ConnectToStream(stream StreamPeer, common_name String, client_options TLSOptions, ) Error {
  classNameV := StringNameFromStr("StreamPeerTLS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("connect_to_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1325480781) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(stream.AsCTypePtr()), gdc.ConstTypePtr(common_name.AsCTypePtr()), gdc.ConstTypePtr(client_options.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StreamPeerTLS) GetStatus() StreamPeerTLSStatus {
  classNameV := StringNameFromStr("StreamPeerTLS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_status")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1128380576) // FIXME: should cache?
  var ret StreamPeerTLSStatus
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StreamPeerTLS) GetStream() StreamPeer {
  classNameV := StringNameFromStr("StreamPeerTLS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2741655269) // FIXME: should cache?
  var ret StreamPeer
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StreamPeerTLS) DisconnectFromStream()  {
  classNameV := StringNameFromStr("StreamPeerTLS")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("disconnect_from_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Signals
