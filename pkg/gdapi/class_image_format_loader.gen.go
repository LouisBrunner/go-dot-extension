// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ImageFormatLoader struct {
  obj gdc.ObjectPtr
}

func (me *ImageFormatLoader) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ImageFormatLoader) BaseClass() string {
  return "ImageFormatLoader"
}

type ImageFormatLoaderLoaderFlags int
const (
  ImageFormatLoaderLoaderFlagsFlagNone ImageFormatLoaderLoaderFlags = 0
  ImageFormatLoaderLoaderFlagsFlagForceLinear ImageFormatLoaderLoaderFlags = 1
  ImageFormatLoaderLoaderFlagsFlagConvertColors ImageFormatLoaderLoaderFlags = 2
)

// TODO: properties

// TODO: signals
