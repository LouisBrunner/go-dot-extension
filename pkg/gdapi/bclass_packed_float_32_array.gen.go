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

type ptrsForPackedFloat32ArrayList struct {
  ctrFn gdc.PtrConstructor
  ctrFromPackedFloat32ArrayFn gdc.PtrConstructor
  ctrFromArrayFn gdc.PtrConstructor
  destructorFn gdc.PtrDestructor
  methodSizeFn gdc.PtrBuiltInMethod
  methodIsEmptyFn gdc.PtrBuiltInMethod
  methodSetFn gdc.PtrBuiltInMethod
  methodPushBackFn gdc.PtrBuiltInMethod
  methodAppendFn gdc.PtrBuiltInMethod
  methodAppendArrayFn gdc.PtrBuiltInMethod
  methodRemoveAtFn gdc.PtrBuiltInMethod
  methodInsertFn gdc.PtrBuiltInMethod
  methodFillFn gdc.PtrBuiltInMethod
  methodResizeFn gdc.PtrBuiltInMethod
  methodClearFn gdc.PtrBuiltInMethod
  methodHasFn gdc.PtrBuiltInMethod
  methodReverseFn gdc.PtrBuiltInMethod
  methodSliceFn gdc.PtrBuiltInMethod
  methodToByteArrayFn gdc.PtrBuiltInMethod
  methodSortFn gdc.PtrBuiltInMethod
  methodBsearchFn gdc.PtrBuiltInMethod
  methodDuplicateFn gdc.PtrBuiltInMethod
  methodFindFn gdc.PtrBuiltInMethod
  methodRfindFn gdc.PtrBuiltInMethod
  methodCountFn gdc.PtrBuiltInMethod
  operatorNotFn gdc.PtrOperatorEvaluator
  operatorInDictionaryFn gdc.PtrOperatorEvaluator
  operatorInArrayFn gdc.PtrOperatorEvaluator
  operatorEqualPackedFloat32ArrayFn gdc.PtrOperatorEvaluator
  operatorNotEqualPackedFloat32ArrayFn gdc.PtrOperatorEvaluator
  operatorAddPackedFloat32ArrayFn gdc.PtrOperatorEvaluator
  toVariantFn gdc.TypeFromVariantConstructorFunc
  fromVariantFn gdc.VariantFromTypeConstructorFunc
}

var ptrsForPackedFloat32Array ptrsForPackedFloat32ArrayList

