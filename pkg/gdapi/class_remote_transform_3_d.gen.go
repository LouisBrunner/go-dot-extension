// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RemoteTransform3D struct {
  obj gdc.ObjectPtr
}

func (me *RemoteTransform3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RemoteTransform3D) BaseClass() string {
  return "RemoteTransform3D"
}
