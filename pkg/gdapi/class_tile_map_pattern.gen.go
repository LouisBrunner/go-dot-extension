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

type ptrsForTileMapPatternList struct {
  fnSetCell gdc.MethodBindPtr
  fnHasCell gdc.MethodBindPtr
  fnRemoveCell gdc.MethodBindPtr
  fnGetCellSourceId gdc.MethodBindPtr
  fnGetCellAtlasCoords gdc.MethodBindPtr
  fnGetCellAlternativeTile gdc.MethodBindPtr
  fnGetUsedCells gdc.MethodBindPtr
  fnGetSize gdc.MethodBindPtr
  fnSetSize gdc.MethodBindPtr
  fnIsEmpty gdc.MethodBindPtr
}

var ptrsForTileMapPattern ptrsForTileMapPatternList

func initTileMapPatternPtrs(iface gdc.Interface) {

  className := StringNameFromStr("TileMapPattern")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_cell")
    defer methodName.Destroy()
    ptrsForTileMapPattern.fnSetCell = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2224802556))
  }
  {
    methodName := StringNameFromStr("has_cell")
    defer methodName.Destroy()
    ptrsForTileMapPattern.fnHasCell = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3900751641))
  }
  {
    methodName := StringNameFromStr("remove_cell")
    defer methodName.Destroy()
    ptrsForTileMapPattern.fnRemoveCell = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4153096796))
  }
  {
    methodName := StringNameFromStr("get_cell_source_id")
    defer methodName.Destroy()
    ptrsForTileMapPattern.fnGetCellSourceId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2485466453))
  }
  {
    methodName := StringNameFromStr("get_cell_atlas_coords")
    defer methodName.Destroy()
    ptrsForTileMapPattern.fnGetCellAtlasCoords = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3050897911))
  }
  {
    methodName := StringNameFromStr("get_cell_alternative_tile")
    defer methodName.Destroy()
    ptrsForTileMapPattern.fnGetCellAlternativeTile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2485466453))
  }
  {
    methodName := StringNameFromStr("get_used_cells")
    defer methodName.Destroy()
    ptrsForTileMapPattern.fnGetUsedCells = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
  {
    methodName := StringNameFromStr("get_size")
    defer methodName.Destroy()
    ptrsForTileMapPattern.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
  }
  {
    methodName := StringNameFromStr("set_size")
    defer methodName.Destroy()
    ptrsForTileMapPattern.fnSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
  }
  {
    methodName := StringNameFromStr("is_empty")
    defer methodName.Destroy()
    ptrsForTileMapPattern.fnIsEmpty = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

type TileMapPattern struct {
  Resource
}

func (me *TileMapPattern) BaseClass() string {
  return "TileMapPattern"
}

func NewTileMapPattern() *TileMapPattern {
  str := StringNameFromStr("TileMapPattern") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &TileMapPattern{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *TileMapPattern) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TileMapPattern) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TileMapPattern) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *TileMapPattern) SetCell(coords Vector2i, source_id int64, atlas_coords Vector2i, alternative_tile int64, )  {
  cargs := []gdc.ConstTypePtr{coords.AsCTypePtr(), gdc.ConstTypePtr(&source_id) , atlas_coords.AsCTypePtr(), gdc.ConstTypePtr(&alternative_tile) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapPattern.fnSetCell), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMapPattern) HasCell(coords Vector2i, ) bool {
  cargs := []gdc.ConstTypePtr{coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapPattern.fnHasCell), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileMapPattern) RemoveCell(coords Vector2i, update_size bool, )  {
  cargs := []gdc.ConstTypePtr{coords.AsCTypePtr(), gdc.ConstTypePtr(&update_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapPattern.fnRemoveCell), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMapPattern) GetCellSourceId(coords Vector2i, ) int64 {
  cargs := []gdc.ConstTypePtr{coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapPattern.fnGetCellSourceId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileMapPattern) GetCellAtlasCoords(coords Vector2i, ) Vector2i {
  cargs := []gdc.ConstTypePtr{coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapPattern.fnGetCellAtlasCoords), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileMapPattern) GetCellAlternativeTile(coords Vector2i, ) int64 {
  cargs := []gdc.ConstTypePtr{coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapPattern.fnGetCellAlternativeTile), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TileMapPattern) GetUsedCells() []Vector2i {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapPattern.fnGetUsedCells), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Vector2i](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *TileMapPattern) GetSize() Vector2i {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapPattern.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TileMapPattern) SetSize(size Vector2i, )  {
  cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapPattern.fnSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TileMapPattern) IsEmpty() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTileMapPattern.fnIsEmpty), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

// Signals
