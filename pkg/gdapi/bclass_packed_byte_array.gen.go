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

func  (me *PackedByteArray) Size() int {
  panic("TODO: implement")
}

func  (me *PackedByteArray) IsEmpty() bool {
  panic("TODO: implement")
}

func  (me *PackedByteArray) Set(index int, value int, )  {
  panic("TODO: implement")
}

func  (me *PackedByteArray) PushBack(value int, ) bool {
  panic("TODO: implement")
}

func  (me *PackedByteArray) Append(value int, ) bool {
  panic("TODO: implement")
}

func  (me *PackedByteArray) AppendArray(array PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *PackedByteArray) RemoveAt(index int, )  {
  panic("TODO: implement")
}

func  (me *PackedByteArray) Insert(at_index int, value int, ) int {
  panic("TODO: implement")
}

func  (me *PackedByteArray) Fill(value int, )  {
  panic("TODO: implement")
}

func  (me *PackedByteArray) Resize(new_size int, ) int {
  panic("TODO: implement")
}

func  (me *PackedByteArray) Clear()  {
  panic("TODO: implement")
}

func  (me *PackedByteArray) Has(value int, ) bool {
  panic("TODO: implement")
}

func  (me *PackedByteArray) Reverse()  {
  panic("TODO: implement")
}

func  (me *PackedByteArray) Slice(begin int, end int, ) PackedByteArray {
  panic("TODO: implement")
}

func  (me *PackedByteArray) Sort()  {
  panic("TODO: implement")
}

func  (me *PackedByteArray) Bsearch(value int, before bool, ) int {
  panic("TODO: implement")
}

func  (me *PackedByteArray) Duplicate() PackedByteArray {
  panic("TODO: implement")
}

func  (me *PackedByteArray) Find(value int, from int, ) int {
  panic("TODO: implement")
}

func  (me *PackedByteArray) Rfind(value int, from int, ) int {
  panic("TODO: implement")
}

func  (me *PackedByteArray) Count(value int, ) int {
  panic("TODO: implement")
}

func  (me *PackedByteArray) GetStringFromAscii() String {
  panic("TODO: implement")
}

func  (me *PackedByteArray) GetStringFromUtf8() String {
  panic("TODO: implement")
}

func  (me *PackedByteArray) GetStringFromUtf16() String {
  panic("TODO: implement")
}

func  (me *PackedByteArray) GetStringFromUtf32() String {
  panic("TODO: implement")
}

func  (me *PackedByteArray) GetStringFromWchar() String {
  panic("TODO: implement")
}

func  (me *PackedByteArray) HexEncode() String {
  panic("TODO: implement")
}

func  (me *PackedByteArray) Compress(compression_mode int, ) PackedByteArray {
  panic("TODO: implement")
}

func  (me *PackedByteArray) Decompress(buffer_size int, compression_mode int, ) PackedByteArray {
  panic("TODO: implement")
}

func  (me *PackedByteArray) DecompressDynamic(max_output_size int, compression_mode int, ) PackedByteArray {
  panic("TODO: implement")
}

func  (me *PackedByteArray) DecodeU8(byte_offset int, ) int {
  panic("TODO: implement")
}

func  (me *PackedByteArray) DecodeS8(byte_offset int, ) int {
  panic("TODO: implement")
}

func  (me *PackedByteArray) DecodeU16(byte_offset int, ) int {
  panic("TODO: implement")
}

func  (me *PackedByteArray) DecodeS16(byte_offset int, ) int {
  panic("TODO: implement")
}

func  (me *PackedByteArray) DecodeU32(byte_offset int, ) int {
  panic("TODO: implement")
}

func  (me *PackedByteArray) DecodeS32(byte_offset int, ) int {
  panic("TODO: implement")
}

func  (me *PackedByteArray) DecodeU64(byte_offset int, ) int {
  panic("TODO: implement")
}

func  (me *PackedByteArray) DecodeS64(byte_offset int, ) int {
  panic("TODO: implement")
}

func  (me *PackedByteArray) DecodeHalf(byte_offset int, ) float32 {
  panic("TODO: implement")
}

func  (me *PackedByteArray) DecodeFloat(byte_offset int, ) float32 {
  panic("TODO: implement")
}

func  (me *PackedByteArray) DecodeDouble(byte_offset int, ) float32 {
  panic("TODO: implement")
}

func  (me *PackedByteArray) HasEncodedVar(byte_offset int, allow_objects bool, ) bool {
  panic("TODO: implement")
}

func  (me *PackedByteArray) DecodeVar(byte_offset int, allow_objects bool, ) Variant {
  panic("TODO: implement")
}

func  (me *PackedByteArray) DecodeVarSize(byte_offset int, allow_objects bool, ) int {
  panic("TODO: implement")
}

func  (me *PackedByteArray) ToInt32Array() PackedInt32Array {
  panic("TODO: implement")
}

func  (me *PackedByteArray) ToInt64Array() PackedInt64Array {
  panic("TODO: implement")
}

func  (me *PackedByteArray) ToFloat32Array() PackedFloat32Array {
  panic("TODO: implement")
}

func  (me *PackedByteArray) ToFloat64Array() PackedFloat64Array {
  panic("TODO: implement")
}

func  (me *PackedByteArray) EncodeU8(byte_offset int, value int, )  {
  panic("TODO: implement")
}

func  (me *PackedByteArray) EncodeS8(byte_offset int, value int, )  {
  panic("TODO: implement")
}

func  (me *PackedByteArray) EncodeU16(byte_offset int, value int, )  {
  panic("TODO: implement")
}

func  (me *PackedByteArray) EncodeS16(byte_offset int, value int, )  {
  panic("TODO: implement")
}

func  (me *PackedByteArray) EncodeU32(byte_offset int, value int, )  {
  panic("TODO: implement")
}

func  (me *PackedByteArray) EncodeS32(byte_offset int, value int, )  {
  panic("TODO: implement")
}

func  (me *PackedByteArray) EncodeU64(byte_offset int, value int, )  {
  panic("TODO: implement")
}

func  (me *PackedByteArray) EncodeS64(byte_offset int, value int, )  {
  panic("TODO: implement")
}

func  (me *PackedByteArray) EncodeHalf(byte_offset int, value float32, )  {
  panic("TODO: implement")
}

func  (me *PackedByteArray) EncodeFloat(byte_offset int, value float32, )  {
  panic("TODO: implement")
}

func  (me *PackedByteArray) EncodeDouble(byte_offset int, value float32, )  {
  panic("TODO: implement")
}

func  (me *PackedByteArray) EncodeVar(byte_offset int, value Variant, allow_objects bool, ) int {
  panic("TODO: implement")
}

// Operators

func (me *PackedByteArray) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *PackedByteArray) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *PackedByteArray) Not() bool {
  panic("TODO: implement")
}

func (me *PackedByteArray) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *PackedByteArray) InArray(right Array) bool {
  panic("TODO: implement")
}

func (me *PackedByteArray) EqualsPackedByteArray(right PackedByteArray) bool {
  panic("TODO: implement")
}

func (me *PackedByteArray) NotEqualsPackedByteArray(right PackedByteArray) bool {
  panic("TODO: implement")
}

func (me *PackedByteArray) AddPackedByteArray(right PackedByteArray) PackedByteArray {
  panic("TODO: implement")
}

// TODO: members (bclass)
