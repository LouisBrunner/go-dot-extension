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

type ptrsForGeometryInstance3DList struct {
	fnSetMaterialOverride           gdc.MethodBindPtr
	fnGetMaterialOverride           gdc.MethodBindPtr
	fnSetMaterialOverlay            gdc.MethodBindPtr
	fnGetMaterialOverlay            gdc.MethodBindPtr
	fnSetCastShadowsSetting         gdc.MethodBindPtr
	fnGetCastShadowsSetting         gdc.MethodBindPtr
	fnSetLodBias                    gdc.MethodBindPtr
	fnGetLodBias                    gdc.MethodBindPtr
	fnSetTransparency               gdc.MethodBindPtr
	fnGetTransparency               gdc.MethodBindPtr
	fnSetVisibilityRangeEndMargin   gdc.MethodBindPtr
	fnGetVisibilityRangeEndMargin   gdc.MethodBindPtr
	fnSetVisibilityRangeEnd         gdc.MethodBindPtr
	fnGetVisibilityRangeEnd         gdc.MethodBindPtr
	fnSetVisibilityRangeBeginMargin gdc.MethodBindPtr
	fnGetVisibilityRangeBeginMargin gdc.MethodBindPtr
	fnSetVisibilityRangeBegin       gdc.MethodBindPtr
	fnGetVisibilityRangeBegin       gdc.MethodBindPtr
	fnSetVisibilityRangeFadeMode    gdc.MethodBindPtr
	fnGetVisibilityRangeFadeMode    gdc.MethodBindPtr
	fnSetInstanceShaderParameter    gdc.MethodBindPtr
	fnGetInstanceShaderParameter    gdc.MethodBindPtr
	fnSetExtraCullMargin            gdc.MethodBindPtr
	fnGetExtraCullMargin            gdc.MethodBindPtr
	fnSetLightmapScale              gdc.MethodBindPtr
	fnGetLightmapScale              gdc.MethodBindPtr
	fnSetGiMode                     gdc.MethodBindPtr
	fnGetGiMode                     gdc.MethodBindPtr
	fnSetIgnoreOcclusionCulling     gdc.MethodBindPtr
	fnIsIgnoringOcclusionCulling    gdc.MethodBindPtr
	fnSetCustomAabb                 gdc.MethodBindPtr
	fnGetCustomAabb                 gdc.MethodBindPtr
}

var ptrsForGeometryInstance3D ptrsForGeometryInstance3DList

func initGeometryInstance3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("GeometryInstance3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_material_override")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnSetMaterialOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2757459619))
	}
	{
		methodName := StringNameFromStr("get_material_override")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnGetMaterialOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 5934680))
	}
	{
		methodName := StringNameFromStr("set_material_overlay")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnSetMaterialOverlay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2757459619))
	}
	{
		methodName := StringNameFromStr("get_material_overlay")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnGetMaterialOverlay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 5934680))
	}
	{
		methodName := StringNameFromStr("set_cast_shadows_setting")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnSetCastShadowsSetting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 856677339))
	}
	{
		methodName := StringNameFromStr("get_cast_shadows_setting")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnGetCastShadowsSetting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3383019359))
	}
	{
		methodName := StringNameFromStr("set_lod_bias")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnSetLodBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_lod_bias")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnGetLodBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_transparency")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnSetTransparency = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_transparency")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnGetTransparency = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_visibility_range_end_margin")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnSetVisibilityRangeEndMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_visibility_range_end_margin")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnGetVisibilityRangeEndMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_visibility_range_end")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnSetVisibilityRangeEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_visibility_range_end")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnGetVisibilityRangeEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_visibility_range_begin_margin")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnSetVisibilityRangeBeginMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_visibility_range_begin_margin")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnGetVisibilityRangeBeginMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_visibility_range_begin")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnSetVisibilityRangeBegin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_visibility_range_begin")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnGetVisibilityRangeBegin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_visibility_range_fade_mode")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnSetVisibilityRangeFadeMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1440117808))
	}
	{
		methodName := StringNameFromStr("get_visibility_range_fade_mode")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnGetVisibilityRangeFadeMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2067221882))
	}
	{
		methodName := StringNameFromStr("set_instance_shader_parameter")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnSetInstanceShaderParameter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3776071444))
	}
	{
		methodName := StringNameFromStr("get_instance_shader_parameter")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnGetInstanceShaderParameter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2760726917))
	}
	{
		methodName := StringNameFromStr("set_extra_cull_margin")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnSetExtraCullMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_extra_cull_margin")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnGetExtraCullMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_lightmap_scale")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnSetLightmapScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2462696582))
	}
	{
		methodName := StringNameFromStr("get_lightmap_scale")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnGetLightmapScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 798767852))
	}
	{
		methodName := StringNameFromStr("set_gi_mode")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnSetGiMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2548557163))
	}
	{
		methodName := StringNameFromStr("get_gi_mode")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnGetGiMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2188566509))
	}
	{
		methodName := StringNameFromStr("set_ignore_occlusion_culling")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnSetIgnoreOcclusionCulling = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_ignoring_occlusion_culling")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnIsIgnoringOcclusionCulling = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("set_custom_aabb")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnSetCustomAabb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 259215842))
	}
	{
		methodName := StringNameFromStr("get_custom_aabb")
		defer methodName.Destroy()
		ptrsForGeometryInstance3D.fnGetCustomAabb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1068685055))
	}

}

