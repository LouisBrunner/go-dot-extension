// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ImageFormatLoaderExtension struct {
  obj gdc.ObjectPtr
}

func createImageFormatLoaderExtension(obj gdc.ObjectPtr) *ImageFormatLoaderExtension {
  return &ImageFormatLoaderExtension{
    obj: obj,
  }
}

func (me *ImageFormatLoaderExtension) BaseClass() string {
  return "ImageFormatLoaderExtension"
}
