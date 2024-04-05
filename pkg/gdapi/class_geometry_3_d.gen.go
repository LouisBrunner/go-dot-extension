// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForGeometry3DList struct {
  fnComputeConvexMeshPoints gdc.MethodBindPtr
  fnBuildBoxPlanes gdc.MethodBindPtr
  fnBuildCylinderPlanes gdc.MethodBindPtr
  fnBuildCapsulePlanes gdc.MethodBindPtr
  fnGetClosestPointsBetweenSegments gdc.MethodBindPtr
  fnGetClosestPointToSegment gdc.MethodBindPtr
  fnGetClosestPointToSegmentUncapped gdc.MethodBindPtr
  fnGetTriangleBarycentricCoords gdc.MethodBindPtr
  fnRayIntersectsTriangle gdc.MethodBindPtr
  fnSegmentIntersectsTriangle gdc.MethodBindPtr
  fnSegmentIntersectsSphere gdc.MethodBindPtr
  fnSegmentIntersectsCylinder gdc.MethodBindPtr
  fnSegmentIntersectsConvex gdc.MethodBindPtr
  fnClipPolygon gdc.MethodBindPtr
}

var ptrsForGeometry3D ptrsForGeometry3DList

func initGeometry3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("Geometry3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("compute_convex_mesh_points")
    defer methodName.Destroy()
    ptrsForGeometry3D.fnComputeConvexMeshPoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1936902142))
  }
  {
    methodName := StringNameFromStr("build_box_planes")
    defer methodName.Destroy()
    ptrsForGeometry3D.fnBuildBoxPlanes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3622277145))
  }
  {
    methodName := StringNameFromStr("build_cylinder_planes")
    defer methodName.Destroy()
    ptrsForGeometry3D.fnBuildCylinderPlanes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 449920067))
  }
  {
    methodName := StringNameFromStr("build_capsule_planes")
    defer methodName.Destroy()
    ptrsForGeometry3D.fnBuildCapsulePlanes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2113592876))
  }
  {
    methodName := StringNameFromStr("get_closest_points_between_segments")
    defer methodName.Destroy()
    ptrsForGeometry3D.fnGetClosestPointsBetweenSegments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1056373962))
  }
  {
    methodName := StringNameFromStr("get_closest_point_to_segment")
    defer methodName.Destroy()
    ptrsForGeometry3D.fnGetClosestPointToSegment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2168193209))
  }
  {
    methodName := StringNameFromStr("get_closest_point_to_segment_uncapped")
    defer methodName.Destroy()
    ptrsForGeometry3D.fnGetClosestPointToSegmentUncapped = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2168193209))
  }
  {
    methodName := StringNameFromStr("get_triangle_barycentric_coords")
    defer methodName.Destroy()
    ptrsForGeometry3D.fnGetTriangleBarycentricCoords = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1362048029))
  }
  {
    methodName := StringNameFromStr("ray_intersects_triangle")
    defer methodName.Destroy()
    ptrsForGeometry3D.fnRayIntersectsTriangle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1718655448))
  }
  {
    methodName := StringNameFromStr("segment_intersects_triangle")
    defer methodName.Destroy()
    ptrsForGeometry3D.fnSegmentIntersectsTriangle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1718655448))
  }
  {
    methodName := StringNameFromStr("segment_intersects_sphere")
    defer methodName.Destroy()
    ptrsForGeometry3D.fnSegmentIntersectsSphere = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4080141172))
  }
  {
    methodName := StringNameFromStr("segment_intersects_cylinder")
    defer methodName.Destroy()
    ptrsForGeometry3D.fnSegmentIntersectsCylinder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2361316491))
  }
  {
    methodName := StringNameFromStr("segment_intersects_convex")
    defer methodName.Destroy()
    ptrsForGeometry3D.fnSegmentIntersectsConvex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 537425332))
  }
  {
    methodName := StringNameFromStr("clip_polygon")
    defer methodName.Destroy()
    ptrsForGeometry3D.fnClipPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2603188319))
  }
}

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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&planes) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector3Array()
  pinner.Pin(&planes)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry3D.fnComputeConvexMeshPoints), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry3D) BuildBoxPlanes(extents Vector3, ) []Plane {
  cargs := []gdc.ConstTypePtr{extents.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry3D.fnBuildBoxPlanes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Plane](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *Geometry3D) BuildCylinderPlanes(radius float64, height float64, sides int64, axis Vector3Axis, ) []Plane {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius) , gdc.ConstTypePtr(&height) , gdc.ConstTypePtr(&sides) , gdc.ConstTypePtr(&axis) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()
  pinner.Pin(&radius)
  pinner.Pin(&height)
  pinner.Pin(&sides)
  pinner.Pin(&axis)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry3D.fnBuildCylinderPlanes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Plane](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *Geometry3D) BuildCapsulePlanes(radius float64, height float64, sides int64, lats int64, axis Vector3Axis, ) []Plane {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius) , gdc.ConstTypePtr(&height) , gdc.ConstTypePtr(&sides) , gdc.ConstTypePtr(&lats) , gdc.ConstTypePtr(&axis) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()
  pinner.Pin(&radius)
  pinner.Pin(&height)
  pinner.Pin(&sides)
  pinner.Pin(&lats)
  pinner.Pin(&axis)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry3D.fnBuildCapsulePlanes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Plane](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *Geometry3D) GetClosestPointsBetweenSegments(p1 Vector3, p2 Vector3, q1 Vector3, q2 Vector3, ) PackedVector3Array {
  cargs := []gdc.ConstTypePtr{p1.AsCTypePtr(), p2.AsCTypePtr(), q1.AsCTypePtr(), q2.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector3Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry3D.fnGetClosestPointsBetweenSegments), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry3D) GetClosestPointToSegment(point Vector3, s1 Vector3, s2 Vector3, ) Vector3 {
  cargs := []gdc.ConstTypePtr{point.AsCTypePtr(), s1.AsCTypePtr(), s2.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry3D.fnGetClosestPointToSegment), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry3D) GetClosestPointToSegmentUncapped(point Vector3, s1 Vector3, s2 Vector3, ) Vector3 {
  cargs := []gdc.ConstTypePtr{point.AsCTypePtr(), s1.AsCTypePtr(), s2.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry3D.fnGetClosestPointToSegmentUncapped), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry3D) GetTriangleBarycentricCoords(point Vector3, a Vector3, b Vector3, c Vector3, ) Vector3 {
  cargs := []gdc.ConstTypePtr{point.AsCTypePtr(), a.AsCTypePtr(), b.AsCTypePtr(), c.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry3D.fnGetTriangleBarycentricCoords), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry3D) RayIntersectsTriangle(from Vector3, dir Vector3, a Vector3, b Vector3, c Vector3, ) Variant {
  cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), dir.AsCTypePtr(), a.AsCTypePtr(), b.AsCTypePtr(), c.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry3D.fnRayIntersectsTriangle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry3D) SegmentIntersectsTriangle(from Vector3, to Vector3, a Vector3, b Vector3, c Vector3, ) Variant {
  cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr(), a.AsCTypePtr(), b.AsCTypePtr(), c.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry3D.fnSegmentIntersectsTriangle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry3D) SegmentIntersectsSphere(from Vector3, to Vector3, sphere_position Vector3, sphere_radius float64, ) PackedVector3Array {
  cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr(), sphere_position.AsCTypePtr(), gdc.ConstTypePtr(&sphere_radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector3Array()
  pinner.Pin(&sphere_radius)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry3D.fnSegmentIntersectsSphere), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry3D) SegmentIntersectsCylinder(from Vector3, to Vector3, height float64, radius float64, ) PackedVector3Array {
  cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr(), gdc.ConstTypePtr(&height) , gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector3Array()
  pinner.Pin(&height)
  pinner.Pin(&radius)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry3D.fnSegmentIntersectsCylinder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry3D) SegmentIntersectsConvex(from Vector3, to Vector3, planes []Plane, ) PackedVector3Array {
  cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr(), gdc.ConstTypePtr(&planes) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector3Array()
  pinner.Pin(&planes)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry3D.fnSegmentIntersectsConvex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry3D) ClipPolygon(points PackedVector3Array, plane Plane, ) PackedVector3Array {
  cargs := []gdc.ConstTypePtr{points.AsCTypePtr(), plane.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector3Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry3D.fnClipPolygon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
