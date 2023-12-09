// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFTexture struct {
  obj gdc.ObjectPtr
}

func (me *GLTFTexture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFTexture) BaseClass() string {
  return "GLTFTexture"
}



// Enums

func (me *GLTFTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GLTFTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *GLTFTexture) GetSrcImage()  {
  panic("TODO: implement")
}

func  (me *GLTFTexture) SetSrcImage(src_image int, )  {
  panic("TODO: implement")
}

func  (me *GLTFTexture) GetSampler()  {
  panic("TODO: implement")
}

func  (me *GLTFTexture) SetSampler(sampler int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
