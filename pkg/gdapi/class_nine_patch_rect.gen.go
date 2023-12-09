// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NinePatchRect struct {
  obj gdc.ObjectPtr
}

func (me *NinePatchRect) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NinePatchRect) BaseClass() string {
  return "NinePatchRect"
}



// Enums

type NinePatchRectAxisStretchMode int
const (
  NinePatchRectAxisStretchModeAxisStretchModeStretch NinePatchRectAxisStretchMode = 0
  NinePatchRectAxisStretchModeAxisStretchModeTile NinePatchRectAxisStretchMode = 1
  NinePatchRectAxisStretchModeAxisStretchModeTileFit NinePatchRectAxisStretchMode = 2
)

func (me *NinePatchRect) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NinePatchRect) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *NinePatchRect) SetTexture(texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *NinePatchRect) GetTexture()  {
  panic("TODO: implement")
}

func  (me *NinePatchRect) SetPatchMargin(margin Side, value int, )  {
  panic("TODO: implement")
}

func  (me *NinePatchRect) GetPatchMargin(margin Side, )  {
  panic("TODO: implement")
}

func  (me *NinePatchRect) SetRegionRect(rect Rect2, )  {
  panic("TODO: implement")
}

func  (me *NinePatchRect) GetRegionRect()  {
  panic("TODO: implement")
}

func  (me *NinePatchRect) SetDrawCenter(draw_center bool, )  {
  panic("TODO: implement")
}

func  (me *NinePatchRect) IsDrawCenterEnabled()  {
  panic("TODO: implement")
}

func  (me *NinePatchRect) SetHAxisStretchMode(mode NinePatchRectAxisStretchMode, )  {
  panic("TODO: implement")
}

func  (me *NinePatchRect) GetHAxisStretchMode()  {
  panic("TODO: implement")
}

func  (me *NinePatchRect) SetVAxisStretchMode(mode NinePatchRectAxisStretchMode, )  {
  panic("TODO: implement")
}

func  (me *NinePatchRect) GetVAxisStretchMode()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
