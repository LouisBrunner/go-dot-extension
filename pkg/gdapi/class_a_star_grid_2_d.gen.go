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

type ptrsForAStarGrid2DList struct {
	fnXEstimateCost               gdc.MethodBindPtr
	fnXComputeCost                gdc.MethodBindPtr
	fnSetRegion                   gdc.MethodBindPtr
	fnGetRegion                   gdc.MethodBindPtr
	fnSetSize                     gdc.MethodBindPtr
	fnGetSize                     gdc.MethodBindPtr
	fnSetOffset                   gdc.MethodBindPtr
	fnGetOffset                   gdc.MethodBindPtr
	fnSetCellSize                 gdc.MethodBindPtr
	fnGetCellSize                 gdc.MethodBindPtr
	fnIsInBounds                  gdc.MethodBindPtr
	fnIsInBoundsv                 gdc.MethodBindPtr
	fnIsDirty                     gdc.MethodBindPtr
	fnUpdate                      gdc.MethodBindPtr
	fnSetJumpingEnabled           gdc.MethodBindPtr
	fnIsJumpingEnabled            gdc.MethodBindPtr
	fnSetDiagonalMode             gdc.MethodBindPtr
	fnGetDiagonalMode             gdc.MethodBindPtr
	fnSetDefaultComputeHeuristic  gdc.MethodBindPtr
	fnGetDefaultComputeHeuristic  gdc.MethodBindPtr
	fnSetDefaultEstimateHeuristic gdc.MethodBindPtr
	fnGetDefaultEstimateHeuristic gdc.MethodBindPtr
	fnSetPointSolid               gdc.MethodBindPtr
	fnIsPointSolid                gdc.MethodBindPtr
	fnSetPointWeightScale         gdc.MethodBindPtr
	fnGetPointWeightScale         gdc.MethodBindPtr
	fnFillSolidRegion             gdc.MethodBindPtr
	fnFillWeightScaleRegion       gdc.MethodBindPtr
	fnClear                       gdc.MethodBindPtr
	fnGetPointPosition            gdc.MethodBindPtr
	fnGetPointPath                gdc.MethodBindPtr
	fnGetIdPath                   gdc.MethodBindPtr
}

var ptrsForAStarGrid2D ptrsForAStarGrid2DList

func initAStarGrid2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AStarGrid2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_region")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnSetRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1763793166))
	}
	{
		methodName := StringNameFromStr("get_region")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnGetRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 410525958))
	}
	{
		methodName := StringNameFromStr("set_size")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
	}
	{
		methodName := StringNameFromStr("get_size")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
	}
	{
		methodName := StringNameFromStr("set_offset")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnSetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_offset")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnGetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_cell_size")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnSetCellSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_cell_size")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnGetCellSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("is_in_bounds")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnIsInBounds = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2522259332))
	}
	{
		methodName := StringNameFromStr("is_in_boundsv")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnIsInBoundsv = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3900751641))
	}
	{
		methodName := StringNameFromStr("is_dirty")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnIsDirty = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("update")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_jumping_enabled")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnSetJumpingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_jumping_enabled")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnIsJumpingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_diagonal_mode")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnSetDiagonalMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1017829798))
	}
	{
		methodName := StringNameFromStr("get_diagonal_mode")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnGetDiagonalMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3129282674))
	}
	{
		methodName := StringNameFromStr("set_default_compute_heuristic")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnSetDefaultComputeHeuristic = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1044375519))
	}
	{
		methodName := StringNameFromStr("get_default_compute_heuristic")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnGetDefaultComputeHeuristic = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2074731422))
	}
	{
		methodName := StringNameFromStr("set_default_estimate_heuristic")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnSetDefaultEstimateHeuristic = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1044375519))
	}
	{
		methodName := StringNameFromStr("get_default_estimate_heuristic")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnGetDefaultEstimateHeuristic = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2074731422))
	}
	{
		methodName := StringNameFromStr("set_point_solid")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnSetPointSolid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1765703753))
	}
	{
		methodName := StringNameFromStr("is_point_solid")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnIsPointSolid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3900751641))
	}
	{
		methodName := StringNameFromStr("set_point_weight_scale")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnSetPointWeightScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2262553149))
	}
	{
		methodName := StringNameFromStr("get_point_weight_scale")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnGetPointWeightScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 719993801))
	}
	{
		methodName := StringNameFromStr("fill_solid_region")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnFillSolidRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2261970063))
	}
	{
		methodName := StringNameFromStr("fill_weight_scale_region")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnFillWeightScaleRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2793244083))
	}
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_point_position")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnGetPointPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 108438297))
	}
	{
		methodName := StringNameFromStr("get_point_path")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnGetPointPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 690373547))
	}
	{
		methodName := StringNameFromStr("get_id_path")
		defer methodName.Destroy()
		ptrsForAStarGrid2D.fnGetIdPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1989391000))
	}

}

type AStarGrid2D struct {
	RefCounted
}

func (me *AStarGrid2D) BaseClass() string {
	return "AStarGrid2D"
}

