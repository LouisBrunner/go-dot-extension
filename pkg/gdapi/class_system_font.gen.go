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

type ptrsForSystemFontList struct {
	fnSetAntialiasing                    gdc.MethodBindPtr
	fnGetAntialiasing                    gdc.MethodBindPtr
	fnSetGenerateMipmaps                 gdc.MethodBindPtr
	fnGetGenerateMipmaps                 gdc.MethodBindPtr
	fnSetAllowSystemFallback             gdc.MethodBindPtr
	fnIsAllowSystemFallback              gdc.MethodBindPtr
	fnSetForceAutohinter                 gdc.MethodBindPtr
	fnIsForceAutohinter                  gdc.MethodBindPtr
	fnSetHinting                         gdc.MethodBindPtr
	fnGetHinting                         gdc.MethodBindPtr
	fnSetSubpixelPositioning             gdc.MethodBindPtr
	fnGetSubpixelPositioning             gdc.MethodBindPtr
	fnSetMultichannelSignedDistanceField gdc.MethodBindPtr
	fnIsMultichannelSignedDistanceField  gdc.MethodBindPtr
	fnSetMsdfPixelRange                  gdc.MethodBindPtr
	fnGetMsdfPixelRange                  gdc.MethodBindPtr
	fnSetMsdfSize                        gdc.MethodBindPtr
	fnGetMsdfSize                        gdc.MethodBindPtr
	fnSetOversampling                    gdc.MethodBindPtr
	fnGetOversampling                    gdc.MethodBindPtr
	fnGetFontNames                       gdc.MethodBindPtr
	fnSetFontNames                       gdc.MethodBindPtr
	fnGetFontItalic                      gdc.MethodBindPtr
	fnSetFontItalic                      gdc.MethodBindPtr
	fnSetFontWeight                      gdc.MethodBindPtr
	fnSetFontStretch                     gdc.MethodBindPtr
}

var ptrsForSystemFont ptrsForSystemFontList

func initSystemFontPtrs(iface gdc.Interface) {

	className := StringNameFromStr("SystemFont")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_antialiasing")
		defer methodName.Destroy()
		ptrsForSystemFont.fnSetAntialiasing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1669900))
	}
	{
		methodName := StringNameFromStr("get_antialiasing")
		defer methodName.Destroy()
		ptrsForSystemFont.fnGetAntialiasing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4262718649))
	}
	{
		methodName := StringNameFromStr("set_generate_mipmaps")
		defer methodName.Destroy()
		ptrsForSystemFont.fnSetGenerateMipmaps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_generate_mipmaps")
		defer methodName.Destroy()
		ptrsForSystemFont.fnGetGenerateMipmaps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_allow_system_fallback")
		defer methodName.Destroy()
		ptrsForSystemFont.fnSetAllowSystemFallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_allow_system_fallback")
		defer methodName.Destroy()
		ptrsForSystemFont.fnIsAllowSystemFallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_force_autohinter")
		defer methodName.Destroy()
		ptrsForSystemFont.fnSetForceAutohinter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_force_autohinter")
		defer methodName.Destroy()
		ptrsForSystemFont.fnIsForceAutohinter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_hinting")
		defer methodName.Destroy()
		ptrsForSystemFont.fnSetHinting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1827459492))
	}
	{
		methodName := StringNameFromStr("get_hinting")
		defer methodName.Destroy()
		ptrsForSystemFont.fnGetHinting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3683214614))
	}
	{
		methodName := StringNameFromStr("set_subpixel_positioning")
		defer methodName.Destroy()
		ptrsForSystemFont.fnSetSubpixelPositioning = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4225742182))
	}
	{
		methodName := StringNameFromStr("get_subpixel_positioning")
		defer methodName.Destroy()
		ptrsForSystemFont.fnGetSubpixelPositioning = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1069238588))
	}
	{
		methodName := StringNameFromStr("set_multichannel_signed_distance_field")
		defer methodName.Destroy()
		ptrsForSystemFont.fnSetMultichannelSignedDistanceField = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_multichannel_signed_distance_field")
		defer methodName.Destroy()
		ptrsForSystemFont.fnIsMultichannelSignedDistanceField = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_msdf_pixel_range")
		defer methodName.Destroy()
		ptrsForSystemFont.fnSetMsdfPixelRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_msdf_pixel_range")
		defer methodName.Destroy()
		ptrsForSystemFont.fnGetMsdfPixelRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_msdf_size")
		defer methodName.Destroy()
		ptrsForSystemFont.fnSetMsdfSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_msdf_size")
		defer methodName.Destroy()
		ptrsForSystemFont.fnGetMsdfSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_oversampling")
		defer methodName.Destroy()
		ptrsForSystemFont.fnSetOversampling = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_oversampling")
		defer methodName.Destroy()
		ptrsForSystemFont.fnGetOversampling = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_font_names")
		defer methodName.Destroy()
		ptrsForSystemFont.fnGetFontNames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("set_font_names")
		defer methodName.Destroy()
		ptrsForSystemFont.fnSetFontNames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4015028928))
	}
	{
		methodName := StringNameFromStr("get_font_italic")
		defer methodName.Destroy()
		ptrsForSystemFont.fnGetFontItalic = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_font_italic")
		defer methodName.Destroy()
		ptrsForSystemFont.fnSetFontItalic = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_font_weight")
		defer methodName.Destroy()
		ptrsForSystemFont.fnSetFontWeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_font_stretch")
		defer methodName.Destroy()
		ptrsForSystemFont.fnSetFontStretch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}

}

