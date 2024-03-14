// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type OpenXRIPBinding struct {
  Resource
}

func (me *OpenXRIPBinding) BaseClass() string {
  return "OpenXRIPBinding"
}



// Enums

func (me *OpenXRIPBinding) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *OpenXRIPBinding) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *OpenXRIPBinding) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *OpenXRIPBinding) SetAction(action OpenXRAction, )  {
  classNameV := StringNameFromStr("OpenXRIPBinding")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_action")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 349361333) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(action.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OpenXRIPBinding) GetAction() OpenXRAction {
  classNameV := StringNameFromStr("OpenXRIPBinding")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_action")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4072409085) // FIXME: should cache?
  var ret OpenXRAction
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRIPBinding) GetPathCount() int {
  classNameV := StringNameFromStr("OpenXRIPBinding")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRIPBinding) SetPaths(paths PackedStringArray, )  {
  classNameV := StringNameFromStr("OpenXRIPBinding")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_paths")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4015028928) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(paths.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OpenXRIPBinding) GetPaths() PackedStringArray {
  classNameV := StringNameFromStr("OpenXRIPBinding")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_paths")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRIPBinding) HasPath(path String, ) bool {
  classNameV := StringNameFromStr("OpenXRIPBinding")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRIPBinding) AddPath(path String, )  {
  classNameV := StringNameFromStr("OpenXRIPBinding")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OpenXRIPBinding) RemovePath(path String, )  {
  classNameV := StringNameFromStr("OpenXRIPBinding")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
