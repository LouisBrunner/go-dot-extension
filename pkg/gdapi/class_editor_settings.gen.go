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

type EditorSettings struct {
  Resource
}

func (me *EditorSettings) BaseClass() string {
  return "EditorSettings"
}

func NewEditorSettings() *EditorSettings {
  str := StringNameFromStr("EditorSettings") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorSettings{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Constants

var (
  EditorSettingsNotificationEditorSettingsChanged = "10000" // TODO: construct correctly
)

// Enums

func (me *EditorSettings) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorSettings) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorSettings) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EditorSettings) HasSetting(name String, ) bool {
  classNameV := StringNameFromStr("EditorSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_setting")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorSettings) SetSetting(name String, value Variant, )  {
  classNameV := StringNameFromStr("EditorSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_setting")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 402577236) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorSettings) GetSetting(name String, ) Variant {
  classNameV := StringNameFromStr("EditorSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_setting")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1868160156) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorSettings) Erase(property String, )  {
  classNameV := StringNameFromStr("EditorSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("erase")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{property.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorSettings) SetInitialValue(name StringName, value Variant, update_current bool, )  {
  classNameV := StringNameFromStr("EditorSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_initial_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1529169264) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), value.AsCTypePtr(), gdc.ConstTypePtr(&update_current) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorSettings) AddPropertyInfo(info Dictionary, )  {
  classNameV := StringNameFromStr("EditorSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_property_info")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155329257) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{info.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorSettings) SetProjectMetadata(section String, key String, data Variant, )  {
  classNameV := StringNameFromStr("EditorSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_project_metadata")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2504492430) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{section.AsCTypePtr(), key.AsCTypePtr(), data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorSettings) GetProjectMetadata(section String, key String, default_ Variant, ) Variant {
  classNameV := StringNameFromStr("EditorSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_project_metadata")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 89809366) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{section.AsCTypePtr(), key.AsCTypePtr(), default_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorSettings) SetFavorites(dirs PackedStringArray, )  {
  classNameV := StringNameFromStr("EditorSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_favorites")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4015028928) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{dirs.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorSettings) GetFavorites() PackedStringArray {
  classNameV := StringNameFromStr("EditorSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_favorites")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorSettings) SetRecentDirs(dirs PackedStringArray, )  {
  classNameV := StringNameFromStr("EditorSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_recent_dirs")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4015028928) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{dirs.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorSettings) GetRecentDirs() PackedStringArray {
  classNameV := StringNameFromStr("EditorSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_recent_dirs")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorSettings) SetBuiltinActionOverride(name String, actions_list []InputEvent, )  {
  classNameV := StringNameFromStr("EditorSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_builtin_action_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1209351045) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&actions_list) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorSettings) CheckChangedSettingsInGroup(setting_prefix String, ) bool {
  classNameV := StringNameFromStr("EditorSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("check_changed_settings_in_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{setting_prefix.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorSettings) GetChangedSettings() PackedStringArray {
  classNameV := StringNameFromStr("EditorSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_changed_settings")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorSettings) MarkSettingChanged(setting String, )  {
  classNameV := StringNameFromStr("EditorSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("mark_setting_changed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{setting.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals

type EditorSettingsSettingsChangedSignalFn func()

func (me *EditorSettings) ConnectSettingsChanged(subs SignalSubscribers, fn EditorSettingsSettingsChangedSignalFn) {
  sig := StringNameFromStr("settings_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorSettings) DisconnectSettingsChanged(subs SignalSubscribers, fn EditorSettingsSettingsChangedSignalFn) {
  sig := StringNameFromStr("settings_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