func initPackedFloat32ArrayPtrs(iface gdc.Interface) {
  ptrsForPackedFloat32Array.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePackedFloat32Array, 0))
  ptrsForPackedFloat32Array.ctrFromPackedFloat32ArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePackedFloat32Array, 1))
  ptrsForPackedFloat32Array.ctrFromArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePackedFloat32Array, 2))
  ptrsForPackedFloat32Array.destructorFn = ensurePtr(iface.VariantGetPtrDestructor(gdc.VariantTypePackedFloat32Array))
  {
    methodName := StringNameFromStr("size")
    defer methodName.Destroy()
    ptrsForPackedFloat32Array.methodSizeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat32Array, methodName.AsCPtr(), 3173160232))
  }
  {
    methodName := StringNameFromStr("is_empty")
    defer methodName.Destroy()
    ptrsForPackedFloat32Array.methodIsEmptyFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat32Array, methodName.AsCPtr(), 3918633141))
  }
  {
    methodName := StringNameFromStr("set")
    defer methodName.Destroy()
    ptrsForPackedFloat32Array.methodSetFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat32Array, methodName.AsCPtr(), 1113000516))
  }
  {
    methodName := StringNameFromStr("push_back")
    defer methodName.Destroy()
    ptrsForPackedFloat32Array.methodPushBackFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat32Array, methodName.AsCPtr(), 4094791666))
  }
  {
    methodName := StringNameFromStr("append")
    defer methodName.Destroy()
    ptrsForPackedFloat32Array.methodAppendFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat32Array, methodName.AsCPtr(), 4094791666))
  }
  {
    methodName := StringNameFromStr("append_array")
    defer methodName.Destroy()
    ptrsForPackedFloat32Array.methodAppendArrayFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat32Array, methodName.AsCPtr(), 2981316639))
  }
  {
    methodName := StringNameFromStr("remove_at")
    defer methodName.Destroy()
    ptrsForPackedFloat32Array.methodRemoveAtFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat32Array, methodName.AsCPtr(), 2823966027))
  }
  {
    methodName := StringNameFromStr("insert")
    defer methodName.Destroy()
    ptrsForPackedFloat32Array.methodInsertFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat32Array, methodName.AsCPtr(), 1379903876))
  }
  {
    methodName := StringNameFromStr("fill")
    defer methodName.Destroy()
    ptrsForPackedFloat32Array.methodFillFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat32Array, methodName.AsCPtr(), 833936903))
  }
  {
    methodName := StringNameFromStr("resize")
    defer methodName.Destroy()
    ptrsForPackedFloat32Array.methodResizeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat32Array, methodName.AsCPtr(), 848867239))
  }
  {
    methodName := StringNameFromStr("clear")
    defer methodName.Destroy()
    ptrsForPackedFloat32Array.methodClearFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat32Array, methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("has")
    defer methodName.Destroy()
    ptrsForPackedFloat32Array.methodHasFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat32Array, methodName.AsCPtr(), 1296369134))
  }
  {
    methodName := StringNameFromStr("reverse")
    defer methodName.Destroy()
    ptrsForPackedFloat32Array.methodReverseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat32Array, methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("slice")
    defer methodName.Destroy()
    ptrsForPackedFloat32Array.methodSliceFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat32Array, methodName.AsCPtr(), 1418229160))
  }
  {
    methodName := StringNameFromStr("to_byte_array")
    defer methodName.Destroy()
    ptrsForPackedFloat32Array.methodToByteArrayFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat32Array, methodName.AsCPtr(), 247621236))
  }
  {
    methodName := StringNameFromStr("sort")
    defer methodName.Destroy()
    ptrsForPackedFloat32Array.methodSortFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat32Array, methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("bsearch")
    defer methodName.Destroy()
    ptrsForPackedFloat32Array.methodBsearchFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat32Array, methodName.AsCPtr(), 1188816338))
  }
  {
    methodName := StringNameFromStr("duplicate")
    defer methodName.Destroy()
    ptrsForPackedFloat32Array.methodDuplicateFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat32Array, methodName.AsCPtr(), 831114784))
  }
  {
    methodName := StringNameFromStr("find")
    defer methodName.Destroy()
    ptrsForPackedFloat32Array.methodFindFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat32Array, methodName.AsCPtr(), 1343150241))
  }
  {
    methodName := StringNameFromStr("rfind")
    defer methodName.Destroy()
    ptrsForPackedFloat32Array.methodRfindFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat32Array, methodName.AsCPtr(), 1343150241))
  }
  {
    methodName := StringNameFromStr("count")
    defer methodName.Destroy()
    ptrsForPackedFloat32Array.methodCountFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat32Array, methodName.AsCPtr(), 2859915090))
  }
  ptrsForPackedFloat32Array.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypePackedFloat32Array, gdc.VariantTypeNil))
  ptrsForPackedFloat32Array.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypePackedFloat32Array, gdc.VariantTypeDictionary))
  ptrsForPackedFloat32Array.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypePackedFloat32Array, gdc.VariantTypeArray))
  ptrsForPackedFloat32Array.operatorEqualPackedFloat32ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypePackedFloat32Array, gdc.VariantTypePackedFloat32Array))
  ptrsForPackedFloat32Array.operatorNotEqualPackedFloat32ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypePackedFloat32Array, gdc.VariantTypePackedFloat32Array))
  ptrsForPackedFloat32Array.operatorAddPackedFloat32ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, gdc.VariantTypePackedFloat32Array, gdc.VariantTypePackedFloat32Array))
  ptrsForPackedFloat32Array.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypePackedFloat32Array))
  ptrsForPackedFloat32Array.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypePackedFloat32Array))
}

type PackedFloat32Array struct {
  data   *[classSizePackedFloat32Array]byte
  iface  gdc.Interface
  pinner runtime.Pinner
}

// Enums

// Constructors
func newPackedFloat32Array() *PackedFloat32Array {
  me := &PackedFloat32Array{
    data:   new([classSizePackedFloat32Array]byte),
    iface:  giface,
  }
  me.pinner.Pin(me)
  me.pinner.Pin(me.AsTypePtr())
  return me
}

