// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TileMap struct {
  obj gdc.ObjectPtr
}

func (me *TileMap) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TileMap) BaseClass() string {
  return "TileMap"
}

type TileMapVisibilityMode int
const (
  TileMapVisibilityModeVisibilityModeDefault TileMapVisibilityMode = 0
  TileMapVisibilityModeVisibilityModeForceHide TileMapVisibilityMode = 2
  TileMapVisibilityModeVisibilityModeForceShow TileMapVisibilityMode = 1
)

func  (me *TileMap) XUseTileDataRuntimeUpdate(layer int, coords Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) XTileDataRuntimeUpdate(layer int, coords Vector2i, tile_data TileData, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) SetTileset(tileset TileSet, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) GetTileset() { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) SetQuadrantSize(size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) GetQuadrantSize() { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) GetLayersCount() { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) AddLayer(to_position int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) MoveLayer(layer int, to_position int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) RemoveLayer(layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) SetLayerName(layer int, name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) GetLayerName(layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) SetLayerEnabled(layer int, enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) IsLayerEnabled(layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) SetLayerModulate(layer int, modulate Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) GetLayerModulate(layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) SetLayerYSortEnabled(layer int, y_sort_enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) IsLayerYSortEnabled(layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) SetLayerYSortOrigin(layer int, y_sort_origin int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) GetLayerYSortOrigin(layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) SetLayerZIndex(layer int, z_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) GetLayerZIndex(layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) SetCollisionAnimatable(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) IsCollisionAnimatable() { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) SetCollisionVisibilityMode(collision_visibility_mode TileMapVisibilityMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) GetCollisionVisibilityMode() { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) SetNavigationVisibilityMode(navigation_visibility_mode TileMapVisibilityMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) GetNavigationVisibilityMode() { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) SetNavigationMap(layer int, map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) GetNavigationMap(layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) SetCell(layer int, coords Vector2i, source_id int, atlas_coords Vector2i, alternative_tile int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) EraseCell(layer int, coords Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) GetCellSourceId(layer int, coords Vector2i, use_proxies bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) GetCellAtlasCoords(layer int, coords Vector2i, use_proxies bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) GetCellAlternativeTile(layer int, coords Vector2i, use_proxies bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) GetCellTileData(layer int, coords Vector2i, use_proxies bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) GetCoordsForBodyRid(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) GetLayerForBodyRid(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) GetPattern(layer int, coords_array Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) MapPattern(position_in_tilemap Vector2i, coords_in_pattern Vector2i, pattern TileMapPattern, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) SetPattern(layer int, position Vector2i, pattern TileMapPattern, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) SetCellsTerrainConnect(layer int, cells Vector2i, terrain_set int, terrain int, ignore_empty_terrains bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) SetCellsTerrainPath(layer int, path Vector2i, terrain_set int, terrain int, ignore_empty_terrains bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) FixInvalidTiles() { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) ClearLayer(layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) Clear() { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) ForceUpdate(layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) GetSurroundingCells(coords Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) GetUsedCells(layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) GetUsedCellsById(layer int, source_id int, atlas_coords Vector2i, alternative_tile int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) GetUsedRect() { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) MapToLocal(map_position Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) LocalToMap(local_position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *TileMap) GetNeighborCell(coords Vector2i, neighbor TileSetCellNeighbor, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
