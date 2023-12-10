// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ScriptExtension struct {
  obj gdc.ObjectPtr
}

func (me *ScriptExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ScriptExtension) BaseClass() string {
  return "ScriptExtension"
}



// Enums

func (me *ScriptExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ScriptExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ScriptExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties