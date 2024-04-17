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

type ptrsForLightmapGIList struct {
	fnSetLightData               gdc.MethodBindPtr
	fnGetLightData               gdc.MethodBindPtr
	fnSetBakeQuality             gdc.MethodBindPtr
	fnGetBakeQuality             gdc.MethodBindPtr
	fnSetBounces                 gdc.MethodBindPtr
	fnGetBounces                 gdc.MethodBindPtr
	fnSetBounceIndirectEnergy    gdc.MethodBindPtr
	fnGetBounceIndirectEnergy    gdc.MethodBindPtr
	fnSetGenerateProbes          gdc.MethodBindPtr
	fnGetGenerateProbes          gdc.MethodBindPtr
	fnSetBias                    gdc.MethodBindPtr
	fnGetBias                    gdc.MethodBindPtr
	fnSetEnvironmentMode         gdc.MethodBindPtr
	fnGetEnvironmentMode         gdc.MethodBindPtr
	fnSetEnvironmentCustomSky    gdc.MethodBindPtr
	fnGetEnvironmentCustomSky    gdc.MethodBindPtr
	fnSetEnvironmentCustomColor  gdc.MethodBindPtr
	fnGetEnvironmentCustomColor  gdc.MethodBindPtr
	fnSetEnvironmentCustomEnergy gdc.MethodBindPtr
	fnGetEnvironmentCustomEnergy gdc.MethodBindPtr
	fnSetTexelScale              gdc.MethodBindPtr
	fnGetTexelScale              gdc.MethodBindPtr
	fnSetMaxTextureSize          gdc.MethodBindPtr
	fnGetMaxTextureSize          gdc.MethodBindPtr
	fnSetUseDenoiser             gdc.MethodBindPtr
	fnIsUsingDenoiser            gdc.MethodBindPtr
	fnSetDenoiserStrength        gdc.MethodBindPtr
	fnGetDenoiserStrength        gdc.MethodBindPtr
	fnSetInterior                gdc.MethodBindPtr
	fnIsInterior                 gdc.MethodBindPtr
	fnSetDirectional             gdc.MethodBindPtr
	fnIsDirectional              gdc.MethodBindPtr
	fnSetUseTextureForBounces    gdc.MethodBindPtr
	fnIsUsingTextureForBounces   gdc.MethodBindPtr
	fnSetCameraAttributes        gdc.MethodBindPtr
	fnGetCameraAttributes        gdc.MethodBindPtr
}

var ptrsForLightmapGI ptrsForLightmapGIList

func initLightmapGIPtrs(iface gdc.Interface) {

	className := StringNameFromStr("LightmapGI")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_light_data")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnSetLightData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1790597277))
	}
	{
		methodName := StringNameFromStr("get_light_data")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnGetLightData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 290354153))
	}
	{
		methodName := StringNameFromStr("set_bake_quality")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnSetBakeQuality = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1192215803))
	}
	{
		methodName := StringNameFromStr("get_bake_quality")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnGetBakeQuality = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 688832735))
	}
	{
		methodName := StringNameFromStr("set_bounces")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnSetBounces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_bounces")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnGetBounces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_bounce_indirect_energy")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnSetBounceIndirectEnergy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_bounce_indirect_energy")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnGetBounceIndirectEnergy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_generate_probes")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnSetGenerateProbes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 549981046))
	}
	{
		methodName := StringNameFromStr("get_generate_probes")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnGetGenerateProbes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3930596226))
	}
	{
		methodName := StringNameFromStr("set_bias")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnSetBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_bias")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnGetBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_environment_mode")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnSetEnvironmentMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2282650285))
	}
	{
		methodName := StringNameFromStr("get_environment_mode")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnGetEnvironmentMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4128646479))
	}
	{
		methodName := StringNameFromStr("set_environment_custom_sky")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnSetEnvironmentCustomSky = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3336722921))
	}
	{
		methodName := StringNameFromStr("get_environment_custom_sky")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnGetEnvironmentCustomSky = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1177136966))
	}
	{
		methodName := StringNameFromStr("set_environment_custom_color")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnSetEnvironmentCustomColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_environment_custom_color")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnGetEnvironmentCustomColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_environment_custom_energy")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnSetEnvironmentCustomEnergy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_environment_custom_energy")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnGetEnvironmentCustomEnergy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_texel_scale")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnSetTexelScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_texel_scale")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnGetTexelScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_max_texture_size")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnSetMaxTextureSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_max_texture_size")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnGetMaxTextureSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_use_denoiser")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnSetUseDenoiser = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_using_denoiser")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnIsUsingDenoiser = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_denoiser_strength")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnSetDenoiserStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_denoiser_strength")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnGetDenoiserStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_interior")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnSetInterior = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_interior")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnIsInterior = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_directional")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnSetDirectional = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_directional")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnIsDirectional = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_use_texture_for_bounces")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnSetUseTextureForBounces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_using_texture_for_bounces")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnIsUsingTextureForBounces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_camera_attributes")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnSetCameraAttributes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2817810567))
	}
	{
		methodName := StringNameFromStr("get_camera_attributes")
		defer methodName.Destroy()
		ptrsForLightmapGI.fnGetCameraAttributes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3921283215))
	}

}

