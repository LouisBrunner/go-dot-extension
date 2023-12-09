// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CanvasItemMaterial struct {
  obj gdc.ObjectPtr
}

func (me *CanvasItemMaterial) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CanvasItemMaterial) BaseClass() string {
  return "CanvasItemMaterial"
}



// Enums

type CanvasItemMaterialBlendMode int
const (
  CanvasItemMaterialBlendModeBlendModeMix CanvasItemMaterialBlendMode = 0
  CanvasItemMaterialBlendModeBlendModeAdd CanvasItemMaterialBlendMode = 1
  CanvasItemMaterialBlendModeBlendModeSub CanvasItemMaterialBlendMode = 2
  CanvasItemMaterialBlendModeBlendModeMul CanvasItemMaterialBlendMode = 3
  CanvasItemMaterialBlendModeBlendModePremultAlpha CanvasItemMaterialBlendMode = 4
)

type CanvasItemMaterialLightMode int
const (
  CanvasItemMaterialLightModeLightModeNormal CanvasItemMaterialLightMode = 0
  CanvasItemMaterialLightModeLightModeUnshaded CanvasItemMaterialLightMode = 1
  CanvasItemMaterialLightModeLightModeLightOnly CanvasItemMaterialLightMode = 2
)

func (me *CanvasItemMaterial) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CanvasItemMaterial) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CanvasItemMaterial) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CanvasItemMaterial) SetBlendMode(blend_mode CanvasItemMaterialBlendMode, )  {
  panic("TODO: implement")
}

func  (me *CanvasItemMaterial) GetBlendMode()  {
  panic("TODO: implement")
}

func  (me *CanvasItemMaterial) SetLightMode(light_mode CanvasItemMaterialLightMode, )  {
  panic("TODO: implement")
}

func  (me *CanvasItemMaterial) GetLightMode()  {
  panic("TODO: implement")
}

func  (me *CanvasItemMaterial) SetParticlesAnimation(particles_anim bool, )  {
  panic("TODO: implement")
}

func  (me *CanvasItemMaterial) GetParticlesAnimation()  {
  panic("TODO: implement")
}

func  (me *CanvasItemMaterial) SetParticlesAnimHFrames(frames int, )  {
  panic("TODO: implement")
}

func  (me *CanvasItemMaterial) GetParticlesAnimHFrames()  {
  panic("TODO: implement")
}

func  (me *CanvasItemMaterial) SetParticlesAnimVFrames(frames int, )  {
  panic("TODO: implement")
}

func  (me *CanvasItemMaterial) GetParticlesAnimVFrames()  {
  panic("TODO: implement")
}

func  (me *CanvasItemMaterial) SetParticlesAnimLoop(loop bool, )  {
  panic("TODO: implement")
}

func  (me *CanvasItemMaterial) GetParticlesAnimLoop()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
