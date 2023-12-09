// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PacketPeerStream struct {
  obj gdc.ObjectPtr
}

func (me *PacketPeerStream) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PacketPeerStream) BaseClass() string {
  return "PacketPeerStream"
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
  panic("TODO: implement")
}

func  (me *PacketPeerStream) GetStreamPeer()  {
  panic("TODO: implement")
}

func  (me *PacketPeerStream) SetInputBufferMaxSize(max_size_bytes int, )  {
  panic("TODO: implement")
}

func  (me *PacketPeerStream) SetOutputBufferMaxSize(max_size_bytes int, )  {
  panic("TODO: implement")
}

func  (me *PacketPeerStream) GetInputBufferMaxSize()  {
  panic("TODO: implement")
}

func  (me *PacketPeerStream) GetOutputBufferMaxSize()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