func NewAStarGrid2D() *AStarGrid2D {
	str := StringNameFromStr("AStarGrid2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AStarGrid2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type AStarGrid2DHeuristic int

const (
	AStarGrid2DHeuristicHeuristicEuclidean AStarGrid2DHeuristic = 0
	AStarGrid2DHeuristicHeuristicManhattan AStarGrid2DHeuristic = 1
	AStarGrid2DHeuristicHeuristicOctile    AStarGrid2DHeuristic = 2
	AStarGrid2DHeuristicHeuristicChebyshev AStarGrid2DHeuristic = 3
	AStarGrid2DHeuristicHeuristicMax       AStarGrid2DHeuristic = 4
)

type AStarGrid2DDiagonalMode int

const (
	AStarGrid2DDiagonalModeDiagonalModeAlways             AStarGrid2DDiagonalMode = 0
	AStarGrid2DDiagonalModeDiagonalModeNever              AStarGrid2DDiagonalMode = 1
	AStarGrid2DDiagonalModeDiagonalModeAtLeastOneWalkable AStarGrid2DDiagonalMode = 2
	AStarGrid2DDiagonalModeDiagonalModeOnlyIfNoObstacles  AStarGrid2DDiagonalMode = 3
	AStarGrid2DDiagonalModeDiagonalModeMax                AStarGrid2DDiagonalMode = 4
)

func (me *AStarGrid2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AStarGrid2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AStarGrid2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AStarGrid2D) SetRegion(region Rect2i) {
	cargs := []gdc.ConstTypePtr{region.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnSetRegion), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AStarGrid2D) GetRegion() Rect2i {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnGetRegion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AStarGrid2D) SetSize(size Vector2i) {
	cargs := []gdc.ConstTypePtr{size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AStarGrid2D) GetSize() Vector2i {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AStarGrid2D) SetOffset(offset Vector2) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnSetOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AStarGrid2D) GetOffset() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnGetOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AStarGrid2D) SetCellSize(cell_size Vector2) {
	cargs := []gdc.ConstTypePtr{cell_size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnSetCellSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AStarGrid2D) GetCellSize() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnGetCellSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AStarGrid2D) IsInBounds(x int64, y int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&x)
	pinner.Pin(&y)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnIsInBounds), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AStarGrid2D) IsInBoundsv(id Vector2i) bool {
	cargs := []gdc.ConstTypePtr{id.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnIsInBoundsv), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AStarGrid2D) IsDirty() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnIsDirty), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AStarGrid2D) Update() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnUpdate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AStarGrid2D) SetJumpingEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnSetJumpingEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AStarGrid2D) IsJumpingEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnIsJumpingEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AStarGrid2D) SetDiagonalMode(mode AStarGrid2DDiagonalMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnSetDiagonalMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AStarGrid2D) GetDiagonalMode() AStarGrid2DDiagonalMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AStarGrid2DDiagonalMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnGetDiagonalMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *AStarGrid2D) SetDefaultComputeHeuristic(heuristic AStarGrid2DHeuristic) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&heuristic)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnSetDefaultComputeHeuristic), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AStarGrid2D) GetDefaultComputeHeuristic() AStarGrid2DHeuristic {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AStarGrid2DHeuristic

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnGetDefaultComputeHeuristic), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *AStarGrid2D) SetDefaultEstimateHeuristic(heuristic AStarGrid2DHeuristic) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&heuristic)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnSetDefaultEstimateHeuristic), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AStarGrid2D) GetDefaultEstimateHeuristic() AStarGrid2DHeuristic {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AStarGrid2DHeuristic

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnGetDefaultEstimateHeuristic), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *AStarGrid2D) SetPointSolid(id Vector2i, solid bool) {
	cargs := []gdc.ConstTypePtr{id.AsCTypePtr(), gdc.ConstTypePtr(&solid)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnSetPointSolid), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AStarGrid2D) IsPointSolid(id Vector2i) bool {
	cargs := []gdc.ConstTypePtr{id.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnIsPointSolid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AStarGrid2D) SetPointWeightScale(id Vector2i, weight_scale float64) {
	cargs := []gdc.ConstTypePtr{id.AsCTypePtr(), gdc.ConstTypePtr(&weight_scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnSetPointWeightScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AStarGrid2D) GetPointWeightScale(id Vector2i) float64 {
	cargs := []gdc.ConstTypePtr{id.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnGetPointWeightScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AStarGrid2D) FillSolidRegion(region Rect2i, solid bool) {
	cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&solid)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnFillSolidRegion), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AStarGrid2D) FillWeightScaleRegion(region Rect2i, weight_scale float64) {
	cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&weight_scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnFillWeightScaleRegion), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AStarGrid2D) Clear() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AStarGrid2D) GetPointPosition(id Vector2i) Vector2 {
	cargs := []gdc.ConstTypePtr{id.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnGetPointPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AStarGrid2D) GetPointPath(from_id Vector2i, to_id Vector2i) PackedVector2Array {
	cargs := []gdc.ConstTypePtr{from_id.AsCTypePtr(), to_id.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector2Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnGetPointPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AStarGrid2D) GetIdPath(from_id Vector2i, to_id Vector2i) []Vector2i {
	cargs := []gdc.ConstTypePtr{from_id.AsCTypePtr(), to_id.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStarGrid2D.fnGetIdPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Vector2i](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
