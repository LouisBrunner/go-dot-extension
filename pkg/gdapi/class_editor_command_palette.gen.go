// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForEditorCommandPaletteList struct {
  fnAddCommand gdc.MethodBindPtr
  fnRemoveCommand gdc.MethodBindPtr
}

var ptrsForEditorCommandPalette ptrsForEditorCommandPaletteList

func initEditorCommandPalettePtrs(iface gdc.Interface) {

  className := StringNameFromStr("EditorCommandPalette")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("add_command")
    defer methodName.Destroy()
    ptrsForEditorCommandPalette.fnAddCommand = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 864043298))
  }
  {
    methodName := StringNameFromStr("remove_command")
    defer methodName.Destroy()
    ptrsForEditorCommandPalette.fnRemoveCommand = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
}

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
  cargs := []gdc.ConstTypePtr{command_name.AsCTypePtr(), key_name.AsCTypePtr(), binded_callable.AsCTypePtr(), shortcut_text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorCommandPalette.fnAddCommand), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorCommandPalette) RemoveCommand(key_name String, )  {
  cargs := []gdc.ConstTypePtr{key_name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorCommandPalette.fnRemoveCommand), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
