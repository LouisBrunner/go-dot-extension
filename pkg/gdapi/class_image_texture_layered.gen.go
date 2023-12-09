// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ImageTextureLayered struct {
  obj gdc.ObjectPtr
}

func (me *ImageTextureLayered) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ImageTextureLayered) BaseClass() string {
  return "ImageTextureLayered"
}



// Enums

func (me *ImageTextureLayered) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ImageTextureLayered) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ImageTextureLayered) CreateFromImages(images Image, )  {
  panic("TODO: implement")
}

func  (me *ImageTextureLayered) UpdateLayer(image Image, layer int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
