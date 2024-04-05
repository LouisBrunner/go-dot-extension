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

type ptrsForOpenXRInteractionProfileList struct {
  fnSetInteractionProfilePath gdc.MethodBindPtr
  fnGetInteractionProfilePath gdc.MethodBindPtr
  fnGetBindingCount gdc.MethodBindPtr
  fnGetBinding gdc.MethodBindPtr
  fnSetBindings gdc.MethodBindPtr
  fnGetBindings gdc.MethodBindPtr
}

var ptrsForOpenXRInteractionProfile ptrsForOpenXRInteractionProfileList

func initOpenXRInteractionProfilePtrs(iface gdc.Interface) {

  className := StringNameFromStr("OpenXRInteractionProfile")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_interaction_profile_path")
    defer methodName.Destroy()
    ptrsForOpenXRInteractionProfile.fnSetInteractionProfilePath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_interaction_profile_path")
    defer methodName.Destroy()
    ptrsForOpenXRInteractionProfile.fnGetInteractionProfilePath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("get_binding_count")
    defer methodName.Destroy()
    ptrsForOpenXRInteractionProfile.fnGetBindingCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_binding")
    defer methodName.Destroy()
    ptrsForOpenXRInteractionProfile.fnGetBinding = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3934429652))
  }
  {
    methodName := StringNameFromStr("set_bindings")
    defer methodName.Destroy()
    ptrsForOpenXRInteractionProfile.fnSetBindings = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("get_bindings")
    defer methodName.Destroy()
    ptrsForOpenXRInteractionProfile.fnGetBindings = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
}

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
  cargs := []gdc.ConstTypePtr{interaction_profile_path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInteractionProfile.fnSetInteractionProfilePath), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OpenXRInteractionProfile) GetInteractionProfilePath() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInteractionProfile.fnGetInteractionProfilePath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *OpenXRInteractionProfile) GetBindingCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInteractionProfile.fnGetBindingCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OpenXRInteractionProfile) GetBinding(index int64, ) OpenXRIPBinding {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewOpenXRIPBinding()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInteractionProfile.fnGetBinding), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *OpenXRInteractionProfile) SetBindings(bindings Array, )  {
  cargs := []gdc.ConstTypePtr{bindings.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInteractionProfile.fnSetBindings), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OpenXRInteractionProfile) GetBindings() Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInteractionProfile.fnGetBindings), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
