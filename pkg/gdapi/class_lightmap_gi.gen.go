// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  LightmapGIBakeQualityBakeQualityLow LightmapGIBakeQuality = 0
  LightmapGIBakeQualityBakeQualityMedium LightmapGIBakeQuality = 1
  LightmapGIBakeQualityBakeQualityHigh LightmapGIBakeQuality = 2
  LightmapGIBakeQualityBakeQualityUltra LightmapGIBakeQuality = 3
)

type LightmapGIGenerateProbes int
const (
  LightmapGIGenerateProbesGenerateProbesDisabled LightmapGIGenerateProbes = 0
  LightmapGIGenerateProbesGenerateProbesSubdiv4 LightmapGIGenerateProbes = 1
  LightmapGIGenerateProbesGenerateProbesSubdiv8 LightmapGIGenerateProbes = 2
  LightmapGIGenerateProbesGenerateProbesSubdiv16 LightmapGIGenerateProbes = 3
  LightmapGIGenerateProbesGenerateProbesSubdiv32 LightmapGIGenerateProbes = 4
)

type LightmapGIBakeError int
const (
  LightmapGIBakeErrorBakeErrorOk LightmapGIBakeError = 0
  LightmapGIBakeErrorBakeErrorNoSceneRoot LightmapGIBakeError = 1
  LightmapGIBakeErrorBakeErrorForeignData LightmapGIBakeError = 2
  LightmapGIBakeErrorBakeErrorNoLightmapper LightmapGIBakeError = 3
  LightmapGIBakeErrorBakeErrorNoSavePath LightmapGIBakeError = 4
  LightmapGIBakeErrorBakeErrorNoMeshes LightmapGIBakeError = 5
  LightmapGIBakeErrorBakeErrorMeshesInvalid LightmapGIBakeError = 6
  LightmapGIBakeErrorBakeErrorCantCreateImage LightmapGIBakeError = 7
  LightmapGIBakeErrorBakeErrorUserAborted LightmapGIBakeError = 8
  LightmapGIBakeErrorBakeErrorTextureSizeTooSmall LightmapGIBakeError = 9
)

type LightmapGIEnvironmentMode int
const (
  LightmapGIEnvironmentModeEnvironmentModeDisabled LightmapGIEnvironmentMode = 0
  LightmapGIEnvironmentModeEnvironmentModeScene LightmapGIEnvironmentMode = 1
  LightmapGIEnvironmentModeEnvironmentModeCustomSky LightmapGIEnvironmentMode = 2
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

func  (me *LightmapGI) SetLightData(data LightmapGIData, )  {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_light_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1790597277) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(data.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightmapGI) GetLightData() LightmapGIData {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_light_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 290354153) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewLightmapGIData()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *LightmapGI) SetBakeQuality(bake_quality LightmapGIBakeQuality, )  {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bake_quality")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1192215803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bake_quality), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightmapGI) GetBakeQuality() LightmapGIBakeQuality {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bake_quality")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 688832735) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret LightmapGIBakeQuality

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *LightmapGI) SetBounces(bounces int64, )  {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bounces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bounces), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightmapGI) GetBounces() int64 {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bounces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LightmapGI) SetBounceIndirectEnergy(bounce_indirect_energy float64, )  {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bounce_indirect_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bounce_indirect_energy), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightmapGI) GetBounceIndirectEnergy() float64 {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bounce_indirect_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LightmapGI) SetGenerateProbes(subdivision LightmapGIGenerateProbes, )  {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_generate_probes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 549981046) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&subdivision), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightmapGI) GetGenerateProbes() LightmapGIGenerateProbes {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_generate_probes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3930596226) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret LightmapGIGenerateProbes

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *LightmapGI) SetBias(bias float64, )  {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bias), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightmapGI) GetBias() float64 {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LightmapGI) SetEnvironmentMode(mode LightmapGIEnvironmentMode, )  {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_environment_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2282650285) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightmapGI) GetEnvironmentMode() LightmapGIEnvironmentMode {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_environment_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4128646479) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret LightmapGIEnvironmentMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *LightmapGI) SetEnvironmentCustomSky(sky Sky, )  {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_environment_custom_sky")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3336722921) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(sky.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightmapGI) GetEnvironmentCustomSky() Sky {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_environment_custom_sky")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1177136966) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewSky()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *LightmapGI) SetEnvironmentCustomColor(color Color, )  {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_environment_custom_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightmapGI) GetEnvironmentCustomColor() Color {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_environment_custom_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *LightmapGI) SetEnvironmentCustomEnergy(energy float64, )  {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_environment_custom_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&energy), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightmapGI) GetEnvironmentCustomEnergy() float64 {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_environment_custom_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LightmapGI) SetMaxTextureSize(max_texture_size int64, )  {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_texture_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_texture_size), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightmapGI) GetMaxTextureSize() int64 {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_texture_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LightmapGI) SetUseDenoiser(use_denoiser bool, )  {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_denoiser")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_denoiser), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightmapGI) IsUsingDenoiser() bool {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_denoiser")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LightmapGI) SetDenoiserStrength(denoiser_strength float64, )  {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_denoiser_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&denoiser_strength), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightmapGI) GetDenoiserStrength() float64 {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_denoiser_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LightmapGI) SetInterior(enable bool, )  {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_interior")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightmapGI) IsInterior() bool {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_interior")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LightmapGI) SetDirectional(directional bool, )  {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_directional")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&directional), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightmapGI) IsDirectional() bool {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_directional")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LightmapGI) SetUseTextureForBounces(use_texture_for_bounces bool, )  {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_texture_for_bounces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_texture_for_bounces), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightmapGI) IsUsingTextureForBounces() bool {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_texture_for_bounces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LightmapGI) SetCameraAttributes(camera_attributes CameraAttributes, )  {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_camera_attributes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2817810567) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(camera_attributes.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightmapGI) GetCameraAttributes() CameraAttributes {
  classNameV := StringNameFromStr("LightmapGI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_camera_attributes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3921283215) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewCameraAttributes()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
