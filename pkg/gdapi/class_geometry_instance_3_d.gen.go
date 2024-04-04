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
  GeometryInstance3DShadowCastingSettingShadowCastingSettingOff GeometryInstance3DShadowCastingSetting = 0
  GeometryInstance3DShadowCastingSettingShadowCastingSettingOn GeometryInstance3DShadowCastingSetting = 1
  GeometryInstance3DShadowCastingSettingShadowCastingSettingDoubleSided GeometryInstance3DShadowCastingSetting = 2
  GeometryInstance3DShadowCastingSettingShadowCastingSettingShadowsOnly GeometryInstance3DShadowCastingSetting = 3
)

type GeometryInstance3DGIMode int
const (
  GeometryInstance3DGIModeGiModeDisabled GeometryInstance3DGIMode = 0
  GeometryInstance3DGIModeGiModeStatic GeometryInstance3DGIMode = 1
  GeometryInstance3DGIModeGiModeDynamic GeometryInstance3DGIMode = 2
)

type GeometryInstance3DLightmapScale int
const (
  GeometryInstance3DLightmapScaleLightmapScale1X GeometryInstance3DLightmapScale = 0
  GeometryInstance3DLightmapScaleLightmapScale2X GeometryInstance3DLightmapScale = 1
  GeometryInstance3DLightmapScaleLightmapScale4X GeometryInstance3DLightmapScale = 2
  GeometryInstance3DLightmapScaleLightmapScale8X GeometryInstance3DLightmapScale = 3
  GeometryInstance3DLightmapScaleLightmapScaleMax GeometryInstance3DLightmapScale = 4
)

type GeometryInstance3DVisibilityRangeFadeMode int
const (
  GeometryInstance3DVisibilityRangeFadeModeVisibilityRangeFadeDisabled GeometryInstance3DVisibilityRangeFadeMode = 0
  GeometryInstance3DVisibilityRangeFadeModeVisibilityRangeFadeSelf GeometryInstance3DVisibilityRangeFadeMode = 1
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

func  (me *GeometryInstance3D) SetMaterialOverride(material Material, )  {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_material_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2757459619) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{material.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GeometryInstance3D) GetMaterialOverride() Material {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_material_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 5934680) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMaterial()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GeometryInstance3D) SetMaterialOverlay(material Material, )  {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_material_overlay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2757459619) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{material.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GeometryInstance3D) GetMaterialOverlay() Material {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_material_overlay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 5934680) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMaterial()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GeometryInstance3D) SetCastShadowsSetting(shadow_casting_setting GeometryInstance3DShadowCastingSetting, )  {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cast_shadows_setting")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 856677339) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shadow_casting_setting) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GeometryInstance3D) GetCastShadowsSetting() GeometryInstance3DShadowCastingSetting {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cast_shadows_setting")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3383019359) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret GeometryInstance3DShadowCastingSetting

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *GeometryInstance3D) SetLodBias(bias float64, )  {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_lod_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bias) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GeometryInstance3D) GetLodBias() float64 {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_lod_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GeometryInstance3D) SetTransparency(transparency float64, )  {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_transparency")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&transparency) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GeometryInstance3D) GetTransparency() float64 {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_transparency")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GeometryInstance3D) SetVisibilityRangeEndMargin(distance float64, )  {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visibility_range_end_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GeometryInstance3D) GetVisibilityRangeEndMargin() float64 {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visibility_range_end_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GeometryInstance3D) SetVisibilityRangeEnd(distance float64, )  {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visibility_range_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GeometryInstance3D) GetVisibilityRangeEnd() float64 {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visibility_range_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GeometryInstance3D) SetVisibilityRangeBeginMargin(distance float64, )  {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visibility_range_begin_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GeometryInstance3D) GetVisibilityRangeBeginMargin() float64 {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visibility_range_begin_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GeometryInstance3D) SetVisibilityRangeBegin(distance float64, )  {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visibility_range_begin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GeometryInstance3D) GetVisibilityRangeBegin() float64 {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visibility_range_begin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GeometryInstance3D) SetVisibilityRangeFadeMode(mode GeometryInstance3DVisibilityRangeFadeMode, )  {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visibility_range_fade_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1440117808) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GeometryInstance3D) GetVisibilityRangeFadeMode() GeometryInstance3DVisibilityRangeFadeMode {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visibility_range_fade_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2067221882) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret GeometryInstance3DVisibilityRangeFadeMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *GeometryInstance3D) SetInstanceShaderParameter(name StringName, value Variant, )  {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_instance_shader_parameter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3776071444) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GeometryInstance3D) GetInstanceShaderParameter(name StringName, ) Variant {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_instance_shader_parameter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2760726917) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GeometryInstance3D) SetExtraCullMargin(margin float64, )  {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_extra_cull_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GeometryInstance3D) GetExtraCullMargin() float64 {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_extra_cull_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GeometryInstance3D) SetLightmapScale(scale GeometryInstance3DLightmapScale, )  {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_lightmap_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2462696582) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GeometryInstance3D) GetLightmapScale() GeometryInstance3DLightmapScale {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_lightmap_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 798767852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret GeometryInstance3DLightmapScale

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *GeometryInstance3D) SetGiMode(mode GeometryInstance3DGIMode, )  {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gi_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2548557163) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GeometryInstance3D) GetGiMode() GeometryInstance3DGIMode {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gi_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2188566509) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret GeometryInstance3DGIMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *GeometryInstance3D) SetIgnoreOcclusionCulling(ignore_culling bool, )  {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ignore_occlusion_culling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ignore_culling) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GeometryInstance3D) IsIgnoringOcclusionCulling() bool {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_ignoring_occlusion_culling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GeometryInstance3D) SetCustomAabb(aabb AABB, )  {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_aabb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 259215842) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{aabb.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GeometryInstance3D) GetCustomAabb() AABB {
  classNameV := StringNameFromStr("GeometryInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_aabb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1068685055) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAABB()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
