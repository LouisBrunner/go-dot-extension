// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AStarGrid2D struct {
  obj gdc.ObjectPtr
}

func (me *AStarGrid2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AStarGrid2D) BaseClass() string {
  return "AStarGrid2D"
}



// Enums

type AStarGrid2DHeuristic int
const (
  AStarGrid2DHeuristicHeuristicEuclidean AStarGrid2DHeuristic = 0
  AStarGrid2DHeuristicHeuristicManhattan AStarGrid2DHeuristic = 1
  AStarGrid2DHeuristicHeuristicOctile AStarGrid2DHeuristic = 2
  AStarGrid2DHeuristicHeuristicChebyshev AStarGrid2DHeuristic = 3
  AStarGrid2DHeuristicHeuristicMax AStarGrid2DHeuristic = 4
)

type AStarGrid2DDiagonalMode int
const (
  AStarGrid2DDiagonalModeDiagonalModeAlways AStarGrid2DDiagonalMode = 0
  AStarGrid2DDiagonalModeDiagonalModeNever AStarGrid2DDiagonalMode = 1
  AStarGrid2DDiagonalModeDiagonalModeAtLeastOneWalkable AStarGrid2DDiagonalMode = 2
  AStarGrid2DDiagonalModeDiagonalModeOnlyIfNoObstacles AStarGrid2DDiagonalMode = 3
  AStarGrid2DDiagonalModeDiagonalModeMax AStarGrid2DDiagonalMode = 4
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

func  (me *AStarGrid2D) SetRegion(region Rect2i, )  {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1763793166) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(region.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AStarGrid2D) GetRegion() Rect2i {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 410525958) // FIXME: should cache?
  var ret Rect2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStarGrid2D) SetSize(size Vector2i, )  {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AStarGrid2D) GetSize() Vector2i {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStarGrid2D) SetOffset(offset Vector2, )  {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AStarGrid2D) GetOffset() Vector2 {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStarGrid2D) SetCellSize(cell_size Vector2, )  {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cell_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(cell_size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AStarGrid2D) GetCellSize() Vector2 {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cell_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStarGrid2D) IsInBounds(x int, y int, ) bool {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_in_bounds")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2522259332) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStarGrid2D) IsInBoundsv(id Vector2i, ) bool {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_in_boundsv")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3900751641) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(id.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStarGrid2D) IsDirty() bool {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_dirty")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStarGrid2D) Update()  {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("update")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AStarGrid2D) SetJumpingEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_jumping_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AStarGrid2D) IsJumpingEnabled() bool {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_jumping_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStarGrid2D) SetDiagonalMode(mode AStarGrid2DDiagonalMode, )  {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_diagonal_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1017829798) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AStarGrid2D) GetDiagonalMode() AStarGrid2DDiagonalMode {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_diagonal_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3129282674) // FIXME: should cache?
  var ret AStarGrid2DDiagonalMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStarGrid2D) SetDefaultComputeHeuristic(heuristic AStarGrid2DHeuristic, )  {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_compute_heuristic")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1044375519) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&heuristic), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AStarGrid2D) GetDefaultComputeHeuristic() AStarGrid2DHeuristic {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_default_compute_heuristic")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2074731422) // FIXME: should cache?
  var ret AStarGrid2DHeuristic
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStarGrid2D) SetDefaultEstimateHeuristic(heuristic AStarGrid2DHeuristic, )  {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_estimate_heuristic")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1044375519) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&heuristic), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AStarGrid2D) GetDefaultEstimateHeuristic() AStarGrid2DHeuristic {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_default_estimate_heuristic")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2074731422) // FIXME: should cache?
  var ret AStarGrid2DHeuristic
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStarGrid2D) SetPointSolid(id Vector2i, solid bool, )  {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_point_solid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2825551965) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(id.AsCTypePtr()), gdc.ConstTypePtr(&solid), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AStarGrid2D) IsPointSolid(id Vector2i, ) bool {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_point_solid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3900751641) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(id.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStarGrid2D) SetPointWeightScale(id Vector2i, weight_scale float32, )  {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_point_weight_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2262553149) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(id.AsCTypePtr()), gdc.ConstTypePtr(&weight_scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AStarGrid2D) GetPointWeightScale(id Vector2i, ) float32 {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_weight_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 719993801) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(id.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStarGrid2D) Clear()  {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AStarGrid2D) GetPointPosition(id Vector2i, ) Vector2 {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 108438297) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(id.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStarGrid2D) GetPointPath(from_id Vector2i, to_id Vector2i, ) PackedVector2Array {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 690373547) // FIXME: should cache?
  var ret PackedVector2Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from_id.AsCTypePtr()), gdc.ConstTypePtr(to_id.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStarGrid2D) GetIdPath(from_id Vector2i, to_id Vector2i, ) Vector2i {
  classNameV := StringNameFromStr("AStarGrid2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_id_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1989391000) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from_id.AsCTypePtr()), gdc.ConstTypePtr(to_id.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
