// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func  (me *Noise) GetNoise1D(x float32, )  {
  panic("TODO: implement")
}

func  (me *Noise) GetNoise2D(x float32, y float32, )  {
  panic("TODO: implement")
}

func  (me *Noise) GetNoise2Dv(v Vector2, )  {
  panic("TODO: implement")
}

func  (me *Noise) GetNoise3D(x float32, y float32, z float32, )  {
  panic("TODO: implement")
}

func  (me *Noise) GetNoise3Dv(v Vector3, )  {
  panic("TODO: implement")
}

func  (me *Noise) GetImage(width int, height int, invert bool, in_3d_space bool, normalize bool, )  {
  panic("TODO: implement")
}

func  (me *Noise) GetSeamlessImage(width int, height int, invert bool, in_3d_space bool, skirt float32, normalize bool, )  {
  panic("TODO: implement")
}

func  (me *Noise) GetImage3D(width int, height int, depth int, invert bool, normalize bool, )  {
  panic("TODO: implement")
}

func  (me *Noise) GetSeamlessImage3D(width int, height int, depth int, invert bool, skirt float32, normalize bool, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
