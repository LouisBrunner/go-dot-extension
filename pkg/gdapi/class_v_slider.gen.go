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

type VSlider struct {
  Slider
}

func (me *VSlider) BaseClass() string {
  return "VSlider"
}

func NewVSlider() *VSlider {
  str := StringNameFromStr("VSlider") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VSlider{}
  obj.SetBaseObject(objPtr)
  return obj
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
