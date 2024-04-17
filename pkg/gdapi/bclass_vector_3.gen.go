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

type ptrsForVector3List struct {
	ctrFn                          gdc.PtrConstructor
	ctrFromVector3Fn               gdc.PtrConstructor
	ctrFromVector3IFn              gdc.PtrConstructor
	ctrFromFloat32Float32Float32Fn gdc.PtrConstructor
	methodMinAxisIndexFn           gdc.PtrBuiltInMethod
	methodMaxAxisIndexFn           gdc.PtrBuiltInMethod
	methodAngleToFn                gdc.PtrBuiltInMethod
	methodSignedAngleToFn          gdc.PtrBuiltInMethod
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
	methodInverseFn                gdc.PtrBuiltInMethod
	methodClampFn                  gdc.PtrBuiltInMethod
	methodSnappedFn                gdc.PtrBuiltInMethod
	methodRotatedFn                gdc.PtrBuiltInMethod
	methodLerpFn                   gdc.PtrBuiltInMethod
	methodSlerpFn                  gdc.PtrBuiltInMethod
	methodCubicInterpolateFn       gdc.PtrBuiltInMethod
	methodCubicInterpolateInTimeFn gdc.PtrBuiltInMethod
	methodBezierInterpolateFn      gdc.PtrBuiltInMethod
	methodBezierDerivativeFn       gdc.PtrBuiltInMethod
	methodMoveTowardFn             gdc.PtrBuiltInMethod
	methodDotFn                    gdc.PtrBuiltInMethod
	methodCrossFn                  gdc.PtrBuiltInMethod
	methodOuterFn                  gdc.PtrBuiltInMethod
	methodAbsFn                    gdc.PtrBuiltInMethod
	methodFloorFn                  gdc.PtrBuiltInMethod
	methodCeilFn                   gdc.PtrBuiltInMethod
	methodRoundFn                  gdc.PtrBuiltInMethod
	methodPosmodFn                 gdc.PtrBuiltInMethod
	methodPosmodvFn                gdc.PtrBuiltInMethod
	methodProjectFn                gdc.PtrBuiltInMethod
	methodSlideFn                  gdc.PtrBuiltInMethod
	methodBounceFn                 gdc.PtrBuiltInMethod
	methodReflectFn                gdc.PtrBuiltInMethod
	methodSignFn                   gdc.PtrBuiltInMethod
	methodOctahedronEncodeFn       gdc.PtrBuiltInMethod
	methodOctahedronDecodeFn       gdc.PtrBuiltInMethod
	operatorNegateFn               gdc.PtrOperatorEvaluator
	operatorPositiveFn             gdc.PtrOperatorEvaluator
	operatorNotFn                  gdc.PtrOperatorEvaluator
	operatorMultiplyIntFn          gdc.PtrOperatorEvaluator
	operatorDivideIntFn            gdc.PtrOperatorEvaluator
	operatorMultiplyFloat32Fn      gdc.PtrOperatorEvaluator
	operatorDivideFloat32Fn        gdc.PtrOperatorEvaluator
	operatorEqualVector3Fn         gdc.PtrOperatorEvaluator
	operatorNotEqualVector3Fn      gdc.PtrOperatorEvaluator
	operatorLessVector3Fn          gdc.PtrOperatorEvaluator
	operatorLessEqualVector3Fn     gdc.PtrOperatorEvaluator
	operatorGreaterVector3Fn       gdc.PtrOperatorEvaluator
	operatorGreaterEqualVector3Fn  gdc.PtrOperatorEvaluator
	operatorAddVector3Fn           gdc.PtrOperatorEvaluator
	operatorSubtractVector3Fn      gdc.PtrOperatorEvaluator
	operatorMultiplyVector3Fn      gdc.PtrOperatorEvaluator
	operatorDivideVector3Fn        gdc.PtrOperatorEvaluator
	operatorMultiplyQuaternionFn   gdc.PtrOperatorEvaluator
	operatorMultiplyBasisFn        gdc.PtrOperatorEvaluator
	operatorMultiplyTransform3DFn  gdc.PtrOperatorEvaluator
	operatorInDictionaryFn         gdc.PtrOperatorEvaluator
	operatorInArrayFn              gdc.PtrOperatorEvaluator
	operatorInPackedVector3ArrayFn gdc.PtrOperatorEvaluator
	memberxGetterFn                gdc.PtrGetter
	memberxSetterFn                gdc.PtrSetter
	memberyGetterFn                gdc.PtrGetter
	memberySetterFn                gdc.PtrSetter
	memberzGetterFn                gdc.PtrGetter
	memberzSetterFn                gdc.PtrSetter
	toVariantFn                    gdc.TypeFromVariantConstructorFunc
	fromVariantFn                  gdc.VariantFromTypeConstructorFunc
}

