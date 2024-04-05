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

type ptrsForBasisList struct {
  ctrFn gdc.PtrConstructor
  ctrFromBasisFn gdc.PtrConstructor
  ctrFromQuaternionFn gdc.PtrConstructor
  ctrFromVector3Float32Fn gdc.PtrConstructor
  ctrFromVector3Vector3Vector3Fn gdc.PtrConstructor
  methodInverseFn gdc.PtrBuiltInMethod
  methodTransposedFn gdc.PtrBuiltInMethod
  methodOrthonormalizedFn gdc.PtrBuiltInMethod
  methodDeterminantFn gdc.PtrBuiltInMethod
  methodRotatedFn gdc.PtrBuiltInMethod
  methodScaledFn gdc.PtrBuiltInMethod
  methodGetScaleFn gdc.PtrBuiltInMethod
  methodGetEulerFn gdc.PtrBuiltInMethod
  methodTdotxFn gdc.PtrBuiltInMethod
  methodTdotyFn gdc.PtrBuiltInMethod
  methodTdotzFn gdc.PtrBuiltInMethod
  methodSlerpFn gdc.PtrBuiltInMethod
  methodIsConformalFn gdc.PtrBuiltInMethod
  methodIsEqualApproxFn gdc.PtrBuiltInMethod
  methodIsFiniteFn gdc.PtrBuiltInMethod
  methodGetRotationQuaternionFn gdc.PtrBuiltInMethod
  methodLookingAtFn gdc.PtrBuiltInMethod
  methodFromScaleFn gdc.PtrBuiltInMethod
  methodFromEulerFn gdc.PtrBuiltInMethod
  operatorNotFn gdc.PtrOperatorEvaluator
  operatorMultiplyIntFn gdc.PtrOperatorEvaluator
  operatorMultiplyFloat32Fn gdc.PtrOperatorEvaluator
  operatorMultiplyVector3Fn gdc.PtrOperatorEvaluator
  operatorEqualBasisFn gdc.PtrOperatorEvaluator
  operatorNotEqualBasisFn gdc.PtrOperatorEvaluator
  operatorMultiplyBasisFn gdc.PtrOperatorEvaluator
  operatorInDictionaryFn gdc.PtrOperatorEvaluator
  operatorInArrayFn gdc.PtrOperatorEvaluator
  memberxGetterFn gdc.PtrGetter
  memberxSetterFn gdc.PtrSetter
  memberyGetterFn gdc.PtrGetter
  memberySetterFn gdc.PtrSetter
  memberzGetterFn gdc.PtrGetter
  memberzSetterFn gdc.PtrSetter
  toVariantFn gdc.TypeFromVariantConstructorFunc
  fromVariantFn gdc.VariantFromTypeConstructorFunc
}

var ptrsForBasis ptrsForBasisList

