// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Node struct {
  obj gdc.ObjectPtr
}

func (me *Node) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Node) BaseClass() string {
  return "Node"
}



// Constants

var (
  NodeNotificationEnterTree = "10" // TODO: construct correctly
  NodeNotificationExitTree = "11" // TODO: construct correctly
  NodeNotificationMovedInParent = "12" // TODO: construct correctly
  NodeNotificationReady = "13" // TODO: construct correctly
  NodeNotificationPaused = "14" // TODO: construct correctly
  NodeNotificationUnpaused = "15" // TODO: construct correctly
  NodeNotificationPhysicsProcess = "16" // TODO: construct correctly
  NodeNotificationProcess = "17" // TODO: construct correctly
  NodeNotificationParented = "18" // TODO: construct correctly
  NodeNotificationUnparented = "19" // TODO: construct correctly
  NodeNotificationSceneInstantiated = "20" // TODO: construct correctly
  NodeNotificationDragBegin = "21" // TODO: construct correctly
  NodeNotificationDragEnd = "22" // TODO: construct correctly
  NodeNotificationPathRenamed = "23" // TODO: construct correctly
  NodeNotificationChildOrderChanged = "24" // TODO: construct correctly
  NodeNotificationInternalProcess = "25" // TODO: construct correctly
  NodeNotificationInternalPhysicsProcess = "26" // TODO: construct correctly
  NodeNotificationPostEnterTree = "27" // TODO: construct correctly
  NodeNotificationDisabled = "28" // TODO: construct correctly
  NodeNotificationEnabled = "29" // TODO: construct correctly
  NodeNotificationNodeRecacheRequested = "30" // TODO: construct correctly
  NodeNotificationEditorPreSave = "9001" // TODO: construct correctly
  NodeNotificationEditorPostSave = "9002" // TODO: construct correctly
  NodeNotificationWmMouseEnter = "1002" // TODO: construct correctly
  NodeNotificationWmMouseExit = "1003" // TODO: construct correctly
  NodeNotificationWmWindowFocusIn = "1004" // TODO: construct correctly
  NodeNotificationWmWindowFocusOut = "1005" // TODO: construct correctly
  NodeNotificationWmCloseRequest = "1006" // TODO: construct correctly
  NodeNotificationWmGoBackRequest = "1007" // TODO: construct correctly
  NodeNotificationWmSizeChanged = "1008" // TODO: construct correctly
  NodeNotificationWmDpiChange = "1009" // TODO: construct correctly
  NodeNotificationVpMouseEnter = "1010" // TODO: construct correctly
  NodeNotificationVpMouseExit = "1011" // TODO: construct correctly
  NodeNotificationOsMemoryWarning = "2009" // TODO: construct correctly
  NodeNotificationTranslationChanged = "2010" // TODO: construct correctly
  NodeNotificationWmAbout = "2011" // TODO: construct correctly
  NodeNotificationCrash = "2012" // TODO: construct correctly
  NodeNotificationOsImeUpdate = "2013" // TODO: construct correctly
  NodeNotificationApplicationResumed = "2014" // TODO: construct correctly
  NodeNotificationApplicationPaused = "2015" // TODO: construct correctly
  NodeNotificationApplicationFocusIn = "2016" // TODO: construct correctly
  NodeNotificationApplicationFocusOut = "2017" // TODO: construct correctly
  NodeNotificationTextServerChanged = "2018" // TODO: construct correctly
)

// Enums

type NodeProcessMode int
const (
  NodeProcessModeProcessModeInherit NodeProcessMode = 0
  NodeProcessModeProcessModePausable NodeProcessMode = 1
  NodeProcessModeProcessModeWhenPaused NodeProcessMode = 2
  NodeProcessModeProcessModeAlways NodeProcessMode = 3
  NodeProcessModeProcessModeDisabled NodeProcessMode = 4
)

type NodeProcessThreadGroup int
const (
  NodeProcessThreadGroupProcessThreadGroupInherit NodeProcessThreadGroup = 0
  NodeProcessThreadGroupProcessThreadGroupMainThread NodeProcessThreadGroup = 1
  NodeProcessThreadGroupProcessThreadGroupSubThread NodeProcessThreadGroup = 2
)

type NodeProcessThreadMessages int
const (
  NodeProcessThreadMessagesFlagProcessThreadMessages NodeProcessThreadMessages = 1
  NodeProcessThreadMessagesFlagProcessThreadMessagesPhysics NodeProcessThreadMessages = 2
  NodeProcessThreadMessagesFlagProcessThreadMessagesAll NodeProcessThreadMessages = 3
)

