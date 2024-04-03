// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Translation struct {
  Resource
}

func (me *Translation) BaseClass() string {
  return "Translation"
}

func NewTranslation() *Translation {
  str := StringNameFromStr("Translation") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Translation{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{}
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Translation) AddMessage(src_message StringName, xlated_message StringName, context StringName, )  {
  classNameV := StringNameFromStr("Translation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_message")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3898530326) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(src_message.AsCTypePtr()), gdc.ConstTypePtr(xlated_message.AsCTypePtr()), gdc.ConstTypePtr(context.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Translation) AddPluralMessage(src_message StringName, xlated_messages PackedStringArray, context StringName, )  {
  classNameV := StringNameFromStr("Translation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_plural_message")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2356982266) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(src_message.AsCTypePtr()), gdc.ConstTypePtr(xlated_messages.AsCTypePtr()), gdc.ConstTypePtr(context.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Translation) GetMessage(src_message StringName, context StringName, ) StringName {
  classNameV := StringNameFromStr("Translation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_message")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1829228469) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(src_message.AsCTypePtr()), gdc.ConstTypePtr(context.AsCTypePtr()), }
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Translation) GetPluralMessage(src_message StringName, src_plural_message StringName, n int64, context StringName, ) StringName {
  classNameV := StringNameFromStr("Translation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_plural_message")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 229954002) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(src_message.AsCTypePtr()), gdc.ConstTypePtr(src_plural_message.AsCTypePtr()), gdc.ConstTypePtr(&n), gdc.ConstTypePtr(context.AsCTypePtr()), }
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Translation) EraseMessage(src_message StringName, context StringName, )  {
  classNameV := StringNameFromStr("Translation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("erase_message")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3959009644) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(src_message.AsCTypePtr()), gdc.ConstTypePtr(context.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Translation) GetMessageList() PackedStringArray {
  classNameV := StringNameFromStr("Translation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_message_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Translation) GetTranslatedMessageList() PackedStringArray {
  classNameV := StringNameFromStr("Translation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_translated_message_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Translation) GetMessageCount() int64 {
  classNameV := StringNameFromStr("Translation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_message_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
