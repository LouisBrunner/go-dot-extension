// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TileMap struct {
  obj gdc.ObjectPtr
}

func createTileMap(obj gdc.ObjectPtr) *TileMap {
  return &TileMap{
    obj: obj,
  }
}

func (me *TileMap) BaseClass() string {
  return "TileMap"
}
