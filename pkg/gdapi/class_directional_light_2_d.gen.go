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

type DirectionalLight2D struct {
  Light2D
}

func (me *DirectionalLight2D) BaseClass() string {
  return "DirectionalLight2D"
}

func NewDirectionalLight2D() *DirectionalLight2D {
  str := StringNameFromStr("DirectionalLight2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &DirectionalLight2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *DirectionalLight2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *DirectionalLight2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *DirectionalLight2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *DirectionalLight2D) SetMaxDistance(pixels float64, )  {
  classNameV := StringNameFromStr("DirectionalLight2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pixels) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *DirectionalLight2D) GetMaxDistance() float64 {
  classNameV := StringNameFromStr("DirectionalLight2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
