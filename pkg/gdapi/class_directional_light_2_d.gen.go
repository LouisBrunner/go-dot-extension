// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type DirectionalLight2D struct {
  obj gdc.ObjectPtr
}

func (me *DirectionalLight2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *DirectionalLight2D) BaseClass() string {
  return "DirectionalLight2D"
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

func  (me *DirectionalLight2D) SetMaxDistance(pixels float32, )  {
  classNameV := StringNameFromStr("DirectionalLight2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pixels), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DirectionalLight2D) GetMaxDistance() float32 {
  classNameV := StringNameFromStr("DirectionalLight2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *DirectionalLight2D) GetPropHeight() float32 {
  panic("TODO: implement")
}

func (me *DirectionalLight2D) SetPropHeight(value float32) {
  panic("TODO: implement")
}

func (me *DirectionalLight2D) GetPropMaxDistance() float32 {
  panic("TODO: implement")
}

func (me *DirectionalLight2D) SetPropMaxDistance(value float32) {
  panic("TODO: implement")
}