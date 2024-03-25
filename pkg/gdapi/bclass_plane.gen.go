// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "fmt"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Plane struct {
  data   *[classSizePlane]byte
  iface  gdc.Interface
  pinner runtime.Pinner
}

// Constants

var (
  PlanePlaneYz = "Plane(1, 0, 0, 0)" // TODO: construct correctly
  PlanePlaneXz = "Plane(0, 1, 0, 0)" // TODO: construct correctly
  PlanePlaneXy = "Plane(0, 0, 1, 0)" // TODO: construct correctly
)

// Enums

// Constructors
func newPlane() *Plane {
  me := &Plane{
    data:   new([classSizePlane]byte),
    iface:  giface,
  }
  me.pinner.Pin(me)
  me.pinner.Pin(me.AsTypePtr())
  return me
}

func NewPlane() *Plane {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPlane()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypePlane, 0) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewPlaneFromPlane(from Plane, ) *Plane {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPlane()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypePlane, 1) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewPlaneFromVector3(normal Vector3, ) *Plane {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPlane()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypePlane, 2) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{normal.AsCTypePtr(), }))
  return me
}

func NewPlaneFromVector3Float32(normal Vector3, d float64, ) *Plane {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&d)
  me := newPlane()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypePlane, 3) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{normal.AsCTypePtr(), gdc.ConstTypePtr(&d), }))
  return me
}

func NewPlaneFromVector3Vector3(normal Vector3, point Vector3, ) *Plane {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPlane()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypePlane, 4) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{normal.AsCTypePtr(), point.AsCTypePtr(), }))
  return me
}

func NewPlaneFromVector3Vector3Vector3(point1 Vector3, point2 Vector3, point3 Vector3, ) *Plane {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPlane()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypePlane, 5) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{point1.AsCTypePtr(), point2.AsCTypePtr(), point3.AsCTypePtr(), }))
  return me
}

func NewPlaneFromFloat32Float32Float32Float32(a float64, b float64, c float64, d float64, ) *Plane {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&a)
  pinner.Pin(&b)
  pinner.Pin(&c)
  pinner.Pin(&d)
  me := newPlane()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypePlane, 6) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&a), gdc.ConstTypePtr(&b), gdc.ConstTypePtr(&c), gdc.ConstTypePtr(&d), }))
  return me
}

// Destructor
func (me *Plane) Destroy() {
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsPlane() (*Plane, error) {
	if me.Type() != gdc.VariantTypePlane {
		return nil, fmt.Errorf("variant is not a Plane")
	}
  bclass := newPlane()
	fn := me.iface.GetVariantToTypeConstructor(me.Type())
	me.iface.CallTypeFromVariantConstructorFunc(fn, bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Plane) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  fn := me.iface.GetVariantFromTypeConstructor(me.Type())
  me.iface.CallVariantFromTypeConstructorFunc(fn, va.asUninitialized(), me.AsTypePtr())
  return va
}

// Pointers
func PlaneFromPtr(ptr gdc.ConstTypePtr) *Plane {
  me := newPlane()
  dataFromPtr(me.data[:], ptr)
  return me
}

func (me *Plane) Type() gdc.VariantType {
  return gdc.VariantTypePlane
}

func (me *Plane) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Plane) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Plane) asUninitialized() gdc.UninitializedTypePtr {
  return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *Plane) Normalized() Plane {
  name := StringNameFromStr("normalized")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, name.AsCPtr(), 1051796340) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Plane
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Plane) GetCenter() Vector3 {
  name := StringNameFromStr("get_center")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, name.AsCPtr(), 1776574132) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Vector3
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Plane) IsEqualApprox(to_plane Plane, ) bool {
  name := StringNameFromStr("is_equal_approx")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, name.AsCPtr(), 1150170233) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret bool
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{to_plane.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Plane) IsFinite() bool {
  name := StringNameFromStr("is_finite")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, name.AsCPtr(), 3918633141) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret bool
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Plane) IsPointOver(point Vector3, ) bool {
  name := StringNameFromStr("is_point_over")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, name.AsCPtr(), 1749054343) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret bool
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{point.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Plane) DistanceTo(point Vector3, ) float32 {
  name := StringNameFromStr("distance_to")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, name.AsCPtr(), 1047977935) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret float32
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{point.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Plane) HasPoint(point Vector3, tolerance float32, ) bool {
  name := StringNameFromStr("has_point")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, name.AsCPtr(), 1258189072) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret bool
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{point.AsCTypePtr(), gdc.ConstTypePtr(unsafe.Pointer(&tolerance)), }

  pinner.Pin(&tolerance)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Plane) Project(point Vector3, ) Vector3 {
  name := StringNameFromStr("project")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, name.AsCPtr(), 2923479887) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Vector3
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{point.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Plane) Intersect3(b Plane, c Plane, ) Variant {
  name := StringNameFromStr("intersect_3")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, name.AsCPtr(), 2012052692) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Variant
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{b.AsCTypePtr(), c.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Plane) IntersectsRay(from Vector3, dir Vector3, ) Variant {
  name := StringNameFromStr("intersects_ray")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, name.AsCPtr(), 2048133369) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Variant
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{from.AsCTypePtr(), dir.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Plane) IntersectsSegment(from Vector3, to Vector3, ) Variant {
  name := StringNameFromStr("intersects_segment")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, name.AsCPtr(), 2048133369) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret Variant
  pinner.Pin(&ret)
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

  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&value)

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

  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&value)

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

  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&value)

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

  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&value)

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

  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&value)

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}
