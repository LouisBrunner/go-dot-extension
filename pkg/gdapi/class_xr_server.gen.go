// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type XRServer struct {
  obj gdc.ObjectPtr
}

func createXRServer(obj gdc.ObjectPtr) *XRServer {
  return &XRServer{
    obj: obj,
  }
}

func (me *XRServer) BaseClass() string {
  return "XRServer"
}
