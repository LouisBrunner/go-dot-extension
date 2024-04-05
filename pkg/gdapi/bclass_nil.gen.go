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

type ptrsForNilList struct {
  ctrFn gdc.PtrConstructor
  ctrFromVariantFn gdc.PtrConstructor
  operatorNotFn gdc.PtrOperatorEvaluator
  operatorEqualBoolFn gdc.PtrOperatorEvaluator
  operatorNotEqualBoolFn gdc.PtrOperatorEvaluator
  operatorAndBoolFn gdc.PtrOperatorEvaluator
  operatorOrBoolFn gdc.PtrOperatorEvaluator
  operatorXorBoolFn gdc.PtrOperatorEvaluator
  operatorEqualIntFn gdc.PtrOperatorEvaluator
  operatorNotEqualIntFn gdc.PtrOperatorEvaluator
  operatorAndIntFn gdc.PtrOperatorEvaluator
  operatorOrIntFn gdc.PtrOperatorEvaluator
  operatorXorIntFn gdc.PtrOperatorEvaluator
  operatorEqualFloat32Fn gdc.PtrOperatorEvaluator
  operatorNotEqualFloat32Fn gdc.PtrOperatorEvaluator
  operatorAndFloat32Fn gdc.PtrOperatorEvaluator
  operatorOrFloat32Fn gdc.PtrOperatorEvaluator
  operatorXorFloat32Fn gdc.PtrOperatorEvaluator
  operatorEqualStringFn gdc.PtrOperatorEvaluator
  operatorNotEqualStringFn gdc.PtrOperatorEvaluator
  operatorEqualVector2Fn gdc.PtrOperatorEvaluator
  operatorNotEqualVector2Fn gdc.PtrOperatorEvaluator
  operatorEqualVector2IFn gdc.PtrOperatorEvaluator
  operatorNotEqualVector2IFn gdc.PtrOperatorEvaluator
  operatorEqualRect2Fn gdc.PtrOperatorEvaluator
  operatorNotEqualRect2Fn gdc.PtrOperatorEvaluator
  operatorEqualRect2IFn gdc.PtrOperatorEvaluator
  operatorNotEqualRect2IFn gdc.PtrOperatorEvaluator
  operatorEqualVector3Fn gdc.PtrOperatorEvaluator
  operatorNotEqualVector3Fn gdc.PtrOperatorEvaluator
  operatorEqualVector3IFn gdc.PtrOperatorEvaluator
  operatorNotEqualVector3IFn gdc.PtrOperatorEvaluator
  operatorEqualTransform2DFn gdc.PtrOperatorEvaluator
  operatorNotEqualTransform2DFn gdc.PtrOperatorEvaluator
  operatorEqualVector4Fn gdc.PtrOperatorEvaluator
  operatorNotEqualVector4Fn gdc.PtrOperatorEvaluator
  operatorEqualVector4IFn gdc.PtrOperatorEvaluator
  operatorNotEqualVector4IFn gdc.PtrOperatorEvaluator
  operatorEqualPlaneFn gdc.PtrOperatorEvaluator
  operatorNotEqualPlaneFn gdc.PtrOperatorEvaluator
  operatorEqualQuaternionFn gdc.PtrOperatorEvaluator
  operatorNotEqualQuaternionFn gdc.PtrOperatorEvaluator
  operatorEqualAABBFn gdc.PtrOperatorEvaluator
  operatorNotEqualAABBFn gdc.PtrOperatorEvaluator
  operatorEqualBasisFn gdc.PtrOperatorEvaluator
  operatorNotEqualBasisFn gdc.PtrOperatorEvaluator
  operatorEqualTransform3DFn gdc.PtrOperatorEvaluator
  operatorNotEqualTransform3DFn gdc.PtrOperatorEvaluator
  operatorEqualProjectionFn gdc.PtrOperatorEvaluator
  operatorNotEqualProjectionFn gdc.PtrOperatorEvaluator
  operatorEqualColorFn gdc.PtrOperatorEvaluator
  operatorNotEqualColorFn gdc.PtrOperatorEvaluator
  operatorEqualStringNameFn gdc.PtrOperatorEvaluator
  operatorNotEqualStringNameFn gdc.PtrOperatorEvaluator
  operatorEqualNodePathFn gdc.PtrOperatorEvaluator
  operatorNotEqualNodePathFn gdc.PtrOperatorEvaluator
  operatorEqualRIDFn gdc.PtrOperatorEvaluator
  operatorNotEqualRIDFn gdc.PtrOperatorEvaluator
  operatorEqualObjectFn gdc.PtrOperatorEvaluator
  operatorNotEqualObjectFn gdc.PtrOperatorEvaluator
  operatorAndObjectFn gdc.PtrOperatorEvaluator
  operatorOrObjectFn gdc.PtrOperatorEvaluator
  operatorXorObjectFn gdc.PtrOperatorEvaluator
  operatorEqualCallableFn gdc.PtrOperatorEvaluator
  operatorNotEqualCallableFn gdc.PtrOperatorEvaluator
  operatorEqualSignalFn gdc.PtrOperatorEvaluator
  operatorNotEqualSignalFn gdc.PtrOperatorEvaluator
  operatorEqualDictionaryFn gdc.PtrOperatorEvaluator
  operatorNotEqualDictionaryFn gdc.PtrOperatorEvaluator
  operatorInDictionaryFn gdc.PtrOperatorEvaluator
  operatorEqualArrayFn gdc.PtrOperatorEvaluator
  operatorNotEqualArrayFn gdc.PtrOperatorEvaluator
  operatorInArrayFn gdc.PtrOperatorEvaluator
  operatorEqualPackedByteArrayFn gdc.PtrOperatorEvaluator
  operatorNotEqualPackedByteArrayFn gdc.PtrOperatorEvaluator
  operatorEqualPackedInt32ArrayFn gdc.PtrOperatorEvaluator
  operatorNotEqualPackedInt32ArrayFn gdc.PtrOperatorEvaluator
  operatorEqualPackedInt64ArrayFn gdc.PtrOperatorEvaluator
  operatorNotEqualPackedInt64ArrayFn gdc.PtrOperatorEvaluator
  operatorEqualPackedFloat32ArrayFn gdc.PtrOperatorEvaluator
  operatorNotEqualPackedFloat32ArrayFn gdc.PtrOperatorEvaluator
  operatorEqualPackedFloat64ArrayFn gdc.PtrOperatorEvaluator
  operatorNotEqualPackedFloat64ArrayFn gdc.PtrOperatorEvaluator
  operatorEqualPackedStringArrayFn gdc.PtrOperatorEvaluator
  operatorNotEqualPackedStringArrayFn gdc.PtrOperatorEvaluator
  operatorEqualPackedVector2ArrayFn gdc.PtrOperatorEvaluator
  operatorNotEqualPackedVector2ArrayFn gdc.PtrOperatorEvaluator
  operatorEqualPackedVector3ArrayFn gdc.PtrOperatorEvaluator
  operatorNotEqualPackedVector3ArrayFn gdc.PtrOperatorEvaluator
  operatorEqualPackedColorArrayFn gdc.PtrOperatorEvaluator
  operatorNotEqualPackedColorArrayFn gdc.PtrOperatorEvaluator
}

