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

type ptrsForTileSetAtlasSourceList struct {
  fnSetTexture gdc.MethodBindPtr
  fnGetTexture gdc.MethodBindPtr
  fnSetMargins gdc.MethodBindPtr
  fnGetMargins gdc.MethodBindPtr
  fnSetSeparation gdc.MethodBindPtr
  fnGetSeparation gdc.MethodBindPtr
  fnSetTextureRegionSize gdc.MethodBindPtr
  fnGetTextureRegionSize gdc.MethodBindPtr
  fnSetUseTexturePadding gdc.MethodBindPtr
  fnGetUseTexturePadding gdc.MethodBindPtr
  fnCreateTile gdc.MethodBindPtr
  fnRemoveTile gdc.MethodBindPtr
  fnMoveTileInAtlas gdc.MethodBindPtr
  fnGetTileSizeInAtlas gdc.MethodBindPtr
  fnHasRoomForTile gdc.MethodBindPtr
  fnGetTilesToBeRemovedOnChange gdc.MethodBindPtr
  fnGetTileAtCoords gdc.MethodBindPtr
  fnHasTilesOutsideTexture gdc.MethodBindPtr
  fnClearTilesOutsideTexture gdc.MethodBindPtr
  fnSetTileAnimationColumns gdc.MethodBindPtr
  fnGetTileAnimationColumns gdc.MethodBindPtr
  fnSetTileAnimationSeparation gdc.MethodBindPtr
  fnGetTileAnimationSeparation gdc.MethodBindPtr
  fnSetTileAnimationSpeed gdc.MethodBindPtr
  fnGetTileAnimationSpeed gdc.MethodBindPtr
  fnSetTileAnimationMode gdc.MethodBindPtr
  fnGetTileAnimationMode gdc.MethodBindPtr
  fnSetTileAnimationFramesCount gdc.MethodBindPtr
  fnGetTileAnimationFramesCount gdc.MethodBindPtr
  fnSetTileAnimationFrameDuration gdc.MethodBindPtr
  fnGetTileAnimationFrameDuration gdc.MethodBindPtr
  fnGetTileAnimationTotalDuration gdc.MethodBindPtr
  fnCreateAlternativeTile gdc.MethodBindPtr
  fnRemoveAlternativeTile gdc.MethodBindPtr
  fnSetAlternativeTileId gdc.MethodBindPtr
  fnGetNextAlternativeTileId gdc.MethodBindPtr
  fnGetTileData gdc.MethodBindPtr
  fnGetAtlasGridSize gdc.MethodBindPtr
  fnGetTileTextureRegion gdc.MethodBindPtr
  fnGetRuntimeTexture gdc.MethodBindPtr
  fnGetRuntimeTileTextureRegion gdc.MethodBindPtr
}

var ptrsForTileSetAtlasSource ptrsForTileSetAtlasSourceList

