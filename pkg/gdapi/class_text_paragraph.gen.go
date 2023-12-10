// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextParagraph) SetDirection(direction TextServerDirection, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1418190634) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextParagraph) GetDirection() TextServerDirection {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2516697328) // FIXME: should cache?
  var ret TextServerDirection
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) SetCustomPunctuation(custom_punctuation String, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_punctuation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(custom_punctuation.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextParagraph) GetCustomPunctuation() String {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_punctuation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) SetOrientation(orientation TextServerOrientation, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_orientation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 42823726) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&orientation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextParagraph) GetOrientation() TextServerOrientation {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_orientation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 175768116) // FIXME: should cache?
  var ret TextServerOrientation
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) SetPreserveInvalid(enabled bool, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_preserve_invalid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextParagraph) GetPreserveInvalid() bool {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_preserve_invalid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) SetPreserveControl(enabled bool, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_preserve_control")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextParagraph) GetPreserveControl() bool {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_preserve_control")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) SetBidiOverride(override Array, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bidi_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(override.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextParagraph) SetDropcap(text String, font Font, font_size int, dropcap_margins Rect2, language String, ) bool {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_dropcap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2613124475) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(text.AsCTypePtr()), gdc.ConstTypePtr(font.AsCTypePtr()), gdc.ConstTypePtr(&font_size), gdc.ConstTypePtr(dropcap_margins.AsCTypePtr()), gdc.ConstTypePtr(language.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) ClearDropcap()  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_dropcap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextParagraph) AddString(text String, font Font, font_size int, language String, meta Variant, ) bool {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 867188035) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(text.AsCTypePtr()), gdc.ConstTypePtr(font.AsCTypePtr()), gdc.ConstTypePtr(&font_size), gdc.ConstTypePtr(language.AsCTypePtr()), gdc.ConstTypePtr(meta.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) AddObject(key Variant, size Vector2, inline_align InlineAlignment, length int, baseline float32, ) bool {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_object")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 735420116) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(key.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&inline_align), gdc.ConstTypePtr(&length), gdc.ConstTypePtr(&baseline), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) ResizeObject(key Variant, size Vector2, inline_align InlineAlignment, baseline float32, ) bool {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("resize_object")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 960819067) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(key.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&inline_align), gdc.ConstTypePtr(&baseline), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) SetAlignment(alignment HorizontalAlignment, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2312603777) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextParagraph) GetAlignment() HorizontalAlignment {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 341400642) // FIXME: should cache?
  var ret HorizontalAlignment
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) TabAlign(tab_stops PackedFloat32Array, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tab_align")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2899603908) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(tab_stops.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextParagraph) SetBreakFlags(flags TextServerLineBreakFlag, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_break_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2809697122) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextParagraph) GetBreakFlags() TextServerLineBreakFlag {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_break_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2340632602) // FIXME: should cache?
  var ret TextServerLineBreakFlag
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) SetJustificationFlags(flags TextServerJustificationFlag, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_justification_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2877345813) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextParagraph) GetJustificationFlags() TextServerJustificationFlag {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_justification_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1583363614) // FIXME: should cache?
  var ret TextServerJustificationFlag
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) SetTextOverrunBehavior(overrun_behavior TextServerOverrunBehavior, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text_overrun_behavior")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1008890932) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&overrun_behavior), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextParagraph) GetTextOverrunBehavior() TextServerOverrunBehavior {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text_overrun_behavior")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3779142101) // FIXME: should cache?
  var ret TextServerOverrunBehavior
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) SetWidth(width float32, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextParagraph) GetWidth() float32 {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) GetNonWrappedSize() Vector2 {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_non_wrapped_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) GetSize() Vector2 {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) GetRid() RID {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) GetLineRid(line int, ) RID {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 495598643) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) GetDropcapRid() RID {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_dropcap_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) GetLineCount() int {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) SetMaxLinesVisible(max_lines_visible int, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_lines_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_lines_visible), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextParagraph) GetMaxLinesVisible() int {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_lines_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) GetLineObjects(line int, ) Array {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_objects")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 663333327) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) GetLineObjectRect(line int, key Variant, ) Rect2 {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_object_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 204315017) // FIXME: should cache?
  var ret Rect2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(key.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) GetLineSize(line int, ) Vector2 {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2299179447) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) GetLineRange(line int, ) Vector2i {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 880721226) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) GetLineAscent(line int, ) float32 {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_ascent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) GetLineDescent(line int, ) float32 {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_descent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) GetLineWidth(line int, ) float32 {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) GetLineUnderlinePosition(line int, ) float32 {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_underline_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) GetLineUnderlineThickness(line int, ) float32 {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_underline_thickness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) GetDropcapSize() Vector2 {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_dropcap_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) GetDropcapLines() int {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_dropcap_lines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextParagraph) Draw(canvas RID, pos Vector2, color Color, dc_color Color, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 367324453) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(canvas.AsCTypePtr()), gdc.ConstTypePtr(pos.AsCTypePtr()), gdc.ConstTypePtr(color.AsCTypePtr()), gdc.ConstTypePtr(dc_color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextParagraph) DrawOutline(canvas RID, pos Vector2, outline_size int, color Color, dc_color Color, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2159523405) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(canvas.AsCTypePtr()), gdc.ConstTypePtr(pos.AsCTypePtr()), gdc.ConstTypePtr(&outline_size), gdc.ConstTypePtr(color.AsCTypePtr()), gdc.ConstTypePtr(dc_color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextParagraph) DrawLine(canvas RID, pos Vector2, line int, color Color, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3963848920) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(canvas.AsCTypePtr()), gdc.ConstTypePtr(pos.AsCTypePtr()), gdc.ConstTypePtr(&line), gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextParagraph) DrawLineOutline(canvas RID, pos Vector2, line int, outline_size int, color Color, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_line_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1814903311) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(canvas.AsCTypePtr()), gdc.ConstTypePtr(pos.AsCTypePtr()), gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&outline_size), gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextParagraph) DrawDropcap(canvas RID, pos Vector2, color Color, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_dropcap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1164457837) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(canvas.AsCTypePtr()), gdc.ConstTypePtr(pos.AsCTypePtr()), gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextParagraph) DrawDropcapOutline(canvas RID, pos Vector2, outline_size int, color Color, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_dropcap_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1364491366) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(canvas.AsCTypePtr()), gdc.ConstTypePtr(pos.AsCTypePtr()), gdc.ConstTypePtr(&outline_size), gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextParagraph) HitTest(coords Vector2, ) int {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("hit_test")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3820158470) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(coords.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *TextParagraph) GetPropDirection() int {
  panic("TODO: implement")
}

