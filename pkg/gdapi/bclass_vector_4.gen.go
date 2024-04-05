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

type ptrsForVector4List struct {
  ctrFn gdc.PtrConstructor
  ctrFromVector4Fn gdc.PtrConstructor
  ctrFromVector4IFn gdc.PtrConstructor
  ctrFromFloat32Float32Float32Float32Fn gdc.PtrConstructor
  methodMinAxisIndexFn gdc.PtrBuiltInMethod
  methodMaxAxisIndexFn gdc.PtrBuiltInMethod
  methodLengthFn gdc.PtrBuiltInMethod
  methodLengthSquaredFn gdc.PtrBuiltInMethod
  methodAbsFn gdc.PtrBuiltInMethod
  methodSignFn gdc.PtrBuiltInMethod
  methodFloorFn gdc.PtrBuiltInMethod
  methodCeilFn gdc.PtrBuiltInMethod
  methodRoundFn gdc.PtrBuiltInMethod
  methodLerpFn gdc.PtrBuiltInMethod
  methodCubicInterpolateFn gdc.PtrBuiltInMethod
  methodCubicInterpolateInTimeFn gdc.PtrBuiltInMethod
  methodPosmodFn gdc.PtrBuiltInMethod
  methodPosmodvFn gdc.PtrBuiltInMethod
  methodSnappedFn gdc.PtrBuiltInMethod
  methodClampFn gdc.PtrBuiltInMethod
  methodNormalizedFn gdc.PtrBuiltInMethod
  methodIsNormalizedFn gdc.PtrBuiltInMethod
  methodDirectionToFn gdc.PtrBuiltInMethod
  methodDistanceToFn gdc.PtrBuiltInMethod
  methodDistanceSquaredToFn gdc.PtrBuiltInMethod
  methodDotFn gdc.PtrBuiltInMethod
  methodInverseFn gdc.PtrBuiltInMethod
  methodIsEqualApproxFn gdc.PtrBuiltInMethod
  methodIsZeroApproxFn gdc.PtrBuiltInMethod
  methodIsFiniteFn gdc.PtrBuiltInMethod
  operatorNegateFn gdc.PtrOperatorEvaluator
  operatorPositiveFn gdc.PtrOperatorEvaluator
  operatorNotFn gdc.PtrOperatorEvaluator
  operatorMultiplyIntFn gdc.PtrOperatorEvaluator
  operatorDivideIntFn gdc.PtrOperatorEvaluator
  operatorMultiplyFloat32Fn gdc.PtrOperatorEvaluator
  operatorDivideFloat32Fn gdc.PtrOperatorEvaluator
  operatorEqualVector4Fn gdc.PtrOperatorEvaluator
  operatorNotEqualVector4Fn gdc.PtrOperatorEvaluator
  operatorLessVector4Fn gdc.PtrOperatorEvaluator
  operatorLessEqualVector4Fn gdc.PtrOperatorEvaluator
  operatorGreaterVector4Fn gdc.PtrOperatorEvaluator
  operatorGreaterEqualVector4Fn gdc.PtrOperatorEvaluator
  operatorAddVector4Fn gdc.PtrOperatorEvaluator
  operatorSubtractVector4Fn gdc.PtrOperatorEvaluator
  operatorMultiplyVector4Fn gdc.PtrOperatorEvaluator
  operatorDivideVector4Fn gdc.PtrOperatorEvaluator
  operatorMultiplyProjectionFn gdc.PtrOperatorEvaluator
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

var ptrsForVector4 ptrsForVector4List

func initVector4Ptrs(iface gdc.Interface) {
  ptrsForVector4.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeVector4, 0))
  ptrsForVector4.ctrFromVector4Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeVector4, 1))
  ptrsForVector4.ctrFromVector4IFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeVector4, 2))
  ptrsForVector4.ctrFromFloat32Float32Float32Float32Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeVector4, 3))
  {
    methodName := StringNameFromStr("min_axis_index")
    defer methodName.Destroy()
    ptrsForVector4.methodMinAxisIndexFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, methodName.AsCPtr(), 3173160232))
  }
  {
    methodName := StringNameFromStr("max_axis_index")
    defer methodName.Destroy()
    ptrsForVector4.methodMaxAxisIndexFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, methodName.AsCPtr(), 3173160232))
  }
  {
    methodName := StringNameFromStr("length")
    defer methodName.Destroy()
    ptrsForVector4.methodLengthFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, methodName.AsCPtr(), 466405837))
  }
  {
    methodName := StringNameFromStr("length_squared")
    defer methodName.Destroy()
    ptrsForVector4.methodLengthSquaredFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, methodName.AsCPtr(), 466405837))
  }
  {
    methodName := StringNameFromStr("abs")
    defer methodName.Destroy()
    ptrsForVector4.methodAbsFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, methodName.AsCPtr(), 80860099))
  }
  {
    methodName := StringNameFromStr("sign")
    defer methodName.Destroy()
    ptrsForVector4.methodSignFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, methodName.AsCPtr(), 80860099))
  }
  {
    methodName := StringNameFromStr("floor")
    defer methodName.Destroy()
    ptrsForVector4.methodFloorFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, methodName.AsCPtr(), 80860099))
  }
  {
    methodName := StringNameFromStr("ceil")
    defer methodName.Destroy()
    ptrsForVector4.methodCeilFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, methodName.AsCPtr(), 80860099))
  }
  {
    methodName := StringNameFromStr("round")
    defer methodName.Destroy()
    ptrsForVector4.methodRoundFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, methodName.AsCPtr(), 80860099))
  }
  {
    methodName := StringNameFromStr("lerp")
    defer methodName.Destroy()
    ptrsForVector4.methodLerpFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, methodName.AsCPtr(), 2329757942))
  }
  {
    methodName := StringNameFromStr("cubic_interpolate")
    defer methodName.Destroy()
    ptrsForVector4.methodCubicInterpolateFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, methodName.AsCPtr(), 726768410))
  }
  {
    methodName := StringNameFromStr("cubic_interpolate_in_time")
    defer methodName.Destroy()
    ptrsForVector4.methodCubicInterpolateInTimeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, methodName.AsCPtr(), 681631873))
  }
  {
    methodName := StringNameFromStr("posmod")
    defer methodName.Destroy()
    ptrsForVector4.methodPosmodFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, methodName.AsCPtr(), 3129671720))
  }
  {
    methodName := StringNameFromStr("posmodv")
    defer methodName.Destroy()
    ptrsForVector4.methodPosmodvFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, methodName.AsCPtr(), 2031281584))
  }
  {
    methodName := StringNameFromStr("snapped")
    defer methodName.Destroy()
    ptrsForVector4.methodSnappedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, methodName.AsCPtr(), 2031281584))
  }
  {
    methodName := StringNameFromStr("clamp")
    defer methodName.Destroy()
    ptrsForVector4.methodClampFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, methodName.AsCPtr(), 823915692))
  }
  {
    methodName := StringNameFromStr("normalized")
    defer methodName.Destroy()
    ptrsForVector4.methodNormalizedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, methodName.AsCPtr(), 80860099))
  }
  {
    methodName := StringNameFromStr("is_normalized")
    defer methodName.Destroy()
    ptrsForVector4.methodIsNormalizedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, methodName.AsCPtr(), 3918633141))
  }
  {
    methodName := StringNameFromStr("direction_to")
    defer methodName.Destroy()
    ptrsForVector4.methodDirectionToFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, methodName.AsCPtr(), 2031281584))
  }
  {
    methodName := StringNameFromStr("distance_to")
    defer methodName.Destroy()
    ptrsForVector4.methodDistanceToFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, methodName.AsCPtr(), 3770801042))
  }
  {
    methodName := StringNameFromStr("distance_squared_to")
    defer methodName.Destroy()
    ptrsForVector4.methodDistanceSquaredToFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, methodName.AsCPtr(), 3770801042))
  }
  {
    methodName := StringNameFromStr("dot")
    defer methodName.Destroy()
    ptrsForVector4.methodDotFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, methodName.AsCPtr(), 3770801042))
  }
  {
    methodName := StringNameFromStr("inverse")
    defer methodName.Destroy()
    ptrsForVector4.methodInverseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, methodName.AsCPtr(), 80860099))
  }
  {
    methodName := StringNameFromStr("is_equal_approx")
    defer methodName.Destroy()
    ptrsForVector4.methodIsEqualApproxFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, methodName.AsCPtr(), 88913544))
  }
  {
    methodName := StringNameFromStr("is_zero_approx")
    defer methodName.Destroy()
    ptrsForVector4.methodIsZeroApproxFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, methodName.AsCPtr(), 3918633141))
  }
  {
    methodName := StringNameFromStr("is_finite")
    defer methodName.Destroy()
    ptrsForVector4.methodIsFiniteFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector4, methodName.AsCPtr(), 3918633141))
  }
  ptrsForVector4.operatorNegateFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNegate, gdc.VariantTypeVector4, gdc.VariantTypeNil))
  ptrsForVector4.operatorPositiveFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpPositive, gdc.VariantTypeVector4, gdc.VariantTypeNil))
  ptrsForVector4.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypeVector4, gdc.VariantTypeNil))
  ptrsForVector4.operatorMultiplyIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeVector4, gdc.VariantTypeInt))
  ptrsForVector4.operatorDivideIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeVector4, gdc.VariantTypeInt))
  ptrsForVector4.operatorMultiplyFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeVector4, gdc.VariantTypeFloat))
  ptrsForVector4.operatorDivideFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeVector4, gdc.VariantTypeFloat))
  ptrsForVector4.operatorEqualVector4Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeVector4, gdc.VariantTypeVector4))
  ptrsForVector4.operatorNotEqualVector4Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeVector4, gdc.VariantTypeVector4))
  ptrsForVector4.operatorLessVector4Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLess, gdc.VariantTypeVector4, gdc.VariantTypeVector4))
  ptrsForVector4.operatorLessEqualVector4Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLessEqual, gdc.VariantTypeVector4, gdc.VariantTypeVector4))
  ptrsForVector4.operatorGreaterVector4Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreater, gdc.VariantTypeVector4, gdc.VariantTypeVector4))
  ptrsForVector4.operatorGreaterEqualVector4Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreaterEqual, gdc.VariantTypeVector4, gdc.VariantTypeVector4))
  ptrsForVector4.operatorAddVector4Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, gdc.VariantTypeVector4, gdc.VariantTypeVector4))
  ptrsForVector4.operatorSubtractVector4Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpSubtract, gdc.VariantTypeVector4, gdc.VariantTypeVector4))
  ptrsForVector4.operatorMultiplyVector4Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeVector4, gdc.VariantTypeVector4))
  ptrsForVector4.operatorDivideVector4Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeVector4, gdc.VariantTypeVector4))
  ptrsForVector4.operatorMultiplyProjectionFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeVector4, gdc.VariantTypeProjection))
  ptrsForVector4.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeVector4, gdc.VariantTypeDictionary))
  ptrsForVector4.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeVector4, gdc.VariantTypeArray))
  {
    memberName := StringNameFromStr("x")
    defer memberName.Destroy()
    ptrsForVector4.memberxGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeVector4, memberName.AsCPtr()))
    ptrsForVector4.memberxSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeVector4, memberName.AsCPtr()))
  }
  {
    memberName := StringNameFromStr("y")
    defer memberName.Destroy()
    ptrsForVector4.memberyGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeVector4, memberName.AsCPtr()))
    ptrsForVector4.memberySetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeVector4, memberName.AsCPtr()))
  }
  {
    memberName := StringNameFromStr("z")
    defer memberName.Destroy()
    ptrsForVector4.memberzGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeVector4, memberName.AsCPtr()))
    ptrsForVector4.memberzSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeVector4, memberName.AsCPtr()))
  }
  {
    memberName := StringNameFromStr("w")
    defer memberName.Destroy()
    ptrsForVector4.memberwGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeVector4, memberName.AsCPtr()))
    ptrsForVector4.memberwSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeVector4, memberName.AsCPtr()))
  }
  ptrsForVector4.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypeVector4))
  ptrsForVector4.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypeVector4))
}

