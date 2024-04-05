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

type ptrsForQuaternionList struct {
  ctrFn gdc.PtrConstructor
  ctrFromQuaternionFn gdc.PtrConstructor
  ctrFromBasisFn gdc.PtrConstructor
  ctrFromVector3Float32Fn gdc.PtrConstructor
  ctrFromVector3Vector3Fn gdc.PtrConstructor
  ctrFromFloat32Float32Float32Float32Fn gdc.PtrConstructor
  methodLengthFn gdc.PtrBuiltInMethod
  methodLengthSquaredFn gdc.PtrBuiltInMethod
  methodNormalizedFn gdc.PtrBuiltInMethod
  methodIsNormalizedFn gdc.PtrBuiltInMethod
  methodIsEqualApproxFn gdc.PtrBuiltInMethod
  methodIsFiniteFn gdc.PtrBuiltInMethod
  methodInverseFn gdc.PtrBuiltInMethod
  methodLogFn gdc.PtrBuiltInMethod
  methodExpFn gdc.PtrBuiltInMethod
  methodAngleToFn gdc.PtrBuiltInMethod
  methodDotFn gdc.PtrBuiltInMethod
  methodSlerpFn gdc.PtrBuiltInMethod
  methodSlerpniFn gdc.PtrBuiltInMethod
  methodSphericalCubicInterpolateFn gdc.PtrBuiltInMethod
  methodSphericalCubicInterpolateInTimeFn gdc.PtrBuiltInMethod
  methodGetEulerFn gdc.PtrBuiltInMethod
  methodFromEulerFn gdc.PtrBuiltInMethod
  methodGetAxisFn gdc.PtrBuiltInMethod
  methodGetAngleFn gdc.PtrBuiltInMethod
  operatorNegateFn gdc.PtrOperatorEvaluator
  operatorPositiveFn gdc.PtrOperatorEvaluator
  operatorNotFn gdc.PtrOperatorEvaluator
  operatorMultiplyIntFn gdc.PtrOperatorEvaluator
  operatorDivideIntFn gdc.PtrOperatorEvaluator
  operatorMultiplyFloat32Fn gdc.PtrOperatorEvaluator
  operatorDivideFloat32Fn gdc.PtrOperatorEvaluator
  operatorMultiplyVector3Fn gdc.PtrOperatorEvaluator
  operatorEqualQuaternionFn gdc.PtrOperatorEvaluator
  operatorNotEqualQuaternionFn gdc.PtrOperatorEvaluator
  operatorAddQuaternionFn gdc.PtrOperatorEvaluator
  operatorSubtractQuaternionFn gdc.PtrOperatorEvaluator
  operatorMultiplyQuaternionFn gdc.PtrOperatorEvaluator
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

var ptrsForQuaternion ptrsForQuaternionList

func initQuaternionPtrs(iface gdc.Interface) {
  ptrsForQuaternion.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeQuaternion, 0))
  ptrsForQuaternion.ctrFromQuaternionFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeQuaternion, 1))
  ptrsForQuaternion.ctrFromBasisFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeQuaternion, 2))
  ptrsForQuaternion.ctrFromVector3Float32Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeQuaternion, 3))
  ptrsForQuaternion.ctrFromVector3Vector3Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeQuaternion, 4))
  ptrsForQuaternion.ctrFromFloat32Float32Float32Float32Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeQuaternion, 5))
  {
    methodName := StringNameFromStr("length")
    defer methodName.Destroy()
    ptrsForQuaternion.methodLengthFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, methodName.AsCPtr(), 466405837))
  }
  {
    methodName := StringNameFromStr("length_squared")
    defer methodName.Destroy()
    ptrsForQuaternion.methodLengthSquaredFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, methodName.AsCPtr(), 466405837))
  }
  {
    methodName := StringNameFromStr("normalized")
    defer methodName.Destroy()
    ptrsForQuaternion.methodNormalizedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, methodName.AsCPtr(), 4274879941))
  }
  {
    methodName := StringNameFromStr("is_normalized")
    defer methodName.Destroy()
    ptrsForQuaternion.methodIsNormalizedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, methodName.AsCPtr(), 3918633141))
  }
  {
    methodName := StringNameFromStr("is_equal_approx")
    defer methodName.Destroy()
    ptrsForQuaternion.methodIsEqualApproxFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, methodName.AsCPtr(), 1682156903))
  }
  {
    methodName := StringNameFromStr("is_finite")
    defer methodName.Destroy()
    ptrsForQuaternion.methodIsFiniteFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, methodName.AsCPtr(), 3918633141))
  }
  {
    methodName := StringNameFromStr("inverse")
    defer methodName.Destroy()
    ptrsForQuaternion.methodInverseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, methodName.AsCPtr(), 4274879941))
  }
  {
    methodName := StringNameFromStr("log")
    defer methodName.Destroy()
    ptrsForQuaternion.methodLogFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, methodName.AsCPtr(), 4274879941))
  }
  {
    methodName := StringNameFromStr("exp")
    defer methodName.Destroy()
    ptrsForQuaternion.methodExpFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, methodName.AsCPtr(), 4274879941))
  }
  {
    methodName := StringNameFromStr("angle_to")
    defer methodName.Destroy()
    ptrsForQuaternion.methodAngleToFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, methodName.AsCPtr(), 3244682419))
  }
  {
    methodName := StringNameFromStr("dot")
    defer methodName.Destroy()
    ptrsForQuaternion.methodDotFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, methodName.AsCPtr(), 3244682419))
  }
  {
    methodName := StringNameFromStr("slerp")
    defer methodName.Destroy()
    ptrsForQuaternion.methodSlerpFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, methodName.AsCPtr(), 1773590316))
  }
  {
    methodName := StringNameFromStr("slerpni")
    defer methodName.Destroy()
    ptrsForQuaternion.methodSlerpniFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, methodName.AsCPtr(), 1773590316))
  }
  {
    methodName := StringNameFromStr("spherical_cubic_interpolate")
    defer methodName.Destroy()
    ptrsForQuaternion.methodSphericalCubicInterpolateFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, methodName.AsCPtr(), 2150967576))
  }
  {
    methodName := StringNameFromStr("spherical_cubic_interpolate_in_time")
    defer methodName.Destroy()
    ptrsForQuaternion.methodSphericalCubicInterpolateInTimeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, methodName.AsCPtr(), 1436023539))
  }
  {
    methodName := StringNameFromStr("get_euler")
    defer methodName.Destroy()
    ptrsForQuaternion.methodGetEulerFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, methodName.AsCPtr(), 1394941017))
  }
  {
    methodName := StringNameFromStr("from_euler")
    defer methodName.Destroy()
    ptrsForQuaternion.methodFromEulerFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, methodName.AsCPtr(), 4053467903))
  }
  {
    methodName := StringNameFromStr("get_axis")
    defer methodName.Destroy()
    ptrsForQuaternion.methodGetAxisFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, methodName.AsCPtr(), 1776574132))
  }
  {
    methodName := StringNameFromStr("get_angle")
    defer methodName.Destroy()
    ptrsForQuaternion.methodGetAngleFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeQuaternion, methodName.AsCPtr(), 466405837))
  }
  ptrsForQuaternion.operatorNegateFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNegate, gdc.VariantTypeQuaternion, gdc.VariantTypeNil))
  ptrsForQuaternion.operatorPositiveFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpPositive, gdc.VariantTypeQuaternion, gdc.VariantTypeNil))
  ptrsForQuaternion.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypeQuaternion, gdc.VariantTypeNil))
  ptrsForQuaternion.operatorMultiplyIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeQuaternion, gdc.VariantTypeInt))
  ptrsForQuaternion.operatorDivideIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeQuaternion, gdc.VariantTypeInt))
  ptrsForQuaternion.operatorMultiplyFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeQuaternion, gdc.VariantTypeFloat))
  ptrsForQuaternion.operatorDivideFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeQuaternion, gdc.VariantTypeFloat))
  ptrsForQuaternion.operatorMultiplyVector3Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeQuaternion, gdc.VariantTypeVector3))
  ptrsForQuaternion.operatorEqualQuaternionFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeQuaternion, gdc.VariantTypeQuaternion))
  ptrsForQuaternion.operatorNotEqualQuaternionFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeQuaternion, gdc.VariantTypeQuaternion))
  ptrsForQuaternion.operatorAddQuaternionFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, gdc.VariantTypeQuaternion, gdc.VariantTypeQuaternion))
  ptrsForQuaternion.operatorSubtractQuaternionFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpSubtract, gdc.VariantTypeQuaternion, gdc.VariantTypeQuaternion))
  ptrsForQuaternion.operatorMultiplyQuaternionFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeQuaternion, gdc.VariantTypeQuaternion))
  ptrsForQuaternion.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeQuaternion, gdc.VariantTypeDictionary))
  ptrsForQuaternion.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeQuaternion, gdc.VariantTypeArray))
  {
    memberName := StringNameFromStr("x")
    defer memberName.Destroy()
    ptrsForQuaternion.memberxGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeQuaternion, memberName.AsCPtr()))
    ptrsForQuaternion.memberxSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeQuaternion, memberName.AsCPtr()))
  }
  {
    memberName := StringNameFromStr("y")
    defer memberName.Destroy()
    ptrsForQuaternion.memberyGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeQuaternion, memberName.AsCPtr()))
    ptrsForQuaternion.memberySetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeQuaternion, memberName.AsCPtr()))
  }
  {
    memberName := StringNameFromStr("z")
    defer memberName.Destroy()
    ptrsForQuaternion.memberzGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeQuaternion, memberName.AsCPtr()))
    ptrsForQuaternion.memberzSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeQuaternion, memberName.AsCPtr()))
  }
  {
    memberName := StringNameFromStr("w")
    defer memberName.Destroy()
    ptrsForQuaternion.memberwGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeQuaternion, memberName.AsCPtr()))
    ptrsForQuaternion.memberwSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeQuaternion, memberName.AsCPtr()))
  }
  ptrsForQuaternion.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypeQuaternion))
  ptrsForQuaternion.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypeQuaternion))
}

