// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VSlider struct {
  obj gdc.ObjectPtr
}

func (me *VSlider) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VSlider) BaseClass() string {
  return "VSlider"
}



// Enums

func (me *VSlider) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VSlider) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VSlider) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
