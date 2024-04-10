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

type ptrsForVector2iList struct {
	ctrFn                          gdc.PtrConstructor
	ctrFromVector2IFn              gdc.PtrConstructor
	ctrFromVector2Fn               gdc.PtrConstructor
	ctrFromIntIntFn                gdc.PtrConstructor
	methodAspectFn                 gdc.PtrBuiltInMethod
	methodMaxAxisIndexFn           gdc.PtrBuiltInMethod
	methodMinAxisIndexFn           gdc.PtrBuiltInMethod
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
	operatorEqualVector2IFn        gdc.PtrOperatorEvaluator
	operatorNotEqualVector2IFn     gdc.PtrOperatorEvaluator
	operatorLessVector2IFn         gdc.PtrOperatorEvaluator
	operatorLessEqualVector2IFn    gdc.PtrOperatorEvaluator
	operatorGreaterVector2IFn      gdc.PtrOperatorEvaluator
	operatorGreaterEqualVector2IFn gdc.PtrOperatorEvaluator
	operatorAddVector2IFn          gdc.PtrOperatorEvaluator
	operatorSubtractVector2IFn     gdc.PtrOperatorEvaluator
	operatorMultiplyVector2IFn     gdc.PtrOperatorEvaluator
	operatorDivideVector2IFn       gdc.PtrOperatorEvaluator
	operatorModuleVector2IFn       gdc.PtrOperatorEvaluator
	operatorInDictionaryFn         gdc.PtrOperatorEvaluator
	operatorInArrayFn              gdc.PtrOperatorEvaluator
	memberxGetterFn                gdc.PtrGetter
	memberxSetterFn                gdc.PtrSetter
	memberyGetterFn                gdc.PtrGetter
	memberySetterFn                gdc.PtrSetter
	toVariantFn                    gdc.TypeFromVariantConstructorFunc
	fromVariantFn                  gdc.VariantFromTypeConstructorFunc
}

var ptrsForVector2i ptrsForVector2iList

