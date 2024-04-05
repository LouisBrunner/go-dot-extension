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

type ptrsForPacketPeerExtensionList struct {
  fnXGetPacket gdc.MethodBindPtr
  fnXPutPacket gdc.MethodBindPtr
  fnXGetAvailablePacketCount gdc.MethodBindPtr
  fnXGetMaxPacketSize gdc.MethodBindPtr
}

var ptrsForPacketPeerExtension ptrsForPacketPeerExtensionList

func initPacketPeerExtensionPtrs(iface gdc.Interface) {

  className := StringNameFromStr("PacketPeerExtension")
  defer className.Destroy()
}

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
