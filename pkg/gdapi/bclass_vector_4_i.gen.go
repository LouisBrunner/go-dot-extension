// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "fmt"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Vector4i struct {
  data   *[classSizeVector4i]byte
  iface  gdc.Interface
  pinner runtime.Pinner
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
func newVector4i() *Vector4i {
  me := &Vector4i{
    data:   new([classSizeVector4i]byte),
    iface:  giface,
  }
  me.pinner.Pin(me)
  me.pinner.Pin(me.AsTypePtr())
  return me
}

func NewVector4i() *Vector4i {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newVector4i()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeVector4I, 0) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewVector4iFromVector4I(from Vector4i, ) *Vector4i {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newVector4i()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeVector4I, 1) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewVector4iFromVector4(from Vector4, ) *Vector4i {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newVector4i()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeVector4I, 2) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewVector4iFromIntIntIntInt(x int64, y int64, z int64, w int64, ) *Vector4i {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&x)
  pinner.Pin(&y)
  pinner.Pin(&z)
  pinner.Pin(&w)
  me := newVector4i()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeVector4I, 3) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), gdc.ConstTypePtr(&z), gdc.ConstTypePtr(&w), }))
  return me
}

// Destructor
func (me *Vector4i) Destroy() {
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsVector4i() (*Vector4i, error) {
	if me.Type() != gdc.VariantTypeVector4I {
		return nil, fmt.Errorf("variant is not a Vector4i")
	}
  bclass := newVector4i()
	fn := me.iface.GetVariantToTypeConstructor(me.Type())
	me.iface.CallTypeFromVariantConstructorFunc(fn, bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Vector4i) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  fn := me.iface.GetVariantFromTypeConstructor(me.Type())
  me.iface.CallVariantFromTypeConstructorFunc(fn, va.asUninitialized(), me.AsTypePtr())
  return va
}

// Pointers
func Vector4iFromPtr(ptr gdc.ConstTypePtr) *Vector4i {
  me := newVector4i()
  dataFromPtr(me.data[:], ptr)
  return me
}

func (me *Vector4i) Type() gdc.VariantType {
  return gdc.VariantTypeVector4I
}

func (me *Vector4i) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Vector4i) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Vector4i) asUninitialized() gdc.UninitializedTypePtr {
  return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *Vector4i) MinAxisIndex() int64 {
  name := StringNameFromStr("min_axis_index")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4I, name.AsCPtr(), 3173160232) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Vector4i) MaxAxisIndex() int64 {
  name := StringNameFromStr("max_axis_index")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4I, name.AsCPtr(), 3173160232) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Vector4i) Length() float64 {
  name := StringNameFromStr("length")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4I, name.AsCPtr(), 466405837) // FIXME: should cache?

  ret := NewFloat()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Vector4i) LengthSquared() int64 {
  name := StringNameFromStr("length_squared")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4I, name.AsCPtr(), 3173160232) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Vector4i) Sign() Vector4i {
  name := StringNameFromStr("sign")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4I, name.AsCPtr(), 4134919947) // FIXME: should cache?

  ret := NewVector4i()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Vector4i) Abs() Vector4i {
  name := StringNameFromStr("abs")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4I, name.AsCPtr(), 4134919947) // FIXME: should cache?

  ret := NewVector4i()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Vector4i) Clamp(min Vector4i, max Vector4i, ) Vector4i {
  name := StringNameFromStr("clamp")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4I, name.AsCPtr(), 3046490913) // FIXME: should cache?

  ret := NewVector4i()



  args := []gdc.ConstTypePtr{min.AsCTypePtr(), max.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Vector4i) Snapped(step Vector4i, ) Vector4i {
  name := StringNameFromStr("snapped")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4I, name.AsCPtr(), 1181693102) // FIXME: should cache?

  ret := NewVector4i()


  args := []gdc.ConstTypePtr{step.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

// Operators

func (me *Vector4i) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) Negate() Vector4i {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNegate, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  ret := NewVector4i()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, ret.AsTypePtr())
  return *ret
}

func (me *Vector4i) Positive() Vector4i {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpPositive, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  ret := NewVector4i()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, ret.AsTypePtr())
  return *ret
}

func (me *Vector4i) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) MultiplyInt(rightArg int64) Vector4i {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  ret := NewVector4i()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4i) DivideInt(rightArg int64) Vector4i {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, me.Type(), right.Type()) // FIXME: cache
  ret := NewVector4i()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4i) ModuleInt(rightArg int64) Vector4i {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewVector4i()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4i) MultiplyFloat32(rightArg float64) Vector4 {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  ret := NewVector4()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4i) DivideFloat32(rightArg float64) Vector4 {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, me.Type(), right.Type()) // FIXME: cache
  ret := NewVector4()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4i) EqualVector4I(right Vector4i) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) NotEqualVector4I(right Vector4i) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) LessVector4I(right Vector4i) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLess, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) LessEqualVector4I(right Vector4i) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLessEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) GreaterVector4I(right Vector4i) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreater, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) GreaterEqualVector4I(right Vector4i) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreaterEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) AddVector4I(right Vector4i) Vector4i {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, me.Type(), right.Type()) // FIXME: cache
  ret := NewVector4i()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4i) SubtractVector4I(right Vector4i) Vector4i {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpSubtract, me.Type(), right.Type()) // FIXME: cache
  ret := NewVector4i()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4i) MultiplyVector4I(right Vector4i) Vector4i {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  ret := NewVector4i()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4i) DivideVector4I(right Vector4i) Vector4i {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, me.Type(), right.Type()) // FIXME: cache
  ret := NewVector4i()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4i) ModuleVector4I(right Vector4i) Vector4i {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewVector4i()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4i) InDictionary(right Dictionary) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) InArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

// Members

func (me *Vector4i) X() int64 {
  name := StringNameFromStr("x")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  ret := NewInt()
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) SetX(value int64) {
  name := StringNameFromStr("x")
  defer name.Destroy()
  valueV := NewIntFromInt(value)
  defer valueV.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Vector4i) Y() int64 {
  name := StringNameFromStr("y")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  ret := NewInt()
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) SetY(value int64) {
  name := StringNameFromStr("y")
  defer name.Destroy()
  valueV := NewIntFromInt(value)
  defer valueV.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Vector4i) Z() int64 {
  name := StringNameFromStr("z")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  ret := NewInt()
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) SetZ(value int64) {
  name := StringNameFromStr("z")
  defer name.Destroy()
  valueV := NewIntFromInt(value)
  defer valueV.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Vector4i) W() int64 {
  name := StringNameFromStr("w")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  ret := NewInt()
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) SetW(value int64) {
  name := StringNameFromStr("w")
  defer name.Destroy()
  valueV := NewIntFromInt(value)
  defer valueV.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), valueV.AsCTypePtr())
}
