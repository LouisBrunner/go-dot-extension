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

type ptrsForVector2List struct {
	ctrFn                          gdc.PtrConstructor
	ctrFromVector2Fn               gdc.PtrConstructor
	ctrFromVector2IFn              gdc.PtrConstructor
	ctrFromFloat32Float32Fn        gdc.PtrConstructor
	methodAngleFn                  gdc.PtrBuiltInMethod
	methodAngleToFn                gdc.PtrBuiltInMethod
	methodAngleToPointFn           gdc.PtrBuiltInMethod
	methodDirectionToFn            gdc.PtrBuiltInMethod
	methodDistanceToFn             gdc.PtrBuiltInMethod
	methodDistanceSquaredToFn      gdc.PtrBuiltInMethod
	methodLengthFn                 gdc.PtrBuiltInMethod
	methodLengthSquaredFn          gdc.PtrBuiltInMethod
	methodLimitLengthFn            gdc.PtrBuiltInMethod
	methodNormalizedFn             gdc.PtrBuiltInMethod
	methodIsNormalizedFn           gdc.PtrBuiltInMethod
	methodIsEqualApproxFn          gdc.PtrBuiltInMethod
	methodIsZeroApproxFn           gdc.PtrBuiltInMethod
	methodIsFiniteFn               gdc.PtrBuiltInMethod
	methodPosmodFn                 gdc.PtrBuiltInMethod
	methodPosmodvFn                gdc.PtrBuiltInMethod
	methodProjectFn                gdc.PtrBuiltInMethod
	methodLerpFn                   gdc.PtrBuiltInMethod
	methodSlerpFn                  gdc.PtrBuiltInMethod
	methodCubicInterpolateFn       gdc.PtrBuiltInMethod
	methodCubicInterpolateInTimeFn gdc.PtrBuiltInMethod
	methodBezierInterpolateFn      gdc.PtrBuiltInMethod
	methodBezierDerivativeFn       gdc.PtrBuiltInMethod
	methodMaxAxisIndexFn           gdc.PtrBuiltInMethod
	methodMinAxisIndexFn           gdc.PtrBuiltInMethod
	methodMoveTowardFn             gdc.PtrBuiltInMethod
	methodRotatedFn                gdc.PtrBuiltInMethod
	methodOrthogonalFn             gdc.PtrBuiltInMethod
	methodFloorFn                  gdc.PtrBuiltInMethod
	methodCeilFn                   gdc.PtrBuiltInMethod
	methodRoundFn                  gdc.PtrBuiltInMethod
	methodAspectFn                 gdc.PtrBuiltInMethod
	methodDotFn                    gdc.PtrBuiltInMethod
	methodSlideFn                  gdc.PtrBuiltInMethod
	methodBounceFn                 gdc.PtrBuiltInMethod
	methodReflectFn                gdc.PtrBuiltInMethod
	methodCrossFn                  gdc.PtrBuiltInMethod
	methodAbsFn                    gdc.PtrBuiltInMethod
	methodSignFn                   gdc.PtrBuiltInMethod
	methodClampFn                  gdc.PtrBuiltInMethod
	methodSnappedFn                gdc.PtrBuiltInMethod
	methodFromAngleFn              gdc.PtrBuiltInMethod
	operatorNegateFn               gdc.PtrOperatorEvaluator
	operatorPositiveFn             gdc.PtrOperatorEvaluator
	operatorNotFn                  gdc.PtrOperatorEvaluator
	operatorMultiplyIntFn          gdc.PtrOperatorEvaluator
	operatorDivideIntFn            gdc.PtrOperatorEvaluator
	operatorMultiplyFloat32Fn      gdc.PtrOperatorEvaluator
	operatorDivideFloat32Fn        gdc.PtrOperatorEvaluator
	operatorEqualVector2Fn         gdc.PtrOperatorEvaluator
	operatorNotEqualVector2Fn      gdc.PtrOperatorEvaluator
	operatorLessVector2Fn          gdc.PtrOperatorEvaluator
	operatorLessEqualVector2Fn     gdc.PtrOperatorEvaluator
	operatorGreaterVector2Fn       gdc.PtrOperatorEvaluator
	operatorGreaterEqualVector2Fn  gdc.PtrOperatorEvaluator
	operatorAddVector2Fn           gdc.PtrOperatorEvaluator
	operatorSubtractVector2Fn      gdc.PtrOperatorEvaluator
	operatorMultiplyVector2Fn      gdc.PtrOperatorEvaluator
	operatorDivideVector2Fn        gdc.PtrOperatorEvaluator
	operatorMultiplyTransform2DFn  gdc.PtrOperatorEvaluator
	operatorInDictionaryFn         gdc.PtrOperatorEvaluator
	operatorInArrayFn              gdc.PtrOperatorEvaluator
	operatorInPackedVector2ArrayFn gdc.PtrOperatorEvaluator
	memberxGetterFn                gdc.PtrGetter
	memberxSetterFn                gdc.PtrSetter
	memberyGetterFn                gdc.PtrGetter
	memberySetterFn                gdc.PtrSetter
	toVariantFn                    gdc.TypeFromVariantConstructorFunc
	fromVariantFn                  gdc.VariantFromTypeConstructorFunc
}

