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

func  (me *CanvasItemMaterial) SetBlendMode(blend_mode CanvasItemMaterialBlendMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItemMaterial) GetBlendMode() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItemMaterial) SetLightMode(light_mode CanvasItemMaterialLightMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItemMaterial) GetLightMode() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItemMaterial) SetParticlesAnimation(particles_anim bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItemMaterial) GetParticlesAnimation() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItemMaterial) SetParticlesAnimHFrames(frames int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItemMaterial) GetParticlesAnimHFrames() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItemMaterial) SetParticlesAnimVFrames(frames int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItemMaterial) GetParticlesAnimVFrames() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItemMaterial) SetParticlesAnimLoop(loop bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItemMaterial) GetParticlesAnimLoop() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
