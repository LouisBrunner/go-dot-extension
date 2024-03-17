// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AnimationNode struct {
  Resource
}

func (me *AnimationNode) BaseClass() string {
  return "AnimationNode"
}



// Enums

type AnimationNodeFilterAction int
const (
  AnimationNodeFilterActionFilterIgnore AnimationNodeFilterAction = 0
  AnimationNodeFilterActionFilterPass AnimationNodeFilterAction = 1
  AnimationNodeFilterActionFilterStop AnimationNodeFilterAction = 2
  AnimationNodeFilterActionFilterBlend AnimationNodeFilterAction = 3
)

func (me *AnimationNode) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationNode) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNode) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AnimationNode) AddInput(name String, ) bool {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_input")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2323990056) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNode) RemoveInput(index int, )  {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_input")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNode) SetInputName(input int, name String, ) bool {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_input_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 215573526) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input), gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNode) GetInputName(input int, ) String {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_input_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNode) GetInputCount() int {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_input_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNode) FindInput(name String, ) int {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_input")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1321353865) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNode) SetFilterPath(path NodePath, enable bool, )  {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_filter_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3868023870) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNode) IsPathFiltered(path NodePath, ) bool {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_path_filtered")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 861721659) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNode) SetFilterEnabled(enable bool, )  {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_filter_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNode) IsFilterEnabled() bool {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_filter_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNode) BlendAnimation(animation StringName, time float32, delta float32, seeked bool, is_external_seeking bool, blend float32, looped_flag AnimationLoopedFlag, )  {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("blend_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1630801826) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(animation.AsCTypePtr()), gdc.ConstTypePtr(&time), gdc.ConstTypePtr(&delta), gdc.ConstTypePtr(&seeked), gdc.ConstTypePtr(&is_external_seeking), gdc.ConstTypePtr(&blend), gdc.ConstTypePtr(&looped_flag), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNode) BlendNode(name StringName, node AnimationNode, time float32, seek bool, is_external_seeking bool, blend float32, filter AnimationNodeFilterAction, sync bool, test_only bool, ) float32 {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("blend_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1746075988) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(node.AsCTypePtr()), gdc.ConstTypePtr(&time), gdc.ConstTypePtr(&seek), gdc.ConstTypePtr(&is_external_seeking), gdc.ConstTypePtr(&blend), gdc.ConstTypePtr(&filter), gdc.ConstTypePtr(&sync), gdc.ConstTypePtr(&test_only), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNode) BlendInput(input_index int, time float32, seek bool, is_external_seeking bool, blend float32, filter AnimationNodeFilterAction, sync bool, test_only bool, ) float32 {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("blend_input")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1361527350) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input_index), gdc.ConstTypePtr(&time), gdc.ConstTypePtr(&seek), gdc.ConstTypePtr(&is_external_seeking), gdc.ConstTypePtr(&blend), gdc.ConstTypePtr(&filter), gdc.ConstTypePtr(&sync), gdc.ConstTypePtr(&test_only), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNode) SetParameter(name StringName, value Variant, )  {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_parameter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3776071444) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(value.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNode) GetParameter(name StringName, ) Variant {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_parameter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2760726917) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type AnimationNodeTreeChangedSignalFn func()

func (me *AnimationNode) ConnectTreeChanged(subs SignalSubscribers, fn AnimationNodeTreeChangedSignalFn) {
  sig := StringNameFromStr("tree_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationNode) DisconnectTreeChanged(subs SignalSubscribers, fn AnimationNodeTreeChangedSignalFn) {
  sig := StringNameFromStr("tree_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type AnimationNodeAnimationNodeRenamedSignalFn func(object_id int, old_name String, new_name String, )

func (me *AnimationNode) ConnectAnimationNodeRenamed(subs SignalSubscribers, fn AnimationNodeAnimationNodeRenamedSignalFn) {
  sig := StringNameFromStr("animation_node_renamed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationNode) DisconnectAnimationNodeRenamed(subs SignalSubscribers, fn AnimationNodeAnimationNodeRenamedSignalFn) {
  sig := StringNameFromStr("animation_node_renamed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type AnimationNodeAnimationNodeRemovedSignalFn func(object_id int, name String, )

func (me *AnimationNode) ConnectAnimationNodeRemoved(subs SignalSubscribers, fn AnimationNodeAnimationNodeRemovedSignalFn) {
  sig := StringNameFromStr("animation_node_removed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationNode) DisconnectAnimationNodeRemoved(subs SignalSubscribers, fn AnimationNodeAnimationNodeRemovedSignalFn) {
  sig := StringNameFromStr("animation_node_removed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
