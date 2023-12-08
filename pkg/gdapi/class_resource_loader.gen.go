// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ResourceLoader struct {
  obj gdc.ObjectPtr
}

func createResourceLoader(obj gdc.ObjectPtr) *ResourceLoader {
  return &ResourceLoader{
    obj: obj,
  }
}

func (me *ResourceLoader) BaseClass() string {
  return "ResourceLoader"
}
