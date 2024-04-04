// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type TextParagraph struct {
  RefCounted
}

func (me *TextParagraph) BaseClass() string {
  return "TextParagraph"
}

func NewTextParagraph() *TextParagraph {
  str := StringNameFromStr("TextParagraph") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &TextParagraph{}
  obj.SetBaseObject(objPtr)
  return obj
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
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextParagraph) SetDirection(direction TextServerDirection, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1418190634) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextParagraph) GetDirection() TextServerDirection {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2516697328) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret TextServerDirection

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TextParagraph) SetCustomPunctuation(custom_punctuation String, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_punctuation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{custom_punctuation.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextParagraph) GetCustomPunctuation() String {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_punctuation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextParagraph) SetOrientation(orientation TextServerOrientation, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_orientation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 42823726) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&orientation) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextParagraph) GetOrientation() TextServerOrientation {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_orientation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 175768116) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret TextServerOrientation

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TextParagraph) SetPreserveInvalid(enabled bool, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_preserve_invalid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextParagraph) GetPreserveInvalid() bool {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_preserve_invalid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextParagraph) SetPreserveControl(enabled bool, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_preserve_control")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextParagraph) GetPreserveControl() bool {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_preserve_control")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextParagraph) SetBidiOverride(override Array, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bidi_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{override.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextParagraph) SetDropcap(text String, font Font, font_size int64, dropcap_margins Rect2, language String, ) bool {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_dropcap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2498990330) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), font.AsCTypePtr(), gdc.ConstTypePtr(&font_size) , dropcap_margins.AsCTypePtr(), language.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&font_size)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextParagraph) ClearDropcap()  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_dropcap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextParagraph) AddString(text String, font Font, font_size int64, language String, meta Variant, ) bool {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 621426851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), font.AsCTypePtr(), gdc.ConstTypePtr(&font_size) , language.AsCTypePtr(), meta.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&font_size)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextParagraph) AddObject(key Variant, size Vector2, inline_align InlineAlignment, length int64, baseline float64, ) bool {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_object")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1316529304) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{key.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&inline_align) , gdc.ConstTypePtr(&length) , gdc.ConstTypePtr(&baseline) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&inline_align)
  pinner.Pin(&length)
  pinner.Pin(&baseline)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextParagraph) ResizeObject(key Variant, size Vector2, inline_align InlineAlignment, baseline float64, ) bool {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("resize_object")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2095776372) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{key.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&inline_align) , gdc.ConstTypePtr(&baseline) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&inline_align)
  pinner.Pin(&baseline)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextParagraph) SetAlignment(alignment HorizontalAlignment, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2312603777) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextParagraph) GetAlignment() HorizontalAlignment {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 341400642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret HorizontalAlignment

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TextParagraph) TabAlign(tab_stops PackedFloat32Array, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tab_align")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2899603908) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{tab_stops.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextParagraph) SetBreakFlags(flags TextServerLineBreakFlag, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_break_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2809697122) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextParagraph) GetBreakFlags() TextServerLineBreakFlag {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_break_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2340632602) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret TextServerLineBreakFlag

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TextParagraph) SetJustificationFlags(flags TextServerJustificationFlag, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_justification_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2877345813) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextParagraph) GetJustificationFlags() TextServerJustificationFlag {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_justification_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1583363614) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret TextServerJustificationFlag

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TextParagraph) SetTextOverrunBehavior(overrun_behavior TextServerOverrunBehavior, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text_overrun_behavior")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1008890932) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&overrun_behavior) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextParagraph) GetTextOverrunBehavior() TextServerOverrunBehavior {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text_overrun_behavior")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3779142101) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret TextServerOverrunBehavior

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TextParagraph) SetWidth(width float64, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextParagraph) GetWidth() float64 {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextParagraph) GetNonWrappedSize() Vector2 {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_non_wrapped_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextParagraph) GetSize() Vector2 {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextParagraph) GetRid() RID {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextParagraph) GetLineRid(line int64, ) RID {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 495598643) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()
  pinner.Pin(&line)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextParagraph) GetDropcapRid() RID {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_dropcap_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextParagraph) GetLineCount() int64 {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextParagraph) SetMaxLinesVisible(max_lines_visible int64, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_lines_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_lines_visible) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextParagraph) GetMaxLinesVisible() int64 {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_lines_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextParagraph) GetLineObjects(line int64, ) Array {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_objects")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 663333327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  pinner.Pin(&line)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextParagraph) GetLineObjectRect(line int64, key Variant, ) Rect2 {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_object_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 204315017) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , key.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRect2()
  pinner.Pin(&line)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextParagraph) GetLineSize(line int64, ) Vector2 {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2299179447) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()
  pinner.Pin(&line)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextParagraph) GetLineRange(line int64, ) Vector2i {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 880721226) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()
  pinner.Pin(&line)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextParagraph) GetLineAscent(line int64, ) float64 {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_ascent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&line)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextParagraph) GetLineDescent(line int64, ) float64 {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_descent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&line)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextParagraph) GetLineWidth(line int64, ) float64 {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&line)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextParagraph) GetLineUnderlinePosition(line int64, ) float64 {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_underline_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&line)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextParagraph) GetLineUnderlineThickness(line int64, ) float64 {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_underline_thickness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&line)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextParagraph) GetDropcapSize() Vector2 {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_dropcap_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextParagraph) GetDropcapLines() int64 {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_dropcap_lines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextParagraph) Draw(canvas RID, pos Vector2, color Color, dc_color Color, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1567802413) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{canvas.AsCTypePtr(), pos.AsCTypePtr(), color.AsCTypePtr(), dc_color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextParagraph) DrawOutline(canvas RID, pos Vector2, outline_size int64, color Color, dc_color Color, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1893131224) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{canvas.AsCTypePtr(), pos.AsCTypePtr(), gdc.ConstTypePtr(&outline_size) , color.AsCTypePtr(), dc_color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextParagraph) DrawLine(canvas RID, pos Vector2, line int64, color Color, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1242169894) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{canvas.AsCTypePtr(), pos.AsCTypePtr(), gdc.ConstTypePtr(&line) , color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextParagraph) DrawLineOutline(canvas RID, pos Vector2, line int64, outline_size int64, color Color, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_line_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2664926980) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{canvas.AsCTypePtr(), pos.AsCTypePtr(), gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&outline_size) , color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextParagraph) DrawDropcap(canvas RID, pos Vector2, color Color, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_dropcap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 856975658) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{canvas.AsCTypePtr(), pos.AsCTypePtr(), color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextParagraph) DrawDropcapOutline(canvas RID, pos Vector2, outline_size int64, color Color, )  {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_dropcap_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1343401456) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{canvas.AsCTypePtr(), pos.AsCTypePtr(), gdc.ConstTypePtr(&outline_size) , color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextParagraph) HitTest(coords Vector2, ) int64 {
  classNameV := StringNameFromStr("TextParagraph")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("hit_test")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3820158470) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
