// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PacketPeerStream struct {
  PacketPeer
}

func (me *PacketPeerStream) BaseClass() string {
  return "PacketPeerStream"
}

func NewPacketPeerStream() *PacketPeerStream {
  str := StringNameFromStr("PacketPeerStream") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PacketPeerStream{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *PacketPeerStream) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PacketPeerStream) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PacketPeerStream) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PacketPeerStream) SetStreamPeer(peer StreamPeer, )  {
  classNameV := StringNameFromStr("PacketPeerStream")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stream_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3281897016) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(peer.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PacketPeerStream) GetStreamPeer() StreamPeer {
  classNameV := StringNameFromStr("PacketPeerStream")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stream_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2741655269) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewStreamPeer()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PacketPeerStream) SetInputBufferMaxSize(max_size_bytes int64, )  {
  classNameV := StringNameFromStr("PacketPeerStream")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_input_buffer_max_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_size_bytes), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PacketPeerStream) SetOutputBufferMaxSize(max_size_bytes int64, )  {
  classNameV := StringNameFromStr("PacketPeerStream")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_output_buffer_max_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_size_bytes), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PacketPeerStream) GetInputBufferMaxSize() int64 {
  classNameV := StringNameFromStr("PacketPeerStream")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_input_buffer_max_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PacketPeerStream) GetOutputBufferMaxSize() int64 {
  classNameV := StringNameFromStr("PacketPeerStream")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_output_buffer_max_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
