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

type EditorCommandPalette struct {
  ConfirmationDialog
}

func (me *EditorCommandPalette) BaseClass() string {
  return "EditorCommandPalette"
}

func NewEditorCommandPalette() *EditorCommandPalette {
  str := StringNameFromStr("EditorCommandPalette") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorCommandPalette{}
  obj.SetBaseObject(objPtr)
  return obj
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
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 864043298) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{command_name.AsCTypePtr(), key_name.AsCTypePtr(), binded_callable.AsCTypePtr(), shortcut_text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorCommandPalette) RemoveCommand(key_name String, )  {
  classNameV := StringNameFromStr("EditorCommandPalette")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_command")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{key_name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
