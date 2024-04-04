// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type ImageFormatLoaderExtension struct {
  ImageFormatLoader
}

func (me *ImageFormatLoaderExtension) BaseClass() string {
  return "ImageFormatLoaderExtension"
}

func NewImageFormatLoaderExtension() *ImageFormatLoaderExtension {
  str := StringNameFromStr("ImageFormatLoaderExtension") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ImageFormatLoaderExtension{}
  obj.SetBaseObject(objPtr)
  return obj
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
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ImageFormatLoaderExtension) RemoveFormatLoader()  {
  classNameV := StringNameFromStr("ImageFormatLoaderExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_format_loader")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
