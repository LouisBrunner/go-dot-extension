// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  Light3DParamEnergy Light3DParam = 0
  Light3DParamIndirectEnergy Light3DParam = 1
  Light3DParamVolumetricFogEnergy Light3DParam = 2
  Light3DParamSpecular Light3DParam = 3
  Light3DParamRange Light3DParam = 4
  Light3DParamSize Light3DParam = 5
  Light3DParamAttenuation Light3DParam = 6
  Light3DParamSpotAngle Light3DParam = 7
  Light3DParamSpotAttenuation Light3DParam = 8
  Light3DParamShadowMaxDistance Light3DParam = 9
  Light3DParamShadowSplit1Offset Light3DParam = 10
  Light3DParamShadowSplit2Offset Light3DParam = 11
  Light3DParamShadowSplit3Offset Light3DParam = 12
  Light3DParamShadowFadeStart Light3DParam = 13
  Light3DParamShadowNormalBias Light3DParam = 14
  Light3DParamShadowBias Light3DParam = 15
  Light3DParamShadowPancakeSize Light3DParam = 16
  Light3DParamShadowOpacity Light3DParam = 17
  Light3DParamShadowBlur Light3DParam = 18
  Light3DParamTransmittanceBias Light3DParam = 19
  Light3DParamIntensity Light3DParam = 20
  Light3DParamMax Light3DParam = 21
)

type Light3DBakeMode int
const (
  Light3DBakeDisabled Light3DBakeMode = 0
  Light3DBakeStatic Light3DBakeMode = 1
  Light3DBakeDynamic Light3DBakeMode = 2
)
