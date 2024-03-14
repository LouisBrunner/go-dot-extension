// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorResourcePicker struct {
  HBoxContainer
}

func (me *EditorResourcePicker) BaseClass() string {
  return "EditorResourcePicker"
}



// Enums

func (me *EditorResourcePicker) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorResourcePicker) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorResourcePicker) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EditorResourcePicker) SetBaseType(base_type String, )  {
  classNameV := StringNameFromStr("EditorResourcePicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_base_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(base_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorResourcePicker) GetBaseType() String {
  classNameV := StringNameFromStr("EditorResourcePicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_base_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorResourcePicker) GetAllowedTypes() PackedStringArray {
  classNameV := StringNameFromStr("EditorResourcePicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_allowed_types")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorResourcePicker) SetEditedResource(resource Resource, )  {
  classNameV := StringNameFromStr("EditorResourcePicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_edited_resource")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 968641751) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(resource.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorResourcePicker) GetEditedResource() Resource {
  classNameV := StringNameFromStr("EditorResourcePicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_edited_resource")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2674603643) // FIXME: should cache?
  var ret Resource
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorResourcePicker) SetToggleMode(enable bool, )  {
  classNameV := StringNameFromStr("EditorResourcePicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_toggle_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorResourcePicker) IsToggleMode() bool {
  classNameV := StringNameFromStr("EditorResourcePicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_toggle_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorResourcePicker) SetTogglePressed(pressed bool, )  {
  classNameV := StringNameFromStr("EditorResourcePicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_toggle_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressed), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorResourcePicker) SetEditable(enable bool, )  {
  classNameV := StringNameFromStr("EditorResourcePicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_editable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorResourcePicker) IsEditable() bool {
  classNameV := StringNameFromStr("EditorResourcePicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_editable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type EditorResourcePickerResourceSelectedSignalFn func(resource Resource, inspect bool, )

func (me *EditorResourcePicker) ConnectResourceSelected(subs SignalSubscribers, fn EditorResourcePickerResourceSelectedSignalFn) {
  sig := StringNameFromStr("resource_selected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *EditorResourcePicker) DisconnectResourceSelected(subs SignalSubscribers, fn EditorResourcePickerResourceSelectedSignalFn) {
  sig := StringNameFromStr("resource_selected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type EditorResourcePickerResourceChangedSignalFn func(resource Resource, )

func (me *EditorResourcePicker) ConnectResourceChanged(subs SignalSubscribers, fn EditorResourcePickerResourceChangedSignalFn) {
  sig := StringNameFromStr("resource_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *EditorResourcePicker) DisconnectResourceChanged(subs SignalSubscribers, fn EditorResourcePickerResourceChangedSignalFn) {
  sig := StringNameFromStr("resource_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
