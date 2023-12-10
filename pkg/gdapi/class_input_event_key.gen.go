// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type InputEventKey struct {
  obj gdc.ObjectPtr
}

func (me *InputEventKey) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *InputEventKey) BaseClass() string {
  return "InputEventKey"
}



// Enums

func (me *InputEventKey) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *InputEventKey) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputEventKey) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *InputEventKey) SetPressed(pressed bool, )  {
  classNameV := StringNameFromStr("InputEventKey")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressed), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventKey) SetKeycode(keycode Key, )  {
  classNameV := StringNameFromStr("InputEventKey")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_keycode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 888074362) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&keycode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventKey) GetKeycode() Key {
  classNameV := StringNameFromStr("InputEventKey")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_keycode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1585896689) // FIXME: should cache?
  var ret Key
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEventKey) SetPhysicalKeycode(physical_keycode Key, )  {
  classNameV := StringNameFromStr("InputEventKey")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_physical_keycode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 888074362) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&physical_keycode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventKey) GetPhysicalKeycode() Key {
  classNameV := StringNameFromStr("InputEventKey")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_physical_keycode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1585896689) // FIXME: should cache?
  var ret Key
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEventKey) SetKeyLabel(key_label Key, )  {
  classNameV := StringNameFromStr("InputEventKey")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_key_label")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 888074362) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&key_label), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventKey) GetKeyLabel() Key {
  classNameV := StringNameFromStr("InputEventKey")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_key_label")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1585896689) // FIXME: should cache?
  var ret Key
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEventKey) SetUnicode(unicode int, )  {
  classNameV := StringNameFromStr("InputEventKey")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_unicode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unicode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventKey) GetUnicode() int {
  classNameV := StringNameFromStr("InputEventKey")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_unicode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEventKey) SetEcho(echo bool, )  {
  classNameV := StringNameFromStr("InputEventKey")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_echo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&echo), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventKey) GetKeycodeWithModifiers() Key {
  classNameV := StringNameFromStr("InputEventKey")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_keycode_with_modifiers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1585896689) // FIXME: should cache?
  var ret Key
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEventKey) GetPhysicalKeycodeWithModifiers() Key {
  classNameV := StringNameFromStr("InputEventKey")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_physical_keycode_with_modifiers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1585896689) // FIXME: should cache?
  var ret Key
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEventKey) GetKeyLabelWithModifiers() Key {
  classNameV := StringNameFromStr("InputEventKey")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_key_label_with_modifiers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1585896689) // FIXME: should cache?
  var ret Key
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEventKey) AsTextKeycode() String {
  classNameV := StringNameFromStr("InputEventKey")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("as_text_keycode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEventKey) AsTextPhysicalKeycode() String {
  classNameV := StringNameFromStr("InputEventKey")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("as_text_physical_keycode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEventKey) AsTextKeyLabel() String {
  classNameV := StringNameFromStr("InputEventKey")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("as_text_key_label")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *InputEventKey) GetPropPressed() bool {
  panic("TODO: implement")
}

func (me *InputEventKey) SetPropPressed(value bool) {
  panic("TODO: implement")
}

func (me *InputEventKey) GetPropKeycode() int {
  panic("TODO: implement")
}

func (me *InputEventKey) SetPropKeycode(value int) {
  panic("TODO: implement")
}

func (me *InputEventKey) GetPropPhysicalKeycode() int {
  panic("TODO: implement")
}

func (me *InputEventKey) SetPropPhysicalKeycode(value int) {
  panic("TODO: implement")
}

func (me *InputEventKey) GetPropKeyLabel() int {
  panic("TODO: implement")
}

func (me *InputEventKey) SetPropKeyLabel(value int) {
  panic("TODO: implement")
}

func (me *InputEventKey) GetPropUnicode() int {
  panic("TODO: implement")
}

func (me *InputEventKey) SetPropUnicode(value int) {
  panic("TODO: implement")
}

func (me *InputEventKey) GetPropEcho() bool {
  panic("TODO: implement")
}

func (me *InputEventKey) SetPropEcho(value bool) {
  panic("TODO: implement")
}