type Quaternion struct {
  data   *[classSizeQuaternion]byte
  iface  gdc.Interface
  pinner runtime.Pinner
}

// Constants

var (
  QuaternionIdentity = "Quaternion(0, 0, 0, 1)" // TODO: construct correctly
)

// Enums

// Constructors
func newQuaternion() *Quaternion {
  me := &Quaternion{
    data:   new([classSizeQuaternion]byte),
    iface:  giface,
  }
  me.pinner.Pin(me)
  me.pinner.Pin(me.AsTypePtr())
  return me
}

func NewQuaternion() *Quaternion {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newQuaternion()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForQuaternion.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewQuaternionFromQuaternion(from Quaternion, ) *Quaternion {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newQuaternion()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForQuaternion.ctrFromQuaternionFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewQuaternionFromBasis(from Basis, ) *Quaternion {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newQuaternion()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForQuaternion.ctrFromBasisFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewQuaternionFromVector3Float32(axis Vector3, angle float64, ) *Quaternion {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&angle)
  me := newQuaternion()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForQuaternion.ctrFromVector3Float32Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{axis.AsCTypePtr(), gdc.ConstTypePtr(&angle), }))
  return me
}

func NewQuaternionFromVector3Vector3(arc_from Vector3, arc_to Vector3, ) *Quaternion {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newQuaternion()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForQuaternion.ctrFromVector3Vector3Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{arc_from.AsCTypePtr(), arc_to.AsCTypePtr(), }))
  return me
}

