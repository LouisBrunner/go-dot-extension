// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TextParagraph struct {
  obj gdc.ObjectPtr
}

func (me *TextParagraph) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TextParagraph) BaseClass() string {
  return "TextParagraph"
}



// Enums

func (me *TextParagraph) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TextParagraph) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TextParagraph) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *TextParagraph) Clear()  {
  panic("TODO: implement")
}

func  (me *TextParagraph) SetDirection(direction TextServerDirection, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetDirection()  {
  panic("TODO: implement")
}

func  (me *TextParagraph) SetCustomPunctuation(custom_punctuation String, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetCustomPunctuation()  {
  panic("TODO: implement")
}

func  (me *TextParagraph) SetOrientation(orientation TextServerOrientation, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetOrientation()  {
  panic("TODO: implement")
}

func  (me *TextParagraph) SetPreserveInvalid(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetPreserveInvalid()  {
  panic("TODO: implement")
}

func  (me *TextParagraph) SetPreserveControl(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetPreserveControl()  {
  panic("TODO: implement")
}

func  (me *TextParagraph) SetBidiOverride(override Array, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) SetDropcap(text String, font Font, font_size int, dropcap_margins Rect2, language String, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) ClearDropcap()  {
  panic("TODO: implement")
}

func  (me *TextParagraph) AddString(text String, font Font, font_size int, language String, meta Variant, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) AddObject(key Variant, size Vector2, inline_align InlineAlignment, length int, baseline float32, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) ResizeObject(key Variant, size Vector2, inline_align InlineAlignment, baseline float32, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) SetAlignment(alignment HorizontalAlignment, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetAlignment()  {
  panic("TODO: implement")
}

func  (me *TextParagraph) TabAlign(tab_stops PackedFloat32Array, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) SetBreakFlags(flags TextServerLineBreakFlag, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetBreakFlags()  {
  panic("TODO: implement")
}

func  (me *TextParagraph) SetJustificationFlags(flags TextServerJustificationFlag, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetJustificationFlags()  {
  panic("TODO: implement")
}

func  (me *TextParagraph) SetTextOverrunBehavior(overrun_behavior TextServerOverrunBehavior, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetTextOverrunBehavior()  {
  panic("TODO: implement")
}

func  (me *TextParagraph) SetWidth(width float32, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetWidth()  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetNonWrappedSize()  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetSize()  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetRid()  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetLineRid(line int, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetDropcapRid()  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetLineCount()  {
  panic("TODO: implement")
}

func  (me *TextParagraph) SetMaxLinesVisible(max_lines_visible int, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetMaxLinesVisible()  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetLineObjects(line int, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetLineObjectRect(line int, key Variant, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetLineSize(line int, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetLineRange(line int, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetLineAscent(line int, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetLineDescent(line int, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetLineWidth(line int, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetLineUnderlinePosition(line int, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetLineUnderlineThickness(line int, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetDropcapSize()  {
  panic("TODO: implement")
}

func  (me *TextParagraph) GetDropcapLines()  {
  panic("TODO: implement")
}

func  (me *TextParagraph) Draw(canvas RID, pos Vector2, color Color, dc_color Color, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) DrawOutline(canvas RID, pos Vector2, outline_size int, color Color, dc_color Color, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) DrawLine(canvas RID, pos Vector2, line int, color Color, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) DrawLineOutline(canvas RID, pos Vector2, line int, outline_size int, color Color, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) DrawDropcap(canvas RID, pos Vector2, color Color, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) DrawDropcapOutline(canvas RID, pos Vector2, outline_size int, color Color, )  {
  panic("TODO: implement")
}

func  (me *TextParagraph) HitTest(coords Vector2, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
