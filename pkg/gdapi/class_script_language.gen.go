// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ScriptLanguage struct {
  Object
}

func (me *ScriptLanguage) BaseClass() string {
  return "ScriptLanguage"
}



// Enums

func (me *ScriptLanguage) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ScriptLanguage) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ScriptLanguage) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
