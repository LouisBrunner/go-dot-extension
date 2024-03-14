// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TileMap struct {
  Node2D
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

func  (me *TileMap) SetNavigationMap(layer int, map_ RID, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_navigation_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4040184819) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(map_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) GetNavigationMap(layer int, ) RID {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 495598643) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) ForceUpdate(layer int, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("force_update")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1025054187) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) SetTileset(tileset TileSet, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tileset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 774531446) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(tileset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) GetTileset() TileSet {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tileset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2678226422) // FIXME: should cache?
  var ret TileSet
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) SetRenderingQuadrantSize(size int, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rendering_quadrant_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) GetRenderingQuadrantSize() int {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rendering_quadrant_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) GetLayersCount() int {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_layers_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) AddLayer(to_position int, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&to_position), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) MoveLayer(layer int, to_position int, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(&to_position), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) RemoveLayer(layer int, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) SetLayerName(layer int, name String, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_layer_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) GetLayerName(layer int, ) String {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_layer_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) SetLayerEnabled(layer int, enabled bool, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_layer_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) IsLayerEnabled(layer int, ) bool {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_layer_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) SetLayerModulate(layer int, modulate Color, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_layer_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2878471219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(modulate.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) GetLayerModulate(layer int, ) Color {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_layer_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3457211756) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) SetLayerYSortEnabled(layer int, y_sort_enabled bool, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_layer_y_sort_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(&y_sort_enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) IsLayerYSortEnabled(layer int, ) bool {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_layer_y_sort_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) SetLayerYSortOrigin(layer int, y_sort_origin int, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_layer_y_sort_origin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(&y_sort_origin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) GetLayerYSortOrigin(layer int, ) int {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_layer_y_sort_origin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) SetLayerZIndex(layer int, z_index int, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_layer_z_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(&z_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) GetLayerZIndex(layer int, ) int {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_layer_z_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) SetLayerNavigationEnabled(layer int, enabled bool, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_layer_navigation_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) IsLayerNavigationEnabled(layer int, ) bool {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_layer_navigation_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) SetLayerNavigationMap(layer int, map_ RID, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_layer_navigation_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4040184819) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(map_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) GetLayerNavigationMap(layer int, ) RID {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_layer_navigation_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 495598643) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) SetCollisionAnimatable(enabled bool, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_animatable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) IsCollisionAnimatable() bool {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_collision_animatable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) SetCollisionVisibilityMode(collision_visibility_mode TileMapVisibilityMode, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_visibility_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3193440636) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_visibility_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) GetCollisionVisibilityMode() TileMapVisibilityMode {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_visibility_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2026313073) // FIXME: should cache?
  var ret TileMapVisibilityMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) SetNavigationVisibilityMode(navigation_visibility_mode TileMapVisibilityMode, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_navigation_visibility_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3193440636) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&navigation_visibility_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) GetNavigationVisibilityMode() TileMapVisibilityMode {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_visibility_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2026313073) // FIXME: should cache?
  var ret TileMapVisibilityMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) SetCell(layer int, coords Vector2i, source_id int, atlas_coords Vector2i, alternative_tile int, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cell")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 966713560) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(coords.AsCTypePtr()), gdc.ConstTypePtr(&source_id), gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), gdc.ConstTypePtr(&alternative_tile), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) EraseCell(layer int, coords Vector2i, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("erase_cell")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2311374912) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(coords.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) GetCellSourceId(layer int, coords Vector2i, use_proxies bool, ) int {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cell_source_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 551761942) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(coords.AsCTypePtr()), gdc.ConstTypePtr(&use_proxies), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) GetCellAtlasCoords(layer int, coords Vector2i, use_proxies bool, ) Vector2i {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cell_atlas_coords")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1869815066) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(coords.AsCTypePtr()), gdc.ConstTypePtr(&use_proxies), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) GetCellAlternativeTile(layer int, coords Vector2i, use_proxies bool, ) int {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cell_alternative_tile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 551761942) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(coords.AsCTypePtr()), gdc.ConstTypePtr(&use_proxies), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) GetCellTileData(layer int, coords Vector2i, use_proxies bool, ) TileData {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cell_tile_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2849631287) // FIXME: should cache?
  var ret TileData
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(coords.AsCTypePtr()), gdc.ConstTypePtr(&use_proxies), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) GetCoordsForBodyRid(body RID, ) Vector2i {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_coords_for_body_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 291584212) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) GetLayerForBodyRid(body RID, ) int {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_layer_for_body_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3917799429) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) GetPattern(layer int, coords_array Vector2i, ) TileMapPattern {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pattern")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2833570986) // FIXME: should cache?
  var ret TileMapPattern
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(coords_array.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) MapPattern(position_in_tilemap Vector2i, coords_in_pattern Vector2i, pattern TileMapPattern, ) Vector2i {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_pattern")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1864516957) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position_in_tilemap.AsCTypePtr()), gdc.ConstTypePtr(coords_in_pattern.AsCTypePtr()), gdc.ConstTypePtr(pattern.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) SetPattern(layer int, position Vector2i, pattern TileMapPattern, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pattern")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1195853946) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(position.AsCTypePtr()), gdc.ConstTypePtr(pattern.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) SetCellsTerrainConnect(layer int, cells Vector2i, terrain_set int, terrain int, ignore_empty_terrains bool, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cells_terrain_connect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3578627656) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(cells.AsCTypePtr()), gdc.ConstTypePtr(&terrain_set), gdc.ConstTypePtr(&terrain), gdc.ConstTypePtr(&ignore_empty_terrains), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) SetCellsTerrainPath(layer int, path Vector2i, terrain_set int, terrain int, ignore_empty_terrains bool, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cells_terrain_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3578627656) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(&terrain_set), gdc.ConstTypePtr(&terrain), gdc.ConstTypePtr(&ignore_empty_terrains), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) FixInvalidTiles()  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("fix_invalid_tiles")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) ClearLayer(layer int, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) Clear()  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) UpdateInternals()  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("update_internals")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) NotifyRuntimeTileDataUpdate(layer int, )  {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("notify_runtime_tile_data_update")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1025054187) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMap) GetSurroundingCells(coords Vector2i, ) Vector2i {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_surrounding_cells")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2673526557) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(coords.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) GetUsedCells(layer int, ) Vector2i {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_used_cells")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 663333327) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) GetUsedCellsById(layer int, source_id int, atlas_coords Vector2i, alternative_tile int, ) Vector2i {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_used_cells_by_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2931012785) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(&source_id), gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), gdc.ConstTypePtr(&alternative_tile), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) GetUsedRect() Rect2i {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_used_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 410525958) // FIXME: should cache?
  var ret Rect2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) MapToLocal(map_position Vector2i, ) Vector2 {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_to_local")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 108438297) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(map_position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) LocalToMap(local_position Vector2, ) Vector2i {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("local_to_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 837806996) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(local_position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMap) GetNeighborCell(coords Vector2i, neighbor TileSetCellNeighbor, ) Vector2i {
  classNameV := StringNameFromStr("TileMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_neighbor_cell")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 986575103) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(coords.AsCTypePtr()), gdc.ConstTypePtr(&neighbor), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type TileMapChangedSignalFn func()

func (me *TileMap) ConnectChanged(subs SignalSubscribers, fn TileMapChangedSignalFn) {
  sig := StringNameFromStr("changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *TileMap) DisconnectChanged(subs SignalSubscribers, fn TileMapChangedSignalFn) {
  sig := StringNameFromStr("changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
