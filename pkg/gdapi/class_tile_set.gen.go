// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TileSet struct {
  obj gdc.ObjectPtr
}

func createTileSet(obj gdc.ObjectPtr) *TileSet {
  return &TileSet{
    obj: obj,
  }
}

func (me *TileSet) BaseClass() string {
  return "TileSet"
}
