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

type CameraServerFeedImage int
const (
  CameraServerFeedImageFeedRgbaImage CameraServerFeedImage = 0
  CameraServerFeedImageFeedYcbcrImage CameraServerFeedImage = 0
  CameraServerFeedImageFeedYImage CameraServerFeedImage = 0
  CameraServerFeedImageFeedCbcrImage CameraServerFeedImage = 1
)

func  (me *CameraServer) GetFeed(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CameraServer) GetFeedCount() { // TODO: return value
  // TODO: implement
}

func  (me *CameraServer) Feeds() { // TODO: return value
  // TODO: implement
}

func  (me *CameraServer) AddFeed(feed CameraFeed, ) { // TODO: return value
  // TODO: implement
}

func  (me *CameraServer) RemoveFeed(feed CameraFeed, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
