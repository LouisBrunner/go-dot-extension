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

type ptrsForNoiseTexture2DList struct {
	fnSetWidth              gdc.MethodBindPtr
	fnSetHeight             gdc.MethodBindPtr
	fnSetInvert             gdc.MethodBindPtr
	fnGetInvert             gdc.MethodBindPtr
	fnSetIn3DSpace          gdc.MethodBindPtr
	fnIsIn3DSpace           gdc.MethodBindPtr
	fnSetGenerateMipmaps    gdc.MethodBindPtr
	fnIsGeneratingMipmaps   gdc.MethodBindPtr
	fnSetSeamless           gdc.MethodBindPtr
	fnGetSeamless           gdc.MethodBindPtr
	fnSetSeamlessBlendSkirt gdc.MethodBindPtr
	fnGetSeamlessBlendSkirt gdc.MethodBindPtr
	fnSetAsNormalMap        gdc.MethodBindPtr
	fnIsNormalMap           gdc.MethodBindPtr
	fnSetBumpStrength       gdc.MethodBindPtr
	fnGetBumpStrength       gdc.MethodBindPtr
	fnSetNormalize          gdc.MethodBindPtr
	fnIsNormalized          gdc.MethodBindPtr
	fnSetColorRamp          gdc.MethodBindPtr
	fnGetColorRamp          gdc.MethodBindPtr
	fnSetNoise              gdc.MethodBindPtr
	fnGetNoise              gdc.MethodBindPtr
}

var ptrsForNoiseTexture2D ptrsForNoiseTexture2DList

func initNoiseTexture2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("NoiseTexture2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_width")
		defer methodName.Destroy()
		ptrsForNoiseTexture2D.fnSetWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_height")
		defer methodName.Destroy()
		ptrsForNoiseTexture2D.fnSetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_invert")
		defer methodName.Destroy()
		ptrsForNoiseTexture2D.fnSetInvert = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_invert")
		defer methodName.Destroy()
		ptrsForNoiseTexture2D.fnGetInvert = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_in_3d_space")
		defer methodName.Destroy()
		ptrsForNoiseTexture2D.fnSetIn3DSpace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_in_3d_space")
		defer methodName.Destroy()
		ptrsForNoiseTexture2D.fnIsIn3DSpace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_generate_mipmaps")
		defer methodName.Destroy()
		ptrsForNoiseTexture2D.fnSetGenerateMipmaps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_generating_mipmaps")
		defer methodName.Destroy()
		ptrsForNoiseTexture2D.fnIsGeneratingMipmaps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_seamless")
		defer methodName.Destroy()
		ptrsForNoiseTexture2D.fnSetSeamless = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_seamless")
		defer methodName.Destroy()
		ptrsForNoiseTexture2D.fnGetSeamless = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("set_seamless_blend_skirt")
		defer methodName.Destroy()
		ptrsForNoiseTexture2D.fnSetSeamlessBlendSkirt = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_seamless_blend_skirt")
		defer methodName.Destroy()
		ptrsForNoiseTexture2D.fnGetSeamlessBlendSkirt = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
	}
	{
		methodName := StringNameFromStr("set_as_normal_map")
		defer methodName.Destroy()
		ptrsForNoiseTexture2D.fnSetAsNormalMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_normal_map")
		defer methodName.Destroy()
		ptrsForNoiseTexture2D.fnIsNormalMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("set_bump_strength")
		defer methodName.Destroy()
		ptrsForNoiseTexture2D.fnSetBumpStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_bump_strength")
		defer methodName.Destroy()
		ptrsForNoiseTexture2D.fnGetBumpStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
	}
	{
		methodName := StringNameFromStr("set_normalize")
		defer methodName.Destroy()
		ptrsForNoiseTexture2D.fnSetNormalize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_normalized")
		defer methodName.Destroy()
		ptrsForNoiseTexture2D.fnIsNormalized = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_color_ramp")
		defer methodName.Destroy()
		ptrsForNoiseTexture2D.fnSetColorRamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2756054477))
	}
	{
		methodName := StringNameFromStr("get_color_ramp")
		defer methodName.Destroy()
		ptrsForNoiseTexture2D.fnGetColorRamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 132272999))
	}
	{
		methodName := StringNameFromStr("set_noise")
		defer methodName.Destroy()
		ptrsForNoiseTexture2D.fnSetNoise = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4135492439))
	}
	{
		methodName := StringNameFromStr("get_noise")
		defer methodName.Destroy()
		ptrsForNoiseTexture2D.fnGetNoise = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 185851837))
	}

}

type NoiseTexture2D struct {
	Texture2D
}

func (me *NoiseTexture2D) BaseClass() string {
	return "NoiseTexture2D"
}

func NewNoiseTexture2D() *NoiseTexture2D {
	str := StringNameFromStr("NoiseTexture2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &NoiseTexture2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *NoiseTexture2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *NoiseTexture2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *NoiseTexture2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *NoiseTexture2D) SetWidth(width int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture2D.fnSetWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NoiseTexture2D) SetHeight(height int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture2D.fnSetHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NoiseTexture2D) SetInvert(invert bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&invert)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture2D.fnSetInvert), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NoiseTexture2D) GetInvert() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture2D.fnGetInvert), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NoiseTexture2D) SetIn3DSpace(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture2D.fnSetIn3DSpace), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NoiseTexture2D) IsIn3DSpace() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture2D.fnIsIn3DSpace), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NoiseTexture2D) SetGenerateMipmaps(invert bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&invert)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture2D.fnSetGenerateMipmaps), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NoiseTexture2D) IsGeneratingMipmaps() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture2D.fnIsGeneratingMipmaps), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NoiseTexture2D) SetSeamless(seamless bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seamless)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture2D.fnSetSeamless), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NoiseTexture2D) GetSeamless() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture2D.fnGetSeamless), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NoiseTexture2D) SetSeamlessBlendSkirt(seamless_blend_skirt float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seamless_blend_skirt)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture2D.fnSetSeamlessBlendSkirt), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NoiseTexture2D) GetSeamlessBlendSkirt() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture2D.fnGetSeamlessBlendSkirt), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NoiseTexture2D) SetAsNormalMap(as_normal_map bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&as_normal_map)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture2D.fnSetAsNormalMap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NoiseTexture2D) IsNormalMap() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture2D.fnIsNormalMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NoiseTexture2D) SetBumpStrength(bump_strength float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bump_strength)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture2D.fnSetBumpStrength), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NoiseTexture2D) GetBumpStrength() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture2D.fnGetBumpStrength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NoiseTexture2D) SetNormalize(normalize bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&normalize)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture2D.fnSetNormalize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NoiseTexture2D) IsNormalized() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture2D.fnIsNormalized), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NoiseTexture2D) SetColorRamp(gradient Gradient) {
	cargs := []gdc.ConstTypePtr{gradient.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture2D.fnSetColorRamp), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NoiseTexture2D) GetColorRamp() Gradient {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewGradient()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture2D.fnGetColorRamp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NoiseTexture2D) SetNoise(noise Noise) {
	cargs := []gdc.ConstTypePtr{noise.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture2D.fnSetNoise), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NoiseTexture2D) GetNoise() Noise {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNoise()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoiseTexture2D.fnGetNoise), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