func initBasisPtrs(iface gdc.Interface) {
  ptrsForBasis.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeBasis, 0))
  ptrsForBasis.ctrFromBasisFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeBasis, 1))
  ptrsForBasis.ctrFromQuaternionFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeBasis, 2))
  ptrsForBasis.ctrFromVector3Float32Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeBasis, 3))
  ptrsForBasis.ctrFromVector3Vector3Vector3Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeBasis, 4))
  {
    methodName := StringNameFromStr("inverse")
    defer methodName.Destroy()
    ptrsForBasis.methodInverseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, methodName.AsCPtr(), 594669093))
  }
  {
    methodName := StringNameFromStr("transposed")
    defer methodName.Destroy()
    ptrsForBasis.methodTransposedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, methodName.AsCPtr(), 594669093))
  }
  {
    methodName := StringNameFromStr("orthonormalized")
    defer methodName.Destroy()
    ptrsForBasis.methodOrthonormalizedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, methodName.AsCPtr(), 594669093))
  }
  {
    methodName := StringNameFromStr("determinant")
    defer methodName.Destroy()
    ptrsForBasis.methodDeterminantFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, methodName.AsCPtr(), 466405837))
  }
  {
    methodName := StringNameFromStr("rotated")
    defer methodName.Destroy()
    ptrsForBasis.methodRotatedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, methodName.AsCPtr(), 1998708965))
  }
  {
    methodName := StringNameFromStr("scaled")
    defer methodName.Destroy()
    ptrsForBasis.methodScaledFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, methodName.AsCPtr(), 3934786792))
  }
  {
    methodName := StringNameFromStr("get_scale")
    defer methodName.Destroy()
    ptrsForBasis.methodGetScaleFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, methodName.AsCPtr(), 1776574132))
  }
  {
    methodName := StringNameFromStr("get_euler")
    defer methodName.Destroy()
    ptrsForBasis.methodGetEulerFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, methodName.AsCPtr(), 1394941017))
  }
  {
    methodName := StringNameFromStr("tdotx")
    defer methodName.Destroy()
    ptrsForBasis.methodTdotxFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, methodName.AsCPtr(), 1047977935))
  }
  {
    methodName := StringNameFromStr("tdoty")
    defer methodName.Destroy()
    ptrsForBasis.methodTdotyFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, methodName.AsCPtr(), 1047977935))
  }
  {
    methodName := StringNameFromStr("tdotz")
    defer methodName.Destroy()
    ptrsForBasis.methodTdotzFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, methodName.AsCPtr(), 1047977935))
  }
  {
    methodName := StringNameFromStr("slerp")
    defer methodName.Destroy()
    ptrsForBasis.methodSlerpFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, methodName.AsCPtr(), 3118673011))
  }
  {
    methodName := StringNameFromStr("is_conformal")
    defer methodName.Destroy()
    ptrsForBasis.methodIsConformalFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, methodName.AsCPtr(), 3918633141))
  }
  {
    methodName := StringNameFromStr("is_equal_approx")
    defer methodName.Destroy()
    ptrsForBasis.methodIsEqualApproxFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, methodName.AsCPtr(), 3165333982))
  }
  {
    methodName := StringNameFromStr("is_finite")
    defer methodName.Destroy()
    ptrsForBasis.methodIsFiniteFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, methodName.AsCPtr(), 3918633141))
  }
  {
    methodName := StringNameFromStr("get_rotation_quaternion")
    defer methodName.Destroy()
    ptrsForBasis.methodGetRotationQuaternionFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, methodName.AsCPtr(), 4274879941))
  }
  {
    methodName := StringNameFromStr("looking_at")
    defer methodName.Destroy()
    ptrsForBasis.methodLookingAtFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, methodName.AsCPtr(), 3728732505))
  }
  {
    methodName := StringNameFromStr("from_scale")
    defer methodName.Destroy()
    ptrsForBasis.methodFromScaleFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, methodName.AsCPtr(), 3703240166))
  }
  {
    methodName := StringNameFromStr("from_euler")
    defer methodName.Destroy()
    ptrsForBasis.methodFromEulerFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeBasis, methodName.AsCPtr(), 2802321791))
  }
  ptrsForBasis.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypeBasis, gdc.VariantTypeNil))
  ptrsForBasis.operatorMultiplyIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeBasis, gdc.VariantTypeInt))
  ptrsForBasis.operatorMultiplyFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeBasis, gdc.VariantTypeFloat))
  ptrsForBasis.operatorMultiplyVector3Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeBasis, gdc.VariantTypeVector3))
  ptrsForBasis.operatorEqualBasisFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeBasis, gdc.VariantTypeBasis))
  ptrsForBasis.operatorNotEqualBasisFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeBasis, gdc.VariantTypeBasis))
  ptrsForBasis.operatorMultiplyBasisFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeBasis, gdc.VariantTypeBasis))
  ptrsForBasis.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeBasis, gdc.VariantTypeDictionary))
  ptrsForBasis.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeBasis, gdc.VariantTypeArray))
  {
    memberName := StringNameFromStr("x")
    defer memberName.Destroy()
    ptrsForBasis.memberxGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeBasis, memberName.AsCPtr()))
    ptrsForBasis.memberxSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeBasis, memberName.AsCPtr()))
  }
  {
    memberName := StringNameFromStr("y")
    defer memberName.Destroy()
    ptrsForBasis.memberyGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeBasis, memberName.AsCPtr()))
    ptrsForBasis.memberySetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeBasis, memberName.AsCPtr()))
  }
  {
    memberName := StringNameFromStr("z")
    defer memberName.Destroy()
    ptrsForBasis.memberzGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeBasis, memberName.AsCPtr()))
    ptrsForBasis.memberzSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeBasis, memberName.AsCPtr()))
  }
  ptrsForBasis.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypeBasis))
  ptrsForBasis.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypeBasis))
}

