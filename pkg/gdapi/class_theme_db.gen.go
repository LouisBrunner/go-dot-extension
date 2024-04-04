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

type ThemeDB struct {
  Object
}

func (me *ThemeDB) BaseClass() string {
  return "ThemeDB"
}

func NewThemeDB() *ThemeDB {
  str := StringNameFromStr("ThemeDB") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ThemeDB{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTheme()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ThemeDB) GetProjectTheme() Theme {
  classNameV := StringNameFromStr("ThemeDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_project_theme")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 754276358) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTheme()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ThemeDB) SetFallbackBaseScale(base_scale float64, )  {
  classNameV := StringNameFromStr("ThemeDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fallback_base_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&base_scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ThemeDB) GetFallbackBaseScale() float64 {
  classNameV := StringNameFromStr("ThemeDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fallback_base_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ThemeDB) SetFallbackFont(font Font, )  {
  classNameV := StringNameFromStr("ThemeDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fallback_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1262170328) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{font.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ThemeDB) GetFallbackFont() Font {
  classNameV := StringNameFromStr("ThemeDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fallback_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3656929885) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFont()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ThemeDB) SetFallbackFontSize(font_size int64, )  {
  classNameV := StringNameFromStr("ThemeDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fallback_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&font_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ThemeDB) GetFallbackFontSize() int64 {
  classNameV := StringNameFromStr("ThemeDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fallback_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ThemeDB) SetFallbackIcon(icon Texture2D, )  {
  classNameV := StringNameFromStr("ThemeDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fallback_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{icon.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ThemeDB) GetFallbackIcon() Texture2D {
  classNameV := StringNameFromStr("ThemeDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fallback_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 255860311) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ThemeDB) SetFallbackStylebox(stylebox StyleBox, )  {
  classNameV := StringNameFromStr("ThemeDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fallback_stylebox")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2797200388) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{stylebox.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ThemeDB) GetFallbackStylebox() StyleBox {
  classNameV := StringNameFromStr("ThemeDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fallback_stylebox")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 496040854) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStyleBox()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type ThemeDBFallbackChangedSignalFn func()

func (me *ThemeDB) ConnectFallbackChanged(subs SignalSubscribers, fn ThemeDBFallbackChangedSignalFn) {
  sig := StringNameFromStr("fallback_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ThemeDB) DisconnectFallbackChanged(subs SignalSubscribers, fn ThemeDBFallbackChangedSignalFn) {
  sig := StringNameFromStr("fallback_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
