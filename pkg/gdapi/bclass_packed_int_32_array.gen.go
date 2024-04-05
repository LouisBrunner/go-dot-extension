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

type ptrsForPackedInt32ArrayList struct {
  ctrFn gdc.PtrConstructor
  ctrFromPackedInt32ArrayFn gdc.PtrConstructor
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
  operatorEqualPackedInt32ArrayFn gdc.PtrOperatorEvaluator
  operatorNotEqualPackedInt32ArrayFn gdc.PtrOperatorEvaluator
  operatorAddPackedInt32ArrayFn gdc.PtrOperatorEvaluator
  toVariantFn gdc.TypeFromVariantConstructorFunc
  fromVariantFn gdc.VariantFromTypeConstructorFunc
}

var ptrsForPackedInt32Array ptrsForPackedInt32ArrayList

func initPackedInt32ArrayPtrs(iface gdc.Interface) {
  ptrsForPackedInt32Array.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePackedInt32Array, 0))
  ptrsForPackedInt32Array.ctrFromPackedInt32ArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePackedInt32Array, 1))
  ptrsForPackedInt32Array.ctrFromArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePackedInt32Array, 2))
  ptrsForPackedInt32Array.destructorFn = ensurePtr(iface.VariantGetPtrDestructor(gdc.VariantTypePackedInt32Array))
  {
    methodName := StringNameFromStr("size")
    defer methodName.Destroy()
    ptrsForPackedInt32Array.methodSizeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, methodName.AsCPtr(), 3173160232))
  }
  {
    methodName := StringNameFromStr("is_empty")
    defer methodName.Destroy()
    ptrsForPackedInt32Array.methodIsEmptyFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, methodName.AsCPtr(), 3918633141))
  }
  {
    methodName := StringNameFromStr("set")
    defer methodName.Destroy()
    ptrsForPackedInt32Array.methodSetFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, methodName.AsCPtr(), 3638975848))
  }
  {
    methodName := StringNameFromStr("push_back")
    defer methodName.Destroy()
    ptrsForPackedInt32Array.methodPushBackFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, methodName.AsCPtr(), 694024632))
  }
  {
    methodName := StringNameFromStr("append")
    defer methodName.Destroy()
    ptrsForPackedInt32Array.methodAppendFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, methodName.AsCPtr(), 694024632))
  }
  {
    methodName := StringNameFromStr("append_array")
    defer methodName.Destroy()
    ptrsForPackedInt32Array.methodAppendArrayFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, methodName.AsCPtr(), 1087733270))
  }
  {
    methodName := StringNameFromStr("remove_at")
    defer methodName.Destroy()
    ptrsForPackedInt32Array.methodRemoveAtFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, methodName.AsCPtr(), 2823966027))
  }
  {
    methodName := StringNameFromStr("insert")
    defer methodName.Destroy()
    ptrsForPackedInt32Array.methodInsertFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, methodName.AsCPtr(), 1487112728))
  }
  {
    methodName := StringNameFromStr("fill")
    defer methodName.Destroy()
    ptrsForPackedInt32Array.methodFillFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, methodName.AsCPtr(), 2823966027))
  }
  {
    methodName := StringNameFromStr("resize")
    defer methodName.Destroy()
    ptrsForPackedInt32Array.methodResizeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, methodName.AsCPtr(), 848867239))
  }
  {
    methodName := StringNameFromStr("clear")
    defer methodName.Destroy()
    ptrsForPackedInt32Array.methodClearFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("has")
    defer methodName.Destroy()
    ptrsForPackedInt32Array.methodHasFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, methodName.AsCPtr(), 931488181))
  }
  {
    methodName := StringNameFromStr("reverse")
    defer methodName.Destroy()
    ptrsForPackedInt32Array.methodReverseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("slice")
    defer methodName.Destroy()
    ptrsForPackedInt32Array.methodSliceFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, methodName.AsCPtr(), 1216021098))
  }
  {
    methodName := StringNameFromStr("to_byte_array")
    defer methodName.Destroy()
    ptrsForPackedInt32Array.methodToByteArrayFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, methodName.AsCPtr(), 247621236))
  }
  {
    methodName := StringNameFromStr("sort")
    defer methodName.Destroy()
    ptrsForPackedInt32Array.methodSortFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("bsearch")
    defer methodName.Destroy()
    ptrsForPackedInt32Array.methodBsearchFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, methodName.AsCPtr(), 3380005890))
  }
  {
    methodName := StringNameFromStr("duplicate")
    defer methodName.Destroy()
    ptrsForPackedInt32Array.methodDuplicateFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, methodName.AsCPtr(), 1997843129))
  }
  {
    methodName := StringNameFromStr("find")
    defer methodName.Destroy()
    ptrsForPackedInt32Array.methodFindFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, methodName.AsCPtr(), 2984303840))
  }
  {
    methodName := StringNameFromStr("rfind")
    defer methodName.Destroy()
    ptrsForPackedInt32Array.methodRfindFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, methodName.AsCPtr(), 2984303840))
  }
  {
    methodName := StringNameFromStr("count")
    defer methodName.Destroy()
    ptrsForPackedInt32Array.methodCountFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, methodName.AsCPtr(), 4103005248))
  }
  ptrsForPackedInt32Array.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypePackedInt32Array, gdc.VariantTypeNil))
  ptrsForPackedInt32Array.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypePackedInt32Array, gdc.VariantTypeDictionary))
  ptrsForPackedInt32Array.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypePackedInt32Array, gdc.VariantTypeArray))
  ptrsForPackedInt32Array.operatorEqualPackedInt32ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypePackedInt32Array, gdc.VariantTypePackedInt32Array))
  ptrsForPackedInt32Array.operatorNotEqualPackedInt32ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypePackedInt32Array, gdc.VariantTypePackedInt32Array))
  ptrsForPackedInt32Array.operatorAddPackedInt32ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, gdc.VariantTypePackedInt32Array, gdc.VariantTypePackedInt32Array))
  ptrsForPackedInt32Array.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypePackedInt32Array))
  ptrsForPackedInt32Array.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypePackedInt32Array))
}

