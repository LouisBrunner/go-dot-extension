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

// TODO: needed?
// const (
// // )

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

func  (me *Node) XProcess(delta float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) XPhysicsProcess(delta float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) XEnterTree() { // TODO: return value
  // TODO: implement
}

func  (me *Node) XExitTree() { // TODO: return value
  // TODO: implement
}

func  (me *Node) XReady() { // TODO: return value
  // TODO: implement
}

func  (me *Node) XGetConfigurationWarnings() { // TODO: return value
  // TODO: implement
}

func  (me *Node) XInput(event InputEvent, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) XShortcutInput(event InputEvent, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) XUnhandledInput(event InputEvent, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) XUnhandledKeyInput(event InputEvent, ) { // TODO: return value
  // TODO: implement
}

func  NodePrintOrphanNodes() { // TODO: return value
  // TODO: implement
}

func  (me *Node) AddSibling(sibling Node, force_readable_name bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) SetName(name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetName() { // TODO: return value
  // TODO: implement
}

func  (me *Node) AddChild(node Node, force_readable_name bool, internal NodeInternalMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) RemoveChild(node Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) Reparent(new_parent Node, keep_global_transform bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetChildCount(include_internal bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetChildren(include_internal bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetChild(idx int, include_internal bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) HasNode(path NodePath, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetNode(path NodePath, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetNodeOrNull(path NodePath, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetParent() { // TODO: return value
  // TODO: implement
}

func  (me *Node) FindChild(pattern String, recursive bool, owned bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) FindChildren(pattern String, type_ String, recursive bool, owned bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) FindParent(pattern String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) HasNodeAndResource(path NodePath, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetNodeAndResource(path NodePath, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) IsInsideTree() { // TODO: return value
  // TODO: implement
}

func  (me *Node) IsAncestorOf(node Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) IsGreaterThan(node Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetPath() { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetPathTo(node Node, use_unique_path bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) AddToGroup(group StringName, persistent bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) RemoveFromGroup(group StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) IsInGroup(group StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) MoveChild(child_node Node, to_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetGroups() { // TODO: return value
  // TODO: implement
}

func  (me *Node) SetOwner(owner Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetOwner() { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetIndex(include_internal bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) PrintTree() { // TODO: return value
  // TODO: implement
}

func  (me *Node) PrintTreePretty() { // TODO: return value
  // TODO: implement
}

func  (me *Node) SetSceneFilePath(scene_file_path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetSceneFilePath() { // TODO: return value
  // TODO: implement
}

func  (me *Node) PropagateNotification(what int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) PropagateCall(method StringName, args Array, parent_first bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) SetPhysicsProcess(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetPhysicsProcessDeltaTime() { // TODO: return value
  // TODO: implement
}

func  (me *Node) IsPhysicsProcessing() { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetProcessDeltaTime() { // TODO: return value
  // TODO: implement
}

func  (me *Node) SetProcess(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) SetProcessPriority(priority int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetProcessPriority() { // TODO: return value
  // TODO: implement
}

func  (me *Node) SetPhysicsProcessPriority(priority int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetPhysicsProcessPriority() { // TODO: return value
  // TODO: implement
}

func  (me *Node) IsProcessing() { // TODO: return value
  // TODO: implement
}

func  (me *Node) SetProcessInput(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) IsProcessingInput() { // TODO: return value
  // TODO: implement
}

func  (me *Node) SetProcessShortcutInput(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) IsProcessingShortcutInput() { // TODO: return value
  // TODO: implement
}

func  (me *Node) SetProcessUnhandledInput(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) IsProcessingUnhandledInput() { // TODO: return value
  // TODO: implement
}

func  (me *Node) SetProcessUnhandledKeyInput(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) IsProcessingUnhandledKeyInput() { // TODO: return value
  // TODO: implement
}

func  (me *Node) SetProcessMode(mode NodeProcessMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetProcessMode() { // TODO: return value
  // TODO: implement
}

func  (me *Node) CanProcess() { // TODO: return value
  // TODO: implement
}

func  (me *Node) SetProcessThreadGroup(mode NodeProcessThreadGroup, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetProcessThreadGroup() { // TODO: return value
  // TODO: implement
}

func  (me *Node) SetProcessThreadMessages(flags NodeProcessThreadMessages, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetProcessThreadMessages() { // TODO: return value
  // TODO: implement
}

func  (me *Node) SetProcessThreadGroupOrder(order int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetProcessThreadGroupOrder() { // TODO: return value
  // TODO: implement
}

func  (me *Node) SetDisplayFolded(fold bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) IsDisplayedFolded() { // TODO: return value
  // TODO: implement
}

func  (me *Node) SetProcessInternal(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) IsProcessingInternal() { // TODO: return value
  // TODO: implement
}

func  (me *Node) SetPhysicsProcessInternal(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) IsPhysicsProcessingInternal() { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetWindow() { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetLastExclusiveWindow() { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetTree() { // TODO: return value
  // TODO: implement
}

func  (me *Node) CreateTween() { // TODO: return value
  // TODO: implement
}

func  (me *Node) Duplicate(flags int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) ReplaceBy(node Node, keep_groups bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) SetSceneInstanceLoadPlaceholder(load_placeholder bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetSceneInstanceLoadPlaceholder() { // TODO: return value
  // TODO: implement
}

func  (me *Node) SetEditableInstance(node Node, is_editable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) IsEditableInstance(node Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetViewport() { // TODO: return value
  // TODO: implement
}

func  (me *Node) QueueFree() { // TODO: return value
  // TODO: implement
}

func  (me *Node) RequestReady() { // TODO: return value
  // TODO: implement
}

func  (me *Node) IsNodeReady() { // TODO: return value
  // TODO: implement
}

func  (me *Node) SetMultiplayerAuthority(id int, recursive bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetMultiplayerAuthority() { // TODO: return value
  // TODO: implement
}

func  (me *Node) IsMultiplayerAuthority() { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetMultiplayer() { // TODO: return value
  // TODO: implement
}

func  (me *Node) RpcConfig(method StringName, config Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) SetEditorDescription(editor_description String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) GetEditorDescription() { // TODO: return value
  // TODO: implement
}

func  (me *Node) SetUniqueNameInOwner(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) IsUniqueNameInOwner() { // TODO: return value
  // TODO: implement
}

func  (me *Node) Rpc(method StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) RpcId(peer_id int, method StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) UpdateConfigurationWarnings() { // TODO: return value
  // TODO: implement
}

func  (me *Node) CallDeferredThreadGroup(method StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) SetDeferredThreadGroup(property StringName, value Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) NotifyDeferredThreadGroup(what int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) CallThreadSafe(method StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) SetThreadSafe(property StringName, value Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node) NotifyThreadSafe(what int, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