type Vector4 struct {
  data   *[classSizeVector4]byte
  iface  gdc.Interface
  pinner runtime.Pinner
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
func newVector4() *Vector4 {
  me := &Vector4{
    data:   new([classSizeVector4]byte),
    iface:  giface,
  }
  me.pinner.Pin(me)
  me.pinner.Pin(me.AsTypePtr())
  return me
}

func NewVector4() *Vector4 {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newVector4()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForVector4.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewVector4FromVector4(from Vector4, ) *Vector4 {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newVector4()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForVector4.ctrFromVector4Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewVector4FromVector4I(from Vector4i, ) *Vector4 {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newVector4()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForVector4.ctrFromVector4IFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewVector4FromFloat32Float32Float32Float32(x float64, y float64, z float64, w float64, ) *Vector4 {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&x)
  pinner.Pin(&y)
  pinner.Pin(&z)
  pinner.Pin(&w)
  me := newVector4()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForVector4.ctrFromFloat32Float32Float32Float32Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), gdc.ConstTypePtr(&z), gdc.ConstTypePtr(&w), }))
  return me
}

// Destructor
func (me *Vector4) Destroy() {
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsVector4() (*Vector4, error) {
	if me.Type() != gdc.VariantTypeVector4 {
		return nil, fmt.Errorf("variant is not a Vector4")
	}
  bclass := newVector4()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForVector4.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Vector4) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForVector4.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
  return va
}

