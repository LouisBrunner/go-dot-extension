// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFSkeleton struct {
  obj gdc.ObjectPtr
}

func (me *GLTFSkeleton) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFSkeleton) BaseClass() string {
  return "GLTFSkeleton"
}
