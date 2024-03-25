// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "fmt"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypePackedColorArray, 0) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewPackedColorArrayFromPackedColorArray(from PackedColorArray, ) *PackedColorArray {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedColorArray()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypePackedColorArray, 1) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewPackedColorArrayFromArray(from Array, ) *PackedColorArray {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedColorArray()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypePackedColorArray, 2) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

// Destructor
func (me *PackedColorArray) Destroy() {
  dtr := me.iface.VariantGetPtrDestructor(gdc.VariantTypePackedColorArray)
	me.iface.CallPtrDestructor(dtr, me.AsTypePtr())
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsPackedColorArray() (*PackedColorArray, error) {
	if me.Type() != gdc.VariantTypePackedColorArray {
		return nil, fmt.Errorf("variant is not a PackedColorArray")
	}
  bclass := newPackedColorArray()
	fn := me.iface.GetVariantToTypeConstructor(me.Type())
	me.iface.CallTypeFromVariantConstructorFunc(fn, bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *PackedColorArray) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  fn := me.iface.GetVariantFromTypeConstructor(me.Type())
  me.iface.CallVariantFromTypeConstructorFunc(fn, va.asUninitialized(), me.AsTypePtr())
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

func (me *PackedColorArray) Size() int {
  name := StringNameFromStr("size")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, name.AsCPtr(), 3173160232) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret int
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedColorArray) IsEmpty() bool {
  name := StringNameFromStr("is_empty")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, name.AsCPtr(), 3918633141) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret bool
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedColorArray) Set(index int, value Color, )  {
  name := StringNameFromStr("set")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, name.AsCPtr(), 1444096570) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&index)), value.AsCTypePtr(), }

  pinner.Pin(&index)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedColorArray) PushBack(value Color, ) bool {
  name := StringNameFromStr("push_back")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, name.AsCPtr(), 1007858200) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret bool
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{value.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedColorArray) Append(value Color, ) bool {
  name := StringNameFromStr("append")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, name.AsCPtr(), 1007858200) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret bool
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{value.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedColorArray) AppendArray(array PackedColorArray, )  {
  name := StringNameFromStr("append_array")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, name.AsCPtr(), 798822497) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  args := []gdc.ConstTypePtr{array.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedColorArray) RemoveAt(index int, )  {
  name := StringNameFromStr("remove_at")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, name.AsCPtr(), 2823966027) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&index)), }

  pinner.Pin(&index)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedColorArray) Insert(at_index int, value Color, ) int {
  name := StringNameFromStr("insert")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, name.AsCPtr(), 785289703) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret int
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&at_index)), value.AsCTypePtr(), }

  pinner.Pin(&at_index)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedColorArray) Fill(value Color, )  {
  name := StringNameFromStr("fill")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, name.AsCPtr(), 3730314301) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  args := []gdc.ConstTypePtr{value.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedColorArray) Resize(new_size int, ) int {
  name := StringNameFromStr("resize")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, name.AsCPtr(), 848867239) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret int
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&new_size)), }

  pinner.Pin(&new_size)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedColorArray) Clear()  {
  name := StringNameFromStr("clear")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, name.AsCPtr(), 3218959716) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedColorArray) Has(value Color, ) bool {
  name := StringNameFromStr("has")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, name.AsCPtr(), 3167426256) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret bool
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{value.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedColorArray) Reverse()  {
  name := StringNameFromStr("reverse")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, name.AsCPtr(), 3218959716) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedColorArray) Slice(begin int, end int, ) PackedColorArray {
  name := StringNameFromStr("slice")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, name.AsCPtr(), 2451797139) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret PackedColorArray
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&begin)), gdc.ConstTypePtr(unsafe.Pointer(&end)), }

  pinner.Pin(&begin)
  pinner.Pin(&end)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedColorArray) ToByteArray() PackedByteArray {
  name := StringNameFromStr("to_byte_array")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, name.AsCPtr(), 247621236) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret PackedByteArray
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedColorArray) Sort()  {
  name := StringNameFromStr("sort")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, name.AsCPtr(), 3218959716) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedColorArray) Bsearch(value Color, before bool, ) int {
  name := StringNameFromStr("bsearch")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, name.AsCPtr(), 314143821) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret int
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{value.AsCTypePtr(), gdc.ConstTypePtr(unsafe.Pointer(&before)), }

  pinner.Pin(&before)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedColorArray) Duplicate() PackedColorArray {
  name := StringNameFromStr("duplicate")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, name.AsCPtr(), 1011903421) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret PackedColorArray
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedColorArray) Find(value Color, from int, ) int {
  name := StringNameFromStr("find")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, name.AsCPtr(), 3156095363) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret int
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{value.AsCTypePtr(), gdc.ConstTypePtr(unsafe.Pointer(&from)), }

  pinner.Pin(&from)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedColorArray) Rfind(value Color, from int, ) int {
  name := StringNameFromStr("rfind")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, name.AsCPtr(), 3156095363) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret int
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{value.AsCTypePtr(), gdc.ConstTypePtr(unsafe.Pointer(&from)), }

  pinner.Pin(&from)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedColorArray) Count(value Color, ) int {
  name := StringNameFromStr("count")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedColorArray, name.AsCPtr(), 1682108616) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret int
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{value.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

// Operators

func (me *PackedColorArray) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedColorArray) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedColorArray) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *PackedColorArray) InDictionary(right Dictionary) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedColorArray) InArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedColorArray) EqualPackedColorArray(right PackedColorArray) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedColorArray) NotEqualPackedColorArray(right PackedColorArray) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedColorArray) AddPackedColorArray(right PackedColorArray) PackedColorArray {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, me.Type(), right.Type()) // FIXME: cache
  var ret PackedColorArray
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

// Members