func initVector2iPtrs(iface gdc.Interface) {
	ptrsForVector2i.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeVector2I, 0))
	ptrsForVector2i.ctrFromVector2IFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeVector2I, 1))
	ptrsForVector2i.ctrFromVector2Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeVector2I, 2))
	ptrsForVector2i.ctrFromIntIntFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeVector2I, 3))
	{
		methodName := StringNameFromStr("aspect")
		defer methodName.Destroy()
		ptrsForVector2i.methodAspectFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2I, methodName.AsCPtr(), 466405837))
	}
	{
		methodName := StringNameFromStr("max_axis_index")
		defer methodName.Destroy()
		ptrsForVector2i.methodMaxAxisIndexFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2I, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("min_axis_index")
		defer methodName.Destroy()
		ptrsForVector2i.methodMinAxisIndexFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2I, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("length")
		defer methodName.Destroy()
		ptrsForVector2i.methodLengthFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2I, methodName.AsCPtr(), 466405837))
	}
	{
		methodName := StringNameFromStr("length_squared")
		defer methodName.Destroy()
		ptrsForVector2i.methodLengthSquaredFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2I, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("sign")
		defer methodName.Destroy()
		ptrsForVector2i.methodSignFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2I, methodName.AsCPtr(), 3444277866))
	}
	{
		methodName := StringNameFromStr("abs")
		defer methodName.Destroy()
		ptrsForVector2i.methodAbsFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2I, methodName.AsCPtr(), 3444277866))
	}
	{
		methodName := StringNameFromStr("clamp")
		defer methodName.Destroy()
		ptrsForVector2i.methodClampFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2I, methodName.AsCPtr(), 186568249))
	}
	{
		methodName := StringNameFromStr("snapped")
		defer methodName.Destroy()
		ptrsForVector2i.methodSnappedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeVector2I, methodName.AsCPtr(), 1735278196))
	}
	ptrsForVector2i.operatorNegateFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNegate, gdc.VariantTypeVector2I, gdc.VariantTypeNil))
	ptrsForVector2i.operatorPositiveFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpPositive, gdc.VariantTypeVector2I, gdc.VariantTypeNil))
	ptrsForVector2i.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypeVector2I, gdc.VariantTypeNil))
	ptrsForVector2i.operatorMultiplyIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeVector2I, gdc.VariantTypeInt))
	ptrsForVector2i.operatorDivideIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeVector2I, gdc.VariantTypeInt))
	ptrsForVector2i.operatorModuleIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeVector2I, gdc.VariantTypeInt))
	ptrsForVector2i.operatorMultiplyFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeVector2I, gdc.VariantTypeFloat))
	ptrsForVector2i.operatorDivideFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeVector2I, gdc.VariantTypeFloat))
	ptrsForVector2i.operatorEqualVector2IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeVector2I, gdc.VariantTypeVector2I))
	ptrsForVector2i.operatorNotEqualVector2IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeVector2I, gdc.VariantTypeVector2I))
	ptrsForVector2i.operatorLessVector2IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLess, gdc.VariantTypeVector2I, gdc.VariantTypeVector2I))
	ptrsForVector2i.operatorLessEqualVector2IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLessEqual, gdc.VariantTypeVector2I, gdc.VariantTypeVector2I))
	ptrsForVector2i.operatorGreaterVector2IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreater, gdc.VariantTypeVector2I, gdc.VariantTypeVector2I))
	ptrsForVector2i.operatorGreaterEqualVector2IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreaterEqual, gdc.VariantTypeVector2I, gdc.VariantTypeVector2I))
	ptrsForVector2i.operatorAddVector2IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, gdc.VariantTypeVector2I, gdc.VariantTypeVector2I))
	ptrsForVector2i.operatorSubtractVector2IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpSubtract, gdc.VariantTypeVector2I, gdc.VariantTypeVector2I))
	ptrsForVector2i.operatorMultiplyVector2IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeVector2I, gdc.VariantTypeVector2I))
	ptrsForVector2i.operatorDivideVector2IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeVector2I, gdc.VariantTypeVector2I))
	ptrsForVector2i.operatorModuleVector2IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeVector2I, gdc.VariantTypeVector2I))
	ptrsForVector2i.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeVector2I, gdc.VariantTypeDictionary))
	ptrsForVector2i.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeVector2I, gdc.VariantTypeArray))
	{
		memberName := StringNameFromStr("x")
		defer memberName.Destroy()
		ptrsForVector2i.memberxGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeVector2I, memberName.AsCPtr()))
		ptrsForVector2i.memberxSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeVector2I, memberName.AsCPtr()))
	}
	{
		memberName := StringNameFromStr("y")
		defer memberName.Destroy()
		ptrsForVector2i.memberyGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeVector2I, memberName.AsCPtr()))
		ptrsForVector2i.memberySetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeVector2I, memberName.AsCPtr()))
	}
	ptrsForVector2i.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypeVector2I))
	ptrsForVector2i.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypeVector2I))
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("ZERO")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector2I, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector2i()
		if err != nil {
			panic("Failed to get constant value ZERO: " + err.Error())
		}
		Vector2iZero = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("ONE")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector2I, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector2i()
		if err != nil {
			panic("Failed to get constant value ONE: " + err.Error())
		}
		Vector2iOne = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("MIN")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector2I, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector2i()
		if err != nil {
			panic("Failed to get constant value MIN: " + err.Error())
		}
		Vector2iMin = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("MAX")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector2I, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector2i()
		if err != nil {
			panic("Failed to get constant value MAX: " + err.Error())
		}
		Vector2iMax = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("LEFT")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector2I, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector2i()
		if err != nil {
			panic("Failed to get constant value LEFT: " + err.Error())
		}
		Vector2iLeft = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("RIGHT")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector2I, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector2i()
		if err != nil {
			panic("Failed to get constant value RIGHT: " + err.Error())
		}
		Vector2iRight = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("UP")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector2I, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector2i()
		if err != nil {
			panic("Failed to get constant value UP: " + err.Error())
		}
		Vector2iUp = *cnst
	}
	{
		va := *newVariant()
		defer va.Destroy()
		name := StringNameFromStr("DOWN")
		defer name.Destroy()
		iface.VariantGetConstantValue(gdc.VariantTypeVector2I, name.AsCPtr(), va.asUninitialized())
		cnst, err := va.AsVector2i()
		if err != nil {
			panic("Failed to get constant value DOWN: " + err.Error())
		}
		Vector2iDown = *cnst
	}
}

