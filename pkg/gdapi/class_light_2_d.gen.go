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



// Enums

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

func (me *Light2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Light2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Light2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Light2D) SetEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Light2D) IsEnabled()  {
  panic("TODO: implement")
}

func  (me *Light2D) SetEditorOnly(editor_only bool, )  {
  panic("TODO: implement")
}

func  (me *Light2D) IsEditorOnly()  {
  panic("TODO: implement")
}

func  (me *Light2D) SetColor(color Color, )  {
  panic("TODO: implement")
}

func  (me *Light2D) GetColor()  {
  panic("TODO: implement")
}

func  (me *Light2D) SetEnergy(energy float32, )  {
  panic("TODO: implement")
}

func  (me *Light2D) GetEnergy()  {
  panic("TODO: implement")
}

func  (me *Light2D) SetZRangeMin(z int, )  {
  panic("TODO: implement")
}

func  (me *Light2D) GetZRangeMin()  {
  panic("TODO: implement")
}

func  (me *Light2D) SetZRangeMax(z int, )  {
  panic("TODO: implement")
}

func  (me *Light2D) GetZRangeMax()  {
  panic("TODO: implement")
}

func  (me *Light2D) SetLayerRangeMin(layer int, )  {
  panic("TODO: implement")
}

func  (me *Light2D) GetLayerRangeMin()  {
  panic("TODO: implement")
}

func  (me *Light2D) SetLayerRangeMax(layer int, )  {
  panic("TODO: implement")
}

func  (me *Light2D) GetLayerRangeMax()  {
  panic("TODO: implement")
}

func  (me *Light2D) SetItemCullMask(item_cull_mask int, )  {
  panic("TODO: implement")
}

func  (me *Light2D) GetItemCullMask()  {
  panic("TODO: implement")
}

func  (me *Light2D) SetItemShadowCullMask(item_shadow_cull_mask int, )  {
  panic("TODO: implement")
}

func  (me *Light2D) GetItemShadowCullMask()  {
  panic("TODO: implement")
}

func  (me *Light2D) SetShadowEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Light2D) IsShadowEnabled()  {
  panic("TODO: implement")
}

func  (me *Light2D) SetShadowSmooth(smooth float32, )  {
  panic("TODO: implement")
}

func  (me *Light2D) GetShadowSmooth()  {
  panic("TODO: implement")
}

func  (me *Light2D) SetShadowFilter(filter Light2DShadowFilter, )  {
  panic("TODO: implement")
}

func  (me *Light2D) GetShadowFilter()  {
  panic("TODO: implement")
}

func  (me *Light2D) SetShadowColor(shadow_color Color, )  {
  panic("TODO: implement")
}

func  (me *Light2D) GetShadowColor()  {
  panic("TODO: implement")
}

func  (me *Light2D) SetBlendMode(mode Light2DBlendMode, )  {
  panic("TODO: implement")
}

func  (me *Light2D) GetBlendMode()  {
  panic("TODO: implement")
}

func  (me *Light2D) SetHeight(height float32, )  {
  panic("TODO: implement")
}

func  (me *Light2D) GetHeight()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
