// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Light3D struct {
  obj gdc.ObjectPtr
}

func (me *Light3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Light3D) BaseClass() string {
  return "Light3D"
}

type Light3DParam int
const (
  Light3DParamParamEnergy Light3DParam = 0
  Light3DParamParamIndirectEnergy Light3DParam = 1
  Light3DParamParamVolumetricFogEnergy Light3DParam = 2
  Light3DParamParamSpecular Light3DParam = 3
  Light3DParamParamRange Light3DParam = 4
  Light3DParamParamSize Light3DParam = 5
  Light3DParamParamAttenuation Light3DParam = 6
  Light3DParamParamSpotAngle Light3DParam = 7
  Light3DParamParamSpotAttenuation Light3DParam = 8
  Light3DParamParamShadowMaxDistance Light3DParam = 9
  Light3DParamParamShadowSplit1Offset Light3DParam = 10
  Light3DParamParamShadowSplit2Offset Light3DParam = 11
  Light3DParamParamShadowSplit3Offset Light3DParam = 12
  Light3DParamParamShadowFadeStart Light3DParam = 13
  Light3DParamParamShadowNormalBias Light3DParam = 14
  Light3DParamParamShadowBias Light3DParam = 15
  Light3DParamParamShadowPancakeSize Light3DParam = 16
  Light3DParamParamShadowOpacity Light3DParam = 17
  Light3DParamParamShadowBlur Light3DParam = 18
  Light3DParamParamTransmittanceBias Light3DParam = 19
  Light3DParamParamIntensity Light3DParam = 20
  Light3DParamParamMax Light3DParam = 21
)

type Light3DBakeMode int
const (
  Light3DBakeModeBakeDisabled Light3DBakeMode = 0
  Light3DBakeModeBakeStatic Light3DBakeMode = 1
  Light3DBakeModeBakeDynamic Light3DBakeMode = 2
)

func  (me *Light3D) SetEditorOnly(editor_only bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) IsEditorOnly() { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) SetParam(param Light3DParam, value float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) GetParam(param Light3DParam, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) SetShadow(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) HasShadow() { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) SetNegative(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) IsNegative() { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) SetCullMask(cull_mask int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) GetCullMask() { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) SetEnableDistanceFade(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) IsDistanceFadeEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) SetDistanceFadeBegin(distance float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) GetDistanceFadeBegin() { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) SetDistanceFadeShadow(distance float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) GetDistanceFadeShadow() { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) SetDistanceFadeLength(distance float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) GetDistanceFadeLength() { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) SetColor(color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) GetColor() { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) SetShadowReverseCullFace(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) GetShadowReverseCullFace() { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) SetBakeMode(bake_mode Light3DBakeMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) GetBakeMode() { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) SetProjector(projector Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) GetProjector() { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) SetTemperature(temperature float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) GetTemperature() { // TODO: return value
  // TODO: implement
}

func  (me *Light3D) GetCorrelatedColor() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
