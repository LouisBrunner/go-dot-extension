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

type ptrsForPacketPeerList struct {
  fnGetVar gdc.MethodBindPtr
  fnPutVar gdc.MethodBindPtr
  fnGetPacket gdc.MethodBindPtr
  fnPutPacket gdc.MethodBindPtr
  fnGetPacketError gdc.MethodBindPtr
  fnGetAvailablePacketCount gdc.MethodBindPtr
  fnGetEncodeBufferMaxSize gdc.MethodBindPtr
  fnSetEncodeBufferMaxSize gdc.MethodBindPtr
}

var ptrsForPacketPeer ptrsForPacketPeerList

func initPacketPeerPtrs(iface gdc.Interface) {

  className := StringNameFromStr("PacketPeer")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_var")
    defer methodName.Destroy()
    ptrsForPacketPeer.fnGetVar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3442865206))
  }
  {
    methodName := StringNameFromStr("put_var")
    defer methodName.Destroy()
    ptrsForPacketPeer.fnPutVar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2436251611))
  }
  {
    methodName := StringNameFromStr("get_packet")
    defer methodName.Destroy()
    ptrsForPacketPeer.fnGetPacket = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2115431945))
  }
  {
    methodName := StringNameFromStr("put_packet")
    defer methodName.Destroy()
    ptrsForPacketPeer.fnPutPacket = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 680677267))
  }
  {
    methodName := StringNameFromStr("get_packet_error")
    defer methodName.Destroy()
    ptrsForPacketPeer.fnGetPacketError = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3185525595))
  }
  {
    methodName := StringNameFromStr("get_available_packet_count")
    defer methodName.Destroy()
    ptrsForPacketPeer.fnGetAvailablePacketCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_encode_buffer_max_size")
    defer methodName.Destroy()
    ptrsForPacketPeer.fnGetEncodeBufferMaxSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_encode_buffer_max_size")
    defer methodName.Destroy()
    ptrsForPacketPeer.fnSetEncodeBufferMaxSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
}

type PacketPeer struct {
  RefCounted
}

func (me *PacketPeer) BaseClass() string {
  return "PacketPeer"
}

func NewPacketPeer() *PacketPeer {
  str := StringNameFromStr("PacketPeer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PacketPeer{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow_objects) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()
  pinner.Pin(&allow_objects)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeer.fnGetVar), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PacketPeer) PutVar(var_ Variant, full_objects bool, ) Error {
  cargs := []gdc.ConstTypePtr{var_.AsCTypePtr(), gdc.ConstTypePtr(&full_objects) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&full_objects)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeer.fnPutVar), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *PacketPeer) GetPacket() PackedByteArray {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedByteArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeer.fnGetPacket), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PacketPeer) PutPacket(buffer PackedByteArray, ) Error {
  cargs := []gdc.ConstTypePtr{buffer.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeer.fnPutPacket), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *PacketPeer) GetPacketError() Error {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeer.fnGetPacketError), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *PacketPeer) GetAvailablePacketCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeer.fnGetAvailablePacketCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PacketPeer) GetEncodeBufferMaxSize() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeer.fnGetEncodeBufferMaxSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PacketPeer) SetEncodeBufferMaxSize(max_size int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeer.fnSetEncodeBufferMaxSize), me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
