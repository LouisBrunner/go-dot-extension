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

type ptrsForAnimationNodeStateMachineList struct {
  fnAddNode gdc.MethodBindPtr
  fnReplaceNode gdc.MethodBindPtr
  fnGetNode gdc.MethodBindPtr
  fnRemoveNode gdc.MethodBindPtr
  fnRenameNode gdc.MethodBindPtr
  fnHasNode gdc.MethodBindPtr
  fnGetNodeName gdc.MethodBindPtr
  fnSetNodePosition gdc.MethodBindPtr
  fnGetNodePosition gdc.MethodBindPtr
  fnHasTransition gdc.MethodBindPtr
  fnAddTransition gdc.MethodBindPtr
  fnGetTransition gdc.MethodBindPtr
  fnGetTransitionFrom gdc.MethodBindPtr
  fnGetTransitionTo gdc.MethodBindPtr
  fnGetTransitionCount gdc.MethodBindPtr
  fnRemoveTransitionByIndex gdc.MethodBindPtr
  fnRemoveTransition gdc.MethodBindPtr
  fnSetGraphOffset gdc.MethodBindPtr
  fnGetGraphOffset gdc.MethodBindPtr
  fnSetStateMachineType gdc.MethodBindPtr
  fnGetStateMachineType gdc.MethodBindPtr
  fnSetAllowTransitionToSelf gdc.MethodBindPtr
  fnIsAllowTransitionToSelf gdc.MethodBindPtr
  fnSetResetEnds gdc.MethodBindPtr
  fnAreEndsReset gdc.MethodBindPtr
}

var ptrsForAnimationNodeStateMachine ptrsForAnimationNodeStateMachineList

