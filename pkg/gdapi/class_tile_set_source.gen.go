// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TileSetSource struct {
  obj gdc.ObjectPtr
}

func createTileSetSource(obj gdc.ObjectPtr) *TileSetSource {
  return &TileSetSource{
    obj: obj,
  }
}

func (me *TileSetSource) BaseClass() string {
  return "TileSetSource"
}
