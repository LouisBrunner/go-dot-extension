// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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



// Enums

type CameraServerFeedImage int
const (
  CameraServerFeedImageFeedRgbaImage CameraServerFeedImage = 0
  CameraServerFeedImageFeedYcbcrImage CameraServerFeedImage = 0
  CameraServerFeedImageFeedYImage CameraServerFeedImage = 0
  CameraServerFeedImageFeedCbcrImage CameraServerFeedImage = 1
)

func (me *CameraServer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CameraServer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CameraServer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CameraServer) GetFeed(index int, )  {
  panic("TODO: implement")
}

func  (me *CameraServer) GetFeedCount()  {
  panic("TODO: implement")
}

func  (me *CameraServer) Feeds()  {
  panic("TODO: implement")
}

func  (me *CameraServer) AddFeed(feed CameraFeed, )  {
  panic("TODO: implement")
}

func  (me *CameraServer) RemoveFeed(feed CameraFeed, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
