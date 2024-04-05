// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForLinkButtonList struct {
  fnSetText gdc.MethodBindPtr
  fnGetText gdc.MethodBindPtr
  fnSetTextDirection gdc.MethodBindPtr
  fnGetTextDirection gdc.MethodBindPtr
  fnSetLanguage gdc.MethodBindPtr
  fnGetLanguage gdc.MethodBindPtr
  fnSetUri gdc.MethodBindPtr
  fnGetUri gdc.MethodBindPtr
  fnSetUnderlineMode gdc.MethodBindPtr
  fnGetUnderlineMode gdc.MethodBindPtr
  fnSetStructuredTextBidiOverride gdc.MethodBindPtr
  fnGetStructuredTextBidiOverride gdc.MethodBindPtr
  fnSetStructuredTextBidiOverrideOptions gdc.MethodBindPtr
  fnGetStructuredTextBidiOverrideOptions gdc.MethodBindPtr
}

var ptrsForLinkButton ptrsForLinkButtonList

func initLinkButtonPtrs(iface gdc.Interface) {

  className := StringNameFromStr("LinkButton")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_text")
    defer methodName.Destroy()
    ptrsForLinkButton.fnSetText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_text")
    defer methodName.Destroy()
    ptrsForLinkButton.fnGetText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("set_text_direction")
    defer methodName.Destroy()
    ptrsForLinkButton.fnSetTextDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 119160795))
  }
  {
    methodName := StringNameFromStr("get_text_direction")
    defer methodName.Destroy()
    ptrsForLinkButton.fnGetTextDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 797257663))
  }
  {
    methodName := StringNameFromStr("set_language")
    defer methodName.Destroy()
    ptrsForLinkButton.fnSetLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_language")
    defer methodName.Destroy()
    ptrsForLinkButton.fnGetLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("set_uri")
    defer methodName.Destroy()
    ptrsForLinkButton.fnSetUri = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_uri")
    defer methodName.Destroy()
    ptrsForLinkButton.fnGetUri = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("set_underline_mode")
    defer methodName.Destroy()
    ptrsForLinkButton.fnSetUnderlineMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4032947085))
  }
  {
    methodName := StringNameFromStr("get_underline_mode")
    defer methodName.Destroy()
    ptrsForLinkButton.fnGetUnderlineMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 568343738))
  }
  {
    methodName := StringNameFromStr("set_structured_text_bidi_override")
    defer methodName.Destroy()
    ptrsForLinkButton.fnSetStructuredTextBidiOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 55961453))
  }
  {
    methodName := StringNameFromStr("get_structured_text_bidi_override")
    defer methodName.Destroy()
    ptrsForLinkButton.fnGetStructuredTextBidiOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3385126229))
  }
  {
    methodName := StringNameFromStr("set_structured_text_bidi_override_options")
    defer methodName.Destroy()
    ptrsForLinkButton.fnSetStructuredTextBidiOverrideOptions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("get_structured_text_bidi_override_options")
    defer methodName.Destroy()
    ptrsForLinkButton.fnGetStructuredTextBidiOverrideOptions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
}

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
  cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLinkButton.fnSetText), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LinkButton) GetText() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLinkButton.fnGetText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *LinkButton) SetTextDirection(direction ControlTextDirection, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLinkButton.fnSetTextDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LinkButton) GetTextDirection() ControlTextDirection {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ControlTextDirection

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLinkButton.fnGetTextDirection), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *LinkButton) SetLanguage(language String, )  {
  cargs := []gdc.ConstTypePtr{language.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLinkButton.fnSetLanguage), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LinkButton) GetLanguage() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLinkButton.fnGetLanguage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *LinkButton) SetUri(uri String, )  {
  cargs := []gdc.ConstTypePtr{uri.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLinkButton.fnSetUri), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LinkButton) GetUri() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLinkButton.fnGetUri), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *LinkButton) SetUnderlineMode(underline_mode LinkButtonUnderlineMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&underline_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLinkButton.fnSetUnderlineMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LinkButton) GetUnderlineMode() LinkButtonUnderlineMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret LinkButtonUnderlineMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLinkButton.fnGetUnderlineMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *LinkButton) SetStructuredTextBidiOverride(parser TextServerStructuredTextParser, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&parser) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLinkButton.fnSetStructuredTextBidiOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LinkButton) GetStructuredTextBidiOverride() TextServerStructuredTextParser {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret TextServerStructuredTextParser

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLinkButton.fnGetStructuredTextBidiOverride), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *LinkButton) SetStructuredTextBidiOverrideOptions(args Array, )  {
  cargs := []gdc.ConstTypePtr{args.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLinkButton.fnSetStructuredTextBidiOverrideOptions), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LinkButton) GetStructuredTextBidiOverrideOptions() Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLinkButton.fnGetStructuredTextBidiOverrideOptions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
