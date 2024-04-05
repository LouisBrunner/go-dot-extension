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

type ptrsForHSliderList struct {
}

var ptrsForHSlider ptrsForHSliderList

func initHSliderPtrs(iface gdc.Interface) {

  className := StringNameFromStr("HSlider")
  defer className.Destroy()
}

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
