// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TileMapPattern struct {
  obj gdc.ObjectPtr
}

func (me *TileMapPattern) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
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
  panic("TODO: implement")
}

func  (me *TileMapPattern) HasCell(coords Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileMapPattern) RemoveCell(coords Vector2i, update_size bool, )  {
  panic("TODO: implement")
}

func  (me *TileMapPattern) GetCellSourceId(coords Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileMapPattern) GetCellAtlasCoords(coords Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileMapPattern) GetCellAlternativeTile(coords Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileMapPattern) GetUsedCells()  {
  panic("TODO: implement")
}

func  (me *TileMapPattern) GetSize()  {
  panic("TODO: implement")
}

func  (me *TileMapPattern) SetSize(size Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileMapPattern) IsEmpty()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
