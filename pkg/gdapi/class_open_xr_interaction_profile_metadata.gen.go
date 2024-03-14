// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type OpenXRInteractionProfileMetadata struct {
  Object
}

func (me *OpenXRInteractionProfileMetadata) BaseClass() string {
  return "OpenXRInteractionProfileMetadata"
}



// Enums

func (me *OpenXRInteractionProfileMetadata) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *OpenXRInteractionProfileMetadata) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *OpenXRInteractionProfileMetadata) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *OpenXRInteractionProfileMetadata) RegisterProfileRename(old_name String, new_name String, )  {
  classNameV := StringNameFromStr("OpenXRInteractionProfileMetadata")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("register_profile_rename")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3186203200) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(old_name.AsCTypePtr()), gdc.ConstTypePtr(new_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OpenXRInteractionProfileMetadata) RegisterTopLevelPath(display_name String, openxr_path String, openxr_extension_name String, )  {
  classNameV := StringNameFromStr("OpenXRInteractionProfileMetadata")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("register_top_level_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 254767734) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(display_name.AsCTypePtr()), gdc.ConstTypePtr(openxr_path.AsCTypePtr()), gdc.ConstTypePtr(openxr_extension_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OpenXRInteractionProfileMetadata) RegisterInteractionProfile(display_name String, openxr_path String, openxr_extension_name String, )  {
  classNameV := StringNameFromStr("OpenXRInteractionProfileMetadata")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("register_interaction_profile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 254767734) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(display_name.AsCTypePtr()), gdc.ConstTypePtr(openxr_path.AsCTypePtr()), gdc.ConstTypePtr(openxr_extension_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OpenXRInteractionProfileMetadata) RegisterIoPath(interaction_profile String, display_name String, toplevel_path String, openxr_path String, openxr_extension_name String, action_type OpenXRActionActionType, )  {
  classNameV := StringNameFromStr("OpenXRInteractionProfileMetadata")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("register_io_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3443511926) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(interaction_profile.AsCTypePtr()), gdc.ConstTypePtr(display_name.AsCTypePtr()), gdc.ConstTypePtr(toplevel_path.AsCTypePtr()), gdc.ConstTypePtr(openxr_path.AsCTypePtr()), gdc.ConstTypePtr(openxr_extension_name.AsCTypePtr()), gdc.ConstTypePtr(&action_type), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Signals