type Basis struct {
  data   *[classSizeBasis]byte
  iface  gdc.Interface
  pinner runtime.Pinner
}

// Constants

var (
  BasisIdentity = "Basis(1, 0, 0, 0, 1, 0, 0, 0, 1)" // TODO: construct correctly
  BasisFlipX = "Basis(-1, 0, 0, 0, 1, 0, 0, 0, 1)" // TODO: construct correctly
  BasisFlipY = "Basis(1, 0, 0, 0, -1, 0, 0, 0, 1)" // TODO: construct correctly
  BasisFlipZ = "Basis(1, 0, 0, 0, 1, 0, 0, 0, -1)" // TODO: construct correctly
)

// Enums

// Constructors
func newBasis() *Basis {
  me := &Basis{
    data:   new([classSizeBasis]byte),
    iface:  giface,
  }
  me.pinner.Pin(me)
  me.pinner.Pin(me.AsTypePtr())
  return me
}

func NewBasis() *Basis {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newBasis()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForBasis.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewBasisFromBasis(from Basis, ) *Basis {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newBasis()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForBasis.ctrFromBasisFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewBasisFromQuaternion(from Quaternion, ) *Basis {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newBasis()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForBasis.ctrFromQuaternionFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewBasisFromVector3Float32(axis Vector3, angle float64, ) *Basis {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&angle)
  me := newBasis()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForBasis.ctrFromVector3Float32Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{axis.AsCTypePtr(), gdc.ConstTypePtr(&angle), }))
  return me
}

func NewBasisFromVector3Vector3Vector3(x_axis Vector3, y_axis Vector3, z_axis Vector3, ) *Basis {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newBasis()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForBasis.ctrFromVector3Vector3Vector3Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{x_axis.AsCTypePtr(), y_axis.AsCTypePtr(), z_axis.AsCTypePtr(), }))
  return me
}

// Destructor
func (me *Basis) Destroy() {
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsBasis() (*Basis, error) {
	if me.Type() != gdc.VariantTypeBasis {
		return nil, fmt.Errorf("variant is not a Basis")
	}
  bclass := newBasis()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForBasis.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Basis) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForBasis.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
  return va
}

// Pointers
func BasisFromPtr(ptr gdc.ConstTypePtr) *Basis {
  me := newBasis()
  dataFromPtr(me.data[:], ptr)
  return me
}

func (me *Basis) Type() gdc.VariantType {
  return gdc.VariantTypeBasis
}

