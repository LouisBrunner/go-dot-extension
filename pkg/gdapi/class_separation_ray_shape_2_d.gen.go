// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SeparationRayShape2D struct {
  obj gdc.ObjectPtr
}

func (me *SeparationRayShape2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SeparationRayShape2D) BaseClass() string {
  return "SeparationRayShape2D"
}



// Enums

func (me *SeparationRayShape2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SeparationRayShape2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SeparationRayShape2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SeparationRayShape2D) SetLength(length float32, )  {
  classNameV := StringNameFromStr("SeparationRayShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SeparationRayShape2D) GetLength() float32 {
  classNameV := StringNameFromStr("SeparationRayShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SeparationRayShape2D) SetSlideOnSlope(active bool, )  {
  classNameV := StringNameFromStr("SeparationRayShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_slide_on_slope")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SeparationRayShape2D) GetSlideOnSlope() bool {
  classNameV := StringNameFromStr("SeparationRayShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_slide_on_slope")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
