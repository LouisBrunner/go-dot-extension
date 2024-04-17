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

type ptrsForTransform2DList struct {
	ctrFn                                 gdc.PtrConstructor
	ctrFromTransform2DFn                  gdc.PtrConstructor
	ctrFromFloat32Vector2Fn               gdc.PtrConstructor
	ctrFromFloat32Vector2Float32Vector2Fn gdc.PtrConstructor
	ctrFromVector2Vector2Vector2Fn        gdc.PtrConstructor
	methodInverseFn                       gdc.PtrBuiltInMethod
	methodAffineInverseFn                 gdc.PtrBuiltInMethod
	methodGetRotationFn                   gdc.PtrBuiltInMethod
	methodGetOriginFn                     gdc.PtrBuiltInMethod
	methodGetScaleFn                      gdc.PtrBuiltInMethod
	methodGetSkewFn                       gdc.PtrBuiltInMethod
	methodOrthonormalizedFn               gdc.PtrBuiltInMethod
	methodRotatedFn                       gdc.PtrBuiltInMethod
	methodRotatedLocalFn                  gdc.PtrBuiltInMethod
	methodScaledFn                        gdc.PtrBuiltInMethod
	methodScaledLocalFn                   gdc.PtrBuiltInMethod
	methodTranslatedFn                    gdc.PtrBuiltInMethod
	methodTranslatedLocalFn               gdc.PtrBuiltInMethod
	methodDeterminantFn                   gdc.PtrBuiltInMethod
	methodBasisXformFn                    gdc.PtrBuiltInMethod
	methodBasisXformInvFn                 gdc.PtrBuiltInMethod
	methodInterpolateWithFn               gdc.PtrBuiltInMethod
	methodIsConformalFn                   gdc.PtrBuiltInMethod
	methodIsEqualApproxFn                 gdc.PtrBuiltInMethod
	methodIsFiniteFn                      gdc.PtrBuiltInMethod
	methodLookingAtFn                     gdc.PtrBuiltInMethod
	operatorNotFn                         gdc.PtrOperatorEvaluator
	operatorMultiplyIntFn                 gdc.PtrOperatorEvaluator
	operatorDivideIntFn                   gdc.PtrOperatorEvaluator
	operatorMultiplyFloat32Fn             gdc.PtrOperatorEvaluator
	operatorDivideFloat32Fn               gdc.PtrOperatorEvaluator
	operatorMultiplyVector2Fn             gdc.PtrOperatorEvaluator
	operatorMultiplyRect2Fn               gdc.PtrOperatorEvaluator
	operatorEqualTransform2DFn            gdc.PtrOperatorEvaluator
	operatorNotEqualTransform2DFn         gdc.PtrOperatorEvaluator
	operatorMultiplyTransform2DFn         gdc.PtrOperatorEvaluator
	operatorInDictionaryFn                gdc.PtrOperatorEvaluator
	operatorInArrayFn                     gdc.PtrOperatorEvaluator
	operatorMultiplyPackedVector2ArrayFn  gdc.PtrOperatorEvaluator
	memberxGetterFn                       gdc.PtrGetter
	memberxSetterFn                       gdc.PtrSetter
	memberyGetterFn                       gdc.PtrGetter
	memberySetterFn                       gdc.PtrSetter
	memberoriginGetterFn                  gdc.PtrGetter
	memberoriginSetterFn                  gdc.PtrSetter
	toVariantFn                           gdc.TypeFromVariantConstructorFunc
	fromVariantFn                         gdc.VariantFromTypeConstructorFunc
}

var ptrsForTransform2D ptrsForTransform2DList