type Vector2i struct {
	data   *[classSizeVector2i]byte
	iface  gdc.Interface
	pinner runtime.Pinner
}

// Constants

var (
	Vector2iZero  Vector2i
	Vector2iOne   Vector2i
	Vector2iMin   Vector2i
	Vector2iMax   Vector2i
	Vector2iLeft  Vector2i
	Vector2iRight Vector2i
	Vector2iUp    Vector2i
	Vector2iDown  Vector2i
)

// Enums

type Vector2iAxis int

const (
	Vector2iAxisAxisX Vector2iAxis = 0
	Vector2iAxisAxisY Vector2iAxis = 1
)

// Constructors
func newVector2i() *Vector2i {
	me := &Vector2i{
		data:  new([classSizeVector2i]byte),
		iface: giface,
	}
	me.pinner.Pin(me)
	me.pinner.Pin(me.AsTypePtr())
	return me
}

func NewVector2i() *Vector2i {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newVector2i()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForVector2i.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
	return me
}

func NewVector2iFromVector2I(from Vector2i) *Vector2i {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newVector2i()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForVector2i.ctrFromVector2IFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewVector2iFromVector2(from Vector2) *Vector2i {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newVector2i()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForVector2i.ctrFromVector2Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewVector2iFromIntInt(x int64, y int64) *Vector2i {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	pinner.Pin(&x)
	pinner.Pin(&y)
	me := newVector2i()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForVector2i.ctrFromIntIntFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y)}))
	return me
}

