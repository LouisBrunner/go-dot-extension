// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Translation struct {
  obj gdc.ObjectPtr
}

func (me *Translation) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Translation) BaseClass() string {
  return "Translation"
}



// Enums

func (me *Translation) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Translation) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Translation) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Translation) SetLocale(locale String, )  {
  classNameV := StringNameFromStr("Translation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_locale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(locale.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Translation) GetLocale() String {
  classNameV := StringNameFromStr("Translation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_locale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Translation) AddMessage(src_message StringName, xlated_message StringName, context StringName, )  {
  classNameV := StringNameFromStr("Translation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_message")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 971803314) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(src_message.AsCTypePtr()), gdc.ConstTypePtr(xlated_message.AsCTypePtr()), gdc.ConstTypePtr(context.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Translation) AddPluralMessage(src_message StringName, xlated_messages PackedStringArray, context StringName, )  {
  classNameV := StringNameFromStr("Translation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_plural_message")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 360316719) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(src_message.AsCTypePtr()), gdc.ConstTypePtr(xlated_messages.AsCTypePtr()), gdc.ConstTypePtr(context.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Translation) GetMessage(src_message StringName, context StringName, ) StringName {
  classNameV := StringNameFromStr("Translation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_message")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 58037827) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(src_message.AsCTypePtr()), gdc.ConstTypePtr(context.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Translation) GetPluralMessage(src_message StringName, src_plural_message StringName, n int, context StringName, ) StringName {
  classNameV := StringNameFromStr("Translation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_plural_message")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1333931916) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(src_message.AsCTypePtr()), gdc.ConstTypePtr(src_plural_message.AsCTypePtr()), gdc.ConstTypePtr(&n), gdc.ConstTypePtr(context.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Translation) EraseMessage(src_message StringName, context StringName, )  {
  classNameV := StringNameFromStr("Translation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("erase_message")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3919944288) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(src_message.AsCTypePtr()), gdc.ConstTypePtr(context.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Translation) GetMessageList() PackedStringArray {
  classNameV := StringNameFromStr("Translation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_message_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Translation) GetTranslatedMessageList() PackedStringArray {
  classNameV := StringNameFromStr("Translation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_translated_message_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Translation) GetMessageCount() int {
  classNameV := StringNameFromStr("Translation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_message_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *Translation) GetPropMessages() Dictionary {
  panic("TODO: implement")
}

func (me *Translation) SetPropMessages(value Dictionary) {
  panic("TODO: implement")
}

func (me *Translation) GetPropLocale() String {
  panic("TODO: implement")
}

func (me *Translation) SetPropLocale(value String) {
  panic("TODO: implement")
}