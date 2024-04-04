// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ImageFormatLoader struct {
  RefCounted
}

func (me *ImageFormatLoader) BaseClass() string {
  return "ImageFormatLoader"
}

func NewImageFormatLoader() *ImageFormatLoader {
  str := StringNameFromStr("ImageFormatLoader") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ImageFormatLoader{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type ImageFormatLoaderLoaderFlags int
const (
  ImageFormatLoaderLoaderFlagsFlagNone ImageFormatLoaderLoaderFlags = 0
  ImageFormatLoaderLoaderFlagsFlagForceLinear ImageFormatLoaderLoaderFlags = 1
  ImageFormatLoaderLoaderFlagsFlagConvertColors ImageFormatLoaderLoaderFlags = 2
)

func (me *ImageFormatLoader) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ImageFormatLoader) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ImageFormatLoader) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
