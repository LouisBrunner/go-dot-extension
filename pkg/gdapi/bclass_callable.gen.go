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

type ptrsForCallableList struct {
	ctrFn                          gdc.PtrConstructor
	ctrFromCallableFn              gdc.PtrConstructor
	ctrFromObjectStringNameFn      gdc.PtrConstructor
	destructorFn                   gdc.PtrDestructor
	methodCreateFn                 gdc.PtrBuiltInMethod
	methodCallvFn                  gdc.PtrBuiltInMethod
	methodIsNullFn                 gdc.PtrBuiltInMethod
	methodIsCustomFn               gdc.PtrBuiltInMethod
	methodIsStandardFn             gdc.PtrBuiltInMethod
	methodIsValidFn                gdc.PtrBuiltInMethod
	methodGetObjectFn              gdc.PtrBuiltInMethod
	methodGetObjectIdFn            gdc.PtrBuiltInMethod
	methodGetMethodFn              gdc.PtrBuiltInMethod
	methodGetArgumentCountFn       gdc.PtrBuiltInMethod
	methodGetBoundArgumentsCountFn gdc.PtrBuiltInMethod
	methodGetBoundArgumentsFn      gdc.PtrBuiltInMethod
	methodHashFn                   gdc.PtrBuiltInMethod
	methodBindvFn                  gdc.PtrBuiltInMethod
	methodUnbindFn                 gdc.PtrBuiltInMethod
	methodCallFn                   gdc.PtrBuiltInMethod
	methodCallDeferredFn           gdc.PtrBuiltInMethod
	methodRpcFn                    gdc.PtrBuiltInMethod
	methodRpcIdFn                  gdc.PtrBuiltInMethod
	methodBindFn                   gdc.PtrBuiltInMethod
	operatorNotFn                  gdc.PtrOperatorEvaluator
	operatorEqualCallableFn        gdc.PtrOperatorEvaluator
	operatorNotEqualCallableFn     gdc.PtrOperatorEvaluator
	operatorInDictionaryFn         gdc.PtrOperatorEvaluator
	operatorInArrayFn              gdc.PtrOperatorEvaluator
	toVariantFn                    gdc.TypeFromVariantConstructorFunc
	fromVariantFn                  gdc.VariantFromTypeConstructorFunc
}

var ptrsForCallable ptrsForCallableList

