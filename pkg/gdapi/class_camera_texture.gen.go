// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func  (me *CameraTexture) SetCameraFeedId(feed_id int, )  {
  classNameV := StringNameFromStr("CameraTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_camera_feed_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&feed_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraTexture) GetCameraFeedId() int {
  classNameV := StringNameFromStr("CameraTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_camera_feed_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraTexture) SetWhichFeed(which_feed CameraServerFeedImage, )  {
  classNameV := StringNameFromStr("CameraTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_which_feed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1595299230) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&which_feed), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraTexture) GetWhichFeed() CameraServerFeedImage {
  classNameV := StringNameFromStr("CameraTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_which_feed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 91039457) // FIXME: should cache?
  var ret CameraServerFeedImage
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraTexture) SetCameraActive(active bool, )  {
  classNameV := StringNameFromStr("CameraTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_camera_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraTexture) GetCameraActive() bool {
  classNameV := StringNameFromStr("CameraTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_camera_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *CameraTexture) GetPropCameraFeedId() int {
  panic("TODO: implement")
}

func (me *CameraTexture) SetPropCameraFeedId(value int) {
  panic("TODO: implement")
}

func (me *CameraTexture) GetPropWhichFeed() int {
  panic("TODO: implement")
}

func (me *CameraTexture) SetPropWhichFeed(value int) {
  panic("TODO: implement")
}

func (me *CameraTexture) GetPropCameraIsActive() bool {
  panic("TODO: implement")
}

func (me *CameraTexture) SetPropCameraIsActive(value bool) {
  panic("TODO: implement")
}