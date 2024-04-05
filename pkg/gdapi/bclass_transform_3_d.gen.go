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

type ptrsForTransform3DList struct {
  ctrFn gdc.PtrConstructor
  ctrFromTransform3DFn gdc.PtrConstructor
  ctrFromBasisVector3Fn gdc.PtrConstructor
  ctrFromVector3Vector3Vector3Vector3Fn gdc.PtrConstructor
  ctrFromProjectionFn gdc.PtrConstructor
  methodInverseFn gdc.PtrBuiltInMethod
  methodAffineInverseFn gdc.PtrBuiltInMethod
  methodOrthonormalizedFn gdc.PtrBuiltInMethod
  methodRotatedFn gdc.PtrBuiltInMethod
  methodRotatedLocalFn gdc.PtrBuiltInMethod
  methodScaledFn gdc.PtrBuiltInMethod
  methodScaledLocalFn gdc.PtrBuiltInMethod
  methodTranslatedFn gdc.PtrBuiltInMethod
  methodTranslatedLocalFn gdc.PtrBuiltInMethod
  methodLookingAtFn gdc.PtrBuiltInMethod
  methodInterpolateWithFn gdc.PtrBuiltInMethod
  methodIsEqualApproxFn gdc.PtrBuiltInMethod
  methodIsFiniteFn gdc.PtrBuiltInMethod
  operatorNotFn gdc.PtrOperatorEvaluator
  operatorMultiplyIntFn gdc.PtrOperatorEvaluator
  operatorMultiplyFloat32Fn gdc.PtrOperatorEvaluator
  operatorMultiplyVector3Fn gdc.PtrOperatorEvaluator
  operatorMultiplyPlaneFn gdc.PtrOperatorEvaluator
  operatorMultiplyAABBFn gdc.PtrOperatorEvaluator
  operatorEqualTransform3DFn gdc.PtrOperatorEvaluator
  operatorNotEqualTransform3DFn gdc.PtrOperatorEvaluator
  operatorMultiplyTransform3DFn gdc.PtrOperatorEvaluator
  operatorInDictionaryFn gdc.PtrOperatorEvaluator
  operatorInArrayFn gdc.PtrOperatorEvaluator
  operatorMultiplyPackedVector3ArrayFn gdc.PtrOperatorEvaluator
  memberbasisGetterFn gdc.PtrGetter
  memberbasisSetterFn gdc.PtrSetter
  memberoriginGetterFn gdc.PtrGetter
  memberoriginSetterFn gdc.PtrSetter
  toVariantFn gdc.TypeFromVariantConstructorFunc
  fromVariantFn gdc.VariantFromTypeConstructorFunc
}

var ptrsForTransform3D ptrsForTransform3DList