func initCallablePtrs(iface gdc.Interface) {
	ptrsForCallable.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeCallable, 0))
	ptrsForCallable.ctrFromCallableFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeCallable, 1))
	ptrsForCallable.ctrFromObjectStringNameFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeCallable, 2))
	ptrsForCallable.destructorFn = ensurePtr(iface.VariantGetPtrDestructor(gdc.VariantTypeCallable))
	{
		methodName := StringNameFromStr("create")
		defer methodName.Destroy()
		ptrsForCallable.methodCreateFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeCallable, methodName.AsCPtr(), 1709381114))
	}
	{
		methodName := StringNameFromStr("callv")
		defer methodName.Destroy()
		ptrsForCallable.methodCallvFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeCallable, methodName.AsCPtr(), 413578926))
	}
	{
		methodName := StringNameFromStr("is_null")
		defer methodName.Destroy()
		ptrsForCallable.methodIsNullFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeCallable, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("is_custom")
		defer methodName.Destroy()
		ptrsForCallable.methodIsCustomFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeCallable, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("is_standard")
		defer methodName.Destroy()
		ptrsForCallable.methodIsStandardFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeCallable, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("is_valid")
		defer methodName.Destroy()
		ptrsForCallable.methodIsValidFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeCallable, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("get_object")
		defer methodName.Destroy()
		ptrsForCallable.methodGetObjectFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeCallable, methodName.AsCPtr(), 4008621732))
	}
	{
		methodName := StringNameFromStr("get_object_id")
		defer methodName.Destroy()
		ptrsForCallable.methodGetObjectIdFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeCallable, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("get_method")
		defer methodName.Destroy()
		ptrsForCallable.methodGetMethodFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeCallable, methodName.AsCPtr(), 1825232092))
	}
	{
		methodName := StringNameFromStr("get_argument_count")
		defer methodName.Destroy()
		ptrsForCallable.methodGetArgumentCountFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeCallable, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("get_bound_arguments_count")
		defer methodName.Destroy()
		ptrsForCallable.methodGetBoundArgumentsCountFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeCallable, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("get_bound_arguments")
		defer methodName.Destroy()
		ptrsForCallable.methodGetBoundArgumentsFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeCallable, methodName.AsCPtr(), 4144163970))
	}
	{
		methodName := StringNameFromStr("hash")
		defer methodName.Destroy()
		ptrsForCallable.methodHashFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeCallable, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("bindv")
		defer methodName.Destroy()
		ptrsForCallable.methodBindvFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeCallable, methodName.AsCPtr(), 3564560322))
	}
	{
		methodName := StringNameFromStr("unbind")
		defer methodName.Destroy()
		ptrsForCallable.methodUnbindFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeCallable, methodName.AsCPtr(), 755001590))
	}
	{
		methodName := StringNameFromStr("call")
		defer methodName.Destroy()
		ptrsForCallable.methodCallFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeCallable, methodName.AsCPtr(), 3643564216))
	}
	{
		methodName := StringNameFromStr("call_deferred")
		defer methodName.Destroy()
		ptrsForCallable.methodCallDeferredFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeCallable, methodName.AsCPtr(), 3286317445))
	}
	{
		methodName := StringNameFromStr("rpc")
		defer methodName.Destroy()
		ptrsForCallable.methodRpcFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeCallable, methodName.AsCPtr(), 3286317445))
	}
	{
		methodName := StringNameFromStr("rpc_id")
		defer methodName.Destroy()
		ptrsForCallable.methodRpcIdFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeCallable, methodName.AsCPtr(), 2270047679))
	}
	{
		methodName := StringNameFromStr("bind")
		defer methodName.Destroy()
		ptrsForCallable.methodBindFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeCallable, methodName.AsCPtr(), 3224143119))
	}
	ptrsForCallable.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypeCallable, gdc.VariantTypeNil))
	ptrsForCallable.operatorEqualCallableFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeCallable, gdc.VariantTypeCallable))
	ptrsForCallable.operatorNotEqualCallableFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeCallable, gdc.VariantTypeCallable))
	ptrsForCallable.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeCallable, gdc.VariantTypeDictionary))
	ptrsForCallable.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeCallable, gdc.VariantTypeArray))
	ptrsForCallable.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypeCallable))
	ptrsForCallable.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypeCallable))
}

type Callable struct {
	data   *[classSizeCallable]byte
	iface  gdc.Interface
	pinner runtime.Pinner
}

// Enums

// Constructors
func newCallable() *Callable {
	me := &Callable{
		data:  new([classSizeCallable]byte),
		iface: giface,
	}
	me.pinner.Pin(me)
	me.pinner.Pin(me.AsTypePtr())
	return me
}

func NewCallable() *Callable {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newCallable()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForCallable.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
	return me
}

func NewCallableFromCallable(from Callable) *Callable {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newCallable()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForCallable.ctrFromCallableFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewCallableFromObjectStringName(object Object, method StringName) *Callable {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newCallable()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForCallable.ctrFromObjectStringNameFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{object.AsCTypePtr(), method.AsCTypePtr()}))
	return me
}

