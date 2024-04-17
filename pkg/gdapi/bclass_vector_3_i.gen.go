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

type ptrsForVector3iList struct {
	ctrFn                          gdc.PtrConstructor
	ctrFromVector3IFn              gdc.PtrConstructor
	ctrFromVector3Fn               gdc.PtrConstructor
	ctrFromIntIntIntFn             gdc.PtrConstructor
	methodMinAxisIndexFn           gdc.PtrBuiltInMethod
	methodMaxAxisIndexFn           gdc.PtrBuiltInMethod
	methodDistanceToFn             gdc.PtrBuiltInMethod
	methodDistanceSquaredToFn      gdc.PtrBuiltInMethod
	methodLengthFn                 gdc.PtrBuiltInMethod
	methodLengthSquaredFn          gdc.PtrBuiltInMethod
	methodSignFn                   gdc.PtrBuiltInMethod
	methodAbsFn                    gdc.PtrBuiltInMethod
	methodClampFn                  gdc.PtrBuiltInMethod
	methodSnappedFn                gdc.PtrBuiltInMethod
	operatorNegateFn               gdc.PtrOperatorEvaluator
	operatorPositiveFn             gdc.PtrOperatorEvaluator
	operatorNotFn                  gdc.PtrOperatorEvaluator
	operatorMultiplyIntFn          gdc.PtrOperatorEvaluator
	operatorDivideIntFn            gdc.PtrOperatorEvaluator
	operatorModuleIntFn            gdc.PtrOperatorEvaluator
	operatorMultiplyFloat32Fn      gdc.PtrOperatorEvaluator
	operatorDivideFloat32Fn        gdc.PtrOperatorEvaluator
	operatorEqualVector3IFn        gdc.PtrOperatorEvaluator
	operatorNotEqualVector3IFn     gdc.PtrOperatorEvaluator
	operatorLessVector3IFn         gdc.PtrOperatorEvaluator
	operatorLessEqualVector3IFn    gdc.PtrOperatorEvaluator
	operatorGreaterVector3IFn      gdc.PtrOperatorEvaluator
	operatorGreaterEqualVector3IFn gdc.PtrOperatorEvaluator
	operatorAddVector3IFn          gdc.PtrOperatorEvaluator
	operatorSubtractVector3IFn     gdc.PtrOperatorEvaluator
	operatorMultiplyVector3IFn     gdc.PtrOperatorEvaluator
	operatorDivideVector3IFn       gdc.PtrOperatorEvaluator
	operatorModuleVector3IFn       gdc.PtrOperatorEvaluator
	operatorInDictionaryFn         gdc.PtrOperatorEvaluator
	operatorInArrayFn              gdc.PtrOperatorEvaluator
	memberxGetterFn                gdc.PtrGetter
	memberxSetterFn                gdc.PtrSetter
	memberyGetterFn                gdc.PtrGetter
	memberySetterFn                gdc.PtrSetter
	memberzGetterFn                gdc.PtrGetter
	memberzSetterFn                gdc.PtrSetter
	toVariantFn                    gdc.TypeFromVariantConstructorFunc
	fromVariantFn                  gdc.VariantFromTypeConstructorFunc
}

var ptrsForVector3i ptrsForVector3iList

