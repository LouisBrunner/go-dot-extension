// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorSelection struct {
  obj gdc.ObjectPtr
}

func (me *EditorSelection) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorSelection) BaseClass() string {
  return "EditorSelection"
}



// Enums

func (me *EditorSelection) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorSelection) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorSelection) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EditorSelection) Clear()  {
  classNameV := StringNameFromStr("EditorSelection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorSelection) AddNode(node Node, )  {
  classNameV := StringNameFromStr("EditorSelection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorSelection) RemoveNode(node Node, )  {
  classNameV := StringNameFromStr("EditorSelection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorSelection) GetSelectedNodes() Node {
  classNameV := StringNameFromStr("EditorSelection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selected_nodes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret Node
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorSelection) GetTransformableSelectedNodes() Node {
  classNameV := StringNameFromStr("EditorSelection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_transformable_selected_nodes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret Node
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties
// Signals
// FIXME: can't seem to be able to connect them from this side of the API