// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type FogMaterial struct {
  obj gdc.ObjectPtr
}

func (me *FogMaterial) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *FogMaterial) BaseClass() string {
  return "FogMaterial"
}



// Enums

func (me *FogMaterial) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *FogMaterial) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *FogMaterial) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *FogMaterial) SetDensity(density float32, )  {
  panic("TODO: implement")
}

func  (me *FogMaterial) GetDensity()  {
  panic("TODO: implement")
}

func  (me *FogMaterial) SetAlbedo(albedo Color, )  {
  panic("TODO: implement")
}

func  (me *FogMaterial) GetAlbedo()  {
  panic("TODO: implement")
}

func  (me *FogMaterial) SetEmission(emission Color, )  {
  panic("TODO: implement")
}

func  (me *FogMaterial) GetEmission()  {
  panic("TODO: implement")
}

func  (me *FogMaterial) SetHeightFalloff(height_falloff float32, )  {
  panic("TODO: implement")
}

func  (me *FogMaterial) GetHeightFalloff()  {
  panic("TODO: implement")
}

func  (me *FogMaterial) SetEdgeFade(edge_fade float32, )  {
  panic("TODO: implement")
}

func  (me *FogMaterial) GetEdgeFade()  {
  panic("TODO: implement")
}

func  (me *FogMaterial) SetDensityTexture(density_texture Texture3D, )  {
  panic("TODO: implement")
}

func  (me *FogMaterial) GetDensityTexture()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
