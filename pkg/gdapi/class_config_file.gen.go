// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ConfigFile struct {
  obj gdc.ObjectPtr
}

func (me *ConfigFile) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ConfigFile) BaseClass() string {
  return "ConfigFile"
}



// Enums

func (me *ConfigFile) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ConfigFile) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ConfigFile) SetValue(section String, key String, value Variant, )  {
  panic("TODO: implement")
}

func  (me *ConfigFile) GetValue(section String, key String, default_ Variant, )  {
  panic("TODO: implement")
}

func  (me *ConfigFile) HasSection(section String, )  {
  panic("TODO: implement")
}

func  (me *ConfigFile) HasSectionKey(section String, key String, )  {
  panic("TODO: implement")
}

func  (me *ConfigFile) GetSections()  {
  panic("TODO: implement")
}

func  (me *ConfigFile) GetSectionKeys(section String, )  {
  panic("TODO: implement")
}

func  (me *ConfigFile) EraseSection(section String, )  {
  panic("TODO: implement")
}

func  (me *ConfigFile) EraseSectionKey(section String, key String, )  {
  panic("TODO: implement")
}

func  (me *ConfigFile) Load(path String, )  {
  panic("TODO: implement")
}

func  (me *ConfigFile) Parse(data String, )  {
  panic("TODO: implement")
}

func  (me *ConfigFile) Save(path String, )  {
  panic("TODO: implement")
}

func  (me *ConfigFile) EncodeToText()  {
  panic("TODO: implement")
}

func  (me *ConfigFile) LoadEncrypted(path String, key PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *ConfigFile) LoadEncryptedPass(path String, password String, )  {
  panic("TODO: implement")
}

func  (me *ConfigFile) SaveEncrypted(path String, key PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *ConfigFile) SaveEncryptedPass(path String, password String, )  {
  panic("TODO: implement")
}

func  (me *ConfigFile) Clear()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
