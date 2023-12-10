// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SphereMesh struct {
  obj gdc.ObjectPtr
}

func (me *SphereMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SphereMesh) BaseClass() string {
  return "SphereMesh"
}



// Enums

func (me *SphereMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SphereMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SphereMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SphereMesh) SetRadius(radius float32, )  {
  classNameV := StringNameFromStr("SphereMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SphereMesh) GetRadius() float32 {
  classNameV := StringNameFromStr("SphereMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SphereMesh) SetHeight(height float32, )  {
  classNameV := StringNameFromStr("SphereMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SphereMesh) GetHeight() float32 {
  classNameV := StringNameFromStr("SphereMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SphereMesh) SetRadialSegments(radial_segments int, )  {
  classNameV := StringNameFromStr("SphereMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_radial_segments")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radial_segments), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SphereMesh) GetRadialSegments() int {
  classNameV := StringNameFromStr("SphereMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_radial_segments")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SphereMesh) SetRings(rings int, )  {
  classNameV := StringNameFromStr("SphereMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rings")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rings), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SphereMesh) GetRings() int {
  classNameV := StringNameFromStr("SphereMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rings")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SphereMesh) SetIsHemisphere(is_hemisphere bool, )  {
  classNameV := StringNameFromStr("SphereMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_is_hemisphere")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&is_hemisphere), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SphereMesh) GetIsHemisphere() bool {
  classNameV := StringNameFromStr("SphereMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_is_hemisphere")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *SphereMesh) GetPropRadius() float32 {
  panic("TODO: implement")
}

func (me *SphereMesh) SetPropRadius(value float32) {
  panic("TODO: implement")
}

func (me *SphereMesh) GetPropHeight() float32 {
  panic("TODO: implement")
}

func (me *SphereMesh) SetPropHeight(value float32) {
  panic("TODO: implement")
}

func (me *SphereMesh) GetPropRadialSegments() int {
  panic("TODO: implement")
}

func (me *SphereMesh) SetPropRadialSegments(value int) {
  panic("TODO: implement")
}

func (me *SphereMesh) GetPropRings() int {
  panic("TODO: implement")
}

func (me *SphereMesh) SetPropRings(value int) {
  panic("TODO: implement")
}

func (me *SphereMesh) GetPropIsHemisphere() bool {
  panic("TODO: implement")
}

func (me *SphereMesh) SetPropIsHemisphere(value bool) {
  panic("TODO: implement")
}