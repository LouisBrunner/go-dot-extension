// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ImageFormatLoaderExtension struct {
  obj gdc.ObjectPtr
}

func (me *ImageFormatLoaderExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ImageFormatLoaderExtension) BaseClass() string {
  return "ImageFormatLoaderExtension"
}



// Enums

func (me *ImageFormatLoaderExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ImageFormatLoaderExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ImageFormatLoaderExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ImageFormatLoaderExtension) XGetRecognizedExtensions()  {
  panic("TODO: implement")
}

func  (me *ImageFormatLoaderExtension) XLoadImage(image Image, fileaccess FileAccess, flags ImageFormatLoaderLoaderFlags, scale float32, )  {
  panic("TODO: implement")
}

func  (me *ImageFormatLoaderExtension) AddFormatLoader()  {
  panic("TODO: implement")
}

func  (me *ImageFormatLoaderExtension) RemoveFormatLoader()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
