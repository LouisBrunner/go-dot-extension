// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type StreamPeerTCP struct {
  StreamPeer
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

func  (me *StreamPeerTCP) Bind(port int, host String, ) Error {
  classNameV := StringNameFromStr("StreamPeerTCP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("bind")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3167955072) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port), gdc.ConstTypePtr(host.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StreamPeerTCP) ConnectToHost(host String, port int, ) Error {
  classNameV := StringNameFromStr("StreamPeerTCP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("connect_to_host")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 993915709) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(host.AsCTypePtr()), gdc.ConstTypePtr(&port), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StreamPeerTCP) Poll() Error {
  classNameV := StringNameFromStr("StreamPeerTCP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("poll")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166280745) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StreamPeerTCP) GetStatus() StreamPeerTCPStatus {
  classNameV := StringNameFromStr("StreamPeerTCP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_status")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 859471121) // FIXME: should cache?
  var ret StreamPeerTCPStatus
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StreamPeerTCP) GetConnectedHost() String {
  classNameV := StringNameFromStr("StreamPeerTCP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connected_host")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StreamPeerTCP) GetConnectedPort() int {
  classNameV := StringNameFromStr("StreamPeerTCP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connected_port")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StreamPeerTCP) GetLocalPort() int {
  classNameV := StringNameFromStr("StreamPeerTCP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_local_port")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StreamPeerTCP) DisconnectFromHost()  {
  classNameV := StringNameFromStr("StreamPeerTCP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("disconnect_from_host")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StreamPeerTCP) SetNoDelay(enabled bool, )  {
  classNameV := StringNameFromStr("StreamPeerTCP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_no_delay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Signals
