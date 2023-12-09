// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Geometry3D struct {
  obj gdc.ObjectPtr
}

func (me *Geometry3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Geometry3D) BaseClass() string {
  return "Geometry3D"
}



// Enums

func (me *Geometry3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Geometry3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Geometry3D) BuildBoxPlanes(extents Vector3, )  {
  panic("TODO: implement")
}

func  (me *Geometry3D) BuildCylinderPlanes(radius float32, height float32, sides int, axis Vector3Axis, )  {
  panic("TODO: implement")
}

func  (me *Geometry3D) BuildCapsulePlanes(radius float32, height float32, sides int, lats int, axis Vector3Axis, )  {
  panic("TODO: implement")
}

func  (me *Geometry3D) GetClosestPointsBetweenSegments(p1 Vector3, p2 Vector3, q1 Vector3, q2 Vector3, )  {
  panic("TODO: implement")
}

func  (me *Geometry3D) GetClosestPointToSegment(point Vector3, s1 Vector3, s2 Vector3, )  {
  panic("TODO: implement")
}

func  (me *Geometry3D) GetClosestPointToSegmentUncapped(point Vector3, s1 Vector3, s2 Vector3, )  {
  panic("TODO: implement")
}

func  (me *Geometry3D) RayIntersectsTriangle(from Vector3, dir Vector3, a Vector3, b Vector3, c Vector3, )  {
  panic("TODO: implement")
}

func  (me *Geometry3D) SegmentIntersectsTriangle(from Vector3, to Vector3, a Vector3, b Vector3, c Vector3, )  {
  panic("TODO: implement")
}

func  (me *Geometry3D) SegmentIntersectsSphere(from Vector3, to Vector3, sphere_position Vector3, sphere_radius float32, )  {
  panic("TODO: implement")
}

func  (me *Geometry3D) SegmentIntersectsCylinder(from Vector3, to Vector3, height float32, radius float32, )  {
  panic("TODO: implement")
}

func  (me *Geometry3D) SegmentIntersectsConvex(from Vector3, to Vector3, planes Plane, )  {
  panic("TODO: implement")
}

func  (me *Geometry3D) ClipPolygon(points PackedVector3Array, plane Plane, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
