// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type MultiplayerPeerExtension struct {
  obj gdc.ObjectPtr
}

func (me *MultiplayerPeerExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MultiplayerPeerExtension) BaseClass() string {
  return "MultiplayerPeerExtension"
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

// Properties