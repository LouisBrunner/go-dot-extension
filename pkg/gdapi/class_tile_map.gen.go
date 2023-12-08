// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TileMap struct {
  obj gdc.ObjectPtr
}

func (me *TileMap) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TileMap) BaseClass() string {
  return "TileMap"
}

type TileMapVisibilityMode int
const (
  TileMapVisibilityModeDefault TileMapVisibilityMode = 0
  TileMapVisibilityModeForceHide TileMapVisibilityMode = 2
  TileMapVisibilityModeForceShow TileMapVisibilityMode = 1
)