type SystemFont struct {
	Font
}

func (me *SystemFont) BaseClass() string {
	return "SystemFont"
}

func NewSystemFont() *SystemFont {
	str := StringNameFromStr("SystemFont") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &SystemFont{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *SystemFont) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *SystemFont) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *SystemFont) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *SystemFont) SetAntialiasing(antialiasing TextServerFontAntialiasing) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&antialiasing)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSystemFont.fnSetAntialiasing), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SystemFont) GetAntialiasing() TextServerFontAntialiasing {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerFontAntialiasing

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSystemFont.fnGetAntialiasing), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *SystemFont) SetGenerateMipmaps(generate_mipmaps bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&generate_mipmaps)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSystemFont.fnSetGenerateMipmaps), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SystemFont) GetGenerateMipmaps() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSystemFont.fnGetGenerateMipmaps), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SystemFont) SetAllowSystemFallback(allow_system_fallback bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow_system_fallback)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSystemFont.fnSetAllowSystemFallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SystemFont) IsAllowSystemFallback() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSystemFont.fnIsAllowSystemFallback), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SystemFont) SetForceAutohinter(force_autohinter bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&force_autohinter)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSystemFont.fnSetForceAutohinter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SystemFont) IsForceAutohinter() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSystemFont.fnIsForceAutohinter), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SystemFont) SetHinting(hinting TextServerHinting) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hinting)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSystemFont.fnSetHinting), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SystemFont) GetHinting() TextServerHinting {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerHinting

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSystemFont.fnGetHinting), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *SystemFont) SetSubpixelPositioning(subpixel_positioning TextServerSubpixelPositioning) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&subpixel_positioning)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSystemFont.fnSetSubpixelPositioning), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SystemFont) GetSubpixelPositioning() TextServerSubpixelPositioning {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerSubpixelPositioning

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSystemFont.fnGetSubpixelPositioning), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *SystemFont) SetMultichannelSignedDistanceField(msdf bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&msdf)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSystemFont.fnSetMultichannelSignedDistanceField), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SystemFont) IsMultichannelSignedDistanceField() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSystemFont.fnIsMultichannelSignedDistanceField), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SystemFont) SetMsdfPixelRange(msdf_pixel_range int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&msdf_pixel_range)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSystemFont.fnSetMsdfPixelRange), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SystemFont) GetMsdfPixelRange() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSystemFont.fnGetMsdfPixelRange), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SystemFont) SetMsdfSize(msdf_size int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&msdf_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSystemFont.fnSetMsdfSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SystemFont) GetMsdfSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSystemFont.fnGetMsdfSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SystemFont) SetOversampling(oversampling float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&oversampling)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSystemFont.fnSetOversampling), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SystemFont) GetOversampling() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSystemFont.fnGetOversampling), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SystemFont) GetFontNames() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSystemFont.fnGetFontNames), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SystemFont) SetFontNames(names PackedStringArray) {
	cargs := []gdc.ConstTypePtr{names.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSystemFont.fnSetFontNames), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SystemFont) GetFontItalic() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSystemFont.fnGetFontItalic), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SystemFont) SetFontItalic(italic bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&italic)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSystemFont.fnSetFontItalic), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SystemFont) SetFontWeight(weight int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&weight)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSystemFont.fnSetFontWeight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SystemFont) SetFontStretch(stretch int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stretch)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSystemFont.fnSetFontStretch), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
