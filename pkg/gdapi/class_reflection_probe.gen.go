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



// Enums

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

func (me *ReflectionProbe) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ReflectionProbe) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ReflectionProbe) SetIntensity(intensity float32, )  {
  panic("TODO: implement")
}

func  (me *ReflectionProbe) GetIntensity()  {
  panic("TODO: implement")
}

func  (me *ReflectionProbe) SetAmbientMode(ambient ReflectionProbeAmbientMode, )  {
  panic("TODO: implement")
}

func  (me *ReflectionProbe) GetAmbientMode()  {
  panic("TODO: implement")
}

func  (me *ReflectionProbe) SetAmbientColor(ambient Color, )  {
  panic("TODO: implement")
}

func  (me *ReflectionProbe) GetAmbientColor()  {
  panic("TODO: implement")
}

func  (me *ReflectionProbe) SetAmbientColorEnergy(ambient_energy float32, )  {
  panic("TODO: implement")
}

func  (me *ReflectionProbe) GetAmbientColorEnergy()  {
  panic("TODO: implement")
}

func  (me *ReflectionProbe) SetMaxDistance(max_distance float32, )  {
  panic("TODO: implement")
}

func  (me *ReflectionProbe) GetMaxDistance()  {
  panic("TODO: implement")
}

func  (me *ReflectionProbe) SetMeshLodThreshold(ratio float32, )  {
  panic("TODO: implement")
}

func  (me *ReflectionProbe) GetMeshLodThreshold()  {
  panic("TODO: implement")
}

func  (me *ReflectionProbe) SetSize(size Vector3, )  {
  panic("TODO: implement")
}

func  (me *ReflectionProbe) GetSize()  {
  panic("TODO: implement")
}

func  (me *ReflectionProbe) SetOriginOffset(origin_offset Vector3, )  {
  panic("TODO: implement")
}

func  (me *ReflectionProbe) GetOriginOffset()  {
  panic("TODO: implement")
}

func  (me *ReflectionProbe) SetAsInterior(enable bool, )  {
  panic("TODO: implement")
}

func  (me *ReflectionProbe) IsSetAsInterior()  {
  panic("TODO: implement")
}

func  (me *ReflectionProbe) SetEnableBoxProjection(enable bool, )  {
  panic("TODO: implement")
}

func  (me *ReflectionProbe) IsBoxProjectionEnabled()  {
  panic("TODO: implement")
}

func  (me *ReflectionProbe) SetEnableShadows(enable bool, )  {
  panic("TODO: implement")
}

func  (me *ReflectionProbe) AreShadowsEnabled()  {
  panic("TODO: implement")
}

func  (me *ReflectionProbe) SetCullMask(layers int, )  {
  panic("TODO: implement")
}

func  (me *ReflectionProbe) GetCullMask()  {
  panic("TODO: implement")
}

func  (me *ReflectionProbe) SetUpdateMode(mode ReflectionProbeUpdateMode, )  {
  panic("TODO: implement")
}

func  (me *ReflectionProbe) GetUpdateMode()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
