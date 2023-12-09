// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CameraTexture struct {
  obj gdc.ObjectPtr
}

func (me *CameraTexture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CameraTexture) BaseClass() string {
  return "CameraTexture"
}



// Enums

func (me *CameraTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CameraTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CameraTexture) SetCameraFeedId(feed_id int, )  {
  panic("TODO: implement")
}

func  (me *CameraTexture) GetCameraFeedId()  {
  panic("TODO: implement")
}

func  (me *CameraTexture) SetWhichFeed(which_feed CameraServerFeedImage, )  {
  panic("TODO: implement")
}

func  (me *CameraTexture) GetWhichFeed()  {
  panic("TODO: implement")
}

func  (me *CameraTexture) SetCameraActive(active bool, )  {
  panic("TODO: implement")
}

func  (me *CameraTexture) GetCameraActive()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