func initTransform3DPtrs(iface gdc.Interface) {
  ptrsForTransform3D.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeTransform3D, 0))
  ptrsForTransform3D.ctrFromTransform3DFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeTransform3D, 1))
  ptrsForTransform3D.ctrFromBasisVector3Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeTransform3D, 2))
  ptrsForTransform3D.ctrFromVector3Vector3Vector3Vector3Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeTransform3D, 3))
  ptrsForTransform3D.ctrFromProjectionFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeTransform3D, 4))
  {
    methodName := StringNameFromStr("inverse")
    defer methodName.Destroy()
    ptrsForTransform3D.methodInverseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform3D, methodName.AsCPtr(), 3816817146))
  }
  {
    methodName := StringNameFromStr("affine_inverse")
    defer methodName.Destroy()
    ptrsForTransform3D.methodAffineInverseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform3D, methodName.AsCPtr(), 3816817146))
  }
  {
    methodName := StringNameFromStr("orthonormalized")
    defer methodName.Destroy()
    ptrsForTransform3D.methodOrthonormalizedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform3D, methodName.AsCPtr(), 3816817146))
  }
  {
    methodName := StringNameFromStr("rotated")
    defer methodName.Destroy()
    ptrsForTransform3D.methodRotatedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform3D, methodName.AsCPtr(), 1563203923))
  }
  {
    methodName := StringNameFromStr("rotated_local")
    defer methodName.Destroy()
    ptrsForTransform3D.methodRotatedLocalFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform3D, methodName.AsCPtr(), 1563203923))
  }
  {
    methodName := StringNameFromStr("scaled")
    defer methodName.Destroy()
    ptrsForTransform3D.methodScaledFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform3D, methodName.AsCPtr(), 1405596198))
  }
  {
    methodName := StringNameFromStr("scaled_local")
    defer methodName.Destroy()
    ptrsForTransform3D.methodScaledLocalFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform3D, methodName.AsCPtr(), 1405596198))
  }
  {
    methodName := StringNameFromStr("translated")
    defer methodName.Destroy()
    ptrsForTransform3D.methodTranslatedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform3D, methodName.AsCPtr(), 1405596198))
  }
  {
    methodName := StringNameFromStr("translated_local")
    defer methodName.Destroy()
    ptrsForTransform3D.methodTranslatedLocalFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform3D, methodName.AsCPtr(), 1405596198))
  }
  {
    methodName := StringNameFromStr("looking_at")
    defer methodName.Destroy()
    ptrsForTransform3D.methodLookingAtFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform3D, methodName.AsCPtr(), 90889270))
  }
  {
    methodName := StringNameFromStr("interpolate_with")
    defer methodName.Destroy()
    ptrsForTransform3D.methodInterpolateWithFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform3D, methodName.AsCPtr(), 1786453358))
  }
  {
    methodName := StringNameFromStr("is_equal_approx")
    defer methodName.Destroy()
    ptrsForTransform3D.methodIsEqualApproxFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform3D, methodName.AsCPtr(), 696001652))
  }
  {
    methodName := StringNameFromStr("is_finite")
    defer methodName.Destroy()
    ptrsForTransform3D.methodIsFiniteFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform3D, methodName.AsCPtr(), 3918633141))
  }
  ptrsForTransform3D.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypeTransform3D, gdc.VariantTypeNil))
  ptrsForTransform3D.operatorMultiplyIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeTransform3D, gdc.VariantTypeInt))
  ptrsForTransform3D.operatorMultiplyFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeTransform3D, gdc.VariantTypeFloat))
  ptrsForTransform3D.operatorMultiplyVector3Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeTransform3D, gdc.VariantTypeVector3))
  ptrsForTransform3D.operatorMultiplyPlaneFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeTransform3D, gdc.VariantTypePlane))
  ptrsForTransform3D.operatorMultiplyAABBFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeTransform3D, gdc.VariantTypeAABB))
  ptrsForTransform3D.operatorEqualTransform3DFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeTransform3D, gdc.VariantTypeTransform3D))
  ptrsForTransform3D.operatorNotEqualTransform3DFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeTransform3D, gdc.VariantTypeTransform3D))
  ptrsForTransform3D.operatorMultiplyTransform3DFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeTransform3D, gdc.VariantTypeTransform3D))
  ptrsForTransform3D.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeTransform3D, gdc.VariantTypeDictionary))
  ptrsForTransform3D.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeTransform3D, gdc.VariantTypeArray))
  ptrsForTransform3D.operatorMultiplyPackedVector3ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeTransform3D, gdc.VariantTypePackedVector3Array))
  {
    memberName := StringNameFromStr("basis")
    defer memberName.Destroy()
    ptrsForTransform3D.memberbasisGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeTransform3D, memberName.AsCPtr()))
    ptrsForTransform3D.memberbasisSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeTransform3D, memberName.AsCPtr()))
  }
  {
    memberName := StringNameFromStr("origin")
    defer memberName.Destroy()
    ptrsForTransform3D.memberoriginGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeTransform3D, memberName.AsCPtr()))
    ptrsForTransform3D.memberoriginSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeTransform3D, memberName.AsCPtr()))
  }
  ptrsForTransform3D.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypeTransform3D))
  ptrsForTransform3D.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypeTransform3D))
}

