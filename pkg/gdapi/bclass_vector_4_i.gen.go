// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "fmt"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused imports
var _ = fmt.Sprintf("")

type ptrsForVector4iList struct {
  ctrFn gdc.PtrConstructor
  ctrFromVector4IFn gdc.PtrConstructor
  ctrFromVector4Fn gdc.PtrConstructor
  ctrFromIntIntIntIntFn gdc.PtrConstructor
  methodMinAxisIndexFn gdc.PtrBuiltInMethod
  methodMaxAxisIndexFn gdc.PtrBuiltInMethod
  methodLengthFn gdc.PtrBuiltInMethod
  methodLengthSquaredFn gdc.PtrBuiltInMethod
  methodSignFn gdc.PtrBuiltInMethod
  methodAbsFn gdc.PtrBuiltInMethod
  methodClampFn gdc.PtrBuiltInMethod
  methodSnappedFn gdc.PtrBuiltInMethod
  operatorNegateFn gdc.PtrOperatorEvaluator
  operatorPositiveFn gdc.PtrOperatorEvaluator
  operatorNotFn gdc.PtrOperatorEvaluator
  operatorMultiplyIntFn gdc.PtrOperatorEvaluator
  operatorDivideIntFn gdc.PtrOperatorEvaluator
  operatorModuleIntFn gdc.PtrOperatorEvaluator
  operatorMultiplyFloat32Fn gdc.PtrOperatorEvaluator
  operatorDivideFloat32Fn gdc.PtrOperatorEvaluator
  operatorEqualVector4IFn gdc.PtrOperatorEvaluator
  operatorNotEqualVector4IFn gdc.PtrOperatorEvaluator
  operatorLessVector4IFn gdc.PtrOperatorEvaluator
  operatorLessEqualVector4IFn gdc.PtrOperatorEvaluator
  operatorGreaterVector4IFn gdc.PtrOperatorEvaluator
  operatorGreaterEqualVector4IFn gdc.PtrOperatorEvaluator
  operatorAddVector4IFn gdc.PtrOperatorEvaluator
  operatorSubtractVector4IFn gdc.PtrOperatorEvaluator
  operatorMultiplyVector4IFn gdc.PtrOperatorEvaluator
  operatorDivideVector4IFn gdc.PtrOperatorEvaluator
  operatorModuleVector4IFn gdc.PtrOperatorEvaluator
  operatorInDictionaryFn gdc.PtrOperatorEvaluator
  operatorInArrayFn gdc.PtrOperatorEvaluator
  memberxGetterFn gdc.PtrGetter
  memberxSetterFn gdc.PtrSetter
  memberyGetterFn gdc.PtrGetter
  memberySetterFn gdc.PtrSetter
  memberzGetterFn gdc.PtrGetter
  memberzSetterFn gdc.PtrSetter
  memberwGetterFn gdc.PtrGetter
  memberwSetterFn gdc.PtrSetter
  toVariantFn gdc.TypeFromVariantConstructorFunc
  fromVariantFn gdc.VariantFromTypeConstructorFunc
}

var ptrsForVector4i ptrsForVector4iList

