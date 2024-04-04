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

type OpenXRInteractionProfile struct {
  Resource
}

func (me *OpenXRInteractionProfile) BaseClass() string {
  return "OpenXRInteractionProfile"
}

func NewOpenXRInteractionProfile() *OpenXRInteractionProfile {
  str := StringNameFromStr("OpenXRInteractionProfile") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &OpenXRInteractionProfile{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *OpenXRInteractionProfile) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *OpenXRInteractionProfile) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *OpenXRInteractionProfile) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *OpenXRInteractionProfile) SetInteractionProfilePath(interaction_profile_path String, )  {
  classNameV := StringNameFromStr("OpenXRInteractionProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_interaction_profile_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{interaction_profile_path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OpenXRInteractionProfile) GetInteractionProfilePath() String {
  classNameV := StringNameFromStr("OpenXRInteractionProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_interaction_profile_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *OpenXRInteractionProfile) GetBindingCount() int64 {
  classNameV := StringNameFromStr("OpenXRInteractionProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_binding_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OpenXRInteractionProfile) GetBinding(index int64, ) OpenXRIPBinding {
  classNameV := StringNameFromStr("OpenXRInteractionProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_binding")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3934429652) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewOpenXRIPBinding()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *OpenXRInteractionProfile) SetBindings(bindings Array, )  {
  classNameV := StringNameFromStr("OpenXRInteractionProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bindings")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{bindings.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OpenXRInteractionProfile) GetBindings() Array {
  classNameV := StringNameFromStr("OpenXRInteractionProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bindings")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
