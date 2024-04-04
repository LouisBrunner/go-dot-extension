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

type EditorScriptPicker struct {
  EditorResourcePicker
}

func (me *EditorScriptPicker) BaseClass() string {
  return "EditorScriptPicker"
}

func NewEditorScriptPicker() *EditorScriptPicker {
  str := StringNameFromStr("EditorScriptPicker") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorScriptPicker{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *EditorScriptPicker) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorScriptPicker) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorScriptPicker) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EditorScriptPicker) SetScriptOwner(owner_node Node, )  {
  classNameV := StringNameFromStr("EditorScriptPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_script_owner")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{owner_node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorScriptPicker) GetScriptOwner() Node {
  classNameV := StringNameFromStr("EditorScriptPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_script_owner")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3160264692) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNode()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
