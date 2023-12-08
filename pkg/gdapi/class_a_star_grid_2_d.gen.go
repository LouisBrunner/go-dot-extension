// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  AStarGrid2DHeuristicEuclidean AStarGrid2DHeuristic = 0
  AStarGrid2DHeuristicManhattan AStarGrid2DHeuristic = 1
  AStarGrid2DHeuristicOctile AStarGrid2DHeuristic = 2
  AStarGrid2DHeuristicChebyshev AStarGrid2DHeuristic = 3
  AStarGrid2DHeuristicMax AStarGrid2DHeuristic = 4
)

type AStarGrid2DDiagonalMode int
const (
  AStarGrid2DDiagonalModeAlways AStarGrid2DDiagonalMode = 0
  AStarGrid2DDiagonalModeNever AStarGrid2DDiagonalMode = 1
  AStarGrid2DDiagonalModeAtLeastOneWalkable AStarGrid2DDiagonalMode = 2
  AStarGrid2DDiagonalModeOnlyIfNoObstacles AStarGrid2DDiagonalMode = 3
  AStarGrid2DDiagonalModeMax AStarGrid2DDiagonalMode = 4
)
