// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func  (me *CameraFeed) GetId() int {
  classNameV := StringNameFromStr("CameraFeed")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraFeed) IsActive() bool {
  classNameV := StringNameFromStr("CameraFeed")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraFeed) SetActive(active bool, )  {
  classNameV := StringNameFromStr("CameraFeed")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraFeed) GetName() String {
  classNameV := StringNameFromStr("CameraFeed")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraFeed) GetPosition() CameraFeedFeedPosition {
  classNameV := StringNameFromStr("CameraFeed")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2711679033) // FIXME: should cache?
  var ret CameraFeedFeedPosition
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraFeed) GetTransform() Transform2D {
  classNameV := StringNameFromStr("CameraFeed")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814499831) // FIXME: should cache?
  var ret Transform2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CameraFeed) SetTransform(transform Transform2D, )  {
  classNameV := StringNameFromStr("CameraFeed")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2761652528) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(transform.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CameraFeed) GetDatatype() CameraFeedFeedDataType {
  classNameV := StringNameFromStr("CameraFeed")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_datatype")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1477782850) // FIXME: should cache?
  var ret CameraFeedFeedDataType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *CameraFeed) GetPropFeedIsActive() bool {
  panic("TODO: implement")
}

func (me *CameraFeed) SetPropFeedIsActive(value bool) {
  panic("TODO: implement")
}

func (me *CameraFeed) GetPropFeedTransform() Transform2D {
  panic("TODO: implement")
}

func (me *CameraFeed) SetPropFeedTransform(value Transform2D) {
  panic("TODO: implement")
}