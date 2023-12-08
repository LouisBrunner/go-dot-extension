// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ResourcePreloader struct {
  obj gdc.ObjectPtr
}

func createResourcePreloader(obj gdc.ObjectPtr) *ResourcePreloader {
  return &ResourcePreloader{
    obj: obj,
  }
}

func (me *ResourcePreloader) BaseClass() string {
  return "ResourcePreloader"
}
