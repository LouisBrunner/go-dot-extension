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

type ptrsForAnimationNodeList struct {
	fnXGetChildNodes            gdc.MethodBindPtr
	fnXGetParameterList         gdc.MethodBindPtr
	fnXGetChildByName           gdc.MethodBindPtr
	fnXGetParameterDefaultValue gdc.MethodBindPtr
	fnXIsParameterReadOnly      gdc.MethodBindPtr
	fnXProcess                  gdc.MethodBindPtr
	fnXGetCaption               gdc.MethodBindPtr
	fnXHasFilter                gdc.MethodBindPtr
	fnAddInput                  gdc.MethodBindPtr
	fnRemoveInput               gdc.MethodBindPtr
	fnSetInputName              gdc.MethodBindPtr
	fnGetInputName              gdc.MethodBindPtr
	fnGetInputCount             gdc.MethodBindPtr
	fnFindInput                 gdc.MethodBindPtr
	fnSetFilterPath             gdc.MethodBindPtr
	fnIsPathFiltered            gdc.MethodBindPtr
	fnSetFilterEnabled          gdc.MethodBindPtr
	fnIsFilterEnabled           gdc.MethodBindPtr
	fnBlendAnimation            gdc.MethodBindPtr
	fnBlendNode                 gdc.MethodBindPtr
	fnBlendInput                gdc.MethodBindPtr
	fnSetParameter              gdc.MethodBindPtr
	fnGetParameter              gdc.MethodBindPtr
}

var ptrsForAnimationNode ptrsForAnimationNodeList

func initAnimationNodePtrs(iface gdc.Interface) {

	className := StringNameFromStr("AnimationNode")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("add_input")
		defer methodName.Destroy()
		ptrsForAnimationNode.fnAddInput = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2323990056))
	}
	{
		methodName := StringNameFromStr("remove_input")
		defer methodName.Destroy()
		ptrsForAnimationNode.fnRemoveInput = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_input_name")
		defer methodName.Destroy()
		ptrsForAnimationNode.fnSetInputName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 215573526))
	}
	{
		methodName := StringNameFromStr("get_input_name")
		defer methodName.Destroy()
		ptrsForAnimationNode.fnGetInputName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("get_input_count")
		defer methodName.Destroy()
		ptrsForAnimationNode.fnGetInputCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("find_input")
		defer methodName.Destroy()
		ptrsForAnimationNode.fnFindInput = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1321353865))
	}
	{
		methodName := StringNameFromStr("set_filter_path")
		defer methodName.Destroy()
		ptrsForAnimationNode.fnSetFilterPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3868023870))
	}
	{
		methodName := StringNameFromStr("is_path_filtered")
		defer methodName.Destroy()
		ptrsForAnimationNode.fnIsPathFiltered = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 861721659))
	}
	{
		methodName := StringNameFromStr("set_filter_enabled")
		defer methodName.Destroy()
		ptrsForAnimationNode.fnSetFilterEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_filter_enabled")
		defer methodName.Destroy()
		ptrsForAnimationNode.fnIsFilterEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("blend_animation")
		defer methodName.Destroy()
		ptrsForAnimationNode.fnBlendAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1630801826))
	}
	{
		methodName := StringNameFromStr("blend_node")
		defer methodName.Destroy()
		ptrsForAnimationNode.fnBlendNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1746075988))
	}
	{
		methodName := StringNameFromStr("blend_input")
		defer methodName.Destroy()
		ptrsForAnimationNode.fnBlendInput = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1361527350))
	}
	{
		methodName := StringNameFromStr("set_parameter")
		defer methodName.Destroy()
		ptrsForAnimationNode.fnSetParameter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3776071444))
	}
	{
		methodName := StringNameFromStr("get_parameter")
		defer methodName.Destroy()
		ptrsForAnimationNode.fnGetParameter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2760726917))
	}
}

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
	AnimationNodeFilterActionFilterPass   AnimationNodeFilterAction = 1
	AnimationNodeFilterActionFilterStop   AnimationNodeFilterAction = 2
	AnimationNodeFilterActionFilterBlend  AnimationNodeFilterAction = 3
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

func (me *AnimationNode) AddInput(name String) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNode.fnAddInput), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNode) RemoveInput(index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNode.fnRemoveInput), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNode) SetInputName(input int64, name String) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input), name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&input)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNode.fnSetInputName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNode) GetInputName(input int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&input)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNode.fnGetInputName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationNode) GetInputCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNode.fnGetInputCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNode) FindInput(name String) int64 {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNode.fnFindInput), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNode) SetFilterPath(path NodePath, enable bool) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNode.fnSetFilterPath), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNode) IsPathFiltered(path NodePath) bool {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNode.fnIsPathFiltered), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNode) SetFilterEnabled(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNode.fnSetFilterEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNode) IsFilterEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNode.fnIsFilterEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNode) BlendAnimation(animation StringName, time float64, delta float64, seeked bool, is_external_seeking bool, blend float64, looped_flag AnimationLoopedFlag) {
	cargs := []gdc.ConstTypePtr{animation.AsCTypePtr(), gdc.ConstTypePtr(&time), gdc.ConstTypePtr(&delta), gdc.ConstTypePtr(&seeked), gdc.ConstTypePtr(&is_external_seeking), gdc.ConstTypePtr(&blend), gdc.ConstTypePtr(&looped_flag)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNode.fnBlendAnimation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNode) BlendNode(name StringName, node AnimationNode, time float64, seek bool, is_external_seeking bool, blend float64, filter AnimationNodeFilterAction, sync bool, test_only bool) float64 {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), node.AsCTypePtr(), gdc.ConstTypePtr(&time), gdc.ConstTypePtr(&seek), gdc.ConstTypePtr(&is_external_seeking), gdc.ConstTypePtr(&blend), gdc.ConstTypePtr(&filter), gdc.ConstTypePtr(&sync), gdc.ConstTypePtr(&test_only)}
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

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNode.fnBlendNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNode) BlendInput(input_index int64, time float64, seek bool, is_external_seeking bool, blend float64, filter AnimationNodeFilterAction, sync bool, test_only bool) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input_index), gdc.ConstTypePtr(&time), gdc.ConstTypePtr(&seek), gdc.ConstTypePtr(&is_external_seeking), gdc.ConstTypePtr(&blend), gdc.ConstTypePtr(&filter), gdc.ConstTypePtr(&sync), gdc.ConstTypePtr(&test_only)}
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

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNode.fnBlendInput), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNode) SetParameter(name StringName, value Variant) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNode.fnSetParameter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNode) GetParameter(name StringName) Variant {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNode.fnGetParameter), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
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

type AnimationNodeAnimationNodeRenamedSignalFn func(object_id int, old_name String, new_name String)

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

type AnimationNodeAnimationNodeRemovedSignalFn func(object_id int, name String)

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