// Destructor
func (me *Callable) Destroy() {
	me.iface.CallPtrDestructor(ensurePtr(ptrsForCallable.destructorFn), me.AsTypePtr())
	me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsCallable() (*Callable, error) {
	if me.Type() != gdc.VariantTypeCallable {
		return nil, fmt.Errorf("variant is not a Callable")
	}
	bclass := newCallable()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForCallable.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Callable) AsVariant() *Variant {
	va := newVariant()
	va.inner = me
	me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForCallable.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
	return va
}

// Pointers
func CallableFromPtr(ptr gdc.ConstTypePtr) *Callable {
	me := newCallable()
	dataFromPtr(me.data[:], ptr)
	return me
}

func (me *Callable) ToTypePtr(ptr gdc.TypePtr) {
	dataToPtr(ptr, me.data[:])
}

func (me *Callable) Type() gdc.VariantType {
	return gdc.VariantTypeCallable
}

func (me *Callable) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Callable) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Callable) asUninitialized() gdc.UninitializedTypePtr {
	return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func CallableCreate(variant Variant, method StringName) Callable {
	ret := NewCallable()

	args := []gdc.ConstTypePtr{variant.AsCTypePtr(), method.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForCallable.methodCreateFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Callable) Callv(arguments Array) Variant {
	ret := NewVariant()

	args := []gdc.ConstTypePtr{arguments.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForCallable.methodCallvFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Callable) IsNull() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForCallable.methodIsNullFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Callable) IsCustom() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForCallable.methodIsCustomFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Callable) IsStandard() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForCallable.methodIsStandardFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Callable) IsValid() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForCallable.methodIsValidFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Callable) GetObject() Object {
	ret := NewObject()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForCallable.methodGetObjectFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Callable) GetObjectId() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForCallable.methodGetObjectIdFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Callable) GetMethod() StringName {
	ret := NewStringName()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForCallable.methodGetMethodFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Callable) GetArgumentCount() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForCallable.methodGetArgumentCountFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Callable) GetBoundArgumentsCount() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForCallable.methodGetBoundArgumentsCountFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Callable) GetBoundArguments() Array {
	ret := NewArray()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForCallable.methodGetBoundArgumentsFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Callable) Hash() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForCallable.methodHashFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Callable) Bindv(arguments Array) Callable {
	ret := NewCallable()

	args := []gdc.ConstTypePtr{arguments.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForCallable.methodBindvFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Callable) Unbind(argcount int64) Callable {
	ret := NewCallable()

	varg0 := NewIntFromInt(argcount)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForCallable.methodUnbindFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Callable) Call(varargs ...Variant) Variant {
	ret := NewVariant()

	args := []gdc.ConstTypePtr{}
	for _, arg := range varargs {
		args = append(args, arg.AsCTypePtr())
	}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForCallable.methodCallFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Callable) CallDeferred(varargs ...Variant) {
	args := []gdc.ConstTypePtr{}
	for _, arg := range varargs {
		args = append(args, arg.AsCTypePtr())
	}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForCallable.methodCallDeferredFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Callable) Rpc(varargs ...Variant) {
	args := []gdc.ConstTypePtr{}
	for _, arg := range varargs {
		args = append(args, arg.AsCTypePtr())
	}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForCallable.methodRpcFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Callable) RpcId(peer_id int64, varargs ...Variant) {
	varg0 := NewIntFromInt(peer_id)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}
	for _, arg := range varargs {
		args = append(args, arg.AsCTypePtr())
	}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForCallable.methodRpcIdFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Callable) Bind(varargs ...Variant) Callable {
	ret := NewCallable()

	args := []gdc.ConstTypePtr{}
	for _, arg := range varargs {
		args = append(args, arg.AsCTypePtr())
	}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForCallable.methodBindFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

// Operators

func (me *Callable) EqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Callable) NotEqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Callable) Not() bool {
	opPtr := ptrsForCallable.operatorNotFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return ret.Get()
}

func (me *Callable) EqualCallable(right Callable) bool {
	opPtr := ptrsForCallable.operatorEqualCallableFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Callable) NotEqualCallable(right Callable) bool {
	opPtr := ptrsForCallable.operatorNotEqualCallableFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Callable) InDictionary(right Dictionary) bool {
	opPtr := ptrsForCallable.operatorInDictionaryFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Callable) InArray(right Array) bool {
	opPtr := ptrsForCallable.operatorInArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

// Members
