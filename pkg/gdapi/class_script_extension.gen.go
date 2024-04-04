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

type ScriptExtension struct {
  Script
}

func (me *ScriptExtension) BaseClass() string {
  return "ScriptExtension"
}

func NewScriptExtension() *ScriptExtension {
  str := StringNameFromStr("ScriptExtension") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ScriptExtension{}
  obj.SetBaseObject(objPtr)
  return obj
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

// Signals