var ptrsForVector3 ptrsForVector3List

func initVector3Ptrs(iface gdc.Interface) {
	ptrsForVector3.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeVector3, 0))
	ptrsForVector3.ctrFromVector3Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeVector3, 1))
	ptrsForVector3.ctrFromVector3IFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeVector3, 2))
	ptrsForVector3.ctrFromFloat32Float32Float32Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeVector3, 3))
	{
		methodName := StringNameFromStr("min_axis_index")
		defer methodName.Destroy()
		ptrsForVector3.methodMinAxisIndexFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("max_axis_index")
		defer methodName.Destroy()
		ptrsForVector3.methodMaxAxisIndexFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("angle_to")
		defer methodName.Destroy()
		ptrsForVector3.methodAngleToFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 1047977935))
	}
	{
		methodName := StringNameFromStr("signed_angle_to")
		defer methodName.Destroy()
		ptrsForVector3.methodSignedAngleToFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 2781412522))
	}
	{
		methodName := StringNameFromStr("direction_to")
		defer methodName.Destroy()
		ptrsForVector3.methodDirectionToFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 2923479887))
	}
	{
		methodName := StringNameFromStr("distance_to")
		defer methodName.Destroy()
		ptrsForVector3.methodDistanceToFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 1047977935))
	}
	{
		methodName := StringNameFromStr("distance_squared_to")
		defer methodName.Destroy()
		ptrsForVector3.methodDistanceSquaredToFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 1047977935))
	}
	{
		methodName := StringNameFromStr("length")
		defer methodName.Destroy()
		ptrsForVector3.methodLengthFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 466405837))
	}
	{
		methodName := StringNameFromStr("length_squared")
		defer methodName.Destroy()
		ptrsForVector3.methodLengthSquaredFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 466405837))
	}
	{
		methodName := StringNameFromStr("limit_length")
		defer methodName.Destroy()
		ptrsForVector3.methodLimitLengthFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 514930144))
	}
	{
		methodName := StringNameFromStr("normalized")
		defer methodName.Destroy()
		ptrsForVector3.methodNormalizedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 1776574132))
	}
	{
		methodName := StringNameFromStr("is_normalized")
		defer methodName.Destroy()
		ptrsForVector3.methodIsNormalizedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("is_equal_approx")
		defer methodName.Destroy()
		ptrsForVector3.methodIsEqualApproxFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 1749054343))
	}
	{
		methodName := StringNameFromStr("is_zero_approx")
		defer methodName.Destroy()
		ptrsForVector3.methodIsZeroApproxFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("is_finite")
		defer methodName.Destroy()
		ptrsForVector3.methodIsFiniteFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("inverse")
		defer methodName.Destroy()
		ptrsForVector3.methodInverseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 1776574132))
	}
	{
		methodName := StringNameFromStr("clamp")
		defer methodName.Destroy()
		ptrsForVector3.methodClampFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 4145107892))
	}
	{
		methodName := StringNameFromStr("snapped")
		defer methodName.Destroy()
		ptrsForVector3.methodSnappedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 2923479887))
	}
	{
		methodName := StringNameFromStr("rotated")
		defer methodName.Destroy()
		ptrsForVector3.methodRotatedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 1682608829))
	}
	{
		methodName := StringNameFromStr("lerp")
		defer methodName.Destroy()
		ptrsForVector3.methodLerpFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 1682608829))
	}
	{
		methodName := StringNameFromStr("slerp")
		defer methodName.Destroy()
		ptrsForVector3.methodSlerpFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 1682608829))
	}
	{
		methodName := StringNameFromStr("cubic_interpolate")
		defer methodName.Destroy()
		ptrsForVector3.methodCubicInterpolateFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 2597922253))
	}
	{
		methodName := StringNameFromStr("cubic_interpolate_in_time")
		defer methodName.Destroy()
		ptrsForVector3.methodCubicInterpolateInTimeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 3256682901))
	}
	{
		methodName := StringNameFromStr("bezier_interpolate")
		defer methodName.Destroy()
		ptrsForVector3.methodBezierInterpolateFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 2597922253))
	}
	{
		methodName := StringNameFromStr("bezier_derivative")
		defer methodName.Destroy()
		ptrsForVector3.methodBezierDerivativeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 2597922253))
	}
	{
		methodName := StringNameFromStr("move_toward")
		defer methodName.Destroy()
		ptrsForVector3.methodMoveTowardFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 1682608829))
	}
	{
		methodName := StringNameFromStr("dot")
		defer methodName.Destroy()
		ptrsForVector3.methodDotFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 1047977935))
	}
	{
		methodName := StringNameFromStr("cross")
		defer methodName.Destroy()
		ptrsForVector3.methodCrossFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 2923479887))
	}
	{
		methodName := StringNameFromStr("outer")
		defer methodName.Destroy()
		ptrsForVector3.methodOuterFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 3934786792))
	}
	{
		methodName := StringNameFromStr("abs")
		defer methodName.Destroy()
		ptrsForVector3.methodAbsFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 1776574132))
	}
	{
		methodName := StringNameFromStr("floor")
		defer methodName.Destroy()
		ptrsForVector3.methodFloorFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 1776574132))
	}
	{
		methodName := StringNameFromStr("ceil")
		defer methodName.Destroy()
		ptrsForVector3.methodCeilFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 1776574132))
	}
	{
		methodName := StringNameFromStr("round")
		defer methodName.Destroy()
		ptrsForVector3.methodRoundFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 1776574132))
	}
	{
		methodName := StringNameFromStr("posmod")
		defer methodName.Destroy()
		ptrsForVector3.methodPosmodFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 514930144))
	}
	{
		methodName := StringNameFromStr("posmodv")
		defer methodName.Destroy()
		ptrsForVector3.methodPosmodvFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 2923479887))
	}
	{
		methodName := StringNameFromStr("project")
		defer methodName.Destroy()
		ptrsForVector3.methodProjectFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 2923479887))
	}
	{
		methodName := StringNameFromStr("slide")
		defer methodName.Destroy()
		ptrsForVector3.methodSlideFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 2923479887))
	}
	{
		methodName := StringNameFromStr("bounce")
		defer methodName.Destroy()
		ptrsForVector3.methodBounceFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 2923479887))
	}
	{
		methodName := StringNameFromStr("reflect")
		defer methodName.Destroy()
		ptrsForVector3.methodReflectFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 2923479887))
	}
	{
		methodName := StringNameFromStr("sign")
		defer methodName.Destroy()
		ptrsForVector3.methodSignFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 1776574132))
	}
	{
		methodName := StringNameFromStr("octahedron_encode")
		defer methodName.Destroy()
		ptrsForVector3.methodOctahedronEncodeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 2428350749))
	}
	{
		methodName := StringNameFromStr("octahedron_decode")
		defer methodName.Destroy()
		ptrsForVector3.methodOctahedronDecodeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3, methodName.AsCPtr(), 3991820552))
	}
	ptrsForVector3.operatorNegateFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNegate, gdc.VariantTypeVector3, gdc.VariantTypeNil))
	ptrsForVector3.operatorPositiveFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpPositive, gdc.VariantTypeVector3, gdc.VariantTypeNil))
	ptrsForVector3.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypeVector3, gdc.VariantTypeNil))
	ptrsForVector3.operatorMultiplyIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeVector3, gdc.VariantTypeInt))
	ptrsForVector3.operatorDivideIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeVector3, gdc.VariantTypeInt))
	ptrsForVector3.operatorMultiplyFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeVector3, gdc.VariantTypeFloat))
	ptrsForVector3.operatorDivideFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeVector3, gdc.VariantTypeFloat))
	ptrsForVector3.operatorEqualVector3Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeVector3, gdc.VariantTypeVector3))
	ptrsForVector3.operatorNotEqualVector3Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeVector3, gdc.VariantTypeVector3))
	ptrsForVector3.operatorLessVector3Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLess, gdc.VariantTypeVector3, gdc.VariantTypeVector3))
	ptrsForVector3.operatorLessEqualVector3Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLessEqual, gdc.VariantTypeVector3, gdc.VariantTypeVector3))
	ptrsForVector3.operatorGreaterVector3Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreater, gdc.VariantTypeVector3, gdc.VariantTypeVector3))
	ptrsForVector3.operatorGreaterEqualVector3Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreaterEqual, gdc.VariantTypeVector3, gdc.VariantTypeVector3))
	ptrsForVector3.operatorAddVector3Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, gdc.VariantTypeVector3, gdc.VariantTypeVector3))
	ptrsForVector3.operatorSubtractVector3Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpSubtract, gdc.VariantTypeVector3, gdc.VariantTypeVector3))
	ptrsForVector3.operatorMultiplyVector3Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeVector3, gdc.VariantTypeVector3))
	ptrsForVector3.operatorDivideVector3Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeVector3, gdc.VariantTypeVector3))
	ptrsForVector3.operatorMultiplyQuaternionFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeVector3, gdc.VariantTypeQuaternion))
	ptrsForVector3.operatorMultiplyBasisFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeVector3, gdc.VariantTypeBasis))
	ptrsForVector3.operatorMultiplyTransform3DFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeVector3, gdc.VariantTypeTransform3D))
	ptrsForVector3.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeVector3, gdc.VariantTypeDictionary))
	ptrsForVector3.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeVector3, gdc.VariantTypeArray))
	ptrsForVector3.operatorInPackedVector3ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeVector3, gdc.VariantTypePackedVector3Array))
	{
		memberName := StringNameFromStr("x")
		defer memberName.Destroy()
		ptrsForVector3.memberxGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeVector3, memberName.AsCPtr()))
		ptrsForVector3.memberxSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeVector3, memberName.AsCPtr()))
	}
	{
		memberName := StringNameFromStr("y")
		defer memberName.Destroy()
		ptrsForVector3.memberyGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeVector3, memberName.AsCPtr()))
		ptrsForVector3.memberySetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeVector3, memberName.AsCPtr()))
	}
	{
		memberName := StringNameFromStr("z")
		defer memberName.Destroy()
		ptrsForVector3.memberzGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeVector3, memberName.AsCPtr()))
		ptrsForVector3.memberzSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeVector3, memberName.AsCPtr()))
	}
	ptrsForVector3.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypeVector3))
	ptrsForVector3.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypeVector3))
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("ZERO")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector3, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector3()
		if err != nil {
			panic("Failed to get constant value ZERO: " + err.Error())
		}
		Vector3Zero = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("ONE")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector3, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector3()
		if err != nil {
			panic("Failed to get constant value ONE: " + err.Error())
		}
		Vector3One = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("INF")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector3, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector3()
		if err != nil {
			panic("Failed to get constant value INF: " + err.Error())
		}
		Vector3Inf = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("LEFT")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector3, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector3()
		if err != nil {
			panic("Failed to get constant value LEFT: " + err.Error())
		}
		Vector3Left = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("RIGHT")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector3, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector3()
		if err != nil {
			panic("Failed to get constant value RIGHT: " + err.Error())
		}
		Vector3Right = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("UP")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector3, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector3()
		if err != nil {
			panic("Failed to get constant value UP: " + err.Error())
		}
		Vector3Up = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("DOWN")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector3, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector3()
		if err != nil {
			panic("Failed to get constant value DOWN: " + err.Error())
		}
		Vector3Down = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("FORWARD")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector3, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector3()
		if err != nil {
			panic("Failed to get constant value FORWARD: " + err.Error())
		}
		Vector3Forward = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("BACK")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector3, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector3()
		if err != nil {
			panic("Failed to get constant value BACK: " + err.Error())
		}
		Vector3Back = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("MODEL_LEFT")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector3, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector3()
		if err != nil {
			panic("Failed to get constant value MODEL_LEFT: " + err.Error())
		}
		Vector3ModelLeft = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("MODEL_RIGHT")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector3, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector3()
		if err != nil {
			panic("Failed to get constant value MODEL_RIGHT: " + err.Error())
		}
		Vector3ModelRight = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("MODEL_TOP")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector3, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector3()
		if err != nil {
			panic("Failed to get constant value MODEL_TOP: " + err.Error())
		}
		Vector3ModelTop = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("MODEL_BOTTOM")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector3, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector3()
		if err != nil {
			panic("Failed to get constant value MODEL_BOTTOM: " + err.Error())
		}
		Vector3ModelBottom = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("MODEL_FRONT")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector3, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector3()
		if err != nil {
			panic("Failed to get constant value MODEL_FRONT: " + err.Error())
		}
		Vector3ModelFront = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("MODEL_REAR")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector3, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector3()
		if err != nil {
			panic("Failed to get constant value MODEL_REAR: " + err.Error())
		}
		Vector3ModelRear = *cnst
	}
}

