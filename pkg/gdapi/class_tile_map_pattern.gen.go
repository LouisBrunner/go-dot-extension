// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TileMapPattern struct {
  Resource
}

func (me *TileMapPattern) BaseClass() string {
  return "TileMapPattern"
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

func  (me *TileMapPattern) SetCell(coords Vector2i, source_id int, atlas_coords Vector2i, alternative_tile int, )  {
  classNameV := StringNameFromStr("TileMapPattern")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cell")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2224802556) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(coords.AsCTypePtr()), gdc.ConstTypePtr(&source_id), gdc.ConstTypePtr(atlas_coords.AsCTypePtr()), gdc.ConstTypePtr(&alternative_tile), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMapPattern) HasCell(coords Vector2i, ) bool {
  classNameV := StringNameFromStr("TileMapPattern")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_cell")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3900751641) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(coords.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMapPattern) RemoveCell(coords Vector2i, update_size bool, )  {
  classNameV := StringNameFromStr("TileMapPattern")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_cell")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4153096796) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(coords.AsCTypePtr()), gdc.ConstTypePtr(&update_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMapPattern) GetCellSourceId(coords Vector2i, ) int {
  classNameV := StringNameFromStr("TileMapPattern")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cell_source_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2485466453) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(coords.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMapPattern) GetCellAtlasCoords(coords Vector2i, ) Vector2i {
  classNameV := StringNameFromStr("TileMapPattern")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cell_atlas_coords")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3050897911) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(coords.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMapPattern) GetCellAlternativeTile(coords Vector2i, ) int {
  classNameV := StringNameFromStr("TileMapPattern")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cell_alternative_tile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2485466453) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(coords.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMapPattern) GetUsedCells() Vector2i {
  classNameV := StringNameFromStr("TileMapPattern")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_used_cells")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMapPattern) GetSize() Vector2i {
  classNameV := StringNameFromStr("TileMapPattern")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TileMapPattern) SetSize(size Vector2i, )  {
  classNameV := StringNameFromStr("TileMapPattern")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TileMapPattern) IsEmpty() bool {
  classNameV := StringNameFromStr("TileMapPattern")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_empty")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