func initTransform2DPtrs(iface gdc.Interface) {
	ptrsForTransform2D.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeTransform2D, 0))
	ptrsForTransform2D.ctrFromTransform2DFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeTransform2D, 1))
	ptrsForTransform2D.ctrFromFloat32Vector2Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeTransform2D, 2))
	ptrsForTransform2D.ctrFromFloat32Vector2Float32Vector2Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeTransform2D, 3))
	ptrsForTransform2D.ctrFromVector2Vector2Vector2Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeTransform2D, 4))
	{
		methodName := StringNameFromStr("inverse")
		defer methodName.Destroy()
		ptrsForTransform2D.methodInverseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, methodName.AsCPtr(), 1420440541))
	}
	{
		methodName := StringNameFromStr("affine_inverse")
		defer methodName.Destroy()
		ptrsForTransform2D.methodAffineInverseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, methodName.AsCPtr(), 1420440541))
	}
	{
		methodName := StringNameFromStr("get_rotation")
		defer methodName.Destroy()
		ptrsForTransform2D.methodGetRotationFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, methodName.AsCPtr(), 466405837))
	}
	{
		methodName := StringNameFromStr("get_origin")
		defer methodName.Destroy()
		ptrsForTransform2D.methodGetOriginFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, methodName.AsCPtr(), 2428350749))
	}
	{
		methodName := StringNameFromStr("get_scale")
		defer methodName.Destroy()
		ptrsForTransform2D.methodGetScaleFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, methodName.AsCPtr(), 2428350749))
	}
	{
		methodName := StringNameFromStr("get_skew")
		defer methodName.Destroy()
		ptrsForTransform2D.methodGetSkewFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, methodName.AsCPtr(), 466405837))
	}
	{
		methodName := StringNameFromStr("orthonormalized")
		defer methodName.Destroy()
		ptrsForTransform2D.methodOrthonormalizedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, methodName.AsCPtr(), 1420440541))
	}
	{
		methodName := StringNameFromStr("rotated")
		defer methodName.Destroy()
		ptrsForTransform2D.methodRotatedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, methodName.AsCPtr(), 729597514))
	}
	{
		methodName := StringNameFromStr("rotated_local")
		defer methodName.Destroy()
		ptrsForTransform2D.methodRotatedLocalFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, methodName.AsCPtr(), 729597514))
	}
	{
		methodName := StringNameFromStr("scaled")
		defer methodName.Destroy()
		ptrsForTransform2D.methodScaledFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, methodName.AsCPtr(), 1446323263))
	}
	{
		methodName := StringNameFromStr("scaled_local")
		defer methodName.Destroy()
		ptrsForTransform2D.methodScaledLocalFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, methodName.AsCPtr(), 1446323263))
	}
	{
		methodName := StringNameFromStr("translated")
		defer methodName.Destroy()
		ptrsForTransform2D.methodTranslatedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, methodName.AsCPtr(), 1446323263))
	}
	{
		methodName := StringNameFromStr("translated_local")
		defer methodName.Destroy()
		ptrsForTransform2D.methodTranslatedLocalFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, methodName.AsCPtr(), 1446323263))
	}
	{
		methodName := StringNameFromStr("determinant")
		defer methodName.Destroy()
		ptrsForTransform2D.methodDeterminantFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, methodName.AsCPtr(), 466405837))
	}
	{
		methodName := StringNameFromStr("basis_xform")
		defer methodName.Destroy()
		ptrsForTransform2D.methodBasisXformFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, methodName.AsCPtr(), 2026743667))
	}
	{
		methodName := StringNameFromStr("basis_xform_inv")
		defer methodName.Destroy()
		ptrsForTransform2D.methodBasisXformInvFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, methodName.AsCPtr(), 2026743667))
	}
	{
		methodName := StringNameFromStr("interpolate_with")
		defer methodName.Destroy()
		ptrsForTransform2D.methodInterpolateWithFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, methodName.AsCPtr(), 359399686))
	}
	{
		methodName := StringNameFromStr("is_conformal")
		defer methodName.Destroy()
		ptrsForTransform2D.methodIsConformalFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("is_equal_approx")
		defer methodName.Destroy()
		ptrsForTransform2D.methodIsEqualApproxFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, methodName.AsCPtr(), 3837431929))
	}
	{
		methodName := StringNameFromStr("is_finite")
		defer methodName.Destroy()
		ptrsForTransform2D.methodIsFiniteFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("looking_at")
		defer methodName.Destroy()
		ptrsForTransform2D.methodLookingAtFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeTransform2D, methodName.AsCPtr(), 1446323263))
	}
	ptrsForTransform2D.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypeTransform2D, gdc.VariantTypeNil))
	ptrsForTransform2D.operatorMultiplyIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeTransform2D, gdc.VariantTypeInt))
	ptrsForTransform2D.operatorDivideIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeTransform2D, gdc.VariantTypeInt))
	ptrsForTransform2D.operatorMultiplyFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeTransform2D, gdc.VariantTypeFloat))
	ptrsForTransform2D.operatorDivideFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeTransform2D, gdc.VariantTypeFloat))
	ptrsForTransform2D.operatorMultiplyVector2Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeTransform2D, gdc.VariantTypeVector2))
	ptrsForTransform2D.operatorMultiplyRect2Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeTransform2D, gdc.VariantTypeRect2))
	ptrsForTransform2D.operatorEqualTransform2DFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeTransform2D, gdc.VariantTypeTransform2D))
	ptrsForTransform2D.operatorNotEqualTransform2DFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeTransform2D, gdc.VariantTypeTransform2D))
	ptrsForTransform2D.operatorMultiplyTransform2DFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeTransform2D, gdc.VariantTypeTransform2D))
	ptrsForTransform2D.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeTransform2D, gdc.VariantTypeDictionary))
	ptrsForTransform2D.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeTransform2D, gdc.VariantTypeArray))
	ptrsForTransform2D.operatorMultiplyPackedVector2ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeTransform2D, gdc.VariantTypePackedVector2Array))
	{
		memberName := StringNameFromStr("x")
		defer memberName.Destroy()
		ptrsForTransform2D.memberxGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeTransform2D, memberName.AsCPtr()))
		ptrsForTransform2D.memberxSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeTransform2D, memberName.AsCPtr()))
	}
	{
		memberName := StringNameFromStr("y")
		defer memberName.Destroy()
		ptrsForTransform2D.memberyGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeTransform2D, memberName.AsCPtr()))
		ptrsForTransform2D.memberySetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeTransform2D, memberName.AsCPtr()))
	}
	{
		memberName := StringNameFromStr("origin")
		defer memberName.Destroy()
		ptrsForTransform2D.memberoriginGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeTransform2D, memberName.AsCPtr()))
		ptrsForTransform2D.memberoriginSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeTransform2D, memberName.AsCPtr()))
	}
	ptrsForTransform2D.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypeTransform2D))
	ptrsForTransform2D.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypeTransform2D))
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("IDENTITY")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeTransform2D, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsTransform2D()
		if err != nil {
			panic("Failed to get constant value IDENTITY: " + err.Error())
		}
		Transform2DIdentity = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("FLIP_X")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeTransform2D, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsTransform2D()
		if err != nil {
			panic("Failed to get constant value FLIP_X: " + err.Error())
		}
		Transform2DFlipX = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("FLIP_Y")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeTransform2D, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsTransform2D()
		if err != nil {
			panic("Failed to get constant value FLIP_Y: " + err.Error())
		}
		Transform2DFlipY = *cnst
	}
}

