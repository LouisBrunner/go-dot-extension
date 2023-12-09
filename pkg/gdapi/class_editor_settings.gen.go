// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorSettings struct {
  obj gdc.ObjectPtr
}

func (me *EditorSettings) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorSettings) BaseClass() string {
  return "EditorSettings"
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

func  (me *EditorSettings) HasSetting(name String, )  {
  panic("TODO: implement")
}

func  (me *EditorSettings) SetSetting(name String, value Variant, )  {
  panic("TODO: implement")
}

func  (me *EditorSettings) GetSetting(name String, )  {
  panic("TODO: implement")
}

func  (me *EditorSettings) Erase(property String, )  {
  panic("TODO: implement")
}

func  (me *EditorSettings) SetInitialValue(name StringName, value Variant, update_current bool, )  {
  panic("TODO: implement")
}

func  (me *EditorSettings) AddPropertyInfo(info Dictionary, )  {
  panic("TODO: implement")
}

func  (me *EditorSettings) SetProjectMetadata(section String, key String, data Variant, )  {
  panic("TODO: implement")
}

func  (me *EditorSettings) GetProjectMetadata(section String, key String, default_ Variant, )  {
  panic("TODO: implement")
}

func  (me *EditorSettings) SetFavorites(dirs PackedStringArray, )  {
  panic("TODO: implement")
}

func  (me *EditorSettings) GetFavorites()  {
  panic("TODO: implement")
}

func  (me *EditorSettings) SetRecentDirs(dirs PackedStringArray, )  {
  panic("TODO: implement")
}

func  (me *EditorSettings) GetRecentDirs()  {
  panic("TODO: implement")
}

func  (me *EditorSettings) SetBuiltinActionOverride(name String, actions_list InputEvent, )  {
  panic("TODO: implement")
}

func  (me *EditorSettings) CheckChangedSettingsInGroup(setting_prefix String, )  {
  panic("TODO: implement")
}

func  (me *EditorSettings) GetChangedSettings()  {
  panic("TODO: implement")
}

func  (me *EditorSettings) MarkSettingChanged(setting String, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
