// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CanvasItemMaterial struct {
  Material
}

func (me *CanvasItemMaterial) BaseClass() string {
  return "CanvasItemMaterial"
}

func NewCanvasItemMaterial() *CanvasItemMaterial {
  str := StringNameFromStr("CanvasItemMaterial") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CanvasItemMaterial{}
  obj.SetBaseObject(objPtr)
  return obj
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
  classNameV := StringNameFromStr("CanvasItemMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_blend_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1786054936) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&blend_mode), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItemMaterial) GetBlendMode() CanvasItemMaterialBlendMode {
  classNameV := StringNameFromStr("CanvasItemMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3318684035) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret CanvasItemMaterialBlendMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *CanvasItemMaterial) SetLightMode(light_mode CanvasItemMaterialLightMode, )  {
  classNameV := StringNameFromStr("CanvasItemMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_light_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 628074070) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&light_mode), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItemMaterial) GetLightMode() CanvasItemMaterialLightMode {
  classNameV := StringNameFromStr("CanvasItemMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_light_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3863292382) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret CanvasItemMaterialLightMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *CanvasItemMaterial) SetParticlesAnimation(particles_anim bool, )  {
  classNameV := StringNameFromStr("CanvasItemMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_particles_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&particles_anim), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItemMaterial) GetParticlesAnimation() bool {
  classNameV := StringNameFromStr("CanvasItemMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_particles_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CanvasItemMaterial) SetParticlesAnimHFrames(frames int64, )  {
  classNameV := StringNameFromStr("CanvasItemMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_particles_anim_h_frames")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frames), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItemMaterial) GetParticlesAnimHFrames() int64 {
  classNameV := StringNameFromStr("CanvasItemMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_particles_anim_h_frames")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CanvasItemMaterial) SetParticlesAnimVFrames(frames int64, )  {
  classNameV := StringNameFromStr("CanvasItemMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_particles_anim_v_frames")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frames), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItemMaterial) GetParticlesAnimVFrames() int64 {
  classNameV := StringNameFromStr("CanvasItemMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_particles_anim_v_frames")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CanvasItemMaterial) SetParticlesAnimLoop(loop bool, )  {
  classNameV := StringNameFromStr("CanvasItemMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_particles_anim_loop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loop), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItemMaterial) GetParticlesAnimLoop() bool {
  classNameV := StringNameFromStr("CanvasItemMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_particles_anim_loop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
