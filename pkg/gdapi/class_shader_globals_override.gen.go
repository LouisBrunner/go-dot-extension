// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ShaderGlobalsOverride struct {
  obj gdc.ObjectPtr
}

func (me *ShaderGlobalsOverride) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ShaderGlobalsOverride) BaseClass() string {
  return "ShaderGlobalsOverride"
}



// Enums

func (me *ShaderGlobalsOverride) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ShaderGlobalsOverride) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ShaderGlobalsOverride) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties