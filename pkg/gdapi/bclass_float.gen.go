// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Float struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Enums

// Constructors

func NewFloat() Float {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeFloat))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeFloat, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Float{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewFloatFromFloat32(from float32, ) Float {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeFloat))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeFloat, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&from), }))
  return Float{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewFloatFromInt(from int, ) Float {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeFloat))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeFloat, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&from), }))
  return Float{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewFloatFromBool(from bool, ) Float {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeFloat))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeFloat, 3) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&from), }))
  return Float{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewFloatFromString(from String, ) Float {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeFloat))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeFloat, 4) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Float{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Float) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Float) Type() gdc.VariantType {
  return gdc.VariantTypeFloat
}

func (me *Float) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Float) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

// Operators

func (me *Float) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Float) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Float) UnaryMinus() float32 {
  panic("TODO: implement")
}

func (me *Float) UnaryPlus() float32 {
  panic("TODO: implement")
}

func (me *Float) AndVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Float) OrVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Float) XorVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Float) Not() bool {
  panic("TODO: implement")
}

func (me *Float) AndBool(right bool) bool {
  panic("TODO: implement")
}

func (me *Float) OrBool(right bool) bool {
  panic("TODO: implement")
}

func (me *Float) XorBool(right bool) bool {
  panic("TODO: implement")
}

func (me *Float) EqualsInt(right int) bool {
  panic("TODO: implement")
}

func (me *Float) NotEqualsInt(right int) bool {
  panic("TODO: implement")
}

func (me *Float) LessThanInt(right int) bool {
  panic("TODO: implement")
}

func (me *Float) LessThanOrEqualsInt(right int) bool {
  panic("TODO: implement")
}

func (me *Float) GreaterThanInt(right int) bool {
  panic("TODO: implement")
}

func (me *Float) GreaterThanOrEqualsInt(right int) bool {
  panic("TODO: implement")
}

func (me *Float) AddInt(right int) float32 {
  panic("TODO: implement")
}

func (me *Float) SubtractInt(right int) float32 {
  panic("TODO: implement")
}

func (me *Float) MultiplyInt(right int) float32 {
  panic("TODO: implement")
}

func (me *Float) DivideInt(right int) float32 {
  panic("TODO: implement")
}

func (me *Float) ExponentInt(right int) float32 {
  panic("TODO: implement")
}

func (me *Float) AndInt(right int) bool {
  panic("TODO: implement")
}

func (me *Float) OrInt(right int) bool {
  panic("TODO: implement")
}

func (me *Float) XorInt(right int) bool {
  panic("TODO: implement")
}

func (me *Float) EqualsFloat32(right float32) bool {
  panic("TODO: implement")
}

func (me *Float) NotEqualsFloat32(right float32) bool {
  panic("TODO: implement")
}

func (me *Float) LessThanFloat32(right float32) bool {
  panic("TODO: implement")
}

func (me *Float) LessThanOrEqualsFloat32(right float32) bool {
  panic("TODO: implement")
}

func (me *Float) GreaterThanFloat32(right float32) bool {
  panic("TODO: implement")
}

func (me *Float) GreaterThanOrEqualsFloat32(right float32) bool {
  panic("TODO: implement")
}

func (me *Float) AddFloat32(right float32) float32 {
  panic("TODO: implement")
}

func (me *Float) SubtractFloat32(right float32) float32 {
  panic("TODO: implement")
}

func (me *Float) MultiplyFloat32(right float32) float32 {
  panic("TODO: implement")
}

func (me *Float) DivideFloat32(right float32) float32 {
  panic("TODO: implement")
}

func (me *Float) ExponentFloat32(right float32) float32 {
  panic("TODO: implement")
}

func (me *Float) AndFloat32(right float32) bool {
  panic("TODO: implement")
}

func (me *Float) OrFloat32(right float32) bool {
  panic("TODO: implement")
}

func (me *Float) XorFloat32(right float32) bool {
  panic("TODO: implement")
}

func (me *Float) MultiplyVector2(right Vector2) Vector2 {
  panic("TODO: implement")
}

func (me *Float) MultiplyVector2I(right Vector2i) Vector2 {
  panic("TODO: implement")
}

func (me *Float) MultiplyVector3(right Vector3) Vector3 {
  panic("TODO: implement")
}

func (me *Float) MultiplyVector3I(right Vector3i) Vector3 {
  panic("TODO: implement")
}

func (me *Float) MultiplyVector4(right Vector4) Vector4 {
  panic("TODO: implement")
}

func (me *Float) MultiplyVector4I(right Vector4i) Vector4 {
  panic("TODO: implement")
}

func (me *Float) MultiplyQuaternion(right Quaternion) Quaternion {
  panic("TODO: implement")
}

func (me *Float) MultiplyColor(right Color) Color {
  panic("TODO: implement")
}

func (me *Float) AndObject(right Object) bool {
  panic("TODO: implement")
}

func (me *Float) OrObject(right Object) bool {
  panic("TODO: implement")
}

func (me *Float) XorObject(right Object) bool {
  panic("TODO: implement")
}

func (me *Float) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *Float) InArray(right Array) bool {
  panic("TODO: implement")
}

func (me *Float) InPackedByteArray(right PackedByteArray) bool {
  panic("TODO: implement")
}

func (me *Float) InPackedInt32Array(right PackedInt32Array) bool {
  panic("TODO: implement")
}

func (me *Float) InPackedInt64Array(right PackedInt64Array) bool {
  panic("TODO: implement")
}

func (me *Float) InPackedFloat32Array(right PackedFloat32Array) bool {
  panic("TODO: implement")
}

func (me *Float) InPackedFloat64Array(right PackedFloat64Array) bool {
  panic("TODO: implement")
}

// TODO: members (bclass)