var ptrsForVector2 ptrsForVector2List

func initVector2Ptrs(iface gdc.Interface) {
	ptrsForVector2.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeVector2, 0))
	ptrsForVector2.ctrFromVector2Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeVector2, 1))
	ptrsForVector2.ctrFromVector2IFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeVector2, 2))
	ptrsForVector2.ctrFromFloat32Float32Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeVector2, 3))
	{
		methodName := StringNameFromStr("angle")
		defer methodName.Destroy()
		ptrsForVector2.methodAngleFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 466405837))
	}
	{
		methodName := StringNameFromStr("angle_to")
		defer methodName.Destroy()
		ptrsForVector2.methodAngleToFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 3819070308))
	}
	{
		methodName := StringNameFromStr("angle_to_point")
		defer methodName.Destroy()
		ptrsForVector2.methodAngleToPointFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 3819070308))
	}
	{
		methodName := StringNameFromStr("direction_to")
		defer methodName.Destroy()
		ptrsForVector2.methodDirectionToFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 2026743667))
	}
	{
		methodName := StringNameFromStr("distance_to")
		defer methodName.Destroy()
		ptrsForVector2.methodDistanceToFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 3819070308))
	}
	{
		methodName := StringNameFromStr("distance_squared_to")
		defer methodName.Destroy()
		ptrsForVector2.methodDistanceSquaredToFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 3819070308))
	}
	{
		methodName := StringNameFromStr("length")
		defer methodName.Destroy()
		ptrsForVector2.methodLengthFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 466405837))
	}
	{
		methodName := StringNameFromStr("length_squared")
		defer methodName.Destroy()
		ptrsForVector2.methodLengthSquaredFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 466405837))
	}
	{
		methodName := StringNameFromStr("limit_length")
		defer methodName.Destroy()
		ptrsForVector2.methodLimitLengthFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 2544004089))
	}
	{
		methodName := StringNameFromStr("normalized")
		defer methodName.Destroy()
		ptrsForVector2.methodNormalizedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 2428350749))
	}
	{
		methodName := StringNameFromStr("is_normalized")
		defer methodName.Destroy()
		ptrsForVector2.methodIsNormalizedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("is_equal_approx")
		defer methodName.Destroy()
		ptrsForVector2.methodIsEqualApproxFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 3190634762))
	}
	{
		methodName := StringNameFromStr("is_zero_approx")
		defer methodName.Destroy()
		ptrsForVector2.methodIsZeroApproxFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("is_finite")
		defer methodName.Destroy()
		ptrsForVector2.methodIsFiniteFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("posmod")
		defer methodName.Destroy()
		ptrsForVector2.methodPosmodFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 2544004089))
	}
	{
		methodName := StringNameFromStr("posmodv")
		defer methodName.Destroy()
		ptrsForVector2.methodPosmodvFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 2026743667))
	}
	{
		methodName := StringNameFromStr("project")
		defer methodName.Destroy()
		ptrsForVector2.methodProjectFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 2026743667))
	}
	{
		methodName := StringNameFromStr("lerp")
		defer methodName.Destroy()
		ptrsForVector2.methodLerpFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 4250033116))
	}
	{
		methodName := StringNameFromStr("slerp")
		defer methodName.Destroy()
		ptrsForVector2.methodSlerpFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 4250033116))
	}
	{
		methodName := StringNameFromStr("cubic_interpolate")
		defer methodName.Destroy()
		ptrsForVector2.methodCubicInterpolateFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 193522989))
	}
	{
		methodName := StringNameFromStr("cubic_interpolate_in_time")
		defer methodName.Destroy()
		ptrsForVector2.methodCubicInterpolateInTimeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 1957055074))
	}
	{
		methodName := StringNameFromStr("bezier_interpolate")
		defer methodName.Destroy()
		ptrsForVector2.methodBezierInterpolateFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 193522989))
	}
	{
		methodName := StringNameFromStr("bezier_derivative")
		defer methodName.Destroy()
		ptrsForVector2.methodBezierDerivativeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 193522989))
	}
	{
		methodName := StringNameFromStr("max_axis_index")
		defer methodName.Destroy()
		ptrsForVector2.methodMaxAxisIndexFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("min_axis_index")
		defer methodName.Destroy()
		ptrsForVector2.methodMinAxisIndexFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("move_toward")
		defer methodName.Destroy()
		ptrsForVector2.methodMoveTowardFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 4250033116))
	}
	{
		methodName := StringNameFromStr("rotated")
		defer methodName.Destroy()
		ptrsForVector2.methodRotatedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 2544004089))
	}
	{
		methodName := StringNameFromStr("orthogonal")
		defer methodName.Destroy()
		ptrsForVector2.methodOrthogonalFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 2428350749))
	}
	{
		methodName := StringNameFromStr("floor")
		defer methodName.Destroy()
		ptrsForVector2.methodFloorFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 2428350749))
	}
	{
		methodName := StringNameFromStr("ceil")
		defer methodName.Destroy()
		ptrsForVector2.methodCeilFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 2428350749))
	}
	{
		methodName := StringNameFromStr("round")
		defer methodName.Destroy()
		ptrsForVector2.methodRoundFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 2428350749))
	}
	{
		methodName := StringNameFromStr("aspect")
		defer methodName.Destroy()
		ptrsForVector2.methodAspectFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 466405837))
	}
	{
		methodName := StringNameFromStr("dot")
		defer methodName.Destroy()
		ptrsForVector2.methodDotFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 3819070308))
	}
	{
		methodName := StringNameFromStr("slide")
		defer methodName.Destroy()
		ptrsForVector2.methodSlideFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 2026743667))
	}
	{
		methodName := StringNameFromStr("bounce")
		defer methodName.Destroy()
		ptrsForVector2.methodBounceFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 2026743667))
	}
	{
		methodName := StringNameFromStr("reflect")
		defer methodName.Destroy()
		ptrsForVector2.methodReflectFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 2026743667))
	}
	{
		methodName := StringNameFromStr("cross")
		defer methodName.Destroy()
		ptrsForVector2.methodCrossFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 3819070308))
	}
	{
		methodName := StringNameFromStr("abs")
		defer methodName.Destroy()
		ptrsForVector2.methodAbsFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 2428350749))
	}
	{
		methodName := StringNameFromStr("sign")
		defer methodName.Destroy()
		ptrsForVector2.methodSignFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 2428350749))
	}
	{
		methodName := StringNameFromStr("clamp")
		defer methodName.Destroy()
		ptrsForVector2.methodClampFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 318031021))
	}
	{
		methodName := StringNameFromStr("snapped")
		defer methodName.Destroy()
		ptrsForVector2.methodSnappedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 2026743667))
	}
	{
		methodName := StringNameFromStr("from_angle")
		defer methodName.Destroy()
		ptrsForVector2.methodFromAngleFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2, methodName.AsCPtr(), 889263119))
	}
	ptrsForVector2.operatorNegateFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNegate, gdc.VariantTypeVector2, gdc.VariantTypeNil))
	ptrsForVector2.operatorPositiveFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpPositive, gdc.VariantTypeVector2, gdc.VariantTypeNil))
	ptrsForVector2.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypeVector2, gdc.VariantTypeNil))
	ptrsForVector2.operatorMultiplyIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeVector2, gdc.VariantTypeInt))
	ptrsForVector2.operatorDivideIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeVector2, gdc.VariantTypeInt))
	ptrsForVector2.operatorMultiplyFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeVector2, gdc.VariantTypeFloat))
	ptrsForVector2.operatorDivideFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeVector2, gdc.VariantTypeFloat))
	ptrsForVector2.operatorEqualVector2Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeVector2, gdc.VariantTypeVector2))
	ptrsForVector2.operatorNotEqualVector2Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeVector2, gdc.VariantTypeVector2))
	ptrsForVector2.operatorLessVector2Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLess, gdc.VariantTypeVector2, gdc.VariantTypeVector2))
	ptrsForVector2.operatorLessEqualVector2Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLessEqual, gdc.VariantTypeVector2, gdc.VariantTypeVector2))
	ptrsForVector2.operatorGreaterVector2Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreater, gdc.VariantTypeVector2, gdc.VariantTypeVector2))
	ptrsForVector2.operatorGreaterEqualVector2Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreaterEqual, gdc.VariantTypeVector2, gdc.VariantTypeVector2))
	ptrsForVector2.operatorAddVector2Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, gdc.VariantTypeVector2, gdc.VariantTypeVector2))
	ptrsForVector2.operatorSubtractVector2Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpSubtract, gdc.VariantTypeVector2, gdc.VariantTypeVector2))
	ptrsForVector2.operatorMultiplyVector2Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeVector2, gdc.VariantTypeVector2))
	ptrsForVector2.operatorDivideVector2Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeVector2, gdc.VariantTypeVector2))
	ptrsForVector2.operatorMultiplyTransform2DFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeVector2, gdc.VariantTypeTransform2D))
	ptrsForVector2.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeVector2, gdc.VariantTypeDictionary))
	ptrsForVector2.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeVector2, gdc.VariantTypeArray))
	ptrsForVector2.operatorInPackedVector2ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeVector2, gdc.VariantTypePackedVector2Array))
	{
		memberName := StringNameFromStr("x")
		defer memberName.Destroy()
		ptrsForVector2.memberxGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeVector2, memberName.AsCPtr()))
		ptrsForVector2.memberxSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeVector2, memberName.AsCPtr()))
	}
	{
		memberName := StringNameFromStr("y")
		defer memberName.Destroy()
		ptrsForVector2.memberyGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeVector2, memberName.AsCPtr()))
		ptrsForVector2.memberySetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeVector2, memberName.AsCPtr()))
	}
	ptrsForVector2.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypeVector2))
	ptrsForVector2.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypeVector2))
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("ZERO")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector2, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector2()
		if err != nil {
			panic("Failed to get constant value ZERO: " + err.Error())
		}
		Vector2Zero = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("ONE")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector2, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector2()
		if err != nil {
			panic("Failed to get constant value ONE: " + err.Error())
		}
		Vector2One = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("INF")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector2, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector2()
		if err != nil {
			panic("Failed to get constant value INF: " + err.Error())
		}
		Vector2Inf = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("LEFT")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector2, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector2()
		if err != nil {
			panic("Failed to get constant value LEFT: " + err.Error())
		}
		Vector2Left = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("RIGHT")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector2, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector2()
		if err != nil {
			panic("Failed to get constant value RIGHT: " + err.Error())
		}
		Vector2Right = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("UP")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector2, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector2()
		if err != nil {
			panic("Failed to get constant value UP: " + err.Error())
		}
		Vector2Up = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("DOWN")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector2, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector2()
		if err != nil {
			panic("Failed to get constant value DOWN: " + err.Error())
		}
		Vector2Down = *cnst
	}
}

