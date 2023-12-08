// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CameraServer struct {
  obj gdc.ObjectPtr
}

func (me *CameraServer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CameraServer) BaseClass() string {
  return "CameraServer"
}

type CameraServerFeedImage int
const (
  CameraServerFeedRgbaImage CameraServerFeedImage = 0
  CameraServerFeedYcbcrImage CameraServerFeedImage = 0
  CameraServerFeedYImage CameraServerFeedImage = 0
  CameraServerFeedCbcrImage CameraServerFeedImage = 1
)
