// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GeometryInstance3D struct {
  obj gdc.ObjectPtr
}

func (me *GeometryInstance3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GeometryInstance3D) BaseClass() string {
  return "GeometryInstance3D"
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

func (me *GeometryInstance3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GeometryInstance3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *GeometryInstance3D) SetMaterialOverride(material Material, )  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) GetMaterialOverride()  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) SetMaterialOverlay(material Material, )  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) GetMaterialOverlay()  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) SetCastShadowsSetting(shadow_casting_setting GeometryInstance3DShadowCastingSetting, )  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) GetCastShadowsSetting()  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) SetLodBias(bias float32, )  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) GetLodBias()  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) SetTransparency(transparency float32, )  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) GetTransparency()  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) SetVisibilityRangeEndMargin(distance float32, )  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) GetVisibilityRangeEndMargin()  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) SetVisibilityRangeEnd(distance float32, )  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) GetVisibilityRangeEnd()  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) SetVisibilityRangeBeginMargin(distance float32, )  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) GetVisibilityRangeBeginMargin()  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) SetVisibilityRangeBegin(distance float32, )  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) GetVisibilityRangeBegin()  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) SetVisibilityRangeFadeMode(mode GeometryInstance3DVisibilityRangeFadeMode, )  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) GetVisibilityRangeFadeMode()  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) SetInstanceShaderParameter(name StringName, value Variant, )  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) GetInstanceShaderParameter(name StringName, )  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) SetExtraCullMargin(margin float32, )  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) GetExtraCullMargin()  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) SetLightmapScale(scale GeometryInstance3DLightmapScale, )  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) GetLightmapScale()  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) SetGiMode(mode GeometryInstance3DGIMode, )  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) GetGiMode()  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) SetIgnoreOcclusionCulling(ignore_culling bool, )  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) IsIgnoringOcclusionCulling()  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) SetCustomAabb(aabb AABB, )  {
  panic("TODO: implement")
}

func  (me *GeometryInstance3D) GetCustomAabb()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