func initVector3iPtrs(iface gdc.Interface) {
	ptrsForVector3i.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeVector3I, 0))
	ptrsForVector3i.ctrFromVector3IFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeVector3I, 1))
	ptrsForVector3i.ctrFromVector3Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeVector3I, 2))
	ptrsForVector3i.ctrFromIntIntIntFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeVector3I, 3))
	{
		methodName := StringNameFromStr("min_axis_index")
		defer methodName.Destroy()
		ptrsForVector3i.methodMinAxisIndexFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3I, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("max_axis_index")
		defer methodName.Destroy()
		ptrsForVector3i.methodMaxAxisIndexFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3I, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("distance_to")
		defer methodName.Destroy()
		ptrsForVector3i.methodDistanceToFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3I, methodName.AsCPtr(), 1975170430))
	}
	{
		methodName := StringNameFromStr("distance_squared_to")
		defer methodName.Destroy()
		ptrsForVector3i.methodDistanceSquaredToFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3I, methodName.AsCPtr(), 2947717320))
	}
	{
		methodName := StringNameFromStr("length")
		defer methodName.Destroy()
		ptrsForVector3i.methodLengthFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3I, methodName.AsCPtr(), 466405837))
	}
	{
		methodName := StringNameFromStr("length_squared")
		defer methodName.Destroy()
		ptrsForVector3i.methodLengthSquaredFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3I, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("sign")
		defer methodName.Destroy()
		ptrsForVector3i.methodSignFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3I, methodName.AsCPtr(), 3729604559))
	}
	{
		methodName := StringNameFromStr("abs")
		defer methodName.Destroy()
		ptrsForVector3i.methodAbsFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3I, methodName.AsCPtr(), 3729604559))
	}
	{
		methodName := StringNameFromStr("clamp")
		defer methodName.Destroy()
		ptrsForVector3i.methodClampFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3I, methodName.AsCPtr(), 1086892323))
	}
	{
		methodName := StringNameFromStr("snapped")
		defer methodName.Destroy()
		ptrsForVector3i.methodSnappedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector3I, methodName.AsCPtr(), 1989319750))
	}
	ptrsForVector3i.operatorNegateFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNegate, gdc.VariantTypeVector3I, gdc.VariantTypeNil))
	ptrsForVector3i.operatorPositiveFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpPositive, gdc.VariantTypeVector3I, gdc.VariantTypeNil))
	ptrsForVector3i.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypeVector3I, gdc.VariantTypeNil))
	ptrsForVector3i.operatorMultiplyIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeVector3I, gdc.VariantTypeInt))
	ptrsForVector3i.operatorDivideIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeVector3I, gdc.VariantTypeInt))
	ptrsForVector3i.operatorModuleIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeVector3I, gdc.VariantTypeInt))
	ptrsForVector3i.operatorMultiplyFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeVector3I, gdc.VariantTypeFloat))
	ptrsForVector3i.operatorDivideFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeVector3I, gdc.VariantTypeFloat))
	ptrsForVector3i.operatorEqualVector3IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeVector3I, gdc.VariantTypeVector3I))
	ptrsForVector3i.operatorNotEqualVector3IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeVector3I, gdc.VariantTypeVector3I))
	ptrsForVector3i.operatorLessVector3IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLess, gdc.VariantTypeVector3I, gdc.VariantTypeVector3I))
	ptrsForVector3i.operatorLessEqualVector3IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLessEqual, gdc.VariantTypeVector3I, gdc.VariantTypeVector3I))
	ptrsForVector3i.operatorGreaterVector3IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreater, gdc.VariantTypeVector3I, gdc.VariantTypeVector3I))
	ptrsForVector3i.operatorGreaterEqualVector3IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreaterEqual, gdc.VariantTypeVector3I, gdc.VariantTypeVector3I))
	ptrsForVector3i.operatorAddVector3IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, gdc.VariantTypeVector3I, gdc.VariantTypeVector3I))
	ptrsForVector3i.operatorSubtractVector3IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpSubtract, gdc.VariantTypeVector3I, gdc.VariantTypeVector3I))
	ptrsForVector3i.operatorMultiplyVector3IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeVector3I, gdc.VariantTypeVector3I))
	ptrsForVector3i.operatorDivideVector3IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeVector3I, gdc.VariantTypeVector3I))
	ptrsForVector3i.operatorModuleVector3IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeVector3I, gdc.VariantTypeVector3I))
	ptrsForVector3i.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeVector3I, gdc.VariantTypeDictionary))
	ptrsForVector3i.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeVector3I, gdc.VariantTypeArray))
	{
		memberName := StringNameFromStr("x")
		defer memberName.Destroy()
		ptrsForVector3i.memberxGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeVector3I, memberName.AsCPtr()))
		ptrsForVector3i.memberxSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeVector3I, memberName.AsCPtr()))
	}
	{
		memberName := StringNameFromStr("y")
		defer memberName.Destroy()
		ptrsForVector3i.memberyGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeVector3I, memberName.AsCPtr()))
		ptrsForVector3i.memberySetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeVector3I, memberName.AsCPtr()))
	}
	{
		memberName := StringNameFromStr("z")
		defer memberName.Destroy()
		ptrsForVector3i.memberzGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeVector3I, memberName.AsCPtr()))
		ptrsForVector3i.memberzSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeVector3I, memberName.AsCPtr()))
	}
	ptrsForVector3i.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypeVector3I))
	ptrsForVector3i.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypeVector3I))
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("ZERO")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector3I, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector3i()
		if err != nil {
			panic("Failed to get constant value ZERO: " + err.Error())
		}
		Vector3iZero = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("ONE")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector3I, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector3i()
		if err != nil {
			panic("Failed to get constant value ONE: " + err.Error())
		}
		Vector3iOne = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("MIN")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector3I, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector3i()
		if err != nil {
			panic("Failed to get constant value MIN: " + err.Error())
		}
		Vector3iMin = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("MAX")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector3I, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector3i()
		if err != nil {
			panic("Failed to get constant value MAX: " + err.Error())
		}
		Vector3iMax = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("LEFT")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector3I, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector3i()
		if err != nil {
			panic("Failed to get constant value LEFT: " + err.Error())
		}
		Vector3iLeft = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("RIGHT")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector3I, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector3i()
		if err != nil {
			panic("Failed to get constant value RIGHT: " + err.Error())
		}
		Vector3iRight = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("UP")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector3I, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector3i()
		if err != nil {
			panic("Failed to get constant value UP: " + err.Error())
		}
		Vector3iUp = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("DOWN")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector3I, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector3i()
		if err != nil {
			panic("Failed to get constant value DOWN: " + err.Error())
		}
		Vector3iDown = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("FORWARD")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector3I, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector3i()
		if err != nil {
			panic("Failed to get constant value FORWARD: " + err.Error())
		}
		Vector3iForward = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("BACK")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector3I, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector3i()
		if err != nil {
			panic("Failed to get constant value BACK: " + err.Error())
		}
		Vector3iBack = *cnst
	}
}

