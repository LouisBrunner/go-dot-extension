// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type OpenXRActionMap struct {
  obj gdc.ObjectPtr
}

func (me *OpenXRActionMap) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *OpenXRActionMap) BaseClass() string {
  return "OpenXRActionMap"
}



// Enums

func (me *OpenXRActionMap) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *OpenXRActionMap) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *OpenXRActionMap) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *OpenXRActionMap) SetActionSets(action_sets Array, )  {
  classNameV := StringNameFromStr("OpenXRActionMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_action_sets")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(action_sets.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OpenXRActionMap) GetActionSets() Array {
  classNameV := StringNameFromStr("OpenXRActionMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_action_sets")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRActionMap) GetActionSetCount() int {
  classNameV := StringNameFromStr("OpenXRActionMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_action_set_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRActionMap) FindActionSet(name String, ) OpenXRActionSet {
  classNameV := StringNameFromStr("OpenXRActionMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_action_set")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1888809267) // FIXME: should cache?
  var ret OpenXRActionSet
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRActionMap) GetActionSet(idx int, ) OpenXRActionSet {
  classNameV := StringNameFromStr("OpenXRActionMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_action_set")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1789580336) // FIXME: should cache?
  var ret OpenXRActionSet
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRActionMap) AddActionSet(action_set OpenXRActionSet, )  {
  classNameV := StringNameFromStr("OpenXRActionMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_action_set")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2093310581) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(action_set.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OpenXRActionMap) RemoveActionSet(action_set OpenXRActionSet, )  {
  classNameV := StringNameFromStr("OpenXRActionMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_action_set")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2093310581) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(action_set.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OpenXRActionMap) SetInteractionProfiles(interaction_profiles Array, )  {
  classNameV := StringNameFromStr("OpenXRActionMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_interaction_profiles")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(interaction_profiles.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OpenXRActionMap) GetInteractionProfiles() Array {
  classNameV := StringNameFromStr("OpenXRActionMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_interaction_profiles")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRActionMap) GetInteractionProfileCount() int {
  classNameV := StringNameFromStr("OpenXRActionMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_interaction_profile_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRActionMap) FindInteractionProfile(name String, ) OpenXRInteractionProfile {
  classNameV := StringNameFromStr("OpenXRActionMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_interaction_profile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3095875538) // FIXME: should cache?
  var ret OpenXRInteractionProfile
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRActionMap) GetInteractionProfile(idx int, ) OpenXRInteractionProfile {
  classNameV := StringNameFromStr("OpenXRActionMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_interaction_profile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2546151210) // FIXME: should cache?
  var ret OpenXRInteractionProfile
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRActionMap) AddInteractionProfile(interaction_profile OpenXRInteractionProfile, )  {
  classNameV := StringNameFromStr("OpenXRActionMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_interaction_profile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2697953512) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(interaction_profile.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OpenXRActionMap) RemoveInteractionProfile(interaction_profile OpenXRInteractionProfile, )  {
  classNameV := StringNameFromStr("OpenXRActionMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_interaction_profile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2697953512) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(interaction_profile.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OpenXRActionMap) CreateDefaultActionSets()  {
  classNameV := StringNameFromStr("OpenXRActionMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_default_action_sets")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
