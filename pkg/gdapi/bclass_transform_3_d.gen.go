// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Transform3D struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Constants

var (
  Transform3DIdentity = "Transform3D(1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0)" // TODO: construct correctly
  Transform3DFlipX = "Transform3D(-1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0)" // TODO: construct correctly
  Transform3DFlipY = "Transform3D(1, 0, 0, 0, -1, 0, 0, 0, 1, 0, 0, 0)" // TODO: construct correctly
  Transform3DFlipZ = "Transform3D(1, 0, 0, 0, 1, 0, 0, 0, -1, 0, 0, 0)" // TODO: construct correctly
)

// Enums

// Constructors

func NewTransform3D() Transform3D {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeTransform3D))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeTransform3D, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Transform3D{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewTransform3DFromTransform3D(from Transform3D, ) Transform3D {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeTransform3D))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeTransform3D, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Transform3D{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewTransform3DFromBasisVector3(basis Basis, origin Vector3, ) Transform3D {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeTransform3D))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeTransform3D, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{basis.AsCTypePtr(), origin.AsCTypePtr(), }))
  return Transform3D{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewTransform3DFromVector3Vector3Vector3Vector3(x_axis Vector3, y_axis Vector3, z_axis Vector3, origin Vector3, ) Transform3D {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeTransform3D))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeTransform3D, 3) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{x_axis.AsCTypePtr(), y_axis.AsCTypePtr(), z_axis.AsCTypePtr(), origin.AsCTypePtr(), }))
  return Transform3D{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewTransform3DFromProjection(from Projection, ) Transform3D {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeTransform3D))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeTransform3D, 4) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Transform3D{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Transform3D) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Transform3D) Type() gdc.VariantType {
  return gdc.VariantTypeTransform3D
}

func (me *Transform3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Transform3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func  (me *Transform3D) Inverse() Transform3D {
  panic("TODO: implement")
}

func  (me *Transform3D) AffineInverse() Transform3D {
  panic("TODO: implement")
}

func  (me *Transform3D) Orthonormalized() Transform3D {
  panic("TODO: implement")
}

func  (me *Transform3D) Rotated(axis Vector3, angle float32, ) Transform3D {
  panic("TODO: implement")
}

func  (me *Transform3D) RotatedLocal(axis Vector3, angle float32, ) Transform3D {
  panic("TODO: implement")
}

func  (me *Transform3D) Scaled(scale Vector3, ) Transform3D {
  panic("TODO: implement")
}

func  (me *Transform3D) ScaledLocal(scale Vector3, ) Transform3D {
  panic("TODO: implement")
}

func  (me *Transform3D) Translated(offset Vector3, ) Transform3D {
  panic("TODO: implement")
}

func  (me *Transform3D) TranslatedLocal(offset Vector3, ) Transform3D {
  panic("TODO: implement")
}

func  (me *Transform3D) LookingAt(target Vector3, up Vector3, use_model_front bool, ) Transform3D {
  panic("TODO: implement")
}

func  (me *Transform3D) InterpolateWith(xform Transform3D, weight float32, ) Transform3D {
  panic("TODO: implement")
}

func  (me *Transform3D) IsEqualApprox(xform Transform3D, ) bool {
  panic("TODO: implement")
}

func  (me *Transform3D) IsFinite() bool {
  panic("TODO: implement")
}

// Operators

func (me *Transform3D) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Transform3D) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Transform3D) Not() bool {
  panic("TODO: implement")
}

func (me *Transform3D) MultiplyInt(right int) Transform3D {
  panic("TODO: implement")
}

func (me *Transform3D) MultiplyFloat32(right float32) Transform3D {
  panic("TODO: implement")
}

func (me *Transform3D) MultiplyVector3(right Vector3) Vector3 {
  panic("TODO: implement")
}

func (me *Transform3D) MultiplyPlane(right Plane) Plane {
  panic("TODO: implement")
}

func (me *Transform3D) MultiplyAABB(right AABB) AABB {
  panic("TODO: implement")
}

func (me *Transform3D) EqualsTransform3D(right Transform3D) bool {
  panic("TODO: implement")
}

func (me *Transform3D) NotEqualsTransform3D(right Transform3D) bool {
  panic("TODO: implement")
}

func (me *Transform3D) MultiplyTransform3D(right Transform3D) Transform3D {
  panic("TODO: implement")
}

func (me *Transform3D) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *Transform3D) InArray(right Array) bool {
  panic("TODO: implement")
}

func (me *Transform3D) MultiplyPackedVector3Array(right PackedVector3Array) PackedVector3Array {
  panic("TODO: implement")
}

// TODO: members (bclass)