type Vector3i struct {
	data   *[classSizeVector3i]byte
	iface  gdc.Interface
	pinner runtime.Pinner
}

// Constants

var (
	Vector3iZero    Vector3i
	Vector3iOne     Vector3i
	Vector3iMin     Vector3i
	Vector3iMax     Vector3i
	Vector3iLeft    Vector3i
	Vector3iRight   Vector3i
	Vector3iUp      Vector3i
	Vector3iDown    Vector3i
	Vector3iForward Vector3i
	Vector3iBack    Vector3i
)

// Enums

type Vector3iAxis int

const (
	Vector3iAxisAxisX Vector3iAxis = 0
	Vector3iAxisAxisY Vector3iAxis = 1
	Vector3iAxisAxisZ Vector3iAxis = 2
)

// Constructors
func newVector3i() *Vector3i {
	me := &Vector3i{
		data:  new([classSizeVector3i]byte),
		iface: giface,
	}
	me.pinner.Pin(me)
	me.pinner.Pin(me.AsTypePtr())
	return me
}

func NewVector3i() *Vector3i {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newVector3i()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForVector3i.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
	return me
}

func NewVector3iFromVector3I(from Vector3i) *Vector3i {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newVector3i()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForVector3i.ctrFromVector3IFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewVector3iFromVector3(from Vector3) *Vector3i {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newVector3i()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForVector3i.ctrFromVector3Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewVector3iFromIntIntInt(x int64, y int64, z int64) *Vector3i {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	pinner.Pin(&x)
	pinner.Pin(&y)
	pinner.Pin(&z)
	me := newVector3i()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForVector3i.ctrFromIntIntIntFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), gdc.ConstTypePtr(&z)}))
	return me
}

