// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type HSlider struct {
  Slider
}

func (me *HSlider) BaseClass() string {
  return "HSlider"
}

func NewHSlider() *HSlider {
  str := StringNameFromStr("HSlider") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &HSlider{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *HSlider) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *HSlider) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *HSlider) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
