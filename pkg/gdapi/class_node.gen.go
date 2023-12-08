// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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

const (
  NodeNOTIFICATION_ENTER_TREE = 10
  NodeNOTIFICATION_EXIT_TREE = 11
  NodeNOTIFICATION_MOVED_IN_PARENT = 12
  NodeNOTIFICATION_READY = 13
  NodeNOTIFICATION_PAUSED = 14
  NodeNOTIFICATION_UNPAUSED = 15
  NodeNOTIFICATION_PHYSICS_PROCESS = 16
  NodeNOTIFICATION_PROCESS = 17
  NodeNOTIFICATION_PARENTED = 18
  NodeNOTIFICATION_UNPARENTED = 19
  NodeNOTIFICATION_SCENE_INSTANTIATED = 20
  NodeNOTIFICATION_DRAG_BEGIN = 21
  NodeNOTIFICATION_DRAG_END = 22
  NodeNOTIFICATION_PATH_RENAMED = 23
  NodeNOTIFICATION_CHILD_ORDER_CHANGED = 24
  NodeNOTIFICATION_INTERNAL_PROCESS = 25
  NodeNOTIFICATION_INTERNAL_PHYSICS_PROCESS = 26
  NodeNOTIFICATION_POST_ENTER_TREE = 27
  NodeNOTIFICATION_DISABLED = 28
  NodeNOTIFICATION_ENABLED = 29
  NodeNOTIFICATION_NODE_RECACHE_REQUESTED = 30
  NodeNOTIFICATION_EDITOR_PRE_SAVE = 9001
  NodeNOTIFICATION_EDITOR_POST_SAVE = 9002
  NodeNOTIFICATION_WM_MOUSE_ENTER = 1002
  NodeNOTIFICATION_WM_MOUSE_EXIT = 1003
  NodeNOTIFICATION_WM_WINDOW_FOCUS_IN = 1004
  NodeNOTIFICATION_WM_WINDOW_FOCUS_OUT = 1005
  NodeNOTIFICATION_WM_CLOSE_REQUEST = 1006
  NodeNOTIFICATION_WM_GO_BACK_REQUEST = 1007
  NodeNOTIFICATION_WM_SIZE_CHANGED = 1008
  NodeNOTIFICATION_WM_DPI_CHANGE = 1009
  NodeNOTIFICATION_VP_MOUSE_ENTER = 1010
  NodeNOTIFICATION_VP_MOUSE_EXIT = 1011
  NodeNOTIFICATION_OS_MEMORY_WARNING = 2009
  NodeNOTIFICATION_TRANSLATION_CHANGED = 2010
  NodeNOTIFICATION_WM_ABOUT = 2011
  NodeNOTIFICATION_CRASH = 2012
  NodeNOTIFICATION_OS_IME_UPDATE = 2013
  NodeNOTIFICATION_APPLICATION_RESUMED = 2014
  NodeNOTIFICATION_APPLICATION_PAUSED = 2015
  NodeNOTIFICATION_APPLICATION_FOCUS_IN = 2016
  NodeNOTIFICATION_APPLICATION_FOCUS_OUT = 2017
  NodeNOTIFICATION_TEXT_SERVER_CHANGED = 2018
)

type NodeProcessMode int
const (
  NodeProcessModeInherit NodeProcessMode = 0
  NodeProcessModePausable NodeProcessMode = 1
  NodeProcessModeWhenPaused NodeProcessMode = 2
  NodeProcessModeAlways NodeProcessMode = 3
  NodeProcessModeDisabled NodeProcessMode = 4
)

type NodeProcessThreadGroup int
const (
  NodeProcessThreadGroupInherit NodeProcessThreadGroup = 0
  NodeProcessThreadGroupMainThread NodeProcessThreadGroup = 1
  NodeProcessThreadGroupSubThread NodeProcessThreadGroup = 2
)

type NodeProcessThreadMessages int
const (
  NodeFlagProcessThreadMessages NodeProcessThreadMessages = 1
  NodeFlagProcessThreadMessagesPhysics NodeProcessThreadMessages = 2
  NodeFlagProcessThreadMessagesAll NodeProcessThreadMessages = 3
)

type NodeDuplicateFlags int
const (
  NodeDuplicateSignals NodeDuplicateFlags = 1
  NodeDuplicateGroups NodeDuplicateFlags = 2
  NodeDuplicateScripts NodeDuplicateFlags = 4
  NodeDuplicateUseInstantiation NodeDuplicateFlags = 8
)

type NodeInternalMode int
const (
  NodeInternalModeDisabled NodeInternalMode = 0
  NodeInternalModeFront NodeInternalMode = 1
  NodeInternalModeBack NodeInternalMode = 2
)
