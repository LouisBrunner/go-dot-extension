// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PackedByteArray struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Enums

// Constructors

func NewPackedByteArray() PackedByteArray {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePackedByteArray))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePackedByteArray, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return PackedByteArray{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPackedByteArrayFromPackedByteArray(from PackedByteArray, ) PackedByteArray {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePackedByteArray))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePackedByteArray, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return PackedByteArray{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPackedByteArrayFromArray(from Array, ) PackedByteArray {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePackedByteArray))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePackedByteArray, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return PackedByteArray{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *PackedByteArray) Destroy() {
  if me.ptr == nil {
    return
  }
  dtr := me.iface.VariantGetPtrDestructor(gdc.VariantTypePackedByteArray)
	me.iface.CallPtrDestructor(dtr, gdc.TypePtr(me.ptr))
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *PackedByteArray) Type() gdc.VariantType {
  return gdc.VariantTypePackedByteArray
}

func (me *PackedByteArray) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *PackedByteArray) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func (me *PackedByteArray) Size() int {
  name := StringNameFromStr("size")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3173160232) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{}
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) IsEmpty() bool {
  name := StringNameFromStr("is_empty")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) Set(index int, value int, )  {
  name := StringNameFromStr("set")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3638975848) // FIXME: should cache?

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&value), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) PushBack(value int, ) bool {
  name := StringNameFromStr("push_back")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 694024632) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) Append(value int, ) bool {
  name := StringNameFromStr("append")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 694024632) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) AppendArray(array PackedByteArray, )  {
  name := StringNameFromStr("append_array")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 791097111) // FIXME: should cache?

  args := []gdc.ConstTypePtr{array.AsCTypePtr(), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) RemoveAt(index int, )  {
  name := StringNameFromStr("remove_at")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 2823966027) // FIXME: should cache?

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) Insert(at_index int, value int, ) int {
  name := StringNameFromStr("insert")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 1487112728) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&at_index), gdc.ConstTypePtr(&value), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) Fill(value int, )  {
  name := StringNameFromStr("fill")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 2823966027) // FIXME: should cache?

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) Resize(new_size int, ) int {
  name := StringNameFromStr("resize")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 848867239) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&new_size), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) Clear()  {
  name := StringNameFromStr("clear")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3218959716) // FIXME: should cache?

  args := []gdc.ConstTypePtr{}
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) Has(value int, ) bool {
  name := StringNameFromStr("has")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 931488181) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) Reverse()  {
  name := StringNameFromStr("reverse")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3218959716) // FIXME: should cache?

  args := []gdc.ConstTypePtr{}
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) Slice(begin int, end int, ) PackedByteArray {
  name := StringNameFromStr("slice")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 2278869132) // FIXME: should cache?

  var ret PackedByteArray
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&begin), gdc.ConstTypePtr(&end), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) Sort()  {
  name := StringNameFromStr("sort")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3218959716) // FIXME: should cache?

  args := []gdc.ConstTypePtr{}
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) Bsearch(value int, before bool, ) int {
  name := StringNameFromStr("bsearch")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3380005890) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), gdc.ConstTypePtr(&before), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) Duplicate() PackedByteArray {
  name := StringNameFromStr("duplicate")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 851781288) // FIXME: should cache?

  var ret PackedByteArray
  args := []gdc.ConstTypePtr{}
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) Find(value int, from int, ) int {
  name := StringNameFromStr("find")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 2984303840) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), gdc.ConstTypePtr(&from), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) Rfind(value int, from int, ) int {
  name := StringNameFromStr("rfind")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 2984303840) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), gdc.ConstTypePtr(&from), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) Count(value int, ) int {
  name := StringNameFromStr("count")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 4103005248) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) GetStringFromAscii() String {
  name := StringNameFromStr("get_string_from_ascii")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) GetStringFromUtf8() String {
  name := StringNameFromStr("get_string_from_utf8")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) GetStringFromUtf16() String {
  name := StringNameFromStr("get_string_from_utf16")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) GetStringFromUtf32() String {
  name := StringNameFromStr("get_string_from_utf32")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) GetStringFromWchar() String {
  name := StringNameFromStr("get_string_from_wchar")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) HexEncode() String {
  name := StringNameFromStr("hex_encode")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) Compress(compression_mode int, ) PackedByteArray {
  name := StringNameFromStr("compress")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 1845905913) // FIXME: should cache?

  var ret PackedByteArray
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&compression_mode), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) Decompress(buffer_size int, compression_mode int, ) PackedByteArray {
  name := StringNameFromStr("decompress")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 2278869132) // FIXME: should cache?

  var ret PackedByteArray
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&buffer_size), gdc.ConstTypePtr(&compression_mode), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) DecompressDynamic(max_output_size int, compression_mode int, ) PackedByteArray {
  name := StringNameFromStr("decompress_dynamic")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 2278869132) // FIXME: should cache?

  var ret PackedByteArray
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_output_size), gdc.ConstTypePtr(&compression_mode), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) DecodeU8(byte_offset int, ) int {
  name := StringNameFromStr("decode_u8")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 4103005248) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) DecodeS8(byte_offset int, ) int {
  name := StringNameFromStr("decode_s8")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 4103005248) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) DecodeU16(byte_offset int, ) int {
  name := StringNameFromStr("decode_u16")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 4103005248) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) DecodeS16(byte_offset int, ) int {
  name := StringNameFromStr("decode_s16")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 4103005248) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) DecodeU32(byte_offset int, ) int {
  name := StringNameFromStr("decode_u32")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 4103005248) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) DecodeS32(byte_offset int, ) int {
  name := StringNameFromStr("decode_s32")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 4103005248) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) DecodeU64(byte_offset int, ) int {
  name := StringNameFromStr("decode_u64")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 4103005248) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) DecodeS64(byte_offset int, ) int {
  name := StringNameFromStr("decode_s64")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 4103005248) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) DecodeHalf(byte_offset int, ) float32 {
  name := StringNameFromStr("decode_half")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 1401583798) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) DecodeFloat(byte_offset int, ) float32 {
  name := StringNameFromStr("decode_float")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 1401583798) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) DecodeDouble(byte_offset int, ) float32 {
  name := StringNameFromStr("decode_double")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 1401583798) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) HasEncodedVar(byte_offset int, allow_objects bool, ) bool {
  name := StringNameFromStr("has_encoded_var")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 2914632957) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset), gdc.ConstTypePtr(&allow_objects), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) DecodeVar(byte_offset int, allow_objects bool, ) Variant {
  name := StringNameFromStr("decode_var")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 1740420038) // FIXME: should cache?

  var ret Variant
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset), gdc.ConstTypePtr(&allow_objects), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) DecodeVarSize(byte_offset int, allow_objects bool, ) int {
  name := StringNameFromStr("decode_var_size")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 954237325) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset), gdc.ConstTypePtr(&allow_objects), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) ToInt32Array() PackedInt32Array {
  name := StringNameFromStr("to_int32_array")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3158844420) // FIXME: should cache?

  var ret PackedInt32Array
  args := []gdc.ConstTypePtr{}
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) ToInt64Array() PackedInt64Array {
  name := StringNameFromStr("to_int64_array")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 1961294120) // FIXME: should cache?

  var ret PackedInt64Array
  args := []gdc.ConstTypePtr{}
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) ToFloat32Array() PackedFloat32Array {
  name := StringNameFromStr("to_float32_array")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3575107827) // FIXME: should cache?

  var ret PackedFloat32Array
  args := []gdc.ConstTypePtr{}
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) ToFloat64Array() PackedFloat64Array {
  name := StringNameFromStr("to_float64_array")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 1627308337) // FIXME: should cache?

  var ret PackedFloat64Array
  args := []gdc.ConstTypePtr{}
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedByteArray) EncodeU8(byte_offset int, value int, )  {
  name := StringNameFromStr("encode_u8")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3638975848) // FIXME: should cache?

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset), gdc.ConstTypePtr(&value), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeS8(byte_offset int, value int, )  {
  name := StringNameFromStr("encode_s8")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3638975848) // FIXME: should cache?

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset), gdc.ConstTypePtr(&value), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeU16(byte_offset int, value int, )  {
  name := StringNameFromStr("encode_u16")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3638975848) // FIXME: should cache?

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset), gdc.ConstTypePtr(&value), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeS16(byte_offset int, value int, )  {
  name := StringNameFromStr("encode_s16")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3638975848) // FIXME: should cache?

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset), gdc.ConstTypePtr(&value), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeU32(byte_offset int, value int, )  {
  name := StringNameFromStr("encode_u32")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3638975848) // FIXME: should cache?

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset), gdc.ConstTypePtr(&value), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeS32(byte_offset int, value int, )  {
  name := StringNameFromStr("encode_s32")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3638975848) // FIXME: should cache?

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset), gdc.ConstTypePtr(&value), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeU64(byte_offset int, value int, )  {
  name := StringNameFromStr("encode_u64")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3638975848) // FIXME: should cache?

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset), gdc.ConstTypePtr(&value), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeS64(byte_offset int, value int, )  {
  name := StringNameFromStr("encode_s64")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 3638975848) // FIXME: should cache?

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset), gdc.ConstTypePtr(&value), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeHalf(byte_offset int, value float32, )  {
  name := StringNameFromStr("encode_half")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 1113000516) // FIXME: should cache?

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset), gdc.ConstTypePtr(&value), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeFloat(byte_offset int, value float32, )  {
  name := StringNameFromStr("encode_float")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 1113000516) // FIXME: should cache?

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset), gdc.ConstTypePtr(&value), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeDouble(byte_offset int, value float32, )  {
  name := StringNameFromStr("encode_double")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 1113000516) // FIXME: should cache?

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset), gdc.ConstTypePtr(&value), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeVar(byte_offset int, value Variant, allow_objects bool, ) int {
  name := StringNameFromStr("encode_var")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, name.AsCPtr(), 2604460497) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset), value.AsCTypePtr(), gdc.ConstTypePtr(&allow_objects), }
  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

// Operators

func (me *PackedByteArray) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedByteArray) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedByteArray) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *PackedByteArray) InDictionary(right Dictionary) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedByteArray) InArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedByteArray) EqualPackedByteArray(right PackedByteArray) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedByteArray) NotEqualPackedByteArray(right PackedByteArray) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedByteArray) AddPackedByteArray(right PackedByteArray) PackedByteArray {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, me.Type(), right.Type()) // FIXME: cache
  var ret PackedByteArray
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

// Members
