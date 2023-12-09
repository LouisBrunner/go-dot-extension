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



// Enums

type TileMapVisibilityMode int
const (
  TileMapVisibilityModeVisibilityModeDefault TileMapVisibilityMode = 0
  TileMapVisibilityModeVisibilityModeForceHide TileMapVisibilityMode = 2
  TileMapVisibilityModeVisibilityModeForceShow TileMapVisibilityMode = 1
)

func (me *TileMap) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TileMap) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TileMap) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *TileMap) XUseTileDataRuntimeUpdate(layer int, coords Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileMap) XTileDataRuntimeUpdate(layer int, coords Vector2i, tile_data TileData, )  {
  panic("TODO: implement")
}

func  (me *TileMap) SetTileset(tileset TileSet, )  {
  panic("TODO: implement")
}

func  (me *TileMap) GetTileset()  {
  panic("TODO: implement")
}

func  (me *TileMap) SetQuadrantSize(size int, )  {
  panic("TODO: implement")
}

func  (me *TileMap) GetQuadrantSize()  {
  panic("TODO: implement")
}

func  (me *TileMap) GetLayersCount()  {
  panic("TODO: implement")
}

func  (me *TileMap) AddLayer(to_position int, )  {
  panic("TODO: implement")
}

func  (me *TileMap) MoveLayer(layer int, to_position int, )  {
  panic("TODO: implement")
}

func  (me *TileMap) RemoveLayer(layer int, )  {
  panic("TODO: implement")
}

func  (me *TileMap) SetLayerName(layer int, name String, )  {
  panic("TODO: implement")
}

func  (me *TileMap) GetLayerName(layer int, )  {
  panic("TODO: implement")
}

func  (me *TileMap) SetLayerEnabled(layer int, enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TileMap) IsLayerEnabled(layer int, )  {
  panic("TODO: implement")
}

func  (me *TileMap) SetLayerModulate(layer int, modulate Color, )  {
  panic("TODO: implement")
}

func  (me *TileMap) GetLayerModulate(layer int, )  {
  panic("TODO: implement")
}

func  (me *TileMap) SetLayerYSortEnabled(layer int, y_sort_enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TileMap) IsLayerYSortEnabled(layer int, )  {
  panic("TODO: implement")
}

func  (me *TileMap) SetLayerYSortOrigin(layer int, y_sort_origin int, )  {
  panic("TODO: implement")
}

func  (me *TileMap) GetLayerYSortOrigin(layer int, )  {
  panic("TODO: implement")
}

func  (me *TileMap) SetLayerZIndex(layer int, z_index int, )  {
  panic("TODO: implement")
}

func  (me *TileMap) GetLayerZIndex(layer int, )  {
  panic("TODO: implement")
}

func  (me *TileMap) SetCollisionAnimatable(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TileMap) IsCollisionAnimatable()  {
  panic("TODO: implement")
}

func  (me *TileMap) SetCollisionVisibilityMode(collision_visibility_mode TileMapVisibilityMode, )  {
  panic("TODO: implement")
}

func  (me *TileMap) GetCollisionVisibilityMode()  {
  panic("TODO: implement")
}

func  (me *TileMap) SetNavigationVisibilityMode(navigation_visibility_mode TileMapVisibilityMode, )  {
  panic("TODO: implement")
}

func  (me *TileMap) GetNavigationVisibilityMode()  {
  panic("TODO: implement")
}

func  (me *TileMap) SetNavigationMap(layer int, map_ RID, )  {
  panic("TODO: implement")
}

func  (me *TileMap) GetNavigationMap(layer int, )  {
  panic("TODO: implement")
}

func  (me *TileMap) SetCell(layer int, coords Vector2i, source_id int, atlas_coords Vector2i, alternative_tile int, )  {
  panic("TODO: implement")
}

func  (me *TileMap) EraseCell(layer int, coords Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileMap) GetCellSourceId(layer int, coords Vector2i, use_proxies bool, )  {
  panic("TODO: implement")
}

func  (me *TileMap) GetCellAtlasCoords(layer int, coords Vector2i, use_proxies bool, )  {
  panic("TODO: implement")
}

func  (me *TileMap) GetCellAlternativeTile(layer int, coords Vector2i, use_proxies bool, )  {
  panic("TODO: implement")
}

func  (me *TileMap) GetCellTileData(layer int, coords Vector2i, use_proxies bool, )  {
  panic("TODO: implement")
}

func  (me *TileMap) GetCoordsForBodyRid(body RID, )  {
  panic("TODO: implement")
}

func  (me *TileMap) GetLayerForBodyRid(body RID, )  {
  panic("TODO: implement")
}

func  (me *TileMap) GetPattern(layer int, coords_array Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileMap) MapPattern(position_in_tilemap Vector2i, coords_in_pattern Vector2i, pattern TileMapPattern, )  {
  panic("TODO: implement")
}

func  (me *TileMap) SetPattern(layer int, position Vector2i, pattern TileMapPattern, )  {
  panic("TODO: implement")
}

func  (me *TileMap) SetCellsTerrainConnect(layer int, cells Vector2i, terrain_set int, terrain int, ignore_empty_terrains bool, )  {
  panic("TODO: implement")
}

func  (me *TileMap) SetCellsTerrainPath(layer int, path Vector2i, terrain_set int, terrain int, ignore_empty_terrains bool, )  {
  panic("TODO: implement")
}

func  (me *TileMap) FixInvalidTiles()  {
  panic("TODO: implement")
}

func  (me *TileMap) ClearLayer(layer int, )  {
  panic("TODO: implement")
}

func  (me *TileMap) Clear()  {
  panic("TODO: implement")
}

func  (me *TileMap) ForceUpdate(layer int, )  {
  panic("TODO: implement")
}

func  (me *TileMap) GetSurroundingCells(coords Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileMap) GetUsedCells(layer int, )  {
  panic("TODO: implement")
}

func  (me *TileMap) GetUsedCellsById(layer int, source_id int, atlas_coords Vector2i, alternative_tile int, )  {
  panic("TODO: implement")
}

func  (me *TileMap) GetUsedRect()  {
  panic("TODO: implement")
}

func  (me *TileMap) MapToLocal(map_position Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileMap) LocalToMap(local_position Vector2, )  {
  panic("TODO: implement")
}

func  (me *TileMap) GetNeighborCell(coords Vector2i, neighbor TileSetCellNeighbor, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
