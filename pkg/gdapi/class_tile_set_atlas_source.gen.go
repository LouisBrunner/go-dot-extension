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

type TileSetAtlasSource struct {
  TileSetSource
}

func (me *TileSetAtlasSource) BaseClass() string {
  return "TileSetAtlasSource"
}

func NewTileSetAtlasSource() *TileSetAtlasSource {
  str := StringNameFromStr("TileSetAtlasSource") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &TileSetAtlasSource{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Constants

var (
  TileSetAtlasSourceTransformFlipH = "4096" // TODO: construct correctly
  TileSetAtlasSourceTransformFlipV = "8192" // TODO: construct correctly
  TileSetAtlasSourceTransformTranspose = "16384" // TODO: construct correctly
)

// Enums

type TileSetAtlasSourceTileAnimationMode int
const (
  TileSetAtlasSourceTileAnimationModeTileAnimationModeDefault TileSetAtlasSourceTileAnimationMode = 0
  TileSetAtlasSourceTileAnimationModeTileAnimationModeRandomStartTimes TileSetAtlasSourceTileAnimationMode = 1
  TileSetAtlasSourceTileAnimationModeTileAnimationModeMax TileSetAtlasSourceTileAnimationMode = 2
)

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
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) GetTexture() Texture2D {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileSetAtlasSource) SetMargins(margins Vector2i, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_margins")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{margins.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) GetMargins() Vector2i {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_margins")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileSetAtlasSource) SetSeparation(separation Vector2i, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_separation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{separation.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) GetSeparation() Vector2i {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_separation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileSetAtlasSource) SetTextureRegionSize(texture_region_size Vector2i, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_region_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{texture_region_size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) GetTextureRegionSize() Vector2i {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_region_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileSetAtlasSource) SetUseTexturePadding(use_texture_padding bool, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_texture_padding")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_texture_padding) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) GetUseTexturePadding() bool {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_use_texture_padding")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetAtlasSource) CreateTile(atlas_coords Vector2i, size Vector2i, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_tile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 190528769) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) RemoveTile(atlas_coords Vector2i, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_tile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) MoveTileInAtlas(atlas_coords Vector2i, new_atlas_coords Vector2i, new_size Vector2i, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_tile_in_atlas")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3870111920) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), new_atlas_coords.AsCTypePtr(), new_size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) GetTileSizeInAtlas(atlas_coords Vector2i, ) Vector2i {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tile_size_in_atlas")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3050897911) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileSetAtlasSource) HasRoomForTile(atlas_coords Vector2i, size Vector2i, animation_columns int64, animation_separation Vector2i, frames_count int64, ignored_tile Vector2i, ) bool {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_room_for_tile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3018597268) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&animation_columns) , animation_separation.AsCTypePtr(), gdc.ConstTypePtr(&frames_count) , ignored_tile.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&animation_columns)
  pinner.Pin(&frames_count)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetAtlasSource) GetTilesToBeRemovedOnChange(texture Texture2D, margins Vector2i, separation Vector2i, texture_region_size Vector2i, ) PackedVector2Array {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tiles_to_be_removed_on_change")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1240378054) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), margins.AsCTypePtr(), separation.AsCTypePtr(), texture_region_size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileSetAtlasSource) GetTileAtCoords(atlas_coords Vector2i, ) Vector2i {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tile_at_coords")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3050897911) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileSetAtlasSource) HasTilesOutsideTexture() bool {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_tiles_outside_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetAtlasSource) ClearTilesOutsideTexture()  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_tiles_outside_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) SetTileAnimationColumns(atlas_coords Vector2i, frame_columns int64, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tile_animation_columns")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3200960707) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&frame_columns) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) GetTileAnimationColumns(atlas_coords Vector2i, ) int64 {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tile_animation_columns")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2485466453) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetAtlasSource) SetTileAnimationSeparation(atlas_coords Vector2i, separation Vector2i, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tile_animation_separation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1941061099) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), separation.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) GetTileAnimationSeparation(atlas_coords Vector2i, ) Vector2i {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tile_animation_separation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3050897911) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileSetAtlasSource) SetTileAnimationSpeed(atlas_coords Vector2i, speed float64, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tile_animation_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2262553149) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&speed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) GetTileAnimationSpeed(atlas_coords Vector2i, ) float64 {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tile_animation_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 719993801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetAtlasSource) SetTileAnimationMode(atlas_coords Vector2i, mode TileSetAtlasSourceTileAnimationMode, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tile_animation_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3192753483) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) GetTileAnimationMode(atlas_coords Vector2i, ) TileSetAtlasSourceTileAnimationMode {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tile_animation_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4025349959) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret TileSetAtlasSourceTileAnimationMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TileSetAtlasSource) SetTileAnimationFramesCount(atlas_coords Vector2i, frames_count int64, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tile_animation_frames_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3200960707) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&frames_count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) GetTileAnimationFramesCount(atlas_coords Vector2i, ) int64 {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tile_animation_frames_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2485466453) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetAtlasSource) SetTileAnimationFrameDuration(atlas_coords Vector2i, frame_index int64, duration float64, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tile_animation_frame_duration")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2843487787) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&frame_index) , gdc.ConstTypePtr(&duration) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) GetTileAnimationFrameDuration(atlas_coords Vector2i, frame_index int64, ) float64 {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tile_animation_frame_duration")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1802448425) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&frame_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&frame_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetAtlasSource) GetTileAnimationTotalDuration(atlas_coords Vector2i, ) float64 {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tile_animation_total_duration")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 719993801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetAtlasSource) CreateAlternativeTile(atlas_coords Vector2i, alternative_id_override int64, ) int64 {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_alternative_tile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2226298068) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&alternative_id_override) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&alternative_id_override)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetAtlasSource) RemoveAlternativeTile(atlas_coords Vector2i, alternative_tile int64, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_alternative_tile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3200960707) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&alternative_tile) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) SetAlternativeTileId(atlas_coords Vector2i, alternative_tile int64, new_id int64, )  {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alternative_tile_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1499785778) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&alternative_tile) , gdc.ConstTypePtr(&new_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) GetNextAlternativeTileId(atlas_coords Vector2i, ) int64 {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_next_alternative_tile_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2485466453) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetAtlasSource) GetTileData(atlas_coords Vector2i, alternative_tile int64, ) TileData {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tile_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3534028207) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&alternative_tile) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTileData()
  pinner.Pin(&alternative_tile)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileSetAtlasSource) GetAtlasGridSize() Vector2i {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_atlas_grid_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileSetAtlasSource) GetTileTextureRegion(atlas_coords Vector2i, frame int64, ) Rect2i {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tile_texture_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 241857547) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&frame) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRect2i()
  pinner.Pin(&frame)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileSetAtlasSource) GetRuntimeTexture() Texture2D {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_runtime_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileSetAtlasSource) GetRuntimeTileTextureRegion(atlas_coords Vector2i, frame int64, ) Rect2i {
  classNameV := StringNameFromStr("TileSetAtlasSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_runtime_tile_texture_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 104874263) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&frame) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRect2i()
  pinner.Pin(&frame)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
