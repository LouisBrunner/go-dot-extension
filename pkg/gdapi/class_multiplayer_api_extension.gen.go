// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type MultiplayerAPIExtension struct {
  obj gdc.ObjectPtr
}

func (me *MultiplayerAPIExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MultiplayerAPIExtension) BaseClass() string {
  return "MultiplayerAPIExtension"
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
