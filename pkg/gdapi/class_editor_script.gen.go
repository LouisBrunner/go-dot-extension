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

type ptrsForEditorScriptList struct {
  fnXRun gdc.MethodBindPtr
  fnAddRootNode gdc.MethodBindPtr
  fnGetScene gdc.MethodBindPtr
  fnGetEditorInterface gdc.MethodBindPtr
}

var ptrsForEditorScript ptrsForEditorScriptList

func initEditorScriptPtrs(iface gdc.Interface) {

  className := StringNameFromStr("EditorScript")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("add_root_node")
    defer methodName.Destroy()
    ptrsForEditorScript.fnAddRootNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1078189570))
  }
  {
    methodName := StringNameFromStr("get_scene")
    defer methodName.Destroy()
    ptrsForEditorScript.fnGetScene = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3160264692))
  }
  {
    methodName := StringNameFromStr("get_editor_interface")
    defer methodName.Destroy()
    ptrsForEditorScript.fnGetEditorInterface = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1976662476))
  }
}

type EditorScript struct {
  RefCounted
}

func (me *EditorScript) BaseClass() string {
  return "EditorScript"
}

func NewEditorScript() *EditorScript {
  str := StringNameFromStr("EditorScript") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorScript{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *EditorScript) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorScript) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorScript) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EditorScript) AddRootNode(node Node, )  {
  cargs := []gdc.ConstTypePtr{node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorScript.fnAddRootNode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorScript) GetScene() Node {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNode()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorScript.fnGetScene), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorScript) GetEditorInterface() EditorInterface {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewEditorInterface()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorScript.fnGetEditorInterface), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
