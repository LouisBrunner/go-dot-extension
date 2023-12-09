// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Plane struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Constants

var (
  PlanePlaneYz = "Plane(1, 0, 0, 0)" // TODO: construct correctly
  PlanePlaneXz = "Plane(0, 1, 0, 0)" // TODO: construct correctly
  PlanePlaneXy = "Plane(0, 0, 1, 0)" // TODO: construct correctly
)

// Enums

// Constructors

func NewPlane() Plane {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePlane))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePlane, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Plane{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPlaneFromPlane(from Plane, ) Plane {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePlane))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePlane, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Plane{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPlaneFromVector3(normal Vector3, ) Plane {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePlane))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePlane, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{normal.AsCTypePtr(), }))
  return Plane{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPlaneFromVector3Float32(normal Vector3, d float32, ) Plane {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePlane))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePlane, 3) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{normal.AsCTypePtr(), gdc.ConstTypePtr(&d), }))
  return Plane{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPlaneFromVector3Vector3(normal Vector3, point Vector3, ) Plane {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePlane))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePlane, 4) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{normal.AsCTypePtr(), point.AsCTypePtr(), }))
  return Plane{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPlaneFromVector3Vector3Vector3(point1 Vector3, point2 Vector3, point3 Vector3, ) Plane {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePlane))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePlane, 5) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{point1.AsCTypePtr(), point2.AsCTypePtr(), point3.AsCTypePtr(), }))
  return Plane{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPlaneFromFloat32Float32Float32Float32(a float32, b float32, c float32, d float32, ) Plane {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePlane))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePlane, 6) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&a), gdc.ConstTypePtr(&b), gdc.ConstTypePtr(&c), gdc.ConstTypePtr(&d), }))
  return Plane{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Plane) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Plane) Type() gdc.VariantType {
  return gdc.VariantTypePlane
}

func (me *Plane) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Plane) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func  (me *Plane) Normalized() Plane {
  panic("TODO: implement")
}

func  (me *Plane) GetCenter() Vector3 {
  panic("TODO: implement")
}

func  (me *Plane) IsEqualApprox(to_plane Plane, ) bool {
  panic("TODO: implement")
}

func  (me *Plane) IsFinite() bool {
  panic("TODO: implement")
}

func  (me *Plane) IsPointOver(point Vector3, ) bool {
  panic("TODO: implement")
}

func  (me *Plane) DistanceTo(point Vector3, ) float32 {
  panic("TODO: implement")
}

func  (me *Plane) HasPoint(point Vector3, tolerance float32, ) bool {
  panic("TODO: implement")
}

func  (me *Plane) Project(point Vector3, ) Vector3 {
  panic("TODO: implement")
}

func  (me *Plane) Intersect3(b Plane, c Plane, ) Variant {
  panic("TODO: implement")
}

func  (me *Plane) IntersectsRay(from Vector3, dir Vector3, ) Variant {
  panic("TODO: implement")
}

func  (me *Plane) IntersectsSegment(from Vector3, to Vector3, ) Variant {
  panic("TODO: implement")
}

// Operators

func (me *Plane) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Plane) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Plane) UnaryMinus() Plane {
  panic("TODO: implement")
}

func (me *Plane) UnaryPlus() Plane {
  panic("TODO: implement")
}

func (me *Plane) Not() bool {
  panic("TODO: implement")
}

func (me *Plane) EqualsPlane(right Plane) bool {
  panic("TODO: implement")
}

func (me *Plane) NotEqualsPlane(right Plane) bool {
  panic("TODO: implement")
}

func (me *Plane) MultiplyTransform3D(right Transform3D) Plane {
  panic("TODO: implement")
}

func (me *Plane) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *Plane) InArray(right Array) bool {
  panic("TODO: implement")
}

// TODO: members (bclass)
