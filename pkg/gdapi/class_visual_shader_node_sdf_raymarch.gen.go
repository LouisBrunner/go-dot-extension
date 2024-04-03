// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeSDFRaymarch struct {
  VisualShaderNode
}

func (me *VisualShaderNodeSDFRaymarch) BaseClass() string {
  return "VisualShaderNodeSDFRaymarch"
}

func NewVisualShaderNodeSDFRaymarch() *VisualShaderNodeSDFRaymarch {
  str := StringNameFromStr("VisualShaderNodeSDFRaymarch") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeSDFRaymarch{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeSDFRaymarch) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeSDFRaymarch) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeSDFRaymarch) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
