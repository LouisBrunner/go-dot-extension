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

type ptrsForSceneStateList struct {
	fnGetNodeCount               gdc.MethodBindPtr
	fnGetNodeType                gdc.MethodBindPtr
	fnGetNodeName                gdc.MethodBindPtr
	fnGetNodePath                gdc.MethodBindPtr
	fnGetNodeOwnerPath           gdc.MethodBindPtr
	fnIsNodeInstancePlaceholder  gdc.MethodBindPtr
	fnGetNodeInstancePlaceholder gdc.MethodBindPtr
	fnGetNodeInstance            gdc.MethodBindPtr
	fnGetNodeGroups              gdc.MethodBindPtr
	fnGetNodeIndex               gdc.MethodBindPtr
	fnGetNodePropertyCount       gdc.MethodBindPtr
	fnGetNodePropertyName        gdc.MethodBindPtr
	fnGetNodePropertyValue       gdc.MethodBindPtr
	fnGetConnectionCount         gdc.MethodBindPtr
	fnGetConnectionSource        gdc.MethodBindPtr
	fnGetConnectionSignal        gdc.MethodBindPtr
	fnGetConnectionTarget        gdc.MethodBindPtr
	fnGetConnectionMethod        gdc.MethodBindPtr
	fnGetConnectionFlags         gdc.MethodBindPtr
	fnGetConnectionBinds         gdc.MethodBindPtr
	fnGetConnectionUnbinds       gdc.MethodBindPtr
}

var ptrsForSceneState ptrsForSceneStateList

func initSceneStatePtrs(iface gdc.Interface) {

	className := StringNameFromStr("SceneState")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_node_count")
		defer methodName.Destroy()
		ptrsForSceneState.fnGetNodeCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_node_type")
		defer methodName.Destroy()
		ptrsForSceneState.fnGetNodeType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 659327637))
	}
	{
		methodName := StringNameFromStr("get_node_name")
		defer methodName.Destroy()
		ptrsForSceneState.fnGetNodeName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 659327637))
	}
	{
		methodName := StringNameFromStr("get_node_path")
		defer methodName.Destroy()
		ptrsForSceneState.fnGetNodePath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2272487792))
	}
	{
		methodName := StringNameFromStr("get_node_owner_path")
		defer methodName.Destroy()
		ptrsForSceneState.fnGetNodeOwnerPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 408788394))
	}
	{
		methodName := StringNameFromStr("is_node_instance_placeholder")
		defer methodName.Destroy()
		ptrsForSceneState.fnIsNodeInstancePlaceholder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("get_node_instance_placeholder")
		defer methodName.Destroy()
		ptrsForSceneState.fnGetNodeInstancePlaceholder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("get_node_instance")
		defer methodName.Destroy()
		ptrsForSceneState.fnGetNodeInstance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 511017218))
	}
	{
		methodName := StringNameFromStr("get_node_groups")
		defer methodName.Destroy()
		ptrsForSceneState.fnGetNodeGroups = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 647634434))
	}
	{
		methodName := StringNameFromStr("get_node_index")
		defer methodName.Destroy()
		ptrsForSceneState.fnGetNodeIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("get_node_property_count")
		defer methodName.Destroy()
		ptrsForSceneState.fnGetNodePropertyCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("get_node_property_name")
		defer methodName.Destroy()
		ptrsForSceneState.fnGetNodePropertyName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 351665558))
	}
	{
		methodName := StringNameFromStr("get_node_property_value")
		defer methodName.Destroy()
		ptrsForSceneState.fnGetNodePropertyValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 678354945))
	}
	{
		methodName := StringNameFromStr("get_connection_count")
		defer methodName.Destroy()
		ptrsForSceneState.fnGetConnectionCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_connection_source")
		defer methodName.Destroy()
		ptrsForSceneState.fnGetConnectionSource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 408788394))
	}
	{
		methodName := StringNameFromStr("get_connection_signal")
		defer methodName.Destroy()
		ptrsForSceneState.fnGetConnectionSignal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 659327637))
	}
	{
		methodName := StringNameFromStr("get_connection_target")
		defer methodName.Destroy()
		ptrsForSceneState.fnGetConnectionTarget = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 408788394))
	}
	{
		methodName := StringNameFromStr("get_connection_method")
		defer methodName.Destroy()
		ptrsForSceneState.fnGetConnectionMethod = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 659327637))
	}
	{
		methodName := StringNameFromStr("get_connection_flags")
		defer methodName.Destroy()
		ptrsForSceneState.fnGetConnectionFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("get_connection_binds")
		defer methodName.Destroy()
		ptrsForSceneState.fnGetConnectionBinds = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 663333327))
	}
	{
		methodName := StringNameFromStr("get_connection_unbinds")
		defer methodName.Destroy()
		ptrsForSceneState.fnGetConnectionUnbinds = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}

}