type Transform2D struct {
	data   *[classSizeTransform2D]byte
	iface  gdc.Interface
	pinner runtime.Pinner
}

// Constants

var (
	Transform2DIdentity Transform2D
	Transform2DFlipX    Transform2D
	Transform2DFlipY    Transform2D
)

// Enums

// Constructors
func newTransform2D() *Transform2D {
	me := &Transform2D{
		data:  new([classSizeTransform2D]byte),
		iface: giface,
	}
	me.pinner.Pin(me)
	me.pinner.Pin(me.AsTypePtr())
	return me
}

func NewTransform2D() *Transform2D {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newTransform2D()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForTransform2D.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
	return me
}

func NewTransform2DFromTransform2D(from Transform2D) *Transform2D {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newTransform2D()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForTransform2D.ctrFromTransform2DFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewTransform2DFromFloat32Vector2(rotation float64, position Vector2) *Transform2D {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	pinner.Pin(&rotation)
	me := newTransform2D()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForTransform2D.ctrFromFloat32Vector2Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&rotation), position.AsCTypePtr()}))
	return me
}

func NewTransform2DFromFloat32Vector2Float32Vector2(rotation float64, scale Vector2, skew float64, position Vector2) *Transform2D {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	pinner.Pin(&rotation)
	pinner.Pin(&skew)
	me := newTransform2D()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForTransform2D.ctrFromFloat32Vector2Float32Vector2Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&rotation), scale.AsCTypePtr(), gdc.ConstTypePtr(&skew), position.AsCTypePtr()}))
	return me
}

