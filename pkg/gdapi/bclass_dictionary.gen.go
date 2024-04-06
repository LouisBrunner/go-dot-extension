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

type ptrsForDictionaryList struct {
	ctrFn                        gdc.PtrConstructor
	ctrFromDictionaryFn          gdc.PtrConstructor
	destructorFn                 gdc.PtrDestructor
	methodSizeFn                 gdc.PtrBuiltInMethod
	methodIsEmptyFn              gdc.PtrBuiltInMethod
	methodClearFn                gdc.PtrBuiltInMethod
	methodMergeFn                gdc.PtrBuiltInMethod
	methodHasFn                  gdc.PtrBuiltInMethod
	methodHasAllFn               gdc.PtrBuiltInMethod
	methodFindKeyFn              gdc.PtrBuiltInMethod
	methodEraseFn                gdc.PtrBuiltInMethod
	methodHashFn                 gdc.PtrBuiltInMethod
	methodKeysFn                 gdc.PtrBuiltInMethod
	methodValuesFn               gdc.PtrBuiltInMethod
	methodDuplicateFn            gdc.PtrBuiltInMethod
	methodGetFn                  gdc.PtrBuiltInMethod
	methodMakeReadOnlyFn         gdc.PtrBuiltInMethod
	methodIsReadOnlyFn           gdc.PtrBuiltInMethod
	operatorNotFn                gdc.PtrOperatorEvaluator
	operatorEqualDictionaryFn    gdc.PtrOperatorEvaluator
	operatorNotEqualDictionaryFn gdc.PtrOperatorEvaluator
	operatorInDictionaryFn       gdc.PtrOperatorEvaluator
	operatorInArrayFn            gdc.PtrOperatorEvaluator
	toVariantFn                  gdc.TypeFromVariantConstructorFunc
	fromVariantFn                gdc.VariantFromTypeConstructorFunc
}

var ptrsForDictionary ptrsForDictionaryList

func initDictionaryPtrs(iface gdc.Interface) {
	ptrsForDictionary.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeDictionary, 0))
	ptrsForDictionary.ctrFromDictionaryFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeDictionary, 1))
	ptrsForDictionary.destructorFn = ensurePtr(iface.VariantGetPtrDestructor(gdc.VariantTypeDictionary))
	{
		methodName := StringNameFromStr("size")
		defer methodName.Destroy()
		ptrsForDictionary.methodSizeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeDictionary, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("is_empty")
		defer methodName.Destroy()
		ptrsForDictionary.methodIsEmptyFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeDictionary, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForDictionary.methodClearFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeDictionary, methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("merge")
		defer methodName.Destroy()
		ptrsForDictionary.methodMergeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeDictionary, methodName.AsCPtr(), 2079548978))
	}
	{
		methodName := StringNameFromStr("has")
		defer methodName.Destroy()
		ptrsForDictionary.methodHasFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeDictionary, methodName.AsCPtr(), 3680194679))
	}
	{
		methodName := StringNameFromStr("has_all")
		defer methodName.Destroy()
		ptrsForDictionary.methodHasAllFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeDictionary, methodName.AsCPtr(), 2988181878))
	}
	{
		methodName := StringNameFromStr("find_key")
		defer methodName.Destroy()
		ptrsForDictionary.methodFindKeyFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeDictionary, methodName.AsCPtr(), 1988825835))
	}
	{
		methodName := StringNameFromStr("erase")
		defer methodName.Destroy()
		ptrsForDictionary.methodEraseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeDictionary, methodName.AsCPtr(), 1776646889))
	}
	{
		methodName := StringNameFromStr("hash")
		defer methodName.Destroy()
		ptrsForDictionary.methodHashFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeDictionary, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("keys")
		defer methodName.Destroy()
		ptrsForDictionary.methodKeysFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeDictionary, methodName.AsCPtr(), 4144163970))
	}
	{
		methodName := StringNameFromStr("values")
		defer methodName.Destroy()
		ptrsForDictionary.methodValuesFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeDictionary, methodName.AsCPtr(), 4144163970))
	}
	{
		methodName := StringNameFromStr("duplicate")
		defer methodName.Destroy()
		ptrsForDictionary.methodDuplicateFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeDictionary, methodName.AsCPtr(), 830099069))
	}
	{
		methodName := StringNameFromStr("get")
		defer methodName.Destroy()
		ptrsForDictionary.methodGetFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeDictionary, methodName.AsCPtr(), 2205440559))
	}
	{
		methodName := StringNameFromStr("make_read_only")
		defer methodName.Destroy()
		ptrsForDictionary.methodMakeReadOnlyFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeDictionary, methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("is_read_only")
		defer methodName.Destroy()
		ptrsForDictionary.methodIsReadOnlyFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeDictionary, methodName.AsCPtr(), 3918633141))
	}
	ptrsForDictionary.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypeDictionary, gdc.VariantTypeNil))
	ptrsForDictionary.operatorEqualDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeDictionary, gdc.VariantTypeDictionary))
	ptrsForDictionary.operatorNotEqualDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeDictionary, gdc.VariantTypeDictionary))
	ptrsForDictionary.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeDictionary, gdc.VariantTypeDictionary))
	ptrsForDictionary.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeDictionary, gdc.VariantTypeArray))
	ptrsForDictionary.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypeDictionary))
	ptrsForDictionary.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypeDictionary))
}

