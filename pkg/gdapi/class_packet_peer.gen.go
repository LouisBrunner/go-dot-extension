// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func (me *PacketPeer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PacketPeer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PacketPeer) GetVar(allow_objects bool, )  {
  panic("TODO: implement")
}

func  (me *PacketPeer) PutVar(var_ Variant, full_objects bool, )  {
  panic("TODO: implement")
}

func  (me *PacketPeer) GetPacket()  {
  panic("TODO: implement")
}

func  (me *PacketPeer) PutPacket(buffer PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *PacketPeer) GetPacketError()  {
  panic("TODO: implement")
}

func  (me *PacketPeer) GetAvailablePacketCount()  {
  panic("TODO: implement")
}

func  (me *PacketPeer) GetEncodeBufferMaxSize()  {
  panic("TODO: implement")
}

func  (me *PacketPeer) SetEncodeBufferMaxSize(max_size int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
