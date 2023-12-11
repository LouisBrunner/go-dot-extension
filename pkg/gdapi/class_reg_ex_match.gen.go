// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type RegExMatch struct {
  obj gdc.ObjectPtr
}

func (me *RegExMatch) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RegExMatch) BaseClass() string {
  return "RegExMatch"
}



// Enums

func (me *RegExMatch) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RegExMatch) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RegExMatch) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RegExMatch) GetSubject() String {
  classNameV := StringNameFromStr("RegExMatch")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_subject")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RegExMatch) GetGroupCount() int {
  classNameV := StringNameFromStr("RegExMatch")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_group_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RegExMatch) GetNames() Dictionary {
  classNameV := StringNameFromStr("RegExMatch")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_names")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3102165223) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RegExMatch) GetStrings() PackedStringArray {
  classNameV := StringNameFromStr("RegExMatch")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_strings")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RegExMatch) GetString(name Variant, ) String {
  classNameV := StringNameFromStr("RegExMatch")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 687115856) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RegExMatch) GetStart(name Variant, ) int {
  classNameV := StringNameFromStr("RegExMatch")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_start")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 490464691) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RegExMatch) GetEnd(name Variant, ) int {
  classNameV := StringNameFromStr("RegExMatch")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 490464691) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
