// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TileMapPattern struct {
  obj gdc.ObjectPtr
}

func createTileMapPattern(obj gdc.ObjectPtr) *TileMapPattern {
  return &TileMapPattern{
    obj: obj,
  }
}

func (me *TileMapPattern) BaseClass() string {
  return "TileMapPattern"
}
