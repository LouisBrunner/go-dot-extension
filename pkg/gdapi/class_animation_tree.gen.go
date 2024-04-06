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

type ptrsForAnimationTreeList struct {
	fnSetTreeRoot                  gdc.MethodBindPtr
	fnGetTreeRoot                  gdc.MethodBindPtr
	fnSetAdvanceExpressionBaseNode gdc.MethodBindPtr
	fnGetAdvanceExpressionBaseNode gdc.MethodBindPtr
	fnSetAnimationPlayer           gdc.MethodBindPtr
	fnGetAnimationPlayer           gdc.MethodBindPtr
	fnSetProcessCallback           gdc.MethodBindPtr
	fnGetProcessCallback           gdc.MethodBindPtr
}

var ptrsForAnimationTree ptrsForAnimationTreeList

func initAnimationTreePtrs(iface gdc.Interface) {

	className := StringNameFromStr("AnimationTree")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_tree_root")
		defer methodName.Destroy()
		ptrsForAnimationTree.fnSetTreeRoot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2581683800))
	}
	{
		methodName := StringNameFromStr("get_tree_root")
		defer methodName.Destroy()
		ptrsForAnimationTree.fnGetTreeRoot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4110384712))
	}
	{
		methodName := StringNameFromStr("set_advance_expression_base_node")
		defer methodName.Destroy()
		ptrsForAnimationTree.fnSetAdvanceExpressionBaseNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_advance_expression_base_node")
		defer methodName.Destroy()
		ptrsForAnimationTree.fnGetAdvanceExpressionBaseNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
	}
	{
		methodName := StringNameFromStr("set_animation_player")
		defer methodName.Destroy()
		ptrsForAnimationTree.fnSetAnimationPlayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_animation_player")
		defer methodName.Destroy()
		ptrsForAnimationTree.fnGetAnimationPlayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
	}
	{
		methodName := StringNameFromStr("set_process_callback")
		defer methodName.Destroy()
		ptrsForAnimationTree.fnSetProcessCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1723352826))
	}
	{
		methodName := StringNameFromStr("get_process_callback")
		defer methodName.Destroy()
		ptrsForAnimationTree.fnGetProcessCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 891317132))
	}
}

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
	AnimationTreeAnimationProcessCallbackAnimationProcessIdle    AnimationTreeAnimationProcessCallback = 1
	AnimationTreeAnimationProcessCallbackAnimationProcessManual  AnimationTreeAnimationProcessCallback = 2
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

func (me *AnimationTree) SetTreeRoot(animation_node AnimationRootNode) {
	cargs := []gdc.ConstTypePtr{animation_node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationTree.fnSetTreeRoot), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationTree) GetTreeRoot() AnimationRootNode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewAnimationRootNode()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationTree.fnGetTreeRoot), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationTree) SetAdvanceExpressionBaseNode(path NodePath) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationTree.fnSetAdvanceExpressionBaseNode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationTree) GetAdvanceExpressionBaseNode() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationTree.fnGetAdvanceExpressionBaseNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationTree) SetAnimationPlayer(path NodePath) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationTree.fnSetAnimationPlayer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationTree) GetAnimationPlayer() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationTree.fnGetAnimationPlayer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationTree) SetProcessCallback(mode AnimationTreeAnimationProcessCallback) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationTree.fnSetProcessCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationTree) GetProcessCallback() AnimationTreeAnimationProcessCallback {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AnimationTreeAnimationProcessCallback

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationTree.fnGetProcessCallback), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
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
