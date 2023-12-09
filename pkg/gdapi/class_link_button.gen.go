// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type LinkButton struct {
  obj gdc.ObjectPtr
}

func (me *LinkButton) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *LinkButton) BaseClass() string {
  return "LinkButton"
}



// Enums

type LinkButtonUnderlineMode int
const (
  LinkButtonUnderlineModeUnderlineModeAlways LinkButtonUnderlineMode = 0
  LinkButtonUnderlineModeUnderlineModeOnHover LinkButtonUnderlineMode = 1
  LinkButtonUnderlineModeUnderlineModeNever LinkButtonUnderlineMode = 2
)

func (me *LinkButton) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *LinkButton) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *LinkButton) SetText(text String, )  {
  panic("TODO: implement")
}

func  (me *LinkButton) GetText()  {
  panic("TODO: implement")
}

func  (me *LinkButton) SetTextDirection(direction ControlTextDirection, )  {
  panic("TODO: implement")
}

func  (me *LinkButton) GetTextDirection()  {
  panic("TODO: implement")
}

func  (me *LinkButton) SetLanguage(language String, )  {
  panic("TODO: implement")
}

func  (me *LinkButton) GetLanguage()  {
  panic("TODO: implement")
}

func  (me *LinkButton) SetUri(uri String, )  {
  panic("TODO: implement")
}

func  (me *LinkButton) GetUri()  {
  panic("TODO: implement")
}

func  (me *LinkButton) SetUnderlineMode(underline_mode LinkButtonUnderlineMode, )  {
  panic("TODO: implement")
}

func  (me *LinkButton) GetUnderlineMode()  {
  panic("TODO: implement")
}

func  (me *LinkButton) SetStructuredTextBidiOverride(parser TextServerStructuredTextParser, )  {
  panic("TODO: implement")
}

func  (me *LinkButton) GetStructuredTextBidiOverride()  {
  panic("TODO: implement")
}

func  (me *LinkButton) SetStructuredTextBidiOverrideOptions(args Array, )  {
  panic("TODO: implement")
}

func  (me *LinkButton) GetStructuredTextBidiOverrideOptions()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