func initAnimationNodeStateMachinePtrs(iface gdc.Interface) {

  className := StringNameFromStr("AnimationNodeStateMachine")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("add_node")
    defer methodName.Destroy()
    ptrsForAnimationNodeStateMachine.fnAddNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1980270704))
  }
  {
    methodName := StringNameFromStr("replace_node")
    defer methodName.Destroy()
    ptrsForAnimationNodeStateMachine.fnReplaceNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2559412862))
  }
  {
    methodName := StringNameFromStr("get_node")
    defer methodName.Destroy()
    ptrsForAnimationNodeStateMachine.fnGetNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 625644256))
  }
  {
    methodName := StringNameFromStr("remove_node")
    defer methodName.Destroy()
    ptrsForAnimationNodeStateMachine.fnRemoveNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
  }
  {
    methodName := StringNameFromStr("rename_node")
    defer methodName.Destroy()
    ptrsForAnimationNodeStateMachine.fnRenameNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3740211285))
  }
  {
    methodName := StringNameFromStr("has_node")
    defer methodName.Destroy()
    ptrsForAnimationNodeStateMachine.fnHasNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
  }
  {
    methodName := StringNameFromStr("get_node_name")
    defer methodName.Destroy()
    ptrsForAnimationNodeStateMachine.fnGetNodeName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 739213945))
  }
  {
    methodName := StringNameFromStr("set_node_position")
    defer methodName.Destroy()
    ptrsForAnimationNodeStateMachine.fnSetNodePosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1999414630))
  }
  {
    methodName := StringNameFromStr("get_node_position")
    defer methodName.Destroy()
    ptrsForAnimationNodeStateMachine.fnGetNodePosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3100822709))
  }
  {
    methodName := StringNameFromStr("has_transition")
    defer methodName.Destroy()
    ptrsForAnimationNodeStateMachine.fnHasTransition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 471820014))
  }
  {
    methodName := StringNameFromStr("add_transition")
    defer methodName.Destroy()
    ptrsForAnimationNodeStateMachine.fnAddTransition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 795486887))
  }
  {
    methodName := StringNameFromStr("get_transition")
    defer methodName.Destroy()
    ptrsForAnimationNodeStateMachine.fnGetTransition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4192381260))
  }
  {
    methodName := StringNameFromStr("get_transition_from")
    defer methodName.Destroy()
    ptrsForAnimationNodeStateMachine.fnGetTransitionFrom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 659327637))
  }
  {
    methodName := StringNameFromStr("get_transition_to")
    defer methodName.Destroy()
    ptrsForAnimationNodeStateMachine.fnGetTransitionTo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 659327637))
  }
  {
    methodName := StringNameFromStr("get_transition_count")
    defer methodName.Destroy()
    ptrsForAnimationNodeStateMachine.fnGetTransitionCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("remove_transition_by_index")
    defer methodName.Destroy()
    ptrsForAnimationNodeStateMachine.fnRemoveTransitionByIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("remove_transition")
    defer methodName.Destroy()
    ptrsForAnimationNodeStateMachine.fnRemoveTransition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3740211285))
  }
  {
    methodName := StringNameFromStr("set_graph_offset")
    defer methodName.Destroy()
    ptrsForAnimationNodeStateMachine.fnSetGraphOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("get_graph_offset")
    defer methodName.Destroy()
    ptrsForAnimationNodeStateMachine.fnGetGraphOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
  {
    methodName := StringNameFromStr("set_state_machine_type")
    defer methodName.Destroy()
    ptrsForAnimationNodeStateMachine.fnSetStateMachineType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2584759088))
  }
  {
    methodName := StringNameFromStr("get_state_machine_type")
    defer methodName.Destroy()
    ptrsForAnimationNodeStateMachine.fnGetStateMachineType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1140726469))
  }
  {
    methodName := StringNameFromStr("set_allow_transition_to_self")
    defer methodName.Destroy()
    ptrsForAnimationNodeStateMachine.fnSetAllowTransitionToSelf = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_allow_transition_to_self")
    defer methodName.Destroy()
    ptrsForAnimationNodeStateMachine.fnIsAllowTransitionToSelf = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_reset_ends")
    defer methodName.Destroy()
    ptrsForAnimationNodeStateMachine.fnSetResetEnds = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("are_ends_reset")
    defer methodName.Destroy()
    ptrsForAnimationNodeStateMachine.fnAreEndsReset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

type AnimationNodeStateMachine struct {
  AnimationRootNode
}

func (me *AnimationNodeStateMachine) BaseClass() string {
  return "AnimationNodeStateMachine"
}

func NewAnimationNodeStateMachine() *AnimationNodeStateMachine {
  str := StringNameFromStr("AnimationNodeStateMachine") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AnimationNodeStateMachine{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type AnimationNodeStateMachineStateMachineType int
const (
  AnimationNodeStateMachineStateMachineTypeStateMachineTypeRoot AnimationNodeStateMachineStateMachineType = 0
  AnimationNodeStateMachineStateMachineTypeStateMachineTypeNested AnimationNodeStateMachineStateMachineType = 1
  AnimationNodeStateMachineStateMachineTypeStateMachineTypeGrouped AnimationNodeStateMachineStateMachineType = 2
)

func (me *AnimationNodeStateMachine) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationNodeStateMachine) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeStateMachine) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AnimationNodeStateMachine) AddNode(name StringName, node AnimationNode, position Vector2, )  {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), node.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachine.fnAddNode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachine) ReplaceNode(name StringName, node AnimationNode, )  {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachine.fnReplaceNode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachine) GetNode(name StringName, ) AnimationNode {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAnimationNode()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachine.fnGetNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeStateMachine) RemoveNode(name StringName, )  {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachine.fnRemoveNode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachine) RenameNode(name StringName, new_name StringName, )  {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), new_name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachine.fnRenameNode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachine) HasNode(name StringName, ) bool {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachine.fnHasNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeStateMachine) GetNodeName(node AnimationNode, ) StringName {
  cargs := []gdc.ConstTypePtr{node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachine.fnGetNodeName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeStateMachine) SetNodePosition(name StringName, position Vector2, )  {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachine.fnSetNodePosition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachine) GetNodePosition(name StringName, ) Vector2 {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachine.fnGetNodePosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeStateMachine) HasTransition(from StringName, to StringName, ) bool {
  cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachine.fnHasTransition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeStateMachine) AddTransition(from StringName, to StringName, transition AnimationNodeStateMachineTransition, )  {
  cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr(), transition.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachine.fnAddTransition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachine) GetTransition(idx int64, ) AnimationNodeStateMachineTransition {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAnimationNodeStateMachineTransition()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachine.fnGetTransition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeStateMachine) GetTransitionFrom(idx int64, ) StringName {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachine.fnGetTransitionFrom), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeStateMachine) GetTransitionTo(idx int64, ) StringName {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachine.fnGetTransitionTo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeStateMachine) GetTransitionCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachine.fnGetTransitionCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeStateMachine) RemoveTransitionByIndex(idx int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachine.fnRemoveTransitionByIndex), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachine) RemoveTransition(from StringName, to StringName, )  {
  cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachine.fnRemoveTransition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachine) SetGraphOffset(offset Vector2, )  {
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachine.fnSetGraphOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachine) GetGraphOffset() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachine.fnGetGraphOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeStateMachine) SetStateMachineType(state_machine_type AnimationNodeStateMachineStateMachineType, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&state_machine_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachine.fnSetStateMachineType), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachine) GetStateMachineType() AnimationNodeStateMachineStateMachineType {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AnimationNodeStateMachineStateMachineType

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachine.fnGetStateMachineType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *AnimationNodeStateMachine) SetAllowTransitionToSelf(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachine.fnSetAllowTransitionToSelf), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachine) IsAllowTransitionToSelf() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachine.fnIsAllowTransitionToSelf), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeStateMachine) SetResetEnds(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachine.fnSetResetEnds), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachine) AreEndsReset() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachine.fnAreEndsReset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
