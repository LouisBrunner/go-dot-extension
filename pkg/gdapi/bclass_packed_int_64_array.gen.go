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

type ptrsForPackedInt64ArrayList struct {
  ctrFn gdc.PtrConstructor
  ctrFromPackedInt64ArrayFn gdc.PtrConstructor
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
  operatorEqualPackedInt64ArrayFn gdc.PtrOperatorEvaluator
  operatorNotEqualPackedInt64ArrayFn gdc.PtrOperatorEvaluator
  operatorAddPackedInt64ArrayFn gdc.PtrOperatorEvaluator
  toVariantFn gdc.TypeFromVariantConstructorFunc
  fromVariantFn gdc.VariantFromTypeConstructorFunc
}

var ptrsForPackedInt64Array ptrsForPackedInt64ArrayList

func initPackedInt64ArrayPtrs(iface gdc.Interface) {
  ptrsForPackedInt64Array.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePackedInt64Array, 0))
  ptrsForPackedInt64Array.ctrFromPackedInt64ArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePackedInt64Array, 1))
  ptrsForPackedInt64Array.ctrFromArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePackedInt64Array, 2))
  ptrsForPackedInt64Array.destructorFn = ensurePtr(iface.VariantGetPtrDestructor(gdc.VariantTypePackedInt64Array))
  {
    methodName := StringNameFromStr("size")
    defer methodName.Destroy()
    ptrsForPackedInt64Array.methodSizeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, methodName.AsCPtr(), 3173160232))
  }
  {
    methodName := StringNameFromStr("is_empty")
    defer methodName.Destroy()
    ptrsForPackedInt64Array.methodIsEmptyFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, methodName.AsCPtr(), 3918633141))
  }
  {
    methodName := StringNameFromStr("set")
    defer methodName.Destroy()
    ptrsForPackedInt64Array.methodSetFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, methodName.AsCPtr(), 3638975848))
  }
  {
    methodName := StringNameFromStr("push_back")
    defer methodName.Destroy()
    ptrsForPackedInt64Array.methodPushBackFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, methodName.AsCPtr(), 694024632))
  }
  {
    methodName := StringNameFromStr("append")
    defer methodName.Destroy()
    ptrsForPackedInt64Array.methodAppendFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, methodName.AsCPtr(), 694024632))
  }
  {
    methodName := StringNameFromStr("append_array")
    defer methodName.Destroy()
    ptrsForPackedInt64Array.methodAppendArrayFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, methodName.AsCPtr(), 2090311302))
  }
  {
    methodName := StringNameFromStr("remove_at")
    defer methodName.Destroy()
    ptrsForPackedInt64Array.methodRemoveAtFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, methodName.AsCPtr(), 2823966027))
  }
  {
    methodName := StringNameFromStr("insert")
    defer methodName.Destroy()
    ptrsForPackedInt64Array.methodInsertFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, methodName.AsCPtr(), 1487112728))
  }
  {
    methodName := StringNameFromStr("fill")
    defer methodName.Destroy()
    ptrsForPackedInt64Array.methodFillFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, methodName.AsCPtr(), 2823966027))
  }
  {
    methodName := StringNameFromStr("resize")
    defer methodName.Destroy()
    ptrsForPackedInt64Array.methodResizeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, methodName.AsCPtr(), 848867239))
  }
  {
    methodName := StringNameFromStr("clear")
    defer methodName.Destroy()
    ptrsForPackedInt64Array.methodClearFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("has")
    defer methodName.Destroy()
    ptrsForPackedInt64Array.methodHasFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, methodName.AsCPtr(), 931488181))
  }
  {
    methodName := StringNameFromStr("reverse")
    defer methodName.Destroy()
    ptrsForPackedInt64Array.methodReverseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("slice")
    defer methodName.Destroy()
    ptrsForPackedInt64Array.methodSliceFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, methodName.AsCPtr(), 1726550804))
  }
  {
    methodName := StringNameFromStr("to_byte_array")
    defer methodName.Destroy()
    ptrsForPackedInt64Array.methodToByteArrayFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, methodName.AsCPtr(), 247621236))
  }
  {
    methodName := StringNameFromStr("sort")
    defer methodName.Destroy()
    ptrsForPackedInt64Array.methodSortFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("bsearch")
    defer methodName.Destroy()
    ptrsForPackedInt64Array.methodBsearchFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, methodName.AsCPtr(), 3380005890))
  }
  {
    methodName := StringNameFromStr("duplicate")
    defer methodName.Destroy()
    ptrsForPackedInt64Array.methodDuplicateFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, methodName.AsCPtr(), 2376370016))
  }
  {
    methodName := StringNameFromStr("find")
    defer methodName.Destroy()
    ptrsForPackedInt64Array.methodFindFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, methodName.AsCPtr(), 2984303840))
  }
  {
    methodName := StringNameFromStr("rfind")
    defer methodName.Destroy()
    ptrsForPackedInt64Array.methodRfindFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, methodName.AsCPtr(), 2984303840))
  }
  {
    methodName := StringNameFromStr("count")
    defer methodName.Destroy()
    ptrsForPackedInt64Array.methodCountFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, methodName.AsCPtr(), 4103005248))
  }
  ptrsForPackedInt64Array.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypePackedInt64Array, gdc.VariantTypeNil))
  ptrsForPackedInt64Array.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypePackedInt64Array, gdc.VariantTypeDictionary))
  ptrsForPackedInt64Array.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypePackedInt64Array, gdc.VariantTypeArray))
  ptrsForPackedInt64Array.operatorEqualPackedInt64ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypePackedInt64Array, gdc.VariantTypePackedInt64Array))
  ptrsForPackedInt64Array.operatorNotEqualPackedInt64ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypePackedInt64Array, gdc.VariantTypePackedInt64Array))
  ptrsForPackedInt64Array.operatorAddPackedInt64ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, gdc.VariantTypePackedInt64Array, gdc.VariantTypePackedInt64Array))
  ptrsForPackedInt64Array.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypePackedInt64Array))
  ptrsForPackedInt64Array.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypePackedInt64Array))
}