var ptrsForNil ptrsForNilList

func initNilPtrs(iface gdc.Interface) {
  ptrsForNil.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeNil, 0))
  ptrsForNil.ctrFromVariantFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeNil, 1))
  ptrsForNil.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypeNil, gdc.VariantTypeNil))
  ptrsForNil.operatorEqualBoolFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypeBool))
  ptrsForNil.operatorNotEqualBoolFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypeBool))
  ptrsForNil.operatorAndBoolFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAnd, gdc.VariantTypeNil, gdc.VariantTypeBool))
  ptrsForNil.operatorOrBoolFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpOr, gdc.VariantTypeNil, gdc.VariantTypeBool))
  ptrsForNil.operatorXorBoolFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpXor, gdc.VariantTypeNil, gdc.VariantTypeBool))
  ptrsForNil.operatorEqualIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypeInt))
  ptrsForNil.operatorNotEqualIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypeInt))
  ptrsForNil.operatorAndIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAnd, gdc.VariantTypeNil, gdc.VariantTypeInt))
  ptrsForNil.operatorOrIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpOr, gdc.VariantTypeNil, gdc.VariantTypeInt))
  ptrsForNil.operatorXorIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpXor, gdc.VariantTypeNil, gdc.VariantTypeInt))
  ptrsForNil.operatorEqualFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypeFloat))
  ptrsForNil.operatorNotEqualFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypeFloat))
  ptrsForNil.operatorAndFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAnd, gdc.VariantTypeNil, gdc.VariantTypeFloat))
  ptrsForNil.operatorOrFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpOr, gdc.VariantTypeNil, gdc.VariantTypeFloat))
  ptrsForNil.operatorXorFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpXor, gdc.VariantTypeNil, gdc.VariantTypeFloat))
  ptrsForNil.operatorEqualStringFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypeString))
  ptrsForNil.operatorNotEqualStringFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypeString))
  ptrsForNil.operatorEqualVector2Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypeVector2))
  ptrsForNil.operatorNotEqualVector2Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypeVector2))
  ptrsForNil.operatorEqualVector2IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypeVector2I))
  ptrsForNil.operatorNotEqualVector2IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypeVector2I))
  ptrsForNil.operatorEqualRect2Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypeRect2))
  ptrsForNil.operatorNotEqualRect2Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypeRect2))
  ptrsForNil.operatorEqualRect2IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypeRect2I))
  ptrsForNil.operatorNotEqualRect2IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypeRect2I))
  ptrsForNil.operatorEqualVector3Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypeVector3))
  ptrsForNil.operatorNotEqualVector3Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypeVector3))
  ptrsForNil.operatorEqualVector3IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypeVector3I))
  ptrsForNil.operatorNotEqualVector3IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypeVector3I))
  ptrsForNil.operatorEqualTransform2DFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypeTransform2D))
  ptrsForNil.operatorNotEqualTransform2DFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypeTransform2D))
  ptrsForNil.operatorEqualVector4Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypeVector4))
  ptrsForNil.operatorNotEqualVector4Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypeVector4))
  ptrsForNil.operatorEqualVector4IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypeVector4I))
  ptrsForNil.operatorNotEqualVector4IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypeVector4I))
  ptrsForNil.operatorEqualPlaneFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypePlane))
  ptrsForNil.operatorNotEqualPlaneFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypePlane))
  ptrsForNil.operatorEqualQuaternionFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypeQuaternion))
  ptrsForNil.operatorNotEqualQuaternionFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypeQuaternion))
  ptrsForNil.operatorEqualAABBFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypeAABB))
  ptrsForNil.operatorNotEqualAABBFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypeAABB))
  ptrsForNil.operatorEqualBasisFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypeBasis))
  ptrsForNil.operatorNotEqualBasisFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypeBasis))
  ptrsForNil.operatorEqualTransform3DFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypeTransform3D))
  ptrsForNil.operatorNotEqualTransform3DFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypeTransform3D))
  ptrsForNil.operatorEqualProjectionFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypeProjection))
  ptrsForNil.operatorNotEqualProjectionFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypeProjection))
  ptrsForNil.operatorEqualColorFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypeColor))
  ptrsForNil.operatorNotEqualColorFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypeColor))
  ptrsForNil.operatorEqualStringNameFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypeStringName))
  ptrsForNil.operatorNotEqualStringNameFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypeStringName))
  ptrsForNil.operatorEqualNodePathFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypeNodePath))
  ptrsForNil.operatorNotEqualNodePathFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypeNodePath))
  ptrsForNil.operatorEqualRIDFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypeRID))
  ptrsForNil.operatorNotEqualRIDFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypeRID))
  ptrsForNil.operatorEqualObjectFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypeObject))
  ptrsForNil.operatorNotEqualObjectFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypeObject))
  ptrsForNil.operatorAndObjectFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAnd, gdc.VariantTypeNil, gdc.VariantTypeObject))
  ptrsForNil.operatorOrObjectFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpOr, gdc.VariantTypeNil, gdc.VariantTypeObject))
  ptrsForNil.operatorXorObjectFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpXor, gdc.VariantTypeNil, gdc.VariantTypeObject))
  ptrsForNil.operatorEqualCallableFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypeCallable))
  ptrsForNil.operatorNotEqualCallableFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypeCallable))
  ptrsForNil.operatorEqualSignalFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypeSignal))
  ptrsForNil.operatorNotEqualSignalFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypeSignal))
  ptrsForNil.operatorEqualDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypeDictionary))
  ptrsForNil.operatorNotEqualDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypeDictionary))
  ptrsForNil.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeNil, gdc.VariantTypeDictionary))
  ptrsForNil.operatorEqualArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypeArray))
  ptrsForNil.operatorNotEqualArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypeArray))
  ptrsForNil.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeNil, gdc.VariantTypeArray))
  ptrsForNil.operatorEqualPackedByteArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypePackedByteArray))
  ptrsForNil.operatorNotEqualPackedByteArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypePackedByteArray))
  ptrsForNil.operatorEqualPackedInt32ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypePackedInt32Array))
  ptrsForNil.operatorNotEqualPackedInt32ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypePackedInt32Array))
  ptrsForNil.operatorEqualPackedInt64ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypePackedInt64Array))
  ptrsForNil.operatorNotEqualPackedInt64ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypePackedInt64Array))
  ptrsForNil.operatorEqualPackedFloat32ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypePackedFloat32Array))
  ptrsForNil.operatorNotEqualPackedFloat32ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypePackedFloat32Array))
  ptrsForNil.operatorEqualPackedFloat64ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypePackedFloat64Array))
  ptrsForNil.operatorNotEqualPackedFloat64ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypePackedFloat64Array))
  ptrsForNil.operatorEqualPackedStringArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypePackedStringArray))
  ptrsForNil.operatorNotEqualPackedStringArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypePackedStringArray))
  ptrsForNil.operatorEqualPackedVector2ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypePackedVector2Array))
  ptrsForNil.operatorNotEqualPackedVector2ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypePackedVector2Array))
  ptrsForNil.operatorEqualPackedVector3ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypePackedVector3Array))
  ptrsForNil.operatorNotEqualPackedVector3ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypePackedVector3Array))
  ptrsForNil.operatorEqualPackedColorArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNil, gdc.VariantTypePackedColorArray))
  ptrsForNil.operatorNotEqualPackedColorArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNil, gdc.VariantTypePackedColorArray))
}

