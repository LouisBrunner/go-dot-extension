// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "fmt"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PackedFloat64Array struct {
  data   *[classSizePackedFloat64Array]byte
  iface  gdc.Interface
  pinner runtime.Pinner
}

// Enums

// Constructors
func newPackedFloat64Array() *PackedFloat64Array {
  me := &PackedFloat64Array{
    data:   new([classSizePackedFloat64Array]byte),
    iface:  giface,
  }
  me.pinner.Pin(me)
  me.pinner.Pin(me.AsTypePtr())
  return me
}

func NewPackedFloat64Array() *PackedFloat64Array {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedFloat64Array()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypePackedFloat64Array, 0) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewPackedFloat64ArrayFromPackedFloat64Array(from PackedFloat64Array, ) *PackedFloat64Array {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedFloat64Array()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypePackedFloat64Array, 1) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewPackedFloat64ArrayFromArray(from Array, ) *PackedFloat64Array {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedFloat64Array()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypePackedFloat64Array, 2) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

// Destructor
func (me *PackedFloat64Array) Destroy() {
  dtr := me.iface.VariantGetPtrDestructor(gdc.VariantTypePackedFloat64Array)
	me.iface.CallPtrDestructor(dtr, me.AsTypePtr())
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsPackedFloat64Array() (*PackedFloat64Array, error) {
	if me.Type() != gdc.VariantTypePackedFloat64Array {
		return nil, fmt.Errorf("variant is not a PackedFloat64Array")
	}
  bclass := newPackedFloat64Array()
	fn := me.iface.GetVariantToTypeConstructor(me.Type())
	me.iface.CallTypeFromVariantConstructorFunc(fn, bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *PackedFloat64Array) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  fn := me.iface.GetVariantFromTypeConstructor(me.Type())
  me.iface.CallVariantFromTypeConstructorFunc(fn, va.asUninitialized(), me.AsTypePtr())
  return va
}

// Pointers
func PackedFloat64ArrayFromPtr(ptr gdc.ConstTypePtr) *PackedFloat64Array {
  me := newPackedFloat64Array()
  dataFromPtr(me.data[:], ptr)
  return me
}

func (me *PackedFloat64Array) Type() gdc.VariantType {
  return gdc.VariantTypePackedFloat64Array
}

func (me *PackedFloat64Array) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *PackedFloat64Array) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *PackedFloat64Array) asUninitialized() gdc.UninitializedTypePtr {
  return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *PackedFloat64Array) Size() int {
  name := StringNameFromStr("size")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 3173160232) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret int
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedFloat64Array) IsEmpty() bool {
  name := StringNameFromStr("is_empty")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 3918633141) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret bool
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedFloat64Array) Set(index int, value float32, )  {
  name := StringNameFromStr("set")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 1113000516) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&index)), gdc.ConstTypePtr(unsafe.Pointer(&value)), }

  pinner.Pin(&index)
  pinner.Pin(&value)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat64Array) PushBack(value float32, ) bool {
  name := StringNameFromStr("push_back")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 4094791666) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret bool
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&value)), }

  pinner.Pin(&value)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedFloat64Array) Append(value float32, ) bool {
  name := StringNameFromStr("append")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 4094791666) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret bool
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&value)), }

  pinner.Pin(&value)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedFloat64Array) AppendArray(array PackedFloat64Array, )  {
  name := StringNameFromStr("append_array")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 792078629) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  args := []gdc.ConstTypePtr{array.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat64Array) RemoveAt(index int, )  {
  name := StringNameFromStr("remove_at")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 2823966027) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&index)), }

  pinner.Pin(&index)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat64Array) Insert(at_index int, value float32, ) int {
  name := StringNameFromStr("insert")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 1379903876) // FIXME: should cache?

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

func (me *PackedFloat64Array) Fill(value float32, )  {
  name := StringNameFromStr("fill")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 833936903) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&value)), }

  pinner.Pin(&value)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat64Array) Resize(new_size int, ) int {
  name := StringNameFromStr("resize")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 848867239) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret int
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&new_size)), }

  pinner.Pin(&new_size)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedFloat64Array) Clear()  {
  name := StringNameFromStr("clear")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 3218959716) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat64Array) Has(value float32, ) bool {
  name := StringNameFromStr("has")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 1296369134) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret bool
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&value)), }

  pinner.Pin(&value)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedFloat64Array) Reverse()  {
  name := StringNameFromStr("reverse")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 3218959716) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat64Array) Slice(begin int, end int, ) PackedFloat64Array {
  name := StringNameFromStr("slice")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 2192974324) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret PackedFloat64Array
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(unsafe.Pointer(&begin)), gdc.ConstTypePtr(unsafe.Pointer(&end)), }

  pinner.Pin(&begin)
  pinner.Pin(&end)

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedFloat64Array) ToByteArray() PackedByteArray {
  name := StringNameFromStr("to_byte_array")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 247621236) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret PackedByteArray
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedFloat64Array) Sort()  {
  name := StringNameFromStr("sort")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 3218959716) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat64Array) Bsearch(value float32, before bool, ) int {
  name := StringNameFromStr("bsearch")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 1188816338) // FIXME: should cache?

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

func (me *PackedFloat64Array) Duplicate() PackedFloat64Array {
  name := StringNameFromStr("duplicate")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 949266573) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret PackedFloat64Array
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedFloat64Array) Find(value float32, from int, ) int {
  name := StringNameFromStr("find")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 1343150241) // FIXME: should cache?

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

func (me *PackedFloat64Array) Rfind(value float32, from int, ) int {
  name := StringNameFromStr("rfind")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 1343150241) // FIXME: should cache?

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

func (me *PackedFloat64Array) Count(value float32, ) int {
  name := StringNameFromStr("count")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 2859915090) // FIXME: should cache?

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

func (me *PackedFloat64Array) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedFloat64Array) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedFloat64Array) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *PackedFloat64Array) InDictionary(right Dictionary) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedFloat64Array) InArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedFloat64Array) EqualPackedFloat64Array(right PackedFloat64Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedFloat64Array) NotEqualPackedFloat64Array(right PackedFloat64Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedFloat64Array) AddPackedFloat64Array(right PackedFloat64Array) PackedFloat64Array {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, me.Type(), right.Type()) // FIXME: cache
  var ret PackedFloat64Array
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

// Members
