// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ShaderGlobalsOverride struct {
  Node
}

func (me *ShaderGlobalsOverride) BaseClass() string {
  return "ShaderGlobalsOverride"
}

func NewShaderGlobalsOverride() *ShaderGlobalsOverride {
  str := StringNameFromStr("ShaderGlobalsOverride") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ShaderGlobalsOverride{}
  obj.SetBaseObject(objPtr)
  return obj
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

// Signals
