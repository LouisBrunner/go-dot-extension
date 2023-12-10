// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CylinderMesh struct {
  obj gdc.ObjectPtr
}

func (me *CylinderMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CylinderMesh) BaseClass() string {
  return "CylinderMesh"
}



// Enums

func (me *CylinderMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CylinderMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CylinderMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CylinderMesh) SetTopRadius(radius float32, )  {
  classNameV := StringNameFromStr("CylinderMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_top_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CylinderMesh) GetTopRadius() float32 {
  classNameV := StringNameFromStr("CylinderMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_top_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CylinderMesh) SetBottomRadius(radius float32, )  {
  classNameV := StringNameFromStr("CylinderMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bottom_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CylinderMesh) GetBottomRadius() float32 {
  classNameV := StringNameFromStr("CylinderMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bottom_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CylinderMesh) SetHeight(height float32, )  {
  classNameV := StringNameFromStr("CylinderMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CylinderMesh) GetHeight() float32 {
  classNameV := StringNameFromStr("CylinderMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CylinderMesh) SetRadialSegments(segments int, )  {
  classNameV := StringNameFromStr("CylinderMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_radial_segments")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&segments), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CylinderMesh) GetRadialSegments() int {
  classNameV := StringNameFromStr("CylinderMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_radial_segments")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CylinderMesh) SetRings(rings int, )  {
  classNameV := StringNameFromStr("CylinderMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rings")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rings), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CylinderMesh) GetRings() int {
  classNameV := StringNameFromStr("CylinderMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rings")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CylinderMesh) SetCapTop(cap_top bool, )  {
  classNameV := StringNameFromStr("CylinderMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cap_top")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cap_top), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CylinderMesh) IsCapTop() bool {
  classNameV := StringNameFromStr("CylinderMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_cap_top")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CylinderMesh) SetCapBottom(cap_bottom bool, )  {
  classNameV := StringNameFromStr("CylinderMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cap_bottom")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cap_bottom), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CylinderMesh) IsCapBottom() bool {
  classNameV := StringNameFromStr("CylinderMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_cap_bottom")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *CylinderMesh) GetPropTopRadius() float32 {
  panic("TODO: implement")
}

func (me *CylinderMesh) SetPropTopRadius(value float32) {
  panic("TODO: implement")
}

func (me *CylinderMesh) GetPropBottomRadius() float32 {
  panic("TODO: implement")
}

func (me *CylinderMesh) SetPropBottomRadius(value float32) {
  panic("TODO: implement")
}

func (me *CylinderMesh) GetPropHeight() float32 {
  panic("TODO: implement")
}

func (me *CylinderMesh) SetPropHeight(value float32) {
  panic("TODO: implement")
}

func (me *CylinderMesh) GetPropRadialSegments() int {
  panic("TODO: implement")
}

func (me *CylinderMesh) SetPropRadialSegments(value int) {
  panic("TODO: implement")
}

func (me *CylinderMesh) GetPropRings() int {
  panic("TODO: implement")
}

func (me *CylinderMesh) SetPropRings(value int) {
  panic("TODO: implement")
}

func (me *CylinderMesh) GetPropCapTop() bool {
  panic("TODO: implement")
}

func (me *CylinderMesh) SetPropCapTop(value bool) {
  panic("TODO: implement")
}

func (me *CylinderMesh) GetPropCapBottom() bool {
  panic("TODO: implement")
}

func (me *CylinderMesh) SetPropCapBottom(value bool) {
  panic("TODO: implement")
}