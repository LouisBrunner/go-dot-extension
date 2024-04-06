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

type ptrsForAABBList struct {
	ctrFn                         gdc.PtrConstructor
	ctrFromAABBFn                 gdc.PtrConstructor
	ctrFromVector3Vector3Fn       gdc.PtrConstructor
	methodAbsFn                   gdc.PtrBuiltInMethod
	methodGetCenterFn             gdc.PtrBuiltInMethod
	methodGetVolumeFn             gdc.PtrBuiltInMethod
	methodHasVolumeFn             gdc.PtrBuiltInMethod
	methodHasSurfaceFn            gdc.PtrBuiltInMethod
	methodHasPointFn              gdc.PtrBuiltInMethod
	methodIsEqualApproxFn         gdc.PtrBuiltInMethod
	methodIsFiniteFn              gdc.PtrBuiltInMethod
	methodIntersectsFn            gdc.PtrBuiltInMethod
	methodEnclosesFn              gdc.PtrBuiltInMethod
	methodIntersectsPlaneFn       gdc.PtrBuiltInMethod
	methodIntersectionFn          gdc.PtrBuiltInMethod
	methodMergeFn                 gdc.PtrBuiltInMethod
	methodExpandFn                gdc.PtrBuiltInMethod
	methodGrowFn                  gdc.PtrBuiltInMethod
	methodGetSupportFn            gdc.PtrBuiltInMethod
	methodGetLongestAxisFn        gdc.PtrBuiltInMethod
	methodGetLongestAxisIndexFn   gdc.PtrBuiltInMethod
	methodGetLongestAxisSizeFn    gdc.PtrBuiltInMethod
	methodGetShortestAxisFn       gdc.PtrBuiltInMethod
	methodGetShortestAxisIndexFn  gdc.PtrBuiltInMethod
	methodGetShortestAxisSizeFn   gdc.PtrBuiltInMethod
	methodGetEndpointFn           gdc.PtrBuiltInMethod
	methodIntersectsSegmentFn     gdc.PtrBuiltInMethod
	methodIntersectsRayFn         gdc.PtrBuiltInMethod
	operatorNotFn                 gdc.PtrOperatorEvaluator
	operatorEqualAABBFn           gdc.PtrOperatorEvaluator
	operatorNotEqualAABBFn        gdc.PtrOperatorEvaluator
	operatorMultiplyTransform3DFn gdc.PtrOperatorEvaluator
	operatorInDictionaryFn        gdc.PtrOperatorEvaluator
	operatorInArrayFn             gdc.PtrOperatorEvaluator
	memberpositionGetterFn        gdc.PtrGetter
	memberpositionSetterFn        gdc.PtrSetter
	membersizeGetterFn            gdc.PtrGetter
	membersizeSetterFn            gdc.PtrSetter
	memberendGetterFn             gdc.PtrGetter
	memberendSetterFn             gdc.PtrSetter
	toVariantFn                   gdc.TypeFromVariantConstructorFunc
	fromVariantFn                 gdc.VariantFromTypeConstructorFunc
}

var ptrsForAABB ptrsForAABBList