type PackedInt32Array struct {
  data   *[classSizePackedInt32Array]byte
  iface  gdc.Interface
  pinner runtime.Pinner
}

// Enums

// Constructors
func newPackedInt32Array() *PackedInt32Array {
  me := &PackedInt32Array{
    data:   new([classSizePackedInt32Array]byte),
    iface:  giface,
  }
  me.pinner.Pin(me)
  me.pinner.Pin(me.AsTypePtr())
  return me
}

func NewPackedInt32Array() *PackedInt32Array {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedInt32Array()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForPackedInt32Array.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewPackedInt32ArrayFromPackedInt32Array(from PackedInt32Array, ) *PackedInt32Array {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedInt32Array()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForPackedInt32Array.ctrFromPackedInt32ArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewPackedInt32ArrayFromArray(from Array, ) *PackedInt32Array {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedInt32Array()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForPackedInt32Array.ctrFromArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

// Destructor
func (me *PackedInt32Array) Destroy() {
	me.iface.CallPtrDestructor(ensurePtr(ptrsForPackedInt32Array.destructorFn), me.AsTypePtr())
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsPackedInt32Array() (*PackedInt32Array, error) {
	if me.Type() != gdc.VariantTypePackedInt32Array {
		return nil, fmt.Errorf("variant is not a PackedInt32Array")
	}
  bclass := newPackedInt32Array()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForPackedInt32Array.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *PackedInt32Array) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForPackedInt32Array.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
  return va
}

// Pointers
func PackedInt32ArrayFromPtr(ptr gdc.ConstTypePtr) *PackedInt32Array {
  me := newPackedInt32Array()
  dataFromPtr(me.data[:], ptr)
  return me
}

func (me *PackedInt32Array) Type() gdc.VariantType {
  return gdc.VariantTypePackedInt32Array
}

func (me *PackedInt32Array) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *PackedInt32Array) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *PackedInt32Array) asUninitialized() gdc.UninitializedTypePtr {
  return gdc.UninitializedTypePtr(me.AsTypePtr())
}
func (me *PackedInt32Array) Get(i int64) int32 {
  ret := me.iface.PackedInt32ArrayOperatorIndexConst(me.AsCTypePtr(), gdc.Int(i))
  return int32(*ret)
}

// Methods

func (me *PackedInt32Array) Size() int64 {
  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt32Array.methodSizeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedInt32Array) IsEmpty() bool {
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt32Array.methodIsEmptyFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedInt32Array) Set(index int64, value int64, )  {
  varg0 := NewIntFromInt(index)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt32Array.methodSetFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt32Array) PushBack(value int64, ) bool {
  ret := NewBool()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt32Array.methodPushBackFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedInt32Array) Append(value int64, ) bool {
  ret := NewBool()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt32Array.methodAppendFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedInt32Array) AppendArray(array PackedInt32Array, )  {

  args := []gdc.ConstTypePtr{array.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt32Array.methodAppendArrayFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt32Array) RemoveAt(index int64, )  {
  varg0 := NewIntFromInt(index)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt32Array.methodRemoveAtFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt32Array) Insert(at_index int64, value int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(at_index)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt32Array.methodInsertFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedInt32Array) Fill(value int64, )  {
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt32Array.methodFillFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt32Array) Resize(new_size int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(new_size)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt32Array.methodResizeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedInt32Array) Clear()  {
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt32Array.methodClearFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt32Array) Has(value int64, ) bool {
  ret := NewBool()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt32Array.methodHasFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedInt32Array) Reverse()  {
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt32Array.methodReverseFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt32Array) Slice(begin int64, end int64, ) PackedInt32Array {
  ret := NewPackedInt32Array()

  varg0 := NewIntFromInt(begin)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(end)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt32Array.methodSliceFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedInt32Array) ToByteArray() PackedByteArray {
  ret := NewPackedByteArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt32Array.methodToByteArrayFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedInt32Array) Sort()  {
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt32Array.methodSortFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt32Array) Bsearch(value int64, before bool, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  varg1 := NewBoolFromBool(before)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt32Array.methodBsearchFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedInt32Array) Duplicate() PackedInt32Array {
  ret := NewPackedInt32Array()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt32Array.methodDuplicateFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedInt32Array) Find(value int64, from int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(from)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt32Array.methodFindFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedInt32Array) Rfind(value int64, from int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(from)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt32Array.methodRfindFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedInt32Array) Count(value int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedInt32Array.methodCountFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

// Operators

func (me *PackedInt32Array) EqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedInt32Array) NotEqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedInt32Array) Not() bool {
  opPtr := ptrsForPackedInt32Array.operatorNotFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedInt32Array) InDictionary(right Dictionary) bool {
  opPtr := ptrsForPackedInt32Array.operatorInDictionaryFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedInt32Array) InArray(right Array) bool {
  opPtr := ptrsForPackedInt32Array.operatorInArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedInt32Array) EqualPackedInt32Array(right PackedInt32Array) bool {
  opPtr := ptrsForPackedInt32Array.operatorEqualPackedInt32ArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedInt32Array) NotEqualPackedInt32Array(right PackedInt32Array) bool {
  opPtr := ptrsForPackedInt32Array.operatorNotEqualPackedInt32ArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedInt32Array) AddPackedInt32Array(right PackedInt32Array) PackedInt32Array {
  opPtr := ptrsForPackedInt32Array.operatorAddPackedInt32ArrayFn
  ret := NewPackedInt32Array()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

// Members
