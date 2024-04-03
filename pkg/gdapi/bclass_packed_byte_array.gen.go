// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "fmt"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PackedByteArray struct {
  data   *[classSizePackedByteArray]byte
  iface  gdc.Interface
  pinner runtime.Pinner
}

// Enums

// Constructors
func newPackedByteArray() *PackedByteArray {
  me := &PackedByteArray{
    data:   new([classSizePackedByteArray]byte),
    iface:  giface,
  }
  me.pinner.Pin(me)
  me.pinner.Pin(me.AsTypePtr())
  return me
}

func NewPackedByteArray() *PackedByteArray {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedByteArray()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypePackedByteArray, 0) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewPackedByteArrayFromPackedByteArray(from PackedByteArray, ) *PackedByteArray {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedByteArray()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypePackedByteArray, 1) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewPackedByteArrayFromArray(from Array, ) *PackedByteArray {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedByteArray()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypePackedByteArray, 2) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

// Destructor
func (me *PackedByteArray) Destroy() {
  dtr := me.iface.VariantGetPtrDestructor(gdc.VariantTypePackedByteArray)
	me.iface.CallPtrDestructor(dtr, me.AsTypePtr())
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsPackedByteArray() (*PackedByteArray, error) {
	if me.Type() != gdc.VariantTypePackedByteArray {
		return nil, fmt.Errorf("variant is not a PackedByteArray")
	}
  bclass := newPackedByteArray()
	fn := me.iface.GetVariantToTypeConstructor(me.Type())
	me.iface.CallTypeFromVariantConstructorFunc(fn, bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *PackedByteArray) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  fn := me.iface.GetVariantFromTypeConstructor(me.Type())
  me.iface.CallVariantFromTypeConstructorFunc(fn, va.asUninitialized(), me.AsTypePtr())
  return va
}

// Pointers
func PackedByteArrayFromPtr(ptr gdc.ConstTypePtr) *PackedByteArray {
  me := newPackedByteArray()
  dataFromPtr(me.data[:], ptr)
  return me
}

func (me *PackedByteArray) Type() gdc.VariantType {
  return gdc.VariantTypePackedByteArray
}

func (me *PackedByteArray) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *PackedByteArray) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *PackedByteArray) asUninitialized() gdc.UninitializedTypePtr {
  return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *PackedByteArray) Size() int64 {
  name := StringNameFromStr("size")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3173160232) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) IsEmpty() bool {
  name := StringNameFromStr("is_empty")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3918633141) // FIXME: should cache?

  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) Set(index int64, value int64, )  {
  name := StringNameFromStr("set")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3638975848) // FIXME: should cache?

  varg0 := NewIntFromInt(index)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) PushBack(value int64, ) bool {
  name := StringNameFromStr("push_back")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 694024632) // FIXME: should cache?

  ret := NewBool()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) Append(value int64, ) bool {
  name := StringNameFromStr("append")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 694024632) // FIXME: should cache?

  ret := NewBool()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) AppendArray(array PackedByteArray, )  {
  name := StringNameFromStr("append_array")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 791097111) // FIXME: should cache?


  args := []gdc.ConstTypePtr{array.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) RemoveAt(index int64, )  {
  name := StringNameFromStr("remove_at")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 2823966027) // FIXME: should cache?

  varg0 := NewIntFromInt(index)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) Insert(at_index int64, value int64, ) int64 {
  name := StringNameFromStr("insert")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 1487112728) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(at_index)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) Fill(value int64, )  {
  name := StringNameFromStr("fill")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 2823966027) // FIXME: should cache?

  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) Resize(new_size int64, ) int64 {
  name := StringNameFromStr("resize")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 848867239) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(new_size)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) Clear()  {
  name := StringNameFromStr("clear")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3218959716) // FIXME: should cache?

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) Has(value int64, ) bool {
  name := StringNameFromStr("has")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 931488181) // FIXME: should cache?

  ret := NewBool()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) Reverse()  {
  name := StringNameFromStr("reverse")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3218959716) // FIXME: should cache?

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) Slice(begin int64, end int64, ) PackedByteArray {
  name := StringNameFromStr("slice")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 2278869132) // FIXME: should cache?

  ret := NewPackedByteArray()

  varg0 := NewIntFromInt(begin)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(end)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) Sort()  {
  name := StringNameFromStr("sort")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3218959716) // FIXME: should cache?

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) Bsearch(value int64, before bool, ) int64 {
  name := StringNameFromStr("bsearch")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3380005890) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  varg1 := NewBoolFromBool(before)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) Duplicate() PackedByteArray {
  name := StringNameFromStr("duplicate")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 851781288) // FIXME: should cache?

  ret := NewPackedByteArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) Find(value int64, from int64, ) int64 {
  name := StringNameFromStr("find")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 2984303840) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(from)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) Rfind(value int64, from int64, ) int64 {
  name := StringNameFromStr("rfind")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 2984303840) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(from)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) Count(value int64, ) int64 {
  name := StringNameFromStr("count")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 4103005248) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) GetStringFromAscii() String {
  name := StringNameFromStr("get_string_from_ascii")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) GetStringFromUtf8() String {
  name := StringNameFromStr("get_string_from_utf8")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) GetStringFromUtf16() String {
  name := StringNameFromStr("get_string_from_utf16")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) GetStringFromUtf32() String {
  name := StringNameFromStr("get_string_from_utf32")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) GetStringFromWchar() String {
  name := StringNameFromStr("get_string_from_wchar")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) HexEncode() String {
  name := StringNameFromStr("hex_encode")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) Compress(compression_mode int64, ) PackedByteArray {
  name := StringNameFromStr("compress")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 1845905913) // FIXME: should cache?

  ret := NewPackedByteArray()

  varg0 := NewIntFromInt(compression_mode)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) Decompress(buffer_size int64, compression_mode int64, ) PackedByteArray {
  name := StringNameFromStr("decompress")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 2278869132) // FIXME: should cache?

  ret := NewPackedByteArray()

  varg0 := NewIntFromInt(buffer_size)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(compression_mode)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) DecompressDynamic(max_output_size int64, compression_mode int64, ) PackedByteArray {
  name := StringNameFromStr("decompress_dynamic")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 2278869132) // FIXME: should cache?

  ret := NewPackedByteArray()

  varg0 := NewIntFromInt(max_output_size)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(compression_mode)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) DecodeU8(byte_offset int64, ) int64 {
  name := StringNameFromStr("decode_u8")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 4103005248) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) DecodeS8(byte_offset int64, ) int64 {
  name := StringNameFromStr("decode_s8")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 4103005248) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) DecodeU16(byte_offset int64, ) int64 {
  name := StringNameFromStr("decode_u16")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 4103005248) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) DecodeS16(byte_offset int64, ) int64 {
  name := StringNameFromStr("decode_s16")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 4103005248) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) DecodeU32(byte_offset int64, ) int64 {
  name := StringNameFromStr("decode_u32")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 4103005248) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) DecodeS32(byte_offset int64, ) int64 {
  name := StringNameFromStr("decode_s32")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 4103005248) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) DecodeU64(byte_offset int64, ) int64 {
  name := StringNameFromStr("decode_u64")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 4103005248) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) DecodeS64(byte_offset int64, ) int64 {
  name := StringNameFromStr("decode_s64")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 4103005248) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) DecodeHalf(byte_offset int64, ) float64 {
  name := StringNameFromStr("decode_half")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 1401583798) // FIXME: should cache?

  ret := NewFloat()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) DecodeFloat(byte_offset int64, ) float64 {
  name := StringNameFromStr("decode_float")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 1401583798) // FIXME: should cache?

  ret := NewFloat()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) DecodeDouble(byte_offset int64, ) float64 {
  name := StringNameFromStr("decode_double")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 1401583798) // FIXME: should cache?

  ret := NewFloat()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) HasEncodedVar(byte_offset int64, allow_objects bool, ) bool {
  name := StringNameFromStr("has_encoded_var")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 2914632957) // FIXME: should cache?

  ret := NewBool()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewBoolFromBool(allow_objects)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) DecodeVar(byte_offset int64, allow_objects bool, ) Variant {
  name := StringNameFromStr("decode_var")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 1740420038) // FIXME: should cache?

  ret := NewVariant()

  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewBoolFromBool(allow_objects)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) DecodeVarSize(byte_offset int64, allow_objects bool, ) int64 {
  name := StringNameFromStr("decode_var_size")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 954237325) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewBoolFromBool(allow_objects)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) ToInt32Array() PackedInt32Array {
  name := StringNameFromStr("to_int32_array")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3158844420) // FIXME: should cache?

  ret := NewPackedInt32Array()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) ToInt64Array() PackedInt64Array {
  name := StringNameFromStr("to_int64_array")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 1961294120) // FIXME: should cache?

  ret := NewPackedInt64Array()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) ToFloat32Array() PackedFloat32Array {
  name := StringNameFromStr("to_float32_array")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3575107827) // FIXME: should cache?

  ret := NewPackedFloat32Array()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) ToFloat64Array() PackedFloat64Array {
  name := StringNameFromStr("to_float64_array")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 1627308337) // FIXME: should cache?

  ret := NewPackedFloat64Array()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) EncodeU8(byte_offset int64, value int64, )  {
  name := StringNameFromStr("encode_u8")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3638975848) // FIXME: should cache?

  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeS8(byte_offset int64, value int64, )  {
  name := StringNameFromStr("encode_s8")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3638975848) // FIXME: should cache?

  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeU16(byte_offset int64, value int64, )  {
  name := StringNameFromStr("encode_u16")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3638975848) // FIXME: should cache?

  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeS16(byte_offset int64, value int64, )  {
  name := StringNameFromStr("encode_s16")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3638975848) // FIXME: should cache?

  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeU32(byte_offset int64, value int64, )  {
  name := StringNameFromStr("encode_u32")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3638975848) // FIXME: should cache?

  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeS32(byte_offset int64, value int64, )  {
  name := StringNameFromStr("encode_s32")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3638975848) // FIXME: should cache?

  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeU64(byte_offset int64, value int64, )  {
  name := StringNameFromStr("encode_u64")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3638975848) // FIXME: should cache?

  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeS64(byte_offset int64, value int64, )  {
  name := StringNameFromStr("encode_s64")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3638975848) // FIXME: should cache?

  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeHalf(byte_offset int64, value float64, )  {
  name := StringNameFromStr("encode_half")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 1113000516) // FIXME: should cache?

  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewFloatFromFloat32(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeFloat(byte_offset int64, value float64, )  {
  name := StringNameFromStr("encode_float")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 1113000516) // FIXME: should cache?

  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewFloatFromFloat32(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeDouble(byte_offset int64, value float64, )  {
  name := StringNameFromStr("encode_double")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 1113000516) // FIXME: should cache?

  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewFloatFromFloat32(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeVar(byte_offset int64, value Variant, allow_objects bool, ) int64 {
  name := StringNameFromStr("encode_var")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 2604460497) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()

  varg2 := NewBoolFromBool(allow_objects)
  defer varg2.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), value.AsCTypePtr(), varg2.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

// Operators

func (me *PackedByteArray) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedByteArray) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedByteArray) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedByteArray) InDictionary(right Dictionary) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedByteArray) InArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedByteArray) EqualPackedByteArray(right PackedByteArray) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedByteArray) NotEqualPackedByteArray(right PackedByteArray) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedByteArray) AddPackedByteArray(right PackedByteArray) PackedByteArray {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, me.Type(), right.Type()) // FIXME: cache
  ret := NewPackedByteArray()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

// Members
