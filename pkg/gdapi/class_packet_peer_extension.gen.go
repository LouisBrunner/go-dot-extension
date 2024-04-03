// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PacketPeerExtension struct {
  PacketPeer
}

func (me *PacketPeerExtension) BaseClass() string {
  return "PacketPeerExtension"
}

func NewPacketPeerExtension() *PacketPeerExtension {
  str := StringNameFromStr("PacketPeerExtension") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PacketPeerExtension{}
  obj.SetBaseObject(objPtr)
  return obj
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

// Signals
