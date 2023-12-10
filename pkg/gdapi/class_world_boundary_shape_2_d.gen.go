// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type WorldBoundaryShape2D struct {
  obj gdc.ObjectPtr
}

func (me *WorldBoundaryShape2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *WorldBoundaryShape2D) BaseClass() string {
  return "WorldBoundaryShape2D"
}



// Enums

func (me *WorldBoundaryShape2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *WorldBoundaryShape2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *WorldBoundaryShape2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *WorldBoundaryShape2D) SetNormal(normal Vector2, )  {
  classNameV := StringNameFromStr("WorldBoundaryShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(normal.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *WorldBoundaryShape2D) GetNormal() Vector2 {
  classNameV := StringNameFromStr("WorldBoundaryShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WorldBoundaryShape2D) SetDistance(distance float32, )  {
  classNameV := StringNameFromStr("WorldBoundaryShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *WorldBoundaryShape2D) GetDistance() float32 {
  classNameV := StringNameFromStr("WorldBoundaryShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *WorldBoundaryShape2D) GetPropNormal() Vector2 {
  panic("TODO: implement")
}

func (me *WorldBoundaryShape2D) SetPropNormal(value Vector2) {
  panic("TODO: implement")
}

func (me *WorldBoundaryShape2D) GetPropDistance() float32 {
  panic("TODO: implement")
}

func (me *WorldBoundaryShape2D) SetPropDistance(value float32) {
  panic("TODO: implement")
}