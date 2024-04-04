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

type Curve3D struct {
  Resource
}

func (me *Curve3D) BaseClass() string {
  return "Curve3D"
}

func NewCurve3D() *Curve3D {
  str := StringNameFromStr("Curve3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Curve3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *Curve3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Curve3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Curve3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Curve3D) GetPointCount() int64 {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Curve3D) SetPointCount(count int64, )  {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_point_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve3D) AddPoint(position Vector3, in Vector3, out Vector3, index int64, )  {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2931053748) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), in.AsCTypePtr(), out.AsCTypePtr(), gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve3D) SetPointPosition(idx int64, position Vector3, )  {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_point_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1530502735) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve3D) GetPointPosition(idx int64, ) Vector3 {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 711720468) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Curve3D) SetPointTilt(idx int64, tilt float64, )  {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_point_tilt")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , gdc.ConstTypePtr(&tilt) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve3D) GetPointTilt(idx int64, ) float64 {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_tilt")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Curve3D) SetPointIn(idx int64, position Vector3, )  {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_point_in")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1530502735) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve3D) GetPointIn(idx int64, ) Vector3 {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_in")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 711720468) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Curve3D) SetPointOut(idx int64, position Vector3, )  {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_point_out")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1530502735) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve3D) GetPointOut(idx int64, ) Vector3 {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_out")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 711720468) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Curve3D) RemovePoint(idx int64, )  {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve3D) ClearPoints()  {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve3D) Sample(idx int64, t float64, ) Vector3 {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("sample")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3285246857) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , gdc.ConstTypePtr(&t) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&idx)
  pinner.Pin(&t)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Curve3D) Samplef(fofs float64, ) Vector3 {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("samplef")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2553580215) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fofs) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&fofs)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Curve3D) SetBakeInterval(distance float64, )  {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bake_interval")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve3D) GetBakeInterval() float64 {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bake_interval")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Curve3D) SetUpVectorEnabled(enable bool, )  {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_up_vector_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Curve3D) IsUpVectorEnabled() bool {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_up_vector_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Curve3D) GetBakedLength() float64 {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_baked_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Curve3D) SampleBaked(offset float64, cubic bool, ) Vector3 {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("sample_baked")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1350085894) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset) , gdc.ConstTypePtr(&cubic) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&offset)
  pinner.Pin(&cubic)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Curve3D) SampleBakedWithRotation(offset float64, cubic bool, apply_tilt bool, ) Transform3D {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("sample_baked_with_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1939359131) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset) , gdc.ConstTypePtr(&cubic) , gdc.ConstTypePtr(&apply_tilt) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()
  pinner.Pin(&offset)
  pinner.Pin(&cubic)
  pinner.Pin(&apply_tilt)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Curve3D) SampleBakedUpVector(offset float64, apply_tilt bool, ) Vector3 {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("sample_baked_up_vector")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1362627031) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset) , gdc.ConstTypePtr(&apply_tilt) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&offset)
  pinner.Pin(&apply_tilt)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Curve3D) GetBakedPoints() PackedVector3Array {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_baked_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 497664490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector3Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Curve3D) GetBakedTilts() PackedFloat32Array {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_baked_tilts")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 675695659) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedFloat32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Curve3D) GetBakedUpVectors() PackedVector3Array {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_baked_up_vectors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 497664490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector3Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Curve3D) GetClosestPoint(to_point Vector3, ) Vector3 {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_closest_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 192990374) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{to_point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Curve3D) GetClosestOffset(to_point Vector3, ) float64 {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_closest_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1109078154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{to_point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Curve3D) Tessellate(max_stages int64, tolerance_degrees float64, ) PackedVector3Array {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tessellate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1519759391) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_stages) , gdc.ConstTypePtr(&tolerance_degrees) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector3Array()
  pinner.Pin(&max_stages)
  pinner.Pin(&tolerance_degrees)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Curve3D) TessellateEvenLength(max_stages int64, tolerance_length float64, ) PackedVector3Array {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tessellate_even_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 133237049) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_stages) , gdc.ConstTypePtr(&tolerance_length) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector3Array()
  pinner.Pin(&max_stages)
  pinner.Pin(&tolerance_length)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
