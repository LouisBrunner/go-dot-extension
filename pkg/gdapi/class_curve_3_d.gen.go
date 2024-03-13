// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Curve3D struct {
  obj gdc.ObjectPtr
}

func (me *Curve3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Curve3D) BaseClass() string {
  return "Curve3D"
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

func  (me *Curve3D) GetPointCount() int {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Curve3D) SetPointCount(count int, )  {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_point_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Curve3D) AddPoint(position Vector3, in Vector3, out Vector3, index int, )  {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2931053748) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), gdc.ConstTypePtr(in.AsCTypePtr()), gdc.ConstTypePtr(out.AsCTypePtr()), gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Curve3D) SetPointPosition(idx int, position Vector3, )  {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_point_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1530502735) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Curve3D) GetPointPosition(idx int, ) Vector3 {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 711720468) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Curve3D) SetPointTilt(idx int, tilt float32, )  {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_point_tilt")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&tilt), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Curve3D) GetPointTilt(idx int, ) float32 {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_tilt")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Curve3D) SetPointIn(idx int, position Vector3, )  {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_point_in")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1530502735) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Curve3D) GetPointIn(idx int, ) Vector3 {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_in")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 711720468) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Curve3D) SetPointOut(idx int, position Vector3, )  {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_point_out")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1530502735) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Curve3D) GetPointOut(idx int, ) Vector3 {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_out")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 711720468) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Curve3D) RemovePoint(idx int, )  {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Curve3D) ClearPoints()  {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Curve3D) Sample(idx int, t float32, ) Vector3 {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("sample")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3285246857) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&t), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Curve3D) Samplef(fofs float32, ) Vector3 {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("samplef")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2553580215) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fofs), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Curve3D) SetBakeInterval(distance float32, )  {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bake_interval")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Curve3D) GetBakeInterval() float32 {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bake_interval")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Curve3D) SetUpVectorEnabled(enable bool, )  {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_up_vector_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Curve3D) IsUpVectorEnabled() bool {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_up_vector_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Curve3D) GetBakedLength() float32 {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_baked_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Curve3D) SampleBaked(offset float32, cubic bool, ) Vector3 {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("sample_baked")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1350085894) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset), gdc.ConstTypePtr(&cubic), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Curve3D) SampleBakedWithRotation(offset float32, cubic bool, apply_tilt bool, ) Transform3D {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("sample_baked_with_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1939359131) // FIXME: should cache?
  var ret Transform3D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset), gdc.ConstTypePtr(&cubic), gdc.ConstTypePtr(&apply_tilt), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Curve3D) SampleBakedUpVector(offset float32, apply_tilt bool, ) Vector3 {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("sample_baked_up_vector")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1362627031) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset), gdc.ConstTypePtr(&apply_tilt), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Curve3D) GetBakedPoints() PackedVector3Array {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_baked_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 497664490) // FIXME: should cache?
  var ret PackedVector3Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Curve3D) GetBakedTilts() PackedFloat32Array {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_baked_tilts")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 675695659) // FIXME: should cache?
  var ret PackedFloat32Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Curve3D) GetBakedUpVectors() PackedVector3Array {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_baked_up_vectors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 497664490) // FIXME: should cache?
  var ret PackedVector3Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Curve3D) GetClosestPoint(to_point Vector3, ) Vector3 {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_closest_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 192990374) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(to_point.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Curve3D) GetClosestOffset(to_point Vector3, ) float32 {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_closest_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1109078154) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(to_point.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Curve3D) Tessellate(max_stages int, tolerance_degrees float32, ) PackedVector3Array {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tessellate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1519759391) // FIXME: should cache?
  var ret PackedVector3Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_stages), gdc.ConstTypePtr(&tolerance_degrees), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Curve3D) TessellateEvenLength(max_stages int, tolerance_length float32, ) PackedVector3Array {
  classNameV := StringNameFromStr("Curve3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tessellate_even_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 133237049) // FIXME: should cache?
  var ret PackedVector3Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_stages), gdc.ConstTypePtr(&tolerance_length), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
