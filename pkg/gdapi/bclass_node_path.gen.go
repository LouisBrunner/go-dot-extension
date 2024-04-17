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

type ptrsForNodePathList struct {
	ctrFn                           gdc.PtrConstructor
	ctrFromNodePathFn               gdc.PtrConstructor
	ctrFromStringFn                 gdc.PtrConstructor
	destructorFn                    gdc.PtrDestructor
	methodIsAbsoluteFn              gdc.PtrBuiltInMethod
	methodGetNameCountFn            gdc.PtrBuiltInMethod
	methodGetNameFn                 gdc.PtrBuiltInMethod
	methodGetSubnameCountFn         gdc.PtrBuiltInMethod
	methodHashFn                    gdc.PtrBuiltInMethod
	methodGetSubnameFn              gdc.PtrBuiltInMethod
	methodGetConcatenatedNamesFn    gdc.PtrBuiltInMethod
	methodGetConcatenatedSubnamesFn gdc.PtrBuiltInMethod
	methodGetAsPropertyPathFn       gdc.PtrBuiltInMethod
	methodIsEmptyFn                 gdc.PtrBuiltInMethod
	operatorNotFn                   gdc.PtrOperatorEvaluator
	operatorEqualNodePathFn         gdc.PtrOperatorEvaluator
	operatorNotEqualNodePathFn      gdc.PtrOperatorEvaluator
	operatorInDictionaryFn          gdc.PtrOperatorEvaluator
	operatorInArrayFn               gdc.PtrOperatorEvaluator
	toVariantFn                     gdc.TypeFromVariantConstructorFunc
	fromVariantFn                   gdc.VariantFromTypeConstructorFunc
}

var ptrsForNodePath ptrsForNodePathList

func initNodePathPtrs(iface gdc.Interface) {
	ptrsForNodePath.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeNodePath, 0))
	ptrsForNodePath.ctrFromNodePathFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeNodePath, 1))
	ptrsForNodePath.ctrFromStringFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeNodePath, 2))
	ptrsForNodePath.destructorFn = ensurePtr(iface.VariantGetPtrDestructor(gdc.VariantTypeNodePath))
	{
		methodName := StringNameFromStr("is_absolute")
		defer methodName.Destroy()
		ptrsForNodePath.methodIsAbsoluteFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeNodePath, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("get_name_count")
		defer methodName.Destroy()
		ptrsForNodePath.methodGetNameCountFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeNodePath, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("get_name")
		defer methodName.Destroy()
		ptrsForNodePath.methodGetNameFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeNodePath, methodName.AsCPtr(), 2948586938))
	}
	{
		methodName := StringNameFromStr("get_subname_count")
		defer methodName.Destroy()
		ptrsForNodePath.methodGetSubnameCountFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeNodePath, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("hash")
		defer methodName.Destroy()
		ptrsForNodePath.methodHashFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeNodePath, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("get_subname")
		defer methodName.Destroy()
		ptrsForNodePath.methodGetSubnameFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeNodePath, methodName.AsCPtr(), 2948586938))
	}
	{
		methodName := StringNameFromStr("get_concatenated_names")
		defer methodName.Destroy()
		ptrsForNodePath.methodGetConcatenatedNamesFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeNodePath, methodName.AsCPtr(), 1825232092))
	}
	{
		methodName := StringNameFromStr("get_concatenated_subnames")
		defer methodName.Destroy()
		ptrsForNodePath.methodGetConcatenatedSubnamesFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeNodePath, methodName.AsCPtr(), 1825232092))
	}
	{
		methodName := StringNameFromStr("get_as_property_path")
		defer methodName.Destroy()
		ptrsForNodePath.methodGetAsPropertyPathFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeNodePath, methodName.AsCPtr(), 1598598043))
	}
	{
		methodName := StringNameFromStr("is_empty")
		defer methodName.Destroy()
		ptrsForNodePath.methodIsEmptyFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeNodePath, methodName.AsCPtr(), 3918633141))
	}
	ptrsForNodePath.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypeNodePath, gdc.VariantTypeNil))
	ptrsForNodePath.operatorEqualNodePathFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeNodePath, gdc.VariantTypeNodePath))
	ptrsForNodePath.operatorNotEqualNodePathFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeNodePath, gdc.VariantTypeNodePath))
	ptrsForNodePath.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeNodePath, gdc.VariantTypeDictionary))
	ptrsForNodePath.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeNodePath, gdc.VariantTypeArray))
	ptrsForNodePath.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypeNodePath))
	ptrsForNodePath.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypeNodePath))
}

