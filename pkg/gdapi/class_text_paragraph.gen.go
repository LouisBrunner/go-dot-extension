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

func  (me *TextParagraph) Clear() { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) SetDirection(direction TextServerDirection, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetDirection() { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) SetCustomPunctuation(custom_punctuation String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetCustomPunctuation() { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) SetOrientation(orientation TextServerOrientation, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetOrientation() { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) SetPreserveInvalid(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetPreserveInvalid() { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) SetPreserveControl(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetPreserveControl() { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) SetBidiOverride(override Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) SetDropcap(text String, font Font, font_size int, dropcap_margins Rect2, language String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) ClearDropcap() { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) AddString(text String, font Font, font_size int, language String, meta Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) AddObject(key Variant, size Vector2, inline_align InlineAlignment, length int, baseline float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) ResizeObject(key Variant, size Vector2, inline_align InlineAlignment, baseline float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) SetAlignment(alignment HorizontalAlignment, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetAlignment() { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) TabAlign(tab_stops PackedFloat32Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) SetBreakFlags(flags TextServerLineBreakFlag, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetBreakFlags() { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) SetJustificationFlags(flags TextServerJustificationFlag, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetJustificationFlags() { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) SetTextOverrunBehavior(overrun_behavior TextServerOverrunBehavior, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetTextOverrunBehavior() { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) SetWidth(width float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetWidth() { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetNonWrappedSize() { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetSize() { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetRid() { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetLineRid(line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetDropcapRid() { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetLineCount() { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) SetMaxLinesVisible(max_lines_visible int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetMaxLinesVisible() { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetLineObjects(line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetLineObjectRect(line int, key Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetLineSize(line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetLineRange(line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetLineAscent(line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetLineDescent(line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetLineWidth(line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetLineUnderlinePosition(line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetLineUnderlineThickness(line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetDropcapSize() { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) GetDropcapLines() { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) Draw(canvas RID, pos Vector2, color Color, dc_color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) DrawOutline(canvas RID, pos Vector2, outline_size int, color Color, dc_color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) DrawLine(canvas RID, pos Vector2, line int, color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) DrawLineOutline(canvas RID, pos Vector2, line int, outline_size int, color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) DrawDropcap(canvas RID, pos Vector2, color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) DrawDropcapOutline(canvas RID, pos Vector2, outline_size int, color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextParagraph) HitTest(coords Vector2, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
