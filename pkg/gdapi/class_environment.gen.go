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

type Environment struct {
  Resource
}

func (me *Environment) BaseClass() string {
  return "Environment"
}

func NewEnvironment() *Environment {
  str := StringNameFromStr("Environment") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Environment{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetBackground() EnvironmentBGMode {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_background")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1843210413) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret EnvironmentBGMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Environment) SetSky(sky Sky, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sky")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3336722921) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{sky.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSky() Sky {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sky")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1177136966) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewSky()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Environment) SetSkyCustomFov(scale float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sky_custom_fov")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSkyCustomFov() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sky_custom_fov")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSkyRotation(euler_radians Vector3, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sky_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{euler_radians.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSkyRotation() Vector3 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sky_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Environment) SetBgColor(color Color, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bg_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetBgColor() Color {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bg_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Environment) SetBgEnergyMultiplier(energy float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bg_energy_multiplier")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&energy) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetBgEnergyMultiplier() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bg_energy_multiplier")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetBgIntensity(energy float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bg_intensity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&energy) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetBgIntensity() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bg_intensity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetCanvasMaxLayer(layer int64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_canvas_max_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetCanvasMaxLayer() int64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_canvas_max_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetCameraFeedId(id int64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_camera_feed_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetCameraFeedId() int64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_camera_feed_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetAmbientLightColor(color Color, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ambient_light_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetAmbientLightColor() Color {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ambient_light_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Environment) SetAmbientSource(source EnvironmentAmbientSource, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ambient_source")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2607780160) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetAmbientSource() EnvironmentAmbientSource {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ambient_source")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 67453933) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret EnvironmentAmbientSource

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Environment) SetAmbientLightEnergy(energy float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ambient_light_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&energy) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetAmbientLightEnergy() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ambient_light_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetAmbientLightSkyContribution(ratio float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ambient_light_sky_contribution")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetAmbientLightSkyContribution() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ambient_light_sky_contribution")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetReflectionSource(source EnvironmentReflectionSource, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_reflection_source")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 299673197) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetReflectionSource() EnvironmentReflectionSource {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_reflection_source")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 777700713) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret EnvironmentReflectionSource

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Environment) SetTonemapper(mode EnvironmentToneMapper, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tonemapper")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1509116664) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetTonemapper() EnvironmentToneMapper {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tonemapper")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2908408137) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret EnvironmentToneMapper

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Environment) SetTonemapExposure(exposure float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tonemap_exposure")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exposure) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetTonemapExposure() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tonemap_exposure")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetTonemapWhite(white float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tonemap_white")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&white) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetTonemapWhite() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tonemap_white")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsrEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssr_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) IsSsrEnabled() bool {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_ssr_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsrMaxSteps(max_steps int64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssr_max_steps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_steps) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsrMaxSteps() int64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssr_max_steps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsrFadeIn(fade_in float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssr_fade_in")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fade_in) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsrFadeIn() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssr_fade_in")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsrFadeOut(fade_out float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssr_fade_out")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fade_out) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsrFadeOut() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssr_fade_out")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsrDepthTolerance(depth_tolerance float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssr_depth_tolerance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&depth_tolerance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsrDepthTolerance() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssr_depth_tolerance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsaoEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssao_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) IsSsaoEnabled() bool {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_ssao_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsaoRadius(radius float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssao_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsaoRadius() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssao_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsaoIntensity(intensity float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssao_intensity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&intensity) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsaoIntensity() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssao_intensity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsaoPower(power float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssao_power")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&power) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsaoPower() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssao_power")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsaoDetail(detail float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssao_detail")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&detail) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsaoDetail() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssao_detail")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsaoHorizon(horizon float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssao_horizon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&horizon) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsaoHorizon() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssao_horizon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsaoSharpness(sharpness float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssao_sharpness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sharpness) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsaoSharpness() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssao_sharpness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsaoDirectLightAffect(amount float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssao_direct_light_affect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsaoDirectLightAffect() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssao_direct_light_affect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsaoAoChannelAffect(amount float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssao_ao_channel_affect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsaoAoChannelAffect() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssao_ao_channel_affect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsilEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssil_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) IsSsilEnabled() bool {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_ssil_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsilRadius(radius float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssil_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsilRadius() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssil_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsilIntensity(intensity float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssil_intensity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&intensity) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsilIntensity() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssil_intensity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsilSharpness(sharpness float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssil_sharpness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sharpness) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsilSharpness() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssil_sharpness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSsilNormalRejection(normal_rejection float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ssil_normal_rejection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&normal_rejection) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSsilNormalRejection() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ssil_normal_rejection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSdfgiEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sdfgi_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) IsSdfgiEnabled() bool {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_sdfgi_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSdfgiCascades(amount int64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sdfgi_cascades")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSdfgiCascades() int64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sdfgi_cascades")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSdfgiMinCellSize(size float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sdfgi_min_cell_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSdfgiMinCellSize() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sdfgi_min_cell_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSdfgiMaxDistance(distance float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sdfgi_max_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSdfgiMaxDistance() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sdfgi_max_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSdfgiCascade0Distance(distance float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sdfgi_cascade0_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSdfgiCascade0Distance() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sdfgi_cascade0_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSdfgiYScale(scale EnvironmentSDFGIYScale, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sdfgi_y_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3608608372) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSdfgiYScale() EnvironmentSDFGIYScale {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sdfgi_y_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2568002245) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret EnvironmentSDFGIYScale

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Environment) SetSdfgiUseOcclusion(enable bool, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sdfgi_use_occlusion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) IsSdfgiUsingOcclusion() bool {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_sdfgi_using_occlusion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSdfgiBounceFeedback(amount float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sdfgi_bounce_feedback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSdfgiBounceFeedback() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sdfgi_bounce_feedback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSdfgiReadSkyLight(enable bool, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sdfgi_read_sky_light")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) IsSdfgiReadingSkyLight() bool {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_sdfgi_reading_sky_light")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSdfgiEnergy(amount float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sdfgi_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSdfgiEnergy() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sdfgi_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSdfgiNormalBias(bias float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sdfgi_normal_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bias) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSdfgiNormalBias() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sdfgi_normal_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetSdfgiProbeBias(bias float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sdfgi_probe_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bias) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetSdfgiProbeBias() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sdfgi_probe_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetGlowEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glow_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) IsGlowEnabled() bool {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_glow_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetGlowLevel(idx int64, intensity float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glow_level")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , gdc.ConstTypePtr(&intensity) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetGlowLevel(idx int64, ) float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glow_level")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetGlowNormalized(normalize bool, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glow_normalized")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&normalize) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) IsGlowNormalized() bool {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_glow_normalized")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetGlowIntensity(intensity float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glow_intensity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&intensity) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetGlowIntensity() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glow_intensity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetGlowStrength(strength float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glow_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&strength) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetGlowStrength() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glow_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetGlowMix(mix float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glow_mix")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mix) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetGlowMix() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glow_mix")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetGlowBloom(amount float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glow_bloom")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetGlowBloom() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glow_bloom")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetGlowBlendMode(mode EnvironmentGlowBlendMode, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glow_blend_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2561587761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetGlowBlendMode() EnvironmentGlowBlendMode {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glow_blend_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1529667332) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret EnvironmentGlowBlendMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Environment) SetGlowHdrBleedThreshold(threshold float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glow_hdr_bleed_threshold")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&threshold) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetGlowHdrBleedThreshold() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glow_hdr_bleed_threshold")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetGlowHdrBleedScale(scale float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glow_hdr_bleed_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetGlowHdrBleedScale() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glow_hdr_bleed_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetGlowHdrLuminanceCap(amount float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glow_hdr_luminance_cap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetGlowHdrLuminanceCap() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glow_hdr_luminance_cap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetGlowMapStrength(strength float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glow_map_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&strength) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetGlowMapStrength() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glow_map_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetGlowMap(mode Texture, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glow_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1790811099) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{mode.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetGlowMap() Texture {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glow_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4037048985) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Environment) SetFogEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fog_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) IsFogEnabled() bool {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_fog_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetFogLightColor(light_color Color, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fog_light_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{light_color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetFogLightColor() Color {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fog_light_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Environment) SetFogLightEnergy(light_energy float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fog_light_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&light_energy) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetFogLightEnergy() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fog_light_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetFogSunScatter(sun_scatter float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fog_sun_scatter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sun_scatter) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetFogSunScatter() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fog_sun_scatter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetFogDensity(density float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fog_density")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&density) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetFogDensity() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fog_density")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetFogHeight(height float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fog_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetFogHeight() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fog_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetFogHeightDensity(height_density float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fog_height_density")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height_density) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetFogHeightDensity() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fog_height_density")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetFogAerialPerspective(aerial_perspective float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fog_aerial_perspective")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&aerial_perspective) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetFogAerialPerspective() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fog_aerial_perspective")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetFogSkyAffect(sky_affect float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fog_sky_affect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sky_affect) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetFogSkyAffect() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fog_sky_affect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetVolumetricFogEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volumetric_fog_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) IsVolumetricFogEnabled() bool {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_volumetric_fog_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetVolumetricFogEmission(color Color, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volumetric_fog_emission")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetVolumetricFogEmission() Color {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_volumetric_fog_emission")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Environment) SetVolumetricFogAlbedo(color Color, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volumetric_fog_albedo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetVolumetricFogAlbedo() Color {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_volumetric_fog_albedo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Environment) SetVolumetricFogDensity(density float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volumetric_fog_density")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&density) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetVolumetricFogDensity() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_volumetric_fog_density")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetVolumetricFogEmissionEnergy(begin float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volumetric_fog_emission_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&begin) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetVolumetricFogEmissionEnergy() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_volumetric_fog_emission_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetVolumetricFogAnisotropy(anisotropy float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volumetric_fog_anisotropy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&anisotropy) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetVolumetricFogAnisotropy() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_volumetric_fog_anisotropy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetVolumetricFogLength(length float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volumetric_fog_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetVolumetricFogLength() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_volumetric_fog_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetVolumetricFogDetailSpread(detail_spread float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volumetric_fog_detail_spread")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&detail_spread) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetVolumetricFogDetailSpread() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_volumetric_fog_detail_spread")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetVolumetricFogGiInject(gi_inject float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volumetric_fog_gi_inject")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gi_inject) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetVolumetricFogGiInject() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_volumetric_fog_gi_inject")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetVolumetricFogAmbientInject(enabled float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volumetric_fog_ambient_inject")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetVolumetricFogAmbientInject() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_volumetric_fog_ambient_inject")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetVolumetricFogSkyAffect(sky_affect float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volumetric_fog_sky_affect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sky_affect) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetVolumetricFogSkyAffect() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_volumetric_fog_sky_affect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetVolumetricFogTemporalReprojectionEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volumetric_fog_temporal_reprojection_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) IsVolumetricFogTemporalReprojectionEnabled() bool {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_volumetric_fog_temporal_reprojection_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetVolumetricFogTemporalReprojectionAmount(temporal_reprojection_amount float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volumetric_fog_temporal_reprojection_amount")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&temporal_reprojection_amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetVolumetricFogTemporalReprojectionAmount() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_volumetric_fog_temporal_reprojection_amount")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetAdjustmentEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_adjustment_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) IsAdjustmentEnabled() bool {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_adjustment_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetAdjustmentBrightness(brightness float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_adjustment_brightness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&brightness) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetAdjustmentBrightness() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_adjustment_brightness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetAdjustmentContrast(contrast float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_adjustment_contrast")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contrast) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetAdjustmentContrast() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_adjustment_contrast")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetAdjustmentSaturation(saturation float64, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_adjustment_saturation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&saturation) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetAdjustmentSaturation() float64 {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_adjustment_saturation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Environment) SetAdjustmentColorCorrection(color_correction Texture, )  {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_adjustment_color_correction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1790811099) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{color_correction.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Environment) GetAdjustmentColorCorrection() Texture {
  classNameV := StringNameFromStr("Environment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_adjustment_color_correction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4037048985) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
