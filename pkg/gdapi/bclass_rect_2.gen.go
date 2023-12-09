// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Rect2 struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Enums

// Constructors

func NewRect2() Rect2 {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeRect2))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeRect2, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Rect2{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewRect2FromRect2(from Rect2, ) Rect2 {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeRect2))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeRect2, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Rect2{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewRect2FromRect2I(from Rect2i, ) Rect2 {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeRect2))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeRect2, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Rect2{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewRect2FromVector2Vector2(position Vector2, size Vector2, ) Rect2 {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeRect2))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeRect2, 3) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{position.AsCTypePtr(), size.AsCTypePtr(), }))
  return Rect2{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewRect2FromFloat32Float32Float32Float32(x float32, y float32, width float32, height float32, ) Rect2 {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeRect2))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeRect2, 4) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), }))
  return Rect2{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Rect2) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Rect2) Type() gdc.VariantType {
  return gdc.VariantTypeRect2
}

func (me *Rect2) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Rect2) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func  (me *Rect2) GetCenter() Vector2 {
  panic("TODO: implement")
}

func  (me *Rect2) GetArea() float32 {
  panic("TODO: implement")
}

func  (me *Rect2) HasArea() bool {
  panic("TODO: implement")
}

func  (me *Rect2) HasPoint(point Vector2, ) bool {
  panic("TODO: implement")
}

func  (me *Rect2) IsEqualApprox(rect Rect2, ) bool {
  panic("TODO: implement")
}

func  (me *Rect2) IsFinite() bool {
  panic("TODO: implement")
}

func  (me *Rect2) Intersects(b Rect2, include_borders bool, ) bool {
  panic("TODO: implement")
}

func  (me *Rect2) Encloses(b Rect2, ) bool {
  panic("TODO: implement")
}

func  (me *Rect2) Intersection(b Rect2, ) Rect2 {
  panic("TODO: implement")
}

func  (me *Rect2) Merge(b Rect2, ) Rect2 {
  panic("TODO: implement")
}

func  (me *Rect2) Expand(to Vector2, ) Rect2 {
  panic("TODO: implement")
}

func  (me *Rect2) Grow(amount float32, ) Rect2 {
  panic("TODO: implement")
}

func  (me *Rect2) GrowSide(side int, amount float32, ) Rect2 {
  panic("TODO: implement")
}

func  (me *Rect2) GrowIndividual(left float32, top float32, right float32, bottom float32, ) Rect2 {
  panic("TODO: implement")
}

func  (me *Rect2) Abs() Rect2 {
  panic("TODO: implement")
}

// Operators

func (me *Rect2) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Rect2) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Rect2) Not() bool {
  panic("TODO: implement")
}

func (me *Rect2) EqualsRect2(right Rect2) bool {
  panic("TODO: implement")
}

func (me *Rect2) NotEqualsRect2(right Rect2) bool {
  panic("TODO: implement")
}

func (me *Rect2) MultiplyTransform2D(right Transform2D) Rect2 {
  panic("TODO: implement")
}

func (me *Rect2) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *Rect2) InArray(right Array) bool {
  panic("TODO: implement")
}

// TODO: members (bclass)
