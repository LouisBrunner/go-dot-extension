// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PacketPeer struct {
  obj gdc.ObjectPtr
}

func (me *PacketPeer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PacketPeer) BaseClass() string {
  return "PacketPeer"
}



// Enums

func (me *PacketPeer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PacketPeer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PacketPeer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PacketPeer) GetVar(allow_objects bool, ) Variant {
  classNameV := StringNameFromStr("PacketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_var")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3442865206) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow_objects), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PacketPeer) PutVar(var_ Variant, full_objects bool, ) Error {
  classNameV := StringNameFromStr("PacketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("put_var")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2436251611) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(var_.AsCTypePtr()), gdc.ConstTypePtr(&full_objects), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PacketPeer) GetPacket() PackedByteArray {
  classNameV := StringNameFromStr("PacketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_packet")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2115431945) // FIXME: should cache?
  var ret PackedByteArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PacketPeer) PutPacket(buffer PackedByteArray, ) Error {
  classNameV := StringNameFromStr("PacketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("put_packet")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 680677267) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(buffer.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PacketPeer) GetPacketError() Error {
  classNameV := StringNameFromStr("PacketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_packet_error")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3185525595) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PacketPeer) GetAvailablePacketCount() int {
  classNameV := StringNameFromStr("PacketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_available_packet_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PacketPeer) GetEncodeBufferMaxSize() int {
  classNameV := StringNameFromStr("PacketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_encode_buffer_max_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PacketPeer) SetEncodeBufferMaxSize(max_size int, )  {
  classNameV := StringNameFromStr("PacketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_encode_buffer_max_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
