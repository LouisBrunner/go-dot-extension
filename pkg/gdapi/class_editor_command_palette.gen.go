// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorCommandPalette struct {
  obj gdc.ObjectPtr
}

func (me *EditorCommandPalette) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorCommandPalette) BaseClass() string {
  return "EditorCommandPalette"
}



// Enums

func (me *EditorCommandPalette) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorCommandPalette) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorCommandPalette) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EditorCommandPalette) AddCommand(command_name String, key_name String, binded_callable Callable, shortcut_text String, )  {
  classNameV := StringNameFromStr("EditorCommandPalette")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_command")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3664614892) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(command_name.AsCTypePtr()), gdc.ConstTypePtr(key_name.AsCTypePtr()), gdc.ConstTypePtr(binded_callable.AsCTypePtr()), gdc.ConstTypePtr(shortcut_text.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorCommandPalette) RemoveCommand(key_name String, )  {
  classNameV := StringNameFromStr("EditorCommandPalette")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_command")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(key_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties