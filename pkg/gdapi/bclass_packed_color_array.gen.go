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

type ptrsForPackedColorArrayList struct {
  ctrFn gdc.PtrConstructor
  ctrFromPackedColorArrayFn gdc.PtrConstructor
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
  operatorEqualPackedColorArrayFn gdc.PtrOperatorEvaluator
  operatorNotEqualPackedColorArrayFn gdc.PtrOperatorEvaluator
  operatorAddPackedColorArrayFn gdc.PtrOperatorEvaluator
  toVariantFn gdc.TypeFromVariantConstructorFunc
  fromVariantFn gdc.VariantFromTypeConstructorFunc
}

var ptrsForPackedColorArray ptrsForPackedColorArrayList

func initPackedColorArrayPtrs(iface gdc.Interface) {
  ptrsForPackedColorArray.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePackedColorArray, 0))
  ptrsForPackedColorArray.ctrFromPackedColorArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePackedColorArray, 1))
  ptrsForPackedColorArray.ctrFromArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePackedColorArray, 2))
  ptrsForPackedColorArray.destructorFn = ensurePtr(iface.VariantGetPtrDestructor(gdc.VariantTypePackedColorArray))
  {
    methodName := StringNameFromStr("size")
    defer methodName.Destroy()
    ptrsForPackedColorArray.methodSizeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, methodName.AsCPtr(), 3173160232))
  }
  {
    methodName := StringNameFromStr("is_empty")
    defer methodName.Destroy()
    ptrsForPackedColorArray.methodIsEmptyFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, methodName.AsCPtr(), 3918633141))
  }
  {
    methodName := StringNameFromStr("set")
    defer methodName.Destroy()
    ptrsForPackedColorArray.methodSetFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, methodName.AsCPtr(), 1444096570))
  }
  {
    methodName := StringNameFromStr("push_back")
    defer methodName.Destroy()
    ptrsForPackedColorArray.methodPushBackFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, methodName.AsCPtr(), 1007858200))
  }
  {
    methodName := StringNameFromStr("append")
    defer methodName.Destroy()
    ptrsForPackedColorArray.methodAppendFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, methodName.AsCPtr(), 1007858200))
  }
  {
    methodName := StringNameFromStr("append_array")
    defer methodName.Destroy()
    ptrsForPackedColorArray.methodAppendArrayFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, methodName.AsCPtr(), 798822497))
  }
  {
    methodName := StringNameFromStr("remove_at")
    defer methodName.Destroy()
    ptrsForPackedColorArray.methodRemoveAtFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, methodName.AsCPtr(), 2823966027))
  }
  {
    methodName := StringNameFromStr("insert")
    defer methodName.Destroy()
    ptrsForPackedColorArray.methodInsertFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, methodName.AsCPtr(), 785289703))
  }
  {
    methodName := StringNameFromStr("fill")
    defer methodName.Destroy()
    ptrsForPackedColorArray.methodFillFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, methodName.AsCPtr(), 3730314301))
  }
  {
    methodName := StringNameFromStr("resize")
    defer methodName.Destroy()
    ptrsForPackedColorArray.methodResizeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, methodName.AsCPtr(), 848867239))
  }
  {
    methodName := StringNameFromStr("clear")
    defer methodName.Destroy()
    ptrsForPackedColorArray.methodClearFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("has")
    defer methodName.Destroy()
    ptrsForPackedColorArray.methodHasFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, methodName.AsCPtr(), 3167426256))
  }
  {
    methodName := StringNameFromStr("reverse")
    defer methodName.Destroy()
    ptrsForPackedColorArray.methodReverseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("slice")
    defer methodName.Destroy()
    ptrsForPackedColorArray.methodSliceFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, methodName.AsCPtr(), 2451797139))
  }
  {
    methodName := StringNameFromStr("to_byte_array")
    defer methodName.Destroy()
    ptrsForPackedColorArray.methodToByteArrayFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, methodName.AsCPtr(), 247621236))
  }
  {
    methodName := StringNameFromStr("sort")
    defer methodName.Destroy()
    ptrsForPackedColorArray.methodSortFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("bsearch")
    defer methodName.Destroy()
    ptrsForPackedColorArray.methodBsearchFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, methodName.AsCPtr(), 314143821))
  }
  {
    methodName := StringNameFromStr("duplicate")
    defer methodName.Destroy()
    ptrsForPackedColorArray.methodDuplicateFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, methodName.AsCPtr(), 1011903421))
  }
  {
    methodName := StringNameFromStr("find")
    defer methodName.Destroy()
    ptrsForPackedColorArray.methodFindFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, methodName.AsCPtr(), 3156095363))
  }
  {
    methodName := StringNameFromStr("rfind")
    defer methodName.Destroy()
    ptrsForPackedColorArray.methodRfindFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, methodName.AsCPtr(), 3156095363))
  }
  {
    methodName := StringNameFromStr("count")
    defer methodName.Destroy()
    ptrsForPackedColorArray.methodCountFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, methodName.AsCPtr(), 1682108616))
  }
  ptrsForPackedColorArray.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypePackedColorArray, gdc.VariantTypeNil))
  ptrsForPackedColorArray.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypePackedColorArray, gdc.VariantTypeDictionary))
  ptrsForPackedColorArray.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypePackedColorArray, gdc.VariantTypeArray))
  ptrsForPackedColorArray.operatorEqualPackedColorArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypePackedColorArray, gdc.VariantTypePackedColorArray))
  ptrsForPackedColorArray.operatorNotEqualPackedColorArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypePackedColorArray, gdc.VariantTypePackedColorArray))
  ptrsForPackedColorArray.operatorAddPackedColorArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, gdc.VariantTypePackedColorArray, gdc.VariantTypePackedColorArray))
  ptrsForPackedColorArray.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypePackedColorArray))
  ptrsForPackedColorArray.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypePackedColorArray))
}

type PackedColorArray struct {
  data   *[classSizePackedColorArray]byte
  iface  gdc.Interface
  pinner runtime.Pinner
}

// Enums

// Constructors
func newPackedColorArray() *PackedColorArray {
  me := &PackedColorArray{
    data:   new([classSizePackedColorArray]byte),
    iface:  giface,
  }
  me.pinner.Pin(me)
  me.pinner.Pin(me.AsTypePtr())
  return me
}

func NewPackedColorArray() *PackedColorArray {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedColorArray()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForPackedColorArray.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewPackedColorArrayFromPackedColorArray(from PackedColorArray, ) *PackedColorArray {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedColorArray()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForPackedColorArray.ctrFromPackedColorArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewPackedColorArrayFromArray(from Array, ) *PackedColorArray {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedColorArray()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForPackedColorArray.ctrFromArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

// Destructor
func (me *PackedColorArray) Destroy() {
	me.iface.CallPtrDestructor(ensurePtr(ptrsForPackedColorArray.destructorFn), me.AsTypePtr())
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsPackedColorArray() (*PackedColorArray, error) {
	if me.Type() != gdc.VariantTypePackedColorArray {
		return nil, fmt.Errorf("variant is not a PackedColorArray")
	}
  bclass := newPackedColorArray()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForPackedColorArray.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *PackedColorArray) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForPackedColorArray.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
  return va
}

// Pointers
func PackedColorArrayFromPtr(ptr gdc.ConstTypePtr) *PackedColorArray {
  me := newPackedColorArray()
  dataFromPtr(me.data[:], ptr)
  return me
}

func (me *PackedColorArray) Type() gdc.VariantType {
  return gdc.VariantTypePackedColorArray
}

func (me *PackedColorArray) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *PackedColorArray) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *PackedColorArray) asUninitialized() gdc.UninitializedTypePtr {
  return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *PackedColorArray) Size() int64 {
  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedColorArray.methodSizeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedColorArray) IsEmpty() bool {
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedColorArray.methodIsEmptyFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedColorArray) Set(index int64, value Color, )  {
  varg0 := NewIntFromInt(index)
  defer varg0.Destroy()

  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), value.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedColorArray.methodSetFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedColorArray) PushBack(value Color, ) bool {
  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{value.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedColorArray.methodPushBackFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedColorArray) Append(value Color, ) bool {
  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{value.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedColorArray.methodAppendFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedColorArray) AppendArray(array PackedColorArray, )  {

  args := []gdc.ConstTypePtr{array.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedColorArray.methodAppendArrayFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedColorArray) RemoveAt(index int64, )  {
  varg0 := NewIntFromInt(index)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedColorArray.methodRemoveAtFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedColorArray) Insert(at_index int64, value Color, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(at_index)
  defer varg0.Destroy()

  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), value.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedColorArray.methodInsertFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedColorArray) Fill(value Color, )  {

  args := []gdc.ConstTypePtr{value.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedColorArray.methodFillFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedColorArray) Resize(new_size int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(new_size)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedColorArray.methodResizeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedColorArray) Clear()  {
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedColorArray.methodClearFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedColorArray) Has(value Color, ) bool {
  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{value.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedColorArray.methodHasFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedColorArray) Reverse()  {
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedColorArray.methodReverseFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedColorArray) Slice(begin int64, end int64, ) PackedColorArray {
  ret := NewPackedColorArray()

  varg0 := NewIntFromInt(begin)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(end)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedColorArray.methodSliceFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedColorArray) ToByteArray() PackedByteArray {
  ret := NewPackedByteArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedColorArray.methodToByteArrayFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedColorArray) Sort()  {
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedColorArray.methodSortFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedColorArray) Bsearch(value Color, before bool, ) int64 {
  ret := NewInt()
  defer ret.Destroy()

  varg1 := NewBoolFromBool(before)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{value.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedColorArray.methodBsearchFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedColorArray) Duplicate() PackedColorArray {
  ret := NewPackedColorArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedColorArray.methodDuplicateFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedColorArray) Find(value Color, from int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()

  varg1 := NewIntFromInt(from)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{value.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedColorArray.methodFindFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedColorArray) Rfind(value Color, from int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()

  varg1 := NewIntFromInt(from)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{value.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedColorArray.methodRfindFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedColorArray) Count(value Color, ) int64 {
  ret := NewInt()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{value.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedColorArray.methodCountFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

// Operators

func (me *PackedColorArray) EqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedColorArray) NotEqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedColorArray) Not() bool {
  opPtr := ptrsForPackedColorArray.operatorNotFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedColorArray) InDictionary(right Dictionary) bool {
  opPtr := ptrsForPackedColorArray.operatorInDictionaryFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedColorArray) InArray(right Array) bool {
  opPtr := ptrsForPackedColorArray.operatorInArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedColorArray) EqualPackedColorArray(right PackedColorArray) bool {
  opPtr := ptrsForPackedColorArray.operatorEqualPackedColorArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedColorArray) NotEqualPackedColorArray(right PackedColorArray) bool {
  opPtr := ptrsForPackedColorArray.operatorNotEqualPackedColorArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedColorArray) AddPackedColorArray(right PackedColorArray) PackedColorArray {
  opPtr := ptrsForPackedColorArray.operatorAddPackedColorArrayFn
  ret := NewPackedColorArray()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

// Members
