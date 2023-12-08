// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type BoneAttachment3D struct {
  obj gdc.ObjectPtr
}

func createBoneAttachment3D(obj gdc.ObjectPtr) *BoneAttachment3D {
  return &BoneAttachment3D{
    obj: obj,
  }
}

func (me *BoneAttachment3D) BaseClass() string {
  return "BoneAttachment3D"
}
