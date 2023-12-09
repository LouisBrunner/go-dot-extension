// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ProjectSettings struct {
  obj gdc.ObjectPtr
}

func (me *ProjectSettings) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ProjectSettings) BaseClass() string {
  return "ProjectSettings"
}



// Enums

func (me *ProjectSettings) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ProjectSettings) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ProjectSettings) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ProjectSettings) HasSetting(name String, )  {
  panic("TODO: implement")
}

func  (me *ProjectSettings) SetSetting(name String, value Variant, )  {
  panic("TODO: implement")
}

func  (me *ProjectSettings) GetSetting(name String, default_value Variant, )  {
  panic("TODO: implement")
}

func  (me *ProjectSettings) GetSettingWithOverride(name StringName, )  {
  panic("TODO: implement")
}

func  (me *ProjectSettings) GetGlobalClassList()  {
  panic("TODO: implement")
}

func  (me *ProjectSettings) SetOrder(name String, position int, )  {
  panic("TODO: implement")
}

func  (me *ProjectSettings) GetOrder(name String, )  {
  panic("TODO: implement")
}

func  (me *ProjectSettings) SetInitialValue(name String, value Variant, )  {
  panic("TODO: implement")
}

func  (me *ProjectSettings) SetAsBasic(name String, basic bool, )  {
  panic("TODO: implement")
}

func  (me *ProjectSettings) SetAsInternal(name String, internal bool, )  {
  panic("TODO: implement")
}

func  (me *ProjectSettings) AddPropertyInfo(hint Dictionary, )  {
  panic("TODO: implement")
}

func  (me *ProjectSettings) SetRestartIfChanged(name String, restart bool, )  {
  panic("TODO: implement")
}

func  (me *ProjectSettings) Clear(name String, )  {
  panic("TODO: implement")
}

func  (me *ProjectSettings) LocalizePath(path String, )  {
  panic("TODO: implement")
}

func  (me *ProjectSettings) GlobalizePath(path String, )  {
  panic("TODO: implement")
}

func  (me *ProjectSettings) Save()  {
  panic("TODO: implement")
}

func  (me *ProjectSettings) LoadResourcePack(pack String, replace_files bool, offset int, )  {
  panic("TODO: implement")
}

func  (me *ProjectSettings) SaveCustom(file String, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