func (me *Basis) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Basis) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Basis) asUninitialized() gdc.UninitializedTypePtr {
  return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *Basis) Inverse() Basis {
  ret := NewBasis()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForBasis.methodInverseFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Basis) Transposed() Basis {
  ret := NewBasis()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForBasis.methodTransposedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Basis) Orthonormalized() Basis {
  ret := NewBasis()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForBasis.methodOrthonormalizedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Basis) Determinant() float64 {
  ret := NewFloat()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForBasis.methodDeterminantFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Basis) Rotated(axis Vector3, angle float64, ) Basis {
  ret := NewBasis()


  varg1 := NewFloatFromFloat32(angle)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{axis.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForBasis.methodRotatedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Basis) Scaled(scale Vector3, ) Basis {
  ret := NewBasis()


  args := []gdc.ConstTypePtr{scale.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForBasis.methodScaledFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Basis) GetScale() Vector3 {
  ret := NewVector3()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForBasis.methodGetScaleFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Basis) GetEuler(order int64, ) Vector3 {
  ret := NewVector3()

  varg0 := NewIntFromInt(order)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForBasis.methodGetEulerFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Basis) Tdotx(with Vector3, ) float64 {
  ret := NewFloat()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{with.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForBasis.methodTdotxFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Basis) Tdoty(with Vector3, ) float64 {
  ret := NewFloat()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{with.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForBasis.methodTdotyFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Basis) Tdotz(with Vector3, ) float64 {
  ret := NewFloat()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{with.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForBasis.methodTdotzFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Basis) Slerp(to Basis, weight float64, ) Basis {
  ret := NewBasis()


  varg1 := NewFloatFromFloat32(weight)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForBasis.methodSlerpFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Basis) IsConformal() bool {
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForBasis.methodIsConformalFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Basis) IsEqualApprox(b Basis, ) bool {
  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{b.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForBasis.methodIsEqualApproxFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Basis) IsFinite() bool {
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForBasis.methodIsFiniteFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Basis) GetRotationQuaternion() Quaternion {
  ret := NewQuaternion()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForBasis.methodGetRotationQuaternionFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func BasisLookingAt(target Vector3, up Vector3, use_model_front bool, ) Basis {
  ret := NewBasis()



  varg2 := NewBoolFromBool(use_model_front)
  defer varg2.Destroy()
  args := []gdc.ConstTypePtr{target.AsCTypePtr(), up.AsCTypePtr(), varg2.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForBasis.methodLookingAtFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func BasisFromScale(scale Vector3, ) Basis {
  ret := NewBasis()


  args := []gdc.ConstTypePtr{scale.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForBasis.methodFromScaleFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func BasisFromEuler(euler Vector3, order int64, ) Basis {
  ret := NewBasis()


  varg1 := NewIntFromInt(order)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{euler.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForBasis.methodFromEulerFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

// Operators

func (me *Basis) EqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Basis) NotEqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Basis) Not() bool {
  opPtr := ptrsForBasis.operatorNotFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
  return ret.Get()
}

func (me *Basis) MultiplyInt(rightArg int64) Basis {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  opPtr := ptrsForBasis.operatorMultiplyIntFn
  ret := NewBasis()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Basis) MultiplyFloat32(rightArg float64) Basis {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  opPtr := ptrsForBasis.operatorMultiplyFloat32Fn
  ret := NewBasis()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Basis) MultiplyVector3(right Vector3) Vector3 {
  opPtr := ptrsForBasis.operatorMultiplyVector3Fn
  ret := NewVector3()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Basis) EqualBasis(right Basis) bool {
  opPtr := ptrsForBasis.operatorEqualBasisFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Basis) NotEqualBasis(right Basis) bool {
  opPtr := ptrsForBasis.operatorNotEqualBasisFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Basis) MultiplyBasis(right Basis) Basis {
  opPtr := ptrsForBasis.operatorMultiplyBasisFn
  ret := NewBasis()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Basis) InDictionary(right Dictionary) bool {
  opPtr := ptrsForBasis.operatorInDictionaryFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Basis) InArray(right Array) bool {
  opPtr := ptrsForBasis.operatorInArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

// Members

func (me *Basis) X() Vector3 {
  ret := NewVector3()
  me.iface.CallPtrGetter(ensurePtr(ptrsForBasis.memberxGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Basis) SetX(value Vector3) {
  valueV := value
  me.iface.CallPtrSetter(ensurePtr(ptrsForBasis.memberxSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Basis) Y() Vector3 {
  ret := NewVector3()
  me.iface.CallPtrGetter(ensurePtr(ptrsForBasis.memberyGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Basis) SetY(value Vector3) {
  valueV := value
  me.iface.CallPtrSetter(ensurePtr(ptrsForBasis.memberySetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Basis) Z() Vector3 {
  ret := NewVector3()
  me.iface.CallPtrGetter(ensurePtr(ptrsForBasis.memberzGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Basis) SetZ(value Vector3) {
  valueV := value
  me.iface.CallPtrSetter(ensurePtr(ptrsForBasis.memberzSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}
