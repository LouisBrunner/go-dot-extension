// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TextLine struct {
  obj gdc.ObjectPtr
}

func (me *TextLine) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TextLine) BaseClass() string {
  return "TextLine"
}



// Enums

func (me *TextLine) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TextLine) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TextLine) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *TextLine) Clear()  {
  panic("TODO: implement")
}

func  (me *TextLine) SetDirection(direction TextServerDirection, )  {
  panic("TODO: implement")
}

func  (me *TextLine) GetDirection()  {
  panic("TODO: implement")
}

func  (me *TextLine) SetOrientation(orientation TextServerOrientation, )  {
  panic("TODO: implement")
}

func  (me *TextLine) GetOrientation()  {
  panic("TODO: implement")
}

func  (me *TextLine) SetPreserveInvalid(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TextLine) GetPreserveInvalid()  {
  panic("TODO: implement")
}

func  (me *TextLine) SetPreserveControl(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TextLine) GetPreserveControl()  {
  panic("TODO: implement")
}

func  (me *TextLine) SetBidiOverride(override Array, )  {
  panic("TODO: implement")
}

func  (me *TextLine) AddString(text String, font Font, font_size int, language String, meta Variant, )  {
  panic("TODO: implement")
}

func  (me *TextLine) AddObject(key Variant, size Vector2, inline_align InlineAlignment, length int, baseline float32, )  {
  panic("TODO: implement")
}

func  (me *TextLine) ResizeObject(key Variant, size Vector2, inline_align InlineAlignment, baseline float32, )  {
  panic("TODO: implement")
}

func  (me *TextLine) SetWidth(width float32, )  {
  panic("TODO: implement")
}

func  (me *TextLine) GetWidth()  {
  panic("TODO: implement")
}

func  (me *TextLine) SetHorizontalAlignment(alignment HorizontalAlignment, )  {
  panic("TODO: implement")
}

func  (me *TextLine) GetHorizontalAlignment()  {
  panic("TODO: implement")
}

func  (me *TextLine) TabAlign(tab_stops PackedFloat32Array, )  {
  panic("TODO: implement")
}

func  (me *TextLine) SetFlags(flags TextServerJustificationFlag, )  {
  panic("TODO: implement")
}

func  (me *TextLine) GetFlags()  {
  panic("TODO: implement")
}

func  (me *TextLine) SetTextOverrunBehavior(overrun_behavior TextServerOverrunBehavior, )  {
  panic("TODO: implement")
}

func  (me *TextLine) GetTextOverrunBehavior()  {
  panic("TODO: implement")
}

func  (me *TextLine) GetObjects()  {
  panic("TODO: implement")
}

func  (me *TextLine) GetObjectRect(key Variant, )  {
  panic("TODO: implement")
}

func  (me *TextLine) GetSize()  {
  panic("TODO: implement")
}

func  (me *TextLine) GetRid()  {
  panic("TODO: implement")
}

func  (me *TextLine) GetLineAscent()  {
  panic("TODO: implement")
}

func  (me *TextLine) GetLineDescent()  {
  panic("TODO: implement")
}

func  (me *TextLine) GetLineWidth()  {
  panic("TODO: implement")
}

func  (me *TextLine) GetLineUnderlinePosition()  {
  panic("TODO: implement")
}

func  (me *TextLine) GetLineUnderlineThickness()  {
  panic("TODO: implement")
}

func  (me *TextLine) Draw(canvas RID, pos Vector2, color Color, )  {
  panic("TODO: implement")
}

func  (me *TextLine) DrawOutline(canvas RID, pos Vector2, outline_size int, color Color, )  {
  panic("TODO: implement")
}

func  (me *TextLine) HitTest(coords float32, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
