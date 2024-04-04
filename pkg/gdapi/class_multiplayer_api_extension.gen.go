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

type MultiplayerAPIExtension struct {
  MultiplayerAPI
}

func (me *MultiplayerAPIExtension) BaseClass() string {
  return "MultiplayerAPIExtension"
}

func NewMultiplayerAPIExtension() *MultiplayerAPIExtension {
  str := StringNameFromStr("MultiplayerAPIExtension") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &MultiplayerAPIExtension{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *MultiplayerAPIExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MultiplayerAPIExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MultiplayerAPIExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
