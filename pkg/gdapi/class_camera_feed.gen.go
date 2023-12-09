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



// Enums

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

func (me *CameraFeed) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CameraFeed) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CameraFeed) GetId()  {
  panic("TODO: implement")
}

func  (me *CameraFeed) IsActive()  {
  panic("TODO: implement")
}

func  (me *CameraFeed) SetActive(active bool, )  {
  panic("TODO: implement")
}

func  (me *CameraFeed) GetName()  {
  panic("TODO: implement")
}

func  (me *CameraFeed) GetPosition()  {
  panic("TODO: implement")
}

func  (me *CameraFeed) GetTransform()  {
  panic("TODO: implement")
}

func  (me *CameraFeed) SetTransform(transform Transform2D, )  {
  panic("TODO: implement")
}

func  (me *CameraFeed) GetDatatype()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
