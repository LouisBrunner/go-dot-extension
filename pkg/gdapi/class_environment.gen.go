// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Environment struct {
  obj gdc.ObjectPtr
}

func (me *Environment) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Environment) BaseClass() string {
  return "Environment"
}



// Enums

type EnvironmentBGMode int
const (
  EnvironmentBGModeBgClearColor EnvironmentBGMode = 0
  EnvironmentBGModeBgColor EnvironmentBGMode = 1
  EnvironmentBGModeBgSky EnvironmentBGMode = 2
  EnvironmentBGModeBgCanvas EnvironmentBGMode = 3
  EnvironmentBGModeBgKeep EnvironmentBGMode = 4
  EnvironmentBGModeBgCameraFeed EnvironmentBGMode = 5
  EnvironmentBGModeBgMax EnvironmentBGMode = 6
)

type EnvironmentAmbientSource int
const (
  EnvironmentAmbientSourceAmbientSourceBg EnvironmentAmbientSource = 0
  EnvironmentAmbientSourceAmbientSourceDisabled EnvironmentAmbientSource = 1
  EnvironmentAmbientSourceAmbientSourceColor EnvironmentAmbientSource = 2
  EnvironmentAmbientSourceAmbientSourceSky EnvironmentAmbientSource = 3
)

type EnvironmentReflectionSource int
const (
  EnvironmentReflectionSourceReflectionSourceBg EnvironmentReflectionSource = 0
  EnvironmentReflectionSourceReflectionSourceDisabled EnvironmentReflectionSource = 1
  EnvironmentReflectionSourceReflectionSourceSky EnvironmentReflectionSource = 2
)

type EnvironmentToneMapper int
const (
  EnvironmentToneMapperToneMapperLinear EnvironmentToneMapper = 0
  EnvironmentToneMapperToneMapperReinhardt EnvironmentToneMapper = 1
  EnvironmentToneMapperToneMapperFilmic EnvironmentToneMapper = 2
  EnvironmentToneMapperToneMapperAces EnvironmentToneMapper = 3
)

type EnvironmentGlowBlendMode int
const (
  EnvironmentGlowBlendModeGlowBlendModeAdditive EnvironmentGlowBlendMode = 0
  EnvironmentGlowBlendModeGlowBlendModeScreen EnvironmentGlowBlendMode = 1
  EnvironmentGlowBlendModeGlowBlendModeSoftlight EnvironmentGlowBlendMode = 2
  EnvironmentGlowBlendModeGlowBlendModeReplace EnvironmentGlowBlendMode = 3
  EnvironmentGlowBlendModeGlowBlendModeMix EnvironmentGlowBlendMode = 4
)

type EnvironmentSDFGIYScale int
const (
  EnvironmentSDFGIYScaleSdfgiYScale50Percent EnvironmentSDFGIYScale = 0
  EnvironmentSDFGIYScaleSdfgiYScale75Percent EnvironmentSDFGIYScale = 1
  EnvironmentSDFGIYScaleSdfgiYScale100Percent EnvironmentSDFGIYScale = 2
)

