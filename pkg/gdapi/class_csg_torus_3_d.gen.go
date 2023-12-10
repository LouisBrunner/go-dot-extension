// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CSGTorus3D struct {
  obj gdc.ObjectPtr
}

func (me *CSGTorus3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CSGTorus3D) BaseClass() string {
  return "CSGTorus3D"
}



// Enums

func (me *CSGTorus3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CSGTorus3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CSGTorus3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CSGTorus3D) SetInnerRadius(radius float32, )  {
  classNameV := StringNameFromStr("CSGTorus3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_inner_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CSGTorus3D) GetInnerRadius() float32 {
  classNameV := StringNameFromStr("CSGTorus3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_inner_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CSGTorus3D) SetOuterRadius(radius float32, )  {
  classNameV := StringNameFromStr("CSGTorus3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_outer_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CSGTorus3D) GetOuterRadius() float32 {
  classNameV := StringNameFromStr("CSGTorus3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_outer_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CSGTorus3D) SetSides(sides int, )  {
  classNameV := StringNameFromStr("CSGTorus3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sides")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sides), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CSGTorus3D) GetSides() int {
  classNameV := StringNameFromStr("CSGTorus3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sides")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CSGTorus3D) SetRingSides(sides int, )  {
  classNameV := StringNameFromStr("CSGTorus3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ring_sides")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sides), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CSGTorus3D) GetRingSides() int {
  classNameV := StringNameFromStr("CSGTorus3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ring_sides")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CSGTorus3D) SetMaterial(material Material, )  {
  classNameV := StringNameFromStr("CSGTorus3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2757459619) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(material.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CSGTorus3D) GetMaterial() Material {
  classNameV := StringNameFromStr("CSGTorus3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 5934680) // FIXME: should cache?
  var ret Material
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CSGTorus3D) SetSmoothFaces(smooth_faces bool, )  {
  classNameV := StringNameFromStr("CSGTorus3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_smooth_faces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&smooth_faces), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CSGTorus3D) GetSmoothFaces() bool {
  classNameV := StringNameFromStr("CSGTorus3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_smooth_faces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *CSGTorus3D) GetPropInnerRadius() float32 {
  panic("TODO: implement")
}

func (me *CSGTorus3D) SetPropInnerRadius(value float32) {
  panic("TODO: implement")
}

func (me *CSGTorus3D) GetPropOuterRadius() float32 {
  panic("TODO: implement")
}

func (me *CSGTorus3D) SetPropOuterRadius(value float32) {
  panic("TODO: implement")
}

func (me *CSGTorus3D) GetPropSides() int {
  panic("TODO: implement")
}

func (me *CSGTorus3D) SetPropSides(value int) {
  panic("TODO: implement")
}

func (me *CSGTorus3D) GetPropRingSides() int {
  panic("TODO: implement")
}

func (me *CSGTorus3D) SetPropRingSides(value int) {
  panic("TODO: implement")
}

func (me *CSGTorus3D) GetPropSmoothFaces() bool {
  panic("TODO: implement")
}

func (me *CSGTorus3D) SetPropSmoothFaces(value bool) {
  panic("TODO: implement")
}

func (me *CSGTorus3D) GetPropMaterial() any {
  panic("TODO: implement")
}

func (me *CSGTorus3D) SetPropMaterial(value any) {
  panic("TODO: implement")
}