type NodePath struct {
	//data   *[classSizeNodePath]byte
	data   unsafe.Pointer
	iface  gdc.Interface
	pinner runtime.Pinner
}

// Enums

// Constructors
func newNodePath() *NodePath {
	me := &NodePath{
		//data:   new([classSizeNodePath]byte),
		data:  giface.MemAlloc(classSizeNodePath),
		iface: giface,
	}
	me.pinner.Pin(me)
	me.pinner.Pin(me.AsTypePtr())
	return me
}

func NewNodePath() *NodePath {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newNodePath()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForNodePath.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
	return me
}

func NewNodePathFromNodePath(from NodePath) *NodePath {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newNodePath()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForNodePath.ctrFromNodePathFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewNodePathFromString(from String) *NodePath {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newNodePath()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForNodePath.ctrFromStringFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

// Destructor
func (me *NodePath) Destroy() {
	//me.iface.CallPtrDestructor(ensurePtr(ptrsForNodePath.destructorFn), me.AsTypePtr())
	//me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsNodePath() (*NodePath, error) {
	if me.Type() != gdc.VariantTypeNodePath {
		return nil, fmt.Errorf("variant is not a NodePath")
	}
	bclass := newNodePath()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForNodePath.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *NodePath) AsVariant() *Variant {
	va := newVariant()
	va.inner = me
	me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForNodePath.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
	return va
}

// Pointers
func NodePathFromPtr(ptr gdc.ConstTypePtr) *NodePath {
	me := newNodePath()
	dataCopy(me.data, unsafe.Pointer(ptr), classSizeNodePath)
	return me
}

func (me *NodePath) ToTypePtr(ptr gdc.TypePtr) {
	dataCopy(unsafe.Pointer(ptr), me.data, classSizeNodePath)
}

func (me *NodePath) Type() gdc.VariantType {
	return gdc.VariantTypeNodePath
}

func (me *NodePath) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.data)
}

func (me *NodePath) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *NodePath) asUninitialized() gdc.UninitializedTypePtr {
	return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *NodePath) IsAbsolute() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForNodePath.methodIsAbsoluteFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *NodePath) GetNameCount() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForNodePath.methodGetNameCountFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *NodePath) GetName(idx int64) StringName {
	ret := NewStringName()

	varg0 := NewIntFromInt(idx)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForNodePath.methodGetNameFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *NodePath) GetSubnameCount() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForNodePath.methodGetSubnameCountFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *NodePath) Hash() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForNodePath.methodHashFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *NodePath) GetSubname(idx int64) StringName {
	ret := NewStringName()

	varg0 := NewIntFromInt(idx)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForNodePath.methodGetSubnameFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *NodePath) GetConcatenatedNames() StringName {
	ret := NewStringName()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForNodePath.methodGetConcatenatedNamesFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *NodePath) GetConcatenatedSubnames() StringName {
	ret := NewStringName()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForNodePath.methodGetConcatenatedSubnamesFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *NodePath) GetAsPropertyPath() NodePath {
	ret := NewNodePath()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForNodePath.methodGetAsPropertyPathFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *NodePath) IsEmpty() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForNodePath.methodIsEmptyFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

// Operators

func (me *NodePath) EqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *NodePath) NotEqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *NodePath) Not() bool {
	opPtr := ptrsForNodePath.operatorNotFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return ret.Get()
}

func (me *NodePath) EqualNodePath(right NodePath) bool {
	opPtr := ptrsForNodePath.operatorEqualNodePathFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *NodePath) NotEqualNodePath(right NodePath) bool {
	opPtr := ptrsForNodePath.operatorNotEqualNodePathFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *NodePath) InDictionary(right Dictionary) bool {
	opPtr := ptrsForNodePath.operatorInDictionaryFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *NodePath) InArray(right Array) bool {
	opPtr := ptrsForNodePath.operatorInArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

// Members
