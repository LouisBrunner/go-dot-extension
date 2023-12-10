// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TubeTrailMesh struct {
  obj gdc.ObjectPtr
}

func (me *TubeTrailMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TubeTrailMesh) BaseClass() string {
  return "TubeTrailMesh"
}



// Enums

func (me *TubeTrailMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TubeTrailMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TubeTrailMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *TubeTrailMesh) SetRadius(radius float32, )  {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TubeTrailMesh) GetRadius() float32 {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TubeTrailMesh) SetRadialSteps(radial_steps int, )  {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_radial_steps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radial_steps), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TubeTrailMesh) GetRadialSteps() int {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_radial_steps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TubeTrailMesh) SetSections(sections int, )  {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sections), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TubeTrailMesh) GetSections() int {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TubeTrailMesh) SetSectionLength(section_length float32, )  {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_section_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&section_length), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TubeTrailMesh) GetSectionLength() float32 {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_section_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TubeTrailMesh) SetSectionRings(section_rings int, )  {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_section_rings")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&section_rings), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TubeTrailMesh) GetSectionRings() int {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_section_rings")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TubeTrailMesh) SetCapTop(cap_top bool, )  {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cap_top")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cap_top), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TubeTrailMesh) IsCapTop() bool {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_cap_top")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TubeTrailMesh) SetCapBottom(cap_bottom bool, )  {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cap_bottom")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cap_bottom), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TubeTrailMesh) IsCapBottom() bool {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_cap_bottom")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TubeTrailMesh) SetCurve(curve Curve, )  {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 270443179) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(curve.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TubeTrailMesh) GetCurve() Curve {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2460114913) // FIXME: should cache?
  var ret Curve
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *TubeTrailMesh) GetPropRadius() float32 {
  panic("TODO: implement")
}

func (me *TubeTrailMesh) SetPropRadius(value float32) {
  panic("TODO: implement")
}

func (me *TubeTrailMesh) GetPropRadialSteps() int {
  panic("TODO: implement")
}

func (me *TubeTrailMesh) SetPropRadialSteps(value int) {
  panic("TODO: implement")
}

func (me *TubeTrailMesh) GetPropSections() int {
  panic("TODO: implement")
}

func (me *TubeTrailMesh) SetPropSections(value int) {
  panic("TODO: implement")
}

func (me *TubeTrailMesh) GetPropSectionLength() float32 {
  panic("TODO: implement")
}

func (me *TubeTrailMesh) SetPropSectionLength(value float32) {
  panic("TODO: implement")
}

func (me *TubeTrailMesh) GetPropSectionRings() int {
  panic("TODO: implement")
}

func (me *TubeTrailMesh) SetPropSectionRings(value int) {
  panic("TODO: implement")
}

func (me *TubeTrailMesh) GetPropCapTop() bool {
  panic("TODO: implement")
}

func (me *TubeTrailMesh) SetPropCapTop(value bool) {
  panic("TODO: implement")
}

func (me *TubeTrailMesh) GetPropCapBottom() bool {
  panic("TODO: implement")
}

func (me *TubeTrailMesh) SetPropCapBottom(value bool) {
  panic("TODO: implement")
}

func (me *TubeTrailMesh) GetPropCurve() Curve {
  panic("TODO: implement")
}

func (me *TubeTrailMesh) SetPropCurve(value Curve) {
  panic("TODO: implement")
}