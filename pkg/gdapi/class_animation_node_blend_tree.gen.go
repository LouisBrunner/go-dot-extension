// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AnimationNodeBlendTree struct {
  obj gdc.ObjectPtr
}

func (me *AnimationNodeBlendTree) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationNodeBlendTree) BaseClass() string {
  return "AnimationNodeBlendTree"
}



// Constants

var (
  AnimationNodeBlendTreeConnectionOk = "0" // TODO: construct correctly
  AnimationNodeBlendTreeConnectionErrorNoInput = "1" // TODO: construct correctly
  AnimationNodeBlendTreeConnectionErrorNoInputIndex = "2" // TODO: construct correctly
  AnimationNodeBlendTreeConnectionErrorNoOutput = "3" // TODO: construct correctly
  AnimationNodeBlendTreeConnectionErrorSameNode = "4" // TODO: construct correctly
  AnimationNodeBlendTreeConnectionErrorConnectionExists = "5" // TODO: construct correctly
)

// Enums

func (me *AnimationNodeBlendTree) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationNodeBlendTree) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeBlendTree) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AnimationNodeBlendTree) AddNode(name StringName, node AnimationNode, position Vector2, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1980270704) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(node.AsCTypePtr()), gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendTree) GetNode(name StringName, ) AnimationNode {
  classNameV := StringNameFromStr("AnimationNodeBlendTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 625644256) // FIXME: should cache?
  var ret AnimationNode
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeBlendTree) RemoveNode(name StringName, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendTree) RenameNode(name StringName, new_name StringName, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rename_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3740211285) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(new_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendTree) HasNode(name StringName, ) bool {
  classNameV := StringNameFromStr("AnimationNodeBlendTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeBlendTree) ConnectNode(input_node StringName, input_index int, output_node StringName, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("connect_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2168001410) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(input_node.AsCTypePtr()), gdc.ConstTypePtr(&input_index), gdc.ConstTypePtr(output_node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendTree) DisconnectNode(input_node StringName, input_index int, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("disconnect_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2415702435) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(input_node.AsCTypePtr()), gdc.ConstTypePtr(&input_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendTree) SetNodePosition(name StringName, position Vector2, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_node_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1999414630) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendTree) GetNodePosition(name StringName, ) Vector2 {
  classNameV := StringNameFromStr("AnimationNodeBlendTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3100822709) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeBlendTree) SetGraphOffset(offset Vector2, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_graph_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendTree) GetGraphOffset() Vector2 {
  classNameV := StringNameFromStr("AnimationNodeBlendTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_graph_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type AnimationNodeBlendTreeNodeChangedSignalFn func(node_name StringName, )

func (me *AnimationNodeBlendTree) ConnectNodeChanged(subs SignalSubscribers, fn AnimationNodeBlendTreeNodeChangedSignalFn) {
  sig := StringNameFromStr("node_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *AnimationNodeBlendTree) DisconnectNodeChanged(subs SignalSubscribers, fn AnimationNodeBlendTreeNodeChangedSignalFn) {
  sig := StringNameFromStr("node_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
