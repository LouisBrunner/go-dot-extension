// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type DisplayServer struct {
  obj gdc.ObjectPtr
}

func createDisplayServer(obj gdc.ObjectPtr) *DisplayServer {
  return &DisplayServer{
    obj: obj,
  }
}

func (me *DisplayServer) BaseClass() string {
  return "DisplayServer"
}
