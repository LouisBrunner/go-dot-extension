// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  CameraFeedFeedDataTypeFeedNoimage CameraFeedFeedDataType = 0
  CameraFeedFeedDataTypeFeedRgb CameraFeedFeedDataType = 1
  CameraFeedFeedDataTypeFeedYcbcr CameraFeedFeedDataType = 2
  CameraFeedFeedDataTypeFeedYcbcrSep CameraFeedFeedDataType = 3
)

type CameraFeedFeedPosition int
const (
  CameraFeedFeedPositionFeedUnspecified CameraFeedFeedPosition = 0
  CameraFeedFeedPositionFeedFront CameraFeedFeedPosition = 1
  CameraFeedFeedPositionFeedBack CameraFeedFeedPosition = 2
)

func  (me *CameraFeed) GetId() { // TODO: return value
  // TODO: implement
}

func  (me *CameraFeed) IsActive() { // TODO: return value
  // TODO: implement
}

func  (me *CameraFeed) SetActive(active bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CameraFeed) GetName() { // TODO: return value
  // TODO: implement
}

func  (me *CameraFeed) GetPosition() { // TODO: return value
  // TODO: implement
}

func  (me *CameraFeed) GetTransform() { // TODO: return value
  // TODO: implement
}

func  (me *CameraFeed) SetTransform(transform Transform2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *CameraFeed) GetDatatype() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
