// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CanvasItemMaterial struct {
  obj gdc.ObjectPtr
}

func (me *CanvasItemMaterial) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CanvasItemMaterial) BaseClass() string {
  return "CanvasItemMaterial"
}

type CanvasItemMaterialBlendMode int
const (
  CanvasItemMaterialBlendModeMix CanvasItemMaterialBlendMode = 0
  CanvasItemMaterialBlendModeAdd CanvasItemMaterialBlendMode = 1
  CanvasItemMaterialBlendModeSub CanvasItemMaterialBlendMode = 2
  CanvasItemMaterialBlendModeMul CanvasItemMaterialBlendMode = 3
  CanvasItemMaterialBlendModePremultAlpha CanvasItemMaterialBlendMode = 4
)

type CanvasItemMaterialLightMode int
const (
  CanvasItemMaterialLightModeNormal CanvasItemMaterialLightMode = 0
  CanvasItemMaterialLightModeUnshaded CanvasItemMaterialLightMode = 1
  CanvasItemMaterialLightModeLightOnly CanvasItemMaterialLightMode = 2
)
