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

type AnimationNodeBlendTree struct {
  AnimationRootNode
}

func (me *AnimationNodeBlendTree) BaseClass() string {
  return "AnimationNodeBlendTree"
}

func NewAnimationNodeBlendTree() *AnimationNodeBlendTree {
  str := StringNameFromStr("AnimationNodeBlendTree") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AnimationNodeBlendTree{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), node.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendTree) GetNode(name StringName, ) AnimationNode {
  classNameV := StringNameFromStr("AnimationNodeBlendTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 625644256) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAnimationNode()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeBlendTree) RemoveNode(name StringName, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendTree) RenameNode(name StringName, new_name StringName, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rename_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3740211285) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), new_name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendTree) HasNode(name StringName, ) bool {
  classNameV := StringNameFromStr("AnimationNodeBlendTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeBlendTree) ConnectNode(input_node StringName, input_index int64, output_node StringName, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("connect_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2168001410) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{input_node.AsCTypePtr(), gdc.ConstTypePtr(&input_index) , output_node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendTree) DisconnectNode(input_node StringName, input_index int64, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("disconnect_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2415702435) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{input_node.AsCTypePtr(), gdc.ConstTypePtr(&input_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendTree) SetNodePosition(name StringName, position Vector2, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_node_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1999414630) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendTree) GetNodePosition(name StringName, ) Vector2 {
  classNameV := StringNameFromStr("AnimationNodeBlendTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3100822709) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeBlendTree) SetGraphOffset(offset Vector2, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_graph_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendTree) GetGraphOffset() Vector2 {
  classNameV := StringNameFromStr("AnimationNodeBlendTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_graph_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type AnimationNodeBlendTreeNodeChangedSignalFn func(node_name StringName, )

func (me *AnimationNodeBlendTree) ConnectNodeChanged(subs SignalSubscribers, fn AnimationNodeBlendTreeNodeChangedSignalFn) {
  sig := StringNameFromStr("node_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationNodeBlendTree) DisconnectNodeChanged(subs SignalSubscribers, fn AnimationNodeBlendTreeNodeChangedSignalFn) {
  sig := StringNameFromStr("node_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
