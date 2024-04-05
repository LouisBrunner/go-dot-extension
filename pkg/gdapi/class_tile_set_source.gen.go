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

type ptrsForTileSetSourceList struct {
  fnGetTilesCount gdc.MethodBindPtr
  fnGetTileId gdc.MethodBindPtr
  fnHasTile gdc.MethodBindPtr
  fnGetAlternativeTilesCount gdc.MethodBindPtr
  fnGetAlternativeTileId gdc.MethodBindPtr
  fnHasAlternativeTile gdc.MethodBindPtr
}

var ptrsForTileSetSource ptrsForTileSetSourceList

func initTileSetSourcePtrs(iface gdc.Interface) {

  className := StringNameFromStr("TileSetSource")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_tiles_count")
    defer methodName.Destroy()
    ptrsForTileSetSource.fnGetTilesCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_tile_id")
    defer methodName.Destroy()
    ptrsForTileSetSource.fnGetTileId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 880721226))
  }
  {
    methodName := StringNameFromStr("has_tile")
    defer methodName.Destroy()
    ptrsForTileSetSource.fnHasTile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3900751641))
  }
  {
    methodName := StringNameFromStr("get_alternative_tiles_count")
    defer methodName.Destroy()
    ptrsForTileSetSource.fnGetAlternativeTilesCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2485466453))
  }
  {
    methodName := StringNameFromStr("get_alternative_tile_id")
    defer methodName.Destroy()
    ptrsForTileSetSource.fnGetAlternativeTileId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 89881719))
  }
  {
    methodName := StringNameFromStr("has_alternative_tile")
    defer methodName.Destroy()
    ptrsForTileSetSource.fnHasAlternativeTile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1073731340))
  }
}

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
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetSource.fnGetTilesCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetSource) GetTileId(index int64, ) Vector2i {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetSource.fnGetTileId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileSetSource) HasTile(atlas_coords Vector2i, ) bool {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetSource.fnHasTile), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetSource) GetAlternativeTilesCount(atlas_coords Vector2i, ) int64 {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetSource.fnGetAlternativeTilesCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetSource) GetAlternativeTileId(atlas_coords Vector2i, index int64, ) int64 {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetSource.fnGetAlternativeTileId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileSetSource) HasAlternativeTile(atlas_coords Vector2i, alternative_tile int64, ) bool {
  cargs := []gdc.ConstTypePtr{atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&alternative_tile) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&alternative_tile)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileSetSource.fnHasAlternativeTile), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

// Signals