func (me *TextParagraph) SetPropDirection(value int) {
  panic("TODO: implement")
}

func (me *TextParagraph) GetPropCustomPunctuation() String {
  panic("TODO: implement")
}

func (me *TextParagraph) SetPropCustomPunctuation(value String) {
  panic("TODO: implement")
}

func (me *TextParagraph) GetPropOrientation() int {
  panic("TODO: implement")
}

func (me *TextParagraph) SetPropOrientation(value int) {
  panic("TODO: implement")
}

func (me *TextParagraph) GetPropPreserveInvalid() bool {
  panic("TODO: implement")
}

func (me *TextParagraph) SetPropPreserveInvalid(value bool) {
  panic("TODO: implement")
}

func (me *TextParagraph) GetPropPreserveControl() bool {
  panic("TODO: implement")
}

func (me *TextParagraph) SetPropPreserveControl(value bool) {
  panic("TODO: implement")
}

func (me *TextParagraph) GetPropAlignment() int {
  panic("TODO: implement")
}

func (me *TextParagraph) SetPropAlignment(value int) {
  panic("TODO: implement")
}

func (me *TextParagraph) GetPropBreakFlags() int {
  panic("TODO: implement")
}

func (me *TextParagraph) SetPropBreakFlags(value int) {
  panic("TODO: implement")
}

func (me *TextParagraph) GetPropJustificationFlags() int {
  panic("TODO: implement")
}

func (me *TextParagraph) SetPropJustificationFlags(value int) {
  panic("TODO: implement")
}

func (me *TextParagraph) GetPropTextOverrunBehavior() int {
  panic("TODO: implement")
}

func (me *TextParagraph) SetPropTextOverrunBehavior(value int) {
  panic("TODO: implement")
}

func (me *TextParagraph) GetPropWidth() float32 {
  panic("TODO: implement")
}

func (me *TextParagraph) SetPropWidth(value float32) {
  panic("TODO: implement")
}

func (me *TextParagraph) GetPropMaxLinesVisible() int {
  panic("TODO: implement")
}

func (me *TextParagraph) SetPropMaxLinesVisible(value int) {
  panic("TODO: implement")
}