type Vector3 struct {
	data   *[classSizeVector3]byte
	iface  gdc.Interface
	pinner runtime.Pinner
}

// Constants

var (
	Vector3Zero        Vector3
	Vector3One         Vector3
	Vector3Inf         Vector3
	Vector3Left        Vector3
	Vector3Right       Vector3
	Vector3Up          Vector3
	Vector3Down        Vector3
	Vector3Forward     Vector3
	Vector3Back        Vector3
	Vector3ModelLeft   Vector3
	Vector3ModelRight  Vector3
	Vector3ModelTop    Vector3
	Vector3ModelBottom Vector3
	Vector3ModelFront  Vector3
	Vector3ModelRear   Vector3
)

// Enums

type Vector3Axis int

const (
	Vector3AxisAxisX Vector3Axis = 0
	Vector3AxisAxisY Vector3Axis = 1
	Vector3AxisAxisZ Vector3Axis = 2
)

// Constructors
func newVector3() *Vector3 {
	me := &Vector3{
		data:  new([classSizeVector3]byte),
		iface: giface,
	}
	me.pinner.Pin(me)
	me.pinner.Pin(me.AsTypePtr())
	return me
}

func NewVector3() *Vector3 {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newVector3()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForVector3.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
	return me
}

func NewVector3FromVector3(from Vector3) *Vector3 {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newVector3()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForVector3.ctrFromVector3Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewVector3FromVector3I(from Vector3i) *Vector3 {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newVector3()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForVector3.ctrFromVector3IFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewVector3FromFloat32Float32Float32(x float64, y float64, z float64) *Vector3 {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	pinner.Pin(&x)
	pinner.Pin(&y)
	pinner.Pin(&z)
	me := newVector3()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForVector3.ctrFromFloat32Float32Float32Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), gdc.ConstTypePtr(&z)}))
	return me
}

