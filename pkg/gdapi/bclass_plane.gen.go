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

type ptrsForPlaneList struct {
	ctrFn                                 gdc.PtrConstructor
	ctrFromPlaneFn                        gdc.PtrConstructor
	ctrFromVector3Fn                      gdc.PtrConstructor
	ctrFromVector3Float32Fn               gdc.PtrConstructor
	ctrFromVector3Vector3Fn               gdc.PtrConstructor
	ctrFromVector3Vector3Vector3Fn        gdc.PtrConstructor
	ctrFromFloat32Float32Float32Float32Fn gdc.PtrConstructor
	methodNormalizedFn                    gdc.PtrBuiltInMethod
	methodGetCenterFn                     gdc.PtrBuiltInMethod
	methodIsEqualApproxFn                 gdc.PtrBuiltInMethod
	methodIsFiniteFn                      gdc.PtrBuiltInMethod
	methodIsPointOverFn                   gdc.PtrBuiltInMethod
	methodDistanceToFn                    gdc.PtrBuiltInMethod
	methodHasPointFn                      gdc.PtrBuiltInMethod
	methodProjectFn                       gdc.PtrBuiltInMethod
	methodIntersect3Fn                    gdc.PtrBuiltInMethod
	methodIntersectsRayFn                 gdc.PtrBuiltInMethod
	methodIntersectsSegmentFn             gdc.PtrBuiltInMethod
	operatorNegateFn                      gdc.PtrOperatorEvaluator
	operatorPositiveFn                    gdc.PtrOperatorEvaluator
	operatorNotFn                         gdc.PtrOperatorEvaluator
	operatorEqualPlaneFn                  gdc.PtrOperatorEvaluator
	operatorNotEqualPlaneFn               gdc.PtrOperatorEvaluator
	operatorMultiplyTransform3DFn         gdc.PtrOperatorEvaluator
	operatorInDictionaryFn                gdc.PtrOperatorEvaluator
	operatorInArrayFn                     gdc.PtrOperatorEvaluator
	memberxGetterFn                       gdc.PtrGetter
	memberxSetterFn                       gdc.PtrSetter
	memberyGetterFn                       gdc.PtrGetter
	memberySetterFn                       gdc.PtrSetter
	memberzGetterFn                       gdc.PtrGetter
	memberzSetterFn                       gdc.PtrSetter
	memberdGetterFn                       gdc.PtrGetter
	memberdSetterFn                       gdc.PtrSetter
	membernormalGetterFn                  gdc.PtrGetter
	membernormalSetterFn                  gdc.PtrSetter
	toVariantFn                           gdc.TypeFromVariantConstructorFunc
	fromVariantFn                         gdc.VariantFromTypeConstructorFunc
}

var ptrsForPlane ptrsForPlaneList