type LightmapGI struct {
	VisualInstance3D
}

func (me *LightmapGI) BaseClass() string {
	return "LightmapGI"
}

func NewLightmapGI() *LightmapGI {
	str := StringNameFromStr("LightmapGI") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &LightmapGI{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type LightmapGIBakeQuality int

const (
	LightmapGIBakeQualityBakeQualityLow    LightmapGIBakeQuality = 0
	LightmapGIBakeQualityBakeQualityMedium LightmapGIBakeQuality = 1
	LightmapGIBakeQualityBakeQualityHigh   LightmapGIBakeQuality = 2
	LightmapGIBakeQualityBakeQualityUltra  LightmapGIBakeQuality = 3
)

type LightmapGIGenerateProbes int

const (
	LightmapGIGenerateProbesGenerateProbesDisabled LightmapGIGenerateProbes = 0
	LightmapGIGenerateProbesGenerateProbesSubdiv4  LightmapGIGenerateProbes = 1
	LightmapGIGenerateProbesGenerateProbesSubdiv8  LightmapGIGenerateProbes = 2
	LightmapGIGenerateProbesGenerateProbesSubdiv16 LightmapGIGenerateProbes = 3
	LightmapGIGenerateProbesGenerateProbesSubdiv32 LightmapGIGenerateProbes = 4
)

type LightmapGIBakeError int

const (
	LightmapGIBakeErrorBakeErrorOk                  LightmapGIBakeError = 0
	LightmapGIBakeErrorBakeErrorNoSceneRoot         LightmapGIBakeError = 1
	LightmapGIBakeErrorBakeErrorForeignData         LightmapGIBakeError = 2
	LightmapGIBakeErrorBakeErrorNoLightmapper       LightmapGIBakeError = 3
	LightmapGIBakeErrorBakeErrorNoSavePath          LightmapGIBakeError = 4
	LightmapGIBakeErrorBakeErrorNoMeshes            LightmapGIBakeError = 5
	LightmapGIBakeErrorBakeErrorMeshesInvalid       LightmapGIBakeError = 6
	LightmapGIBakeErrorBakeErrorCantCreateImage     LightmapGIBakeError = 7
	LightmapGIBakeErrorBakeErrorUserAborted         LightmapGIBakeError = 8
	LightmapGIBakeErrorBakeErrorTextureSizeTooSmall LightmapGIBakeError = 9
)

type LightmapGIEnvironmentMode int

const (
	LightmapGIEnvironmentModeEnvironmentModeDisabled    LightmapGIEnvironmentMode = 0
	LightmapGIEnvironmentModeEnvironmentModeScene       LightmapGIEnvironmentMode = 1
	LightmapGIEnvironmentModeEnvironmentModeCustomSky   LightmapGIEnvironmentMode = 2
	LightmapGIEnvironmentModeEnvironmentModeCustomColor LightmapGIEnvironmentMode = 3
)

func (me *LightmapGI) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *LightmapGI) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *LightmapGI) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *LightmapGI) SetLightData(data LightmapGIData) {
	cargs := []gdc.ConstTypePtr{data.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnSetLightData), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LightmapGI) GetLightData() LightmapGIData {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewLightmapGIData()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnGetLightData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *LightmapGI) SetBakeQuality(bake_quality LightmapGIBakeQuality) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bake_quality)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnSetBakeQuality), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LightmapGI) GetBakeQuality() LightmapGIBakeQuality {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret LightmapGIBakeQuality

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnGetBakeQuality), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *LightmapGI) SetBounces(bounces int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bounces)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnSetBounces), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LightmapGI) GetBounces() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnGetBounces), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *LightmapGI) SetBounceIndirectEnergy(bounce_indirect_energy float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bounce_indirect_energy)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnSetBounceIndirectEnergy), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LightmapGI) GetBounceIndirectEnergy() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnGetBounceIndirectEnergy), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *LightmapGI) SetGenerateProbes(subdivision LightmapGIGenerateProbes) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&subdivision)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnSetGenerateProbes), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LightmapGI) GetGenerateProbes() LightmapGIGenerateProbes {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret LightmapGIGenerateProbes

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnGetGenerateProbes), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *LightmapGI) SetBias(bias float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bias)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnSetBias), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LightmapGI) GetBias() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnGetBias), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *LightmapGI) SetEnvironmentMode(mode LightmapGIEnvironmentMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnSetEnvironmentMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LightmapGI) GetEnvironmentMode() LightmapGIEnvironmentMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret LightmapGIEnvironmentMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnGetEnvironmentMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *LightmapGI) SetEnvironmentCustomSky(sky Sky) {
	cargs := []gdc.ConstTypePtr{sky.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnSetEnvironmentCustomSky), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LightmapGI) GetEnvironmentCustomSky() Sky {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewSky()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnGetEnvironmentCustomSky), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *LightmapGI) SetEnvironmentCustomColor(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnSetEnvironmentCustomColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LightmapGI) GetEnvironmentCustomColor() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnGetEnvironmentCustomColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *LightmapGI) SetEnvironmentCustomEnergy(energy float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&energy)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnSetEnvironmentCustomEnergy), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LightmapGI) GetEnvironmentCustomEnergy() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnGetEnvironmentCustomEnergy), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *LightmapGI) SetTexelScale(texel_scale float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&texel_scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnSetTexelScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LightmapGI) GetTexelScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnGetTexelScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *LightmapGI) SetMaxTextureSize(max_texture_size int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_texture_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnSetMaxTextureSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LightmapGI) GetMaxTextureSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnGetMaxTextureSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *LightmapGI) SetUseDenoiser(use_denoiser bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_denoiser)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnSetUseDenoiser), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LightmapGI) IsUsingDenoiser() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnIsUsingDenoiser), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *LightmapGI) SetDenoiserStrength(denoiser_strength float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&denoiser_strength)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnSetDenoiserStrength), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LightmapGI) GetDenoiserStrength() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnGetDenoiserStrength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *LightmapGI) SetInterior(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnSetInterior), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LightmapGI) IsInterior() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnIsInterior), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *LightmapGI) SetDirectional(directional bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&directional)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnSetDirectional), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LightmapGI) IsDirectional() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnIsDirectional), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *LightmapGI) SetUseTextureForBounces(use_texture_for_bounces bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_texture_for_bounces)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnSetUseTextureForBounces), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LightmapGI) IsUsingTextureForBounces() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnIsUsingTextureForBounces), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *LightmapGI) SetCameraAttributes(camera_attributes CameraAttributes) {
	cargs := []gdc.ConstTypePtr{camera_attributes.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnSetCameraAttributes), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LightmapGI) GetCameraAttributes() CameraAttributes {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCameraAttributes()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGI.fnGetCameraAttributes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
