// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ThemeDB struct {
  obj gdc.ObjectPtr
}

func (me *ThemeDB) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
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

func (me *ThemeDB) GetPropFallbackBaseScale() float32 {
  panic("TODO: implement")
}

func (me *ThemeDB) SetPropFallbackBaseScale(value float32) {
  panic("TODO: implement")
}

func (me *ThemeDB) GetPropFallbackFont() Font {
  panic("TODO: implement")
}

func (me *ThemeDB) SetPropFallbackFont(value Font) {
  panic("TODO: implement")
}

func (me *ThemeDB) GetPropFallbackFontSize() int {
  panic("TODO: implement")
}

func (me *ThemeDB) SetPropFallbackFontSize(value int) {
  panic("TODO: implement")
}

func (me *ThemeDB) GetPropFallbackIcon() Texture2D {
  panic("TODO: implement")
}

func (me *ThemeDB) SetPropFallbackIcon(value Texture2D) {
  panic("TODO: implement")
}

func (me *ThemeDB) GetPropFallbackStylebox() StyleBox {
  panic("TODO: implement")
}

func (me *ThemeDB) SetPropFallbackStylebox(value StyleBox) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API