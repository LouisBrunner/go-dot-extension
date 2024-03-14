// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TileSetSource struct {
  Resource
}

func (me *TileSetSource) BaseClass() string {
  return "TileSetSource"
}



// Enums

func (me *TileSetSource) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TileSetSource) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TileSetSource) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *TileSetSource) GetTilesCount() int {
  classNameV := StringNameFromStr("TileSetSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tiles_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetSource) GetTileId(index int, ) Vector2i {
  classNameV := StringNameFromStr("TileSetSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tile_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 880721226) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetSource) HasTile(atlas_coords Vector2i, ) bool {
  classNameV := StringNameFromStr("TileSetSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_tile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3900751641) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetSource) GetAlternativeTilesCount(atlas_coords Vector2i, ) int {
  classNameV := StringNameFromStr("TileSetSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alternative_tiles_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2485466453) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetSource) GetAlternativeTileId(atlas_coords Vector2i, index int, ) int {
  classNameV := StringNameFromStr("TileSetSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alternative_tile_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 89881719) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileSetSource) HasAlternativeTile(atlas_coords Vector2i, alternative_tile int, ) bool {
  classNameV := StringNameFromStr("TileSetSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_alternative_tile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1073731340) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), gdc.ConstTypePtr(&alternative_tile), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
