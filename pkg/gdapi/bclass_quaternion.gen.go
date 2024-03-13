// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Quaternion struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Constants

var (
  QuaternionIdentity = "Quaternion(0, 0, 0, 1)" // TODO: construct correctly
)

// Enums

// Constructors

func NewQuaternion() Quaternion {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeQuaternion))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeQuaternion, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Quaternion{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewQuaternionFromQuaternion(from Quaternion, ) Quaternion {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeQuaternion))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeQuaternion, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Quaternion{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewQuaternionFromBasis(from Basis, ) Quaternion {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeQuaternion))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeQuaternion, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Quaternion{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewQuaternionFromVector3Float32(axis Vector3, angle float32, ) Quaternion {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeQuaternion))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeQuaternion, 3) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{axis.AsCTypePtr(), gdc.ConstTypePtr(&angle), }))
  return Quaternion{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewQuaternionFromVector3Vector3(arc_from Vector3, arc_to Vector3, ) Quaternion {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeQuaternion))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeQuaternion, 4) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{arc_from.AsCTypePtr(), arc_to.AsCTypePtr(), }))
  return Quaternion{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewQuaternionFromFloat32Float32Float32Float32(x float32, y float32, z float32, w float32, ) Quaternion {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeQuaternion))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeQuaternion, 5) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), gdc.ConstTypePtr(&z), gdc.ConstTypePtr(&w), }))
  return Quaternion{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Quaternion) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Quaternion) Type() gdc.VariantType {
  return gdc.VariantTypeQuaternion
}

func (me *Quaternion) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Quaternion) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func (me *Quaternion) Length() float32 {
  name := StringNameFromStr("length")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, name.AsCPtr(), 466405837) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Quaternion) LengthSquared() float32 {
  name := StringNameFromStr("length_squared")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, name.AsCPtr(), 466405837) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Quaternion) Normalized() Quaternion {
  name := StringNameFromStr("normalized")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, name.AsCPtr(), 4274879941) // FIXME: should cache?

  var ret Quaternion
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Quaternion) IsNormalized() bool {
  name := StringNameFromStr("is_normalized")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Quaternion) IsEqualApprox(to Quaternion, ) bool {
  name := StringNameFromStr("is_equal_approx")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, name.AsCPtr(), 1682156903) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Quaternion) IsFinite() bool {
  name := StringNameFromStr("is_finite")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Quaternion) Inverse() Quaternion {
  name := StringNameFromStr("inverse")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, name.AsCPtr(), 4274879941) // FIXME: should cache?

  var ret Quaternion
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Quaternion) Log() Quaternion {
  name := StringNameFromStr("log")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, name.AsCPtr(), 4274879941) // FIXME: should cache?

  var ret Quaternion
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Quaternion) Exp() Quaternion {
  name := StringNameFromStr("exp")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, name.AsCPtr(), 4274879941) // FIXME: should cache?

  var ret Quaternion
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Quaternion) AngleTo(to Quaternion, ) float32 {
  name := StringNameFromStr("angle_to")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, name.AsCPtr(), 3244682419) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Quaternion) Dot(with Quaternion, ) float32 {
  name := StringNameFromStr("dot")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, name.AsCPtr(), 3244682419) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{with.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Quaternion) Slerp(to Quaternion, weight float32, ) Quaternion {
  name := StringNameFromStr("slerp")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, name.AsCPtr(), 1773590316) // FIXME: should cache?

  var ret Quaternion
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), gdc.ConstTypePtr(&weight), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Quaternion) Slerpni(to Quaternion, weight float32, ) Quaternion {
  name := StringNameFromStr("slerpni")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, name.AsCPtr(), 1773590316) // FIXME: should cache?

  var ret Quaternion
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), gdc.ConstTypePtr(&weight), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Quaternion) SphericalCubicInterpolate(b Quaternion, pre_a Quaternion, post_b Quaternion, weight float32, ) Quaternion {
  name := StringNameFromStr("spherical_cubic_interpolate")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, name.AsCPtr(), 2150967576) // FIXME: should cache?

  var ret Quaternion
  args := []gdc.ConstTypePtr{b.AsCTypePtr(), pre_a.AsCTypePtr(), post_b.AsCTypePtr(), gdc.ConstTypePtr(&weight), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Quaternion) SphericalCubicInterpolateInTime(b Quaternion, pre_a Quaternion, post_b Quaternion, weight float32, b_t float32, pre_a_t float32, post_b_t float32, ) Quaternion {
  name := StringNameFromStr("spherical_cubic_interpolate_in_time")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, name.AsCPtr(), 1436023539) // FIXME: should cache?

  var ret Quaternion
  args := []gdc.ConstTypePtr{b.AsCTypePtr(), pre_a.AsCTypePtr(), post_b.AsCTypePtr(), gdc.ConstTypePtr(&weight), gdc.ConstTypePtr(&b_t), gdc.ConstTypePtr(&pre_a_t), gdc.ConstTypePtr(&post_b_t), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Quaternion) GetEuler(order int, ) Vector3 {
  name := StringNameFromStr("get_euler")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, name.AsCPtr(), 1394941017) // FIXME: should cache?

  var ret Vector3
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&order), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func QuaternionFromEuler(euler Vector3, ) Quaternion {
  name := StringNameFromStr("from_euler")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, name.AsCPtr(), 4053467903) // FIXME: should cache?

  var ret Quaternion
  args := []gdc.ConstTypePtr{euler.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, nil, unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Quaternion) GetAxis() Vector3 {
  name := StringNameFromStr("get_axis")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, name.AsCPtr(), 1776574132) // FIXME: should cache?

  var ret Vector3
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Quaternion) GetAngle() float32 {
  name := StringNameFromStr("get_angle")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, name.AsCPtr(), 466405837) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

