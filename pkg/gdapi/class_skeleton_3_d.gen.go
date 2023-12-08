// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Skeleton3D struct {
  obj gdc.ObjectPtr
}

func (me *Skeleton3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Skeleton3D) BaseClass() string {
  return "Skeleton3D"
}

const (
  Skeleton3DNOTIFICATION_UPDATE_SKELETON = 50
)
