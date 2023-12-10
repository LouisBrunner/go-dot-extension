// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type LightOccluder2D struct {
  obj gdc.ObjectPtr
}

func (me *LightOccluder2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *LightOccluder2D) BaseClass() string {
  return "LightOccluder2D"
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
  var ret OccluderPolygon2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LightOccluder2D) SetOccluderLightMask(mask int, )  {
  classNameV := StringNameFromStr("LightOccluder2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_occluder_light_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LightOccluder2D) GetOccluderLightMask() int {
  classNameV := StringNameFromStr("LightOccluder2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_occluder_light_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *LightOccluder2D) GetPropOccluder() OccluderPolygon2D {
  panic("TODO: implement")
}

func (me *LightOccluder2D) SetPropOccluder(value OccluderPolygon2D) {
  panic("TODO: implement")
}

func (me *LightOccluder2D) GetPropSdfCollision() bool {
  panic("TODO: implement")
}

func (me *LightOccluder2D) SetPropSdfCollision(value bool) {
  panic("TODO: implement")
}

func (me *LightOccluder2D) GetPropOccluderLightMask() int {
  panic("TODO: implement")
}

func (me *LightOccluder2D) SetPropOccluderLightMask(value int) {
  panic("TODO: implement")
}