// Destructor
func (me *Vector2i) Destroy() {
	me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsVector2i() (*Vector2i, error) {
	if me.Type() != gdc.VariantTypeVector2I {
		return nil, fmt.Errorf("variant is not a Vector2i")
	}
	bclass := newVector2i()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForVector2i.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Vector2i) AsVariant() *Variant {
	va := newVariant()
	va.inner = me
	me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForVector2i.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
	return va
}

// Pointers
func Vector2iFromPtr(ptr gdc.ConstTypePtr) *Vector2i {
	me := newVector2i()
	dataFromPtr(me.data[:], ptr)
	return me
}

func (me *Vector2i) ToTypePtr(ptr gdc.TypePtr) {
	dataToPtr(ptr, me.data[:])
}

func (me *Vector2i) Type() gdc.VariantType {
	return gdc.VariantTypeVector2I
}

func (me *Vector2i) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Vector2i) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Vector2i) asUninitialized() gdc.UninitializedTypePtr {
	return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *Vector2i) Aspect() float64 {
	ret := NewFloat()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2i.methodAspectFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector2i) MaxAxisIndex() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2i.methodMaxAxisIndexFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector2i) MinAxisIndex() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2i.methodMinAxisIndexFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector2i) Length() float64 {
	ret := NewFloat()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2i.methodLengthFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector2i) LengthSquared() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2i.methodLengthSquaredFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Vector2i) Sign() Vector2i {
	ret := NewVector2i()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2i.methodSignFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector2i) Abs() Vector2i {
	ret := NewVector2i()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2i.methodAbsFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector2i) Clamp(min Vector2i, max Vector2i) Vector2i {
	ret := NewVector2i()

	args := []gdc.ConstTypePtr{min.AsCTypePtr(), max.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2i.methodClampFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Vector2i) Snapped(step Vector2i) Vector2i {
	ret := NewVector2i()

	args := []gdc.ConstTypePtr{step.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForVector2i.methodSnappedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

// Operators

func (me *Vector2i) EqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector2i) NotEqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector2i) Negate() Vector2i {
	opPtr := ptrsForVector2i.operatorNegateFn
	ret := NewVector2i()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return *ret
}

func (me *Vector2i) Positive() Vector2i {
	opPtr := ptrsForVector2i.operatorPositiveFn
	ret := NewVector2i()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return *ret
}

func (me *Vector2i) Not() bool {
	opPtr := ptrsForVector2i.operatorNotFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector2i) MultiplyInt(rightArg int64) Vector2i {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForVector2i.operatorMultiplyIntFn
	ret := NewVector2i()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector2i) DivideInt(rightArg int64) Vector2i {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForVector2i.operatorDivideIntFn
	ret := NewVector2i()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector2i) ModuleInt(rightArg int64) Vector2i {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForVector2i.operatorModuleIntFn
	ret := NewVector2i()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector2i) MultiplyFloat32(rightArg float64) Vector2 {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForVector2i.operatorMultiplyFloat32Fn
	ret := NewVector2()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector2i) DivideFloat32(rightArg float64) Vector2 {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForVector2i.operatorDivideFloat32Fn
	ret := NewVector2()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector2i) EqualVector2I(right Vector2i) bool {
	opPtr := ptrsForVector2i.operatorEqualVector2IFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector2i) NotEqualVector2I(right Vector2i) bool {
	opPtr := ptrsForVector2i.operatorNotEqualVector2IFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector2i) LessVector2I(right Vector2i) bool {
	opPtr := ptrsForVector2i.operatorLessVector2IFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector2i) LessEqualVector2I(right Vector2i) bool {
	opPtr := ptrsForVector2i.operatorLessEqualVector2IFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector2i) GreaterVector2I(right Vector2i) bool {
	opPtr := ptrsForVector2i.operatorGreaterVector2IFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector2i) GreaterEqualVector2I(right Vector2i) bool {
	opPtr := ptrsForVector2i.operatorGreaterEqualVector2IFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector2i) AddVector2I(right Vector2i) Vector2i {
	opPtr := ptrsForVector2i.operatorAddVector2IFn
	ret := NewVector2i()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector2i) SubtractVector2I(right Vector2i) Vector2i {
	opPtr := ptrsForVector2i.operatorSubtractVector2IFn
	ret := NewVector2i()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector2i) MultiplyVector2I(right Vector2i) Vector2i {
	opPtr := ptrsForVector2i.operatorMultiplyVector2IFn
	ret := NewVector2i()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector2i) DivideVector2I(right Vector2i) Vector2i {
	opPtr := ptrsForVector2i.operatorDivideVector2IFn
	ret := NewVector2i()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector2i) ModuleVector2I(right Vector2i) Vector2i {
	opPtr := ptrsForVector2i.operatorModuleVector2IFn
	ret := NewVector2i()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Vector2i) InDictionary(right Dictionary) bool {
	opPtr := ptrsForVector2i.operatorInDictionaryFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector2i) InArray(right Array) bool {
	opPtr := ptrsForVector2i.operatorInArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

// Members

func (me *Vector2i) X() int64 {
	ret := NewInt()
	me.iface.CallPtrGetter(ensurePtr(ptrsForVector2i.memberxGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector2i) SetX(value int64) {
	valueV := NewIntFromInt(value)
	defer valueV.Destroy()
	me.iface.CallPtrSetter(ensurePtr(ptrsForVector2i.memberxSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Vector2i) Y() int64 {
	ret := NewInt()
	me.iface.CallPtrGetter(ensurePtr(ptrsForVector2i.memberyGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Vector2i) SetY(value int64) {
	valueV := NewIntFromInt(value)
	defer valueV.Destroy()
	me.iface.CallPtrSetter(ensurePtr(ptrsForVector2i.memberySetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}
