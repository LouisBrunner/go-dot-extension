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



// Enums

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

func (me *Light3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Light3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Light3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Light3D) SetEditorOnly(editor_only bool, )  {
  panic("TODO: implement")
}

func  (me *Light3D) IsEditorOnly()  {
  panic("TODO: implement")
}

func  (me *Light3D) SetParam(param Light3DParam, value float32, )  {
  panic("TODO: implement")
}

func  (me *Light3D) GetParam(param Light3DParam, )  {
  panic("TODO: implement")
}

func  (me *Light3D) SetShadow(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Light3D) HasShadow()  {
  panic("TODO: implement")
}

func  (me *Light3D) SetNegative(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Light3D) IsNegative()  {
  panic("TODO: implement")
}

func  (me *Light3D) SetCullMask(cull_mask int, )  {
  panic("TODO: implement")
}

func  (me *Light3D) GetCullMask()  {
  panic("TODO: implement")
}

func  (me *Light3D) SetEnableDistanceFade(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Light3D) IsDistanceFadeEnabled()  {
  panic("TODO: implement")
}

func  (me *Light3D) SetDistanceFadeBegin(distance float32, )  {
  panic("TODO: implement")
}

func  (me *Light3D) GetDistanceFadeBegin()  {
  panic("TODO: implement")
}

func  (me *Light3D) SetDistanceFadeShadow(distance float32, )  {
  panic("TODO: implement")
}

func  (me *Light3D) GetDistanceFadeShadow()  {
  panic("TODO: implement")
}

func  (me *Light3D) SetDistanceFadeLength(distance float32, )  {
  panic("TODO: implement")
}

func  (me *Light3D) GetDistanceFadeLength()  {
  panic("TODO: implement")
}

func  (me *Light3D) SetColor(color Color, )  {
  panic("TODO: implement")
}

func  (me *Light3D) GetColor()  {
  panic("TODO: implement")
}

func  (me *Light3D) SetShadowReverseCullFace(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Light3D) GetShadowReverseCullFace()  {
  panic("TODO: implement")
}

func  (me *Light3D) SetBakeMode(bake_mode Light3DBakeMode, )  {
  panic("TODO: implement")
}

func  (me *Light3D) GetBakeMode()  {
  panic("TODO: implement")
}

func  (me *Light3D) SetProjector(projector Texture2D, )  {
  panic("TODO: implement")
}

func  (me *Light3D) GetProjector()  {
  panic("TODO: implement")
}

func  (me *Light3D) SetTemperature(temperature float32, )  {
  panic("TODO: implement")
}

func  (me *Light3D) GetTemperature()  {
  panic("TODO: implement")
}

func  (me *Light3D) GetCorrelatedColor()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