type NodeDuplicateFlags int
const (
  NodeDuplicateFlagsDuplicateSignals NodeDuplicateFlags = 1
  NodeDuplicateFlagsDuplicateGroups NodeDuplicateFlags = 2
  NodeDuplicateFlagsDuplicateScripts NodeDuplicateFlags = 4
  NodeDuplicateFlagsDuplicateUseInstantiation NodeDuplicateFlags = 8
)

type NodeInternalMode int
const (
  NodeInternalModeInternalModeDisabled NodeInternalMode = 0
  NodeInternalModeInternalModeFront NodeInternalMode = 1
  NodeInternalModeInternalModeBack NodeInternalMode = 2
)

func (me *Node) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Node) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Node) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  NodePrintOrphanNodes()  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("print_orphan_nodes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), nil)
}

func  (me *Node) AddSibling(sibling Node, force_readable_name bool, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_sibling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2570952461) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(sibling.AsCTypePtr()), gdc.ConstTypePtr(&force_readable_name), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) SetName(name String, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) GetName() StringName {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) AddChild(node Node, force_readable_name bool, internal NodeInternalMode, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_child")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3070154285) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), gdc.ConstTypePtr(&force_readable_name), gdc.ConstTypePtr(&internal), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) RemoveChild(node Node, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_child")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) Reparent(new_parent Node, keep_global_transform bool, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("reparent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2570952461) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(new_parent.AsCTypePtr()), gdc.ConstTypePtr(&keep_global_transform), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) GetChildCount(include_internal bool, ) int {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_child_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 894402480) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&include_internal), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) GetChildren(include_internal bool, ) Node {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_children")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 873284517) // FIXME: should cache?
  var ret Node
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&include_internal), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) GetChild(idx int, include_internal bool, ) Node {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_child")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 541253412) // FIXME: should cache?
  var ret Node
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&include_internal), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) HasNode(path NodePath, ) bool {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 861721659) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) GetNode(path NodePath, ) Node {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2734337346) // FIXME: should cache?
  var ret Node
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) GetNodeOrNull(path NodePath, ) Node {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_or_null")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2734337346) // FIXME: should cache?
  var ret Node
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) GetParent() Node {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_parent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3160264692) // FIXME: should cache?
  var ret Node
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) FindChild(pattern String, recursive bool, owned bool, ) Node {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_child")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4253159453) // FIXME: should cache?
  var ret Node
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(pattern.AsCTypePtr()), gdc.ConstTypePtr(&recursive), gdc.ConstTypePtr(&owned), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) FindChildren(pattern String, type_ String, recursive bool, owned bool, ) Node {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_children")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1585018254) // FIXME: should cache?
  var ret Node
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(pattern.AsCTypePtr()), gdc.ConstTypePtr(type_.AsCTypePtr()), gdc.ConstTypePtr(&recursive), gdc.ConstTypePtr(&owned), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) FindParent(pattern String, ) Node {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_parent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1140089439) // FIXME: should cache?
  var ret Node
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(pattern.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) HasNodeAndResource(path NodePath, ) bool {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_node_and_resource")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 861721659) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) GetNodeAndResource(path NodePath, ) Array {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_and_resource")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 502563882) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) IsInsideTree() bool {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_inside_tree")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) IsAncestorOf(node Node, ) bool {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_ancestor_of")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3093956946) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) IsGreaterThan(node Node, ) bool {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_greater_than")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3093956946) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) GetPath() NodePath {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) GetPathTo(node Node, use_unique_path bool, ) NodePath {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path_to")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 498846349) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), gdc.ConstTypePtr(&use_unique_path), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) AddToGroup(group StringName, persistent bool, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_to_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3683006648) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(group.AsCTypePtr()), gdc.ConstTypePtr(&persistent), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) RemoveFromGroup(group StringName, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_from_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(group.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) IsInGroup(group StringName, ) bool {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_in_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(group.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) MoveChild(child_node Node, to_index int, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_child")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3315886247) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(child_node.AsCTypePtr()), gdc.ConstTypePtr(&to_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) GetGroups() StringName {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_groups")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) SetOwner(owner Node, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_owner")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(owner.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) GetOwner() Node {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_owner")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3160264692) // FIXME: should cache?
  var ret Node
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) GetIndex(include_internal bool, ) int {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 894402480) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&include_internal), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) PrintTree()  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("print_tree")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) PrintTreePretty()  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("print_tree_pretty")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) SetSceneFilePath(scene_file_path String, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scene_file_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(scene_file_path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) GetSceneFilePath() String {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scene_file_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) PropagateNotification(what int, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("propagate_notification")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&what), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) PropagateCall(method StringName, args Array, parent_first bool, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("propagate_call")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1667910434) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(method.AsCTypePtr()), gdc.ConstTypePtr(args.AsCTypePtr()), gdc.ConstTypePtr(&parent_first), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) SetPhysicsProcess(enable bool, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_physics_process")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) GetPhysicsProcessDeltaTime() float32 {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_physics_process_delta_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) IsPhysicsProcessing() bool {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_physics_processing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) GetProcessDeltaTime() float32 {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_process_delta_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) SetProcess(enable bool, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_process")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) SetProcessPriority(priority int, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_process_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) GetProcessPriority() int {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_process_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) SetPhysicsProcessPriority(priority int, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_physics_process_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) GetPhysicsProcessPriority() int {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_physics_process_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) IsProcessing() bool {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_processing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) SetProcessInput(enable bool, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_process_input")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) IsProcessingInput() bool {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_processing_input")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) SetProcessShortcutInput(enable bool, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_process_shortcut_input")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) IsProcessingShortcutInput() bool {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_processing_shortcut_input")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) SetProcessUnhandledInput(enable bool, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_process_unhandled_input")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) IsProcessingUnhandledInput() bool {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_processing_unhandled_input")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) SetProcessUnhandledKeyInput(enable bool, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_process_unhandled_key_input")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) IsProcessingUnhandledKeyInput() bool {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_processing_unhandled_key_input")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) SetProcessMode(mode NodeProcessMode, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_process_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1841290486) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) GetProcessMode() NodeProcessMode {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_process_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 739966102) // FIXME: should cache?
  var ret NodeProcessMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) CanProcess() bool {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("can_process")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) SetProcessThreadGroup(mode NodeProcessThreadGroup, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_process_thread_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2275442745) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) GetProcessThreadGroup() NodeProcessThreadGroup {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_process_thread_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1866404740) // FIXME: should cache?
  var ret NodeProcessThreadGroup
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) SetProcessThreadMessages(flags NodeProcessThreadMessages, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_process_thread_messages")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1357280998) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) GetProcessThreadMessages() NodeProcessThreadMessages {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_process_thread_messages")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4228993612) // FIXME: should cache?
  var ret NodeProcessThreadMessages
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) SetProcessThreadGroupOrder(order int, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_process_thread_group_order")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&order), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) GetProcessThreadGroupOrder() int {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_process_thread_group_order")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) SetDisplayFolded(fold bool, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_display_folded")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fold), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) IsDisplayedFolded() bool {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_displayed_folded")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) SetProcessInternal(enable bool, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_process_internal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) IsProcessingInternal() bool {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_processing_internal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) SetPhysicsProcessInternal(enable bool, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_physics_process_internal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) IsPhysicsProcessingInternal() bool {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_physics_processing_internal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) GetWindow() Window {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_window")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1757182445) // FIXME: should cache?
  var ret Window
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) GetLastExclusiveWindow() Window {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_last_exclusive_window")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1757182445) // FIXME: should cache?
  var ret Window
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) GetTree() SceneTree {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tree")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2958820483) // FIXME: should cache?
  var ret SceneTree
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) CreateTween() Tween {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_tween")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3426978995) // FIXME: should cache?
  var ret Tween
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) Duplicate(flags int, ) Node {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("duplicate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3511555459) // FIXME: should cache?
  var ret Node
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) ReplaceBy(node Node, keep_groups bool, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("replace_by")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2570952461) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), gdc.ConstTypePtr(&keep_groups), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) SetSceneInstanceLoadPlaceholder(load_placeholder bool, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scene_instance_load_placeholder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&load_placeholder), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) GetSceneInstanceLoadPlaceholder() bool {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scene_instance_load_placeholder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) SetEditableInstance(node Node, is_editable bool, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_editable_instance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2731852923) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), gdc.ConstTypePtr(&is_editable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) IsEditableInstance(node Node, ) bool {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_editable_instance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3093956946) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) GetViewport() Viewport {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_viewport")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3596683776) // FIXME: should cache?
  var ret Viewport
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) QueueFree()  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("queue_free")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) RequestReady()  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("request_ready")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) IsNodeReady() bool {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_node_ready")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) SetMultiplayerAuthority(id int, recursive bool, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_multiplayer_authority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4023243586) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&recursive), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) GetMultiplayerAuthority() int {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_multiplayer_authority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) IsMultiplayerAuthority() bool {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_multiplayer_authority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) GetMultiplayer() MultiplayerAPI {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_multiplayer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 406750475) // FIXME: should cache?
  var ret MultiplayerAPI
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) RpcConfig(method StringName, config Variant, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rpc_config")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3776071444) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(method.AsCTypePtr()), gdc.ConstTypePtr(config.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) SetEditorDescription(editor_description String, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_editor_description")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(editor_description.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) GetEditorDescription() String {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_editor_description")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) SetUniqueNameInOwner(enable bool, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_unique_name_in_owner")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) IsUniqueNameInOwner() bool {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_unique_name_in_owner")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Node) Rpc(method StringName, ) Error {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rpc")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4047867050) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstVariantPtr{gdc.ConstVariantPtr(method.AsCTypePtr()), }
  err := &gdc.CallError{}
  giface.ObjectMethodBindCall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), gdc.UninitializedVariantPtr(&ret), err)
  if err.Error != gdc.CallOk {
    panic(err) // TODO: return `err`?
  }

  return ret
}

