// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Rect2i struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Enums

// Constructors

func NewRect2i() Rect2i {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeRect2i))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeRect2I, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Rect2i{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewRect2iFromRect2I(from Rect2i, ) Rect2i {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeRect2i))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeRect2I, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Rect2i{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewRect2iFromRect2(from Rect2, ) Rect2i {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeRect2i))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeRect2I, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Rect2i{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewRect2iFromVector2IVector2I(position Vector2i, size Vector2i, ) Rect2i {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeRect2i))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeRect2I, 3) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{position.AsCTypePtr(), size.AsCTypePtr(), }))
  return Rect2i{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewRect2iFromIntIntIntInt(x int, y int, width int, height int, ) Rect2i {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeRect2i))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeRect2I, 4) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), }))
  return Rect2i{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Rect2i) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Rect2i) Type() gdc.VariantType {
  return gdc.VariantTypeRect2I
}

func (me *Rect2i) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Rect2i) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func  (me *Rect2i) GetCenter() Vector2i {
  panic("TODO: implement")
}

func  (me *Rect2i) GetArea() int {
  panic("TODO: implement")
}

func  (me *Rect2i) HasArea() bool {
  panic("TODO: implement")
}

func  (me *Rect2i) HasPoint(point Vector2i, ) bool {
  panic("TODO: implement")
}

func  (me *Rect2i) Intersects(b Rect2i, ) bool {
  panic("TODO: implement")
}

func  (me *Rect2i) Encloses(b Rect2i, ) bool {
  panic("TODO: implement")
}

func  (me *Rect2i) Intersection(b Rect2i, ) Rect2i {
  panic("TODO: implement")
}

func  (me *Rect2i) Merge(b Rect2i, ) Rect2i {
  panic("TODO: implement")
}

func  (me *Rect2i) Expand(to Vector2i, ) Rect2i {
  panic("TODO: implement")
}

func  (me *Rect2i) Grow(amount int, ) Rect2i {
  panic("TODO: implement")
}

func  (me *Rect2i) GrowSide(side int, amount int, ) Rect2i {
  panic("TODO: implement")
}

func  (me *Rect2i) GrowIndividual(left int, top int, right int, bottom int, ) Rect2i {
  panic("TODO: implement")
}

func  (me *Rect2i) Abs() Rect2i {
  panic("TODO: implement")
}

// Operators

func (me *Rect2i) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Rect2i) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Rect2i) Not() bool {
  panic("TODO: implement")
}

func (me *Rect2i) EqualsRect2I(right Rect2i) bool {
  panic("TODO: implement")
}

func (me *Rect2i) NotEqualsRect2I(right Rect2i) bool {
  panic("TODO: implement")
}

func (me *Rect2i) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *Rect2i) InArray(right Array) bool {
  panic("TODO: implement")
}

// TODO: members (bclass)
