// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SurfaceTool struct {
  obj gdc.ObjectPtr
}

func (me *SurfaceTool) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SurfaceTool) BaseClass() string {
  return "SurfaceTool"
}

type SurfaceToolCustomFormat int
const (
  SurfaceToolCustomRgba8Unorm SurfaceToolCustomFormat = 0
  SurfaceToolCustomRgba8Snorm SurfaceToolCustomFormat = 1
  SurfaceToolCustomRgHalf SurfaceToolCustomFormat = 2
  SurfaceToolCustomRgbaHalf SurfaceToolCustomFormat = 3
  SurfaceToolCustomRFloat SurfaceToolCustomFormat = 4
  SurfaceToolCustomRgFloat SurfaceToolCustomFormat = 5
  SurfaceToolCustomRgbFloat SurfaceToolCustomFormat = 6
  SurfaceToolCustomRgbaFloat SurfaceToolCustomFormat = 7
  SurfaceToolCustomMax SurfaceToolCustomFormat = 8
)

type SurfaceToolSkinWeightCount int
const (
  SurfaceToolSkin4Weights SurfaceToolSkinWeightCount = 0
  SurfaceToolSkin8Weights SurfaceToolSkinWeightCount = 1
)
