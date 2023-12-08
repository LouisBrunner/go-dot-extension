// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TileSet struct {
  obj gdc.ObjectPtr
}

func (me *TileSet) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TileSet) BaseClass() string {
  return "TileSet"
}

type TileSetTileShape int
const (
  TileSetTileShapeSquare TileSetTileShape = 0
  TileSetTileShapeIsometric TileSetTileShape = 1
  TileSetTileShapeHalfOffsetSquare TileSetTileShape = 2
  TileSetTileShapeHexagon TileSetTileShape = 3
)

type TileSetTileLayout int
const (
  TileSetTileLayoutStacked TileSetTileLayout = 0
  TileSetTileLayoutStackedOffset TileSetTileLayout = 1
  TileSetTileLayoutStairsRight TileSetTileLayout = 2
  TileSetTileLayoutStairsDown TileSetTileLayout = 3
  TileSetTileLayoutDiamondRight TileSetTileLayout = 4
  TileSetTileLayoutDiamondDown TileSetTileLayout = 5
)

type TileSetTileOffsetAxis int
const (
  TileSetTileOffsetAxisHorizontal TileSetTileOffsetAxis = 0
  TileSetTileOffsetAxisVertical TileSetTileOffsetAxis = 1
)

type TileSetCellNeighbor int
const (
  TileSetCellNeighborRightSide TileSetCellNeighbor = 0
  TileSetCellNeighborRightCorner TileSetCellNeighbor = 1
  TileSetCellNeighborBottomRightSide TileSetCellNeighbor = 2
  TileSetCellNeighborBottomRightCorner TileSetCellNeighbor = 3
  TileSetCellNeighborBottomSide TileSetCellNeighbor = 4
  TileSetCellNeighborBottomCorner TileSetCellNeighbor = 5
  TileSetCellNeighborBottomLeftSide TileSetCellNeighbor = 6
  TileSetCellNeighborBottomLeftCorner TileSetCellNeighbor = 7
  TileSetCellNeighborLeftSide TileSetCellNeighbor = 8
  TileSetCellNeighborLeftCorner TileSetCellNeighbor = 9
  TileSetCellNeighborTopLeftSide TileSetCellNeighbor = 10
  TileSetCellNeighborTopLeftCorner TileSetCellNeighbor = 11
  TileSetCellNeighborTopSide TileSetCellNeighbor = 12
  TileSetCellNeighborTopCorner TileSetCellNeighbor = 13
  TileSetCellNeighborTopRightSide TileSetCellNeighbor = 14
  TileSetCellNeighborTopRightCorner TileSetCellNeighbor = 15
)

type TileSetTerrainMode int
const (
  TileSetTerrainModeMatchCornersAndSides TileSetTerrainMode = 0
  TileSetTerrainModeMatchCorners TileSetTerrainMode = 1
  TileSetTerrainModeMatchSides TileSetTerrainMode = 2
)
