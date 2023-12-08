// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Light2D struct {
  obj gdc.ObjectPtr
}

func (me *Light2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Light2D) BaseClass() string {
  return "Light2D"
}

type Light2DShadowFilter int
const (
  Light2DShadowFilterShadowFilterNone Light2DShadowFilter = 0
  Light2DShadowFilterShadowFilterPcf5 Light2DShadowFilter = 1
  Light2DShadowFilterShadowFilterPcf13 Light2DShadowFilter = 2
)

type Light2DBlendMode int
const (
  Light2DBlendModeBlendModeAdd Light2DBlendMode = 0
  Light2DBlendModeBlendModeSub Light2DBlendMode = 1
  Light2DBlendModeBlendModeMix Light2DBlendMode = 2
)

func  (me *Light2D) SetEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) IsEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) SetEditorOnly(editor_only bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) IsEditorOnly() { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) SetColor(color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) GetColor() { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) SetEnergy(energy float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) GetEnergy() { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) SetZRangeMin(z int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) GetZRangeMin() { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) SetZRangeMax(z int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) GetZRangeMax() { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) SetLayerRangeMin(layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) GetLayerRangeMin() { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) SetLayerRangeMax(layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) GetLayerRangeMax() { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) SetItemCullMask(item_cull_mask int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) GetItemCullMask() { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) SetItemShadowCullMask(item_shadow_cull_mask int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) GetItemShadowCullMask() { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) SetShadowEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) IsShadowEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) SetShadowSmooth(smooth float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) GetShadowSmooth() { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) SetShadowFilter(filter Light2DShadowFilter, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) GetShadowFilter() { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) SetShadowColor(shadow_color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) GetShadowColor() { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) SetBlendMode(mode Light2DBlendMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) GetBlendMode() { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) SetHeight(height float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light2D) GetHeight() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
