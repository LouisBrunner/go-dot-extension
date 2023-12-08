// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamMicrophone struct {
  obj gdc.ObjectPtr
}

func (me *AudioStreamMicrophone) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioStreamMicrophone) BaseClass() string {
  return "AudioStreamMicrophone"
}
