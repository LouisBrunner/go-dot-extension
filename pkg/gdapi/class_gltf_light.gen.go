// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFLight struct {
  obj gdc.ObjectPtr
}

func (me *GLTFLight) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFLight) BaseClass() string {
  return "GLTFLight"
}



// Enums

func (me *GLTFLight) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GLTFLight) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  GLTFLightFromNode(light_node Light3D, )  {
  panic("TODO: implement")
}

func  (me *GLTFLight) ToNode()  {
  panic("TODO: implement")
}

func  GLTFLightFromDictionary(dictionary Dictionary, )  {
  panic("TODO: implement")
}

func  (me *GLTFLight) ToDictionary()  {
  panic("TODO: implement")
}

func  (me *GLTFLight) GetColor()  {
  panic("TODO: implement")
}

func  (me *GLTFLight) SetColor(color Color, )  {
  panic("TODO: implement")
}

func  (me *GLTFLight) GetIntensity()  {
  panic("TODO: implement")
}

func  (me *GLTFLight) SetIntensity(intensity float32, )  {
  panic("TODO: implement")
}

func  (me *GLTFLight) GetLightType()  {
  panic("TODO: implement")
}

func  (me *GLTFLight) SetLightType(light_type String, )  {
  panic("TODO: implement")
}

func  (me *GLTFLight) GetRange()  {
  panic("TODO: implement")
}

func  (me *GLTFLight) SetRange(range_ float32, )  {
  panic("TODO: implement")
}

func  (me *GLTFLight) GetInnerConeAngle()  {
  panic("TODO: implement")
}

func  (me *GLTFLight) SetInnerConeAngle(inner_cone_angle float32, )  {
  panic("TODO: implement")
}

func  (me *GLTFLight) GetOuterConeAngle()  {
  panic("TODO: implement")
}

func  (me *GLTFLight) SetOuterConeAngle(outer_cone_angle float32, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