func initAABBPtrs(iface gdc.Interface) {
	ptrsForAABB.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeAABB, 0))
	ptrsForAABB.ctrFromAABBFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeAABB, 1))
	ptrsForAABB.ctrFromVector3Vector3Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeAABB, 2))
	{
		methodName := StringNameFromStr("abs")
		defer methodName.Destroy()
		ptrsForAABB.methodAbsFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, methodName.AsCPtr(), 1576868580))
	}
	{
		methodName := StringNameFromStr("get_center")
		defer methodName.Destroy()
		ptrsForAABB.methodGetCenterFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, methodName.AsCPtr(), 1776574132))
	}
	{
		methodName := StringNameFromStr("get_volume")
		defer methodName.Destroy()
		ptrsForAABB.methodGetVolumeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, methodName.AsCPtr(), 466405837))
	}
	{
		methodName := StringNameFromStr("has_volume")
		defer methodName.Destroy()
		ptrsForAABB.methodHasVolumeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("has_surface")
		defer methodName.Destroy()
		ptrsForAABB.methodHasSurfaceFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("has_point")
		defer methodName.Destroy()
		ptrsForAABB.methodHasPointFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, methodName.AsCPtr(), 1749054343))
	}
	{
		methodName := StringNameFromStr("is_equal_approx")
		defer methodName.Destroy()
		ptrsForAABB.methodIsEqualApproxFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, methodName.AsCPtr(), 299946684))
	}
	{
		methodName := StringNameFromStr("is_finite")
		defer methodName.Destroy()
		ptrsForAABB.methodIsFiniteFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("intersects")
		defer methodName.Destroy()
		ptrsForAABB.methodIntersectsFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, methodName.AsCPtr(), 299946684))
	}
	{
		methodName := StringNameFromStr("encloses")
		defer methodName.Destroy()
		ptrsForAABB.methodEnclosesFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, methodName.AsCPtr(), 299946684))
	}
	{
		methodName := StringNameFromStr("intersects_plane")
		defer methodName.Destroy()
		ptrsForAABB.methodIntersectsPlaneFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, methodName.AsCPtr(), 1150170233))
	}
	{
		methodName := StringNameFromStr("intersection")
		defer methodName.Destroy()
		ptrsForAABB.methodIntersectionFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, methodName.AsCPtr(), 1271470306))
	}
	{
		methodName := StringNameFromStr("merge")
		defer methodName.Destroy()
		ptrsForAABB.methodMergeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, methodName.AsCPtr(), 1271470306))
	}
	{
		methodName := StringNameFromStr("expand")
		defer methodName.Destroy()
		ptrsForAABB.methodExpandFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, methodName.AsCPtr(), 2851643018))
	}
	{
		methodName := StringNameFromStr("grow")
		defer methodName.Destroy()
		ptrsForAABB.methodGrowFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, methodName.AsCPtr(), 239217291))
	}
	{
		methodName := StringNameFromStr("get_support")
		defer methodName.Destroy()
		ptrsForAABB.methodGetSupportFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, methodName.AsCPtr(), 2923479887))
	}
	{
		methodName := StringNameFromStr("get_longest_axis")
		defer methodName.Destroy()
		ptrsForAABB.methodGetLongestAxisFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, methodName.AsCPtr(), 1776574132))
	}
	{
		methodName := StringNameFromStr("get_longest_axis_index")
		defer methodName.Destroy()
		ptrsForAABB.methodGetLongestAxisIndexFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("get_longest_axis_size")
		defer methodName.Destroy()
		ptrsForAABB.methodGetLongestAxisSizeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, methodName.AsCPtr(), 466405837))
	}
	{
		methodName := StringNameFromStr("get_shortest_axis")
		defer methodName.Destroy()
		ptrsForAABB.methodGetShortestAxisFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, methodName.AsCPtr(), 1776574132))
	}
	{
		methodName := StringNameFromStr("get_shortest_axis_index")
		defer methodName.Destroy()
		ptrsForAABB.methodGetShortestAxisIndexFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("get_shortest_axis_size")
		defer methodName.Destroy()
		ptrsForAABB.methodGetShortestAxisSizeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, methodName.AsCPtr(), 466405837))
	}
	{
		methodName := StringNameFromStr("get_endpoint")
		defer methodName.Destroy()
		ptrsForAABB.methodGetEndpointFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, methodName.AsCPtr(), 1394941017))
	}
	{
		methodName := StringNameFromStr("intersects_segment")
		defer methodName.Destroy()
		ptrsForAABB.methodIntersectsSegmentFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, methodName.AsCPtr(), 2048133369))
	}
	{
		methodName := StringNameFromStr("intersects_ray")
		defer methodName.Destroy()
		ptrsForAABB.methodIntersectsRayFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeAABB, methodName.AsCPtr(), 2048133369))
	}
	ptrsForAABB.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypeAABB, gdc.VariantTypeNil))
	ptrsForAABB.operatorEqualAABBFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeAABB, gdc.VariantTypeAABB))
	ptrsForAABB.operatorNotEqualAABBFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeAABB, gdc.VariantTypeAABB))
	ptrsForAABB.operatorMultiplyTransform3DFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeAABB, gdc.VariantTypeTransform3D))
	ptrsForAABB.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeAABB, gdc.VariantTypeDictionary))
	ptrsForAABB.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeAABB, gdc.VariantTypeArray))
	{
		memberName := StringNameFromStr("position")
		defer memberName.Destroy()
		ptrsForAABB.memberpositionGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeAABB, memberName.AsCPtr()))
		ptrsForAABB.memberpositionSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeAABB, memberName.AsCPtr()))
	}
	{
		memberName := StringNameFromStr("size")
		defer memberName.Destroy()
		ptrsForAABB.membersizeGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeAABB, memberName.AsCPtr()))
		ptrsForAABB.membersizeSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeAABB, memberName.AsCPtr()))
	}
	{
		memberName := StringNameFromStr("end")
		defer memberName.Destroy()
		ptrsForAABB.memberendGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeAABB, memberName.AsCPtr()))
		ptrsForAABB.memberendSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeAABB, memberName.AsCPtr()))
	}
	ptrsForAABB.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypeAABB))
	ptrsForAABB.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypeAABB))
}

