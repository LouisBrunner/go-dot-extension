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

type OpenXRActionSet struct {
  Resource
}

func (me *OpenXRActionSet) BaseClass() string {
  return "OpenXRActionSet"
}

func NewOpenXRActionSet() *OpenXRActionSet {
  str := StringNameFromStr("OpenXRActionSet") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &OpenXRActionSet{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *OpenXRActionSet) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *OpenXRActionSet) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *OpenXRActionSet) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *OpenXRActionSet) SetLocalizedName(localized_name String, )  {
  classNameV := StringNameFromStr("OpenXRActionSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_localized_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{localized_name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OpenXRActionSet) GetLocalizedName() String {
  classNameV := StringNameFromStr("OpenXRActionSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_localized_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *OpenXRActionSet) SetPriority(priority int64, )  {
  classNameV := StringNameFromStr("OpenXRActionSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OpenXRActionSet) GetPriority() int64 {
  classNameV := StringNameFromStr("OpenXRActionSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OpenXRActionSet) GetActionCount() int64 {
  classNameV := StringNameFromStr("OpenXRActionSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_action_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OpenXRActionSet) SetActions(actions Array, )  {
  classNameV := StringNameFromStr("OpenXRActionSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_actions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{actions.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OpenXRActionSet) GetActions() Array {
  classNameV := StringNameFromStr("OpenXRActionSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_actions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *OpenXRActionSet) AddAction(action OpenXRAction, )  {
  classNameV := StringNameFromStr("OpenXRActionSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_action")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 349361333) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OpenXRActionSet) RemoveAction(action OpenXRAction, )  {
  classNameV := StringNameFromStr("OpenXRActionSet")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_action")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 349361333) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
