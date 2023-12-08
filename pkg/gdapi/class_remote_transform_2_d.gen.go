// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RemoteTransform2D struct {
  obj gdc.ObjectPtr
}

func createRemoteTransform2D(obj gdc.ObjectPtr) *RemoteTransform2D {
  return &RemoteTransform2D{
    obj: obj,
  }
}

func (me *RemoteTransform2D) BaseClass() string {
  return "RemoteTransform2D"
}
