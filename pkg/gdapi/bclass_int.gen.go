// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Int struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Enums

// Constructors

func NewInt() Int {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeInt))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeInt, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Int{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewIntFromInt(from int, ) Int {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeInt))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeInt, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&from), }))
  return Int{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewIntFromFloat32(from float32, ) Int {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeInt))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeInt, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&from), }))
  return Int{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewIntFromBool(from bool, ) Int {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeInt))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeInt, 3) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&from), }))
  return Int{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewIntFromString(from String, ) Int {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeInt))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeInt, 4) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Int{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Int) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Int) Type() gdc.VariantType {
  return gdc.VariantTypeInt
}

func (me *Int) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Int) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

// Operators

func (me *Int) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Int) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Int) UnaryMinus() int {
  panic("TODO: implement")
}

func (me *Int) UnaryPlus() int {
  panic("TODO: implement")
}

func (me *Int) BitwiseNot() int {
  panic("TODO: implement")
}

func (me *Int) AndVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Int) OrVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Int) XorVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Int) Not() bool {
  panic("TODO: implement")
}

func (me *Int) AndBool(right bool) bool {
  panic("TODO: implement")
}

func (me *Int) OrBool(right bool) bool {
  panic("TODO: implement")
}

func (me *Int) XorBool(right bool) bool {
  panic("TODO: implement")
}

func (me *Int) EqualsInt(right int) bool {
  panic("TODO: implement")
}

func (me *Int) NotEqualsInt(right int) bool {
  panic("TODO: implement")
}

func (me *Int) LessThanInt(right int) bool {
  panic("TODO: implement")
}

func (me *Int) LessThanOrEqualsInt(right int) bool {
  panic("TODO: implement")
}

func (me *Int) GreaterThanInt(right int) bool {
  panic("TODO: implement")
}

func (me *Int) GreaterThanOrEqualsInt(right int) bool {
  panic("TODO: implement")
}

func (me *Int) AddInt(right int) int {
  panic("TODO: implement")
}

func (me *Int) SubtractInt(right int) int {
  panic("TODO: implement")
}

func (me *Int) MultiplyInt(right int) int {
  panic("TODO: implement")
}

func (me *Int) DivideInt(right int) int {
  panic("TODO: implement")
}

func (me *Int) ModuloInt(right int) int {
  panic("TODO: implement")
}

func (me *Int) ExponentInt(right int) int {
  panic("TODO: implement")
}

func (me *Int) LeftShiftInt(right int) int {
  panic("TODO: implement")
}

func (me *Int) RightShiftInt(right int) int {
  panic("TODO: implement")
}

func (me *Int) BitwiseAndInt(right int) int {
  panic("TODO: implement")
}

func (me *Int) BitwiseOrInt(right int) int {
  panic("TODO: implement")
}

func (me *Int) BitwiseXorInt(right int) int {
  panic("TODO: implement")
}

func (me *Int) AndInt(right int) bool {
  panic("TODO: implement")
}

func (me *Int) OrInt(right int) bool {
  panic("TODO: implement")
}

func (me *Int) XorInt(right int) bool {
  panic("TODO: implement")
}

func (me *Int) EqualsFloat32(right float32) bool {
  panic("TODO: implement")
}

func (me *Int) NotEqualsFloat32(right float32) bool {
  panic("TODO: implement")
}

func (me *Int) LessThanFloat32(right float32) bool {
  panic("TODO: implement")
}

func (me *Int) LessThanOrEqualsFloat32(right float32) bool {
  panic("TODO: implement")
}

func (me *Int) GreaterThanFloat32(right float32) bool {
  panic("TODO: implement")
}

func (me *Int) GreaterThanOrEqualsFloat32(right float32) bool {
  panic("TODO: implement")
}

func (me *Int) AddFloat32(right float32) float32 {
  panic("TODO: implement")
}

func (me *Int) SubtractFloat32(right float32) float32 {
  panic("TODO: implement")
}

func (me *Int) MultiplyFloat32(right float32) float32 {
  panic("TODO: implement")
}

func (me *Int) DivideFloat32(right float32) float32 {
  panic("TODO: implement")
}

func (me *Int) ExponentFloat32(right float32) float32 {
  panic("TODO: implement")
}

func (me *Int) AndFloat32(right float32) bool {
  panic("TODO: implement")
}

func (me *Int) OrFloat32(right float32) bool {
  panic("TODO: implement")
}

func (me *Int) XorFloat32(right float32) bool {
  panic("TODO: implement")
}

func (me *Int) MultiplyVector2(right Vector2) Vector2 {
  panic("TODO: implement")
}

func (me *Int) MultiplyVector2I(right Vector2i) Vector2i {
  panic("TODO: implement")
}

func (me *Int) MultiplyVector3(right Vector3) Vector3 {
  panic("TODO: implement")
}

func (me *Int) MultiplyVector3I(right Vector3i) Vector3i {
  panic("TODO: implement")
}

func (me *Int) MultiplyVector4(right Vector4) Vector4 {
  panic("TODO: implement")
}

func (me *Int) MultiplyVector4I(right Vector4i) Vector4i {
  panic("TODO: implement")
}

func (me *Int) MultiplyQuaternion(right Quaternion) Quaternion {
  panic("TODO: implement")
}

func (me *Int) MultiplyColor(right Color) Color {
  panic("TODO: implement")
}

func (me *Int) AndObject(right Object) bool {
  panic("TODO: implement")
}

func (me *Int) OrObject(right Object) bool {
  panic("TODO: implement")
}

func (me *Int) XorObject(right Object) bool {
  panic("TODO: implement")
}

func (me *Int) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *Int) InArray(right Array) bool {
  panic("TODO: implement")
}

func (me *Int) InPackedByteArray(right PackedByteArray) bool {
  panic("TODO: implement")
}

func (me *Int) InPackedInt32Array(right PackedInt32Array) bool {
  panic("TODO: implement")
}

func (me *Int) InPackedInt64Array(right PackedInt64Array) bool {
  panic("TODO: implement")
}

func (me *Int) InPackedFloat32Array(right PackedFloat32Array) bool {
  panic("TODO: implement")
}

func (me *Int) InPackedFloat64Array(right PackedFloat64Array) bool {
  panic("TODO: implement")
}

// TODO: members (bclass)
