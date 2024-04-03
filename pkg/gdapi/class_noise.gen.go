// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Noise struct {
  Resource
}

func (me *Noise) BaseClass() string {
  return "Noise"
}

func NewNoise() *Noise {
  str := StringNameFromStr("Noise") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Noise{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *Noise) GetNoise1D(x float64, ) float64 {
  classNameV := StringNameFromStr("Noise")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_noise_1d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3919130443) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&x), }
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Noise) GetNoise2D(x float64, y float64, ) float64 {
  classNameV := StringNameFromStr("Noise")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_noise_2d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2753205203) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), }
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Noise) GetNoise2Dv(v Vector2, ) float64 {
  classNameV := StringNameFromStr("Noise")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_noise_2dv")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2276447920) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(v.AsCTypePtr()), }
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Noise) GetNoise3D(x float64, y float64, z float64, ) float64 {
  classNameV := StringNameFromStr("Noise")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_noise_3d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 973811851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), gdc.ConstTypePtr(&z), }
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Noise) GetNoise3Dv(v Vector3, ) float64 {
  classNameV := StringNameFromStr("Noise")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_noise_3dv")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1109078154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(v.AsCTypePtr()), }
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Noise) GetImage(width int64, height int64, invert bool, in_3d_space bool, normalize bool, ) Image {
  classNameV := StringNameFromStr("Noise")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_image")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3180683109) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), gdc.ConstTypePtr(&invert), gdc.ConstTypePtr(&in_3d_space), gdc.ConstTypePtr(&normalize), }
  ret := NewImage()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Noise) GetSeamlessImage(width int64, height int64, invert bool, in_3d_space bool, skirt float64, normalize bool, ) Image {
  classNameV := StringNameFromStr("Noise")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_seamless_image")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2770743602) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), gdc.ConstTypePtr(&invert), gdc.ConstTypePtr(&in_3d_space), gdc.ConstTypePtr(&skirt), gdc.ConstTypePtr(&normalize), }
  ret := NewImage()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Noise) GetImage3D(width int64, height int64, depth int64, invert bool, normalize bool, ) []Image {
  classNameV := StringNameFromStr("Noise")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_image_3d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3977814329) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), gdc.ConstTypePtr(&depth), gdc.ConstTypePtr(&invert), gdc.ConstTypePtr(&normalize), }
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[Image](ret)
}

func  (me *Noise) GetSeamlessImage3D(width int64, height int64, depth int64, invert bool, skirt float64, normalize bool, ) []Image {
  classNameV := StringNameFromStr("Noise")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_seamless_image_3d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 451006340) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), gdc.ConstTypePtr(&depth), gdc.ConstTypePtr(&invert), gdc.ConstTypePtr(&skirt), gdc.ConstTypePtr(&normalize), }
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[Image](ret)
}

// Signals
