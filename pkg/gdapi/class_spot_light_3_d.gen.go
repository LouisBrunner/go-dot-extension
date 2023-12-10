// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SpotLight3D struct {
  obj gdc.ObjectPtr
}

func (me *SpotLight3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SpotLight3D) BaseClass() string {
  return "SpotLight3D"
}



// Enums

func (me *SpotLight3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SpotLight3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SpotLight3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties

func (me *SpotLight3D) GetPropSpotRange() float32 {
  panic("TODO: implement")
}

func (me *SpotLight3D) SetPropSpotRange(value float32) {
  panic("TODO: implement")
}

func (me *SpotLight3D) GetPropSpotAttenuation() float32 {
  panic("TODO: implement")
}

func (me *SpotLight3D) SetPropSpotAttenuation(value float32) {
  panic("TODO: implement")
}

func (me *SpotLight3D) GetPropSpotAngle() float32 {
  panic("TODO: implement")
}

func (me *SpotLight3D) SetPropSpotAngle(value float32) {
  panic("TODO: implement")
}

func (me *SpotLight3D) GetPropSpotAngleAttenuation() float32 {
  panic("TODO: implement")
}

func (me *SpotLight3D) SetPropSpotAngleAttenuation(value float32) {
  panic("TODO: implement")
}