func initPlanePtrs(iface gdc.Interface) {
	ptrsForPlane.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePlane, 0))
	ptrsForPlane.ctrFromPlaneFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePlane, 1))
	ptrsForPlane.ctrFromVector3Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePlane, 2))
	ptrsForPlane.ctrFromVector3Float32Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePlane, 3))
	ptrsForPlane.ctrFromVector3Vector3Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePlane, 4))
	ptrsForPlane.ctrFromVector3Vector3Vector3Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePlane, 5))
	ptrsForPlane.ctrFromFloat32Float32Float32Float32Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePlane, 6))
	{
		methodName := StringNameFromStr("normalized")
		defer methodName.Destroy()
		ptrsForPlane.methodNormalizedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, methodName.AsCPtr(), 1051796340))
	}
	{
		methodName := StringNameFromStr("get_center")
		defer methodName.Destroy()
		ptrsForPlane.methodGetCenterFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, methodName.AsCPtr(), 1776574132))
	}
	{
		methodName := StringNameFromStr("is_equal_approx")
		defer methodName.Destroy()
		ptrsForPlane.methodIsEqualApproxFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, methodName.AsCPtr(), 1150170233))
	}
	{
		methodName := StringNameFromStr("is_finite")
		defer methodName.Destroy()
		ptrsForPlane.methodIsFiniteFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("is_point_over")
		defer methodName.Destroy()
		ptrsForPlane.methodIsPointOverFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, methodName.AsCPtr(), 1749054343))
	}
	{
		methodName := StringNameFromStr("distance_to")
		defer methodName.Destroy()
		ptrsForPlane.methodDistanceToFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, methodName.AsCPtr(), 1047977935))
	}
	{
		methodName := StringNameFromStr("has_point")
		defer methodName.Destroy()
		ptrsForPlane.methodHasPointFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, methodName.AsCPtr(), 1258189072))
	}
	{
		methodName := StringNameFromStr("project")
		defer methodName.Destroy()
		ptrsForPlane.methodProjectFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, methodName.AsCPtr(), 2923479887))
	}
	{
		methodName := StringNameFromStr("intersect_3")
		defer methodName.Destroy()
		ptrsForPlane.methodIntersect3Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, methodName.AsCPtr(), 2012052692))
	}
	{
		methodName := StringNameFromStr("intersects_ray")
		defer methodName.Destroy()
		ptrsForPlane.methodIntersectsRayFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, methodName.AsCPtr(), 2048133369))
	}
	{
		methodName := StringNameFromStr("intersects_segment")
		defer methodName.Destroy()
		ptrsForPlane.methodIntersectsSegmentFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePlane, methodName.AsCPtr(), 2048133369))
	}
	ptrsForPlane.operatorNegateFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNegate, gdc.VariantTypePlane, gdc.VariantTypeNil))
	ptrsForPlane.operatorPositiveFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpPositive, gdc.VariantTypePlane, gdc.VariantTypeNil))
	ptrsForPlane.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypePlane, gdc.VariantTypeNil))
	ptrsForPlane.operatorEqualPlaneFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypePlane, gdc.VariantTypePlane))
	ptrsForPlane.operatorNotEqualPlaneFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypePlane, gdc.VariantTypePlane))
	ptrsForPlane.operatorMultiplyTransform3DFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypePlane, gdc.VariantTypeTransform3D))
	ptrsForPlane.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypePlane, gdc.VariantTypeDictionary))
	ptrsForPlane.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypePlane, gdc.VariantTypeArray))
	{
		memberName := StringNameFromStr("x")
		defer memberName.Destroy()
		ptrsForPlane.memberxGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypePlane, memberName.AsCPtr()))
		ptrsForPlane.memberxSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypePlane, memberName.AsCPtr()))
	}
	{
		memberName := StringNameFromStr("y")
		defer memberName.Destroy()
		ptrsForPlane.memberyGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypePlane, memberName.AsCPtr()))
		ptrsForPlane.memberySetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypePlane, memberName.AsCPtr()))
	}
	{
		memberName := StringNameFromStr("z")
		defer memberName.Destroy()
		ptrsForPlane.memberzGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypePlane, memberName.AsCPtr()))
		ptrsForPlane.memberzSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypePlane, memberName.AsCPtr()))
	}
	{
		memberName := StringNameFromStr("d")
		defer memberName.Destroy()
		ptrsForPlane.memberdGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypePlane, memberName.AsCPtr()))
		ptrsForPlane.memberdSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypePlane, memberName.AsCPtr()))
	}
	{
		memberName := StringNameFromStr("normal")
		defer memberName.Destroy()
		ptrsForPlane.membernormalGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypePlane, memberName.AsCPtr()))
		ptrsForPlane.membernormalSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypePlane, memberName.AsCPtr()))
	}
	ptrsForPlane.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypePlane))
	ptrsForPlane.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypePlane))
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("PLANE_YZ")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypePlane, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsPlane()
		if err != nil {
			panic("Failed to get constant value PLANE_YZ: " + err.Error())
		}
		PlanePlaneYz = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("PLANE_XZ")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypePlane, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsPlane()
		if err != nil {
			panic("Failed to get constant value PLANE_XZ: " + err.Error())
		}
		PlanePlaneXz = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("PLANE_XY")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypePlane, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsPlane()
		if err != nil {
			panic("Failed to get constant value PLANE_XY: " + err.Error())
		}
		PlanePlaneXy = *cnst
	}
}