// Destructor
func (me *Vector3i) Destroy() {
	me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsVector3i() (*Vector3i, error) {
	if me.Type() != gdc.VariantTypeVector3I {
		return nil, fmt.Errorf("variant is not a Vector3i")
	}
	bclass := newVector3i()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForVector3i.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Vector3i) AsVariant() *Variant {
	va := newVariant()
	va.inner = me
	me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForVector3i.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
	return va
}

// Pointers
func Vector3iFromPtr(ptr gdc.ConstTypePtr) *Vector3i {
	me := newVector3i()
	dataFromPtr(me.data[:], ptr)
	return me
}

func (me *Vector3i) ToTypePtr(ptr gdc.TypePtr) {
	dataToPtr(ptr, me.data[:])
}

func (me *Vector3i) Type() gdc.VariantType {
	return gdc.VariantTypeVector3I
}

func (me *Vector3i) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Vector3i) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Vector3i) asUninitialized() gdc.UninitializedTypePtr {
	return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *Vector3i) MinAxisIndex() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3i.methodMinAxisIndexFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector3i) MaxAxisIndex() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3i.methodMaxAxisIndexFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector3i) DistanceTo(to Vector3i) float64 {
	ret := NewFloat()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{to.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3i.methodDistanceToFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector3i) DistanceSquaredTo(to Vector3i) int64 {
	ret := NewInt()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{to.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3i.methodDistanceSquaredToFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector3i) Length() float64 {
	ret := NewFloat()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3i.methodLengthFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector3i) LengthSquared() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3i.methodLengthSquaredFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector3i) Sign() Vector3i {
	ret := NewVector3i()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3i.methodSignFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3i) Abs() Vector3i {
	ret := NewVector3i()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3i.methodAbsFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3i) Clamp(min Vector3i, max Vector3i) Vector3i {
	ret := NewVector3i()

	args := []gdc.ConstTypePtr{min.AsCTypePtr(), max.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3i.methodClampFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector3i) Snapped(step Vector3i) Vector3i {
	ret := NewVector3i()

	args := []gdc.ConstTypePtr{step.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector3i.methodSnappedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

// Operators

func (me *Vector3i) EqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector3i) NotEqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector3i) Negate() Vector3i {
	opPtr := ptrsForVector3i.operatorNegateFn
	ret := NewVector3i()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return *ret
}

func (me *Vector3i) Positive() Vector3i {
	opPtr := ptrsForVector3i.operatorPositiveFn
	ret := NewVector3i()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return *ret
}

func (me *Vector3i) Not() bool {
	opPtr := ptrsForVector3i.operatorNotFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector3i) MultiplyInt(rightArg int64) Vector3i {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForVector3i.operatorMultiplyIntFn
	ret := NewVector3i()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector3i) DivideInt(rightArg int64) Vector3i {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForVector3i.operatorDivideIntFn
	ret := NewVector3i()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector3i) ModuleInt(rightArg int64) Vector3i {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForVector3i.operatorModuleIntFn
	ret := NewVector3i()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector3i) MultiplyFloat32(rightArg float64) Vector3 {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForVector3i.operatorMultiplyFloat32Fn
	ret := NewVector3()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector3i) DivideFloat32(rightArg float64) Vector3 {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForVector3i.operatorDivideFloat32Fn
	ret := NewVector3()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector3i) EqualVector3I(right Vector3i) bool {
	opPtr := ptrsForVector3i.operatorEqualVector3IFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector3i) NotEqualVector3I(right Vector3i) bool {
	opPtr := ptrsForVector3i.operatorNotEqualVector3IFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector3i) LessVector3I(right Vector3i) bool {
	opPtr := ptrsForVector3i.operatorLessVector3IFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector3i) LessEqualVector3I(right Vector3i) bool {
	opPtr := ptrsForVector3i.operatorLessEqualVector3IFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector3i) GreaterVector3I(right Vector3i) bool {
	opPtr := ptrsForVector3i.operatorGreaterVector3IFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector3i) GreaterEqualVector3I(right Vector3i) bool {
	opPtr := ptrsForVector3i.operatorGreaterEqualVector3IFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector3i) AddVector3I(right Vector3i) Vector3i {
	opPtr := ptrsForVector3i.operatorAddVector3IFn
	ret := NewVector3i()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector3i) SubtractVector3I(right Vector3i) Vector3i {
	opPtr := ptrsForVector3i.operatorSubtractVector3IFn
	ret := NewVector3i()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector3i) MultiplyVector3I(right Vector3i) Vector3i {
	opPtr := ptrsForVector3i.operatorMultiplyVector3IFn
	ret := NewVector3i()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector3i) DivideVector3I(right Vector3i) Vector3i {
	opPtr := ptrsForVector3i.operatorDivideVector3IFn
	ret := NewVector3i()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector3i) ModuleVector3I(right Vector3i) Vector3i {
	opPtr := ptrsForVector3i.operatorModuleVector3IFn
	ret := NewVector3i()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector3i) InDictionary(right Dictionary) bool {
	opPtr := ptrsForVector3i.operatorInDictionaryFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector3i) InArray(right Array) bool {
	opPtr := ptrsForVector3i.operatorInArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

// Members

func (me *Vector3i) X() int64 {
	ret := NewInt()
	me.iface.CallPtrGetter(ensurePtr(ptrsForVector3i.memberxGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector3i) SetX(value int64) {
	valueV := NewIntFromInt(value)
	defer valueV.Destroy()
	me.iface.CallPtrSetter(ensurePtr(ptrsForVector3i.memberxSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Vector3i) Y() int64 {
	ret := NewInt()
	me.iface.CallPtrGetter(ensurePtr(ptrsForVector3i.memberyGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector3i) SetY(value int64) {
	valueV := NewIntFromInt(value)
	defer valueV.Destroy()
	me.iface.CallPtrSetter(ensurePtr(ptrsForVector3i.memberySetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Vector3i) Z() int64 {
	ret := NewInt()
	me.iface.CallPtrGetter(ensurePtr(ptrsForVector3i.memberzGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector3i) SetZ(value int64) {
	valueV := NewIntFromInt(value)
	defer valueV.Destroy()
	me.iface.CallPtrSetter(ensurePtr(ptrsForVector3i.memberzSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}