type SceneState struct {
	RefCounted
}

func (me *SceneState) BaseClass() string {
	return "SceneState"
}

func NewSceneState() *SceneState {
	str := StringNameFromStr("SceneState") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &SceneState{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type SceneStateGenEditState int

const (
	SceneStateGenEditStateGenEditStateDisabled      SceneStateGenEditState = 0
	SceneStateGenEditStateGenEditStateInstance      SceneStateGenEditState = 1
	SceneStateGenEditStateGenEditStateMain          SceneStateGenEditState = 2
	SceneStateGenEditStateGenEditStateMainInherited SceneStateGenEditState = 3
)

func (me *SceneState) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *SceneState) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *SceneState) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *SceneState) GetNodeCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneState.fnGetNodeCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SceneState) GetNodeType(idx int64) StringName {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneState.fnGetNodeType), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SceneState) GetNodeName(idx int64) StringName {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneState.fnGetNodeName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SceneState) GetNodePath(idx int64, for_parent bool) NodePath {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&for_parent)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()
	pinner.Pin(&idx)
	pinner.Pin(&for_parent)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneState.fnGetNodePath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SceneState) GetNodeOwnerPath(idx int64) NodePath {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneState.fnGetNodeOwnerPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SceneState) IsNodeInstancePlaceholder(idx int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneState.fnIsNodeInstancePlaceholder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SceneState) GetNodeInstancePlaceholder(idx int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneState.fnGetNodeInstancePlaceholder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SceneState) GetNodeInstance(idx int64) PackedScene {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedScene()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneState.fnGetNodeInstance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SceneState) GetNodeGroups(idx int64) PackedStringArray {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneState.fnGetNodeGroups), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SceneState) GetNodeIndex(idx int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneState.fnGetNodeIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SceneState) GetNodePropertyCount(idx int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneState.fnGetNodePropertyCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SceneState) GetNodePropertyName(idx int64, prop_idx int64) StringName {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&prop_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()
	pinner.Pin(&idx)
	pinner.Pin(&prop_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneState.fnGetNodePropertyName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SceneState) GetNodePropertyValue(idx int64, prop_idx int64) Variant {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&prop_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()
	pinner.Pin(&idx)
	pinner.Pin(&prop_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneState.fnGetNodePropertyValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SceneState) GetConnectionCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneState.fnGetConnectionCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SceneState) GetConnectionSource(idx int64) NodePath {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneState.fnGetConnectionSource), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SceneState) GetConnectionSignal(idx int64) StringName {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneState.fnGetConnectionSignal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SceneState) GetConnectionTarget(idx int64) NodePath {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneState.fnGetConnectionTarget), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SceneState) GetConnectionMethod(idx int64) StringName {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneState.fnGetConnectionMethod), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SceneState) GetConnectionFlags(idx int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneState.fnGetConnectionFlags), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SceneState) GetConnectionBinds(idx int64) Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneState.fnGetConnectionBinds), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SceneState) GetConnectionUnbinds(idx int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneState.fnGetConnectionUnbinds), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Signals
