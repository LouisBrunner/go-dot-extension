// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SceneState struct {
  obj gdc.ObjectPtr
}

func (me *SceneState) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SceneState) BaseClass() string {
  return "SceneState"
}



// Enums

type SceneStateGenEditState int
const (
  SceneStateGenEditStateGenEditStateDisabled SceneStateGenEditState = 0
  SceneStateGenEditStateGenEditStateInstance SceneStateGenEditState = 1
  SceneStateGenEditStateGenEditStateMain SceneStateGenEditState = 2
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

func  (me *SceneState) GetNodeCount() int {
  classNameV := StringNameFromStr("SceneState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneState) GetNodeType(idx int, ) StringName {
  classNameV := StringNameFromStr("SceneState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 659327637) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneState) GetNodeName(idx int, ) StringName {
  classNameV := StringNameFromStr("SceneState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 659327637) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneState) GetNodePath(idx int, for_parent bool, ) NodePath {
  classNameV := StringNameFromStr("SceneState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2272487792) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&for_parent), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneState) GetNodeOwnerPath(idx int, ) NodePath {
  classNameV := StringNameFromStr("SceneState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_owner_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 408788394) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneState) IsNodeInstancePlaceholder(idx int, ) bool {
  classNameV := StringNameFromStr("SceneState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_node_instance_placeholder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneState) GetNodeInstancePlaceholder(idx int, ) String {
  classNameV := StringNameFromStr("SceneState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_instance_placeholder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneState) GetNodeInstance(idx int, ) PackedScene {
  classNameV := StringNameFromStr("SceneState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_instance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 511017218) // FIXME: should cache?
  var ret PackedScene
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneState) GetNodeGroups(idx int, ) PackedStringArray {
  classNameV := StringNameFromStr("SceneState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_groups")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 647634434) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneState) GetNodeIndex(idx int, ) int {
  classNameV := StringNameFromStr("SceneState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneState) GetNodePropertyCount(idx int, ) int {
  classNameV := StringNameFromStr("SceneState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_property_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneState) GetNodePropertyName(idx int, prop_idx int, ) StringName {
  classNameV := StringNameFromStr("SceneState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_property_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 351665558) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&prop_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneState) GetNodePropertyValue(idx int, prop_idx int, ) Variant {
  classNameV := StringNameFromStr("SceneState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_property_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 678354945) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&prop_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneState) GetConnectionCount() int {
  classNameV := StringNameFromStr("SceneState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneState) GetConnectionSource(idx int, ) NodePath {
  classNameV := StringNameFromStr("SceneState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_source")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 408788394) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneState) GetConnectionSignal(idx int, ) StringName {
  classNameV := StringNameFromStr("SceneState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_signal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 659327637) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneState) GetConnectionTarget(idx int, ) NodePath {
  classNameV := StringNameFromStr("SceneState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_target")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 408788394) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneState) GetConnectionMethod(idx int, ) StringName {
  classNameV := StringNameFromStr("SceneState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_method")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 659327637) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneState) GetConnectionFlags(idx int, ) int {
  classNameV := StringNameFromStr("SceneState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneState) GetConnectionBinds(idx int, ) Array {
  classNameV := StringNameFromStr("SceneState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_binds")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 663333327) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneState) GetConnectionUnbinds(idx int, ) int {
  classNameV := StringNameFromStr("SceneState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_unbinds")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