func initTileSetAtlasSourcePtrs(iface gdc.Interface) {

  className := StringNameFromStr("TileSetAtlasSource")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_texture")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnSetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
  }
  {
    methodName := StringNameFromStr("get_texture")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnGetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
  }
  {
    methodName := StringNameFromStr("set_margins")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnSetMargins = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
  }
  {
    methodName := StringNameFromStr("get_margins")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnGetMargins = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
  }
  {
    methodName := StringNameFromStr("set_separation")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnSetSeparation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
  }
  {
    methodName := StringNameFromStr("get_separation")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnGetSeparation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
  }
  {
    methodName := StringNameFromStr("set_texture_region_size")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnSetTextureRegionSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
  }
  {
    methodName := StringNameFromStr("get_texture_region_size")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnGetTextureRegionSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
  }
  {
    methodName := StringNameFromStr("set_use_texture_padding")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnSetUseTexturePadding = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_use_texture_padding")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnGetUseTexturePadding = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("create_tile")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnCreateTile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 190528769))
  }
  {
    methodName := StringNameFromStr("remove_tile")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnRemoveTile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
  }
  {
    methodName := StringNameFromStr("move_tile_in_atlas")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnMoveTileInAtlas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3870111920))
  }
  {
    methodName := StringNameFromStr("get_tile_size_in_atlas")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnGetTileSizeInAtlas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3050897911))
  }
  {
    methodName := StringNameFromStr("has_room_for_tile")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnHasRoomForTile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3018597268))
  }
  {
    methodName := StringNameFromStr("get_tiles_to_be_removed_on_change")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnGetTilesToBeRemovedOnChange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1240378054))
  }
  {
    methodName := StringNameFromStr("get_tile_at_coords")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnGetTileAtCoords = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3050897911))
  }
  {
    methodName := StringNameFromStr("has_tiles_outside_texture")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnHasTilesOutsideTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("clear_tiles_outside_texture")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnClearTilesOutsideTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_tile_animation_columns")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnSetTileAnimationColumns = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3200960707))
  }
  {
    methodName := StringNameFromStr("get_tile_animation_columns")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnGetTileAnimationColumns = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2485466453))
  }
  {
    methodName := StringNameFromStr("set_tile_animation_separation")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnSetTileAnimationSeparation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1941061099))
  }
  {
    methodName := StringNameFromStr("get_tile_animation_separation")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnGetTileAnimationSeparation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3050897911))
  }
  {
    methodName := StringNameFromStr("set_tile_animation_speed")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnSetTileAnimationSpeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2262553149))
  }
  {
    methodName := StringNameFromStr("get_tile_animation_speed")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnGetTileAnimationSpeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 719993801))
  }
  {
    methodName := StringNameFromStr("set_tile_animation_mode")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnSetTileAnimationMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3192753483))
  }
  {
    methodName := StringNameFromStr("get_tile_animation_mode")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnGetTileAnimationMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4025349959))
  }
  {
    methodName := StringNameFromStr("set_tile_animation_frames_count")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnSetTileAnimationFramesCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3200960707))
  }
  {
    methodName := StringNameFromStr("get_tile_animation_frames_count")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnGetTileAnimationFramesCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2485466453))
  }
  {
    methodName := StringNameFromStr("set_tile_animation_frame_duration")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnSetTileAnimationFrameDuration = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2843487787))
  }
  {
    methodName := StringNameFromStr("get_tile_animation_frame_duration")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnGetTileAnimationFrameDuration = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1802448425))
  }
  {
    methodName := StringNameFromStr("get_tile_animation_total_duration")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnGetTileAnimationTotalDuration = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 719993801))
  }
  {
    methodName := StringNameFromStr("create_alternative_tile")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnCreateAlternativeTile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2226298068))
  }
  {
    methodName := StringNameFromStr("remove_alternative_tile")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnRemoveAlternativeTile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3200960707))
  }
  {
    methodName := StringNameFromStr("set_alternative_tile_id")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnSetAlternativeTileId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1499785778))
  }
  {
    methodName := StringNameFromStr("get_next_alternative_tile_id")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnGetNextAlternativeTileId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2485466453))
  }
  {
    methodName := StringNameFromStr("get_tile_data")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnGetTileData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3534028207))
  }
  {
    methodName := StringNameFromStr("get_atlas_grid_size")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnGetAtlasGridSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
  }
  {
    methodName := StringNameFromStr("get_tile_texture_region")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnGetTileTextureRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 241857547))
  }
  {
    methodName := StringNameFromStr("get_runtime_texture")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnGetRuntimeTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
  }
  {
    methodName := StringNameFromStr("get_runtime_tile_texture_region")
    defer methodName.Destroy()
    ptrsForTileSetAtlasSource.fnGetRuntimeTileTextureRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 104874263))
  }
}

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
  TileSetAtlasSourceTransformFlipH = 4096
  TileSetAtlasSourceTransformFlipV = 8192
  TileSetAtlasSourceTransformTranspose = 16384
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
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnSetTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) GetTexture() Texture2D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnGetTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileSetAtlasSource) SetMargins(margins Vector2i, )  {
  cargs := []gdc.ConstTypePtr{margins.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnSetMargins), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) GetMargins() Vector2i {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnGetMargins), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileSetAtlasSource) SetSeparation(separation Vector2i, )  {
  cargs := []gdc.ConstTypePtr{separation.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnSetSeparation), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) GetSeparation() Vector2i {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnGetSeparation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileSetAtlasSource) SetTextureRegionSize(texture_region_size Vector2i, )  {
  cargs := []gdc.ConstTypePtr{texture_region_size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnSetTextureRegionSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) GetTextureRegionSize() Vector2i {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnGetTextureRegionSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileSetAtlasSource) SetUseTexturePadding(use_texture_padding bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_texture_padding) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnSetUseTexturePadding), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) GetUseTexturePadding() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnGetUseTexturePadding), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetAtlasSource) CreateTile(atlas_coords Vector2i, size Vector2i, )  {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnCreateTile), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) RemoveTile(atlas_coords Vector2i, )  {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnRemoveTile), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) MoveTileInAtlas(atlas_coords Vector2i, new_atlas_coords Vector2i, new_size Vector2i, )  {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), new_atlas_coords.AsCTypePtr(), new_size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnMoveTileInAtlas), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) GetTileSizeInAtlas(atlas_coords Vector2i, ) Vector2i {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnGetTileSizeInAtlas), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileSetAtlasSource) HasRoomForTile(atlas_coords Vector2i, size Vector2i, animation_columns int64, animation_separation Vector2i, frames_count int64, ignored_tile Vector2i, ) bool {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&animation_columns) , animation_separation.AsCTypePtr(), gdc.ConstTypePtr(&frames_count) , ignored_tile.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&animation_columns)
  pinner.Pin(&frames_count)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnHasRoomForTile), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetAtlasSource) GetTilesToBeRemovedOnChange(texture Texture2D, margins Vector2i, separation Vector2i, texture_region_size Vector2i, ) PackedVector2Array {
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), margins.AsCTypePtr(), separation.AsCTypePtr(), texture_region_size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnGetTilesToBeRemovedOnChange), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileSetAtlasSource) GetTileAtCoords(atlas_coords Vector2i, ) Vector2i {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnGetTileAtCoords), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileSetAtlasSource) HasTilesOutsideTexture() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnHasTilesOutsideTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetAtlasSource) ClearTilesOutsideTexture()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnClearTilesOutsideTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) SetTileAnimationColumns(atlas_coords Vector2i, frame_columns int64, )  {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&frame_columns) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnSetTileAnimationColumns), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) GetTileAnimationColumns(atlas_coords Vector2i, ) int64 {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnGetTileAnimationColumns), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetAtlasSource) SetTileAnimationSeparation(atlas_coords Vector2i, separation Vector2i, )  {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), separation.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnSetTileAnimationSeparation), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) GetTileAnimationSeparation(atlas_coords Vector2i, ) Vector2i {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnGetTileAnimationSeparation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileSetAtlasSource) SetTileAnimationSpeed(atlas_coords Vector2i, speed float64, )  {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&speed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnSetTileAnimationSpeed), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) GetTileAnimationSpeed(atlas_coords Vector2i, ) float64 {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnGetTileAnimationSpeed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetAtlasSource) SetTileAnimationMode(atlas_coords Vector2i, mode TileSetAtlasSourceTileAnimationMode, )  {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnSetTileAnimationMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) GetTileAnimationMode(atlas_coords Vector2i, ) TileSetAtlasSourceTileAnimationMode {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret TileSetAtlasSourceTileAnimationMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnGetTileAnimationMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TileSetAtlasSource) SetTileAnimationFramesCount(atlas_coords Vector2i, frames_count int64, )  {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&frames_count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnSetTileAnimationFramesCount), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) GetTileAnimationFramesCount(atlas_coords Vector2i, ) int64 {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnGetTileAnimationFramesCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetAtlasSource) SetTileAnimationFrameDuration(atlas_coords Vector2i, frame_index int64, duration float64, )  {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&frame_index) , gdc.ConstTypePtr(&duration) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnSetTileAnimationFrameDuration), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) GetTileAnimationFrameDuration(atlas_coords Vector2i, frame_index int64, ) float64 {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&frame_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&frame_index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnGetTileAnimationFrameDuration), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetAtlasSource) GetTileAnimationTotalDuration(atlas_coords Vector2i, ) float64 {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnGetTileAnimationTotalDuration), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetAtlasSource) CreateAlternativeTile(atlas_coords Vector2i, alternative_id_override int64, ) int64 {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&alternative_id_override) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&alternative_id_override)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnCreateAlternativeTile), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetAtlasSource) RemoveAlternativeTile(atlas_coords Vector2i, alternative_tile int64, )  {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&alternative_tile) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnRemoveAlternativeTile), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) SetAlternativeTileId(atlas_coords Vector2i, alternative_tile int64, new_id int64, )  {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&alternative_tile) , gdc.ConstTypePtr(&new_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnSetAlternativeTileId), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileSetAtlasSource) GetNextAlternativeTileId(atlas_coords Vector2i, ) int64 {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnGetNextAlternativeTileId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetAtlasSource) GetTileData(atlas_coords Vector2i, alternative_tile int64, ) TileData {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&alternative_tile) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTileData()
  pinner.Pin(&alternative_tile)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnGetTileData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileSetAtlasSource) GetAtlasGridSize() Vector2i {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnGetAtlasGridSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileSetAtlasSource) GetTileTextureRegion(atlas_coords Vector2i, frame int64, ) Rect2i {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&frame) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRect2i()
  pinner.Pin(&frame)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnGetTileTextureRegion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileSetAtlasSource) GetRuntimeTexture() Texture2D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnGetRuntimeTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileSetAtlasSource) GetRuntimeTileTextureRegion(atlas_coords Vector2i, frame int64, ) Rect2i {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&frame) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRect2i()
  pinner.Pin(&frame)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetAtlasSource.fnGetRuntimeTileTextureRegion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
