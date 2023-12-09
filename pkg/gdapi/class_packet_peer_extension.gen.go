// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PacketPeerExtension struct {
  obj gdc.ObjectPtr
}

func (me *PacketPeerExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PacketPeerExtension) BaseClass() string {
  return "PacketPeerExtension"
}



// Enums

func (me *PacketPeerExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PacketPeerExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PacketPeerExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PacketPeerExtension) XGetPacket(r_buffer **uint8, r_buffer_size *int32, )  {
  panic("TODO: implement")
}

func  (me *PacketPeerExtension) XPutPacket(p_buffer *uint8, p_buffer_size int, )  {
  panic("TODO: implement")
}

func  (me *PacketPeerExtension) XGetAvailablePacketCount()  {
  panic("TODO: implement")
}

func  (me *PacketPeerExtension) XGetMaxPacketSize()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