type Plane struct {
	data   *[classSizePlane]byte
	iface  gdc.Interface
	pinner runtime.Pinner
}

// Constants

var (
	PlanePlaneYz Plane
	PlanePlaneXz Plane
	PlanePlaneXy Plane
)

// Enums

// Constructors
func newPlane() *Plane {
	me := &Plane{
		data:  new([classSizePlane]byte),
		iface: giface,
	}
	me.pinner.Pin(me)
	me.pinner.Pin(me.AsTypePtr())
	return me
}

func NewPlane() *Plane {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newPlane()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForPlane.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
	return me
}

func NewPlaneFromPlane(from Plane) *Plane {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newPlane()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForPlane.ctrFromPlaneFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewPlaneFromVector3(normal Vector3) *Plane {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newPlane()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForPlane.ctrFromVector3Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{normal.AsCTypePtr()}))
	return me
}

func NewPlaneFromVector3Float32(normal Vector3, d float64) *Plane {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	pinner.Pin(&d)
	me := newPlane()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForPlane.ctrFromVector3Float32Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{normal.AsCTypePtr(), gdc.ConstTypePtr(&d)}))
	return me
}

func NewPlaneFromVector3Vector3(normal Vector3, point Vector3) *Plane {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newPlane()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForPlane.ctrFromVector3Vector3Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{normal.AsCTypePtr(), point.AsCTypePtr()}))
	return me
}

func NewPlaneFromVector3Vector3Vector3(point1 Vector3, point2 Vector3, point3 Vector3) *Plane {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newPlane()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForPlane.ctrFromVector3Vector3Vector3Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{point1.AsCTypePtr(), point2.AsCTypePtr(), point3.AsCTypePtr()}))
	return me
}

func NewPlaneFromFloat32Float32Float32Float32(a float64, b float64, c float64, d float64) *Plane {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	pinner.Pin(&a)
	pinner.Pin(&b)
	pinner.Pin(&c)
	pinner.Pin(&d)
	me := newPlane()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForPlane.ctrFromFloat32Float32Float32Float32Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&a), gdc.ConstTypePtr(&b), gdc.ConstTypePtr(&c), gdc.ConstTypePtr(&d)}))
	return me
}

