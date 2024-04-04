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

type AnimationTree struct {
  AnimationMixer
}

func (me *AnimationTree) BaseClass() string {
  return "AnimationTree"
}

func NewAnimationTree() *AnimationTree {
  str := StringNameFromStr("AnimationTree") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AnimationTree{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type AnimationTreeAnimationProcessCallback int
const (
  AnimationTreeAnimationProcessCallbackAnimationProcessPhysics AnimationTreeAnimationProcessCallback = 0
  AnimationTreeAnimationProcessCallbackAnimationProcessIdle AnimationTreeAnimationProcessCallback = 1
  AnimationTreeAnimationProcessCallbackAnimationProcessManual AnimationTreeAnimationProcessCallback = 2
)

func (me *AnimationTree) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationTree) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationTree) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AnimationTree) SetTreeRoot(animation_node AnimationRootNode, )  {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tree_root")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2581683800) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{animation_node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationTree) GetTreeRoot() AnimationRootNode {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tree_root")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4110384712) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAnimationRootNode()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationTree) SetAdvanceExpressionBaseNode(path NodePath, )  {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_advance_expression_base_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationTree) GetAdvanceExpressionBaseNode() NodePath {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_advance_expression_base_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationTree) SetAnimationPlayer(path NodePath, )  {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_animation_player")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationTree) GetAnimationPlayer() NodePath {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_animation_player")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationTree) SetProcessCallback(mode AnimationTreeAnimationProcessCallback, )  {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_process_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1723352826) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationTree) GetProcessCallback() AnimationTreeAnimationProcessCallback {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_process_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 891317132) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AnimationTreeAnimationProcessCallback

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type AnimationTreeAnimationPlayerChangedSignalFn func()

func (me *AnimationTree) ConnectAnimationPlayerChanged(subs SignalSubscribers, fn AnimationTreeAnimationPlayerChangedSignalFn) {
  sig := StringNameFromStr("animation_player_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationTree) DisconnectAnimationPlayerChanged(subs SignalSubscribers, fn AnimationTreeAnimationPlayerChangedSignalFn) {
  sig := StringNameFromStr("animation_player_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
