// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "fmt"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Vector3i struct {
  data   *[classSizeVector3i]byte
  iface  gdc.Interface
  pinner runtime.Pinner
}

// Constants

var (
  Vector3iZero = "Vector3i(0, 0, 0)" // TODO: construct correctly
  Vector3iOne = "Vector3i(1, 1, 1)" // TODO: construct correctly
  Vector3iMin = "Vector3i(-2147483648, -2147483648, -2147483648)" // TODO: construct correctly
  Vector3iMax = "Vector3i(2147483647, 2147483647, 2147483647)" // TODO: construct correctly
  Vector3iLeft = "Vector3i(-1, 0, 0)" // TODO: construct correctly
  Vector3iRight = "Vector3i(1, 0, 0)" // TODO: construct correctly
  Vector3iUp = "Vector3i(0, 1, 0)" // TODO: construct correctly
  Vector3iDown = "Vector3i(0, -1, 0)" // TODO: construct correctly
  Vector3iForward = "Vector3i(0, 0, -1)" // TODO: construct correctly
  Vector3iBack = "Vector3i(0, 0, 1)" // TODO: construct correctly
)

// Enums

type Vector3iAxis int
const (
  Vector3iAxisAxisX Vector3iAxis = 0
  Vector3iAxisAxisY Vector3iAxis = 1
  Vector3iAxisAxisZ Vector3iAxis = 2
)

// Constructors
func newVector3i() *Vector3i {
  me := &Vector3i{
    data:   new([classSizeVector3i]byte),
    iface:  giface,
  }
  me.pinner.Pin(me)
  me.pinner.Pin(me.AsTypePtr())
  return me
}

func NewVector3i() *Vector3i {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newVector3i()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeVector3I, 0) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewVector3iFromVector3I(from Vector3i, ) *Vector3i {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newVector3i()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeVector3I, 1) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewVector3iFromVector3(from Vector3, ) *Vector3i {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newVector3i()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeVector3I, 2) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewVector3iFromIntIntInt(x int64, y int64, z int64, ) *Vector3i {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&x)
  pinner.Pin(&y)
  pinner.Pin(&z)
  me := newVector3i()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeVector3I, 3) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), gdc.ConstTypePtr(&z), }))
  return me
}

// Destructor
func (me *Vector3i) Destroy() {
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsVector3i() (*Vector3i, error) {
	if me.Type() != gdc.VariantTypeVector3I {
		return nil, fmt.Errorf("variant is not a Vector3i")
	}
  bclass := newVector3i()
	fn := me.iface.GetVariantToTypeConstructor(me.Type())
	me.iface.CallTypeFromVariantConstructorFunc(fn, bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Vector3i) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  fn := me.iface.GetVariantFromTypeConstructor(me.Type())
  me.iface.CallVariantFromTypeConstructorFunc(fn, va.asUninitialized(), me.AsTypePtr())
  return va
}

// Pointers
func Vector3iFromPtr(ptr gdc.ConstTypePtr) *Vector3i {
  me := newVector3i()
  dataFromPtr(me.data[:], ptr)
  return me
}

func (me *Vector3i) Type() gdc.VariantType {
  return gdc.VariantTypeVector3I
}

func (me *Vector3i) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Vector3i) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Vector3i) asUninitialized() gdc.UninitializedTypePtr {
  return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *Vector3i) MinAxisIndex() int64 {
  name := StringNameFromStr("min_axis_index")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3I, name.AsCPtr(), 3173160232) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Vector3i) MaxAxisIndex() int64 {
  name := StringNameFromStr("max_axis_index")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3I, name.AsCPtr(), 3173160232) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Vector3i) Length() float64 {
  name := StringNameFromStr("length")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3I, name.AsCPtr(), 466405837) // FIXME: should cache?

  ret := NewFloat()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Vector3i) LengthSquared() int64 {
  name := StringNameFromStr("length_squared")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3I, name.AsCPtr(), 3173160232) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Vector3i) Sign() Vector3i {
  name := StringNameFromStr("sign")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3I, name.AsCPtr(), 3729604559) // FIXME: should cache?

  ret := NewVector3i()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Vector3i) Abs() Vector3i {
  name := StringNameFromStr("abs")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3I, name.AsCPtr(), 3729604559) // FIXME: should cache?

  ret := NewVector3i()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Vector3i) Clamp(min Vector3i, max Vector3i, ) Vector3i {
  name := StringNameFromStr("clamp")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3I, name.AsCPtr(), 1086892323) // FIXME: should cache?

  ret := NewVector3i()



  args := []gdc.ConstTypePtr{min.AsCTypePtr(), max.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Vector3i) Snapped(step Vector3i, ) Vector3i {
  name := StringNameFromStr("snapped")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3I, name.AsCPtr(), 1989319750) // FIXME: should cache?

  ret := NewVector3i()


  args := []gdc.ConstTypePtr{step.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

// Operators

func (me *Vector3i) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector3i) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector3i) Negate() Vector3i {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNegate, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  ret := NewVector3i()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, ret.AsTypePtr())
  return *ret
}

