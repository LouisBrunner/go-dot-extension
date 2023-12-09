// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Transform2D struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Constants

var (
  Transform2DIdentity = "Transform2D(1, 0, 0, 1, 0, 0)" // TODO: construct correctly
  Transform2DFlipX = "Transform2D(-1, 0, 0, 1, 0, 0)" // TODO: construct correctly
  Transform2DFlipY = "Transform2D(1, 0, 0, -1, 0, 0)" // TODO: construct correctly
)

// Enums

// Constructors

func NewTransform2D() Transform2D {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeTransform2D))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeTransform2D, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Transform2D{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewTransform2DFromTransform2D(from Transform2D, ) Transform2D {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeTransform2D))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeTransform2D, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Transform2D{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewTransform2DFromFloat32Vector2(rotation float32, position Vector2, ) Transform2D {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeTransform2D))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeTransform2D, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&rotation), position.AsCTypePtr(), }))
  return Transform2D{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewTransform2DFromFloat32Vector2Float32Vector2(rotation float32, scale Vector2, skew float32, position Vector2, ) Transform2D {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeTransform2D))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeTransform2D, 3) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&rotation), scale.AsCTypePtr(), gdc.ConstTypePtr(&skew), position.AsCTypePtr(), }))
  return Transform2D{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewTransform2DFromVector2Vector2Vector2(x_axis Vector2, y_axis Vector2, origin Vector2, ) Transform2D {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeTransform2D))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeTransform2D, 4) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{x_axis.AsCTypePtr(), y_axis.AsCTypePtr(), origin.AsCTypePtr(), }))
  return Transform2D{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Transform2D) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Transform2D) Type() gdc.VariantType {
  return gdc.VariantTypeTransform2D
}

func (me *Transform2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Transform2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func  (me *Transform2D) Inverse() Transform2D {
  panic("TODO: implement")
}

func  (me *Transform2D) AffineInverse() Transform2D {
  panic("TODO: implement")
}

func  (me *Transform2D) GetRotation() float32 {
  panic("TODO: implement")
}

func  (me *Transform2D) GetOrigin() Vector2 {
  panic("TODO: implement")
}

func  (me *Transform2D) GetScale() Vector2 {
  panic("TODO: implement")
}

func  (me *Transform2D) GetSkew() float32 {
  panic("TODO: implement")
}

func  (me *Transform2D) Orthonormalized() Transform2D {
  panic("TODO: implement")
}

func  (me *Transform2D) Rotated(angle float32, ) Transform2D {
  panic("TODO: implement")
}

func  (me *Transform2D) RotatedLocal(angle float32, ) Transform2D {
  panic("TODO: implement")
}

func  (me *Transform2D) Scaled(scale Vector2, ) Transform2D {
  panic("TODO: implement")
}

func  (me *Transform2D) ScaledLocal(scale Vector2, ) Transform2D {
  panic("TODO: implement")
}

func  (me *Transform2D) Translated(offset Vector2, ) Transform2D {
  panic("TODO: implement")
}

func  (me *Transform2D) TranslatedLocal(offset Vector2, ) Transform2D {
  panic("TODO: implement")
}

func  (me *Transform2D) Determinant() float32 {
  panic("TODO: implement")
}

func  (me *Transform2D) BasisXform(v Vector2, ) Vector2 {
  panic("TODO: implement")
}

func  (me *Transform2D) BasisXformInv(v Vector2, ) Vector2 {
  panic("TODO: implement")
}

func  (me *Transform2D) InterpolateWith(xform Transform2D, weight float32, ) Transform2D {
  panic("TODO: implement")
}

func  (me *Transform2D) IsEqualApprox(xform Transform2D, ) bool {
  panic("TODO: implement")
}

func  (me *Transform2D) IsFinite() bool {
  panic("TODO: implement")
}

func  (me *Transform2D) LookingAt(target Vector2, ) Transform2D {
  panic("TODO: implement")
}

// Operators

func (me *Transform2D) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Transform2D) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Transform2D) Not() bool {
  panic("TODO: implement")
}

func (me *Transform2D) MultiplyInt(right int) Transform2D {
  panic("TODO: implement")
}

func (me *Transform2D) MultiplyFloat32(right float32) Transform2D {
  panic("TODO: implement")
}

func (me *Transform2D) MultiplyVector2(right Vector2) Vector2 {
  panic("TODO: implement")
}

func (me *Transform2D) MultiplyRect2(right Rect2) Rect2 {
  panic("TODO: implement")
}

func (me *Transform2D) EqualsTransform2D(right Transform2D) bool {
  panic("TODO: implement")
}

func (me *Transform2D) NotEqualsTransform2D(right Transform2D) bool {
  panic("TODO: implement")
}

func (me *Transform2D) MultiplyTransform2D(right Transform2D) Transform2D {
  panic("TODO: implement")
}

func (me *Transform2D) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *Transform2D) InArray(right Array) bool {
  panic("TODO: implement")
}

func (me *Transform2D) MultiplyPackedVector2Array(right PackedVector2Array) PackedVector2Array {
  panic("TODO: implement")
}

// TODO: members (bclass)
