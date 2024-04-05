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

type ptrsForMultiplayerPeerExtensionList struct {
  fnXGetPacket gdc.MethodBindPtr
  fnXPutPacket gdc.MethodBindPtr
  fnXGetAvailablePacketCount gdc.MethodBindPtr
  fnXGetMaxPacketSize gdc.MethodBindPtr
  fnXGetPacketScript gdc.MethodBindPtr
  fnXPutPacketScript gdc.MethodBindPtr
  fnXGetPacketChannel gdc.MethodBindPtr
  fnXGetPacketMode gdc.MethodBindPtr
  fnXSetTransferChannel gdc.MethodBindPtr
  fnXGetTransferChannel gdc.MethodBindPtr
  fnXSetTransferMode gdc.MethodBindPtr
  fnXGetTransferMode gdc.MethodBindPtr
  fnXSetTargetPeer gdc.MethodBindPtr
  fnXGetPacketPeer gdc.MethodBindPtr
  fnXIsServer gdc.MethodBindPtr
  fnXPoll gdc.MethodBindPtr
  fnXClose gdc.MethodBindPtr
  fnXDisconnectPeer gdc.MethodBindPtr
  fnXGetUniqueId gdc.MethodBindPtr
  fnXSetRefuseNewConnections gdc.MethodBindPtr
  fnXIsRefusingNewConnections gdc.MethodBindPtr
  fnXIsServerRelaySupported gdc.MethodBindPtr
  fnXGetConnectionStatus gdc.MethodBindPtr
}

var ptrsForMultiplayerPeerExtension ptrsForMultiplayerPeerExtensionList

func initMultiplayerPeerExtensionPtrs(iface gdc.Interface) {

  className := StringNameFromStr("MultiplayerPeerExtension")
  defer className.Destroy()
}

type MultiplayerPeerExtension struct {
  MultiplayerPeer
}

func (me *MultiplayerPeerExtension) BaseClass() string {
  return "MultiplayerPeerExtension"
}

func NewMultiplayerPeerExtension() *MultiplayerPeerExtension {
  str := StringNameFromStr("MultiplayerPeerExtension") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &MultiplayerPeerExtension{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *MultiplayerPeerExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MultiplayerPeerExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MultiplayerPeerExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
