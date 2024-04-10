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

type ptrsForRect2iList struct {
	ctrFn                     gdc.PtrConstructor
	ctrFromRect2IFn           gdc.PtrConstructor
	ctrFromRect2Fn            gdc.PtrConstructor
	ctrFromVector2IVector2IFn gdc.PtrConstructor
	ctrFromIntIntIntIntFn     gdc.PtrConstructor
	methodGetCenterFn         gdc.PtrBuiltInMethod
	methodGetAreaFn           gdc.PtrBuiltInMethod
	methodHasAreaFn           gdc.PtrBuiltInMethod
	methodHasPointFn          gdc.PtrBuiltInMethod
	methodIntersectsFn        gdc.PtrBuiltInMethod
	methodEnclosesFn          gdc.PtrBuiltInMethod
	methodIntersectionFn      gdc.PtrBuiltInMethod
	methodMergeFn             gdc.PtrBuiltInMethod
	methodExpandFn            gdc.PtrBuiltInMethod
	methodGrowFn              gdc.PtrBuiltInMethod
	methodGrowSideFn          gdc.PtrBuiltInMethod
	methodGrowIndividualFn    gdc.PtrBuiltInMethod
	methodAbsFn               gdc.PtrBuiltInMethod
	operatorNotFn             gdc.PtrOperatorEvaluator
	operatorEqualRect2IFn     gdc.PtrOperatorEvaluator
	operatorNotEqualRect2IFn  gdc.PtrOperatorEvaluator
	operatorInDictionaryFn    gdc.PtrOperatorEvaluator
	operatorInArrayFn         gdc.PtrOperatorEvaluator
	memberpositionGetterFn    gdc.PtrGetter
	memberpositionSetterFn    gdc.PtrSetter
	membersizeGetterFn        gdc.PtrGetter
	membersizeSetterFn        gdc.PtrSetter
	memberendGetterFn         gdc.PtrGetter
	memberendSetterFn         gdc.PtrSetter
	toVariantFn               gdc.TypeFromVariantConstructorFunc
	fromVariantFn             gdc.VariantFromTypeConstructorFunc
}

var ptrsForRect2i ptrsForRect2iList

