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

type TubeTrailMesh struct {
  PrimitiveMesh
}

func (me *TubeTrailMesh) BaseClass() string {
  return "TubeTrailMesh"
}

func NewTubeTrailMesh() *TubeTrailMesh {
  str := StringNameFromStr("TubeTrailMesh") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &TubeTrailMesh{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *TubeTrailMesh) SetRadius(radius float64, )  {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TubeTrailMesh) GetRadius() float64 {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TubeTrailMesh) SetRadialSteps(radial_steps int64, )  {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_radial_steps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radial_steps) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TubeTrailMesh) GetRadialSteps() int64 {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_radial_steps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TubeTrailMesh) SetSections(sections int64, )  {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sections) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TubeTrailMesh) GetSections() int64 {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TubeTrailMesh) SetSectionLength(section_length float64, )  {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_section_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&section_length) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TubeTrailMesh) GetSectionLength() float64 {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_section_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TubeTrailMesh) SetSectionRings(section_rings int64, )  {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_section_rings")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&section_rings) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TubeTrailMesh) GetSectionRings() int64 {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_section_rings")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TubeTrailMesh) SetCapTop(cap_top bool, )  {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cap_top")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cap_top) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TubeTrailMesh) IsCapTop() bool {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_cap_top")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TubeTrailMesh) SetCapBottom(cap_bottom bool, )  {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cap_bottom")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cap_bottom) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TubeTrailMesh) IsCapBottom() bool {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_cap_bottom")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TubeTrailMesh) SetCurve(curve Curve, )  {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 270443179) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{curve.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TubeTrailMesh) GetCurve() Curve {
  classNameV := StringNameFromStr("TubeTrailMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2460114913) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCurve()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