func NewPackedFloat32Array() *PackedFloat32Array {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedFloat32Array()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForPackedFloat32Array.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewPackedFloat32ArrayFromPackedFloat32Array(from PackedFloat32Array, ) *PackedFloat32Array {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedFloat32Array()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForPackedFloat32Array.ctrFromPackedFloat32ArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewPackedFloat32ArrayFromArray(from Array, ) *PackedFloat32Array {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedFloat32Array()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForPackedFloat32Array.ctrFromArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

// Destructor
func (me *PackedFloat32Array) Destroy() {
	me.iface.CallPtrDestructor(ensurePtr(ptrsForPackedFloat32Array.destructorFn), me.AsTypePtr())
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsPackedFloat32Array() (*PackedFloat32Array, error) {
	if me.Type() != gdc.VariantTypePackedFloat32Array {
		return nil, fmt.Errorf("variant is not a PackedFloat32Array")
	}
  bclass := newPackedFloat32Array()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForPackedFloat32Array.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *PackedFloat32Array) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForPackedFloat32Array.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
  return va
}

// Pointers
func PackedFloat32ArrayFromPtr(ptr gdc.ConstTypePtr) *PackedFloat32Array {
  me := newPackedFloat32Array()
  dataFromPtr(me.data[:], ptr)
  return me
}

func (me *PackedFloat32Array) Type() gdc.VariantType {
  return gdc.VariantTypePackedFloat32Array
}

func (me *PackedFloat32Array) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *PackedFloat32Array) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *PackedFloat32Array) asUninitialized() gdc.UninitializedTypePtr {
  return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *PackedFloat32Array) Size() int64 {
  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat32Array.methodSizeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedFloat32Array) IsEmpty() bool {
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat32Array.methodIsEmptyFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedFloat32Array) Set(index int64, value float64, )  {
  varg0 := NewIntFromInt(index)
  defer varg0.Destroy()
  varg1 := NewFloatFromFloat32(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat32Array.methodSetFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat32Array) PushBack(value float64, ) bool {
  ret := NewBool()
  defer ret.Destroy()
  varg0 := NewFloatFromFloat32(value)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat32Array.methodPushBackFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedFloat32Array) Append(value float64, ) bool {
  ret := NewBool()
  defer ret.Destroy()
  varg0 := NewFloatFromFloat32(value)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat32Array.methodAppendFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedFloat32Array) AppendArray(array PackedFloat32Array, )  {

  args := []gdc.ConstTypePtr{array.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat32Array.methodAppendArrayFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat32Array) RemoveAt(index int64, )  {
  varg0 := NewIntFromInt(index)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat32Array.methodRemoveAtFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat32Array) Insert(at_index int64, value float64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(at_index)
  defer varg0.Destroy()
  varg1 := NewFloatFromFloat32(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat32Array.methodInsertFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedFloat32Array) Fill(value float64, )  {
  varg0 := NewFloatFromFloat32(value)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat32Array.methodFillFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat32Array) Resize(new_size int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(new_size)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat32Array.methodResizeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedFloat32Array) Clear()  {
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat32Array.methodClearFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat32Array) Has(value float64, ) bool {
  ret := NewBool()
  defer ret.Destroy()
  varg0 := NewFloatFromFloat32(value)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat32Array.methodHasFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedFloat32Array) Reverse()  {
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat32Array.methodReverseFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat32Array) Slice(begin int64, end int64, ) PackedFloat32Array {
  ret := NewPackedFloat32Array()

  varg0 := NewIntFromInt(begin)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(end)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat32Array.methodSliceFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedFloat32Array) ToByteArray() PackedByteArray {
  ret := NewPackedByteArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat32Array.methodToByteArrayFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedFloat32Array) Sort()  {
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat32Array.methodSortFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat32Array) Bsearch(value float64, before bool, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewFloatFromFloat32(value)
  defer varg0.Destroy()
  varg1 := NewBoolFromBool(before)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat32Array.methodBsearchFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedFloat32Array) Duplicate() PackedFloat32Array {
  ret := NewPackedFloat32Array()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat32Array.methodDuplicateFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedFloat32Array) Find(value float64, from int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewFloatFromFloat32(value)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(from)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat32Array.methodFindFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedFloat32Array) Rfind(value float64, from int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewFloatFromFloat32(value)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(from)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat32Array.methodRfindFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedFloat32Array) Count(value float64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewFloatFromFloat32(value)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat32Array.methodCountFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

// Operators

func (me *PackedFloat32Array) EqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedFloat32Array) NotEqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedFloat32Array) Not() bool {
  opPtr := ptrsForPackedFloat32Array.operatorNotFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedFloat32Array) InDictionary(right Dictionary) bool {
  opPtr := ptrsForPackedFloat32Array.operatorInDictionaryFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedFloat32Array) InArray(right Array) bool {
  opPtr := ptrsForPackedFloat32Array.operatorInArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedFloat32Array) EqualPackedFloat32Array(right PackedFloat32Array) bool {
  opPtr := ptrsForPackedFloat32Array.operatorEqualPackedFloat32ArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedFloat32Array) NotEqualPackedFloat32Array(right PackedFloat32Array) bool {
  opPtr := ptrsForPackedFloat32Array.operatorNotEqualPackedFloat32ArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedFloat32Array) AddPackedFloat32Array(right PackedFloat32Array) PackedFloat32Array {
  opPtr := ptrsForPackedFloat32Array.operatorAddPackedFloat32ArrayFn
  ret := NewPackedFloat32Array()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

// Members
