// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WebXRInterface struct {
  obj gdc.ObjectPtr
}

func createWebXRInterface(obj gdc.ObjectPtr) *WebXRInterface {
  return &WebXRInterface{
    obj: obj,
  }
}

func (me *WebXRInterface) BaseClass() string {
  return "WebXRInterface"
}
