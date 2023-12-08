// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TileSetAtlasSource struct {
  obj gdc.ObjectPtr
}

func createTileSetAtlasSource(obj gdc.ObjectPtr) *TileSetAtlasSource {
  return &TileSetAtlasSource{
    obj: obj,
  }
}

func (me *TileSetAtlasSource) BaseClass() string {
  return "TileSetAtlasSource"
}