type Dictionary struct {
	data   *[classSizeDictionary]byte
	iface  gdc.Interface
	pinner runtime.Pinner
}

// Enums

// Constructors
func newDictionary() *Dictionary {
	me := &Dictionary{
		data:  new([classSizeDictionary]byte),
		iface: giface,
	}
	me.pinner.Pin(me)
	me.pinner.Pin(me.AsTypePtr())
	return me
}

func NewDictionary() *Dictionary {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newDictionary()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForDictionary.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
	return me
}

func NewDictionaryFromDictionary(from Dictionary) *Dictionary {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newDictionary()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForDictionary.ctrFromDictionaryFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

// Destructor
func (me *Dictionary) Destroy() {
	me.iface.CallPtrDestructor(ensurePtr(ptrsForDictionary.destructorFn), me.AsTypePtr())
	me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsDictionary() (*Dictionary, error) {
	if me.Type() != gdc.VariantTypeDictionary {
		return nil, fmt.Errorf("variant is not a Dictionary")
	}
	bclass := newDictionary()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForDictionary.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Dictionary) AsVariant() *Variant {
	va := newVariant()
	va.inner = me
	me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForDictionary.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
	return va
}

// Pointers
func DictionaryFromPtr(ptr gdc.ConstTypePtr) *Dictionary {
	me := newDictionary()
	dataFromPtr(me.data[:], ptr)
	return me
}

func (me *Dictionary) Type() gdc.VariantType {
	return gdc.VariantTypeDictionary
}

func (me *Dictionary) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Dictionary) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Dictionary) asUninitialized() gdc.UninitializedTypePtr {
	return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *Dictionary) Size() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForDictionary.methodSizeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Dictionary) IsEmpty() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForDictionary.methodIsEmptyFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Dictionary) Clear() {
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForDictionary.methodClearFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Dictionary) Merge(dictionary Dictionary, overwrite bool) {

	varg1 := NewBoolFromBool(overwrite)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{dictionary.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForDictionary.methodMergeFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Dictionary) Has(key Variant) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{key.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForDictionary.methodHasFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Dictionary) HasAll(keys Array) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{keys.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForDictionary.methodHasAllFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Dictionary) FindKey(value Variant) Variant {
	ret := NewVariant()

	args := []gdc.ConstTypePtr{value.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForDictionary.methodFindKeyFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Dictionary) Erase(key Variant) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{key.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForDictionary.methodEraseFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Dictionary) Hash() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForDictionary.methodHashFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Dictionary) Keys() Array {
	ret := NewArray()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForDictionary.methodKeysFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Dictionary) Values() Array {
	ret := NewArray()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForDictionary.methodValuesFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Dictionary) Duplicate(deep bool) Dictionary {
	ret := NewDictionary()

	varg0 := NewBoolFromBool(deep)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForDictionary.methodDuplicateFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Dictionary) Get(key Variant, default_ Variant) Variant {
	ret := NewVariant()

	args := []gdc.ConstTypePtr{key.AsCTypePtr(), default_.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForDictionary.methodGetFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Dictionary) MakeReadOnly() {
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForDictionary.methodMakeReadOnlyFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Dictionary) IsReadOnly() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForDictionary.methodIsReadOnlyFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

// Operators

func (me *Dictionary) EqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Dictionary) NotEqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Dictionary) Not() bool {
	opPtr := ptrsForDictionary.operatorNotFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return ret.Get()
}

func (me *Dictionary) EqualDictionary(right Dictionary) bool {
	opPtr := ptrsForDictionary.operatorEqualDictionaryFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Dictionary) NotEqualDictionary(right Dictionary) bool {
	opPtr := ptrsForDictionary.operatorNotEqualDictionaryFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Dictionary) InDictionary(right Dictionary) bool {
	opPtr := ptrsForDictionary.operatorInDictionaryFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Dictionary) InArray(right Array) bool {
	opPtr := ptrsForDictionary.operatorInArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

// Members
