// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func  (me *ImageFormatLoaderExtension) AddFormatLoader()  {
  classNameV := StringNameFromStr("ImageFormatLoaderExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_format_loader")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ImageFormatLoaderExtension) RemoveFormatLoader()  {
  classNameV := StringNameFromStr("ImageFormatLoaderExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_format_loader")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties