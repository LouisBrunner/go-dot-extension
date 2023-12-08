// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CameraFeed struct {
  obj gdc.ObjectPtr
}

func (me *CameraFeed) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CameraFeed) BaseClass() string {
  return "CameraFeed"
}

type CameraFeedFeedDataType int
const (
  CameraFeedFeedNoimage CameraFeedFeedDataType = 0
  CameraFeedFeedRgb CameraFeedFeedDataType = 1
  CameraFeedFeedYcbcr CameraFeedFeedDataType = 2
  CameraFeedFeedYcbcrSep CameraFeedFeedDataType = 3
)

type CameraFeedFeedPosition int
const (
  CameraFeedFeedUnspecified CameraFeedFeedPosition = 0
  CameraFeedFeedFront CameraFeedFeedPosition = 1
  CameraFeedFeedBack CameraFeedFeedPosition = 2
)
