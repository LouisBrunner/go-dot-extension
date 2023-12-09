// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func (me *Node) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Node) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Node) XProcess(delta float32, )  {
  panic("TODO: implement")
}

func  (me *Node) XPhysicsProcess(delta float32, )  {
  panic("TODO: implement")
}

func  (me *Node) XEnterTree()  {
  panic("TODO: implement")
}

func  (me *Node) XExitTree()  {
  panic("TODO: implement")
}

func  (me *Node) XReady()  {
  panic("TODO: implement")
}

func  (me *Node) XGetConfigurationWarnings()  {
  panic("TODO: implement")
}

func  (me *Node) XInput(event InputEvent, )  {
  panic("TODO: implement")
}

func  (me *Node) XShortcutInput(event InputEvent, )  {
  panic("TODO: implement")
}

func  (me *Node) XUnhandledInput(event InputEvent, )  {
  panic("TODO: implement")
}

func  (me *Node) XUnhandledKeyInput(event InputEvent, )  {
  panic("TODO: implement")
}

func  NodePrintOrphanNodes()  {
  panic("TODO: implement")
}

func  (me *Node) AddSibling(sibling Node, force_readable_name bool, )  {
  panic("TODO: implement")
}

func  (me *Node) SetName(name String, )  {
  panic("TODO: implement")
}

func  (me *Node) GetName()  {
  panic("TODO: implement")
}

func  (me *Node) AddChild(node Node, force_readable_name bool, internal NodeInternalMode, )  {
  panic("TODO: implement")
}

func  (me *Node) RemoveChild(node Node, )  {
  panic("TODO: implement")
}

func  (me *Node) Reparent(new_parent Node, keep_global_transform bool, )  {
  panic("TODO: implement")
}

func  (me *Node) GetChildCount(include_internal bool, )  {
  panic("TODO: implement")
}

func  (me *Node) GetChildren(include_internal bool, )  {
  panic("TODO: implement")
}

func  (me *Node) GetChild(idx int, include_internal bool, )  {
  panic("TODO: implement")
}

func  (me *Node) HasNode(path NodePath, )  {
  panic("TODO: implement")
}

func  (me *Node) GetNode(path NodePath, )  {
  panic("TODO: implement")
}

func  (me *Node) GetNodeOrNull(path NodePath, )  {
  panic("TODO: implement")
}

func  (me *Node) GetParent()  {
  panic("TODO: implement")
}

func  (me *Node) FindChild(pattern String, recursive bool, owned bool, )  {
  panic("TODO: implement")
}

func  (me *Node) FindChildren(pattern String, type_ String, recursive bool, owned bool, )  {
  panic("TODO: implement")
}

func  (me *Node) FindParent(pattern String, )  {
  panic("TODO: implement")
}

func  (me *Node) HasNodeAndResource(path NodePath, )  {
  panic("TODO: implement")
}

func  (me *Node) GetNodeAndResource(path NodePath, )  {
  panic("TODO: implement")
}

func  (me *Node) IsInsideTree()  {
  panic("TODO: implement")
}

func  (me *Node) IsAncestorOf(node Node, )  {
  panic("TODO: implement")
}

func  (me *Node) IsGreaterThan(node Node, )  {
  panic("TODO: implement")
}

func  (me *Node) GetPath()  {
  panic("TODO: implement")
}

func  (me *Node) GetPathTo(node Node, use_unique_path bool, )  {
  panic("TODO: implement")
}

func  (me *Node) AddToGroup(group StringName, persistent bool, )  {
  panic("TODO: implement")
}

func  (me *Node) RemoveFromGroup(group StringName, )  {
  panic("TODO: implement")
}

func  (me *Node) IsInGroup(group StringName, )  {
  panic("TODO: implement")
}

func  (me *Node) MoveChild(child_node Node, to_index int, )  {
  panic("TODO: implement")
}

func  (me *Node) GetGroups()  {
  panic("TODO: implement")
}

func  (me *Node) SetOwner(owner Node, )  {
  panic("TODO: implement")
}

func  (me *Node) GetOwner()  {
  panic("TODO: implement")
}

func  (me *Node) GetIndex(include_internal bool, )  {
  panic("TODO: implement")
}

func  (me *Node) PrintTree()  {
  panic("TODO: implement")
}

func  (me *Node) PrintTreePretty()  {
  panic("TODO: implement")
}

func  (me *Node) SetSceneFilePath(scene_file_path String, )  {
  panic("TODO: implement")
}

func  (me *Node) GetSceneFilePath()  {
  panic("TODO: implement")
}

func  (me *Node) PropagateNotification(what int, )  {
  panic("TODO: implement")
}

func  (me *Node) PropagateCall(method StringName, args Array, parent_first bool, )  {
  panic("TODO: implement")
}

func  (me *Node) SetPhysicsProcess(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Node) GetPhysicsProcessDeltaTime()  {
  panic("TODO: implement")
}

func  (me *Node) IsPhysicsProcessing()  {
  panic("TODO: implement")
}

func  (me *Node) GetProcessDeltaTime()  {
  panic("TODO: implement")
}

func  (me *Node) SetProcess(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Node) SetProcessPriority(priority int, )  {
  panic("TODO: implement")
}

func  (me *Node) GetProcessPriority()  {
  panic("TODO: implement")
}