func  (me *Node) RpcId(peer_id int, method StringName, ) Error {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rpc_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 361499283) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstVariantPtr{gdc.ConstVariantPtr(&peer_id), gdc.ConstVariantPtr(method.AsCTypePtr()), }
  err := &gdc.CallError{}
  giface.ObjectMethodBindCall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), gdc.UninitializedVariantPtr(&ret), err)
  if err.Error != gdc.CallOk {
    panic(err) // TODO: return `err`?
  }

  return ret
}

func  (me *Node) UpdateConfigurationWarnings()  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("update_configuration_warnings")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) CallDeferredThreadGroup(method StringName, ) Variant {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("call_deferred_thread_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3400424181) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstVariantPtr{gdc.ConstVariantPtr(method.AsCTypePtr()), }
  err := &gdc.CallError{}
  giface.ObjectMethodBindCall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), gdc.UninitializedVariantPtr(&ret), err)
  if err.Error != gdc.CallOk {
    panic(err) // TODO: return `err`?
  }

  return ret
}

func  (me *Node) SetDeferredThreadGroup(property StringName, value Variant, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_deferred_thread_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3776071444) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(property.AsCTypePtr()), gdc.ConstTypePtr(value.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) NotifyDeferredThreadGroup(what int, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("notify_deferred_thread_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&what), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) CallThreadSafe(method StringName, ) Variant {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("call_thread_safe")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3400424181) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstVariantPtr{gdc.ConstVariantPtr(method.AsCTypePtr()), }
  err := &gdc.CallError{}
  giface.ObjectMethodBindCall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), gdc.UninitializedVariantPtr(&ret), err)
  if err.Error != gdc.CallOk {
    panic(err) // TODO: return `err`?
  }

  return ret
}

