// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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

type GeometryInstance3DShadowCastingSetting int
const (
  GeometryInstance3DShadowCastingSettingOff GeometryInstance3DShadowCastingSetting = 0
  GeometryInstance3DShadowCastingSettingOn GeometryInstance3DShadowCastingSetting = 1
  GeometryInstance3DShadowCastingSettingDoubleSided GeometryInstance3DShadowCastingSetting = 2
  GeometryInstance3DShadowCastingSettingShadowsOnly GeometryInstance3DShadowCastingSetting = 3
)

type GeometryInstance3DGIMode int
const (
  GeometryInstance3DGiModeDisabled GeometryInstance3DGIMode = 0
  GeometryInstance3DGiModeStatic GeometryInstance3DGIMode = 1
  GeometryInstance3DGiModeDynamic GeometryInstance3DGIMode = 2
)

type GeometryInstance3DLightmapScale int
const (
  GeometryInstance3DLightmapScale1X GeometryInstance3DLightmapScale = 0
  GeometryInstance3DLightmapScale2X GeometryInstance3DLightmapScale = 1
  GeometryInstance3DLightmapScale4X GeometryInstance3DLightmapScale = 2
  GeometryInstance3DLightmapScale8X GeometryInstance3DLightmapScale = 3
  GeometryInstance3DLightmapScaleMax GeometryInstance3DLightmapScale = 4
)

type GeometryInstance3DVisibilityRangeFadeMode int
const (
  GeometryInstance3DVisibilityRangeFadeDisabled GeometryInstance3DVisibilityRangeFadeMode = 0
  GeometryInstance3DVisibilityRangeFadeSelf GeometryInstance3DVisibilityRangeFadeMode = 1
  GeometryInstance3DVisibilityRangeFadeDependencies GeometryInstance3DVisibilityRangeFadeMode = 2
)