type Vector2 struct {
	data   *[classSizeVector2]byte
	iface  gdc.Interface
	pinner runtime.Pinner
}

// Constants

var (
	Vector2Zero  Vector2
	Vector2One   Vector2
	Vector2Inf   Vector2
	Vector2Left  Vector2
	Vector2Right Vector2
	Vector2Up    Vector2
	Vector2Down  Vector2
)

// Enums

type Vector2Axis int

const (
	Vector2AxisAxisX Vector2Axis = 0
	Vector2AxisAxisY Vector2Axis = 1
)

// Constructors
func newVector2() *Vector2 {
	me := &Vector2{
		data:  new([classSizeVector2]byte),
		iface: giface,
	}
	me.pinner.Pin(me)
	me.pinner.Pin(me.AsTypePtr())
	return me
}

func NewVector2() *Vector2 {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newVector2()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForVector2.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
	return me
}

func NewVector2FromVector2(from Vector2) *Vector2 {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newVector2()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForVector2.ctrFromVector2Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewVector2FromVector2I(from Vector2i) *Vector2 {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newVector2()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForVector2.ctrFromVector2IFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewVector2FromFloat32Float32(x float64, y float64) *Vector2 {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	pinner.Pin(&x)
	pinner.Pin(&y)
	me := newVector2()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForVector2.ctrFromFloat32Float32Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y)}))
	return me
}

