// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func (me *TileSetAtlasSource) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TileSetAtlasSource) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *TileSetAtlasSource) SetTexture(texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) GetTexture()  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) SetMargins(margins Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) GetMargins()  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) SetSeparation(separation Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) GetSeparation()  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) SetTextureRegionSize(texture_region_size Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) GetTextureRegionSize()  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) SetUseTexturePadding(use_texture_padding bool, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) GetUseTexturePadding()  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) CreateTile(atlas_coords Vector2i, size Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) RemoveTile(atlas_coords Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) MoveTileInAtlas(atlas_coords Vector2i, new_atlas_coords Vector2i, new_size Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) GetTileSizeInAtlas(atlas_coords Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) HasRoomForTile(atlas_coords Vector2i, size Vector2i, animation_columns int, animation_separation Vector2i, frames_count int, ignored_tile Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) GetTilesToBeRemovedOnChange(texture Texture2D, margins Vector2i, separation Vector2i, texture_region_size Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) GetTileAtCoords(atlas_coords Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) SetTileAnimationColumns(atlas_coords Vector2i, frame_columns int, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) GetTileAnimationColumns(atlas_coords Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) SetTileAnimationSeparation(atlas_coords Vector2i, separation Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) GetTileAnimationSeparation(atlas_coords Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) SetTileAnimationSpeed(atlas_coords Vector2i, speed float32, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) GetTileAnimationSpeed(atlas_coords Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) SetTileAnimationFramesCount(atlas_coords Vector2i, frames_count int, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) GetTileAnimationFramesCount(atlas_coords Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) SetTileAnimationFrameDuration(atlas_coords Vector2i, frame_index int, duration float32, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) GetTileAnimationFrameDuration(atlas_coords Vector2i, frame_index int, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) GetTileAnimationTotalDuration(atlas_coords Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) CreateAlternativeTile(atlas_coords Vector2i, alternative_id_override int, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) RemoveAlternativeTile(atlas_coords Vector2i, alternative_tile int, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) SetAlternativeTileId(atlas_coords Vector2i, alternative_tile int, new_id int, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) GetNextAlternativeTileId(atlas_coords Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) GetTileData(atlas_coords Vector2i, alternative_tile int, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) GetAtlasGridSize()  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) GetTileTextureRegion(atlas_coords Vector2i, frame int, )  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) GetRuntimeTexture()  {
  panic("TODO: implement")
}

func  (me *TileSetAtlasSource) GetRuntimeTileTextureRegion(atlas_coords Vector2i, frame int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