type PackedInt64Array struct {
  data   *[classSizePackedInt64Array]byte
  iface  gdc.Interface
  pinner runtime.Pinner
}

// Enums

// Constructors
func newPackedInt64Array() *PackedInt64Array {
  me := &PackedInt64Array{
    data:   new([classSizePackedInt64Array]byte),
    iface:  giface,
  }
  me.pinner.Pin(me)
  me.pinner.Pin(me.AsTypePtr())
  return me
}

func NewPackedInt64Array() *PackedInt64Array {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedInt64Array()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForPackedInt64Array.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewPackedInt64ArrayFromPackedInt64Array(from PackedInt64Array, ) *PackedInt64Array {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedInt64Array()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForPackedInt64Array.ctrFromPackedInt64ArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewPackedInt64ArrayFromArray(from Array, ) *PackedInt64Array {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedInt64Array()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForPackedInt64Array.ctrFromArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

// Destructor
func (me *PackedInt64Array) Destroy() {
	me.iface.CallPtrDestructor(ensurePtr(ptrsForPackedInt64Array.destructorFn), me.AsTypePtr())
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsPackedInt64Array() (*PackedInt64Array, error) {
	if me.Type() != gdc.VariantTypePackedInt64Array {
		return nil, fmt.Errorf("variant is not a PackedInt64Array")
	}
  bclass := newPackedInt64Array()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForPackedInt64Array.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *PackedInt64Array) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForPackedInt64Array.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
  return va
}

// Pointers
func PackedInt64ArrayFromPtr(ptr gdc.ConstTypePtr) *PackedInt64Array {
  me := newPackedInt64Array()
  dataFromPtr(me.data[:], ptr)
  return me
}

func (me *PackedInt64Array) Type() gdc.VariantType {
  return gdc.VariantTypePackedInt64Array
}

func (me *PackedInt64Array) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *PackedInt64Array) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *PackedInt64Array) asUninitialized() gdc.UninitializedTypePtr {
  return gdc.UninitializedTypePtr(me.AsTypePtr())
}
func (me *PackedInt64Array) Get(i int64) int64 {
  ret := me.iface.PackedInt64ArrayOperatorIndexConst(me.AsCTypePtr(), gdc.Int(i))
  return *ret
}

// Methods

func (me *PackedInt64Array) Size() int64 {
  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt64Array.methodSizeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedInt64Array) IsEmpty() bool {
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt64Array.methodIsEmptyFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedInt64Array) Set(index int64, value int64, )  {
  varg0 := NewIntFromInt(index)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt64Array.methodSetFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt64Array) PushBack(value int64, ) bool {
  ret := NewBool()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt64Array.methodPushBackFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedInt64Array) Append(value int64, ) bool {
  ret := NewBool()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt64Array.methodAppendFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedInt64Array) AppendArray(array PackedInt64Array, )  {

  args := []gdc.ConstTypePtr{array.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt64Array.methodAppendArrayFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt64Array) RemoveAt(index int64, )  {
  varg0 := NewIntFromInt(index)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt64Array.methodRemoveAtFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt64Array) Insert(at_index int64, value int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(at_index)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt64Array.methodInsertFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedInt64Array) Fill(value int64, )  {
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt64Array.methodFillFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt64Array) Resize(new_size int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(new_size)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt64Array.methodResizeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedInt64Array) Clear()  {
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt64Array.methodClearFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt64Array) Has(value int64, ) bool {
  ret := NewBool()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt64Array.methodHasFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedInt64Array) Reverse()  {
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt64Array.methodReverseFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt64Array) Slice(begin int64, end int64, ) PackedInt64Array {
  ret := NewPackedInt64Array()

  varg0 := NewIntFromInt(begin)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(end)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt64Array.methodSliceFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedInt64Array) ToByteArray() PackedByteArray {
  ret := NewPackedByteArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt64Array.methodToByteArrayFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedInt64Array) Sort()  {
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt64Array.methodSortFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt64Array) Bsearch(value int64, before bool, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  varg1 := NewBoolFromBool(before)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt64Array.methodBsearchFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedInt64Array) Duplicate() PackedInt64Array {
  ret := NewPackedInt64Array()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt64Array.methodDuplicateFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedInt64Array) Find(value int64, from int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(from)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt64Array.methodFindFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedInt64Array) Rfind(value int64, from int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(from)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt64Array.methodRfindFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedInt64Array) Count(value int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt64Array.methodCountFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

// Operators

func (me *PackedInt64Array) EqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedInt64Array) NotEqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedInt64Array) Not() bool {
  opPtr := ptrsForPackedInt64Array.operatorNotFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedInt64Array) InDictionary(right Dictionary) bool {
  opPtr := ptrsForPackedInt64Array.operatorInDictionaryFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedInt64Array) InArray(right Array) bool {
  opPtr := ptrsForPackedInt64Array.operatorInArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedInt64Array) EqualPackedInt64Array(right PackedInt64Array) bool {
  opPtr := ptrsForPackedInt64Array.operatorEqualPackedInt64ArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedInt64Array) NotEqualPackedInt64Array(right PackedInt64Array) bool {
  opPtr := ptrsForPackedInt64Array.operatorNotEqualPackedInt64ArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedInt64Array) AddPackedInt64Array(right PackedInt64Array) PackedInt64Array {
  opPtr := ptrsForPackedInt64Array.operatorAddPackedInt64ArrayFn
  ret := NewPackedInt64Array()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

// Members
