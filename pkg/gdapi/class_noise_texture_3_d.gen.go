// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type NoiseTexture3D struct {
  obj gdc.ObjectPtr
}

func (me *NoiseTexture3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NoiseTexture3D) BaseClass() string {
  return "NoiseTexture3D"
}



// Enums

func (me *NoiseTexture3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *NoiseTexture3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NoiseTexture3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *NoiseTexture3D) SetWidth(width int, )  {
  classNameV := StringNameFromStr("NoiseTexture3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NoiseTexture3D) SetHeight(height int, )  {
  classNameV := StringNameFromStr("NoiseTexture3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NoiseTexture3D) SetDepth(depth int, )  {
  classNameV := StringNameFromStr("NoiseTexture3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_depth")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&depth), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NoiseTexture3D) SetInvert(invert bool, )  {
  classNameV := StringNameFromStr("NoiseTexture3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_invert")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&invert), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NoiseTexture3D) GetInvert() bool {
  classNameV := StringNameFromStr("NoiseTexture3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_invert")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NoiseTexture3D) SetSeamless(seamless bool, )  {
  classNameV := StringNameFromStr("NoiseTexture3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_seamless")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seamless), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NoiseTexture3D) GetSeamless() bool {
  classNameV := StringNameFromStr("NoiseTexture3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_seamless")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NoiseTexture3D) SetSeamlessBlendSkirt(seamless_blend_skirt float32, )  {
  classNameV := StringNameFromStr("NoiseTexture3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_seamless_blend_skirt")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seamless_blend_skirt), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NoiseTexture3D) GetSeamlessBlendSkirt() float32 {
  classNameV := StringNameFromStr("NoiseTexture3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_seamless_blend_skirt")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NoiseTexture3D) SetNormalize(normalize bool, )  {
  classNameV := StringNameFromStr("NoiseTexture3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_normalize")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&normalize), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NoiseTexture3D) IsNormalized() bool {
  classNameV := StringNameFromStr("NoiseTexture3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_normalized")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NoiseTexture3D) SetColorRamp(gradient Gradient, )  {
  classNameV := StringNameFromStr("NoiseTexture3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color_ramp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2756054477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(gradient.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NoiseTexture3D) GetColorRamp() Gradient {
  classNameV := StringNameFromStr("NoiseTexture3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color_ramp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 132272999) // FIXME: should cache?
  var ret Gradient
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NoiseTexture3D) SetNoise(noise Noise, )  {
  classNameV := StringNameFromStr("NoiseTexture3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_noise")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4135492439) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(noise.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NoiseTexture3D) GetNoise() Noise {
  classNameV := StringNameFromStr("NoiseTexture3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_noise")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 185851837) // FIXME: should cache?
  var ret Noise
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *NoiseTexture3D) GetPropWidth() int {
  panic("TODO: implement")
}

func (me *NoiseTexture3D) SetPropWidth(value int) {
  panic("TODO: implement")
}

func (me *NoiseTexture3D) GetPropHeight() int {
  panic("TODO: implement")
}

func (me *NoiseTexture3D) SetPropHeight(value int) {
  panic("TODO: implement")
}

func (me *NoiseTexture3D) GetPropDepth() int {
  panic("TODO: implement")
}

func (me *NoiseTexture3D) SetPropDepth(value int) {
  panic("TODO: implement")
}

func (me *NoiseTexture3D) GetPropInvert() bool {
  panic("TODO: implement")
}

func (me *NoiseTexture3D) SetPropInvert(value bool) {
  panic("TODO: implement")
}

func (me *NoiseTexture3D) GetPropSeamless() bool {
  panic("TODO: implement")
}

func (me *NoiseTexture3D) SetPropSeamless(value bool) {
  panic("TODO: implement")
}

func (me *NoiseTexture3D) GetPropSeamlessBlendSkirt() float32 {
  panic("TODO: implement")
}

func (me *NoiseTexture3D) SetPropSeamlessBlendSkirt(value float32) {
  panic("TODO: implement")
}

func (me *NoiseTexture3D) GetPropNormalize() bool {
  panic("TODO: implement")
}

func (me *NoiseTexture3D) SetPropNormalize(value bool) {
  panic("TODO: implement")
}

func (me *NoiseTexture3D) GetPropColorRamp() Gradient {
  panic("TODO: implement")
}

func (me *NoiseTexture3D) SetPropColorRamp(value Gradient) {
  panic("TODO: implement")
}

func (me *NoiseTexture3D) GetPropNoise() Noise {
  panic("TODO: implement")
}

func (me *NoiseTexture3D) SetPropNoise(value Noise) {
  panic("TODO: implement")
}