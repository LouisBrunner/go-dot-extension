// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TileSetSource struct {
  obj gdc.ObjectPtr
}

func (me *TileSetSource) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TileSetSource) BaseClass() string {
  return "TileSetSource"
}



// Enums

func (me *TileSetSource) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TileSetSource) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *TileSetSource) GetTilesCount()  {
  panic("TODO: implement")
}

func  (me *TileSetSource) GetTileId(index int, )  {
  panic("TODO: implement")
}

func  (me *TileSetSource) HasTile(atlas_coords Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileSetSource) GetAlternativeTilesCount(atlas_coords Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TileSetSource) GetAlternativeTileId(atlas_coords Vector2i, index int, )  {
  panic("TODO: implement")
}

func  (me *TileSetSource) HasAlternativeTile(atlas_coords Vector2i, alternative_tile int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
