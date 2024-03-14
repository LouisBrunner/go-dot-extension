// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ThemeDB struct {
  Object
}

func (me *ThemeDB) BaseClass() string {
  return "ThemeDB"
}



// Enums

func (me *ThemeDB) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ThemeDB) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ThemeDB) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ThemeDB) GetDefaultTheme() Theme {
  classNameV := StringNameFromStr("ThemeDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_default_theme")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 754276358) // FIXME: should cache?
  var ret Theme
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ThemeDB) GetProjectTheme() Theme {
  classNameV := StringNameFromStr("ThemeDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_project_theme")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 754276358) // FIXME: should cache?
  var ret Theme
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ThemeDB) SetFallbackBaseScale(base_scale float32, )  {
  classNameV := StringNameFromStr("ThemeDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fallback_base_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&base_scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ThemeDB) GetFallbackBaseScale() float32 {
  classNameV := StringNameFromStr("ThemeDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fallback_base_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ThemeDB) SetFallbackFont(font Font, )  {
  classNameV := StringNameFromStr("ThemeDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fallback_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1262170328) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ThemeDB) GetFallbackFont() Font {
  classNameV := StringNameFromStr("ThemeDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fallback_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3656929885) // FIXME: should cache?
  var ret Font
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ThemeDB) SetFallbackFontSize(font_size int, )  {
  classNameV := StringNameFromStr("ThemeDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fallback_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&font_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ThemeDB) GetFallbackFontSize() int {
  classNameV := StringNameFromStr("ThemeDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fallback_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ThemeDB) SetFallbackIcon(icon Texture2D, )  {
  classNameV := StringNameFromStr("ThemeDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fallback_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(icon.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ThemeDB) GetFallbackIcon() Texture2D {
  classNameV := StringNameFromStr("ThemeDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fallback_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 255860311) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ThemeDB) SetFallbackStylebox(stylebox StyleBox, )  {
  classNameV := StringNameFromStr("ThemeDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fallback_stylebox")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2797200388) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(stylebox.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ThemeDB) GetFallbackStylebox() StyleBox {
  classNameV := StringNameFromStr("ThemeDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fallback_stylebox")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 496040854) // FIXME: should cache?
  var ret StyleBox
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type ThemeDBFallbackChangedSignalFn func()

func (me *ThemeDB) ConnectFallbackChanged(subs SignalSubscribers, fn ThemeDBFallbackChangedSignalFn) {
  sig := StringNameFromStr("fallback_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *ThemeDB) DisconnectFallbackChanged(subs SignalSubscribers, fn ThemeDBFallbackChangedSignalFn) {
  sig := StringNameFromStr("fallback_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
