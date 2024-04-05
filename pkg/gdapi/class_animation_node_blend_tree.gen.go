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

type ptrsForAnimationNodeBlendTreeList struct {
  fnAddNode gdc.MethodBindPtr
  fnGetNode gdc.MethodBindPtr
  fnRemoveNode gdc.MethodBindPtr
  fnRenameNode gdc.MethodBindPtr
  fnHasNode gdc.MethodBindPtr
  fnConnectNode gdc.MethodBindPtr
  fnDisconnectNode gdc.MethodBindPtr
  fnSetNodePosition gdc.MethodBindPtr
  fnGetNodePosition gdc.MethodBindPtr
  fnSetGraphOffset gdc.MethodBindPtr
  fnGetGraphOffset gdc.MethodBindPtr
}

var ptrsForAnimationNodeBlendTree ptrsForAnimationNodeBlendTreeList

func initAnimationNodeBlendTreePtrs(iface gdc.Interface) {

  className := StringNameFromStr("AnimationNodeBlendTree")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("add_node")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendTree.fnAddNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1980270704))
  }
  {
    methodName := StringNameFromStr("get_node")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendTree.fnGetNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 625644256))
  }
  {
    methodName := StringNameFromStr("remove_node")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendTree.fnRemoveNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
  }
  {
    methodName := StringNameFromStr("rename_node")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendTree.fnRenameNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3740211285))
  }
  {
    methodName := StringNameFromStr("has_node")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendTree.fnHasNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
  }
  {
    methodName := StringNameFromStr("connect_node")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendTree.fnConnectNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2168001410))
  }
  {
    methodName := StringNameFromStr("disconnect_node")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendTree.fnDisconnectNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2415702435))
  }
  {
    methodName := StringNameFromStr("set_node_position")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendTree.fnSetNodePosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1999414630))
  }
  {
    methodName := StringNameFromStr("get_node_position")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendTree.fnGetNodePosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3100822709))
  }
  {
    methodName := StringNameFromStr("set_graph_offset")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendTree.fnSetGraphOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("get_graph_offset")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendTree.fnGetGraphOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
}

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
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), node.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendTree.fnAddNode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendTree) GetNode(name StringName, ) AnimationNode {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAnimationNode()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendTree.fnGetNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeBlendTree) RemoveNode(name StringName, )  {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendTree.fnRemoveNode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendTree) RenameNode(name StringName, new_name StringName, )  {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), new_name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendTree.fnRenameNode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendTree) HasNode(name StringName, ) bool {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendTree.fnHasNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeBlendTree) ConnectNode(input_node StringName, input_index int64, output_node StringName, )  {
  cargs := []gdc.ConstTypePtr{input_node.AsCTypePtr(), gdc.ConstTypePtr(&input_index) , output_node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendTree.fnConnectNode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendTree) DisconnectNode(input_node StringName, input_index int64, )  {
  cargs := []gdc.ConstTypePtr{input_node.AsCTypePtr(), gdc.ConstTypePtr(&input_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendTree.fnDisconnectNode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendTree) SetNodePosition(name StringName, position Vector2, )  {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendTree.fnSetNodePosition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendTree) GetNodePosition(name StringName, ) Vector2 {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendTree.fnGetNodePosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeBlendTree) SetGraphOffset(offset Vector2, )  {
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendTree.fnSetGraphOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendTree) GetGraphOffset() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendTree.fnGetGraphOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
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