type GeometryInstance3D struct {
	VisualInstance3D
}

func (me *GeometryInstance3D) BaseClass() string {
	return "GeometryInstance3D"
}

func NewGeometryInstance3D() *GeometryInstance3D {
	str := StringNameFromStr("GeometryInstance3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GeometryInstance3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type GeometryInstance3DShadowCastingSetting int

const (
	GeometryInstance3DShadowCastingSettingShadowCastingSettingOff         GeometryInstance3DShadowCastingSetting = 0
	GeometryInstance3DShadowCastingSettingShadowCastingSettingOn          GeometryInstance3DShadowCastingSetting = 1
	GeometryInstance3DShadowCastingSettingShadowCastingSettingDoubleSided GeometryInstance3DShadowCastingSetting = 2
	GeometryInstance3DShadowCastingSettingShadowCastingSettingShadowsOnly GeometryInstance3DShadowCastingSetting = 3
)

type GeometryInstance3DGIMode int

const (
	GeometryInstance3DGIModeGiModeDisabled GeometryInstance3DGIMode = 0
	GeometryInstance3DGIModeGiModeStatic   GeometryInstance3DGIMode = 1
	GeometryInstance3DGIModeGiModeDynamic  GeometryInstance3DGIMode = 2
)

type GeometryInstance3DLightmapScale int

const (
	GeometryInstance3DLightmapScaleLightmapScale1X  GeometryInstance3DLightmapScale = 0
	GeometryInstance3DLightmapScaleLightmapScale2X  GeometryInstance3DLightmapScale = 1
	GeometryInstance3DLightmapScaleLightmapScale4X  GeometryInstance3DLightmapScale = 2
	GeometryInstance3DLightmapScaleLightmapScale8X  GeometryInstance3DLightmapScale = 3
	GeometryInstance3DLightmapScaleLightmapScaleMax GeometryInstance3DLightmapScale = 4
)

type GeometryInstance3DVisibilityRangeFadeMode int

const (
	GeometryInstance3DVisibilityRangeFadeModeVisibilityRangeFadeDisabled     GeometryInstance3DVisibilityRangeFadeMode = 0
	GeometryInstance3DVisibilityRangeFadeModeVisibilityRangeFadeSelf         GeometryInstance3DVisibilityRangeFadeMode = 1
	GeometryInstance3DVisibilityRangeFadeModeVisibilityRangeFadeDependencies GeometryInstance3DVisibilityRangeFadeMode = 2
)

func (me *GeometryInstance3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GeometryInstance3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GeometryInstance3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *GeometryInstance3D) SetMaterialOverride(material Material) {
	cargs := []gdc.ConstTypePtr{material.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnSetMaterialOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GeometryInstance3D) GetMaterialOverride() Material {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMaterial()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnGetMaterialOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GeometryInstance3D) SetMaterialOverlay(material Material) {
	cargs := []gdc.ConstTypePtr{material.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnSetMaterialOverlay), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GeometryInstance3D) GetMaterialOverlay() Material {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMaterial()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnGetMaterialOverlay), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GeometryInstance3D) SetCastShadowsSetting(shadow_casting_setting GeometryInstance3DShadowCastingSetting) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shadow_casting_setting)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnSetCastShadowsSetting), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GeometryInstance3D) GetCastShadowsSetting() GeometryInstance3DShadowCastingSetting {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret GeometryInstance3DShadowCastingSetting

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnGetCastShadowsSetting), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *GeometryInstance3D) SetLodBias(bias float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bias)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnSetLodBias), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GeometryInstance3D) GetLodBias() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnGetLodBias), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GeometryInstance3D) SetTransparency(transparency float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&transparency)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnSetTransparency), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GeometryInstance3D) GetTransparency() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnGetTransparency), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GeometryInstance3D) SetVisibilityRangeEndMargin(distance float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnSetVisibilityRangeEndMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GeometryInstance3D) GetVisibilityRangeEndMargin() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnGetVisibilityRangeEndMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GeometryInstance3D) SetVisibilityRangeEnd(distance float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnSetVisibilityRangeEnd), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GeometryInstance3D) GetVisibilityRangeEnd() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnGetVisibilityRangeEnd), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GeometryInstance3D) SetVisibilityRangeBeginMargin(distance float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnSetVisibilityRangeBeginMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GeometryInstance3D) GetVisibilityRangeBeginMargin() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnGetVisibilityRangeBeginMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GeometryInstance3D) SetVisibilityRangeBegin(distance float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnSetVisibilityRangeBegin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GeometryInstance3D) GetVisibilityRangeBegin() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnGetVisibilityRangeBegin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GeometryInstance3D) SetVisibilityRangeFadeMode(mode GeometryInstance3DVisibilityRangeFadeMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnSetVisibilityRangeFadeMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GeometryInstance3D) GetVisibilityRangeFadeMode() GeometryInstance3DVisibilityRangeFadeMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret GeometryInstance3DVisibilityRangeFadeMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnGetVisibilityRangeFadeMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *GeometryInstance3D) SetInstanceShaderParameter(name StringName, value Variant) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnSetInstanceShaderParameter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GeometryInstance3D) GetInstanceShaderParameter(name StringName) Variant {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnGetInstanceShaderParameter), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GeometryInstance3D) SetExtraCullMargin(margin float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnSetExtraCullMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GeometryInstance3D) GetExtraCullMargin() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnGetExtraCullMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GeometryInstance3D) SetLightmapScale(scale GeometryInstance3DLightmapScale) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnSetLightmapScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GeometryInstance3D) GetLightmapScale() GeometryInstance3DLightmapScale {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret GeometryInstance3DLightmapScale

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnGetLightmapScale), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *GeometryInstance3D) SetGiMode(mode GeometryInstance3DGIMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnSetGiMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GeometryInstance3D) GetGiMode() GeometryInstance3DGIMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret GeometryInstance3DGIMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnGetGiMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *GeometryInstance3D) SetIgnoreOcclusionCulling(ignore_culling bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ignore_culling)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnSetIgnoreOcclusionCulling), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GeometryInstance3D) IsIgnoringOcclusionCulling() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnIsIgnoringOcclusionCulling), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GeometryInstance3D) SetCustomAabb(aabb AABB) {
	cargs := []gdc.ConstTypePtr{aabb.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnSetCustomAabb), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GeometryInstance3D) GetCustomAabb() AABB {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewAABB()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeometryInstance3D.fnGetCustomAabb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
