// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func  (me *CameraServer) GetFeed(index int, ) CameraFeed {
  classNameV := StringNameFromStr("CameraServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_feed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 361927068) // FIXME: should cache?
  var ret CameraFeed
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraServer) GetFeedCount() int {
  classNameV := StringNameFromStr("CameraServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_feed_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraServer) Feeds() CameraFeed {
  classNameV := StringNameFromStr("CameraServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("feeds")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret CameraFeed
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraServer) AddFeed(feed CameraFeed, )  {
  classNameV := StringNameFromStr("CameraServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_feed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3204782488) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(feed.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraServer) RemoveFeed(feed CameraFeed, )  {
  classNameV := StringNameFromStr("CameraServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_feed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3204782488) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(feed.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties
// Signals
// FIXME: can't seem to be able to connect them from this side of the API