type AABB struct {
	data   *[classSizeAABB]byte
	iface  gdc.Interface
	pinner runtime.Pinner
}

// Enums

// Constructors
func newAABB() *AABB {
	me := &AABB{
		data:  new([classSizeAABB]byte),
		iface: giface,
	}
	me.pinner.Pin(me)
	me.pinner.Pin(me.AsTypePtr())
	return me
}

func NewAABB() *AABB {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newAABB()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForAABB.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
	return me
}

func NewAABBFromAABB(from AABB) *AABB {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newAABB()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForAABB.ctrFromAABBFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewAABBFromVector3Vector3(position Vector3, size Vector3) *AABB {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newAABB()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForAABB.ctrFromVector3Vector3Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{position.AsCTypePtr(), size.AsCTypePtr()}))
	return me
}

// Destructor
func (me *AABB) Destroy() {
	me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsAABB() (*AABB, error) {
	if me.Type() != gdc.VariantTypeAABB {
		return nil, fmt.Errorf("variant is not a AABB")
	}
	bclass := newAABB()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForAABB.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *AABB) AsVariant() *Variant {
	va := newVariant()
	va.inner = me
	me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForAABB.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
	return va
}

// Pointers
func AABBFromPtr(ptr gdc.ConstTypePtr) *AABB {
	me := newAABB()
	dataFromPtr(me.data[:], ptr)
	return me
}

func (me *AABB) Type() gdc.VariantType {
	return gdc.VariantTypeAABB
}

func (me *AABB) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *AABB) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *AABB) asUninitialized() gdc.UninitializedTypePtr {
	return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *AABB) Abs() AABB {
	ret := NewAABB()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForAABB.methodAbsFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *AABB) GetCenter() Vector3 {
	ret := NewVector3()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForAABB.methodGetCenterFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *AABB) GetVolume() float64 {
	ret := NewFloat()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForAABB.methodGetVolumeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *AABB) HasVolume() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForAABB.methodHasVolumeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *AABB) HasSurface() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForAABB.methodHasSurfaceFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *AABB) HasPoint(point Vector3) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{point.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForAABB.methodHasPointFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *AABB) IsEqualApprox(aabb AABB) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{aabb.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForAABB.methodIsEqualApproxFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *AABB) IsFinite() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForAABB.methodIsFiniteFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *AABB) Intersects(with AABB) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{with.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForAABB.methodIntersectsFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *AABB) Encloses(with AABB) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{with.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForAABB.methodEnclosesFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *AABB) IntersectsPlane(plane Plane) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{plane.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForAABB.methodIntersectsPlaneFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *AABB) Intersection(with AABB) AABB {
	ret := NewAABB()

	args := []gdc.ConstTypePtr{with.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForAABB.methodIntersectionFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *AABB) Merge(with AABB) AABB {
	ret := NewAABB()

	args := []gdc.ConstTypePtr{with.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForAABB.methodMergeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *AABB) Expand(to_point Vector3) AABB {
	ret := NewAABB()

	args := []gdc.ConstTypePtr{to_point.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForAABB.methodExpandFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *AABB) Grow(by float64) AABB {
	ret := NewAABB()

	varg0 := NewFloatFromFloat32(by)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForAABB.methodGrowFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *AABB) GetSupport(dir Vector3) Vector3 {
	ret := NewVector3()

	args := []gdc.ConstTypePtr{dir.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForAABB.methodGetSupportFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *AABB) GetLongestAxis() Vector3 {
	ret := NewVector3()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForAABB.methodGetLongestAxisFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *AABB) GetLongestAxisIndex() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForAABB.methodGetLongestAxisIndexFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *AABB) GetLongestAxisSize() float64 {
	ret := NewFloat()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForAABB.methodGetLongestAxisSizeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *AABB) GetShortestAxis() Vector3 {
	ret := NewVector3()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForAABB.methodGetShortestAxisFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *AABB) GetShortestAxisIndex() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForAABB.methodGetShortestAxisIndexFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *AABB) GetShortestAxisSize() float64 {
	ret := NewFloat()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForAABB.methodGetShortestAxisSizeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *AABB) GetEndpoint(idx int64) Vector3 {
	ret := NewVector3()

	varg0 := NewIntFromInt(idx)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForAABB.methodGetEndpointFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *AABB) IntersectsSegment(from Vector3, to Vector3) Variant {
	ret := NewVariant()

	args := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForAABB.methodIntersectsSegmentFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *AABB) IntersectsRay(from Vector3, dir Vector3) Variant {
	ret := NewVariant()

	args := []gdc.ConstTypePtr{from.AsCTypePtr(), dir.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForAABB.methodIntersectsRayFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

// Operators

func (me *AABB) EqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *AABB) NotEqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *AABB) Not() bool {
	opPtr := ptrsForAABB.operatorNotFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return ret.Get()
}

func (me *AABB) EqualAABB(right AABB) bool {
	opPtr := ptrsForAABB.operatorEqualAABBFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *AABB) NotEqualAABB(right AABB) bool {
	opPtr := ptrsForAABB.operatorNotEqualAABBFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *AABB) MultiplyTransform3D(right Transform3D) AABB {
	opPtr := ptrsForAABB.operatorMultiplyTransform3DFn
	ret := NewAABB()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *AABB) InDictionary(right Dictionary) bool {
	opPtr := ptrsForAABB.operatorInDictionaryFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *AABB) InArray(right Array) bool {
	opPtr := ptrsForAABB.operatorInArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

// Members

func (me *AABB) Position() Vector3 {
	ret := NewVector3()
	me.iface.CallPtrGetter(ensurePtr(ptrsForAABB.memberpositionGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *AABB) SetPosition(value Vector3) {
	valueV := value
	me.iface.CallPtrSetter(ensurePtr(ptrsForAABB.memberpositionSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *AABB) Size() Vector3 {
	ret := NewVector3()
	me.iface.CallPtrGetter(ensurePtr(ptrsForAABB.membersizeGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *AABB) SetSize(value Vector3) {
	valueV := value
	me.iface.CallPtrSetter(ensurePtr(ptrsForAABB.membersizeSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *AABB) End() Vector3 {
	ret := NewVector3()
	me.iface.CallPtrGetter(ensurePtr(ptrsForAABB.memberendGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *AABB) SetEnd(value Vector3) {
	valueV := value
	me.iface.CallPtrSetter(ensurePtr(ptrsForAABB.memberendSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}
