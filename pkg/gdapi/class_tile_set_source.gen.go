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

func NewTileSetSource() *TileSetSource {
  str := StringNameFromStr("TileSetSource") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &TileSetSource{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *TileSetSource) GetTilesCount() int64 {
  classNameV := StringNameFromStr("TileSetSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tiles_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetSource) GetTileId(index int64, ) Vector2i {
  classNameV := StringNameFromStr("TileSetSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tile_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 880721226) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileSetSource) HasTile(atlas_coords Vector2i, ) bool {
  classNameV := StringNameFromStr("TileSetSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_tile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3900751641) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetSource) GetAlternativeTilesCount(atlas_coords Vector2i, ) int64 {
  classNameV := StringNameFromStr("TileSetSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alternative_tiles_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2485466453) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), }
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetSource) GetAlternativeTileId(atlas_coords Vector2i, index int64, ) int64 {
  classNameV := StringNameFromStr("TileSetSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alternative_tile_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 89881719) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), gdc.ConstTypePtr(&index), }
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetSource) HasAlternativeTile(atlas_coords Vector2i, alternative_tile int64, ) bool {
  classNameV := StringNameFromStr("TileSetSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_alternative_tile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1073731340) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), gdc.ConstTypePtr(&alternative_tile), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

// Signals
