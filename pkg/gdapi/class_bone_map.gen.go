// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type BoneMap struct {
  obj gdc.ObjectPtr
}

func createBoneMap(obj gdc.ObjectPtr) *BoneMap {
  return &BoneMap{
    obj: obj,
  }
}

func (me *BoneMap) BaseClass() string {
  return "BoneMap"
}
