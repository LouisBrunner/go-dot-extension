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

type CameraFeed struct {
  RefCounted
}

func (me *CameraFeed) BaseClass() string {
  return "CameraFeed"
}

func NewCameraFeed() *CameraFeed {
  str := StringNameFromStr("CameraFeed") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CameraFeed{}
  obj.SetBaseObject(objPtr)
  return obj
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

func (me *CameraFeed) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CameraFeed) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CameraFeed) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CameraFeed) GetId() int64 {
  classNameV := StringNameFromStr("CameraFeed")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraFeed) IsActive() bool {
  classNameV := StringNameFromStr("CameraFeed")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraFeed) SetActive(active bool, )  {
  classNameV := StringNameFromStr("CameraFeed")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraFeed) GetName() String {
  classNameV := StringNameFromStr("CameraFeed")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CameraFeed) GetPosition() CameraFeedFeedPosition {
  classNameV := StringNameFromStr("CameraFeed")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2711679033) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret CameraFeedFeedPosition

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *CameraFeed) GetTransform() Transform2D {
  classNameV := StringNameFromStr("CameraFeed")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814499831) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CameraFeed) SetTransform(transform Transform2D, )  {
  classNameV := StringNameFromStr("CameraFeed")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2761652528) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraFeed) GetDatatype() CameraFeedFeedDataType {
  classNameV := StringNameFromStr("CameraFeed")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_datatype")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1477782850) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret CameraFeedFeedDataType

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
