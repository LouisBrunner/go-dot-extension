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

type ptrsForGeometry2DList struct {
	fnIsPointInCircle                  gdc.MethodBindPtr
	fnSegmentIntersectsCircle          gdc.MethodBindPtr
	fnSegmentIntersectsSegment         gdc.MethodBindPtr
	fnLineIntersectsLine               gdc.MethodBindPtr
	fnGetClosestPointsBetweenSegments  gdc.MethodBindPtr
	fnGetClosestPointToSegment         gdc.MethodBindPtr
	fnGetClosestPointToSegmentUncapped gdc.MethodBindPtr
	fnPointIsInsideTriangle            gdc.MethodBindPtr
	fnIsPolygonClockwise               gdc.MethodBindPtr
	fnIsPointInPolygon                 gdc.MethodBindPtr
	fnTriangulatePolygon               gdc.MethodBindPtr
	fnTriangulateDelaunay              gdc.MethodBindPtr
	fnConvexHull                       gdc.MethodBindPtr
	fnDecomposePolygonInConvex         gdc.MethodBindPtr
	fnMergePolygons                    gdc.MethodBindPtr
	fnClipPolygons                     gdc.MethodBindPtr
	fnIntersectPolygons                gdc.MethodBindPtr
	fnExcludePolygons                  gdc.MethodBindPtr
	fnClipPolylineWithPolygon          gdc.MethodBindPtr
	fnIntersectPolylineWithPolygon     gdc.MethodBindPtr
	fnOffsetPolygon                    gdc.MethodBindPtr
	fnOffsetPolyline                   gdc.MethodBindPtr
	fnMakeAtlas                        gdc.MethodBindPtr
}

var ptrsForGeometry2D ptrsForGeometry2DList

func initGeometry2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Geometry2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("is_point_in_circle")
		defer methodName.Destroy()
		ptrsForGeometry2D.fnIsPointInCircle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2929491703))
	}
	{
		methodName := StringNameFromStr("segment_intersects_circle")
		defer methodName.Destroy()
		ptrsForGeometry2D.fnSegmentIntersectsCircle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1356928167))
	}
	{
		methodName := StringNameFromStr("segment_intersects_segment")
		defer methodName.Destroy()
		ptrsForGeometry2D.fnSegmentIntersectsSegment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2058025344))
	}
	{
		methodName := StringNameFromStr("line_intersects_line")
		defer methodName.Destroy()
		ptrsForGeometry2D.fnLineIntersectsLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2058025344))
	}
	{
		methodName := StringNameFromStr("get_closest_points_between_segments")
		defer methodName.Destroy()
		ptrsForGeometry2D.fnGetClosestPointsBetweenSegments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3344690961))
	}
	{
		methodName := StringNameFromStr("get_closest_point_to_segment")
		defer methodName.Destroy()
		ptrsForGeometry2D.fnGetClosestPointToSegment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4172901909))
	}
	{
		methodName := StringNameFromStr("get_closest_point_to_segment_uncapped")
		defer methodName.Destroy()
		ptrsForGeometry2D.fnGetClosestPointToSegmentUncapped = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4172901909))
	}
	{
		methodName := StringNameFromStr("point_is_inside_triangle")
		defer methodName.Destroy()
		ptrsForGeometry2D.fnPointIsInsideTriangle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1025948137))
	}
	{
		methodName := StringNameFromStr("is_polygon_clockwise")
		defer methodName.Destroy()
		ptrsForGeometry2D.fnIsPolygonClockwise = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1361156557))
	}
	{
		methodName := StringNameFromStr("is_point_in_polygon")
		defer methodName.Destroy()
		ptrsForGeometry2D.fnIsPointInPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 738277916))
	}
	{
		methodName := StringNameFromStr("triangulate_polygon")
		defer methodName.Destroy()
		ptrsForGeometry2D.fnTriangulatePolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1389921771))
	}
	{
		methodName := StringNameFromStr("triangulate_delaunay")
		defer methodName.Destroy()
		ptrsForGeometry2D.fnTriangulateDelaunay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1389921771))
	}
	{
		methodName := StringNameFromStr("convex_hull")
		defer methodName.Destroy()
		ptrsForGeometry2D.fnConvexHull = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2004331998))
	}
	{
		methodName := StringNameFromStr("decompose_polygon_in_convex")
		defer methodName.Destroy()
		ptrsForGeometry2D.fnDecomposePolygonInConvex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3982393695))
	}
	{
		methodName := StringNameFromStr("merge_polygons")
		defer methodName.Destroy()
		ptrsForGeometry2D.fnMergePolygons = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3637387053))
	}
	{
		methodName := StringNameFromStr("clip_polygons")
		defer methodName.Destroy()
		ptrsForGeometry2D.fnClipPolygons = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3637387053))
	}
	{
		methodName := StringNameFromStr("intersect_polygons")
		defer methodName.Destroy()
		ptrsForGeometry2D.fnIntersectPolygons = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3637387053))
	}
	{
		methodName := StringNameFromStr("exclude_polygons")
		defer methodName.Destroy()
		ptrsForGeometry2D.fnExcludePolygons = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3637387053))
	}
	{
		methodName := StringNameFromStr("clip_polyline_with_polygon")
		defer methodName.Destroy()
		ptrsForGeometry2D.fnClipPolylineWithPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3637387053))
	}
	{
		methodName := StringNameFromStr("intersect_polyline_with_polygon")
		defer methodName.Destroy()
		ptrsForGeometry2D.fnIntersectPolylineWithPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3637387053))
	}
	{
		methodName := StringNameFromStr("offset_polygon")
		defer methodName.Destroy()
		ptrsForGeometry2D.fnOffsetPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1275354010))
	}
	{
		methodName := StringNameFromStr("offset_polyline")
		defer methodName.Destroy()
		ptrsForGeometry2D.fnOffsetPolyline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2328231778))
	}
	{
		methodName := StringNameFromStr("make_atlas")
		defer methodName.Destroy()
		ptrsForGeometry2D.fnMakeAtlas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1337682371))
	}

}

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
	Geometry2DPolyBooleanOperationOperationUnion        Geometry2DPolyBooleanOperation = 0
	Geometry2DPolyBooleanOperationOperationDifference   Geometry2DPolyBooleanOperation = 1
	Geometry2DPolyBooleanOperationOperationIntersection Geometry2DPolyBooleanOperation = 2
	Geometry2DPolyBooleanOperationOperationXor          Geometry2DPolyBooleanOperation = 3
)