func (me *Vector3i) Positive() Vector3i {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpPositive, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  ret := NewVector3i()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, ret.AsTypePtr())
  return *ret
}

func (me *Vector3i) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector3i) MultiplyInt(rightArg int64) Vector3i {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  ret := NewVector3i()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector3i) DivideInt(rightArg int64) Vector3i {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, me.Type(), right.Type()) // FIXME: cache
  ret := NewVector3i()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector3i) ModuleInt(rightArg int64) Vector3i {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewVector3i()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector3i) MultiplyFloat32(rightArg float64) Vector3 {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  ret := NewVector3()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector3i) DivideFloat32(rightArg float64) Vector3 {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, me.Type(), right.Type()) // FIXME: cache
  ret := NewVector3()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector3i) EqualVector3I(right Vector3i) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector3i) NotEqualVector3I(right Vector3i) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector3i) LessVector3I(right Vector3i) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLess, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector3i) LessEqualVector3I(right Vector3i) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLessEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector3i) GreaterVector3I(right Vector3i) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreater, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector3i) GreaterEqualVector3I(right Vector3i) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreaterEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector3i) AddVector3I(right Vector3i) Vector3i {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, me.Type(), right.Type()) // FIXME: cache
  ret := NewVector3i()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector3i) SubtractVector3I(right Vector3i) Vector3i {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpSubtract, me.Type(), right.Type()) // FIXME: cache
  ret := NewVector3i()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector3i) MultiplyVector3I(right Vector3i) Vector3i {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, me.Type(), right.Type()) // FIXME: cache
  ret := NewVector3i()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector3i) DivideVector3I(right Vector3i) Vector3i {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, me.Type(), right.Type()) // FIXME: cache
  ret := NewVector3i()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector3i) ModuleVector3I(right Vector3i) Vector3i {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewVector3i()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector3i) InDictionary(right Dictionary) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector3i) InArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

// Members

func (me *Vector3i) X() int64 {
  name := StringNameFromStr("x")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  ret := NewInt()
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector3i) SetX(value int64) {
  name := StringNameFromStr("x")
  defer name.Destroy()
  valueV := NewIntFromInt(value)
  defer valueV.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Vector3i) Y() int64 {
  name := StringNameFromStr("y")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  ret := NewInt()
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector3i) SetY(value int64) {
  name := StringNameFromStr("y")
  defer name.Destroy()
  valueV := NewIntFromInt(value)
  defer valueV.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Vector3i) Z() int64 {
  name := StringNameFromStr("z")
  defer name.Destroy()

  getter := me.iface.VariantGetPtrGetter(me.Type(), name.AsCPtr()) // FIXME: cache
  ret := NewInt()
  me.iface.CallPtrGetter(getter, me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector3i) SetZ(value int64) {
  name := StringNameFromStr("z")
  defer name.Destroy()
  valueV := NewIntFromInt(value)
  defer valueV.Destroy()

  setter := me.iface.VariantGetPtrSetter(me.Type(), name.AsCPtr()) // FIXME: cache
  me.iface.CallPtrSetter(setter, me.AsTypePtr(), valueV.AsCTypePtr())
}
