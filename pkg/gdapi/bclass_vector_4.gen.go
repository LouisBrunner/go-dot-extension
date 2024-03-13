// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Vector4 struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Constants

var (
  Vector4Zero = "Vector4(0, 0, 0, 0)" // TODO: construct correctly
  Vector4One = "Vector4(1, 1, 1, 1)" // TODO: construct correctly
  Vector4Inf = "Vector4(inf, inf, inf, inf)" // TODO: construct correctly
)

// Enums

type Vector4Axis int
const (
  Vector4AxisAxisX Vector4Axis = 0
  Vector4AxisAxisY Vector4Axis = 1
  Vector4AxisAxisZ Vector4Axis = 2
  Vector4AxisAxisW Vector4Axis = 3
)

// Constructors

func NewVector4() Vector4 {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector4))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector4, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Vector4{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewVector4FromVector4(from Vector4, ) Vector4 {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector4))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector4, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Vector4{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewVector4FromVector4I(from Vector4i, ) Vector4 {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector4))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector4, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Vector4{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewVector4FromFloat32Float32Float32Float32(x float32, y float32, z float32, w float32, ) Vector4 {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector4))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector4, 3) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), gdc.ConstTypePtr(&z), gdc.ConstTypePtr(&w), }))
  return Vector4{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Vector4) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Vector4) Type() gdc.VariantType {
  return gdc.VariantTypeVector4
}