type Geometry2DPolyJoinType int

const (
	Geometry2DPolyJoinTypeJoinSquare Geometry2DPolyJoinType = 0
	Geometry2DPolyJoinTypeJoinRound  Geometry2DPolyJoinType = 1
	Geometry2DPolyJoinTypeJoinMiter  Geometry2DPolyJoinType = 2
)

type Geometry2DPolyEndType int

const (
	Geometry2DPolyEndTypeEndPolygon Geometry2DPolyEndType = 0
	Geometry2DPolyEndTypeEndJoined  Geometry2DPolyEndType = 1
	Geometry2DPolyEndTypeEndButt    Geometry2DPolyEndType = 2
	Geometry2DPolyEndTypeEndSquare  Geometry2DPolyEndType = 3
	Geometry2DPolyEndTypeEndRound   Geometry2DPolyEndType = 4
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

func (me *Geometry2D) IsPointInCircle(point Vector2, circle_position Vector2, circle_radius float64) bool {
	cargs := []gdc.ConstTypePtr{point.AsCTypePtr(), circle_position.AsCTypePtr(), gdc.ConstTypePtr(&circle_radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&circle_radius)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry2D.fnIsPointInCircle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Geometry2D) SegmentIntersectsCircle(segment_from Vector2, segment_to Vector2, circle_position Vector2, circle_radius float64) float64 {
	cargs := []gdc.ConstTypePtr{segment_from.AsCTypePtr(), segment_to.AsCTypePtr(), circle_position.AsCTypePtr(), gdc.ConstTypePtr(&circle_radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&circle_radius)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry2D.fnSegmentIntersectsCircle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Geometry2D) SegmentIntersectsSegment(from_a Vector2, to_a Vector2, from_b Vector2, to_b Vector2) Variant {
	cargs := []gdc.ConstTypePtr{from_a.AsCTypePtr(), to_a.AsCTypePtr(), from_b.AsCTypePtr(), to_b.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry2D.fnSegmentIntersectsSegment), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Geometry2D) LineIntersectsLine(from_a Vector2, dir_a Vector2, from_b Vector2, dir_b Vector2) Variant {
	cargs := []gdc.ConstTypePtr{from_a.AsCTypePtr(), dir_a.AsCTypePtr(), from_b.AsCTypePtr(), dir_b.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry2D.fnLineIntersectsLine), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Geometry2D) GetClosestPointsBetweenSegments(p1 Vector2, q1 Vector2, p2 Vector2, q2 Vector2) PackedVector2Array {
	cargs := []gdc.ConstTypePtr{p1.AsCTypePtr(), q1.AsCTypePtr(), p2.AsCTypePtr(), q2.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector2Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry2D.fnGetClosestPointsBetweenSegments), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Geometry2D) GetClosestPointToSegment(point Vector2, s1 Vector2, s2 Vector2) Vector2 {
	cargs := []gdc.ConstTypePtr{point.AsCTypePtr(), s1.AsCTypePtr(), s2.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry2D.fnGetClosestPointToSegment), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Geometry2D) GetClosestPointToSegmentUncapped(point Vector2, s1 Vector2, s2 Vector2) Vector2 {
	cargs := []gdc.ConstTypePtr{point.AsCTypePtr(), s1.AsCTypePtr(), s2.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry2D.fnGetClosestPointToSegmentUncapped), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Geometry2D) PointIsInsideTriangle(point Vector2, a Vector2, b Vector2, c Vector2) bool {
	cargs := []gdc.ConstTypePtr{point.AsCTypePtr(), a.AsCTypePtr(), b.AsCTypePtr(), c.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry2D.fnPointIsInsideTriangle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Geometry2D) IsPolygonClockwise(polygon PackedVector2Array) bool {
	cargs := []gdc.ConstTypePtr{polygon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry2D.fnIsPolygonClockwise), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Geometry2D) IsPointInPolygon(point Vector2, polygon PackedVector2Array) bool {
	cargs := []gdc.ConstTypePtr{point.AsCTypePtr(), polygon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry2D.fnIsPointInPolygon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Geometry2D) TriangulatePolygon(polygon PackedVector2Array) PackedInt32Array {
	cargs := []gdc.ConstTypePtr{polygon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry2D.fnTriangulatePolygon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Geometry2D) TriangulateDelaunay(points PackedVector2Array) PackedInt32Array {
	cargs := []gdc.ConstTypePtr{points.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry2D.fnTriangulateDelaunay), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Geometry2D) ConvexHull(points PackedVector2Array) PackedVector2Array {
	cargs := []gdc.ConstTypePtr{points.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector2Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry2D.fnConvexHull), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Geometry2D) DecomposePolygonInConvex(polygon PackedVector2Array) []PackedVector2Array {
	cargs := []gdc.ConstTypePtr{polygon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry2D.fnDecomposePolygonInConvex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[PackedVector2Array](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Geometry2D) MergePolygons(polygon_a PackedVector2Array, polygon_b PackedVector2Array) []PackedVector2Array {
	cargs := []gdc.ConstTypePtr{polygon_a.AsCTypePtr(), polygon_b.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry2D.fnMergePolygons), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[PackedVector2Array](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Geometry2D) ClipPolygons(polygon_a PackedVector2Array, polygon_b PackedVector2Array) []PackedVector2Array {
	cargs := []gdc.ConstTypePtr{polygon_a.AsCTypePtr(), polygon_b.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry2D.fnClipPolygons), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[PackedVector2Array](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Geometry2D) IntersectPolygons(polygon_a PackedVector2Array, polygon_b PackedVector2Array) []PackedVector2Array {
	cargs := []gdc.ConstTypePtr{polygon_a.AsCTypePtr(), polygon_b.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry2D.fnIntersectPolygons), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[PackedVector2Array](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Geometry2D) ExcludePolygons(polygon_a PackedVector2Array, polygon_b PackedVector2Array) []PackedVector2Array {
	cargs := []gdc.ConstTypePtr{polygon_a.AsCTypePtr(), polygon_b.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry2D.fnExcludePolygons), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[PackedVector2Array](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Geometry2D) ClipPolylineWithPolygon(polyline PackedVector2Array, polygon PackedVector2Array) []PackedVector2Array {
	cargs := []gdc.ConstTypePtr{polyline.AsCTypePtr(), polygon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry2D.fnClipPolylineWithPolygon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[PackedVector2Array](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Geometry2D) IntersectPolylineWithPolygon(polyline PackedVector2Array, polygon PackedVector2Array) []PackedVector2Array {
	cargs := []gdc.ConstTypePtr{polyline.AsCTypePtr(), polygon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry2D.fnIntersectPolylineWithPolygon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[PackedVector2Array](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Geometry2D) OffsetPolygon(polygon PackedVector2Array, delta float64, join_type Geometry2DPolyJoinType) []PackedVector2Array {
	cargs := []gdc.ConstTypePtr{polygon.AsCTypePtr(), gdc.ConstTypePtr(&delta), gdc.ConstTypePtr(&join_type)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()
	pinner.Pin(&delta)
	pinner.Pin(&join_type)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry2D.fnOffsetPolygon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[PackedVector2Array](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Geometry2D) OffsetPolyline(polyline PackedVector2Array, delta float64, join_type Geometry2DPolyJoinType, end_type Geometry2DPolyEndType) []PackedVector2Array {
	cargs := []gdc.ConstTypePtr{polyline.AsCTypePtr(), gdc.ConstTypePtr(&delta), gdc.ConstTypePtr(&join_type), gdc.ConstTypePtr(&end_type)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()
	pinner.Pin(&delta)
	pinner.Pin(&join_type)
	pinner.Pin(&end_type)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry2D.fnOffsetPolyline), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[PackedVector2Array](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Geometry2D) MakeAtlas(sizes PackedVector2Array) Dictionary {
	cargs := []gdc.ConstTypePtr{sizes.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometry2D.fnMakeAtlas), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
