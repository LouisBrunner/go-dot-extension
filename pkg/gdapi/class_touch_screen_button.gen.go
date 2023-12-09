// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TouchScreenButton struct {
  obj gdc.ObjectPtr
}

func (me *TouchScreenButton) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TouchScreenButton) BaseClass() string {
  return "TouchScreenButton"
}



// Enums

type TouchScreenButtonVisibilityMode int
const (
  TouchScreenButtonVisibilityModeVisibilityAlways TouchScreenButtonVisibilityMode = 0
  TouchScreenButtonVisibilityModeVisibilityTouchscreenOnly TouchScreenButtonVisibilityMode = 1
)

func (me *TouchScreenButton) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TouchScreenButton) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TouchScreenButton) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *TouchScreenButton) SetTextureNormal(texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *TouchScreenButton) GetTextureNormal()  {
  panic("TODO: implement")
}

func  (me *TouchScreenButton) SetTexturePressed(texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *TouchScreenButton) GetTexturePressed()  {
  panic("TODO: implement")
}

func  (me *TouchScreenButton) SetBitmask(bitmask BitMap, )  {
  panic("TODO: implement")
}

func  (me *TouchScreenButton) GetBitmask()  {
  panic("TODO: implement")
}

func  (me *TouchScreenButton) SetShape(shape Shape2D, )  {
  panic("TODO: implement")
}

func  (me *TouchScreenButton) GetShape()  {
  panic("TODO: implement")
}

func  (me *TouchScreenButton) SetShapeCentered(bool bool, )  {
  panic("TODO: implement")
}

func  (me *TouchScreenButton) IsShapeCentered()  {
  panic("TODO: implement")
}

func  (me *TouchScreenButton) SetShapeVisible(bool bool, )  {
  panic("TODO: implement")
}

func  (me *TouchScreenButton) IsShapeVisible()  {
  panic("TODO: implement")
}

func  (me *TouchScreenButton) SetAction(action String, )  {
  panic("TODO: implement")
}

func  (me *TouchScreenButton) GetAction()  {
  panic("TODO: implement")
}

func  (me *TouchScreenButton) SetVisibilityMode(mode TouchScreenButtonVisibilityMode, )  {
  panic("TODO: implement")
}

func  (me *TouchScreenButton) GetVisibilityMode()  {
  panic("TODO: implement")
}

func  (me *TouchScreenButton) SetPassbyPress(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TouchScreenButton) IsPassbyPressEnabled()  {
  panic("TODO: implement")
}

func  (me *TouchScreenButton) IsPressed()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
