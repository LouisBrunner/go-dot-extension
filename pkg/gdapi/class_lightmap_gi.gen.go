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

func  (me *LightmapGI) SetLightData(data LightmapGIData, ) { // TODO: return value
  // TODO: implement
}

func  (me *LightmapGI) GetLightData() { // TODO: return value
  // TODO: implement
}

func  (me *LightmapGI) SetBakeQuality(bake_quality LightmapGIBakeQuality, ) { // TODO: return value
  // TODO: implement
}

func  (me *LightmapGI) GetBakeQuality() { // TODO: return value
  // TODO: implement
}

func  (me *LightmapGI) SetBounces(bounces int, ) { // TODO: return value
  // TODO: implement
}

func  (me *LightmapGI) GetBounces() { // TODO: return value
  // TODO: implement
}

func  (me *LightmapGI) SetGenerateProbes(subdivision LightmapGIGenerateProbes, ) { // TODO: return value
  // TODO: implement
}

func  (me *LightmapGI) GetGenerateProbes() { // TODO: return value
  // TODO: implement
}

func  (me *LightmapGI) SetBias(bias float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *LightmapGI) GetBias() { // TODO: return value
  // TODO: implement
}

func  (me *LightmapGI) SetEnvironmentMode(mode LightmapGIEnvironmentMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *LightmapGI) GetEnvironmentMode() { // TODO: return value
  // TODO: implement
}

func  (me *LightmapGI) SetEnvironmentCustomSky(sky Sky, ) { // TODO: return value
  // TODO: implement
}

func  (me *LightmapGI) GetEnvironmentCustomSky() { // TODO: return value
  // TODO: implement
}

func  (me *LightmapGI) SetEnvironmentCustomColor(color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *LightmapGI) GetEnvironmentCustomColor() { // TODO: return value
  // TODO: implement
}

func  (me *LightmapGI) SetEnvironmentCustomEnergy(energy float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *LightmapGI) GetEnvironmentCustomEnergy() { // TODO: return value
  // TODO: implement
}

func  (me *LightmapGI) SetMaxTextureSize(max_texture_size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *LightmapGI) GetMaxTextureSize() { // TODO: return value
  // TODO: implement
}

func  (me *LightmapGI) SetUseDenoiser(use_denoiser bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *LightmapGI) IsUsingDenoiser() { // TODO: return value
  // TODO: implement
}

func  (me *LightmapGI) SetInterior(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *LightmapGI) IsInterior() { // TODO: return value
  // TODO: implement
}

func  (me *LightmapGI) SetDirectional(directional bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *LightmapGI) IsDirectional() { // TODO: return value
  // TODO: implement
}

func  (me *LightmapGI) SetCameraAttributes(camera_attributes CameraAttributes, ) { // TODO: return value
  // TODO: implement
}

func  (me *LightmapGI) GetCameraAttributes() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
