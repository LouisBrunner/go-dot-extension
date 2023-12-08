// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  TileSetTileShapeTileShapeSquare TileSetTileShape = 0
  TileSetTileShapeTileShapeIsometric TileSetTileShape = 1
  TileSetTileShapeTileShapeHalfOffsetSquare TileSetTileShape = 2
  TileSetTileShapeTileShapeHexagon TileSetTileShape = 3
)

type TileSetTileLayout int
const (
  TileSetTileLayoutTileLayoutStacked TileSetTileLayout = 0
  TileSetTileLayoutTileLayoutStackedOffset TileSetTileLayout = 1
  TileSetTileLayoutTileLayoutStairsRight TileSetTileLayout = 2
  TileSetTileLayoutTileLayoutStairsDown TileSetTileLayout = 3
  TileSetTileLayoutTileLayoutDiamondRight TileSetTileLayout = 4
  TileSetTileLayoutTileLayoutDiamondDown TileSetTileLayout = 5
)

type TileSetTileOffsetAxis int
const (
  TileSetTileOffsetAxisTileOffsetAxisHorizontal TileSetTileOffsetAxis = 0
  TileSetTileOffsetAxisTileOffsetAxisVertical TileSetTileOffsetAxis = 1
)

type TileSetCellNeighbor int
const (
  TileSetCellNeighborCellNeighborRightSide TileSetCellNeighbor = 0
  TileSetCellNeighborCellNeighborRightCorner TileSetCellNeighbor = 1
  TileSetCellNeighborCellNeighborBottomRightSide TileSetCellNeighbor = 2
  TileSetCellNeighborCellNeighborBottomRightCorner TileSetCellNeighbor = 3
  TileSetCellNeighborCellNeighborBottomSide TileSetCellNeighbor = 4
  TileSetCellNeighborCellNeighborBottomCorner TileSetCellNeighbor = 5
  TileSetCellNeighborCellNeighborBottomLeftSide TileSetCellNeighbor = 6
  TileSetCellNeighborCellNeighborBottomLeftCorner TileSetCellNeighbor = 7
  TileSetCellNeighborCellNeighborLeftSide TileSetCellNeighbor = 8
  TileSetCellNeighborCellNeighborLeftCorner TileSetCellNeighbor = 9
  TileSetCellNeighborCellNeighborTopLeftSide TileSetCellNeighbor = 10
  TileSetCellNeighborCellNeighborTopLeftCorner TileSetCellNeighbor = 11
  TileSetCellNeighborCellNeighborTopSide TileSetCellNeighbor = 12
  TileSetCellNeighborCellNeighborTopCorner TileSetCellNeighbor = 13
  TileSetCellNeighborCellNeighborTopRightSide TileSetCellNeighbor = 14
  TileSetCellNeighborCellNeighborTopRightCorner TileSetCellNeighbor = 15
)

type TileSetTerrainMode int
const (
  TileSetTerrainModeTerrainModeMatchCornersAndSides TileSetTerrainMode = 0
  TileSetTerrainModeTerrainModeMatchCorners TileSetTerrainMode = 1
  TileSetTerrainModeTerrainModeMatchSides TileSetTerrainMode = 2
)

func  (me *TileSet) GetNextSourceId() { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) AddSource(source TileSetSource, atlas_source_id_override int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) RemoveSource(source_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) SetSourceId(source_id int, new_source_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetSourceCount() { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetSourceId(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) HasSource(source_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetSource(source_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) SetTileShape(shape TileSetTileShape, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetTileShape() { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) SetTileLayout(layout TileSetTileLayout, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetTileLayout() { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) SetTileOffsetAxis(alignment TileSetTileOffsetAxis, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetTileOffsetAxis() { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) SetTileSize(size Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetTileSize() { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) SetUvClipping(uv_clipping bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) IsUvClipping() { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetOcclusionLayersCount() { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) AddOcclusionLayer(to_position int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) MoveOcclusionLayer(layer_index int, to_position int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) RemoveOcclusionLayer(layer_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) SetOcclusionLayerLightMask(layer_index int, light_mask int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetOcclusionLayerLightMask(layer_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) SetOcclusionLayerSdfCollision(layer_index int, sdf_collision bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetOcclusionLayerSdfCollision(layer_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetPhysicsLayersCount() { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) AddPhysicsLayer(to_position int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) MovePhysicsLayer(layer_index int, to_position int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) RemovePhysicsLayer(layer_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) SetPhysicsLayerCollisionLayer(layer_index int, layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetPhysicsLayerCollisionLayer(layer_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) SetPhysicsLayerCollisionMask(layer_index int, mask int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetPhysicsLayerCollisionMask(layer_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) SetPhysicsLayerPhysicsMaterial(layer_index int, physics_material PhysicsMaterial, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetPhysicsLayerPhysicsMaterial(layer_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetTerrainSetsCount() { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) AddTerrainSet(to_position int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) MoveTerrainSet(terrain_set int, to_position int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) RemoveTerrainSet(terrain_set int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) SetTerrainSetMode(terrain_set int, mode TileSetTerrainMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetTerrainSetMode(terrain_set int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetTerrainsCount(terrain_set int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) AddTerrain(terrain_set int, to_position int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) MoveTerrain(terrain_set int, terrain_index int, to_position int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) RemoveTerrain(terrain_set int, terrain_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) SetTerrainName(terrain_set int, terrain_index int, name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetTerrainName(terrain_set int, terrain_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) SetTerrainColor(terrain_set int, terrain_index int, color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetTerrainColor(terrain_set int, terrain_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetNavigationLayersCount() { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) AddNavigationLayer(to_position int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) MoveNavigationLayer(layer_index int, to_position int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) RemoveNavigationLayer(layer_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) SetNavigationLayerLayers(layer_index int, layers int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetNavigationLayerLayers(layer_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) SetNavigationLayerLayerValue(layer_index int, layer_number int, value bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetNavigationLayerLayerValue(layer_index int, layer_number int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetCustomDataLayersCount() { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) AddCustomDataLayer(to_position int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) MoveCustomDataLayer(layer_index int, to_position int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) RemoveCustomDataLayer(layer_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetCustomDataLayerByName(layer_name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) SetCustomDataLayerName(layer_index int, layer_name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetCustomDataLayerName(layer_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) SetCustomDataLayerType(layer_index int, layer_type VariantType, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetCustomDataLayerType(layer_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) SetSourceLevelTileProxy(source_from int, source_to int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetSourceLevelTileProxy(source_from int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) HasSourceLevelTileProxy(source_from int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) RemoveSourceLevelTileProxy(source_from int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) SetCoordsLevelTileProxy(p_source_from int, coords_from Vector2i, source_to int, coords_to Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetCoordsLevelTileProxy(source_from int, coords_from Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) HasCoordsLevelTileProxy(source_from int, coords_from Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) RemoveCoordsLevelTileProxy(source_from int, coords_from Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) SetAlternativeLevelTileProxy(source_from int, coords_from Vector2i, alternative_from int, source_to int, coords_to Vector2i, alternative_to int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetAlternativeLevelTileProxy(source_from int, coords_from Vector2i, alternative_from int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) HasAlternativeLevelTileProxy(source_from int, coords_from Vector2i, alternative_from int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) RemoveAlternativeLevelTileProxy(source_from int, coords_from Vector2i, alternative_from int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) MapTileProxy(source_from int, coords_from Vector2i, alternative_from int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) CleanupInvalidTileProxies() { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) ClearTileProxies() { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) AddPattern(pattern TileMapPattern, index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetPattern(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) RemovePattern(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileSet) GetPatternsCount() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
