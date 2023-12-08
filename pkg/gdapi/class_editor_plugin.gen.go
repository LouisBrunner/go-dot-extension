// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorPlugin struct {
  obj gdc.ObjectPtr
}

func (me *EditorPlugin) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorPlugin) BaseClass() string {
  return "EditorPlugin"
}

type EditorPluginCustomControlContainer int
const (
  EditorPluginContainerToolbar EditorPluginCustomControlContainer = 0
  EditorPluginContainerSpatialEditorMenu EditorPluginCustomControlContainer = 1
  EditorPluginContainerSpatialEditorSideLeft EditorPluginCustomControlContainer = 2
  EditorPluginContainerSpatialEditorSideRight EditorPluginCustomControlContainer = 3
  EditorPluginContainerSpatialEditorBottom EditorPluginCustomControlContainer = 4
  EditorPluginContainerCanvasEditorMenu EditorPluginCustomControlContainer = 5
  EditorPluginContainerCanvasEditorSideLeft EditorPluginCustomControlContainer = 6
  EditorPluginContainerCanvasEditorSideRight EditorPluginCustomControlContainer = 7
  EditorPluginContainerCanvasEditorBottom EditorPluginCustomControlContainer = 8
  EditorPluginContainerInspectorBottom EditorPluginCustomControlContainer = 9
  EditorPluginContainerProjectSettingTabLeft EditorPluginCustomControlContainer = 10
  EditorPluginContainerProjectSettingTabRight EditorPluginCustomControlContainer = 11
)

type EditorPluginDockSlot int
const (
  EditorPluginDockSlotLeftUl EditorPluginDockSlot = 0
  EditorPluginDockSlotLeftBl EditorPluginDockSlot = 1
  EditorPluginDockSlotLeftUr EditorPluginDockSlot = 2
  EditorPluginDockSlotLeftBr EditorPluginDockSlot = 3
  EditorPluginDockSlotRightUl EditorPluginDockSlot = 4
  EditorPluginDockSlotRightBl EditorPluginDockSlot = 5
  EditorPluginDockSlotRightUr EditorPluginDockSlot = 6
  EditorPluginDockSlotRightBr EditorPluginDockSlot = 7
  EditorPluginDockSlotMax EditorPluginDockSlot = 8
)

type EditorPluginAfterGUIInput int
const (
  EditorPluginAfterGuiInputPass EditorPluginAfterGUIInput = 0
  EditorPluginAfterGuiInputStop EditorPluginAfterGUIInput = 1
  EditorPluginAfterGuiInputCustom EditorPluginAfterGUIInput = 2
)
