// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

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