func (me *Environment) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Environment) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Environment) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Environment) SetBackground(mode EnvironmentBGMode, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_background")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4071623990) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetBackground() EnvironmentBGMode {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_background")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1843210413) // FIXME: should cache?
  var ret EnvironmentBGMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSky(sky Sky, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sky")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3336722921) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(sky.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSky() Sky {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sky")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1177136966) // FIXME: should cache?
  var ret Sky
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSkyCustomFov(scale float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sky_custom_fov")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSkyCustomFov() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sky_custom_fov")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSkyRotation(euler_radians Vector3, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sky_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(euler_radians.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSkyRotation() Vector3 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sky_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetBgColor(color Color, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bg_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetBgColor() Color {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bg_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetBgEnergyMultiplier(energy float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bg_energy_multiplier")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&energy), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetBgEnergyMultiplier() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bg_energy_multiplier")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetBgIntensity(energy float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bg_intensity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&energy), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetBgIntensity() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bg_intensity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetCanvasMaxLayer(layer int, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_canvas_max_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetCanvasMaxLayer() int {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_canvas_max_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetCameraFeedId(id int, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_camera_feed_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetCameraFeedId() int {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_camera_feed_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetAmbientLightColor(color Color, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ambient_light_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetAmbientLightColor() Color {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ambient_light_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetAmbientSource(source EnvironmentAmbientSource, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ambient_source")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2607780160) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetAmbientSource() EnvironmentAmbientSource {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ambient_source")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 67453933) // FIXME: should cache?
  var ret EnvironmentAmbientSource
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetAmbientLightEnergy(energy float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ambient_light_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&energy), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetAmbientLightEnergy() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ambient_light_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetAmbientLightSkyContribution(ratio float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ambient_light_sky_contribution")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetAmbientLightSkyContribution() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ambient_light_sky_contribution")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetReflectionSource(source EnvironmentReflectionSource, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_reflection_source")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 299673197) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetReflectionSource() EnvironmentReflectionSource {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_reflection_source")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 777700713) // FIXME: should cache?
  var ret EnvironmentReflectionSource
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetTonemapper(mode EnvironmentToneMapper, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tonemapper")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1509116664) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetTonemapper() EnvironmentToneMapper {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tonemapper")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2908408137) // FIXME: should cache?
  var ret EnvironmentToneMapper
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetTonemapExposure(exposure float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tonemap_exposure")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exposure), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetTonemapExposure() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tonemap_exposure")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetTonemapWhite(white float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tonemap_white")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&white), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetTonemapWhite() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tonemap_white")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSsrEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssr_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) IsSsrEnabled() bool {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_ssr_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSsrMaxSteps(max_steps int, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssr_max_steps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_steps), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSsrMaxSteps() int {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssr_max_steps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSsrFadeIn(fade_in float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssr_fade_in")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fade_in), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSsrFadeIn() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssr_fade_in")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSsrFadeOut(fade_out float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssr_fade_out")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fade_out), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSsrFadeOut() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssr_fade_out")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSsrDepthTolerance(depth_tolerance float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssr_depth_tolerance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&depth_tolerance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSsrDepthTolerance() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssr_depth_tolerance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSsaoEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssao_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) IsSsaoEnabled() bool {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_ssao_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSsaoRadius(radius float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssao_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSsaoRadius() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssao_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSsaoIntensity(intensity float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssao_intensity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&intensity), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSsaoIntensity() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssao_intensity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSsaoPower(power float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssao_power")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&power), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSsaoPower() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssao_power")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSsaoDetail(detail float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssao_detail")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&detail), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSsaoDetail() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssao_detail")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSsaoHorizon(horizon float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssao_horizon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&horizon), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSsaoHorizon() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssao_horizon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSsaoSharpness(sharpness float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssao_sharpness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sharpness), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSsaoSharpness() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssao_sharpness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSsaoDirectLightAffect(amount float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssao_direct_light_affect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSsaoDirectLightAffect() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssao_direct_light_affect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSsaoAoChannelAffect(amount float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssao_ao_channel_affect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSsaoAoChannelAffect() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssao_ao_channel_affect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSsilEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssil_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) IsSsilEnabled() bool {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_ssil_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSsilRadius(radius float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssil_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSsilRadius() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssil_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSsilIntensity(intensity float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssil_intensity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&intensity), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSsilIntensity() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssil_intensity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSsilSharpness(sharpness float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssil_sharpness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sharpness), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSsilSharpness() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssil_sharpness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSsilNormalRejection(normal_rejection float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssil_normal_rejection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&normal_rejection), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSsilNormalRejection() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssil_normal_rejection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSdfgiEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sdfgi_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) IsSdfgiEnabled() bool {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_sdfgi_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSdfgiCascades(amount int, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sdfgi_cascades")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSdfgiCascades() int {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sdfgi_cascades")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSdfgiMinCellSize(size float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sdfgi_min_cell_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSdfgiMinCellSize() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sdfgi_min_cell_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSdfgiMaxDistance(distance float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sdfgi_max_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSdfgiMaxDistance() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sdfgi_max_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSdfgiCascade0Distance(distance float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sdfgi_cascade0_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSdfgiCascade0Distance() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sdfgi_cascade0_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSdfgiYScale(scale EnvironmentSDFGIYScale, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sdfgi_y_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3608608372) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSdfgiYScale() EnvironmentSDFGIYScale {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sdfgi_y_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2568002245) // FIXME: should cache?
  var ret EnvironmentSDFGIYScale
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSdfgiUseOcclusion(enable bool, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sdfgi_use_occlusion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) IsSdfgiUsingOcclusion() bool {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_sdfgi_using_occlusion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSdfgiBounceFeedback(amount float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sdfgi_bounce_feedback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSdfgiBounceFeedback() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sdfgi_bounce_feedback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSdfgiReadSkyLight(enable bool, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sdfgi_read_sky_light")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) IsSdfgiReadingSkyLight() bool {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_sdfgi_reading_sky_light")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSdfgiEnergy(amount float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sdfgi_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSdfgiEnergy() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sdfgi_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSdfgiNormalBias(bias float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sdfgi_normal_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bias), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSdfgiNormalBias() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sdfgi_normal_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetSdfgiProbeBias(bias float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sdfgi_probe_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bias), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetSdfgiProbeBias() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sdfgi_probe_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetGlowEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glow_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) IsGlowEnabled() bool {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_glow_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetGlowLevel(idx int, intensity float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glow_level")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&intensity), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetGlowLevel(idx int, ) float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glow_level")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetGlowNormalized(normalize bool, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glow_normalized")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&normalize), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) IsGlowNormalized() bool {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_glow_normalized")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetGlowIntensity(intensity float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glow_intensity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&intensity), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetGlowIntensity() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glow_intensity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetGlowStrength(strength float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glow_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&strength), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetGlowStrength() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glow_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetGlowMix(mix float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glow_mix")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mix), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetGlowMix() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glow_mix")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetGlowBloom(amount float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glow_bloom")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetGlowBloom() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glow_bloom")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetGlowBlendMode(mode EnvironmentGlowBlendMode, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glow_blend_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2561587761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetGlowBlendMode() EnvironmentGlowBlendMode {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glow_blend_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1529667332) // FIXME: should cache?
  var ret EnvironmentGlowBlendMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetGlowHdrBleedThreshold(threshold float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glow_hdr_bleed_threshold")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&threshold), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetGlowHdrBleedThreshold() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glow_hdr_bleed_threshold")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetGlowHdrBleedScale(scale float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glow_hdr_bleed_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetGlowHdrBleedScale() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glow_hdr_bleed_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetGlowHdrLuminanceCap(amount float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glow_hdr_luminance_cap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetGlowHdrLuminanceCap() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glow_hdr_luminance_cap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetGlowMapStrength(strength float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glow_map_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&strength), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetGlowMapStrength() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glow_map_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetGlowMap(mode Texture, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glow_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1790811099) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(mode.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetGlowMap() Texture {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glow_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4037048985) // FIXME: should cache?
  var ret Texture
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetFogEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fog_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) IsFogEnabled() bool {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_fog_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetFogLightColor(light_color Color, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fog_light_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(light_color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetFogLightColor() Color {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fog_light_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetFogLightEnergy(light_energy float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fog_light_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&light_energy), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetFogLightEnergy() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fog_light_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetFogSunScatter(sun_scatter float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fog_sun_scatter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sun_scatter), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetFogSunScatter() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fog_sun_scatter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetFogDensity(density float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fog_density")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&density), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetFogDensity() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fog_density")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetFogHeight(height float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fog_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetFogHeight() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fog_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetFogHeightDensity(height_density float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fog_height_density")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height_density), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetFogHeightDensity() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fog_height_density")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetFogAerialPerspective(aerial_perspective float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fog_aerial_perspective")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&aerial_perspective), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetFogAerialPerspective() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fog_aerial_perspective")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetFogSkyAffect(sky_affect float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fog_sky_affect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sky_affect), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetFogSkyAffect() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fog_sky_affect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetVolumetricFogEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volumetric_fog_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) IsVolumetricFogEnabled() bool {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_volumetric_fog_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetVolumetricFogEmission(color Color, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volumetric_fog_emission")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetVolumetricFogEmission() Color {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_volumetric_fog_emission")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetVolumetricFogAlbedo(color Color, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volumetric_fog_albedo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetVolumetricFogAlbedo() Color {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_volumetric_fog_albedo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetVolumetricFogDensity(density float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volumetric_fog_density")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&density), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetVolumetricFogDensity() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_volumetric_fog_density")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetVolumetricFogEmissionEnergy(begin float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volumetric_fog_emission_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&begin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetVolumetricFogEmissionEnergy() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_volumetric_fog_emission_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetVolumetricFogAnisotropy(anisotropy float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volumetric_fog_anisotropy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&anisotropy), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetVolumetricFogAnisotropy() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_volumetric_fog_anisotropy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetVolumetricFogLength(length float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volumetric_fog_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetVolumetricFogLength() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_volumetric_fog_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetVolumetricFogDetailSpread(detail_spread float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volumetric_fog_detail_spread")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&detail_spread), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetVolumetricFogDetailSpread() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_volumetric_fog_detail_spread")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetVolumetricFogGiInject(gi_inject float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volumetric_fog_gi_inject")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gi_inject), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetVolumetricFogGiInject() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_volumetric_fog_gi_inject")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetVolumetricFogAmbientInject(enabled float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volumetric_fog_ambient_inject")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetVolumetricFogAmbientInject() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_volumetric_fog_ambient_inject")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetVolumetricFogSkyAffect(sky_affect float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volumetric_fog_sky_affect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sky_affect), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetVolumetricFogSkyAffect() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_volumetric_fog_sky_affect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetVolumetricFogTemporalReprojectionEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volumetric_fog_temporal_reprojection_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) IsVolumetricFogTemporalReprojectionEnabled() bool {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_volumetric_fog_temporal_reprojection_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetVolumetricFogTemporalReprojectionAmount(temporal_reprojection_amount float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volumetric_fog_temporal_reprojection_amount")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&temporal_reprojection_amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetVolumetricFogTemporalReprojectionAmount() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_volumetric_fog_temporal_reprojection_amount")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetAdjustmentEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_adjustment_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) IsAdjustmentEnabled() bool {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_adjustment_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetAdjustmentBrightness(brightness float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_adjustment_brightness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&brightness), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetAdjustmentBrightness() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_adjustment_brightness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetAdjustmentContrast(contrast float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_adjustment_contrast")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contrast), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetAdjustmentContrast() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_adjustment_contrast")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetAdjustmentSaturation(saturation float32, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_adjustment_saturation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&saturation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetAdjustmentSaturation() float32 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_adjustment_saturation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Environment) SetAdjustmentColorCorrection(color_correction Texture, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_adjustment_color_correction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1790811099) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color_correction.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Environment) GetAdjustmentColorCorrection() Texture {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_adjustment_color_correction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4037048985) // FIXME: should cache?
  var ret Texture
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *Environment) GetPropBackgroundMode() int {
  panic("TODO: implement")
}

func (me *Environment) SetPropBackgroundMode(value int) {
  panic("TODO: implement")
}

func (me *Environment) GetPropBackgroundColor() Color {
  panic("TODO: implement")
}

func (me *Environment) SetPropBackgroundColor(value Color) {
  panic("TODO: implement")
}

func (me *Environment) GetPropBackgroundEnergyMultiplier() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropBackgroundEnergyMultiplier(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropBackgroundIntensity() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropBackgroundIntensity(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropBackgroundCanvasMaxLayer() int {
  panic("TODO: implement")
}

func (me *Environment) SetPropBackgroundCanvasMaxLayer(value int) {
  panic("TODO: implement")
}

func (me *Environment) GetPropBackgroundCameraFeedId() int {
  panic("TODO: implement")
}

func (me *Environment) SetPropBackgroundCameraFeedId(value int) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSky() Sky {
  panic("TODO: implement")
}

func (me *Environment) SetPropSky(value Sky) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSkyCustomFov() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropSkyCustomFov(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSkyRotation() Vector3 {
  panic("TODO: implement")
}

func (me *Environment) SetPropSkyRotation(value Vector3) {
  panic("TODO: implement")
}

func (me *Environment) GetPropAmbientLightSource() int {
  panic("TODO: implement")
}

func (me *Environment) SetPropAmbientLightSource(value int) {
  panic("TODO: implement")
}

func (me *Environment) GetPropAmbientLightColor() Color {
  panic("TODO: implement")
}

func (me *Environment) SetPropAmbientLightColor(value Color) {
  panic("TODO: implement")
}

func (me *Environment) GetPropAmbientLightSkyContribution() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropAmbientLightSkyContribution(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropAmbientLightEnergy() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropAmbientLightEnergy(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropReflectedLightSource() int {
  panic("TODO: implement")
}

func (me *Environment) SetPropReflectedLightSource(value int) {
  panic("TODO: implement")
}

func (me *Environment) GetPropTonemapMode() int {
  panic("TODO: implement")
}

func (me *Environment) SetPropTonemapMode(value int) {
  panic("TODO: implement")
}

func (me *Environment) GetPropTonemapExposure() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropTonemapExposure(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropTonemapWhite() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropTonemapWhite(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSsrEnabled() bool {
  panic("TODO: implement")
}

func (me *Environment) SetPropSsrEnabled(value bool) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSsrMaxSteps() int {
  panic("TODO: implement")
}

func (me *Environment) SetPropSsrMaxSteps(value int) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSsrFadeIn() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropSsrFadeIn(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSsrFadeOut() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropSsrFadeOut(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSsrDepthTolerance() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropSsrDepthTolerance(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSsaoEnabled() bool {
  panic("TODO: implement")
}

func (me *Environment) SetPropSsaoEnabled(value bool) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSsaoRadius() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropSsaoRadius(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSsaoIntensity() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropSsaoIntensity(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSsaoPower() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropSsaoPower(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSsaoDetail() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropSsaoDetail(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSsaoHorizon() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropSsaoHorizon(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSsaoSharpness() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropSsaoSharpness(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSsaoLightAffect() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropSsaoLightAffect(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSsaoAoChannelAffect() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropSsaoAoChannelAffect(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSsilEnabled() bool {
  panic("TODO: implement")
}

func (me *Environment) SetPropSsilEnabled(value bool) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSsilRadius() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropSsilRadius(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSsilIntensity() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropSsilIntensity(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSsilSharpness() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropSsilSharpness(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSsilNormalRejection() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropSsilNormalRejection(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSdfgiEnabled() bool {
  panic("TODO: implement")
}

func (me *Environment) SetPropSdfgiEnabled(value bool) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSdfgiUseOcclusion() bool {
  panic("TODO: implement")
}

func (me *Environment) SetPropSdfgiUseOcclusion(value bool) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSdfgiReadSkyLight() bool {
  panic("TODO: implement")
}

func (me *Environment) SetPropSdfgiReadSkyLight(value bool) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSdfgiBounceFeedback() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropSdfgiBounceFeedback(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSdfgiCascades() int {
  panic("TODO: implement")
}

func (me *Environment) SetPropSdfgiCascades(value int) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSdfgiMinCellSize() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropSdfgiMinCellSize(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSdfgiCascade0Distance() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropSdfgiCascade0Distance(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSdfgiMaxDistance() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropSdfgiMaxDistance(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSdfgiYScale() int {
  panic("TODO: implement")
}

func (me *Environment) SetPropSdfgiYScale(value int) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSdfgiEnergy() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropSdfgiEnergy(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSdfgiNormalBias() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropSdfgiNormalBias(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropSdfgiProbeBias() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropSdfgiProbeBias(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropGlowEnabled() bool {
  panic("TODO: implement")
}

func (me *Environment) SetPropGlowEnabled(value bool) {
  panic("TODO: implement")
}

func (me *Environment) GetPropGlowNormalized() bool {
  panic("TODO: implement")
}

func (me *Environment) SetPropGlowNormalized(value bool) {
  panic("TODO: implement")
}

func (me *Environment) GetPropGlowIntensity() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropGlowIntensity(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropGlowStrength() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropGlowStrength(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropGlowMix() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropGlowMix(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropGlowBloom() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropGlowBloom(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropGlowBlendMode() int {
  panic("TODO: implement")
}

func (me *Environment) SetPropGlowBlendMode(value int) {
  panic("TODO: implement")
}

func (me *Environment) GetPropGlowHdrThreshold() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropGlowHdrThreshold(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropGlowHdrScale() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropGlowHdrScale(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropGlowHdrLuminanceCap() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropGlowHdrLuminanceCap(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropGlowMapStrength() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropGlowMapStrength(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropGlowMap() Texture2D {
  panic("TODO: implement")
}

func (me *Environment) SetPropGlowMap(value Texture2D) {
  panic("TODO: implement")
}

func (me *Environment) GetPropFogEnabled() bool {
  panic("TODO: implement")
}

func (me *Environment) SetPropFogEnabled(value bool) {
  panic("TODO: implement")
}

func (me *Environment) GetPropFogLightColor() Color {
  panic("TODO: implement")
}

func (me *Environment) SetPropFogLightColor(value Color) {
  panic("TODO: implement")
}

func (me *Environment) GetPropFogLightEnergy() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropFogLightEnergy(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropFogSunScatter() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropFogSunScatter(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropFogDensity() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropFogDensity(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropFogAerialPerspective() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropFogAerialPerspective(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropFogSkyAffect() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropFogSkyAffect(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropFogHeight() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropFogHeight(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropFogHeightDensity() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropFogHeightDensity(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropVolumetricFogEnabled() bool {
  panic("TODO: implement")
}

func (me *Environment) SetPropVolumetricFogEnabled(value bool) {
  panic("TODO: implement")
}

func (me *Environment) GetPropVolumetricFogDensity() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropVolumetricFogDensity(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropVolumetricFogAlbedo() Color {
  panic("TODO: implement")
}

func (me *Environment) SetPropVolumetricFogAlbedo(value Color) {
  panic("TODO: implement")
}

func (me *Environment) GetPropVolumetricFogEmission() Color {
  panic("TODO: implement")
}

func (me *Environment) SetPropVolumetricFogEmission(value Color) {
  panic("TODO: implement")
}

func (me *Environment) GetPropVolumetricFogEmissionEnergy() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropVolumetricFogEmissionEnergy(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropVolumetricFogGiInject() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropVolumetricFogGiInject(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropVolumetricFogAnisotropy() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropVolumetricFogAnisotropy(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropVolumetricFogLength() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropVolumetricFogLength(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropVolumetricFogDetailSpread() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropVolumetricFogDetailSpread(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropVolumetricFogAmbientInject() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropVolumetricFogAmbientInject(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropVolumetricFogSkyAffect() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropVolumetricFogSkyAffect(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropVolumetricFogTemporalReprojectionEnabled() bool {
  panic("TODO: implement")
}

func (me *Environment) SetPropVolumetricFogTemporalReprojectionEnabled(value bool) {
  panic("TODO: implement")
}

func (me *Environment) GetPropVolumetricFogTemporalReprojectionAmount() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropVolumetricFogTemporalReprojectionAmount(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropAdjustmentEnabled() bool {
  panic("TODO: implement")
}

func (me *Environment) SetPropAdjustmentEnabled(value bool) {
  panic("TODO: implement")
}

func (me *Environment) GetPropAdjustmentBrightness() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropAdjustmentBrightness(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropAdjustmentContrast() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropAdjustmentContrast(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropAdjustmentSaturation() float32 {
  panic("TODO: implement")
}

func (me *Environment) SetPropAdjustmentSaturation(value float32) {
  panic("TODO: implement")
}

func (me *Environment) GetPropAdjustmentColorCorrection() any {
  panic("TODO: implement")
}

func (me *Environment) SetPropAdjustmentColorCorrection(value any) {
  panic("TODO: implement")
}