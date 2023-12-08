// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFSkeleton struct {
  obj gdc.ObjectPtr
}

func createGLTFSkeleton(obj gdc.ObjectPtr) *GLTFSkeleton {
  return &GLTFSkeleton{
    obj: obj,
  }
}

func (me *GLTFSkeleton) BaseClass() string {
  return "GLTFSkeleton"
}
