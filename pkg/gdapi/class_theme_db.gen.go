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

type ptrsForThemeDBList struct {
	fnGetDefaultTheme      gdc.MethodBindPtr
	fnGetProjectTheme      gdc.MethodBindPtr
	fnSetFallbackBaseScale gdc.MethodBindPtr
	fnGetFallbackBaseScale gdc.MethodBindPtr
	fnSetFallbackFont      gdc.MethodBindPtr
	fnGetFallbackFont      gdc.MethodBindPtr
	fnSetFallbackFontSize  gdc.MethodBindPtr
	fnGetFallbackFontSize  gdc.MethodBindPtr
	fnSetFallbackIcon      gdc.MethodBindPtr
	fnGetFallbackIcon      gdc.MethodBindPtr
	fnSetFallbackStylebox  gdc.MethodBindPtr
	fnGetFallbackStylebox  gdc.MethodBindPtr
}

var ptrsForThemeDB ptrsForThemeDBList

func initThemeDBPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ThemeDB")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_default_theme")
		defer methodName.Destroy()
		ptrsForThemeDB.fnGetDefaultTheme = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 754276358))
	}
	{
		methodName := StringNameFromStr("get_project_theme")
		defer methodName.Destroy()
		ptrsForThemeDB.fnGetProjectTheme = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 754276358))
	}
	{
		methodName := StringNameFromStr("set_fallback_base_scale")
		defer methodName.Destroy()
		ptrsForThemeDB.fnSetFallbackBaseScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_fallback_base_scale")
		defer methodName.Destroy()
		ptrsForThemeDB.fnGetFallbackBaseScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
	}
	{
		methodName := StringNameFromStr("set_fallback_font")
		defer methodName.Destroy()
		ptrsForThemeDB.fnSetFallbackFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1262170328))
	}
	{
		methodName := StringNameFromStr("get_fallback_font")
		defer methodName.Destroy()
		ptrsForThemeDB.fnGetFallbackFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3656929885))
	}
	{
		methodName := StringNameFromStr("set_fallback_font_size")
		defer methodName.Destroy()
		ptrsForThemeDB.fnSetFallbackFontSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_fallback_font_size")
		defer methodName.Destroy()
		ptrsForThemeDB.fnGetFallbackFontSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("set_fallback_icon")
		defer methodName.Destroy()
		ptrsForThemeDB.fnSetFallbackIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_fallback_icon")
		defer methodName.Destroy()
		ptrsForThemeDB.fnGetFallbackIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 255860311))
	}
	{
		methodName := StringNameFromStr("set_fallback_stylebox")
		defer methodName.Destroy()
		ptrsForThemeDB.fnSetFallbackStylebox = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2797200388))
	}
	{
		methodName := StringNameFromStr("get_fallback_stylebox")
		defer methodName.Destroy()
		ptrsForThemeDB.fnGetFallbackStylebox = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 496040854))
	}

}

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

func (me *ThemeDB) GetDefaultTheme() Theme {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTheme()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForThemeDB.fnGetDefaultTheme), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ThemeDB) GetProjectTheme() Theme {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTheme()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForThemeDB.fnGetProjectTheme), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ThemeDB) SetFallbackBaseScale(base_scale float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&base_scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForThemeDB.fnSetFallbackBaseScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ThemeDB) GetFallbackBaseScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForThemeDB.fnGetFallbackBaseScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ThemeDB) SetFallbackFont(font Font) {
	cargs := []gdc.ConstTypePtr{font.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForThemeDB.fnSetFallbackFont), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ThemeDB) GetFallbackFont() Font {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFont()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForThemeDB.fnGetFallbackFont), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ThemeDB) SetFallbackFontSize(font_size int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&font_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForThemeDB.fnSetFallbackFontSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ThemeDB) GetFallbackFontSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForThemeDB.fnGetFallbackFontSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ThemeDB) SetFallbackIcon(icon Texture2D) {
	cargs := []gdc.ConstTypePtr{icon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForThemeDB.fnSetFallbackIcon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ThemeDB) GetFallbackIcon() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForThemeDB.fnGetFallbackIcon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ThemeDB) SetFallbackStylebox(stylebox StyleBox) {
	cargs := []gdc.ConstTypePtr{stylebox.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForThemeDB.fnSetFallbackStylebox), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ThemeDB) GetFallbackStylebox() StyleBox {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStyleBox()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForThemeDB.fnGetFallbackStylebox), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
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