// Destructor
func (me *Vector2) Destroy() {
	me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsVector2() (*Vector2, error) {
	if me.Type() != gdc.VariantTypeVector2 {
		return nil, fmt.Errorf("variant is not a Vector2")
	}
	bclass := newVector2()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForVector2.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Vector2) AsVariant() *Variant {
	va := newVariant()
	va.inner = me
	me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForVector2.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
	return va
}

// Pointers
func Vector2FromPtr(ptr gdc.ConstTypePtr) *Vector2 {
	me := newVector2()
	dataFromPtr(me.data[:], ptr)
	return me
}

func (me *Vector2) ToTypePtr(ptr gdc.TypePtr) {
	dataToPtr(ptr, me.data[:])
}

func (me *Vector2) Type() gdc.VariantType {
	return gdc.VariantTypeVector2
}

func (me *Vector2) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Vector2) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Vector2) asUninitialized() gdc.UninitializedTypePtr {
	return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *Vector2) Angle() float64 {
	ret := NewFloat()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodAngleFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector2) AngleTo(to Vector2) float64 {
	ret := NewFloat()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{to.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodAngleToFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector2) AngleToPoint(to Vector2) float64 {
	ret := NewFloat()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{to.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodAngleToPointFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector2) DirectionTo(to Vector2) Vector2 {
	ret := NewVector2()

	args := []gdc.ConstTypePtr{to.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodDirectionToFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector2) DistanceTo(to Vector2) float64 {
	ret := NewFloat()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{to.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodDistanceToFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector2) DistanceSquaredTo(to Vector2) float64 {
	ret := NewFloat()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{to.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodDistanceSquaredToFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector2) Length() float64 {
	ret := NewFloat()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodLengthFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector2) LengthSquared() float64 {
	ret := NewFloat()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodLengthSquaredFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector2) LimitLength(length float64) Vector2 {
	ret := NewVector2()

	varg0 := NewFloatFromFloat32(length)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodLimitLengthFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector2) Normalized() Vector2 {
	ret := NewVector2()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodNormalizedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector2) IsNormalized() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodIsNormalizedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector2) IsEqualApprox(to Vector2) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{to.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodIsEqualApproxFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector2) IsZeroApprox() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodIsZeroApproxFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector2) IsFinite() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodIsFiniteFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector2) Posmod(mod float64) Vector2 {
	ret := NewVector2()

	varg0 := NewFloatFromFloat32(mod)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodPosmodFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector2) Posmodv(modv Vector2) Vector2 {
	ret := NewVector2()

	args := []gdc.ConstTypePtr{modv.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodPosmodvFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector2) Project(b Vector2) Vector2 {
	ret := NewVector2()

	args := []gdc.ConstTypePtr{b.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodProjectFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector2) Lerp(to Vector2, weight float64) Vector2 {
	ret := NewVector2()

	varg1 := NewFloatFromFloat32(weight)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{to.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodLerpFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector2) Slerp(to Vector2, weight float64) Vector2 {
	ret := NewVector2()

	varg1 := NewFloatFromFloat32(weight)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{to.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodSlerpFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector2) CubicInterpolate(b Vector2, pre_a Vector2, post_b Vector2, weight float64) Vector2 {
	ret := NewVector2()

	varg3 := NewFloatFromFloat32(weight)
	defer varg3.Destroy()
	args := []gdc.ConstTypePtr{b.AsCTypePtr(), pre_a.AsCTypePtr(), post_b.AsCTypePtr(), varg3.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodCubicInterpolateFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector2) CubicInterpolateInTime(b Vector2, pre_a Vector2, post_b Vector2, weight float64, b_t float64, pre_a_t float64, post_b_t float64) Vector2 {
	ret := NewVector2()

	varg3 := NewFloatFromFloat32(weight)
	defer varg3.Destroy()
	varg4 := NewFloatFromFloat32(b_t)
	defer varg4.Destroy()
	varg5 := NewFloatFromFloat32(pre_a_t)
	defer varg5.Destroy()
	varg6 := NewFloatFromFloat32(post_b_t)
	defer varg6.Destroy()
	args := []gdc.ConstTypePtr{b.AsCTypePtr(), pre_a.AsCTypePtr(), post_b.AsCTypePtr(), varg3.AsCTypePtr(), varg4.AsCTypePtr(), varg5.AsCTypePtr(), varg6.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodCubicInterpolateInTimeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector2) BezierInterpolate(control_1 Vector2, control_2 Vector2, end Vector2, t float64) Vector2 {
	ret := NewVector2()

	varg3 := NewFloatFromFloat32(t)
	defer varg3.Destroy()
	args := []gdc.ConstTypePtr{control_1.AsCTypePtr(), control_2.AsCTypePtr(), end.AsCTypePtr(), varg3.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodBezierInterpolateFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector2) BezierDerivative(control_1 Vector2, control_2 Vector2, end Vector2, t float64) Vector2 {
	ret := NewVector2()

	varg3 := NewFloatFromFloat32(t)
	defer varg3.Destroy()
	args := []gdc.ConstTypePtr{control_1.AsCTypePtr(), control_2.AsCTypePtr(), end.AsCTypePtr(), varg3.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodBezierDerivativeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector2) MaxAxisIndex() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodMaxAxisIndexFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector2) MinAxisIndex() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodMinAxisIndexFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector2) MoveToward(to Vector2, delta float64) Vector2 {
	ret := NewVector2()

	varg1 := NewFloatFromFloat32(delta)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{to.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodMoveTowardFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector2) Rotated(angle float64) Vector2 {
	ret := NewVector2()

	varg0 := NewFloatFromFloat32(angle)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodRotatedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector2) Orthogonal() Vector2 {
	ret := NewVector2()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodOrthogonalFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector2) Floor() Vector2 {
	ret := NewVector2()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodFloorFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector2) Ceil() Vector2 {
	ret := NewVector2()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodCeilFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector2) Round() Vector2 {
	ret := NewVector2()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodRoundFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector2) Aspect() float64 {
	ret := NewFloat()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodAspectFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector2) Dot(with Vector2) float64 {
	ret := NewFloat()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{with.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodDotFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector2) Slide(n Vector2) Vector2 {
	ret := NewVector2()

	args := []gdc.ConstTypePtr{n.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodSlideFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector2) Bounce(n Vector2) Vector2 {
	ret := NewVector2()

	args := []gdc.ConstTypePtr{n.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodBounceFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector2) Reflect(line Vector2) Vector2 {
	ret := NewVector2()

	args := []gdc.ConstTypePtr{line.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodReflectFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector2) Cross(with Vector2) float64 {
	ret := NewFloat()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{with.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodCrossFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector2) Abs() Vector2 {
	ret := NewVector2()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodAbsFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector2) Sign() Vector2 {
	ret := NewVector2()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodSignFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector2) Clamp(min Vector2, max Vector2) Vector2 {
	ret := NewVector2()

	args := []gdc.ConstTypePtr{min.AsCTypePtr(), max.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodClampFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector2) Snapped(step Vector2) Vector2 {
	ret := NewVector2()

	args := []gdc.ConstTypePtr{step.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodSnappedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func Vector2FromAngle(angle float64) Vector2 {
	ret := NewVector2()

	varg0 := NewFloatFromFloat32(angle)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2.methodFromAngleFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

// Operators

func (me *Vector2) EqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector2) NotEqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector2) Negate() Vector2 {
	opPtr := ptrsForVector2.operatorNegateFn
	ret := NewVector2()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return *ret
}

func (me *Vector2) Positive() Vector2 {
	opPtr := ptrsForVector2.operatorPositiveFn
	ret := NewVector2()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return *ret
}

func (me *Vector2) Not() bool {
	opPtr := ptrsForVector2.operatorNotFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector2) MultiplyInt(rightArg int64) Vector2 {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForVector2.operatorMultiplyIntFn
	ret := NewVector2()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector2) DivideInt(rightArg int64) Vector2 {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForVector2.operatorDivideIntFn
	ret := NewVector2()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector2) MultiplyFloat32(rightArg float64) Vector2 {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForVector2.operatorMultiplyFloat32Fn
	ret := NewVector2()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector2) DivideFloat32(rightArg float64) Vector2 {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForVector2.operatorDivideFloat32Fn
	ret := NewVector2()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector2) EqualVector2(right Vector2) bool {
	opPtr := ptrsForVector2.operatorEqualVector2Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector2) NotEqualVector2(right Vector2) bool {
	opPtr := ptrsForVector2.operatorNotEqualVector2Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector2) LessVector2(right Vector2) bool {
	opPtr := ptrsForVector2.operatorLessVector2Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector2) LessEqualVector2(right Vector2) bool {
	opPtr := ptrsForVector2.operatorLessEqualVector2Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector2) GreaterVector2(right Vector2) bool {
	opPtr := ptrsForVector2.operatorGreaterVector2Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector2) GreaterEqualVector2(right Vector2) bool {
	opPtr := ptrsForVector2.operatorGreaterEqualVector2Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector2) AddVector2(right Vector2) Vector2 {
	opPtr := ptrsForVector2.operatorAddVector2Fn
	ret := NewVector2()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector2) SubtractVector2(right Vector2) Vector2 {
	opPtr := ptrsForVector2.operatorSubtractVector2Fn
	ret := NewVector2()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector2) MultiplyVector2(right Vector2) Vector2 {
	opPtr := ptrsForVector2.operatorMultiplyVector2Fn
	ret := NewVector2()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector2) DivideVector2(right Vector2) Vector2 {
	opPtr := ptrsForVector2.operatorDivideVector2Fn
	ret := NewVector2()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector2) MultiplyTransform2D(right Transform2D) Vector2 {
	opPtr := ptrsForVector2.operatorMultiplyTransform2DFn
	ret := NewVector2()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector2) InDictionary(right Dictionary) bool {
	opPtr := ptrsForVector2.operatorInDictionaryFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector2) InArray(right Array) bool {
	opPtr := ptrsForVector2.operatorInArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector2) InPackedVector2Array(right PackedVector2Array) bool {
	opPtr := ptrsForVector2.operatorInPackedVector2ArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

// Members

func (me *Vector2) X() float64 {
	ret := NewFloat()
	me.iface.CallPtrGetter(ensurePtr(ptrsForVector2.memberxGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector2) SetX(value float64) {
	valueV := NewFloatFromFloat32(value)
	defer valueV.Destroy()
	me.iface.CallPtrSetter(ensurePtr(ptrsForVector2.memberxSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Vector2) Y() float64 {
	ret := NewFloat()
	me.iface.CallPtrGetter(ensurePtr(ptrsForVector2.memberyGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector2) SetY(value float64) {
	valueV := NewFloatFromFloat32(value)
	defer valueV.Destroy()
	me.iface.CallPtrSetter(ensurePtr(ptrsForVector2.memberySetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}