// Destructor
func (me *Plane) Destroy() {
	me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsPlane() (*Plane, error) {
	if me.Type() != gdc.VariantTypePlane {
		return nil, fmt.Errorf("variant is not a Plane")
	}
	bclass := newPlane()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForPlane.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Plane) AsVariant() *Variant {
	va := newVariant()
	va.inner = me
	me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForPlane.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
	return va
}

// Pointers
func PlaneFromPtr(ptr gdc.ConstTypePtr) *Plane {
	me := newPlane()
	dataFromPtr(me.data[:], ptr)
	return me
}

func (me *Plane) ToTypePtr(ptr gdc.TypePtr) {
	dataToPtr(ptr, me.data[:])
}

func (me *Plane) Type() gdc.VariantType {
	return gdc.VariantTypePlane
}

func (me *Plane) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Plane) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Plane) asUninitialized() gdc.UninitializedTypePtr {
	return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *Plane) Normalized() Plane {
	ret := NewPlane()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPlane.methodNormalizedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Plane) GetCenter() Vector3 {
	ret := NewVector3()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPlane.methodGetCenterFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Plane) IsEqualApprox(to_plane Plane) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{to_plane.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPlane.methodIsEqualApproxFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Plane) IsFinite() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPlane.methodIsFiniteFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Plane) IsPointOver(point Vector3) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{point.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPlane.methodIsPointOverFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Plane) DistanceTo(point Vector3) float64 {
	ret := NewFloat()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{point.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPlane.methodDistanceToFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Plane) HasPoint(point Vector3, tolerance float64) bool {
	ret := NewBool()
	defer ret.Destroy()

	varg1 := NewFloatFromFloat32(tolerance)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{point.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPlane.methodHasPointFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Plane) Project(point Vector3) Vector3 {
	ret := NewVector3()

	args := []gdc.ConstTypePtr{point.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPlane.methodProjectFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Plane) Intersect3(b Plane, c Plane) Variant {
	ret := NewVariant()

	args := []gdc.ConstTypePtr{b.AsCTypePtr(), c.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPlane.methodIntersect3Fn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Plane) IntersectsRay(from Vector3, dir Vector3) Variant {
	ret := NewVariant()

	args := []gdc.ConstTypePtr{from.AsCTypePtr(), dir.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPlane.methodIntersectsRayFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Plane) IntersectsSegment(from Vector3, to Vector3) Variant {
	ret := NewVariant()

	args := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPlane.methodIntersectsSegmentFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

// Operators

func (me *Plane) EqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Plane) NotEqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Plane) Negate() Plane {
	opPtr := ptrsForPlane.operatorNegateFn
	ret := NewPlane()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return *ret
}

func (me *Plane) Positive() Plane {
	opPtr := ptrsForPlane.operatorPositiveFn
	ret := NewPlane()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return *ret
}

func (me *Plane) Not() bool {
	opPtr := ptrsForPlane.operatorNotFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return ret.Get()
}

func (me *Plane) EqualPlane(right Plane) bool {
	opPtr := ptrsForPlane.operatorEqualPlaneFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Plane) NotEqualPlane(right Plane) bool {
	opPtr := ptrsForPlane.operatorNotEqualPlaneFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Plane) MultiplyTransform3D(right Transform3D) Plane {
	opPtr := ptrsForPlane.operatorMultiplyTransform3DFn
	ret := NewPlane()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Plane) InDictionary(right Dictionary) bool {
	opPtr := ptrsForPlane.operatorInDictionaryFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Plane) InArray(right Array) bool {
	opPtr := ptrsForPlane.operatorInArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

// Members

func (me *Plane) X() float64 {
	ret := NewFloat()
	me.iface.CallPtrGetter(ensurePtr(ptrsForPlane.memberxGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Plane) SetX(value float64) {
	valueV := NewFloatFromFloat32(value)
	defer valueV.Destroy()
	me.iface.CallPtrSetter(ensurePtr(ptrsForPlane.memberxSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Plane) Y() float64 {
	ret := NewFloat()
	me.iface.CallPtrGetter(ensurePtr(ptrsForPlane.memberyGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Plane) SetY(value float64) {
	valueV := NewFloatFromFloat32(value)
	defer valueV.Destroy()
	me.iface.CallPtrSetter(ensurePtr(ptrsForPlane.memberySetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Plane) Z() float64 {
	ret := NewFloat()
	me.iface.CallPtrGetter(ensurePtr(ptrsForPlane.memberzGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Plane) SetZ(value float64) {
	valueV := NewFloatFromFloat32(value)
	defer valueV.Destroy()
	me.iface.CallPtrSetter(ensurePtr(ptrsForPlane.memberzSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Plane) D() float64 {
	ret := NewFloat()
	me.iface.CallPtrGetter(ensurePtr(ptrsForPlane.memberdGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Plane) SetD(value float64) {
	valueV := NewFloatFromFloat32(value)
	defer valueV.Destroy()
	me.iface.CallPtrSetter(ensurePtr(ptrsForPlane.memberdSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Plane) Normal() Vector3 {
	ret := NewVector3()
	me.iface.CallPtrGetter(ensurePtr(ptrsForPlane.membernormalGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Plane) SetNormal(value Vector3) {
	valueV := value
	me.iface.CallPtrSetter(ensurePtr(ptrsForPlane.membernormalSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}
