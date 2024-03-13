// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Vector4i struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Constants

var (
  Vector4iZero = "Vector4i(0, 0, 0, 0)" // TODO: construct correctly
  Vector4iOne = "Vector4i(1, 1, 1, 1)" // TODO: construct correctly
  Vector4iMin = "Vector4i(-2147483648, -2147483648, -2147483648, -2147483648)" // TODO: construct correctly
  Vector4iMax = "Vector4i(2147483647, 2147483647, 2147483647, 2147483647)" // TODO: construct correctly
)

// Enums

type Vector4iAxis int
const (
  Vector4iAxisAxisX Vector4iAxis = 0
  Vector4iAxisAxisY Vector4iAxis = 1
  Vector4iAxisAxisZ Vector4iAxis = 2
  Vector4iAxisAxisW Vector4iAxis = 3
)

// Constructors

func NewVector4i() Vector4i {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector4i))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector4I, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Vector4i{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewVector4iFromVector4I(from Vector4i, ) Vector4i {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector4i))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector4I, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Vector4i{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewVector4iFromVector4(from Vector4, ) Vector4i {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector4i))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector4I, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Vector4i{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewVector4iFromIntIntIntInt(x int, y int, z int, w int, ) Vector4i {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector4i))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector4I, 3) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), gdc.ConstTypePtr(&z), gdc.ConstTypePtr(&w), }))
  return Vector4i{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Vector4i) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Vector4i) Type() gdc.VariantType {
  return gdc.VariantTypeVector4I
}

func (me *Vector4i) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Vector4i) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func (me *Vector4i) MinAxisIndex() int {
  name := StringNameFromStr("min_axis_index")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4I, name.AsCPtr(), 3173160232) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4i) MaxAxisIndex() int {
  name := StringNameFromStr("max_axis_index")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4I, name.AsCPtr(), 3173160232) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4i) Length() float32 {
  name := StringNameFromStr("length")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4I, name.AsCPtr(), 466405837) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4i) LengthSquared() int {
  name := StringNameFromStr("length_squared")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4I, name.AsCPtr(), 3173160232) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4i) Sign() Vector4i {
  name := StringNameFromStr("sign")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4I, name.AsCPtr(), 4134919947) // FIXME: should cache?

  var ret Vector4i
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4i) Abs() Vector4i {
  name := StringNameFromStr("abs")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4I, name.AsCPtr(), 4134919947) // FIXME: should cache?

  var ret Vector4i
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4i) Clamp(min Vector4i, max Vector4i, ) Vector4i {
  name := StringNameFromStr("clamp")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4I, name.AsCPtr(), 3046490913) // FIXME: should cache?

  var ret Vector4i
  args := []gdc.ConstTypePtr{min.AsCTypePtr(), max.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Vector4i) Snapped(step Vector4i, ) Vector4i {
  name := StringNameFromStr("snapped")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4I, name.AsCPtr(), 1181693102) // FIXME: should cache?

  var ret Vector4i
  args := []gdc.ConstTypePtr{step.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

// Operators

func (me *Vector4i) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4i) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4i) Negate() Vector4i {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNegate, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret Vector4i
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4i) Positive() Vector4i {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpPositive, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret Vector4i
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4i) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4i) MultiplyInt(rightArg int) Vector4i {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Vector4i
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4i) DivideInt(rightArg int) Vector4i {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, me.Type(), right.Type()) // FIXME: cache
  var ret Vector4i
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4i) ModuleInt(rightArg int) Vector4i {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret Vector4i
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4i) MultiplyFloat32(rightArg float32) Vector4 {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Vector4
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4i) DivideFloat32(rightArg float32) Vector4 {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, me.Type(), right.Type()) // FIXME: cache
  var ret Vector4
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4i) EqualVector4I(right Vector4i) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4i) NotEqualVector4I(right Vector4i) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4i) LessVector4I(right Vector4i) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLess, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4i) LessEqualVector4I(right Vector4i) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLessEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4i) GreaterVector4I(right Vector4i) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreater, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4i) GreaterEqualVector4I(right Vector4i) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreaterEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4i) AddVector4I(right Vector4i) Vector4i {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, me.Type(), right.Type()) // FIXME: cache
  var ret Vector4i
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4i) SubtractVector4I(right Vector4i) Vector4i {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpSubtract, me.Type(), right.Type()) // FIXME: cache
  var ret Vector4i
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4i) MultiplyVector4I(right Vector4i) Vector4i {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  var ret Vector4i
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4i) DivideVector4I(right Vector4i) Vector4i {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, me.Type(), right.Type()) // FIXME: cache
  var ret Vector4i
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4i) ModuleVector4I(right Vector4i) Vector4i {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret Vector4i
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4i) InDictionary(right Dictionary) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4i) InArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

// Members

func (me *Vector4i) X() int {
  name := StringNameFromStr("x")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret int
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4i) SetX(value int) {
  name := StringNameFromStr("x")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}

func (me *Vector4i) Y() int {
  name := StringNameFromStr("y")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret int
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4i) SetY(value int) {
  name := StringNameFromStr("y")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}

func (me *Vector4i) Z() int {
  name := StringNameFromStr("z")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret int
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4i) SetZ(value int) {
  name := StringNameFromStr("z")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}

func (me *Vector4i) W() int {
  name := StringNameFromStr("w")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  var ret int
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Vector4i) SetW(value int) {
  name := StringNameFromStr("w")
  defer name.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), gdc.ConstTypePtr(&value))
}
