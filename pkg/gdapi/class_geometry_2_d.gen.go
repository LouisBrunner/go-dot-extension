// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type Geometry2D struct {
  Object
}

func (me *Geometry2D) BaseClass() string {
  return "Geometry2D"
}

func NewGeometry2D() *Geometry2D {
  str := StringNameFromStr("Geometry2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Geometry2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type Geometry2DPolyBooleanOperation int
const (
  Geometry2DPolyBooleanOperationOperationUnion Geometry2DPolyBooleanOperation = 0
  Geometry2DPolyBooleanOperationOperationDifference Geometry2DPolyBooleanOperation = 1
  Geometry2DPolyBooleanOperationOperationIntersection Geometry2DPolyBooleanOperation = 2
  Geometry2DPolyBooleanOperationOperationXor Geometry2DPolyBooleanOperation = 3
)

type Geometry2DPolyJoinType int
const (
  Geometry2DPolyJoinTypeJoinSquare Geometry2DPolyJoinType = 0
  Geometry2DPolyJoinTypeJoinRound Geometry2DPolyJoinType = 1
  Geometry2DPolyJoinTypeJoinMiter Geometry2DPolyJoinType = 2
)

type Geometry2DPolyEndType int
const (
  Geometry2DPolyEndTypeEndPolygon Geometry2DPolyEndType = 0
  Geometry2DPolyEndTypeEndJoined Geometry2DPolyEndType = 1
  Geometry2DPolyEndTypeEndButt Geometry2DPolyEndType = 2
  Geometry2DPolyEndTypeEndSquare Geometry2DPolyEndType = 3
  Geometry2DPolyEndTypeEndRound Geometry2DPolyEndType = 4
)

func (me *Geometry2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Geometry2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Geometry2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Geometry2D) IsPointInCircle(point Vector2, circle_position Vector2, circle_radius float64, ) bool {
  classNameV := StringNameFromStr("Geometry2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_point_in_circle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2929491703) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{point.AsCTypePtr(), circle_position.AsCTypePtr(), gdc.ConstTypePtr(&circle_radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&circle_radius)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Geometry2D) SegmentIntersectsCircle(segment_from Vector2, segment_to Vector2, circle_position Vector2, circle_radius float64, ) float64 {
  classNameV := StringNameFromStr("Geometry2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("segment_intersects_circle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1356928167) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{segment_from.AsCTypePtr(), segment_to.AsCTypePtr(), circle_position.AsCTypePtr(), gdc.ConstTypePtr(&circle_radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&circle_radius)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Geometry2D) SegmentIntersectsSegment(from_a Vector2, to_a Vector2, from_b Vector2, to_b Vector2, ) Variant {
  classNameV := StringNameFromStr("Geometry2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("segment_intersects_segment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2058025344) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{from_a.AsCTypePtr(), to_a.AsCTypePtr(), from_b.AsCTypePtr(), to_b.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry2D) LineIntersectsLine(from_a Vector2, dir_a Vector2, from_b Vector2, dir_b Vector2, ) Variant {
  classNameV := StringNameFromStr("Geometry2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("line_intersects_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2058025344) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{from_a.AsCTypePtr(), dir_a.AsCTypePtr(), from_b.AsCTypePtr(), dir_b.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry2D) GetClosestPointsBetweenSegments(p1 Vector2, q1 Vector2, p2 Vector2, q2 Vector2, ) PackedVector2Array {
  classNameV := StringNameFromStr("Geometry2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_closest_points_between_segments")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3344690961) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{p1.AsCTypePtr(), q1.AsCTypePtr(), p2.AsCTypePtr(), q2.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry2D) GetClosestPointToSegment(point Vector2, s1 Vector2, s2 Vector2, ) Vector2 {
  classNameV := StringNameFromStr("Geometry2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_closest_point_to_segment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4172901909) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{point.AsCTypePtr(), s1.AsCTypePtr(), s2.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry2D) GetClosestPointToSegmentUncapped(point Vector2, s1 Vector2, s2 Vector2, ) Vector2 {
  classNameV := StringNameFromStr("Geometry2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_closest_point_to_segment_uncapped")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4172901909) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{point.AsCTypePtr(), s1.AsCTypePtr(), s2.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry2D) PointIsInsideTriangle(point Vector2, a Vector2, b Vector2, c Vector2, ) bool {
  classNameV := StringNameFromStr("Geometry2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("point_is_inside_triangle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1025948137) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{point.AsCTypePtr(), a.AsCTypePtr(), b.AsCTypePtr(), c.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Geometry2D) IsPolygonClockwise(polygon PackedVector2Array, ) bool {
  classNameV := StringNameFromStr("Geometry2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_polygon_clockwise")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1361156557) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{polygon.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Geometry2D) IsPointInPolygon(point Vector2, polygon PackedVector2Array, ) bool {
  classNameV := StringNameFromStr("Geometry2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_point_in_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 738277916) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{point.AsCTypePtr(), polygon.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Geometry2D) TriangulatePolygon(polygon PackedVector2Array, ) PackedInt32Array {
  classNameV := StringNameFromStr("Geometry2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("triangulate_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1389921771) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{polygon.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry2D) TriangulateDelaunay(points PackedVector2Array, ) PackedInt32Array {
  classNameV := StringNameFromStr("Geometry2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("triangulate_delaunay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1389921771) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{points.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry2D) ConvexHull(points PackedVector2Array, ) PackedVector2Array {
  classNameV := StringNameFromStr("Geometry2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("convex_hull")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2004331998) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{points.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Geometry2D) DecomposePolygonInConvex(polygon PackedVector2Array, ) []PackedVector2Array {
  classNameV := StringNameFromStr("Geometry2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("decompose_polygon_in_convex")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3982393695) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{polygon.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[PackedVector2Array](ret)
}

func  (me *Geometry2D) MergePolygons(polygon_a PackedVector2Array, polygon_b PackedVector2Array, ) []PackedVector2Array {
  classNameV := StringNameFromStr("Geometry2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("merge_polygons")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3637387053) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{polygon_a.AsCTypePtr(), polygon_b.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[PackedVector2Array](ret)
}

func  (me *Geometry2D) ClipPolygons(polygon_a PackedVector2Array, polygon_b PackedVector2Array, ) []PackedVector2Array {
  classNameV := StringNameFromStr("Geometry2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clip_polygons")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3637387053) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{polygon_a.AsCTypePtr(), polygon_b.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[PackedVector2Array](ret)
}

func  (me *Geometry2D) IntersectPolygons(polygon_a PackedVector2Array, polygon_b PackedVector2Array, ) []PackedVector2Array {
  classNameV := StringNameFromStr("Geometry2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("intersect_polygons")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3637387053) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{polygon_a.AsCTypePtr(), polygon_b.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[PackedVector2Array](ret)
}

func  (me *Geometry2D) ExcludePolygons(polygon_a PackedVector2Array, polygon_b PackedVector2Array, ) []PackedVector2Array {
  classNameV := StringNameFromStr("Geometry2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("exclude_polygons")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3637387053) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{polygon_a.AsCTypePtr(), polygon_b.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[PackedVector2Array](ret)
}

func  (me *Geometry2D) ClipPolylineWithPolygon(polyline PackedVector2Array, polygon PackedVector2Array, ) []PackedVector2Array {
  classNameV := StringNameFromStr("Geometry2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clip_polyline_with_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3637387053) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{polyline.AsCTypePtr(), polygon.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[PackedVector2Array](ret)
}

func  (me *Geometry2D) IntersectPolylineWithPolygon(polyline PackedVector2Array, polygon PackedVector2Array, ) []PackedVector2Array {
  classNameV := StringNameFromStr("Geometry2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("intersect_polyline_with_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3637387053) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{polyline.AsCTypePtr(), polygon.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[PackedVector2Array](ret)
}

func  (me *Geometry2D) OffsetPolygon(polygon PackedVector2Array, delta float64, join_type Geometry2DPolyJoinType, ) []PackedVector2Array {
  classNameV := StringNameFromStr("Geometry2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("offset_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1275354010) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{polygon.AsCTypePtr(), gdc.ConstTypePtr(&delta) , gdc.ConstTypePtr(&join_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()
  pinner.Pin(&delta)
  pinner.Pin(&join_type)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[PackedVector2Array](ret)
}

func  (me *Geometry2D) OffsetPolyline(polyline PackedVector2Array, delta float64, join_type Geometry2DPolyJoinType, end_type Geometry2DPolyEndType, ) []PackedVector2Array {
  classNameV := StringNameFromStr("Geometry2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("offset_polyline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2328231778) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{polyline.AsCTypePtr(), gdc.ConstTypePtr(&delta) , gdc.ConstTypePtr(&join_type) , gdc.ConstTypePtr(&end_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()
  pinner.Pin(&delta)
  pinner.Pin(&join_type)
  pinner.Pin(&end_type)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[PackedVector2Array](ret)
}

func  (me *Geometry2D) MakeAtlas(sizes PackedVector2Array, ) Dictionary {
  classNameV := StringNameFromStr("Geometry2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_atlas")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1337682371) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{sizes.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
