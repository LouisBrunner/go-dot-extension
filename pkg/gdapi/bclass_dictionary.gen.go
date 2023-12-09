// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Dictionary struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Enums

// Constructors

func NewDictionary() Dictionary {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeDictionary))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeDictionary, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Dictionary{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewDictionaryFromDictionary(from Dictionary, ) Dictionary {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeDictionary))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeDictionary, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Dictionary{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Dictionary) Destroy() {
  if me.ptr == nil {
    return
  }
  dtr := me.iface.VariantGetPtrDestructor(gdc.VariantTypeDictionary)
	me.iface.CallPtrDestructor(dtr, gdc.TypePtr(me.ptr))
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Dictionary) Type() gdc.VariantType {
  return gdc.VariantTypeDictionary
}

func (me *Dictionary) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Dictionary) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func  (me *Dictionary) Size() int {
  panic("TODO: implement")
}

func  (me *Dictionary) IsEmpty() bool {
  panic("TODO: implement")
}

func  (me *Dictionary) Clear()  {
  panic("TODO: implement")
}

func  (me *Dictionary) Merge(dictionary Dictionary, overwrite bool, )  {
  panic("TODO: implement")
}

func  (me *Dictionary) Has(key Variant, ) bool {
  panic("TODO: implement")
}

func  (me *Dictionary) HasAll(keys Array, ) bool {
  panic("TODO: implement")
}

func  (me *Dictionary) FindKey(value Variant, ) Variant {
  panic("TODO: implement")
}

func  (me *Dictionary) Erase(key Variant, ) bool {
  panic("TODO: implement")
}

func  (me *Dictionary) Hash() int {
  panic("TODO: implement")
}

func  (me *Dictionary) Keys() Array {
  panic("TODO: implement")
}

func  (me *Dictionary) Values() Array {
  panic("TODO: implement")
}

func  (me *Dictionary) Duplicate(deep bool, ) Dictionary {
  panic("TODO: implement")
}

func  (me *Dictionary) Get(key Variant, default_ Variant, ) Variant {
  panic("TODO: implement")
}

func  (me *Dictionary) MakeReadOnly()  {
  panic("TODO: implement")
}

func  (me *Dictionary) IsReadOnly() bool {
  panic("TODO: implement")
}

// Operators

func (me *Dictionary) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Dictionary) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Dictionary) Not() bool {
  panic("TODO: implement")
}

func (me *Dictionary) EqualsDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *Dictionary) NotEqualsDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *Dictionary) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *Dictionary) InArray(right Array) bool {
  panic("TODO: implement")
}

// TODO: members (bclass)