func NewQuaternionFromFloat32Float32Float32Float32(x float64, y float64, z float64, w float64, ) *Quaternion {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&x)
  pinner.Pin(&y)
  pinner.Pin(&z)
  pinner.Pin(&w)
  me := newQuaternion()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForQuaternion.ctrFromFloat32Float32Float32Float32Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), gdc.ConstTypePtr(&z), gdc.ConstTypePtr(&w), }))
  return me
}

// Destructor
func (me *Quaternion) Destroy() {
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsQuaternion() (*Quaternion, error) {
	if me.Type() != gdc.VariantTypeQuaternion {
		return nil, fmt.Errorf("variant is not a Quaternion")
	}
  bclass := newQuaternion()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForQuaternion.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Quaternion) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForQuaternion.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
  return va
}

// Pointers
func QuaternionFromPtr(ptr gdc.ConstTypePtr) *Quaternion {
  me := newQuaternion()
  dataFromPtr(me.data[:], ptr)
  return me
}

func (me *Quaternion) Type() gdc.VariantType {
  return gdc.VariantTypeQuaternion
}

func (me *Quaternion) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Quaternion) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Quaternion) asUninitialized() gdc.UninitializedTypePtr {
  return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *Quaternion) Length() float64 {
  ret := NewFloat()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForQuaternion.methodLengthFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Quaternion) LengthSquared() float64 {
  ret := NewFloat()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForQuaternion.methodLengthSquaredFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Quaternion) Normalized() Quaternion {
  ret := NewQuaternion()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForQuaternion.methodNormalizedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Quaternion) IsNormalized() bool {
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForQuaternion.methodIsNormalizedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Quaternion) IsEqualApprox(to Quaternion, ) bool {
  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForQuaternion.methodIsEqualApproxFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Quaternion) IsFinite() bool {
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForQuaternion.methodIsFiniteFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Quaternion) Inverse() Quaternion {
  ret := NewQuaternion()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForQuaternion.methodInverseFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Quaternion) Log() Quaternion {
  ret := NewQuaternion()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForQuaternion.methodLogFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Quaternion) Exp() Quaternion {
  ret := NewQuaternion()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForQuaternion.methodExpFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Quaternion) AngleTo(to Quaternion, ) float64 {
  ret := NewFloat()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForQuaternion.methodAngleToFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Quaternion) Dot(with Quaternion, ) float64 {
  ret := NewFloat()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{with.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForQuaternion.methodDotFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Quaternion) Slerp(to Quaternion, weight float64, ) Quaternion {
  ret := NewQuaternion()


  varg1 := NewFloatFromFloat32(weight)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForQuaternion.methodSlerpFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Quaternion) Slerpni(to Quaternion, weight float64, ) Quaternion {
  ret := NewQuaternion()


  varg1 := NewFloatFromFloat32(weight)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForQuaternion.methodSlerpniFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Quaternion) SphericalCubicInterpolate(b Quaternion, pre_a Quaternion, post_b Quaternion, weight float64, ) Quaternion {
  ret := NewQuaternion()




  varg3 := NewFloatFromFloat32(weight)
  defer varg3.Destroy()
  args := []gdc.ConstTypePtr{b.AsCTypePtr(), pre_a.AsCTypePtr(), post_b.AsCTypePtr(), varg3.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForQuaternion.methodSphericalCubicInterpolateFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Quaternion) SphericalCubicInterpolateInTime(b Quaternion, pre_a Quaternion, post_b Quaternion, weight float64, b_t float64, pre_a_t float64, post_b_t float64, ) Quaternion {
  ret := NewQuaternion()




  varg3 := NewFloatFromFloat32(weight)
  defer varg3.Destroy()
  varg4 := NewFloatFromFloat32(b_t)
  defer varg4.Destroy()
  varg5 := NewFloatFromFloat32(pre_a_t)
  defer varg5.Destroy()
  varg6 := NewFloatFromFloat32(post_b_t)
  defer varg6.Destroy()
  args := []gdc.ConstTypePtr{b.AsCTypePtr(), pre_a.AsCTypePtr(), post_b.AsCTypePtr(), varg3.AsCTypePtr(), varg4.AsCTypePtr(), varg5.AsCTypePtr(), varg6.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForQuaternion.methodSphericalCubicInterpolateInTimeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Quaternion) GetEuler(order int64, ) Vector3 {
  ret := NewVector3()

  varg0 := NewIntFromInt(order)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForQuaternion.methodGetEulerFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func QuaternionFromEuler(euler Vector3, ) Quaternion {
  ret := NewQuaternion()


  args := []gdc.ConstTypePtr{euler.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForQuaternion.methodFromEulerFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Quaternion) GetAxis() Vector3 {
  ret := NewVector3()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForQuaternion.methodGetAxisFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Quaternion) GetAngle() float64 {
  ret := NewFloat()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForQuaternion.methodGetAngleFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

// Operators

func (me *Quaternion) EqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Quaternion) NotEqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Quaternion) Negate() Quaternion {
  opPtr := ptrsForQuaternion.operatorNegateFn
  ret := NewQuaternion()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
  return *ret
}

func (me *Quaternion) Positive() Quaternion {
  opPtr := ptrsForQuaternion.operatorPositiveFn
  ret := NewQuaternion()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
  return *ret
}

func (me *Quaternion) Not() bool {
  opPtr := ptrsForQuaternion.operatorNotFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
  return ret.Get()
}

func (me *Quaternion) MultiplyInt(rightArg int64) Quaternion {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  opPtr := ptrsForQuaternion.operatorMultiplyIntFn
  ret := NewQuaternion()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Quaternion) DivideInt(rightArg int64) Quaternion {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  opPtr := ptrsForQuaternion.operatorDivideIntFn
  ret := NewQuaternion()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Quaternion) MultiplyFloat32(rightArg float64) Quaternion {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  opPtr := ptrsForQuaternion.operatorMultiplyFloat32Fn
  ret := NewQuaternion()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Quaternion) DivideFloat32(rightArg float64) Quaternion {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  opPtr := ptrsForQuaternion.operatorDivideFloat32Fn
  ret := NewQuaternion()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Quaternion) MultiplyVector3(right Vector3) Vector3 {
  opPtr := ptrsForQuaternion.operatorMultiplyVector3Fn
  ret := NewVector3()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Quaternion) EqualQuaternion(right Quaternion) bool {
  opPtr := ptrsForQuaternion.operatorEqualQuaternionFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Quaternion) NotEqualQuaternion(right Quaternion) bool {
  opPtr := ptrsForQuaternion.operatorNotEqualQuaternionFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Quaternion) AddQuaternion(right Quaternion) Quaternion {
  opPtr := ptrsForQuaternion.operatorAddQuaternionFn
  ret := NewQuaternion()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Quaternion) SubtractQuaternion(right Quaternion) Quaternion {
  opPtr := ptrsForQuaternion.operatorSubtractQuaternionFn
  ret := NewQuaternion()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Quaternion) MultiplyQuaternion(right Quaternion) Quaternion {
  opPtr := ptrsForQuaternion.operatorMultiplyQuaternionFn
  ret := NewQuaternion()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Quaternion) InDictionary(right Dictionary) bool {
  opPtr := ptrsForQuaternion.operatorInDictionaryFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Quaternion) InArray(right Array) bool {
  opPtr := ptrsForQuaternion.operatorInArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

// Members

func (me *Quaternion) X() float64 {
  ret := NewFloat()
  me.iface.CallPtrGetter(ensurePtr(ptrsForQuaternion.memberxGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Quaternion) SetX(value float64) {
  valueV := NewFloatFromFloat32(value)
  defer valueV.Destroy()
  me.iface.CallPtrSetter(ensurePtr(ptrsForQuaternion.memberxSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Quaternion) Y() float64 {
  ret := NewFloat()
  me.iface.CallPtrGetter(ensurePtr(ptrsForQuaternion.memberyGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Quaternion) SetY(value float64) {
  valueV := NewFloatFromFloat32(value)
  defer valueV.Destroy()
  me.iface.CallPtrSetter(ensurePtr(ptrsForQuaternion.memberySetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Quaternion) Z() float64 {
  ret := NewFloat()
  me.iface.CallPtrGetter(ensurePtr(ptrsForQuaternion.memberzGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Quaternion) SetZ(value float64) {
  valueV := NewFloatFromFloat32(value)
  defer valueV.Destroy()
  me.iface.CallPtrSetter(ensurePtr(ptrsForQuaternion.memberzSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Quaternion) W() float64 {
  ret := NewFloat()
  me.iface.CallPtrGetter(ensurePtr(ptrsForQuaternion.memberwGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Quaternion) SetW(value float64) {
  valueV := NewFloatFromFloat32(value)
  defer valueV.Destroy()
  me.iface.CallPtrSetter(ensurePtr(ptrsForQuaternion.memberwSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}
