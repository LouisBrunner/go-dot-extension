// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Geometry2D struct {
  obj gdc.ObjectPtr
}

func (me *Geometry2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Geometry2D) BaseClass() string {
  return "Geometry2D"
}

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

func  (me *Geometry2D) IsPointInCircle(point Vector2, circle_position Vector2, circle_radius float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Geometry2D) SegmentIntersectsCircle(segment_from Vector2, segment_to Vector2, circle_position Vector2, circle_radius float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Geometry2D) SegmentIntersectsSegment(from_a Vector2, to_a Vector2, from_b Vector2, to_b Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Geometry2D) LineIntersectsLine(from_a Vector2, dir_a Vector2, from_b Vector2, dir_b Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Geometry2D) GetClosestPointsBetweenSegments(p1 Vector2, q1 Vector2, p2 Vector2, q2 Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Geometry2D) GetClosestPointToSegment(point Vector2, s1 Vector2, s2 Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Geometry2D) GetClosestPointToSegmentUncapped(point Vector2, s1 Vector2, s2 Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Geometry2D) PointIsInsideTriangle(point Vector2, a Vector2, b Vector2, c Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Geometry2D) IsPolygonClockwise(polygon PackedVector2Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *Geometry2D) IsPointInPolygon(point Vector2, polygon PackedVector2Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *Geometry2D) TriangulatePolygon(polygon PackedVector2Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *Geometry2D) TriangulateDelaunay(points PackedVector2Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *Geometry2D) ConvexHull(points PackedVector2Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *Geometry2D) DecomposePolygonInConvex(polygon PackedVector2Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *Geometry2D) MergePolygons(polygon_a PackedVector2Array, polygon_b PackedVector2Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *Geometry2D) ClipPolygons(polygon_a PackedVector2Array, polygon_b PackedVector2Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *Geometry2D) IntersectPolygons(polygon_a PackedVector2Array, polygon_b PackedVector2Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *Geometry2D) ExcludePolygons(polygon_a PackedVector2Array, polygon_b PackedVector2Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *Geometry2D) ClipPolylineWithPolygon(polyline PackedVector2Array, polygon PackedVector2Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *Geometry2D) IntersectPolylineWithPolygon(polyline PackedVector2Array, polygon PackedVector2Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *Geometry2D) OffsetPolygon(polygon PackedVector2Array, delta float32, join_type Geometry2DPolyJoinType, ) { // TODO: return value
  // TODO: implement
}

func  (me *Geometry2D) OffsetPolyline(polyline PackedVector2Array, delta float32, join_type Geometry2DPolyJoinType, end_type Geometry2DPolyEndType, ) { // TODO: return value
  // TODO: implement
}

func  (me *Geometry2D) MakeAtlas(sizes PackedVector2Array, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
