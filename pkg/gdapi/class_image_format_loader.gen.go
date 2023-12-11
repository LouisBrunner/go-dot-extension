// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ImageFormatLoader struct {
  obj gdc.ObjectPtr
}

func (me *ImageFormatLoader) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ImageFormatLoader) BaseClass() string {
  return "ImageFormatLoader"
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
