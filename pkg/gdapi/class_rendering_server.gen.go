// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RenderingServer struct {
  obj gdc.ObjectPtr
}

func createRenderingServer(obj gdc.ObjectPtr) *RenderingServer {
  return &RenderingServer{
    obj: obj,
  }
}

func (me *RenderingServer) BaseClass() string {
  return "RenderingServer"
}
