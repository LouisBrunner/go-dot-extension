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

type AnimationNode struct {
  Resource
}

func (me *AnimationNode) BaseClass() string {
  return "AnimationNode"
}

func NewAnimationNode() *AnimationNode {
  str := StringNameFromStr("AnimationNode") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AnimationNode{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNode) RemoveInput(index int64, )  {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_input")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNode) SetInputName(input int64, name String, ) bool {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_input_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 215573526) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input) , name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&input)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNode) GetInputName(input int64, ) String {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_input_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&input)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNode) GetInputCount() int64 {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_input_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNode) FindInput(name String, ) int64 {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_input")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1321353865) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNode) SetFilterPath(path NodePath, enable bool, )  {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_filter_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3868023870) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNode) IsPathFiltered(path NodePath, ) bool {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_path_filtered")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 861721659) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNode) SetFilterEnabled(enable bool, )  {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_filter_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNode) IsFilterEnabled() bool {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_filter_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNode) BlendAnimation(animation StringName, time float64, delta float64, seeked bool, is_external_seeking bool, blend float64, looped_flag AnimationLoopedFlag, )  {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("blend_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1630801826) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{animation.AsCTypePtr(), gdc.ConstTypePtr(&time) , gdc.ConstTypePtr(&delta) , gdc.ConstTypePtr(&seeked) , gdc.ConstTypePtr(&is_external_seeking) , gdc.ConstTypePtr(&blend) , gdc.ConstTypePtr(&looped_flag) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNode) BlendNode(name StringName, node AnimationNode, time float64, seek bool, is_external_seeking bool, blend float64, filter AnimationNodeFilterAction, sync bool, test_only bool, ) float64 {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("blend_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1746075988) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), node.AsCTypePtr(), gdc.ConstTypePtr(&time) , gdc.ConstTypePtr(&seek) , gdc.ConstTypePtr(&is_external_seeking) , gdc.ConstTypePtr(&blend) , gdc.ConstTypePtr(&filter) , gdc.ConstTypePtr(&sync) , gdc.ConstTypePtr(&test_only) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&time)
  pinner.Pin(&seek)
  pinner.Pin(&is_external_seeking)
  pinner.Pin(&blend)
  pinner.Pin(&filter)
  pinner.Pin(&sync)
  pinner.Pin(&test_only)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNode) BlendInput(input_index int64, time float64, seek bool, is_external_seeking bool, blend float64, filter AnimationNodeFilterAction, sync bool, test_only bool, ) float64 {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("blend_input")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1361527350) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input_index) , gdc.ConstTypePtr(&time) , gdc.ConstTypePtr(&seek) , gdc.ConstTypePtr(&is_external_seeking) , gdc.ConstTypePtr(&blend) , gdc.ConstTypePtr(&filter) , gdc.ConstTypePtr(&sync) , gdc.ConstTypePtr(&test_only) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&input_index)
  pinner.Pin(&time)
  pinner.Pin(&seek)
  pinner.Pin(&is_external_seeking)
  pinner.Pin(&blend)
  pinner.Pin(&filter)
  pinner.Pin(&sync)
  pinner.Pin(&test_only)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNode) SetParameter(name StringName, value Variant, )  {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_parameter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3776071444) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNode) GetParameter(name StringName, ) Variant {
  classNameV := StringNameFromStr("AnimationNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_parameter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2760726917) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
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