// Destructor
func (me *Vector3) Destroy() {
	me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsVector3() (*Vector3, error) {
	if me.Type() != gdc.VariantTypeVector3 {
		return nil, fmt.Errorf("variant is not a Vector3")
	}
	bclass := newVector3()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForVector3.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Vector3) AsVariant() *Variant {
	va := newVariant()
	va.inner = me
	me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForVector3.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
	return va
}

// Pointers
func Vector3FromPtr(ptr gdc.ConstTypePtr) *Vector3 {
	me := newVector3()
	dataFromPtr(me.data[:], ptr)
	return me
}

func (me *Vector3) ToTypePtr(ptr gdc.TypePtr) {
	dataToPtr(ptr, me.data[:])
}

func (me *Vector3) Type() gdc.VariantType {
	return gdc.VariantTypeVector3
}

func (me *Vector3) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Vector3) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Vector3) asUninitialized() gdc.UninitializedTypePtr {
	return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *Vector3) MinAxisIndex() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodMinAxisIndexFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector3) MaxAxisIndex() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodMaxAxisIndexFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector3) AngleTo(to Vector3) float64 {
	ret := NewFloat()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{to.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodAngleToFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector3) SignedAngleTo(to Vector3, axis Vector3) float64 {
	ret := NewFloat()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{to.AsCTypePtr(), axis.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodSignedAngleToFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector3) DirectionTo(to Vector3) Vector3 {
	ret := NewVector3()

	args := []gdc.ConstTypePtr{to.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodDirectionToFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3) DistanceTo(to Vector3) float64 {
	ret := NewFloat()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{to.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodDistanceToFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector3) DistanceSquaredTo(to Vector3) float64 {
	ret := NewFloat()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{to.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodDistanceSquaredToFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector3) Length() float64 {
	ret := NewFloat()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodLengthFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector3) LengthSquared() float64 {
	ret := NewFloat()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodLengthSquaredFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector3) LimitLength(length float64) Vector3 {
	ret := NewVector3()

	varg0 := NewFloatFromFloat32(length)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodLimitLengthFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3) Normalized() Vector3 {
	ret := NewVector3()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodNormalizedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3) IsNormalized() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodIsNormalizedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector3) IsEqualApprox(to Vector3) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{to.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodIsEqualApproxFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector3) IsZeroApprox() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodIsZeroApproxFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector3) IsFinite() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodIsFiniteFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector3) Inverse() Vector3 {
	ret := NewVector3()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodInverseFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3) Clamp(min Vector3, max Vector3) Vector3 {
	ret := NewVector3()

	args := []gdc.ConstTypePtr{min.AsCTypePtr(), max.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodClampFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3) Snapped(step Vector3) Vector3 {
	ret := NewVector3()

	args := []gdc.ConstTypePtr{step.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodSnappedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3) Rotated(axis Vector3, angle float64) Vector3 {
	ret := NewVector3()

	varg1 := NewFloatFromFloat32(angle)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{axis.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodRotatedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3) Lerp(to Vector3, weight float64) Vector3 {
	ret := NewVector3()

	varg1 := NewFloatFromFloat32(weight)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{to.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodLerpFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3) Slerp(to Vector3, weight float64) Vector3 {
	ret := NewVector3()

	varg1 := NewFloatFromFloat32(weight)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{to.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodSlerpFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3) CubicInterpolate(b Vector3, pre_a Vector3, post_b Vector3, weight float64) Vector3 {
	ret := NewVector3()

	varg3 := NewFloatFromFloat32(weight)
	defer varg3.Destroy()
	args := []gdc.ConstTypePtr{b.AsCTypePtr(), pre_a.AsCTypePtr(), post_b.AsCTypePtr(), varg3.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodCubicInterpolateFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3) CubicInterpolateInTime(b Vector3, pre_a Vector3, post_b Vector3, weight float64, b_t float64, pre_a_t float64, post_b_t float64) Vector3 {
	ret := NewVector3()

	varg3 := NewFloatFromFloat32(weight)
	defer varg3.Destroy()
	varg4 := NewFloatFromFloat32(b_t)
	defer varg4.Destroy()
	varg5 := NewFloatFromFloat32(pre_a_t)
	defer varg5.Destroy()
	varg6 := NewFloatFromFloat32(post_b_t)
	defer varg6.Destroy()
	args := []gdc.ConstTypePtr{b.AsCTypePtr(), pre_a.AsCTypePtr(), post_b.AsCTypePtr(), varg3.AsCTypePtr(), varg4.AsCTypePtr(), varg5.AsCTypePtr(), varg6.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodCubicInterpolateInTimeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3) BezierInterpolate(control_1 Vector3, control_2 Vector3, end Vector3, t float64) Vector3 {
	ret := NewVector3()

	varg3 := NewFloatFromFloat32(t)
	defer varg3.Destroy()
	args := []gdc.ConstTypePtr{control_1.AsCTypePtr(), control_2.AsCTypePtr(), end.AsCTypePtr(), varg3.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodBezierInterpolateFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3) BezierDerivative(control_1 Vector3, control_2 Vector3, end Vector3, t float64) Vector3 {
	ret := NewVector3()

	varg3 := NewFloatFromFloat32(t)
	defer varg3.Destroy()
	args := []gdc.ConstTypePtr{control_1.AsCTypePtr(), control_2.AsCTypePtr(), end.AsCTypePtr(), varg3.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodBezierDerivativeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3) MoveToward(to Vector3, delta float64) Vector3 {
	ret := NewVector3()

	varg1 := NewFloatFromFloat32(delta)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{to.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodMoveTowardFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3) Dot(with Vector3) float64 {
	ret := NewFloat()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{with.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodDotFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector3) Cross(with Vector3) Vector3 {
	ret := NewVector3()

	args := []gdc.ConstTypePtr{with.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodCrossFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3) Outer(with Vector3) Basis {
	ret := NewBasis()

	args := []gdc.ConstTypePtr{with.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodOuterFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3) Abs() Vector3 {
	ret := NewVector3()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodAbsFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3) Floor() Vector3 {
	ret := NewVector3()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodFloorFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3) Ceil() Vector3 {
	ret := NewVector3()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodCeilFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3) Round() Vector3 {
	ret := NewVector3()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodRoundFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3) Posmod(mod float64) Vector3 {
	ret := NewVector3()

	varg0 := NewFloatFromFloat32(mod)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodPosmodFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3) Posmodv(modv Vector3) Vector3 {
	ret := NewVector3()

	args := []gdc.ConstTypePtr{modv.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodPosmodvFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3) Project(b Vector3) Vector3 {
	ret := NewVector3()

	args := []gdc.ConstTypePtr{b.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodProjectFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3) Slide(n Vector3) Vector3 {
	ret := NewVector3()

	args := []gdc.ConstTypePtr{n.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodSlideFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3) Bounce(n Vector3) Vector3 {
	ret := NewVector3()

	args := []gdc.ConstTypePtr{n.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodBounceFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3) Reflect(direction Vector3) Vector3 {
	ret := NewVector3()

	args := []gdc.ConstTypePtr{direction.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodReflectFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3) Sign() Vector3 {
	ret := NewVector3()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodSignFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3) OctahedronEncode() Vector2 {
	ret := NewVector2()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodOctahedronEncodeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func Vector3OctahedronDecode(uv Vector2) Vector3 {
	ret := NewVector3()

	args := []gdc.ConstTypePtr{uv.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3.methodOctahedronDecodeFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

// Operators

func (me *Vector3) EqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector3) NotEqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector3) Negate() Vector3 {
	opPtr := ptrsForVector3.operatorNegateFn
	ret := NewVector3()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return *ret
}

func (me *Vector3) Positive() Vector3 {
	opPtr := ptrsForVector3.operatorPositiveFn
	ret := NewVector3()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return *ret
}

func (me *Vector3) Not() bool {
	opPtr := ptrsForVector3.operatorNotFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector3) MultiplyInt(rightArg int64) Vector3 {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForVector3.operatorMultiplyIntFn
	ret := NewVector3()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector3) DivideInt(rightArg int64) Vector3 {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForVector3.operatorDivideIntFn
	ret := NewVector3()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector3) MultiplyFloat32(rightArg float64) Vector3 {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForVector3.operatorMultiplyFloat32Fn
	ret := NewVector3()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector3) DivideFloat32(rightArg float64) Vector3 {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForVector3.operatorDivideFloat32Fn
	ret := NewVector3()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector3) EqualVector3(right Vector3) bool {
	opPtr := ptrsForVector3.operatorEqualVector3Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector3) NotEqualVector3(right Vector3) bool {
	opPtr := ptrsForVector3.operatorNotEqualVector3Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector3) LessVector3(right Vector3) bool {
	opPtr := ptrsForVector3.operatorLessVector3Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector3) LessEqualVector3(right Vector3) bool {
	opPtr := ptrsForVector3.operatorLessEqualVector3Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector3) GreaterVector3(right Vector3) bool {
	opPtr := ptrsForVector3.operatorGreaterVector3Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector3) GreaterEqualVector3(right Vector3) bool {
	opPtr := ptrsForVector3.operatorGreaterEqualVector3Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector3) AddVector3(right Vector3) Vector3 {
	opPtr := ptrsForVector3.operatorAddVector3Fn
	ret := NewVector3()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector3) SubtractVector3(right Vector3) Vector3 {
	opPtr := ptrsForVector3.operatorSubtractVector3Fn
	ret := NewVector3()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector3) MultiplyVector3(right Vector3) Vector3 {
	opPtr := ptrsForVector3.operatorMultiplyVector3Fn
	ret := NewVector3()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector3) DivideVector3(right Vector3) Vector3 {
	opPtr := ptrsForVector3.operatorDivideVector3Fn
	ret := NewVector3()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector3) MultiplyQuaternion(right Quaternion) Vector3 {
	opPtr := ptrsForVector3.operatorMultiplyQuaternionFn
	ret := NewVector3()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector3) MultiplyBasis(right Basis) Vector3 {
	opPtr := ptrsForVector3.operatorMultiplyBasisFn
	ret := NewVector3()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector3) MultiplyTransform3D(right Transform3D) Vector3 {
	opPtr := ptrsForVector3.operatorMultiplyTransform3DFn
	ret := NewVector3()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector3) InDictionary(right Dictionary) bool {
	opPtr := ptrsForVector3.operatorInDictionaryFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector3) InArray(right Array) bool {
	opPtr := ptrsForVector3.operatorInArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector3) InPackedVector3Array(right PackedVector3Array) bool {
	opPtr := ptrsForVector3.operatorInPackedVector3ArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

// Members

func (me *Vector3) X() float64 {
	ret := NewFloat()
	me.iface.CallPtrGetter(ensurePtr(ptrsForVector3.memberxGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector3) SetX(value float64) {
	valueV := NewFloatFromFloat32(value)
	defer valueV.Destroy()
	me.iface.CallPtrSetter(ensurePtr(ptrsForVector3.memberxSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Vector3) Y() float64 {
	ret := NewFloat()
	me.iface.CallPtrGetter(ensurePtr(ptrsForVector3.memberyGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector3) SetY(value float64) {
	valueV := NewFloatFromFloat32(value)
	defer valueV.Destroy()
	me.iface.CallPtrSetter(ensurePtr(ptrsForVector3.memberySetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Vector3) Z() float64 {
	ret := NewFloat()
	me.iface.CallPtrGetter(ensurePtr(ptrsForVector3.memberzGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector3) SetZ(value float64) {
	valueV := NewFloatFromFloat32(value)
	defer valueV.Destroy()
	me.iface.CallPtrSetter(ensurePtr(ptrsForVector3.memberzSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}