func initRect2iPtrs(iface gdc.Interface) {
	ptrsForRect2i.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeRect2I, 0))
	ptrsForRect2i.ctrFromRect2IFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeRect2I, 1))
	ptrsForRect2i.ctrFromRect2Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeRect2I, 2))
	ptrsForRect2i.ctrFromVector2IVector2IFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeRect2I, 3))
	ptrsForRect2i.ctrFromIntIntIntIntFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeRect2I, 4))
	{
		methodName := StringNameFromStr("get_center")
		defer methodName.Destroy()
		ptrsForRect2i.methodGetCenterFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2I, methodName.AsCPtr(), 3444277866))
	}
	{
		methodName := StringNameFromStr("get_area")
		defer methodName.Destroy()
		ptrsForRect2i.methodGetAreaFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2I, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("has_area")
		defer methodName.Destroy()
		ptrsForRect2i.methodHasAreaFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2I, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("has_point")
		defer methodName.Destroy()
		ptrsForRect2i.methodHasPointFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2I, methodName.AsCPtr(), 328189994))
	}
	{
		methodName := StringNameFromStr("intersects")
		defer methodName.Destroy()
		ptrsForRect2i.methodIntersectsFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2I, methodName.AsCPtr(), 3434691493))
	}
	{
		methodName := StringNameFromStr("encloses")
		defer methodName.Destroy()
		ptrsForRect2i.methodEnclosesFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2I, methodName.AsCPtr(), 3434691493))
	}
	{
		methodName := StringNameFromStr("intersection")
		defer methodName.Destroy()
		ptrsForRect2i.methodIntersectionFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2I, methodName.AsCPtr(), 717431873))
	}
	{
		methodName := StringNameFromStr("merge")
		defer methodName.Destroy()
		ptrsForRect2i.methodMergeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2I, methodName.AsCPtr(), 717431873))
	}
	{
		methodName := StringNameFromStr("expand")
		defer methodName.Destroy()
		ptrsForRect2i.methodExpandFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2I, methodName.AsCPtr(), 1355196872))
	}
	{
		methodName := StringNameFromStr("grow")
		defer methodName.Destroy()
		ptrsForRect2i.methodGrowFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2I, methodName.AsCPtr(), 1578070074))
	}
	{
		methodName := StringNameFromStr("grow_side")
		defer methodName.Destroy()
		ptrsForRect2i.methodGrowSideFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2I, methodName.AsCPtr(), 3191154199))
	}
	{
		methodName := StringNameFromStr("grow_individual")
		defer methodName.Destroy()
		ptrsForRect2i.methodGrowIndividualFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2I, methodName.AsCPtr(), 1893743416))
	}
	{
		methodName := StringNameFromStr("abs")
		defer methodName.Destroy()
		ptrsForRect2i.methodAbsFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRect2I, methodName.AsCPtr(), 1469025700))
	}
	ptrsForRect2i.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypeRect2I, gdc.VariantTypeNil))
	ptrsForRect2i.operatorEqualRect2IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeRect2I, gdc.VariantTypeRect2I))
	ptrsForRect2i.operatorNotEqualRect2IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeRect2I, gdc.VariantTypeRect2I))
	ptrsForRect2i.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeRect2I, gdc.VariantTypeDictionary))
	ptrsForRect2i.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeRect2I, gdc.VariantTypeArray))
	{
		memberName := StringNameFromStr("position")
		defer memberName.Destroy()
		ptrsForRect2i.memberpositionGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeRect2I, memberName.AsCPtr()))
		ptrsForRect2i.memberpositionSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeRect2I, memberName.AsCPtr()))
	}
	{
		memberName := StringNameFromStr("size")
		defer memberName.Destroy()
		ptrsForRect2i.membersizeGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeRect2I, memberName.AsCPtr()))
		ptrsForRect2i.membersizeSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeRect2I, memberName.AsCPtr()))
	}
	{
		memberName := StringNameFromStr("end")
		defer memberName.Destroy()
		ptrsForRect2i.memberendGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeRect2I, memberName.AsCPtr()))
		ptrsForRect2i.memberendSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeRect2I, memberName.AsCPtr()))
	}
	ptrsForRect2i.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypeRect2I))
	ptrsForRect2i.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypeRect2I))
}

type Rect2i struct {
	data   *[classSizeRect2i]byte
	iface  gdc.Interface
	pinner runtime.Pinner
}

// Enums

// Constructors
func newRect2i() *Rect2i {
	me := &Rect2i{
		data:  new([classSizeRect2i]byte),
		iface: giface,
	}
	me.pinner.Pin(me)
	me.pinner.Pin(me.AsTypePtr())
	return me
}

func NewRect2i() *Rect2i {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newRect2i()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForRect2i.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
	return me
}

func NewRect2iFromRect2I(from Rect2i) *Rect2i {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newRect2i()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForRect2i.ctrFromRect2IFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewRect2iFromRect2(from Rect2) *Rect2i {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newRect2i()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForRect2i.ctrFromRect2Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewRect2iFromVector2IVector2I(position Vector2i, size Vector2i) *Rect2i {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newRect2i()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForRect2i.ctrFromVector2IVector2IFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{position.AsCTypePtr(), size.AsCTypePtr()}))
	return me
}

func NewRect2iFromIntIntIntInt(x int64, y int64, width int64, height int64) *Rect2i {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	pinner.Pin(&x)
	pinner.Pin(&y)
	pinner.Pin(&width)
	pinner.Pin(&height)
	me := newRect2i()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForRect2i.ctrFromIntIntIntIntFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height)}))
	return me
}