func initVector4iPtrs(iface gdc.Interface) {
  ptrsForVector4i.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeVector4I, 0))
  ptrsForVector4i.ctrFromVector4IFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeVector4I, 1))
  ptrsForVector4i.ctrFromVector4Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeVector4I, 2))
  ptrsForVector4i.ctrFromIntIntIntIntFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeVector4I, 3))
  {
    methodName := StringNameFromStr("min_axis_index")
    defer methodName.Destroy()
    ptrsForVector4i.methodMinAxisIndexFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4I, methodName.AsCPtr(), 3173160232))
  }
  {
    methodName := StringNameFromStr("max_axis_index")
    defer methodName.Destroy()
    ptrsForVector4i.methodMaxAxisIndexFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4I, methodName.AsCPtr(), 3173160232))
  }
  {
    methodName := StringNameFromStr("length")
    defer methodName.Destroy()
    ptrsForVector4i.methodLengthFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4I, methodName.AsCPtr(), 466405837))
  }
  {
    methodName := StringNameFromStr("length_squared")
    defer methodName.Destroy()
    ptrsForVector4i.methodLengthSquaredFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4I, methodName.AsCPtr(), 3173160232))
  }
  {
    methodName := StringNameFromStr("sign")
    defer methodName.Destroy()
    ptrsForVector4i.methodSignFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4I, methodName.AsCPtr(), 4134919947))
  }
  {
    methodName := StringNameFromStr("abs")
    defer methodName.Destroy()
    ptrsForVector4i.methodAbsFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4I, methodName.AsCPtr(), 4134919947))
  }
  {
    methodName := StringNameFromStr("clamp")
    defer methodName.Destroy()
    ptrsForVector4i.methodClampFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4I, methodName.AsCPtr(), 3046490913))
  }
  {
    methodName := StringNameFromStr("snapped")
    defer methodName.Destroy()
    ptrsForVector4i.methodSnappedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4I, methodName.AsCPtr(), 1181693102))
  }
  ptrsForVector4i.operatorNegateFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNegate, gdc.VariantTypeVector4I, gdc.VariantTypeNil))
  ptrsForVector4i.operatorPositiveFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpPositive, gdc.VariantTypeVector4I, gdc.VariantTypeNil))
  ptrsForVector4i.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypeVector4I, gdc.VariantTypeNil))
  ptrsForVector4i.operatorMultiplyIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeVector4I, gdc.VariantTypeInt))
  ptrsForVector4i.operatorDivideIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeVector4I, gdc.VariantTypeInt))
  ptrsForVector4i.operatorModuleIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeVector4I, gdc.VariantTypeInt))
  ptrsForVector4i.operatorMultiplyFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeVector4I, gdc.VariantTypeFloat))
  ptrsForVector4i.operatorDivideFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeVector4I, gdc.VariantTypeFloat))
  ptrsForVector4i.operatorEqualVector4IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeVector4I, gdc.VariantTypeVector4I))
  ptrsForVector4i.operatorNotEqualVector4IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeVector4I, gdc.VariantTypeVector4I))
  ptrsForVector4i.operatorLessVector4IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLess, gdc.VariantTypeVector4I, gdc.VariantTypeVector4I))
  ptrsForVector4i.operatorLessEqualVector4IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLessEqual, gdc.VariantTypeVector4I, gdc.VariantTypeVector4I))
  ptrsForVector4i.operatorGreaterVector4IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreater, gdc.VariantTypeVector4I, gdc.VariantTypeVector4I))
  ptrsForVector4i.operatorGreaterEqualVector4IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreaterEqual, gdc.VariantTypeVector4I, gdc.VariantTypeVector4I))
  ptrsForVector4i.operatorAddVector4IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, gdc.VariantTypeVector4I, gdc.VariantTypeVector4I))
  ptrsForVector4i.operatorSubtractVector4IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpSubtract, gdc.VariantTypeVector4I, gdc.VariantTypeVector4I))
  ptrsForVector4i.operatorMultiplyVector4IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeVector4I, gdc.VariantTypeVector4I))
  ptrsForVector4i.operatorDivideVector4IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeVector4I, gdc.VariantTypeVector4I))
  ptrsForVector4i.operatorModuleVector4IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeVector4I, gdc.VariantTypeVector4I))
  ptrsForVector4i.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeVector4I, gdc.VariantTypeDictionary))
  ptrsForVector4i.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeVector4I, gdc.VariantTypeArray))
  {
    memberName := StringNameFromStr("x")
    defer memberName.Destroy()
    ptrsForVector4i.memberxGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeVector4I, memberName.AsCPtr()))
    ptrsForVector4i.memberxSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeVector4I, memberName.AsCPtr()))
  }
  {
    memberName := StringNameFromStr("y")
    defer memberName.Destroy()
    ptrsForVector4i.memberyGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeVector4I, memberName.AsCPtr()))
    ptrsForVector4i.memberySetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeVector4I, memberName.AsCPtr()))
  }
  {
    memberName := StringNameFromStr("z")
    defer memberName.Destroy()
    ptrsForVector4i.memberzGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeVector4I, memberName.AsCPtr()))
    ptrsForVector4i.memberzSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeVector4I, memberName.AsCPtr()))
  }
  {
    memberName := StringNameFromStr("w")
    defer memberName.Destroy()
    ptrsForVector4i.memberwGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeVector4I, memberName.AsCPtr()))
    ptrsForVector4i.memberwSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeVector4I, memberName.AsCPtr()))
  }
  ptrsForVector4i.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypeVector4I))
  ptrsForVector4i.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypeVector4I))
}

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
  me.iface.CallPtrConstructor(ensurePtr(ptrsForVector4i.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewVector4iFromVector4I(from Vector4i, ) *Vector4i {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newVector4i()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForVector4i.ctrFromVector4IFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewVector4iFromVector4(from Vector4, ) *Vector4i {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newVector4i()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForVector4i.ctrFromVector4Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
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
  me.iface.CallPtrConstructor(ensurePtr(ptrsForVector4i.ctrFromIntIntIntIntFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), gdc.ConstTypePtr(&z), gdc.ConstTypePtr(&w), }))
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
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForVector4i.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Vector4i) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForVector4i.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
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
  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4i.methodMinAxisIndexFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Vector4i) MaxAxisIndex() int64 {
  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4i.methodMaxAxisIndexFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Vector4i) Length() float64 {
  ret := NewFloat()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4i.methodLengthFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Vector4i) LengthSquared() int64 {
  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4i.methodLengthSquaredFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Vector4i) Sign() Vector4i {
  ret := NewVector4i()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4i.methodSignFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Vector4i) Abs() Vector4i {
  ret := NewVector4i()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4i.methodAbsFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Vector4i) Clamp(min Vector4i, max Vector4i, ) Vector4i {
  ret := NewVector4i()



  args := []gdc.ConstTypePtr{min.AsCTypePtr(), max.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4i.methodClampFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Vector4i) Snapped(step Vector4i, ) Vector4i {
  ret := NewVector4i()


  args := []gdc.ConstTypePtr{step.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4i.methodSnappedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

// Operators

func (me *Vector4i) EqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) NotEqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) Negate() Vector4i {
  opPtr := ptrsForVector4i.operatorNegateFn
  ret := NewVector4i()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
  return *ret
}

func (me *Vector4i) Positive() Vector4i {
  opPtr := ptrsForVector4i.operatorPositiveFn
  ret := NewVector4i()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
  return *ret
}

func (me *Vector4i) Not() bool {
  opPtr := ptrsForVector4i.operatorNotFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) MultiplyInt(rightArg int64) Vector4i {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  opPtr := ptrsForVector4i.operatorMultiplyIntFn
  ret := NewVector4i()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4i) DivideInt(rightArg int64) Vector4i {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  opPtr := ptrsForVector4i.operatorDivideIntFn
  ret := NewVector4i()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4i) ModuleInt(rightArg int64) Vector4i {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  opPtr := ptrsForVector4i.operatorModuleIntFn
  ret := NewVector4i()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4i) MultiplyFloat32(rightArg float64) Vector4 {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  opPtr := ptrsForVector4i.operatorMultiplyFloat32Fn
  ret := NewVector4()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4i) DivideFloat32(rightArg float64) Vector4 {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  opPtr := ptrsForVector4i.operatorDivideFloat32Fn
  ret := NewVector4()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4i) EqualVector4I(right Vector4i) bool {
  opPtr := ptrsForVector4i.operatorEqualVector4IFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) NotEqualVector4I(right Vector4i) bool {
  opPtr := ptrsForVector4i.operatorNotEqualVector4IFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) LessVector4I(right Vector4i) bool {
  opPtr := ptrsForVector4i.operatorLessVector4IFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) LessEqualVector4I(right Vector4i) bool {
  opPtr := ptrsForVector4i.operatorLessEqualVector4IFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) GreaterVector4I(right Vector4i) bool {
  opPtr := ptrsForVector4i.operatorGreaterVector4IFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) GreaterEqualVector4I(right Vector4i) bool {
  opPtr := ptrsForVector4i.operatorGreaterEqualVector4IFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) AddVector4I(right Vector4i) Vector4i {
  opPtr := ptrsForVector4i.operatorAddVector4IFn
  ret := NewVector4i()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4i) SubtractVector4I(right Vector4i) Vector4i {
  opPtr := ptrsForVector4i.operatorSubtractVector4IFn
  ret := NewVector4i()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4i) MultiplyVector4I(right Vector4i) Vector4i {
  opPtr := ptrsForVector4i.operatorMultiplyVector4IFn
  ret := NewVector4i()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4i) DivideVector4I(right Vector4i) Vector4i {
  opPtr := ptrsForVector4i.operatorDivideVector4IFn
  ret := NewVector4i()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4i) ModuleVector4I(right Vector4i) Vector4i {
  opPtr := ptrsForVector4i.operatorModuleVector4IFn
  ret := NewVector4i()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4i) InDictionary(right Dictionary) bool {
  opPtr := ptrsForVector4i.operatorInDictionaryFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) InArray(right Array) bool {
  opPtr := ptrsForVector4i.operatorInArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

// Members

func (me *Vector4i) X() int64 {
  ret := NewInt()
  me.iface.CallPtrGetter(ensurePtr(ptrsForVector4i.memberxGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) SetX(value int64) {
  valueV := NewIntFromInt(value)
  defer valueV.Destroy()
  me.iface.CallPtrSetter(ensurePtr(ptrsForVector4i.memberxSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Vector4i) Y() int64 {
  ret := NewInt()
  me.iface.CallPtrGetter(ensurePtr(ptrsForVector4i.memberyGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) SetY(value int64) {
  valueV := NewIntFromInt(value)
  defer valueV.Destroy()
  me.iface.CallPtrSetter(ensurePtr(ptrsForVector4i.memberySetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Vector4i) Z() int64 {
  ret := NewInt()
  me.iface.CallPtrGetter(ensurePtr(ptrsForVector4i.memberzGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) SetZ(value int64) {
  valueV := NewIntFromInt(value)
  defer valueV.Destroy()
  me.iface.CallPtrSetter(ensurePtr(ptrsForVector4i.memberzSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Vector4i) W() int64 {
  ret := NewInt()
  me.iface.CallPtrGetter(ensurePtr(ptrsForVector4i.memberwGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4i) SetW(value int64) {
  valueV := NewIntFromInt(value)
  defer valueV.Destroy()
  me.iface.CallPtrSetter(ensurePtr(ptrsForVector4i.memberwSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}
