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

type ptrsForPackedVector2ArrayList struct {
  ctrFn gdc.PtrConstructor
  ctrFromPackedVector2ArrayFn gdc.PtrConstructor
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
  operatorMultiplyTransform2DFn gdc.PtrOperatorEvaluator
  operatorInDictionaryFn gdc.PtrOperatorEvaluator
  operatorInArrayFn gdc.PtrOperatorEvaluator
  operatorEqualPackedVector2ArrayFn gdc.PtrOperatorEvaluator
  operatorNotEqualPackedVector2ArrayFn gdc.PtrOperatorEvaluator
  operatorAddPackedVector2ArrayFn gdc.PtrOperatorEvaluator
  toVariantFn gdc.TypeFromVariantConstructorFunc
  fromVariantFn gdc.VariantFromTypeConstructorFunc
}

var ptrsForPackedVector2Array ptrsForPackedVector2ArrayList

func initPackedVector2ArrayPtrs(iface gdc.Interface) {
  ptrsForPackedVector2Array.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePackedVector2Array, 0))
  ptrsForPackedVector2Array.ctrFromPackedVector2ArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePackedVector2Array, 1))
  ptrsForPackedVector2Array.ctrFromArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePackedVector2Array, 2))
  ptrsForPackedVector2Array.destructorFn = ensurePtr(iface.VariantGetPtrDestructor(gdc.VariantTypePackedVector2Array))
  {
    methodName := StringNameFromStr("size")
    defer methodName.Destroy()
    ptrsForPackedVector2Array.methodSizeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector2Array, methodName.AsCPtr(), 3173160232))
  }
  {
    methodName := StringNameFromStr("is_empty")
    defer methodName.Destroy()
    ptrsForPackedVector2Array.methodIsEmptyFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector2Array, methodName.AsCPtr(), 3918633141))
  }
  {
    methodName := StringNameFromStr("set")
    defer methodName.Destroy()
    ptrsForPackedVector2Array.methodSetFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector2Array, methodName.AsCPtr(), 635767250))
  }
  {
    methodName := StringNameFromStr("push_back")
    defer methodName.Destroy()
    ptrsForPackedVector2Array.methodPushBackFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector2Array, methodName.AsCPtr(), 4188891560))
  }
  {
    methodName := StringNameFromStr("append")
    defer methodName.Destroy()
    ptrsForPackedVector2Array.methodAppendFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector2Array, methodName.AsCPtr(), 4188891560))
  }
  {
    methodName := StringNameFromStr("append_array")
    defer methodName.Destroy()
    ptrsForPackedVector2Array.methodAppendArrayFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector2Array, methodName.AsCPtr(), 3887534835))
  }
  {
    methodName := StringNameFromStr("remove_at")
    defer methodName.Destroy()
    ptrsForPackedVector2Array.methodRemoveAtFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector2Array, methodName.AsCPtr(), 2823966027))
  }
  {
    methodName := StringNameFromStr("insert")
    defer methodName.Destroy()
    ptrsForPackedVector2Array.methodInsertFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector2Array, methodName.AsCPtr(), 2225629369))
  }
  {
    methodName := StringNameFromStr("fill")
    defer methodName.Destroy()
    ptrsForPackedVector2Array.methodFillFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector2Array, methodName.AsCPtr(), 3790411178))
  }
  {
    methodName := StringNameFromStr("resize")
    defer methodName.Destroy()
    ptrsForPackedVector2Array.methodResizeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector2Array, methodName.AsCPtr(), 848867239))
  }
  {
    methodName := StringNameFromStr("clear")
    defer methodName.Destroy()
    ptrsForPackedVector2Array.methodClearFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector2Array, methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("has")
    defer methodName.Destroy()
    ptrsForPackedVector2Array.methodHasFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector2Array, methodName.AsCPtr(), 3190634762))
  }
  {
    methodName := StringNameFromStr("reverse")
    defer methodName.Destroy()
    ptrsForPackedVector2Array.methodReverseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector2Array, methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("slice")
    defer methodName.Destroy()
    ptrsForPackedVector2Array.methodSliceFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector2Array, methodName.AsCPtr(), 3864005350))
  }
  {
    methodName := StringNameFromStr("to_byte_array")
    defer methodName.Destroy()
    ptrsForPackedVector2Array.methodToByteArrayFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector2Array, methodName.AsCPtr(), 247621236))
  }
  {
    methodName := StringNameFromStr("sort")
    defer methodName.Destroy()
    ptrsForPackedVector2Array.methodSortFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector2Array, methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("bsearch")
    defer methodName.Destroy()
    ptrsForPackedVector2Array.methodBsearchFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector2Array, methodName.AsCPtr(), 3778035805))
  }
  {
    methodName := StringNameFromStr("duplicate")
    defer methodName.Destroy()
    ptrsForPackedVector2Array.methodDuplicateFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector2Array, methodName.AsCPtr(), 3763646812))
  }
  {
    methodName := StringNameFromStr("find")
    defer methodName.Destroy()
    ptrsForPackedVector2Array.methodFindFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector2Array, methodName.AsCPtr(), 1469606149))
  }
  {
    methodName := StringNameFromStr("rfind")
    defer methodName.Destroy()
    ptrsForPackedVector2Array.methodRfindFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector2Array, methodName.AsCPtr(), 1469606149))
  }
  {
    methodName := StringNameFromStr("count")
    defer methodName.Destroy()
    ptrsForPackedVector2Array.methodCountFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector2Array, methodName.AsCPtr(), 2798848307))
  }
  ptrsForPackedVector2Array.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypePackedVector2Array, gdc.VariantTypeNil))
  ptrsForPackedVector2Array.operatorMultiplyTransform2DFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypePackedVector2Array, gdc.VariantTypeTransform2D))
  ptrsForPackedVector2Array.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypePackedVector2Array, gdc.VariantTypeDictionary))
  ptrsForPackedVector2Array.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypePackedVector2Array, gdc.VariantTypeArray))
  ptrsForPackedVector2Array.operatorEqualPackedVector2ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypePackedVector2Array, gdc.VariantTypePackedVector2Array))
  ptrsForPackedVector2Array.operatorNotEqualPackedVector2ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypePackedVector2Array, gdc.VariantTypePackedVector2Array))
  ptrsForPackedVector2Array.operatorAddPackedVector2ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, gdc.VariantTypePackedVector2Array, gdc.VariantTypePackedVector2Array))
  ptrsForPackedVector2Array.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypePackedVector2Array))
  ptrsForPackedVector2Array.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypePackedVector2Array))
}

