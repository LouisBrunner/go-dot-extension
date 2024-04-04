// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type CameraServer struct {
  Object
}

func (me *CameraServer) BaseClass() string {
  return "CameraServer"
}

func NewCameraServer() *CameraServer {
  str := StringNameFromStr("CameraServer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CameraServer{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *CameraServer) GetFeed(index int64, ) CameraFeed {
  classNameV := StringNameFromStr("CameraServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_feed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 361927068) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCameraFeed()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CameraServer) GetFeedCount() int64 {
  classNameV := StringNameFromStr("CameraServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_feed_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraServer) Feeds() []CameraFeed {
  classNameV := StringNameFromStr("CameraServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("feeds")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[CameraFeed](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *CameraServer) AddFeed(feed CameraFeed, )  {
  classNameV := StringNameFromStr("CameraServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_feed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3204782488) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{feed.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraServer) RemoveFeed(feed CameraFeed, )  {
  classNameV := StringNameFromStr("CameraServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_feed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3204782488) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{feed.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals

type CameraServerCameraFeedAddedSignalFn func(id int, )

func (me *CameraServer) ConnectCameraFeedAdded(subs SignalSubscribers, fn CameraServerCameraFeedAddedSignalFn) {
  sig := StringNameFromStr("camera_feed_added")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *CameraServer) DisconnectCameraFeedAdded(subs SignalSubscribers, fn CameraServerCameraFeedAddedSignalFn) {
  sig := StringNameFromStr("camera_feed_added")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type CameraServerCameraFeedRemovedSignalFn func(id int, )

func (me *CameraServer) ConnectCameraFeedRemoved(subs SignalSubscribers, fn CameraServerCameraFeedRemovedSignalFn) {
  sig := StringNameFromStr("camera_feed_removed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *CameraServer) DisconnectCameraFeedRemoved(subs SignalSubscribers, fn CameraServerCameraFeedRemovedSignalFn) {
  sig := StringNameFromStr("camera_feed_removed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
