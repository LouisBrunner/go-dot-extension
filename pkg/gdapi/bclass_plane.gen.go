// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Plane struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Constants

var (
  PlanePlaneYz = "Plane(1, 0, 0, 0)" // TODO: construct correctly
  PlanePlaneXz = "Plane(0, 1, 0, 0)" // TODO: construct correctly
  PlanePlaneXy = "Plane(0, 0, 1, 0)" // TODO: construct correctly
)

// Enums

// Constructors

func NewPlane() Plane {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePlane))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePlane, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Plane{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPlaneFromPlane(from Plane, ) Plane {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePlane))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePlane, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Plane{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPlaneFromVector3(normal Vector3, ) Plane {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePlane))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePlane, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{normal.AsCTypePtr(), }))
  return Plane{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPlaneFromVector3Float32(normal Vector3, d float32, ) Plane {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePlane))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePlane, 3) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{normal.AsCTypePtr(), gdc.ConstTypePtr(&d), }))
  return Plane{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPlaneFromVector3Vector3(normal Vector3, point Vector3, ) Plane {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePlane))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePlane, 4) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{normal.AsCTypePtr(), point.AsCTypePtr(), }))
  return Plane{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPlaneFromVector3Vector3Vector3(point1 Vector3, point2 Vector3, point3 Vector3, ) Plane {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePlane))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePlane, 5) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{point1.AsCTypePtr(), point2.AsCTypePtr(), point3.AsCTypePtr(), }))
  return Plane{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPlaneFromFloat32Float32Float32Float32(a float32, b float32, c float32, d float32, ) Plane {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePlane))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePlane, 6) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&a), gdc.ConstTypePtr(&b), gdc.ConstTypePtr(&c), gdc.ConstTypePtr(&d), }))
  return Plane{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Plane) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Plane) Type() gdc.VariantType {
  return gdc.VariantTypePlane
}

func (me *Plane) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Plane) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func (me *Plane) Normalized() Plane {
  name := StringNameFromStr("normalized")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, name.AsCPtr(), 1051796340) // FIXME: should cache?

  var ret Plane
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Plane) GetCenter() Vector3 {
  name := StringNameFromStr("get_center")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, name.AsCPtr(), 1776574132) // FIXME: should cache?

  var ret Vector3
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Plane) IsEqualApprox(to_plane Plane, ) bool {
  name := StringNameFromStr("is_equal_approx")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, name.AsCPtr(), 1150170233) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{to_plane.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Plane) IsFinite() bool {
  name := StringNameFromStr("is_finite")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Plane) IsPointOver(point Vector3, ) bool {
  name := StringNameFromStr("is_point_over")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, name.AsCPtr(), 1749054343) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{point.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Plane) DistanceTo(point Vector3, ) float32 {
  name := StringNameFromStr("distance_to")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, name.AsCPtr(), 1047977935) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{point.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Plane) HasPoint(point Vector3, tolerance float32, ) bool {
  name := StringNameFromStr("has_point")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, name.AsCPtr(), 1258189072) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{point.AsCTypePtr(), gdc.ConstTypePtr(&tolerance), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Plane) Project(point Vector3, ) Vector3 {
  name := StringNameFromStr("project")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, name.AsCPtr(), 2923479887) // FIXME: should cache?

  var ret Vector3
  args := []gdc.ConstTypePtr{point.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Plane) Intersect3(b Plane, c Plane, ) Variant {
  name := StringNameFromStr("intersect_3")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, name.AsCPtr(), 2012052692) // FIXME: should cache?

  var ret Variant
  args := []gdc.ConstTypePtr{b.AsCTypePtr(), c.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Plane) IntersectsRay(from Vector3, dir Vector3, ) Variant {
  name := StringNameFromStr("intersects_ray")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, name.AsCPtr(), 2048133369) // FIXME: should cache?

  var ret Variant
  args := []gdc.ConstTypePtr{from.AsCTypePtr(), dir.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Plane) IntersectsSegment(from Vector3, to Vector3, ) Variant {
  name := StringNameFromStr("intersects_segment")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, name.AsCPtr(), 2048133369) // FIXME: should cache?

  var ret Variant
  args := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

// Operators

func (me *Plane) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Plane) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Plane) Negate() Plane {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNegate, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret Plane
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *Plane) Positive() Plane {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpPositive, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret Plane
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *Plane) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *Plane) EqualPlane(right Plane) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Plane) NotEqualPlane(right Plane) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Plane) MultiplyTransform3D(right Transform3D) Plane {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Plane
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Plane) InDictionary(right Dictionary) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Plane) InArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

// Members

func (me *Plane) X() float32 {
  name := StringNameFromStr("x")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret float32
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Plane) SetX(value float32) {
  name := StringNameFromStr("x")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}

func (me *Plane) Y() float32 {
  name := StringNameFromStr("y")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret float32
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Plane) SetY(value float32) {
  name := StringNameFromStr("y")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}

func (me *Plane) Z() float32 {
  name := StringNameFromStr("z")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret float32
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Plane) SetZ(value float32) {
  name := StringNameFromStr("z")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}

func (me *Plane) D() float32 {
  name := StringNameFromStr("d")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret float32
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Plane) SetD(value float32) {
  name := StringNameFromStr("d")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}

func (me *Plane) Normal() Vector3 {
  name := StringNameFromStr("normal")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret Vector3
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Plane) SetNormal(value Vector3) {
  name := StringNameFromStr("normal")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}
