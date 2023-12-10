// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type OfflineMultiplayerPeer struct {
  obj gdc.ObjectPtr
}

func (me *OfflineMultiplayerPeer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *OfflineMultiplayerPeer) BaseClass() string {
  return "OfflineMultiplayerPeer"
}



// Enums

func (me *OfflineMultiplayerPeer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *OfflineMultiplayerPeer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *OfflineMultiplayerPeer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties