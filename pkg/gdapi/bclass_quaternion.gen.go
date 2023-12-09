// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Quaternion struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Constants

var (
  QuaternionIdentity = "Quaternion(0, 0, 0, 1)" // TODO: construct correctly
)

// Enums

// Constructors

func NewQuaternion() Quaternion {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeQuaternion))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeQuaternion, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Quaternion{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewQuaternionFromQuaternion(from Quaternion, ) Quaternion {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeQuaternion))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeQuaternion, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Quaternion{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewQuaternionFromBasis(from Basis, ) Quaternion {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeQuaternion))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeQuaternion, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Quaternion{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewQuaternionFromVector3Float32(axis Vector3, angle float32, ) Quaternion {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeQuaternion))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeQuaternion, 3) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{axis.AsCTypePtr(), gdc.ConstTypePtr(&angle), }))
  return Quaternion{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewQuaternionFromVector3Vector3(arc_from Vector3, arc_to Vector3, ) Quaternion {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeQuaternion))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeQuaternion, 4) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{arc_from.AsCTypePtr(), arc_to.AsCTypePtr(), }))
  return Quaternion{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewQuaternionFromFloat32Float32Float32Float32(x float32, y float32, z float32, w float32, ) Quaternion {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeQuaternion))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeQuaternion, 5) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), gdc.ConstTypePtr(&z), gdc.ConstTypePtr(&w), }))
  return Quaternion{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Quaternion) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Quaternion) Type() gdc.VariantType {
  return gdc.VariantTypeQuaternion
}

func (me *Quaternion) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Quaternion) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func  (me *Quaternion) Length() float32 {
  panic("TODO: implement")
}

func  (me *Quaternion) LengthSquared() float32 {
  panic("TODO: implement")
}

func  (me *Quaternion) Normalized() Quaternion {
  panic("TODO: implement")
}

func  (me *Quaternion) IsNormalized() bool {
  panic("TODO: implement")
}

func  (me *Quaternion) IsEqualApprox(to Quaternion, ) bool {
  panic("TODO: implement")
}

func  (me *Quaternion) IsFinite() bool {
  panic("TODO: implement")
}

func  (me *Quaternion) Inverse() Quaternion {
  panic("TODO: implement")
}

func  (me *Quaternion) Log() Quaternion {
  panic("TODO: implement")
}

func  (me *Quaternion) Exp() Quaternion {
  panic("TODO: implement")
}

func  (me *Quaternion) AngleTo(to Quaternion, ) float32 {
  panic("TODO: implement")
}

func  (me *Quaternion) Dot(with Quaternion, ) float32 {
  panic("TODO: implement")
}

func  (me *Quaternion) Slerp(to Quaternion, weight float32, ) Quaternion {
  panic("TODO: implement")
}

func  (me *Quaternion) Slerpni(to Quaternion, weight float32, ) Quaternion {
  panic("TODO: implement")
}

func  (me *Quaternion) SphericalCubicInterpolate(b Quaternion, pre_a Quaternion, post_b Quaternion, weight float32, ) Quaternion {
  panic("TODO: implement")
}

func  (me *Quaternion) SphericalCubicInterpolateInTime(b Quaternion, pre_a Quaternion, post_b Quaternion, weight float32, b_t float32, pre_a_t float32, post_b_t float32, ) Quaternion {
  panic("TODO: implement")
}

func  (me *Quaternion) GetEuler(order int, ) Vector3 {
  panic("TODO: implement")
}

func  QuaternionFromEuler(euler Vector3, ) Quaternion {
  panic("TODO: implement")
}

func  (me *Quaternion) GetAxis() Vector3 {
  panic("TODO: implement")
}

func  (me *Quaternion) GetAngle() float32 {
  panic("TODO: implement")
}

// Operators

func (me *Quaternion) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Quaternion) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Quaternion) UnaryMinus() Quaternion {
  panic("TODO: implement")
}

func (me *Quaternion) UnaryPlus() Quaternion {
  panic("TODO: implement")
}

func (me *Quaternion) Not() bool {
  panic("TODO: implement")
}

func (me *Quaternion) MultiplyInt(right int) Quaternion {
  panic("TODO: implement")
}

func (me *Quaternion) DivideInt(right int) Quaternion {
  panic("TODO: implement")
}

func (me *Quaternion) MultiplyFloat32(right float32) Quaternion {
  panic("TODO: implement")
}

func (me *Quaternion) DivideFloat32(right float32) Quaternion {
  panic("TODO: implement")
}

func (me *Quaternion) MultiplyVector3(right Vector3) Vector3 {
  panic("TODO: implement")
}

func (me *Quaternion) EqualsQuaternion(right Quaternion) bool {
  panic("TODO: implement")
}

func (me *Quaternion) NotEqualsQuaternion(right Quaternion) bool {
  panic("TODO: implement")
}

func (me *Quaternion) AddQuaternion(right Quaternion) Quaternion {
  panic("TODO: implement")
}

func (me *Quaternion) SubtractQuaternion(right Quaternion) Quaternion {
  panic("TODO: implement")
}

func (me *Quaternion) MultiplyQuaternion(right Quaternion) Quaternion {
  panic("TODO: implement")
}

func (me *Quaternion) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *Quaternion) InArray(right Array) bool {
  panic("TODO: implement")
}

// TODO: members (bclass)
