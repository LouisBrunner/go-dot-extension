// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type LightmapGI struct {
  obj gdc.ObjectPtr
}

func (me *LightmapGI) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *LightmapGI) BaseClass() string {
  return "LightmapGI"
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
  panic("TODO: implement")
}

func  (me *LightmapGI) GetLightData()  {
  panic("TODO: implement")
}

func  (me *LightmapGI) SetBakeQuality(bake_quality LightmapGIBakeQuality, )  {
  panic("TODO: implement")
}

func  (me *LightmapGI) GetBakeQuality()  {
  panic("TODO: implement")
}

func  (me *LightmapGI) SetBounces(bounces int, )  {
  panic("TODO: implement")
}

func  (me *LightmapGI) GetBounces()  {
  panic("TODO: implement")
}

func  (me *LightmapGI) SetGenerateProbes(subdivision LightmapGIGenerateProbes, )  {
  panic("TODO: implement")
}

func  (me *LightmapGI) GetGenerateProbes()  {
  panic("TODO: implement")
}

func  (me *LightmapGI) SetBias(bias float32, )  {
  panic("TODO: implement")
}

func  (me *LightmapGI) GetBias()  {
  panic("TODO: implement")
}

func  (me *LightmapGI) SetEnvironmentMode(mode LightmapGIEnvironmentMode, )  {
  panic("TODO: implement")
}

func  (me *LightmapGI) GetEnvironmentMode()  {
  panic("TODO: implement")
}

func  (me *LightmapGI) SetEnvironmentCustomSky(sky Sky, )  {
  panic("TODO: implement")
}

func  (me *LightmapGI) GetEnvironmentCustomSky()  {
  panic("TODO: implement")
}

func  (me *LightmapGI) SetEnvironmentCustomColor(color Color, )  {
  panic("TODO: implement")
}

func  (me *LightmapGI) GetEnvironmentCustomColor()  {
  panic("TODO: implement")
}

func  (me *LightmapGI) SetEnvironmentCustomEnergy(energy float32, )  {
  panic("TODO: implement")
}

func  (me *LightmapGI) GetEnvironmentCustomEnergy()  {
  panic("TODO: implement")
}

func  (me *LightmapGI) SetMaxTextureSize(max_texture_size int, )  {
  panic("TODO: implement")
}

func  (me *LightmapGI) GetMaxTextureSize()  {
  panic("TODO: implement")
}

func  (me *LightmapGI) SetUseDenoiser(use_denoiser bool, )  {
  panic("TODO: implement")
}

func  (me *LightmapGI) IsUsingDenoiser()  {
  panic("TODO: implement")
}

func  (me *LightmapGI) SetInterior(enable bool, )  {
  panic("TODO: implement")
}

func  (me *LightmapGI) IsInterior()  {
  panic("TODO: implement")
}

func  (me *LightmapGI) SetDirectional(directional bool, )  {
  panic("TODO: implement")
}

func  (me *LightmapGI) IsDirectional()  {
  panic("TODO: implement")
}

func  (me *LightmapGI) SetCameraAttributes(camera_attributes CameraAttributes, )  {
  panic("TODO: implement")
}

func  (me *LightmapGI) GetCameraAttributes()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
