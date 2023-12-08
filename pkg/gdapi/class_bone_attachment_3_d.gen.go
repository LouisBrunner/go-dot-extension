// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type BoneAttachment3D struct {
  obj gdc.ObjectPtr
}

func (me *BoneAttachment3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *BoneAttachment3D) BaseClass() string {
  return "BoneAttachment3D"
}
