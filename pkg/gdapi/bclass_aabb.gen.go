// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AABB struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Enums

// Constructors

func NewAABB() AABB {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeAABB))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeAABB, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return AABB{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewAABBFromAABB(from AABB, ) AABB {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeAABB))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeAABB, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return AABB{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewAABBFromVector3Vector3(position Vector3, size Vector3, ) AABB {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeAABB))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeAABB, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{position.AsCTypePtr(), size.AsCTypePtr(), }))
  return AABB{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *AABB) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *AABB) Type() gdc.VariantType {
  return gdc.VariantTypeAABB
}

func (me *AABB) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *AABB) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func  (me *AABB) Abs() AABB {
  panic("TODO: implement")
}

func  (me *AABB) GetCenter() Vector3 {
  panic("TODO: implement")
}

func  (me *AABB) GetVolume() float32 {
  panic("TODO: implement")
}

func  (me *AABB) HasVolume() bool {
  panic("TODO: implement")
}

func  (me *AABB) HasSurface() bool {
  panic("TODO: implement")
}

func  (me *AABB) HasPoint(point Vector3, ) bool {
  panic("TODO: implement")
}

func  (me *AABB) IsEqualApprox(aabb AABB, ) bool {
  panic("TODO: implement")
}

func  (me *AABB) IsFinite() bool {
  panic("TODO: implement")
}

func  (me *AABB) Intersects(with AABB, ) bool {
  panic("TODO: implement")
}

func  (me *AABB) Encloses(with AABB, ) bool {
  panic("TODO: implement")
}

func  (me *AABB) IntersectsPlane(plane Plane, ) bool {
  panic("TODO: implement")
}

func  (me *AABB) Intersection(with AABB, ) AABB {
  panic("TODO: implement")
}

func  (me *AABB) Merge(with AABB, ) AABB {
  panic("TODO: implement")
}

func  (me *AABB) Expand(to_point Vector3, ) AABB {
  panic("TODO: implement")
}

func  (me *AABB) Grow(by float32, ) AABB {
  panic("TODO: implement")
}

func  (me *AABB) GetSupport(dir Vector3, ) Vector3 {
  panic("TODO: implement")
}

func  (me *AABB) GetLongestAxis() Vector3 {
  panic("TODO: implement")
}

func  (me *AABB) GetLongestAxisIndex() int {
  panic("TODO: implement")
}

func  (me *AABB) GetLongestAxisSize() float32 {
  panic("TODO: implement")
}

func  (me *AABB) GetShortestAxis() Vector3 {
  panic("TODO: implement")
}

func  (me *AABB) GetShortestAxisIndex() int {
  panic("TODO: implement")
}

func  (me *AABB) GetShortestAxisSize() float32 {
  panic("TODO: implement")
}

func  (me *AABB) GetEndpoint(idx int, ) Vector3 {
  panic("TODO: implement")
}

func  (me *AABB) IntersectsSegment(from Vector3, to Vector3, ) Variant {
  panic("TODO: implement")
}

func  (me *AABB) IntersectsRay(from Vector3, dir Vector3, ) Variant {
  panic("TODO: implement")
}

// Operators

func (me *AABB) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *AABB) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *AABB) Not() bool {
  panic("TODO: implement")
}

func (me *AABB) EqualsAABB(right AABB) bool {
  panic("TODO: implement")
}

func (me *AABB) NotEqualsAABB(right AABB) bool {
  panic("TODO: implement")
}

func (me *AABB) MultiplyTransform3D(right Transform3D) AABB {
  panic("TODO: implement")
}

func (me *AABB) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *AABB) InArray(right Array) bool {
  panic("TODO: implement")
}

// TODO: members (bclass)