type Nil struct {
  data   *[classSizeNil]byte
  iface  gdc.Interface
  pinner runtime.Pinner
}

// Enums

// Constructors
func newNil() *Nil {
  me := &Nil{
    data:   new([classSizeNil]byte),
    iface:  giface,
  }
  me.pinner.Pin(me)
  me.pinner.Pin(me.AsTypePtr())
  return me
}

func NewNil() *Nil {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newNil()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForNil.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewNilFromVariant(from Variant, ) *Nil {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newNil()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForNil.ctrFromVariantFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

// Destructor
func (me *Nil) Destroy() {
  me.pinner.Unpin()
}

// Pointers
func NilFromPtr(ptr gdc.ConstTypePtr) *Nil {
  me := newNil()
  dataFromPtr(me.data[:], ptr)
  return me
}

func (me *Nil) Type() gdc.VariantType {
  return gdc.VariantTypeNil
}

func (me *Nil) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Nil) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Nil) asUninitialized() gdc.UninitializedTypePtr {
  return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

// Operators

func (me *Nil) EqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) OrVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpOr, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) Not() bool {
  opPtr := ptrsForNil.operatorNotFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualBool(rightArg bool) bool {
  right := NewBoolFromBool(rightArg)
  defer right.Destroy()

  opPtr := ptrsForNil.operatorEqualBoolFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualBool(rightArg bool) bool {
  right := NewBoolFromBool(rightArg)
  defer right.Destroy()

  opPtr := ptrsForNil.operatorNotEqualBoolFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) AndBool(rightArg bool) bool {
  right := NewBoolFromBool(rightArg)
  defer right.Destroy()

  opPtr := ptrsForNil.operatorAndBoolFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) OrBool(rightArg bool) bool {
  right := NewBoolFromBool(rightArg)
  defer right.Destroy()

  opPtr := ptrsForNil.operatorOrBoolFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) XorBool(rightArg bool) bool {
  right := NewBoolFromBool(rightArg)
  defer right.Destroy()

  opPtr := ptrsForNil.operatorXorBoolFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualInt(rightArg int64) bool {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  opPtr := ptrsForNil.operatorEqualIntFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualInt(rightArg int64) bool {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  opPtr := ptrsForNil.operatorNotEqualIntFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) AndInt(rightArg int64) bool {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  opPtr := ptrsForNil.operatorAndIntFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) OrInt(rightArg int64) bool {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  opPtr := ptrsForNil.operatorOrIntFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) XorInt(rightArg int64) bool {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  opPtr := ptrsForNil.operatorXorIntFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualFloat32(rightArg float64) bool {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  opPtr := ptrsForNil.operatorEqualFloat32Fn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualFloat32(rightArg float64) bool {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  opPtr := ptrsForNil.operatorNotEqualFloat32Fn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) AndFloat32(rightArg float64) bool {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  opPtr := ptrsForNil.operatorAndFloat32Fn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) OrFloat32(rightArg float64) bool {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  opPtr := ptrsForNil.operatorOrFloat32Fn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) XorFloat32(rightArg float64) bool {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  opPtr := ptrsForNil.operatorXorFloat32Fn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualString(right String) bool {
  opPtr := ptrsForNil.operatorEqualStringFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualString(right String) bool {
  opPtr := ptrsForNil.operatorNotEqualStringFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualVector2(right Vector2) bool {
  opPtr := ptrsForNil.operatorEqualVector2Fn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualVector2(right Vector2) bool {
  opPtr := ptrsForNil.operatorNotEqualVector2Fn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualVector2I(right Vector2i) bool {
  opPtr := ptrsForNil.operatorEqualVector2IFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualVector2I(right Vector2i) bool {
  opPtr := ptrsForNil.operatorNotEqualVector2IFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualRect2(right Rect2) bool {
  opPtr := ptrsForNil.operatorEqualRect2Fn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualRect2(right Rect2) bool {
  opPtr := ptrsForNil.operatorNotEqualRect2Fn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualRect2I(right Rect2i) bool {
  opPtr := ptrsForNil.operatorEqualRect2IFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualRect2I(right Rect2i) bool {
  opPtr := ptrsForNil.operatorNotEqualRect2IFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualVector3(right Vector3) bool {
  opPtr := ptrsForNil.operatorEqualVector3Fn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualVector3(right Vector3) bool {
  opPtr := ptrsForNil.operatorNotEqualVector3Fn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualVector3I(right Vector3i) bool {
  opPtr := ptrsForNil.operatorEqualVector3IFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualVector3I(right Vector3i) bool {
  opPtr := ptrsForNil.operatorNotEqualVector3IFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualTransform2D(right Transform2D) bool {
  opPtr := ptrsForNil.operatorEqualTransform2DFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualTransform2D(right Transform2D) bool {
  opPtr := ptrsForNil.operatorNotEqualTransform2DFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualVector4(right Vector4) bool {
  opPtr := ptrsForNil.operatorEqualVector4Fn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualVector4(right Vector4) bool {
  opPtr := ptrsForNil.operatorNotEqualVector4Fn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualVector4I(right Vector4i) bool {
  opPtr := ptrsForNil.operatorEqualVector4IFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualVector4I(right Vector4i) bool {
  opPtr := ptrsForNil.operatorNotEqualVector4IFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualPlane(right Plane) bool {
  opPtr := ptrsForNil.operatorEqualPlaneFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualPlane(right Plane) bool {
  opPtr := ptrsForNil.operatorNotEqualPlaneFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualQuaternion(right Quaternion) bool {
  opPtr := ptrsForNil.operatorEqualQuaternionFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualQuaternion(right Quaternion) bool {
  opPtr := ptrsForNil.operatorNotEqualQuaternionFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualAABB(right AABB) bool {
  opPtr := ptrsForNil.operatorEqualAABBFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualAABB(right AABB) bool {
  opPtr := ptrsForNil.operatorNotEqualAABBFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualBasis(right Basis) bool {
  opPtr := ptrsForNil.operatorEqualBasisFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualBasis(right Basis) bool {
  opPtr := ptrsForNil.operatorNotEqualBasisFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualTransform3D(right Transform3D) bool {
  opPtr := ptrsForNil.operatorEqualTransform3DFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualTransform3D(right Transform3D) bool {
  opPtr := ptrsForNil.operatorNotEqualTransform3DFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualProjection(right Projection) bool {
  opPtr := ptrsForNil.operatorEqualProjectionFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualProjection(right Projection) bool {
  opPtr := ptrsForNil.operatorNotEqualProjectionFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualColor(right Color) bool {
  opPtr := ptrsForNil.operatorEqualColorFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualColor(right Color) bool {
  opPtr := ptrsForNil.operatorNotEqualColorFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualStringName(right StringName) bool {
  opPtr := ptrsForNil.operatorEqualStringNameFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualStringName(right StringName) bool {
  opPtr := ptrsForNil.operatorNotEqualStringNameFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualNodePath(right NodePath) bool {
  opPtr := ptrsForNil.operatorEqualNodePathFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualNodePath(right NodePath) bool {
  opPtr := ptrsForNil.operatorNotEqualNodePathFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualRID(right RID) bool {
  opPtr := ptrsForNil.operatorEqualRIDFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualRID(right RID) bool {
  opPtr := ptrsForNil.operatorNotEqualRIDFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualObject(right Object) bool {
  opPtr := ptrsForNil.operatorEqualObjectFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualObject(right Object) bool {
  opPtr := ptrsForNil.operatorNotEqualObjectFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) AndObject(right Object) bool {
  opPtr := ptrsForNil.operatorAndObjectFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) OrObject(right Object) bool {
  opPtr := ptrsForNil.operatorOrObjectFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) XorObject(right Object) bool {
  opPtr := ptrsForNil.operatorXorObjectFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualCallable(right Callable) bool {
  opPtr := ptrsForNil.operatorEqualCallableFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualCallable(right Callable) bool {
  opPtr := ptrsForNil.operatorNotEqualCallableFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualSignal(right Signal) bool {
  opPtr := ptrsForNil.operatorEqualSignalFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualSignal(right Signal) bool {
  opPtr := ptrsForNil.operatorNotEqualSignalFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualDictionary(right Dictionary) bool {
  opPtr := ptrsForNil.operatorEqualDictionaryFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualDictionary(right Dictionary) bool {
  opPtr := ptrsForNil.operatorNotEqualDictionaryFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) InDictionary(right Dictionary) bool {
  opPtr := ptrsForNil.operatorInDictionaryFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualArray(right Array) bool {
  opPtr := ptrsForNil.operatorEqualArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualArray(right Array) bool {
  opPtr := ptrsForNil.operatorNotEqualArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) InArray(right Array) bool {
  opPtr := ptrsForNil.operatorInArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualPackedByteArray(right PackedByteArray) bool {
  opPtr := ptrsForNil.operatorEqualPackedByteArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualPackedByteArray(right PackedByteArray) bool {
  opPtr := ptrsForNil.operatorNotEqualPackedByteArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualPackedInt32Array(right PackedInt32Array) bool {
  opPtr := ptrsForNil.operatorEqualPackedInt32ArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualPackedInt32Array(right PackedInt32Array) bool {
  opPtr := ptrsForNil.operatorNotEqualPackedInt32ArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualPackedInt64Array(right PackedInt64Array) bool {
  opPtr := ptrsForNil.operatorEqualPackedInt64ArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualPackedInt64Array(right PackedInt64Array) bool {
  opPtr := ptrsForNil.operatorNotEqualPackedInt64ArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualPackedFloat32Array(right PackedFloat32Array) bool {
  opPtr := ptrsForNil.operatorEqualPackedFloat32ArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualPackedFloat32Array(right PackedFloat32Array) bool {
  opPtr := ptrsForNil.operatorNotEqualPackedFloat32ArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualPackedFloat64Array(right PackedFloat64Array) bool {
  opPtr := ptrsForNil.operatorEqualPackedFloat64ArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualPackedFloat64Array(right PackedFloat64Array) bool {
  opPtr := ptrsForNil.operatorNotEqualPackedFloat64ArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualPackedStringArray(right PackedStringArray) bool {
  opPtr := ptrsForNil.operatorEqualPackedStringArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualPackedStringArray(right PackedStringArray) bool {
  opPtr := ptrsForNil.operatorNotEqualPackedStringArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualPackedVector2Array(right PackedVector2Array) bool {
  opPtr := ptrsForNil.operatorEqualPackedVector2ArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualPackedVector2Array(right PackedVector2Array) bool {
  opPtr := ptrsForNil.operatorNotEqualPackedVector2ArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualPackedVector3Array(right PackedVector3Array) bool {
  opPtr := ptrsForNil.operatorEqualPackedVector3ArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualPackedVector3Array(right PackedVector3Array) bool {
  opPtr := ptrsForNil.operatorNotEqualPackedVector3ArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) EqualPackedColorArray(right PackedColorArray) bool {
  opPtr := ptrsForNil.operatorEqualPackedColorArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Nil) NotEqualPackedColorArray(right PackedColorArray) bool {
  opPtr := ptrsForNil.operatorNotEqualPackedColorArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

// Members
