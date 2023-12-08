// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AStarGrid2D struct {
  obj gdc.ObjectPtr
}

func (me *AStarGrid2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AStarGrid2D) BaseClass() string {
  return "AStarGrid2D"
}

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

func  (me *AStarGrid2D) XEstimateCost(from_id Vector2i, to_id Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) XComputeCost(from_id Vector2i, to_id Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) SetRegion(region Rect2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) GetRegion() { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) SetSize(size Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) GetSize() { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) SetOffset(offset Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) GetOffset() { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) SetCellSize(cell_size Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) GetCellSize() { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) IsInBounds(x int, y int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) IsInBoundsv(id Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) IsDirty() { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) Update() { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) SetJumpingEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) IsJumpingEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) SetDiagonalMode(mode AStarGrid2DDiagonalMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) GetDiagonalMode() { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) SetDefaultComputeHeuristic(heuristic AStarGrid2DHeuristic, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) GetDefaultComputeHeuristic() { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) SetDefaultEstimateHeuristic(heuristic AStarGrid2DHeuristic, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) GetDefaultEstimateHeuristic() { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) SetPointSolid(id Vector2i, solid bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) IsPointSolid(id Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) SetPointWeightScale(id Vector2i, weight_scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) GetPointWeightScale(id Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) Clear() { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) GetPointPosition(id Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) GetPointPath(from_id Vector2i, to_id Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStarGrid2D) GetIdPath(from_id Vector2i, to_id Vector2i, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