// Destructor
func (me *Rect2i) Destroy() {
	me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsRect2i() (*Rect2i, error) {
	if me.Type() != gdc.VariantTypeRect2I {
		return nil, fmt.Errorf("variant is not a Rect2i")
	}
	bclass := newRect2i()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForRect2i.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Rect2i) AsVariant() *Variant {
	va := newVariant()
	va.inner = me
	me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForRect2i.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
	return va
}

// Pointers
func Rect2iFromPtr(ptr gdc.ConstTypePtr) *Rect2i {
	me := newRect2i()
	dataFromPtr(me.data[:], ptr)
	return me
}

func (me *Rect2i) ToTypePtr(ptr gdc.TypePtr) {
	dataToPtr(ptr, me.data[:])
}

func (me *Rect2i) Type() gdc.VariantType {
	return gdc.VariantTypeRect2I
}

func (me *Rect2i) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Rect2i) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Rect2i) asUninitialized() gdc.UninitializedTypePtr {
	return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *Rect2i) GetCenter() Vector2i {
	ret := NewVector2i()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2i.methodGetCenterFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Rect2i) GetArea() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2i.methodGetAreaFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Rect2i) HasArea() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2i.methodHasAreaFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Rect2i) HasPoint(point Vector2i) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{point.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2i.methodHasPointFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Rect2i) Intersects(b Rect2i) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{b.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2i.methodIntersectsFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Rect2i) Encloses(b Rect2i) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{b.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2i.methodEnclosesFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Rect2i) Intersection(b Rect2i) Rect2i {
	ret := NewRect2i()

	args := []gdc.ConstTypePtr{b.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2i.methodIntersectionFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Rect2i) Merge(b Rect2i) Rect2i {
	ret := NewRect2i()

	args := []gdc.ConstTypePtr{b.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2i.methodMergeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Rect2i) Expand(to Vector2i) Rect2i {
	ret := NewRect2i()

	args := []gdc.ConstTypePtr{to.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2i.methodExpandFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Rect2i) Grow(amount int64) Rect2i {
	ret := NewRect2i()

	varg0 := NewIntFromInt(amount)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2i.methodGrowFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Rect2i) GrowSide(side int64, amount int64) Rect2i {
	ret := NewRect2i()

	varg0 := NewIntFromInt(side)
	defer varg0.Destroy()
	varg1 := NewIntFromInt(amount)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2i.methodGrowSideFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Rect2i) GrowIndividual(left int64, top int64, right int64, bottom int64) Rect2i {
	ret := NewRect2i()

	varg0 := NewIntFromInt(left)
	defer varg0.Destroy()
	varg1 := NewIntFromInt(top)
	defer varg1.Destroy()
	varg2 := NewIntFromInt(right)
	defer varg2.Destroy()
	varg3 := NewIntFromInt(bottom)
	defer varg3.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), varg3.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2i.methodGrowIndividualFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Rect2i) Abs() Rect2i {
	ret := NewRect2i()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRect2i.methodAbsFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

// Operators

func (me *Rect2i) EqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Rect2i) NotEqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Rect2i) Not() bool {
	opPtr := ptrsForRect2i.operatorNotFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return ret.Get()
}

func (me *Rect2i) EqualRect2I(right Rect2i) bool {
	opPtr := ptrsForRect2i.operatorEqualRect2IFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Rect2i) NotEqualRect2I(right Rect2i) bool {
	opPtr := ptrsForRect2i.operatorNotEqualRect2IFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Rect2i) InDictionary(right Dictionary) bool {
	opPtr := ptrsForRect2i.operatorInDictionaryFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Rect2i) InArray(right Array) bool {
	opPtr := ptrsForRect2i.operatorInArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

// Members

func (me *Rect2i) Position() Vector2i {
	ret := NewVector2i()
	me.iface.CallPtrGetter(ensurePtr(ptrsForRect2i.memberpositionGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Rect2i) SetPosition(value Vector2i) {
	valueV := value
	me.iface.CallPtrSetter(ensurePtr(ptrsForRect2i.memberpositionSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Rect2i) Size() Vector2i {
	ret := NewVector2i()
	me.iface.CallPtrGetter(ensurePtr(ptrsForRect2i.membersizeGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Rect2i) SetSize(value Vector2i) {
	valueV := value
	me.iface.CallPtrSetter(ensurePtr(ptrsForRect2i.membersizeSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Rect2i) End() Vector2i {
	ret := NewVector2i()
	me.iface.CallPtrGetter(ensurePtr(ptrsForRect2i.memberendGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Rect2i) SetEnd(value Vector2i) {
	valueV := value
	me.iface.CallPtrSetter(ensurePtr(ptrsForRect2i.memberendSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}
