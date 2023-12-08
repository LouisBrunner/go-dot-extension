// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ReflectionProbe struct {
  obj gdc.ObjectPtr
}

func (me *ReflectionProbe) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ReflectionProbe) BaseClass() string {
  return "ReflectionProbe"
}

type ReflectionProbeUpdateMode int
const (
  ReflectionProbeUpdateModeUpdateOnce ReflectionProbeUpdateMode = 0
  ReflectionProbeUpdateModeUpdateAlways ReflectionProbeUpdateMode = 1
)

type ReflectionProbeAmbientMode int
const (
  ReflectionProbeAmbientModeAmbientDisabled ReflectionProbeAmbientMode = 0
  ReflectionProbeAmbientModeAmbientEnvironment ReflectionProbeAmbientMode = 1
  ReflectionProbeAmbientModeAmbientColor ReflectionProbeAmbientMode = 2
)

func  (me *ReflectionProbe) SetIntensity(intensity float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *ReflectionProbe) GetIntensity() { // TODO: return value
  // TODO: implement
}

func  (me *ReflectionProbe) SetAmbientMode(ambient ReflectionProbeAmbientMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *ReflectionProbe) GetAmbientMode() { // TODO: return value
  // TODO: implement
}

func  (me *ReflectionProbe) SetAmbientColor(ambient Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *ReflectionProbe) GetAmbientColor() { // TODO: return value
  // TODO: implement
}

func  (me *ReflectionProbe) SetAmbientColorEnergy(ambient_energy float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *ReflectionProbe) GetAmbientColorEnergy() { // TODO: return value
  // TODO: implement
}

func  (me *ReflectionProbe) SetMaxDistance(max_distance float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *ReflectionProbe) GetMaxDistance() { // TODO: return value
  // TODO: implement
}

func  (me *ReflectionProbe) SetMeshLodThreshold(ratio float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *ReflectionProbe) GetMeshLodThreshold() { // TODO: return value
  // TODO: implement
}

func  (me *ReflectionProbe) SetSize(size Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *ReflectionProbe) GetSize() { // TODO: return value
  // TODO: implement
}

func  (me *ReflectionProbe) SetOriginOffset(origin_offset Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *ReflectionProbe) GetOriginOffset() { // TODO: return value
  // TODO: implement
}

func  (me *ReflectionProbe) SetAsInterior(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ReflectionProbe) IsSetAsInterior() { // TODO: return value
  // TODO: implement
}

func  (me *ReflectionProbe) SetEnableBoxProjection(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ReflectionProbe) IsBoxProjectionEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *ReflectionProbe) SetEnableShadows(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ReflectionProbe) AreShadowsEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *ReflectionProbe) SetCullMask(layers int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ReflectionProbe) GetCullMask() { // TODO: return value
  // TODO: implement
}

func  (me *ReflectionProbe) SetUpdateMode(mode ReflectionProbeUpdateMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *ReflectionProbe) GetUpdateMode() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