func  (me *Node) SetThreadSafe(property StringName, value Variant, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_thread_safe")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3776071444) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(property.AsCTypePtr()), gdc.ConstTypePtr(value.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Node) NotifyThreadSafe(what int, )  {
  classNameV := StringNameFromStr("Node")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("notify_thread_safe")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&what), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties

func (me *Node) GetPropName() StringName {
  panic("TODO: implement")
}

func (me *Node) SetPropName(value StringName) {
  panic("TODO: implement")
}

func (me *Node) GetPropUniqueNameInOwner() bool {
  panic("TODO: implement")
}

func (me *Node) SetPropUniqueNameInOwner(value bool) {
  panic("TODO: implement")
}

func (me *Node) GetPropSceneFilePath() String {
  panic("TODO: implement")
}

func (me *Node) SetPropSceneFilePath(value String) {
  panic("TODO: implement")
}

func (me *Node) GetPropOwner() Node {
  panic("TODO: implement")
}

func (me *Node) SetPropOwner(value Node) {
  panic("TODO: implement")
}

func (me *Node) GetPropMultiplayer() MultiplayerAPI {
  panic("TODO: implement")
}

func (me *Node) SetPropMultiplayer(value MultiplayerAPI) {
  panic("TODO: implement")
}

func (me *Node) GetPropProcessMode() int {
  panic("TODO: implement")
}

func (me *Node) SetPropProcessMode(value int) {
  panic("TODO: implement")
}

func (me *Node) GetPropProcessPriority() int {
  panic("TODO: implement")
}

func (me *Node) SetPropProcessPriority(value int) {
  panic("TODO: implement")
}

func (me *Node) GetPropProcessPhysicsPriority() int {
  panic("TODO: implement")
}

func (me *Node) SetPropProcessPhysicsPriority(value int) {
  panic("TODO: implement")
}

func (me *Node) GetPropProcessThreadGroup() int {
  panic("TODO: implement")
}

func (me *Node) SetPropProcessThreadGroup(value int) {
  panic("TODO: implement")
}

func (me *Node) GetPropProcessThreadGroupOrder() int {
  panic("TODO: implement")
}

func (me *Node) SetPropProcessThreadGroupOrder(value int) {
  panic("TODO: implement")
}

func (me *Node) GetPropProcessThreadMessages() int {
  panic("TODO: implement")
}

func (me *Node) SetPropProcessThreadMessages(value int) {
  panic("TODO: implement")
}

func (me *Node) GetPropEditorDescription() String {
  panic("TODO: implement")
}

func (me *Node) SetPropEditorDescription(value String) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API