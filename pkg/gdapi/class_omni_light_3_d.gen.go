// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type OmniLight3D struct {
  obj gdc.ObjectPtr
}

func (me *OmniLight3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *OmniLight3D) BaseClass() string {
  return "OmniLight3D"
}