func (me *Vector4) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Vector4) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func (me *Vector4) MinAxisIndex() int {
  name := StringNameFromStr("min_axis_index")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, name.AsCPtr(), 3173160232) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4) MaxAxisIndex() int {
  name := StringNameFromStr("max_axis_index")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, name.AsCPtr(), 3173160232) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4) Length() float32 {
  name := StringNameFromStr("length")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, name.AsCPtr(), 466405837) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4) LengthSquared() float32 {
  name := StringNameFromStr("length_squared")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, name.AsCPtr(), 466405837) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4) Abs() Vector4 {
  name := StringNameFromStr("abs")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, name.AsCPtr(), 80860099) // FIXME: should cache?

  var ret Vector4
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4) Sign() Vector4 {
  name := StringNameFromStr("sign")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, name.AsCPtr(), 80860099) // FIXME: should cache?

  var ret Vector4
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4) Floor() Vector4 {
  name := StringNameFromStr("floor")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, name.AsCPtr(), 80860099) // FIXME: should cache?

  var ret Vector4
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4) Ceil() Vector4 {
  name := StringNameFromStr("ceil")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, name.AsCPtr(), 80860099) // FIXME: should cache?

  var ret Vector4
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4) Round() Vector4 {
  name := StringNameFromStr("round")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, name.AsCPtr(), 80860099) // FIXME: should cache?

  var ret Vector4
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4) Lerp(to Vector4, weight float32, ) Vector4 {
  name := StringNameFromStr("lerp")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, name.AsCPtr(), 2329757942) // FIXME: should cache?

  var ret Vector4
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), gdc.ConstTypePtr(&weight), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4) CubicInterpolate(b Vector4, pre_a Vector4, post_b Vector4, weight float32, ) Vector4 {
  name := StringNameFromStr("cubic_interpolate")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, name.AsCPtr(), 726768410) // FIXME: should cache?

  var ret Vector4
  args := []gdc.ConstTypePtr{b.AsCTypePtr(), pre_a.AsCTypePtr(), post_b.AsCTypePtr(), gdc.ConstTypePtr(&weight), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4) CubicInterpolateInTime(b Vector4, pre_a Vector4, post_b Vector4, weight float32, b_t float32, pre_a_t float32, post_b_t float32, ) Vector4 {
  name := StringNameFromStr("cubic_interpolate_in_time")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, name.AsCPtr(), 681631873) // FIXME: should cache?

  var ret Vector4
  args := []gdc.ConstTypePtr{b.AsCTypePtr(), pre_a.AsCTypePtr(), post_b.AsCTypePtr(), gdc.ConstTypePtr(&weight), gdc.ConstTypePtr(&b_t), gdc.ConstTypePtr(&pre_a_t), gdc.ConstTypePtr(&post_b_t), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4) Posmod(mod float32, ) Vector4 {
  name := StringNameFromStr("posmod")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, name.AsCPtr(), 3129671720) // FIXME: should cache?

  var ret Vector4
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mod), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4) Posmodv(modv Vector4, ) Vector4 {
  name := StringNameFromStr("posmodv")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, name.AsCPtr(), 2031281584) // FIXME: should cache?

  var ret Vector4
  args := []gdc.ConstTypePtr{modv.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4) Snapped(step Vector4, ) Vector4 {
  name := StringNameFromStr("snapped")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, name.AsCPtr(), 2031281584) // FIXME: should cache?

  var ret Vector4
  args := []gdc.ConstTypePtr{step.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4) Clamp(min Vector4, max Vector4, ) Vector4 {
  name := StringNameFromStr("clamp")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, name.AsCPtr(), 823915692) // FIXME: should cache?

  var ret Vector4
  args := []gdc.ConstTypePtr{min.AsCTypePtr(), max.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4) Normalized() Vector4 {
  name := StringNameFromStr("normalized")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, name.AsCPtr(), 80860099) // FIXME: should cache?

  var ret Vector4
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4) IsNormalized() bool {
  name := StringNameFromStr("is_normalized")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4) DirectionTo(to Vector4, ) Vector4 {
  name := StringNameFromStr("direction_to")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, name.AsCPtr(), 2031281584) // FIXME: should cache?

  var ret Vector4
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4) DistanceTo(to Vector4, ) float32 {
  name := StringNameFromStr("distance_to")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, name.AsCPtr(), 3770801042) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4) DistanceSquaredTo(to Vector4, ) float32 {
  name := StringNameFromStr("distance_squared_to")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, name.AsCPtr(), 3770801042) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4) Dot(with Vector4, ) float32 {
  name := StringNameFromStr("dot")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, name.AsCPtr(), 3770801042) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{with.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4) Inverse() Vector4 {
  name := StringNameFromStr("inverse")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, name.AsCPtr(), 80860099) // FIXME: should cache?

  var ret Vector4
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4) IsEqualApprox(to Vector4, ) bool {
  name := StringNameFromStr("is_equal_approx")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, name.AsCPtr(), 88913544) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4) IsZeroApprox() bool {
  name := StringNameFromStr("is_zero_approx")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4) IsFinite() bool {
  name := StringNameFromStr("is_finite")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

// Operators

func (me *Vector4) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4) Negate() Vector4 {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNegate, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret Vector4
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4) Positive() Vector4 {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpPositive, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret Vector4
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4) MultiplyInt(rightArg int) Vector4 {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Vector4
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4) DivideInt(rightArg int) Vector4 {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, me.Type(), right.Type()) // FIXME: cache
  var ret Vector4
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4) MultiplyFloat32(rightArg float32) Vector4 {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Vector4
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4) DivideFloat32(rightArg float32) Vector4 {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, me.Type(), right.Type()) // FIXME: cache
  var ret Vector4
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4) EqualVector4(right Vector4) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4) NotEqualVector4(right Vector4) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4) LessVector4(right Vector4) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLess, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4) LessEqualVector4(right Vector4) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLessEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4) GreaterVector4(right Vector4) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreater, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4) GreaterEqualVector4(right Vector4) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreaterEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4) AddVector4(right Vector4) Vector4 {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, me.Type(), right.Type()) // FIXME: cache
  var ret Vector4
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4) SubtractVector4(right Vector4) Vector4 {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpSubtract, me.Type(), right.Type()) // FIXME: cache
  var ret Vector4
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4) MultiplyVector4(right Vector4) Vector4 {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Vector4
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4) DivideVector4(right Vector4) Vector4 {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, me.Type(), right.Type()) // FIXME: cache
  var ret Vector4
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4) MultiplyProjection(right Projection) Vector4 {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Vector4
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4) InDictionary(right Dictionary) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4) InArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

// Members

func (me *Vector4) X() float32 {
  name := StringNameFromStr("x")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret float32
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4) SetX(value float32) {
  name := StringNameFromStr("x")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}

func (me *Vector4) Y() float32 {
  name := StringNameFromStr("y")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret float32
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4) SetY(value float32) {
  name := StringNameFromStr("y")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}

func (me *Vector4) Z() float32 {
  name := StringNameFromStr("z")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret float32
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4) SetZ(value float32) {
  name := StringNameFromStr("z")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}

func (me *Vector4) W() float32 {
  name := StringNameFromStr("w")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret float32
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4) SetW(value float32) {
  name := StringNameFromStr("w")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}
