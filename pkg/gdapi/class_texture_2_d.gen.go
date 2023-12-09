// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Texture2D struct {
  obj gdc.ObjectPtr
}

func (me *Texture2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Texture2D) BaseClass() string {
  return "Texture2D"
}



// Enums

func (me *Texture2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Texture2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Texture2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Texture2D) XGetWidth()  {
  panic("TODO: implement")
}

func  (me *Texture2D) XGetHeight()  {
  panic("TODO: implement")
}

func  (me *Texture2D) XIsPixelOpaque(x int, y int, )  {
  panic("TODO: implement")
}

func  (me *Texture2D) XHasAlpha()  {
  panic("TODO: implement")
}

func  (me *Texture2D) XDraw(to_canvas_item RID, pos Vector2, modulate Color, transpose bool, )  {
  panic("TODO: implement")
}

func  (me *Texture2D) XDrawRect(to_canvas_item RID, rect Rect2, tile bool, modulate Color, transpose bool, )  {
  panic("TODO: implement")
}

func  (me *Texture2D) XDrawRectRegion(to_canvas_item RID, rect Rect2, src_rect Rect2, modulate Color, transpose bool, clip_uv bool, )  {
  panic("TODO: implement")
}

func  (me *Texture2D) GetWidth()  {
  panic("TODO: implement")
}

func  (me *Texture2D) GetHeight()  {
  panic("TODO: implement")
}

func  (me *Texture2D) GetSize()  {
  panic("TODO: implement")
}

func  (me *Texture2D) HasAlpha()  {
  panic("TODO: implement")
}

func  (me *Texture2D) Draw(canvas_item RID, position Vector2, modulate Color, transpose bool, )  {
  panic("TODO: implement")
}

func  (me *Texture2D) DrawRect(canvas_item RID, rect Rect2, tile bool, modulate Color, transpose bool, )  {
  panic("TODO: implement")
}

func  (me *Texture2D) DrawRectRegion(canvas_item RID, rect Rect2, src_rect Rect2, modulate Color, transpose bool, clip_uv bool, )  {
  panic("TODO: implement")
}

func  (me *Texture2D) GetImage()  {
  panic("TODO: implement")
}

func  (me *Texture2D) CreatePlaceholder()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
