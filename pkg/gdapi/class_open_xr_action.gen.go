// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type OpenXRAction struct {
  Resource
}

func (me *OpenXRAction) BaseClass() string {
  return "OpenXRAction"
}

func NewOpenXRAction() *OpenXRAction {
  str := StringNameFromStr("OpenXRAction") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &OpenXRAction{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type OpenXRActionActionType int
const (
  OpenXRActionActionTypeOpenxrActionBool OpenXRActionActionType = 0
  OpenXRActionActionTypeOpenxrActionFloat OpenXRActionActionType = 1
  OpenXRActionActionTypeOpenxrActionVector2 OpenXRActionActionType = 2
  OpenXRActionActionTypeOpenxrActionPose OpenXRActionActionType = 3
)

func (me *OpenXRAction) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *OpenXRAction) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *OpenXRAction) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *OpenXRAction) SetLocalizedName(localized_name String, )  {
  classNameV := StringNameFromStr("OpenXRAction")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_localized_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(localized_name.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OpenXRAction) GetLocalizedName() String {
  classNameV := StringNameFromStr("OpenXRAction")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_localized_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *OpenXRAction) SetActionType(action_type OpenXRActionActionType, )  {
  classNameV := StringNameFromStr("OpenXRAction")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_action_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1675238366) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&action_type), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OpenXRAction) GetActionType() OpenXRActionActionType {
  classNameV := StringNameFromStr("OpenXRAction")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_action_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3536542431) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret OpenXRActionActionType

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *OpenXRAction) SetToplevelPaths(toplevel_paths PackedStringArray, )  {
  classNameV := StringNameFromStr("OpenXRAction")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_toplevel_paths")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4015028928) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(toplevel_paths.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OpenXRAction) GetToplevelPaths() PackedStringArray {
  classNameV := StringNameFromStr("OpenXRAction")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_toplevel_paths")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
