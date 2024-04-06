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

type ptrsForLabelSettingsList struct {
	fnSetLineSpacing  gdc.MethodBindPtr
	fnGetLineSpacing  gdc.MethodBindPtr
	fnSetFont         gdc.MethodBindPtr
	fnGetFont         gdc.MethodBindPtr
	fnSetFontSize     gdc.MethodBindPtr
	fnGetFontSize     gdc.MethodBindPtr
	fnSetFontColor    gdc.MethodBindPtr
	fnGetFontColor    gdc.MethodBindPtr
	fnSetOutlineSize  gdc.MethodBindPtr
	fnGetOutlineSize  gdc.MethodBindPtr
	fnSetOutlineColor gdc.MethodBindPtr
	fnGetOutlineColor gdc.MethodBindPtr
	fnSetShadowSize   gdc.MethodBindPtr
	fnGetShadowSize   gdc.MethodBindPtr
	fnSetShadowColor  gdc.MethodBindPtr
	fnGetShadowColor  gdc.MethodBindPtr
	fnSetShadowOffset gdc.MethodBindPtr
	fnGetShadowOffset gdc.MethodBindPtr
}

var ptrsForLabelSettings ptrsForLabelSettingsList

func initLabelSettingsPtrs(iface gdc.Interface) {

	className := StringNameFromStr("LabelSettings")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_line_spacing")
		defer methodName.Destroy()
		ptrsForLabelSettings.fnSetLineSpacing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_line_spacing")
		defer methodName.Destroy()
		ptrsForLabelSettings.fnGetLineSpacing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_font")
		defer methodName.Destroy()
		ptrsForLabelSettings.fnSetFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1262170328))
	}
	{
		methodName := StringNameFromStr("get_font")
		defer methodName.Destroy()
		ptrsForLabelSettings.fnGetFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3229501585))
	}
	{
		methodName := StringNameFromStr("set_font_size")
		defer methodName.Destroy()
		ptrsForLabelSettings.fnSetFontSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_font_size")
		defer methodName.Destroy()
		ptrsForLabelSettings.fnGetFontSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_font_color")
		defer methodName.Destroy()
		ptrsForLabelSettings.fnSetFontColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_font_color")
		defer methodName.Destroy()
		ptrsForLabelSettings.fnGetFontColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_outline_size")
		defer methodName.Destroy()
		ptrsForLabelSettings.fnSetOutlineSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_outline_size")
		defer methodName.Destroy()
		ptrsForLabelSettings.fnGetOutlineSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_outline_color")
		defer methodName.Destroy()
		ptrsForLabelSettings.fnSetOutlineColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_outline_color")
		defer methodName.Destroy()
		ptrsForLabelSettings.fnGetOutlineColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_shadow_size")
		defer methodName.Destroy()
		ptrsForLabelSettings.fnSetShadowSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_shadow_size")
		defer methodName.Destroy()
		ptrsForLabelSettings.fnGetShadowSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_shadow_color")
		defer methodName.Destroy()
		ptrsForLabelSettings.fnSetShadowColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_shadow_color")
		defer methodName.Destroy()
		ptrsForLabelSettings.fnGetShadowColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_shadow_offset")
		defer methodName.Destroy()
		ptrsForLabelSettings.fnSetShadowOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_shadow_offset")
		defer methodName.Destroy()
		ptrsForLabelSettings.fnGetShadowOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}

}

type LabelSettings struct {
	Resource
}

func (me *LabelSettings) BaseClass() string {
	return "LabelSettings"
}

func NewLabelSettings() *LabelSettings {
	str := StringNameFromStr("LabelSettings") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &LabelSettings{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *LabelSettings) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *LabelSettings) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *LabelSettings) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *LabelSettings) SetLineSpacing(spacing float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&spacing)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabelSettings.fnSetLineSpacing), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LabelSettings) GetLineSpacing() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabelSettings.fnGetLineSpacing), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *LabelSettings) SetFont(font Font) {
	cargs := []gdc.ConstTypePtr{font.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabelSettings.fnSetFont), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LabelSettings) GetFont() Font {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFont()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabelSettings.fnGetFont), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *LabelSettings) SetFontSize(size int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabelSettings.fnSetFontSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LabelSettings) GetFontSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabelSettings.fnGetFontSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *LabelSettings) SetFontColor(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabelSettings.fnSetFontColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LabelSettings) GetFontColor() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabelSettings.fnGetFontColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *LabelSettings) SetOutlineSize(size int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabelSettings.fnSetOutlineSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LabelSettings) GetOutlineSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabelSettings.fnGetOutlineSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *LabelSettings) SetOutlineColor(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabelSettings.fnSetOutlineColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LabelSettings) GetOutlineColor() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabelSettings.fnGetOutlineColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *LabelSettings) SetShadowSize(size int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabelSettings.fnSetShadowSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LabelSettings) GetShadowSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabelSettings.fnGetShadowSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *LabelSettings) SetShadowColor(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabelSettings.fnSetShadowColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LabelSettings) GetShadowColor() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabelSettings.fnGetShadowColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *LabelSettings) SetShadowOffset(offset Vector2) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabelSettings.fnSetShadowOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LabelSettings) GetShadowOffset() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabelSettings.fnGetShadowOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