// Pointers
func Vector4FromPtr(ptr gdc.ConstTypePtr) *Vector4 {
  me := newVector4()
  dataFromPtr(me.data[:], ptr)
  return me
}

func (me *Vector4) Type() gdc.VariantType {
  return gdc.VariantTypeVector4
}

func (me *Vector4) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Vector4) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Vector4) asUninitialized() gdc.UninitializedTypePtr {
  return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *Vector4) MinAxisIndex() int64 {
  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4.methodMinAxisIndexFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Vector4) MaxAxisIndex() int64 {
  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4.methodMaxAxisIndexFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Vector4) Length() float64 {
  ret := NewFloat()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4.methodLengthFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Vector4) LengthSquared() float64 {
  ret := NewFloat()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4.methodLengthSquaredFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Vector4) Abs() Vector4 {
  ret := NewVector4()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4.methodAbsFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Vector4) Sign() Vector4 {
  ret := NewVector4()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4.methodSignFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Vector4) Floor() Vector4 {
  ret := NewVector4()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4.methodFloorFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Vector4) Ceil() Vector4 {
  ret := NewVector4()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4.methodCeilFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Vector4) Round() Vector4 {
  ret := NewVector4()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4.methodRoundFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Vector4) Lerp(to Vector4, weight float64, ) Vector4 {
  ret := NewVector4()


  varg1 := NewFloatFromFloat32(weight)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4.methodLerpFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Vector4) CubicInterpolate(b Vector4, pre_a Vector4, post_b Vector4, weight float64, ) Vector4 {
  ret := NewVector4()




  varg3 := NewFloatFromFloat32(weight)
  defer varg3.Destroy()
  args := []gdc.ConstTypePtr{b.AsCTypePtr(), pre_a.AsCTypePtr(), post_b.AsCTypePtr(), varg3.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4.methodCubicInterpolateFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Vector4) CubicInterpolateInTime(b Vector4, pre_a Vector4, post_b Vector4, weight float64, b_t float64, pre_a_t float64, post_b_t float64, ) Vector4 {
  ret := NewVector4()




  varg3 := NewFloatFromFloat32(weight)
  defer varg3.Destroy()
  varg4 := NewFloatFromFloat32(b_t)
  defer varg4.Destroy()
  varg5 := NewFloatFromFloat32(pre_a_t)
  defer varg5.Destroy()
  varg6 := NewFloatFromFloat32(post_b_t)
  defer varg6.Destroy()
  args := []gdc.ConstTypePtr{b.AsCTypePtr(), pre_a.AsCTypePtr(), post_b.AsCTypePtr(), varg3.AsCTypePtr(), varg4.AsCTypePtr(), varg5.AsCTypePtr(), varg6.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4.methodCubicInterpolateInTimeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Vector4) Posmod(mod float64, ) Vector4 {
  ret := NewVector4()

  varg0 := NewFloatFromFloat32(mod)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4.methodPosmodFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Vector4) Posmodv(modv Vector4, ) Vector4 {
  ret := NewVector4()


  args := []gdc.ConstTypePtr{modv.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4.methodPosmodvFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Vector4) Snapped(step Vector4, ) Vector4 {
  ret := NewVector4()


  args := []gdc.ConstTypePtr{step.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4.methodSnappedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Vector4) Clamp(min Vector4, max Vector4, ) Vector4 {
  ret := NewVector4()



  args := []gdc.ConstTypePtr{min.AsCTypePtr(), max.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4.methodClampFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Vector4) Normalized() Vector4 {
  ret := NewVector4()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4.methodNormalizedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Vector4) IsNormalized() bool {
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4.methodIsNormalizedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Vector4) DirectionTo(to Vector4, ) Vector4 {
  ret := NewVector4()


  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4.methodDirectionToFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Vector4) DistanceTo(to Vector4, ) float64 {
  ret := NewFloat()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4.methodDistanceToFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Vector4) DistanceSquaredTo(to Vector4, ) float64 {
  ret := NewFloat()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4.methodDistanceSquaredToFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Vector4) Dot(with Vector4, ) float64 {
  ret := NewFloat()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{with.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4.methodDotFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Vector4) Inverse() Vector4 {
  ret := NewVector4()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4.methodInverseFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Vector4) IsEqualApprox(to Vector4, ) bool {
  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4.methodIsEqualApproxFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Vector4) IsZeroApprox() bool {
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4.methodIsZeroApproxFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Vector4) IsFinite() bool {
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector4.methodIsFiniteFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

// Operators

func (me *Vector4) EqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4) NotEqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4) Negate() Vector4 {
  opPtr := ptrsForVector4.operatorNegateFn
  ret := NewVector4()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
  return *ret
}

func (me *Vector4) Positive() Vector4 {
  opPtr := ptrsForVector4.operatorPositiveFn
  ret := NewVector4()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
  return *ret
}

func (me *Vector4) Not() bool {
  opPtr := ptrsForVector4.operatorNotFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4) MultiplyInt(rightArg int64) Vector4 {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  opPtr := ptrsForVector4.operatorMultiplyIntFn
  ret := NewVector4()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4) DivideInt(rightArg int64) Vector4 {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  opPtr := ptrsForVector4.operatorDivideIntFn
  ret := NewVector4()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4) MultiplyFloat32(rightArg float64) Vector4 {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  opPtr := ptrsForVector4.operatorMultiplyFloat32Fn
  ret := NewVector4()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4) DivideFloat32(rightArg float64) Vector4 {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  opPtr := ptrsForVector4.operatorDivideFloat32Fn
  ret := NewVector4()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4) EqualVector4(right Vector4) bool {
  opPtr := ptrsForVector4.operatorEqualVector4Fn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4) NotEqualVector4(right Vector4) bool {
  opPtr := ptrsForVector4.operatorNotEqualVector4Fn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4) LessVector4(right Vector4) bool {
  opPtr := ptrsForVector4.operatorLessVector4Fn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4) LessEqualVector4(right Vector4) bool {
  opPtr := ptrsForVector4.operatorLessEqualVector4Fn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4) GreaterVector4(right Vector4) bool {
  opPtr := ptrsForVector4.operatorGreaterVector4Fn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4) GreaterEqualVector4(right Vector4) bool {
  opPtr := ptrsForVector4.operatorGreaterEqualVector4Fn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4) AddVector4(right Vector4) Vector4 {
  opPtr := ptrsForVector4.operatorAddVector4Fn
  ret := NewVector4()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4) SubtractVector4(right Vector4) Vector4 {
  opPtr := ptrsForVector4.operatorSubtractVector4Fn
  ret := NewVector4()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4) MultiplyVector4(right Vector4) Vector4 {
  opPtr := ptrsForVector4.operatorMultiplyVector4Fn
  ret := NewVector4()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4) DivideVector4(right Vector4) Vector4 {
  opPtr := ptrsForVector4.operatorDivideVector4Fn
  ret := NewVector4()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4) MultiplyProjection(right Projection) Vector4 {
  opPtr := ptrsForVector4.operatorMultiplyProjectionFn
  ret := NewVector4()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Vector4) InDictionary(right Dictionary) bool {
  opPtr := ptrsForVector4.operatorInDictionaryFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4) InArray(right Array) bool {
  opPtr := ptrsForVector4.operatorInArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

// Members

func (me *Vector4) X() float64 {
  ret := NewFloat()
  me.iface.CallPtrGetter(ensurePtr(ptrsForVector4.memberxGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4) SetX(value float64) {
  valueV := NewFloatFromFloat32(value)
  defer valueV.Destroy()
  me.iface.CallPtrSetter(ensurePtr(ptrsForVector4.memberxSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Vector4) Y() float64 {
  ret := NewFloat()
  me.iface.CallPtrGetter(ensurePtr(ptrsForVector4.memberyGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4) SetY(value float64) {
  valueV := NewFloatFromFloat32(value)
  defer valueV.Destroy()
  me.iface.CallPtrSetter(ensurePtr(ptrsForVector4.memberySetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Vector4) Z() float64 {
  ret := NewFloat()
  me.iface.CallPtrGetter(ensurePtr(ptrsForVector4.memberzGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4) SetZ(value float64) {
  valueV := NewFloatFromFloat32(value)
  defer valueV.Destroy()
  me.iface.CallPtrSetter(ensurePtr(ptrsForVector4.memberzSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Vector4) W() float64 {
  ret := NewFloat()
  me.iface.CallPtrGetter(ensurePtr(ptrsForVector4.memberwGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Vector4) SetW(value float64) {
  valueV := NewFloatFromFloat32(value)
  defer valueV.Destroy()
  me.iface.CallPtrSetter(ensurePtr(ptrsForVector4.memberwSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}