func NewTransform2DFromVector2Vector2Vector2(x_axis Vector2, y_axis Vector2, origin Vector2) *Transform2D {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newTransform2D()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForTransform2D.ctrFromVector2Vector2Vector2Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{x_axis.AsCTypePtr(), y_axis.AsCTypePtr(), origin.AsCTypePtr()}))
	return me
}

// Destructor
func (me *Transform2D) Destroy() {
	me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsTransform2D() (*Transform2D, error) {
	if me.Type() != gdc.VariantTypeTransform2D {
		return nil, fmt.Errorf("variant is not a Transform2D")
	}
	bclass := newTransform2D()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForTransform2D.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Transform2D) AsVariant() *Variant {
	va := newVariant()
	va.inner = me
	me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForTransform2D.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
	return va
}

// Pointers
func Transform2DFromPtr(ptr gdc.ConstTypePtr) *Transform2D {
	me := newTransform2D()
	dataFromPtr(me.data[:], ptr)
	return me
}

func (me *Transform2D) ToTypePtr(ptr gdc.TypePtr) {
	dataToPtr(ptr, me.data[:])
}

func (me *Transform2D) Type() gdc.VariantType {
	return gdc.VariantTypeTransform2D
}

func (me *Transform2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Transform2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Transform2D) asUninitialized() gdc.UninitializedTypePtr {
	return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *Transform2D) Inverse() Transform2D {
	ret := NewTransform2D()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform2D.methodInverseFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Transform2D) AffineInverse() Transform2D {
	ret := NewTransform2D()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform2D.methodAffineInverseFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Transform2D) GetRotation() float64 {
	ret := NewFloat()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform2D.methodGetRotationFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Transform2D) GetOrigin() Vector2 {
	ret := NewVector2()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform2D.methodGetOriginFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Transform2D) GetScale() Vector2 {
	ret := NewVector2()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform2D.methodGetScaleFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Transform2D) GetSkew() float64 {
	ret := NewFloat()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform2D.methodGetSkewFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Transform2D) Orthonormalized() Transform2D {
	ret := NewTransform2D()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform2D.methodOrthonormalizedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Transform2D) Rotated(angle float64) Transform2D {
	ret := NewTransform2D()

	varg0 := NewFloatFromFloat32(angle)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform2D.methodRotatedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Transform2D) RotatedLocal(angle float64) Transform2D {
	ret := NewTransform2D()

	varg0 := NewFloatFromFloat32(angle)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform2D.methodRotatedLocalFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Transform2D) Scaled(scale Vector2) Transform2D {
	ret := NewTransform2D()

	args := []gdc.ConstTypePtr{scale.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform2D.methodScaledFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Transform2D) ScaledLocal(scale Vector2) Transform2D {
	ret := NewTransform2D()

	args := []gdc.ConstTypePtr{scale.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform2D.methodScaledLocalFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Transform2D) Translated(offset Vector2) Transform2D {
	ret := NewTransform2D()

	args := []gdc.ConstTypePtr{offset.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform2D.methodTranslatedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Transform2D) TranslatedLocal(offset Vector2) Transform2D {
	ret := NewTransform2D()

	args := []gdc.ConstTypePtr{offset.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform2D.methodTranslatedLocalFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Transform2D) Determinant() float64 {
	ret := NewFloat()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform2D.methodDeterminantFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Transform2D) BasisXform(v Vector2) Vector2 {
	ret := NewVector2()

	args := []gdc.ConstTypePtr{v.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform2D.methodBasisXformFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Transform2D) BasisXformInv(v Vector2) Vector2 {
	ret := NewVector2()

	args := []gdc.ConstTypePtr{v.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform2D.methodBasisXformInvFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Transform2D) InterpolateWith(xform Transform2D, weight float64) Transform2D {
	ret := NewTransform2D()

	varg1 := NewFloatFromFloat32(weight)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{xform.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform2D.methodInterpolateWithFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Transform2D) IsConformal() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform2D.methodIsConformalFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Transform2D) IsEqualApprox(xform Transform2D) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{xform.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform2D.methodIsEqualApproxFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Transform2D) IsFinite() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform2D.methodIsFiniteFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Transform2D) LookingAt(target Vector2) Transform2D {
	ret := NewTransform2D()

	args := []gdc.ConstTypePtr{target.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForTransform2D.methodLookingAtFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

// Operators

func (me *Transform2D) EqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Transform2D) NotEqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Transform2D) Not() bool {
	opPtr := ptrsForTransform2D.operatorNotFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return ret.Get()
}

func (me *Transform2D) MultiplyInt(rightArg int64) Transform2D {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForTransform2D.operatorMultiplyIntFn
	ret := NewTransform2D()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Transform2D) DivideInt(rightArg int64) Transform2D {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForTransform2D.operatorDivideIntFn
	ret := NewTransform2D()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Transform2D) MultiplyFloat32(rightArg float64) Transform2D {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForTransform2D.operatorMultiplyFloat32Fn
	ret := NewTransform2D()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Transform2D) DivideFloat32(rightArg float64) Transform2D {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForTransform2D.operatorDivideFloat32Fn
	ret := NewTransform2D()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Transform2D) MultiplyVector2(right Vector2) Vector2 {
	opPtr := ptrsForTransform2D.operatorMultiplyVector2Fn
	ret := NewVector2()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Transform2D) MultiplyRect2(right Rect2) Rect2 {
	opPtr := ptrsForTransform2D.operatorMultiplyRect2Fn
	ret := NewRect2()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Transform2D) EqualTransform2D(right Transform2D) bool {
	opPtr := ptrsForTransform2D.operatorEqualTransform2DFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Transform2D) NotEqualTransform2D(right Transform2D) bool {
	opPtr := ptrsForTransform2D.operatorNotEqualTransform2DFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Transform2D) MultiplyTransform2D(right Transform2D) Transform2D {
	opPtr := ptrsForTransform2D.operatorMultiplyTransform2DFn
	ret := NewTransform2D()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Transform2D) InDictionary(right Dictionary) bool {
	opPtr := ptrsForTransform2D.operatorInDictionaryFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Transform2D) InArray(right Array) bool {
	opPtr := ptrsForTransform2D.operatorInArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Transform2D) MultiplyPackedVector2Array(right PackedVector2Array) PackedVector2Array {
	opPtr := ptrsForTransform2D.operatorMultiplyPackedVector2ArrayFn
	ret := NewPackedVector2Array()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

// Members

func (me *Transform2D) X() Vector2 {
	ret := NewVector2()
	me.iface.CallPtrGetter(ensurePtr(ptrsForTransform2D.memberxGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Transform2D) SetX(value Vector2) {
	valueV := value
	me.iface.CallPtrSetter(ensurePtr(ptrsForTransform2D.memberxSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Transform2D) Y() Vector2 {
	ret := NewVector2()
	me.iface.CallPtrGetter(ensurePtr(ptrsForTransform2D.memberyGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Transform2D) SetY(value Vector2) {
	valueV := value
	me.iface.CallPtrSetter(ensurePtr(ptrsForTransform2D.memberySetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Transform2D) Origin() Vector2 {
	ret := NewVector2()
	me.iface.CallPtrGetter(ensurePtr(ptrsForTransform2D.memberoriginGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Transform2D) SetOrigin(value Vector2) {
	valueV := value
	me.iface.CallPtrSetter(ensurePtr(ptrsForTransform2D.memberoriginSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}
