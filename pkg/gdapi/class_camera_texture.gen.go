// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type CameraTexture struct {
  Texture2D
}

func (me *CameraTexture) BaseClass() string {
  return "CameraTexture"
}

func NewCameraTexture() *CameraTexture {
  str := StringNameFromStr("CameraTexture") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CameraTexture{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *CameraTexture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CameraTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CameraTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CameraTexture) SetCameraFeedId(feed_id int64, )  {
  classNameV := StringNameFromStr("CameraTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_camera_feed_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&feed_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraTexture) GetCameraFeedId() int64 {
  classNameV := StringNameFromStr("CameraTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_camera_feed_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraTexture) SetWhichFeed(which_feed CameraServerFeedImage, )  {
  classNameV := StringNameFromStr("CameraTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_which_feed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1595299230) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&which_feed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraTexture) GetWhichFeed() CameraServerFeedImage {
  classNameV := StringNameFromStr("CameraTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_which_feed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 91039457) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret CameraServerFeedImage

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *CameraTexture) SetCameraActive(active bool, )  {
  classNameV := StringNameFromStr("CameraTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_camera_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraTexture) GetCameraActive() bool {
  classNameV := StringNameFromStr("CameraTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_camera_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
