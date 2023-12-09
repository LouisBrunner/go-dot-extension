// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Basis struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Constants

var (
  BasisIdentity = "Basis(1, 0, 0, 0, 1, 0, 0, 0, 1)" // TODO: construct correctly
  BasisFlipX = "Basis(-1, 0, 0, 0, 1, 0, 0, 0, 1)" // TODO: construct correctly
  BasisFlipY = "Basis(1, 0, 0, 0, -1, 0, 0, 0, 1)" // TODO: construct correctly
  BasisFlipZ = "Basis(1, 0, 0, 0, 1, 0, 0, 0, -1)" // TODO: construct correctly
)

// Enums

// Constructors

func NewBasis() Basis {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeBasis))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeBasis, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Basis{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewBasisFromBasis(from Basis, ) Basis {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeBasis))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeBasis, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Basis{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewBasisFromQuaternion(from Quaternion, ) Basis {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeBasis))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeBasis, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Basis{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewBasisFromVector3Float32(axis Vector3, angle float32, ) Basis {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeBasis))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeBasis, 3) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{axis.AsCTypePtr(), gdc.ConstTypePtr(&angle), }))
  return Basis{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewBasisFromVector3Vector3Vector3(x_axis Vector3, y_axis Vector3, z_axis Vector3, ) Basis {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeBasis))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeBasis, 4) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{x_axis.AsCTypePtr(), y_axis.AsCTypePtr(), z_axis.AsCTypePtr(), }))
  return Basis{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Basis) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Basis) Type() gdc.VariantType {
  return gdc.VariantTypeBasis
}

func (me *Basis) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Basis) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func  (me *Basis) Inverse() Basis {
  panic("TODO: implement")
}

func  (me *Basis) Transposed() Basis {
  panic("TODO: implement")
}

func  (me *Basis) Orthonormalized() Basis {
  panic("TODO: implement")
}

func  (me *Basis) Determinant() float32 {
  panic("TODO: implement")
}

func  (me *Basis) Rotated(axis Vector3, angle float32, ) Basis {
  panic("TODO: implement")
}

func  (me *Basis) Scaled(scale Vector3, ) Basis {
  panic("TODO: implement")
}

func  (me *Basis) GetScale() Vector3 {
  panic("TODO: implement")
}

func  (me *Basis) GetEuler(order int, ) Vector3 {
  panic("TODO: implement")
}

func  (me *Basis) Tdotx(with Vector3, ) float32 {
  panic("TODO: implement")
}

func  (me *Basis) Tdoty(with Vector3, ) float32 {
  panic("TODO: implement")
}

func  (me *Basis) Tdotz(with Vector3, ) float32 {
  panic("TODO: implement")
}

func  (me *Basis) Slerp(to Basis, weight float32, ) Basis {
  panic("TODO: implement")
}

func  (me *Basis) IsEqualApprox(b Basis, ) bool {
  panic("TODO: implement")
}

func  (me *Basis) IsFinite() bool {
  panic("TODO: implement")
}

func  (me *Basis) GetRotationQuaternion() Quaternion {
  panic("TODO: implement")
}

func  BasisLookingAt(target Vector3, up Vector3, use_model_front bool, ) Basis {
  panic("TODO: implement")
}

func  BasisFromScale(scale Vector3, ) Basis {
  panic("TODO: implement")
}

func  BasisFromEuler(euler Vector3, order int, ) Basis {
  panic("TODO: implement")
}

// Operators

func (me *Basis) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Basis) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Basis) Not() bool {
  panic("TODO: implement")
}

func (me *Basis) MultiplyInt(right int) Basis {
  panic("TODO: implement")
}

func (me *Basis) MultiplyFloat32(right float32) Basis {
  panic("TODO: implement")
}

func (me *Basis) MultiplyVector3(right Vector3) Vector3 {
  panic("TODO: implement")
}

func (me *Basis) EqualsBasis(right Basis) bool {
  panic("TODO: implement")
}

func (me *Basis) NotEqualsBasis(right Basis) bool {
  panic("TODO: implement")
}

func (me *Basis) MultiplyBasis(right Basis) Basis {
  panic("TODO: implement")
}

func (me *Basis) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *Basis) InArray(right Array) bool {
  panic("TODO: implement")
}

// TODO: members (bclass)
