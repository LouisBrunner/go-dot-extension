// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
