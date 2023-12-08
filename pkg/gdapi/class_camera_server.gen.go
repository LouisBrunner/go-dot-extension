// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CameraServer struct {
  obj gdc.ObjectPtr
}

func createCameraServer(obj gdc.ObjectPtr) *CameraServer {
  return &CameraServer{
    obj: obj,
  }
}

func (me *CameraServer) BaseClass() string {
  return "CameraServer"
}
