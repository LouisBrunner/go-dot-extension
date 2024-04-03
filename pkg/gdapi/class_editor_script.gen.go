// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  classNameV := StringNameFromStr("EditorScript")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_root_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorScript) GetScene() Node {
  classNameV := StringNameFromStr("EditorScript")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3160264692) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewNode()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorScript) GetEditorInterface() EditorInterface {
  classNameV := StringNameFromStr("EditorScript")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_editor_interface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1976662476) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewEditorInterface()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