type Transform3D struct {
  data   *[classSizeTransform3D]byte
  iface  gdc.Interface
  pinner runtime.Pinner
}

// Constants

var (
  Transform3DIdentity = "Transform3D(1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0)" // TODO: construct correctly
  Transform3DFlipX = "Transform3D(-1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0)" // TODO: construct correctly
  Transform3DFlipY = "Transform3D(1, 0, 0, 0, -1, 0, 0, 0, 1, 0, 0, 0)" // TODO: construct correctly
  Transform3DFlipZ = "Transform3D(1, 0, 0, 0, 1, 0, 0, 0, -1, 0, 0, 0)" // TODO: construct correctly
)

// Enums

// Constructors
func newTransform3D() *Transform3D {
  me := &Transform3D{
    data:   new([classSizeTransform3D]byte),
    iface:  giface,
  }
  me.pinner.Pin(me)
  me.pinner.Pin(me.AsTypePtr())
  return me
}

func NewTransform3D() *Transform3D {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newTransform3D()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForTransform3D.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewTransform3DFromTransform3D(from Transform3D, ) *Transform3D {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newTransform3D()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForTransform3D.ctrFromTransform3DFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewTransform3DFromBasisVector3(basis Basis, origin Vector3, ) *Transform3D {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newTransform3D()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForTransform3D.ctrFromBasisVector3Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{basis.AsCTypePtr(), origin.AsCTypePtr(), }))
  return me
}

func NewTransform3DFromVector3Vector3Vector3Vector3(x_axis Vector3, y_axis Vector3, z_axis Vector3, origin Vector3, ) *Transform3D {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newTransform3D()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForTransform3D.ctrFromVector3Vector3Vector3Vector3Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{x_axis.AsCTypePtr(), y_axis.AsCTypePtr(), z_axis.AsCTypePtr(), origin.AsCTypePtr(), }))
  return me
}

