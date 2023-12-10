// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeCustom struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeCustom) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeCustom) BaseClass() string {
  return "VisualShaderNodeCustom"
}



// Enums

func (me *VisualShaderNodeCustom) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeCustom) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeCustom) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties

func (me *VisualShaderNodeCustom) GetPropInitialized() bool {
  panic("TODO: implement")
}

func (me *VisualShaderNodeCustom) SetPropInitialized(value bool) {
  panic("TODO: implement")
}