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

type ptrsForCameraFeedList struct {
  fnGetId gdc.MethodBindPtr
  fnIsActive gdc.MethodBindPtr
  fnSetActive gdc.MethodBindPtr
  fnGetName gdc.MethodBindPtr
  fnGetPosition gdc.MethodBindPtr
  fnGetTransform gdc.MethodBindPtr
  fnSetTransform gdc.MethodBindPtr
  fnGetDatatype gdc.MethodBindPtr
}

var ptrsForCameraFeed ptrsForCameraFeedList

func initCameraFeedPtrs(iface gdc.Interface) {

  className := StringNameFromStr("CameraFeed")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_id")
    defer methodName.Destroy()
    ptrsForCameraFeed.fnGetId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("is_active")
    defer methodName.Destroy()
    ptrsForCameraFeed.fnIsActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_active")
    defer methodName.Destroy()
    ptrsForCameraFeed.fnSetActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_name")
    defer methodName.Destroy()
    ptrsForCameraFeed.fnGetName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("get_position")
    defer methodName.Destroy()
    ptrsForCameraFeed.fnGetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2711679033))
  }
  {
    methodName := StringNameFromStr("get_transform")
    defer methodName.Destroy()
    ptrsForCameraFeed.fnGetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814499831))
  }
  {
    methodName := StringNameFromStr("set_transform")
    defer methodName.Destroy()
    ptrsForCameraFeed.fnSetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2761652528))
  }
  {
    methodName := StringNameFromStr("get_datatype")
    defer methodName.Destroy()
    ptrsForCameraFeed.fnGetDatatype = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1477782850))
  }
}

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
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraFeed.fnGetId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraFeed) IsActive() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraFeed.fnIsActive), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CameraFeed) SetActive(active bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraFeed.fnSetActive), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraFeed) GetName() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraFeed.fnGetName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CameraFeed) GetPosition() CameraFeedFeedPosition {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret CameraFeedFeedPosition

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraFeed.fnGetPosition), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *CameraFeed) GetTransform() Transform2D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform2D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraFeed.fnGetTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CameraFeed) SetTransform(transform Transform2D, )  {
  cargs := []gdc.ConstTypePtr{transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraFeed.fnSetTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CameraFeed) GetDatatype() CameraFeedFeedDataType {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret CameraFeedFeedDataType

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraFeed.fnGetDatatype), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