func  (me *Node) SetPhysicsProcessPriority(priority int, )  {
  panic("TODO: implement")
}

func  (me *Node) GetPhysicsProcessPriority()  {
  panic("TODO: implement")
}

func  (me *Node) IsProcessing()  {
  panic("TODO: implement")
}

func  (me *Node) SetProcessInput(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Node) IsProcessingInput()  {
  panic("TODO: implement")
}

func  (me *Node) SetProcessShortcutInput(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Node) IsProcessingShortcutInput()  {
  panic("TODO: implement")
}

func  (me *Node) SetProcessUnhandledInput(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Node) IsProcessingUnhandledInput()  {
  panic("TODO: implement")
}

func  (me *Node) SetProcessUnhandledKeyInput(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Node) IsProcessingUnhandledKeyInput()  {
  panic("TODO: implement")
}

func  (me *Node) SetProcessMode(mode NodeProcessMode, )  {
  panic("TODO: implement")
}

func  (me *Node) GetProcessMode()  {
  panic("TODO: implement")
}

func  (me *Node) CanProcess()  {
  panic("TODO: implement")
}

func  (me *Node) SetProcessThreadGroup(mode NodeProcessThreadGroup, )  {
  panic("TODO: implement")
}

func  (me *Node) GetProcessThreadGroup()  {
  panic("TODO: implement")
}

func  (me *Node) SetProcessThreadMessages(flags NodeProcessThreadMessages, )  {
  panic("TODO: implement")
}

func  (me *Node) GetProcessThreadMessages()  {
  panic("TODO: implement")
}

func  (me *Node) SetProcessThreadGroupOrder(order int, )  {
  panic("TODO: implement")
}

func  (me *Node) GetProcessThreadGroupOrder()  {
  panic("TODO: implement")
}

func  (me *Node) SetDisplayFolded(fold bool, )  {
  panic("TODO: implement")
}

func  (me *Node) IsDisplayedFolded()  {
  panic("TODO: implement")
}

func  (me *Node) SetProcessInternal(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Node) IsProcessingInternal()  {
  panic("TODO: implement")
}

func  (me *Node) SetPhysicsProcessInternal(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Node) IsPhysicsProcessingInternal()  {
  panic("TODO: implement")
}

func  (me *Node) GetWindow()  {
  panic("TODO: implement")
}

func  (me *Node) GetLastExclusiveWindow()  {
  panic("TODO: implement")
}

func  (me *Node) GetTree()  {
  panic("TODO: implement")
}

func  (me *Node) CreateTween()  {
  panic("TODO: implement")
}

func  (me *Node) Duplicate(flags int, )  {
  panic("TODO: implement")
}

func  (me *Node) ReplaceBy(node Node, keep_groups bool, )  {
  panic("TODO: implement")
}

func  (me *Node) SetSceneInstanceLoadPlaceholder(load_placeholder bool, )  {
  panic("TODO: implement")
}

func  (me *Node) GetSceneInstanceLoadPlaceholder()  {
  panic("TODO: implement")
}

func  (me *Node) SetEditableInstance(node Node, is_editable bool, )  {
  panic("TODO: implement")
}

func  (me *Node) IsEditableInstance(node Node, )  {
  panic("TODO: implement")
}

func  (me *Node) GetViewport()  {
  panic("TODO: implement")
}

func  (me *Node) QueueFree()  {
  panic("TODO: implement")
}

func  (me *Node) RequestReady()  {
  panic("TODO: implement")
}

func  (me *Node) IsNodeReady()  {
  panic("TODO: implement")
}

func  (me *Node) SetMultiplayerAuthority(id int, recursive bool, )  {
  panic("TODO: implement")
}

func  (me *Node) GetMultiplayerAuthority()  {
  panic("TODO: implement")
}

func  (me *Node) IsMultiplayerAuthority()  {
  panic("TODO: implement")
}

func  (me *Node) GetMultiplayer()  {
  panic("TODO: implement")
}

func  (me *Node) RpcConfig(method StringName, config Variant, )  {
  panic("TODO: implement")
}

func  (me *Node) SetEditorDescription(editor_description String, )  {
  panic("TODO: implement")
}

func  (me *Node) GetEditorDescription()  {
  panic("TODO: implement")
}

func  (me *Node) SetUniqueNameInOwner(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Node) IsUniqueNameInOwner()  {
  panic("TODO: implement")
}

func  (me *Node) Rpc(method StringName, )  {
  panic("TODO: implement")
}

func  (me *Node) RpcId(peer_id int, method StringName, )  {
  panic("TODO: implement")
}

func  (me *Node) UpdateConfigurationWarnings()  {
  panic("TODO: implement")
}

func  (me *Node) CallDeferredThreadGroup(method StringName, )  {
  panic("TODO: implement")
}

func  (me *Node) SetDeferredThreadGroup(property StringName, value Variant, )  {
  panic("TODO: implement")
}

func  (me *Node) NotifyDeferredThreadGroup(what int, )  {
  panic("TODO: implement")
}

func  (me *Node) CallThreadSafe(method StringName, )  {
  panic("TODO: implement")
}

func  (me *Node) SetThreadSafe(property StringName, value Variant, )  {
  panic("TODO: implement")
}

func  (me *Node) NotifyThreadSafe(what int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
