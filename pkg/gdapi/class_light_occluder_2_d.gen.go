// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type LightOccluder2D struct {
  Node2D
}

func (me *LightOccluder2D) BaseClass() string {
  return "LightOccluder2D"
}

func NewLightOccluder2D() *LightOccluder2D {
  str := StringNameFromStr("LightOccluder2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &LightOccluder2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *LightOccluder2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *LightOccluder2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *LightOccluder2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *LightOccluder2D) SetOccluderPolygon(polygon OccluderPolygon2D, )  {
  classNameV := StringNameFromStr("LightOccluder2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_occluder_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3258315893) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(polygon.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightOccluder2D) GetOccluderPolygon() OccluderPolygon2D {
  classNameV := StringNameFromStr("LightOccluder2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_occluder_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3962317075) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewOccluderPolygon2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *LightOccluder2D) SetOccluderLightMask(mask int64, )  {
  classNameV := StringNameFromStr("LightOccluder2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_occluder_light_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightOccluder2D) GetOccluderLightMask() int64 {
  classNameV := StringNameFromStr("LightOccluder2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_occluder_light_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LightOccluder2D) SetAsSdfCollision(enable bool, )  {
  classNameV := StringNameFromStr("LightOccluder2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_as_sdf_collision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightOccluder2D) IsSetAsSdfCollision() bool {
  classNameV := StringNameFromStr("LightOccluder2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_set_as_sdf_collision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
