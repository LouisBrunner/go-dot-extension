// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TileSetAtlasSource struct {
  obj gdc.ObjectPtr
}

func (me *TileSetAtlasSource) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TileSetAtlasSource) BaseClass() string {
  return "TileSetAtlasSource"
}



// Enums

func (me *TileSetAtlasSource) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TileSetAtlasSource) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TileSetAtlasSource) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *TileSetAtlasSource) SetTexture(texture Texture2D, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSetAtlasSource) GetTexture() Texture2D {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetAtlasSource) SetMargins(margins Vector2i, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_margins")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(margins.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSetAtlasSource) GetMargins() Vector2i {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_margins")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetAtlasSource) SetSeparation(separation Vector2i, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_separation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(separation.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSetAtlasSource) GetSeparation() Vector2i {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_separation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetAtlasSource) SetTextureRegionSize(texture_region_size Vector2i, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_region_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture_region_size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSetAtlasSource) GetTextureRegionSize() Vector2i {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_region_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetAtlasSource) SetUseTexturePadding(use_texture_padding bool, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_texture_padding")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_texture_padding), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSetAtlasSource) GetUseTexturePadding() bool {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_use_texture_padding")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetAtlasSource) CreateTile(atlas_coords Vector2i, size Vector2i, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_tile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1583819816) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSetAtlasSource) RemoveTile(atlas_coords Vector2i, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_tile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSetAtlasSource) MoveTileInAtlas(atlas_coords Vector2i, new_atlas_coords Vector2i, new_size Vector2i, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_tile_in_atlas")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1375626516) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), gdc.ConstTypePtr(new_atlas_coords.AsCTypePtr()), gdc.ConstTypePtr(new_size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSetAtlasSource) GetTileSizeInAtlas(atlas_coords Vector2i, ) Vector2i {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tile_size_in_atlas")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3050897911) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetAtlasSource) HasRoomForTile(atlas_coords Vector2i, size Vector2i, animation_columns int, animation_separation Vector2i, frames_count int, ignored_tile Vector2i, ) bool {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_room_for_tile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4182444377) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&animation_columns), gdc.ConstTypePtr(animation_separation.AsCTypePtr()), gdc.ConstTypePtr(&frames_count), gdc.ConstTypePtr(ignored_tile.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetAtlasSource) GetTilesToBeRemovedOnChange(texture Texture2D, margins Vector2i, separation Vector2i, texture_region_size Vector2i, ) PackedVector2Array {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tiles_to_be_removed_on_change")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1240378054) // FIXME: should cache?
  var ret PackedVector2Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), gdc.ConstTypePtr(margins.AsCTypePtr()), gdc.ConstTypePtr(separation.AsCTypePtr()), gdc.ConstTypePtr(texture_region_size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetAtlasSource) GetTileAtCoords(atlas_coords Vector2i, ) Vector2i {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tile_at_coords")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3050897911) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetAtlasSource) SetTileAnimationColumns(atlas_coords Vector2i, frame_columns int, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tile_animation_columns")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3200960707) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), gdc.ConstTypePtr(&frame_columns), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSetAtlasSource) GetTileAnimationColumns(atlas_coords Vector2i, ) int {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tile_animation_columns")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2485466453) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetAtlasSource) SetTileAnimationSeparation(atlas_coords Vector2i, separation Vector2i, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tile_animation_separation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1941061099) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), gdc.ConstTypePtr(separation.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSetAtlasSource) GetTileAnimationSeparation(atlas_coords Vector2i, ) Vector2i {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tile_animation_separation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3050897911) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetAtlasSource) SetTileAnimationSpeed(atlas_coords Vector2i, speed float32, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tile_animation_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2262553149) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), gdc.ConstTypePtr(&speed), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSetAtlasSource) GetTileAnimationSpeed(atlas_coords Vector2i, ) float32 {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tile_animation_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 719993801) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetAtlasSource) SetTileAnimationFramesCount(atlas_coords Vector2i, frames_count int, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tile_animation_frames_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3200960707) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), gdc.ConstTypePtr(&frames_count), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSetAtlasSource) GetTileAnimationFramesCount(atlas_coords Vector2i, ) int {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tile_animation_frames_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2485466453) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetAtlasSource) SetTileAnimationFrameDuration(atlas_coords Vector2i, frame_index int, duration float32, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tile_animation_frame_duration")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2843487787) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), gdc.ConstTypePtr(&frame_index), gdc.ConstTypePtr(&duration), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSetAtlasSource) GetTileAnimationFrameDuration(atlas_coords Vector2i, frame_index int, ) float32 {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tile_animation_frame_duration")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1802448425) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), gdc.ConstTypePtr(&frame_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetAtlasSource) GetTileAnimationTotalDuration(atlas_coords Vector2i, ) float32 {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tile_animation_total_duration")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 719993801) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetAtlasSource) CreateAlternativeTile(atlas_coords Vector2i, alternative_id_override int, ) int {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_alternative_tile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3531100812) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), gdc.ConstTypePtr(&alternative_id_override), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetAtlasSource) RemoveAlternativeTile(atlas_coords Vector2i, alternative_tile int, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_alternative_tile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3200960707) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), gdc.ConstTypePtr(&alternative_tile), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSetAtlasSource) SetAlternativeTileId(atlas_coords Vector2i, alternative_tile int, new_id int, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alternative_tile_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1499785778) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), gdc.ConstTypePtr(&alternative_tile), gdc.ConstTypePtr(&new_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileSetAtlasSource) GetNextAlternativeTileId(atlas_coords Vector2i, ) int {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_next_alternative_tile_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2485466453) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetAtlasSource) GetTileData(atlas_coords Vector2i, alternative_tile int, ) TileData {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tile_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3534028207) // FIXME: should cache?
  var ret TileData
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), gdc.ConstTypePtr(&alternative_tile), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetAtlasSource) GetAtlasGridSize() Vector2i {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_atlas_grid_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetAtlasSource) GetTileTextureRegion(atlas_coords Vector2i, frame int, ) Rect2i {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tile_texture_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1321423751) // FIXME: should cache?
  var ret Rect2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), gdc.ConstTypePtr(&frame), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetAtlasSource) GetRuntimeTexture() Texture2D {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_runtime_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetAtlasSource) GetRuntimeTileTextureRegion(atlas_coords Vector2i, frame int, ) Rect2i {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_runtime_tile_texture_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 104874263) // FIXME: should cache?
  var ret Rect2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), gdc.ConstTypePtr(&frame), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *TileSetAtlasSource) GetPropTexture() Texture2D {
  panic("TODO: implement")
}

func (me *TileSetAtlasSource) SetPropTexture(value Texture2D) {
  panic("TODO: implement")
}

func (me *TileSetAtlasSource) GetPropMargins() Vector2i {
  panic("TODO: implement")
}

func (me *TileSetAtlasSource) SetPropMargins(value Vector2i) {
  panic("TODO: implement")
}

func (me *TileSetAtlasSource) GetPropSeparation() Vector2i {
  panic("TODO: implement")
}

func (me *TileSetAtlasSource) SetPropSeparation(value Vector2i) {
  panic("TODO: implement")
}

func (me *TileSetAtlasSource) GetPropTextureRegionSize() Vector2i {
  panic("TODO: implement")
}

func (me *TileSetAtlasSource) SetPropTextureRegionSize(value Vector2i) {
  panic("TODO: implement")
}

func (me *TileSetAtlasSource) GetPropUseTexturePadding() bool {
  panic("TODO: implement")
}

func (me *TileSetAtlasSource) SetPropUseTexturePadding(value bool) {
  panic("TODO: implement")
}