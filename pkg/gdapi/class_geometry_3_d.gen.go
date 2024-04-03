// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Geometry3D struct {
  Object
}

func (me *Geometry3D) BaseClass() string {
  return "Geometry3D"
}

func NewGeometry3D() *Geometry3D {
  str := StringNameFromStr("Geometry3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Geometry3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *Geometry3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Geometry3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Geometry3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Geometry3D) ComputeConvexMeshPoints(planes []Plane, ) PackedVector3Array {
  classNameV := StringNameFromStr("Geometry3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("compute_convex_mesh_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1936902142) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&planes), }
  ret := NewPackedVector3Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry3D) BuildBoxPlanes(extents Vector3, ) []Plane {
  classNameV := StringNameFromStr("Geometry3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("build_box_planes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3622277145) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(extents.AsCTypePtr()), }
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[Plane](ret)
}

func  (me *Geometry3D) BuildCylinderPlanes(radius float64, height float64, sides int64, axis Vector3Axis, ) []Plane {
  classNameV := StringNameFromStr("Geometry3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("build_cylinder_planes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 449920067) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius), gdc.ConstTypePtr(&height), gdc.ConstTypePtr(&sides), gdc.ConstTypePtr(&axis), }
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[Plane](ret)
}

func  (me *Geometry3D) BuildCapsulePlanes(radius float64, height float64, sides int64, lats int64, axis Vector3Axis, ) []Plane {
  classNameV := StringNameFromStr("Geometry3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("build_capsule_planes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2113592876) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius), gdc.ConstTypePtr(&height), gdc.ConstTypePtr(&sides), gdc.ConstTypePtr(&lats), gdc.ConstTypePtr(&axis), }
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[Plane](ret)
}

func  (me *Geometry3D) GetClosestPointsBetweenSegments(p1 Vector3, p2 Vector3, q1 Vector3, q2 Vector3, ) PackedVector3Array {
  classNameV := StringNameFromStr("Geometry3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_closest_points_between_segments")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1056373962) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(p1.AsCTypePtr()), gdc.ConstTypePtr(p2.AsCTypePtr()), gdc.ConstTypePtr(q1.AsCTypePtr()), gdc.ConstTypePtr(q2.AsCTypePtr()), }
  ret := NewPackedVector3Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry3D) GetClosestPointToSegment(point Vector3, s1 Vector3, s2 Vector3, ) Vector3 {
  classNameV := StringNameFromStr("Geometry3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_closest_point_to_segment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2168193209) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(point.AsCTypePtr()), gdc.ConstTypePtr(s1.AsCTypePtr()), gdc.ConstTypePtr(s2.AsCTypePtr()), }
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry3D) GetClosestPointToSegmentUncapped(point Vector3, s1 Vector3, s2 Vector3, ) Vector3 {
  classNameV := StringNameFromStr("Geometry3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_closest_point_to_segment_uncapped")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2168193209) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(point.AsCTypePtr()), gdc.ConstTypePtr(s1.AsCTypePtr()), gdc.ConstTypePtr(s2.AsCTypePtr()), }
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry3D) GetTriangleBarycentricCoords(point Vector3, a Vector3, b Vector3, c Vector3, ) Vector3 {
  classNameV := StringNameFromStr("Geometry3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_triangle_barycentric_coords")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1362048029) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(point.AsCTypePtr()), gdc.ConstTypePtr(a.AsCTypePtr()), gdc.ConstTypePtr(b.AsCTypePtr()), gdc.ConstTypePtr(c.AsCTypePtr()), }
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry3D) RayIntersectsTriangle(from Vector3, dir Vector3, a Vector3, b Vector3, c Vector3, ) Variant {
  classNameV := StringNameFromStr("Geometry3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("ray_intersects_triangle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1718655448) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from.AsCTypePtr()), gdc.ConstTypePtr(dir.AsCTypePtr()), gdc.ConstTypePtr(a.AsCTypePtr()), gdc.ConstTypePtr(b.AsCTypePtr()), gdc.ConstTypePtr(c.AsCTypePtr()), }
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry3D) SegmentIntersectsTriangle(from Vector3, to Vector3, a Vector3, b Vector3, c Vector3, ) Variant {
  classNameV := StringNameFromStr("Geometry3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("segment_intersects_triangle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1718655448) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from.AsCTypePtr()), gdc.ConstTypePtr(to.AsCTypePtr()), gdc.ConstTypePtr(a.AsCTypePtr()), gdc.ConstTypePtr(b.AsCTypePtr()), gdc.ConstTypePtr(c.AsCTypePtr()), }
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry3D) SegmentIntersectsSphere(from Vector3, to Vector3, sphere_position Vector3, sphere_radius float64, ) PackedVector3Array {
  classNameV := StringNameFromStr("Geometry3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("segment_intersects_sphere")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4080141172) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from.AsCTypePtr()), gdc.ConstTypePtr(to.AsCTypePtr()), gdc.ConstTypePtr(sphere_position.AsCTypePtr()), gdc.ConstTypePtr(&sphere_radius), }
  ret := NewPackedVector3Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry3D) SegmentIntersectsCylinder(from Vector3, to Vector3, height float64, radius float64, ) PackedVector3Array {
  classNameV := StringNameFromStr("Geometry3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("segment_intersects_cylinder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2361316491) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from.AsCTypePtr()), gdc.ConstTypePtr(to.AsCTypePtr()), gdc.ConstTypePtr(&height), gdc.ConstTypePtr(&radius), }
  ret := NewPackedVector3Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry3D) SegmentIntersectsConvex(from Vector3, to Vector3, planes []Plane, ) PackedVector3Array {
  classNameV := StringNameFromStr("Geometry3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("segment_intersects_convex")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 537425332) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from.AsCTypePtr()), gdc.ConstTypePtr(to.AsCTypePtr()), gdc.ConstTypePtr(&planes), }
  ret := NewPackedVector3Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry3D) ClipPolygon(points PackedVector3Array, plane Plane, ) PackedVector3Array {
  classNameV := StringNameFromStr("Geometry3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clip_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2603188319) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(points.AsCTypePtr()), gdc.ConstTypePtr(plane.AsCTypePtr()), }
  ret := NewPackedVector3Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
