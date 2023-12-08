// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TileData struct {
  obj gdc.ObjectPtr
}

func createTileData(obj gdc.ObjectPtr) *TileData {
  return &TileData{
    obj: obj,
  }
}

func (me *TileData) BaseClass() string {
  return "TileData"
}