// Operators

func (me *Quaternion) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Quaternion) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Quaternion) Negate() Quaternion {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNegate, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret Quaternion
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *Quaternion) Positive() Quaternion {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpPositive, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret Quaternion
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *Quaternion) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *Quaternion) MultiplyInt(rightArg int) Quaternion {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Quaternion
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Quaternion) DivideInt(rightArg int) Quaternion {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, me.Type(), right.Type()) // FIXME: cache
  var ret Quaternion
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Quaternion) MultiplyFloat32(rightArg float32) Quaternion {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Quaternion
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Quaternion) DivideFloat32(rightArg float32) Quaternion {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, me.Type(), right.Type()) // FIXME: cache
  var ret Quaternion
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Quaternion) MultiplyVector3(right Vector3) Vector3 {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Vector3
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Quaternion) EqualQuaternion(right Quaternion) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Quaternion) NotEqualQuaternion(right Quaternion) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Quaternion) AddQuaternion(right Quaternion) Quaternion {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, me.Type(), right.Type()) // FIXME: cache
  var ret Quaternion
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Quaternion) SubtractQuaternion(right Quaternion) Quaternion {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpSubtract, me.Type(), right.Type()) // FIXME: cache
  var ret Quaternion
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Quaternion) MultiplyQuaternion(right Quaternion) Quaternion {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Quaternion
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Quaternion) InDictionary(right Dictionary) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Quaternion) InArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

// Members

func (me *Quaternion) X() float32 {
  name := StringNameFromStr("x")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret float32
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Quaternion) SetX(value float32) {
  name := StringNameFromStr("x")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}

func (me *Quaternion) Y() float32 {
  name := StringNameFromStr("y")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret float32
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Quaternion) SetY(value float32) {
  name := StringNameFromStr("y")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}

func (me *Quaternion) Z() float32 {
  name := StringNameFromStr("z")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret float32
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Quaternion) SetZ(value float32) {
  name := StringNameFromStr("z")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}

func (me *Quaternion) W() float32 {
  name := StringNameFromStr("w")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret float32
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Quaternion) SetW(value float32) {
  name := StringNameFromStr("w")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}
