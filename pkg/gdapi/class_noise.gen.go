// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Noise struct {
  obj gdc.ObjectPtr
}

func (me *Noise) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Noise) BaseClass() string {
  return "Noise"
}



// Enums

func (me *Noise) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Noise) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Noise) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Noise) GetNoise1D(x float32, ) float32 {
  classNameV := StringNameFromStr("Noise")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_noise_1d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3919130443) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&x), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Noise) GetNoise2D(x float32, y float32, ) float32 {
  classNameV := StringNameFromStr("Noise")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_noise_2d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2753205203) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Noise) GetNoise2Dv(v Vector2, ) float32 {
  classNameV := StringNameFromStr("Noise")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_noise_2dv")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2276447920) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(v.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Noise) GetNoise3D(x float32, y float32, z float32, ) float32 {
  classNameV := StringNameFromStr("Noise")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_noise_3d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 973811851) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), gdc.ConstTypePtr(&z), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Noise) GetNoise3Dv(v Vector3, ) float32 {
  classNameV := StringNameFromStr("Noise")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_noise_3dv")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1109078154) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(v.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Noise) GetImage(width int, height int, invert bool, in_3d_space bool, normalize bool, ) Image {
  classNameV := StringNameFromStr("Noise")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_image")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2569233413) // FIXME: should cache?
  var ret Image
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), gdc.ConstTypePtr(&invert), gdc.ConstTypePtr(&in_3d_space), gdc.ConstTypePtr(&normalize), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Noise) GetSeamlessImage(width int, height int, invert bool, in_3d_space bool, skirt float32, normalize bool, ) Image {
  classNameV := StringNameFromStr("Noise")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_seamless_image")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2210827790) // FIXME: should cache?
  var ret Image
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), gdc.ConstTypePtr(&invert), gdc.ConstTypePtr(&in_3d_space), gdc.ConstTypePtr(&skirt), gdc.ConstTypePtr(&normalize), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Noise) GetImage3D(width int, height int, depth int, invert bool, normalize bool, ) Image {
  classNameV := StringNameFromStr("Noise")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_image_3d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2358868431) // FIXME: should cache?
  var ret Image
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), gdc.ConstTypePtr(&depth), gdc.ConstTypePtr(&invert), gdc.ConstTypePtr(&normalize), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Noise) GetSeamlessImage3D(width int, height int, depth int, invert bool, skirt float32, normalize bool, ) Image {
  classNameV := StringNameFromStr("Noise")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_seamless_image_3d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3328694319) // FIXME: should cache?
  var ret Image
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), gdc.ConstTypePtr(&depth), gdc.ConstTypePtr(&invert), gdc.ConstTypePtr(&skirt), gdc.ConstTypePtr(&normalize), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