func NewTransform3DFromProjection(from Projection, ) *Transform3D {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newTransform3D()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForTransform3D.ctrFromProjectionFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

// Destructor
func (me *Transform3D) Destroy() {
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsTransform3D() (*Transform3D, error) {
	if me.Type() != gdc.VariantTypeTransform3D {
		return nil, fmt.Errorf("variant is not a Transform3D")
	}
  bclass := newTransform3D()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForTransform3D.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Transform3D) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForTransform3D.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
  return va
}

// Pointers
func Transform3DFromPtr(ptr gdc.ConstTypePtr) *Transform3D {
  me := newTransform3D()
  dataFromPtr(me.data[:], ptr)
  return me
}

func (me *Transform3D) Type() gdc.VariantType {
  return gdc.VariantTypeTransform3D
}

func (me *Transform3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Transform3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Transform3D) asUninitialized() gdc.UninitializedTypePtr {
  return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *Transform3D) Inverse() Transform3D {
  ret := NewTransform3D()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform3D.methodInverseFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Transform3D) AffineInverse() Transform3D {
  ret := NewTransform3D()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform3D.methodAffineInverseFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Transform3D) Orthonormalized() Transform3D {
  ret := NewTransform3D()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform3D.methodOrthonormalizedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Transform3D) Rotated(axis Vector3, angle float64, ) Transform3D {
  ret := NewTransform3D()


  varg1 := NewFloatFromFloat32(angle)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{axis.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform3D.methodRotatedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Transform3D) RotatedLocal(axis Vector3, angle float64, ) Transform3D {
  ret := NewTransform3D()


  varg1 := NewFloatFromFloat32(angle)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{axis.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform3D.methodRotatedLocalFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Transform3D) Scaled(scale Vector3, ) Transform3D {
  ret := NewTransform3D()


  args := []gdc.ConstTypePtr{scale.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform3D.methodScaledFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Transform3D) ScaledLocal(scale Vector3, ) Transform3D {
  ret := NewTransform3D()


  args := []gdc.ConstTypePtr{scale.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform3D.methodScaledLocalFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Transform3D) Translated(offset Vector3, ) Transform3D {
  ret := NewTransform3D()


  args := []gdc.ConstTypePtr{offset.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform3D.methodTranslatedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Transform3D) TranslatedLocal(offset Vector3, ) Transform3D {
  ret := NewTransform3D()


  args := []gdc.ConstTypePtr{offset.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform3D.methodTranslatedLocalFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Transform3D) LookingAt(target Vector3, up Vector3, use_model_front bool, ) Transform3D {
  ret := NewTransform3D()



  varg2 := NewBoolFromBool(use_model_front)
  defer varg2.Destroy()
  args := []gdc.ConstTypePtr{target.AsCTypePtr(), up.AsCTypePtr(), varg2.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform3D.methodLookingAtFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Transform3D) InterpolateWith(xform Transform3D, weight float64, ) Transform3D {
  ret := NewTransform3D()


  varg1 := NewFloatFromFloat32(weight)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{xform.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform3D.methodInterpolateWithFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Transform3D) IsEqualApprox(xform Transform3D, ) bool {
  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{xform.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform3D.methodIsEqualApproxFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Transform3D) IsFinite() bool {
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform3D.methodIsFiniteFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

// Operators

func (me *Transform3D) EqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Transform3D) NotEqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Transform3D) Not() bool {
  opPtr := ptrsForTransform3D.operatorNotFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
  return ret.Get()
}

func (me *Transform3D) MultiplyInt(rightArg int64) Transform3D {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  opPtr := ptrsForTransform3D.operatorMultiplyIntFn
  ret := NewTransform3D()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Transform3D) MultiplyFloat32(rightArg float64) Transform3D {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  opPtr := ptrsForTransform3D.operatorMultiplyFloat32Fn
  ret := NewTransform3D()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Transform3D) MultiplyVector3(right Vector3) Vector3 {
  opPtr := ptrsForTransform3D.operatorMultiplyVector3Fn
  ret := NewVector3()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Transform3D) MultiplyPlane(right Plane) Plane {
  opPtr := ptrsForTransform3D.operatorMultiplyPlaneFn
  ret := NewPlane()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Transform3D) MultiplyAABB(right AABB) AABB {
  opPtr := ptrsForTransform3D.operatorMultiplyAABBFn
  ret := NewAABB()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Transform3D) EqualTransform3D(right Transform3D) bool {
  opPtr := ptrsForTransform3D.operatorEqualTransform3DFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Transform3D) NotEqualTransform3D(right Transform3D) bool {
  opPtr := ptrsForTransform3D.operatorNotEqualTransform3DFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Transform3D) MultiplyTransform3D(right Transform3D) Transform3D {
  opPtr := ptrsForTransform3D.operatorMultiplyTransform3DFn
  ret := NewTransform3D()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Transform3D) InDictionary(right Dictionary) bool {
  opPtr := ptrsForTransform3D.operatorInDictionaryFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Transform3D) InArray(right Array) bool {
  opPtr := ptrsForTransform3D.operatorInArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Transform3D) MultiplyPackedVector3Array(right PackedVector3Array) PackedVector3Array {
  opPtr := ptrsForTransform3D.operatorMultiplyPackedVector3ArrayFn
  ret := NewPackedVector3Array()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

// Members

func (me *Transform3D) Basis() Basis {
  ret := NewBasis()
  me.iface.CallPtrGetter(ensurePtr(ptrsForTransform3D.memberbasisGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Transform3D) SetBasis(value Basis) {
  valueV := value
  me.iface.CallPtrSetter(ensurePtr(ptrsForTransform3D.memberbasisSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Transform3D) Origin() Vector3 {
  ret := NewVector3()
  me.iface.CallPtrGetter(ensurePtr(ptrsForTransform3D.memberoriginGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Transform3D) SetOrigin(value Vector3) {
  valueV := value
  me.iface.CallPtrSetter(ensurePtr(ptrsForTransform3D.memberoriginSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}
