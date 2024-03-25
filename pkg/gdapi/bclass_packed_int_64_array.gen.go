// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "fmt"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypePackedInt64Array, 0) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewPackedInt64ArrayFromPackedInt64Array(from PackedInt64Array, ) *PackedInt64Array {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedInt64Array()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypePackedInt64Array, 1) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewPackedInt64ArrayFromArray(from Array, ) *PackedInt64Array {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedInt64Array()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypePackedInt64Array, 2) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

// Destructor
func (me *PackedInt64Array) Destroy() {
  dtr := me.iface.VariantGetPtrDestructor(gdc.VariantTypePackedInt64Array)
	me.iface.CallPtrDestructor(dtr, me.AsTypePtr())
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsPackedInt64Array() (*PackedInt64Array, error) {
	if me.Type() != gdc.VariantTypePackedInt64Array {
		return nil, fmt.Errorf("variant is not a PackedInt64Array")
	}
  bclass := newPackedInt64Array()
	fn := me.iface.GetVariantToTypeConstructor(me.Type())
	me.iface.CallTypeFromVariantConstructorFunc(fn, bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *PackedInt64Array) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  fn := me.iface.GetVariantFromTypeConstructor(me.Type())
  me.iface.CallVariantFromTypeConstructorFunc(fn, va.asUninitialized(), me.AsTypePtr())
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

// Methods

func (me *PackedInt64Array) Size() int {
  name := StringNameFromStr("size")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, name.AsCPtr(), 3173160232) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret int
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedInt64Array) IsEmpty() bool {
  name := StringNameFromStr("is_empty")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, name.AsCPtr(), 3918633141) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret bool
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedInt64Array) Set(index int, value int, )  {
  name := StringNameFromStr("set")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, name.AsCPtr(), 3638975848) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&index)), gdc.ConstTypePtr(unsafe.Pointer(&value)), }

  pinner.Pin(&index)
  pinner.Pin(&value)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt64Array) PushBack(value int, ) bool {
  name := StringNameFromStr("push_back")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, name.AsCPtr(), 694024632) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret bool
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&value)), }

  pinner.Pin(&value)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedInt64Array) Append(value int, ) bool {
  name := StringNameFromStr("append")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, name.AsCPtr(), 694024632) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret bool
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&value)), }

  pinner.Pin(&value)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedInt64Array) AppendArray(array PackedInt64Array, )  {
  name := StringNameFromStr("append_array")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, name.AsCPtr(), 2090311302) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  args := []gdc.ConstTypePtr{array.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt64Array) RemoveAt(index int, )  {
  name := StringNameFromStr("remove_at")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, name.AsCPtr(), 2823966027) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&index)), }

  pinner.Pin(&index)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt64Array) Insert(at_index int, value int, ) int {
  name := StringNameFromStr("insert")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, name.AsCPtr(), 1487112728) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret int
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&at_index)), gdc.ConstTypePtr(unsafe.Pointer(&value)), }

  pinner.Pin(&at_index)
  pinner.Pin(&value)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedInt64Array) Fill(value int, )  {
  name := StringNameFromStr("fill")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, name.AsCPtr(), 2823966027) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&value)), }

  pinner.Pin(&value)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt64Array) Resize(new_size int, ) int {
  name := StringNameFromStr("resize")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, name.AsCPtr(), 848867239) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret int
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&new_size)), }

  pinner.Pin(&new_size)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedInt64Array) Clear()  {
  name := StringNameFromStr("clear")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, name.AsCPtr(), 3218959716) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt64Array) Has(value int, ) bool {
  name := StringNameFromStr("has")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, name.AsCPtr(), 931488181) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret bool
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&value)), }

  pinner.Pin(&value)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedInt64Array) Reverse()  {
  name := StringNameFromStr("reverse")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, name.AsCPtr(), 3218959716) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt64Array) Slice(begin int, end int, ) PackedInt64Array {
  name := StringNameFromStr("slice")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, name.AsCPtr(), 1726550804) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret PackedInt64Array
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&begin)), gdc.ConstTypePtr(unsafe.Pointer(&end)), }

  pinner.Pin(&begin)
  pinner.Pin(&end)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedInt64Array) ToByteArray() PackedByteArray {
  name := StringNameFromStr("to_byte_array")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, name.AsCPtr(), 247621236) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret PackedByteArray
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedInt64Array) Sort()  {
  name := StringNameFromStr("sort")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, name.AsCPtr(), 3218959716) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt64Array) Bsearch(value int, before bool, ) int {
  name := StringNameFromStr("bsearch")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, name.AsCPtr(), 3380005890) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret int
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&value)), gdc.ConstTypePtr(unsafe.Pointer(&before)), }

  pinner.Pin(&value)
  pinner.Pin(&before)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedInt64Array) Duplicate() PackedInt64Array {
  name := StringNameFromStr("duplicate")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, name.AsCPtr(), 2376370016) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret PackedInt64Array
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedInt64Array) Find(value int, from int, ) int {
  name := StringNameFromStr("find")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, name.AsCPtr(), 2984303840) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret int
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&value)), gdc.ConstTypePtr(unsafe.Pointer(&from)), }

  pinner.Pin(&value)
  pinner.Pin(&from)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedInt64Array) Rfind(value int, from int, ) int {
  name := StringNameFromStr("rfind")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, name.AsCPtr(), 2984303840) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret int
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&value)), gdc.ConstTypePtr(unsafe.Pointer(&from)), }

  pinner.Pin(&value)
  pinner.Pin(&from)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedInt64Array) Count(value int, ) int {
  name := StringNameFromStr("count")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt64Array, name.AsCPtr(), 4103005248) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret int
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&value)), }

  pinner.Pin(&value)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

// Operators

func (me *PackedInt64Array) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedInt64Array) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedInt64Array) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *PackedInt64Array) InDictionary(right Dictionary) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedInt64Array) InArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedInt64Array) EqualPackedInt64Array(right PackedInt64Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedInt64Array) NotEqualPackedInt64Array(right PackedInt64Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedInt64Array) AddPackedInt64Array(right PackedInt64Array) PackedInt64Array {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, me.Type(), right.Type()) // FIXME: cache
  var ret PackedInt64Array
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

// Members
