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

type LinkButton struct {
  BaseButton
}

func (me *LinkButton) BaseClass() string {
  return "LinkButton"
}

func NewLinkButton() *LinkButton {
  str := StringNameFromStr("LinkButton") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &LinkButton{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type LinkButtonUnderlineMode int
const (
  LinkButtonUnderlineModeUnderlineModeAlways LinkButtonUnderlineMode = 0
  LinkButtonUnderlineModeUnderlineModeOnHover LinkButtonUnderlineMode = 1
  LinkButtonUnderlineModeUnderlineModeNever LinkButtonUnderlineMode = 2
)

func (me *LinkButton) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *LinkButton) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *LinkButton) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *LinkButton) SetText(text String, )  {
  classNameV := StringNameFromStr("LinkButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LinkButton) GetText() String {
  classNameV := StringNameFromStr("LinkButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *LinkButton) SetTextDirection(direction ControlTextDirection, )  {
  classNameV := StringNameFromStr("LinkButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 119160795) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LinkButton) GetTextDirection() ControlTextDirection {
  classNameV := StringNameFromStr("LinkButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 797257663) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ControlTextDirection

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *LinkButton) SetLanguage(language String, )  {
  classNameV := StringNameFromStr("LinkButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{language.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LinkButton) GetLanguage() String {
  classNameV := StringNameFromStr("LinkButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *LinkButton) SetUri(uri String, )  {
  classNameV := StringNameFromStr("LinkButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_uri")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{uri.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LinkButton) GetUri() String {
  classNameV := StringNameFromStr("LinkButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_uri")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *LinkButton) SetUnderlineMode(underline_mode LinkButtonUnderlineMode, )  {
  classNameV := StringNameFromStr("LinkButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_underline_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4032947085) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&underline_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LinkButton) GetUnderlineMode() LinkButtonUnderlineMode {
  classNameV := StringNameFromStr("LinkButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_underline_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 568343738) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret LinkButtonUnderlineMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *LinkButton) SetStructuredTextBidiOverride(parser TextServerStructuredTextParser, )  {
  classNameV := StringNameFromStr("LinkButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_structured_text_bidi_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 55961453) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&parser) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LinkButton) GetStructuredTextBidiOverride() TextServerStructuredTextParser {
  classNameV := StringNameFromStr("LinkButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_structured_text_bidi_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3385126229) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret TextServerStructuredTextParser

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *LinkButton) SetStructuredTextBidiOverrideOptions(args Array, )  {
  classNameV := StringNameFromStr("LinkButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_structured_text_bidi_override_options")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{args.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LinkButton) GetStructuredTextBidiOverrideOptions() Array {
  classNameV := StringNameFromStr("LinkButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_structured_text_bidi_override_options")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