type PackedVector2Array struct {
  data   *[classSizePackedVector2Array]byte
  iface  gdc.Interface
  pinner runtime.Pinner
}

// Enums

// Constructors
func newPackedVector2Array() *PackedVector2Array {
  me := &PackedVector2Array{
    data:   new([classSizePackedVector2Array]byte),
    iface:  giface,
  }
  me.pinner.Pin(me)
  me.pinner.Pin(me.AsTypePtr())
  return me
}

func NewPackedVector2Array() *PackedVector2Array {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedVector2Array()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForPackedVector2Array.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewPackedVector2ArrayFromPackedVector2Array(from PackedVector2Array, ) *PackedVector2Array {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedVector2Array()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForPackedVector2Array.ctrFromPackedVector2ArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewPackedVector2ArrayFromArray(from Array, ) *PackedVector2Array {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedVector2Array()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForPackedVector2Array.ctrFromArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

// Destructor
func (me *PackedVector2Array) Destroy() {
	me.iface.CallPtrDestructor(ensurePtr(ptrsForPackedVector2Array.destructorFn), me.AsTypePtr())
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsPackedVector2Array() (*PackedVector2Array, error) {
	if me.Type() != gdc.VariantTypePackedVector2Array {
		return nil, fmt.Errorf("variant is not a PackedVector2Array")
	}
  bclass := newPackedVector2Array()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForPackedVector2Array.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *PackedVector2Array) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForPackedVector2Array.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
  return va
}

// Pointers
func PackedVector2ArrayFromPtr(ptr gdc.ConstTypePtr) *PackedVector2Array {
  me := newPackedVector2Array()
  dataFromPtr(me.data[:], ptr)
  return me
}

func (me *PackedVector2Array) Type() gdc.VariantType {
  return gdc.VariantTypePackedVector2Array
}

func (me *PackedVector2Array) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *PackedVector2Array) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *PackedVector2Array) asUninitialized() gdc.UninitializedTypePtr {
  return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *PackedVector2Array) Size() int64 {
  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector2Array.methodSizeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedVector2Array) IsEmpty() bool {
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector2Array.methodIsEmptyFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedVector2Array) Set(index int64, value Vector2, )  {
  varg0 := NewIntFromInt(index)
  defer varg0.Destroy()

  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), value.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector2Array.methodSetFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedVector2Array) PushBack(value Vector2, ) bool {
  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{value.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector2Array.methodPushBackFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedVector2Array) Append(value Vector2, ) bool {
  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{value.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector2Array.methodAppendFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedVector2Array) AppendArray(array PackedVector2Array, )  {

  args := []gdc.ConstTypePtr{array.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector2Array.methodAppendArrayFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedVector2Array) RemoveAt(index int64, )  {
  varg0 := NewIntFromInt(index)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector2Array.methodRemoveAtFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedVector2Array) Insert(at_index int64, value Vector2, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(at_index)
  defer varg0.Destroy()

  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), value.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector2Array.methodInsertFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedVector2Array) Fill(value Vector2, )  {

  args := []gdc.ConstTypePtr{value.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector2Array.methodFillFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedVector2Array) Resize(new_size int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(new_size)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector2Array.methodResizeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedVector2Array) Clear()  {
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector2Array.methodClearFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedVector2Array) Has(value Vector2, ) bool {
  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{value.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector2Array.methodHasFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedVector2Array) Reverse()  {
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector2Array.methodReverseFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedVector2Array) Slice(begin int64, end int64, ) PackedVector2Array {
  ret := NewPackedVector2Array()

  varg0 := NewIntFromInt(begin)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(end)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector2Array.methodSliceFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedVector2Array) ToByteArray() PackedByteArray {
  ret := NewPackedByteArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector2Array.methodToByteArrayFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedVector2Array) Sort()  {
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector2Array.methodSortFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedVector2Array) Bsearch(value Vector2, before bool, ) int64 {
  ret := NewInt()
  defer ret.Destroy()

  varg1 := NewBoolFromBool(before)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{value.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector2Array.methodBsearchFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedVector2Array) Duplicate() PackedVector2Array {
  ret := NewPackedVector2Array()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector2Array.methodDuplicateFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedVector2Array) Find(value Vector2, from int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()

  varg1 := NewIntFromInt(from)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{value.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector2Array.methodFindFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedVector2Array) Rfind(value Vector2, from int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()

  varg1 := NewIntFromInt(from)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{value.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector2Array.methodRfindFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedVector2Array) Count(value Vector2, ) int64 {
  ret := NewInt()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{value.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector2Array.methodCountFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

// Operators

func (me *PackedVector2Array) EqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedVector2Array) NotEqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedVector2Array) Not() bool {
  opPtr := ptrsForPackedVector2Array.operatorNotFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedVector2Array) MultiplyTransform2D(right Transform2D) PackedVector2Array {
  opPtr := ptrsForPackedVector2Array.operatorMultiplyTransform2DFn
  ret := NewPackedVector2Array()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *PackedVector2Array) InDictionary(right Dictionary) bool {
  opPtr := ptrsForPackedVector2Array.operatorInDictionaryFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedVector2Array) InArray(right Array) bool {
  opPtr := ptrsForPackedVector2Array.operatorInArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedVector2Array) EqualPackedVector2Array(right PackedVector2Array) bool {
  opPtr := ptrsForPackedVector2Array.operatorEqualPackedVector2ArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedVector2Array) NotEqualPackedVector2Array(right PackedVector2Array) bool {
  opPtr := ptrsForPackedVector2Array.operatorNotEqualPackedVector2ArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedVector2Array) AddPackedVector2Array(right PackedVector2Array) PackedVector2Array {
  opPtr := ptrsForPackedVector2Array.operatorAddPackedVector2ArrayFn
  ret := NewPackedVector2Array()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

// Members
