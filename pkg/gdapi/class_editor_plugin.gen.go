// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  EditorPluginCustomControlContainerContainerToolbar EditorPluginCustomControlContainer = 0
  EditorPluginCustomControlContainerContainerSpatialEditorMenu EditorPluginCustomControlContainer = 1
  EditorPluginCustomControlContainerContainerSpatialEditorSideLeft EditorPluginCustomControlContainer = 2
  EditorPluginCustomControlContainerContainerSpatialEditorSideRight EditorPluginCustomControlContainer = 3
  EditorPluginCustomControlContainerContainerSpatialEditorBottom EditorPluginCustomControlContainer = 4
  EditorPluginCustomControlContainerContainerCanvasEditorMenu EditorPluginCustomControlContainer = 5
  EditorPluginCustomControlContainerContainerCanvasEditorSideLeft EditorPluginCustomControlContainer = 6
  EditorPluginCustomControlContainerContainerCanvasEditorSideRight EditorPluginCustomControlContainer = 7
  EditorPluginCustomControlContainerContainerCanvasEditorBottom EditorPluginCustomControlContainer = 8
  EditorPluginCustomControlContainerContainerInspectorBottom EditorPluginCustomControlContainer = 9
  EditorPluginCustomControlContainerContainerProjectSettingTabLeft EditorPluginCustomControlContainer = 10
  EditorPluginCustomControlContainerContainerProjectSettingTabRight EditorPluginCustomControlContainer = 11
)

type EditorPluginDockSlot int
const (
  EditorPluginDockSlotDockSlotLeftUl EditorPluginDockSlot = 0
  EditorPluginDockSlotDockSlotLeftBl EditorPluginDockSlot = 1
  EditorPluginDockSlotDockSlotLeftUr EditorPluginDockSlot = 2
  EditorPluginDockSlotDockSlotLeftBr EditorPluginDockSlot = 3
  EditorPluginDockSlotDockSlotRightUl EditorPluginDockSlot = 4
  EditorPluginDockSlotDockSlotRightBl EditorPluginDockSlot = 5
  EditorPluginDockSlotDockSlotRightUr EditorPluginDockSlot = 6
  EditorPluginDockSlotDockSlotRightBr EditorPluginDockSlot = 7
  EditorPluginDockSlotDockSlotMax EditorPluginDockSlot = 8
)

type EditorPluginAfterGUIInput int
const (
  EditorPluginAfterGUIInputAfterGuiInputPass EditorPluginAfterGUIInput = 0
  EditorPluginAfterGUIInputAfterGuiInputStop EditorPluginAfterGUIInput = 1
  EditorPluginAfterGUIInputAfterGuiInputCustom EditorPluginAfterGUIInput = 2
)

func  (me *EditorPlugin) XForwardCanvasGuiInput(event InputEvent, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) XForwardCanvasDrawOverViewport(viewport_control Control, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) XForwardCanvasForceDrawOverViewport(viewport_control Control, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) XForward3DGuiInput(viewport_camera Camera3D, event InputEvent, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) XForward3DDrawOverViewport(viewport_control Control, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) XForward3DForceDrawOverViewport(viewport_control Control, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) XGetPluginName() { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) XGetPluginIcon() { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) XHasMainScreen() { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) XMakeVisible(visible bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) XEdit(object Object, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) XHandles(object Object, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) XGetState() { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) XSetState(state Dictionary, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) XClear() { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) XSaveExternalData() { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) XApplyChanges() { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) XGetBreakpoints() { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) XSetWindowLayout(configuration ConfigFile, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) XGetWindowLayout(configuration ConfigFile, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) XBuild() { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) XEnablePlugin() { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) XDisablePlugin() { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) AddControlToContainer(container EditorPluginCustomControlContainer, control Control, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) AddControlToBottomPanel(control Control, title String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) AddControlToDock(slot EditorPluginDockSlot, control Control, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) RemoveControlFromDocks(control Control, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) RemoveControlFromBottomPanel(control Control, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) RemoveControlFromContainer(container EditorPluginCustomControlContainer, control Control, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) AddToolMenuItem(name String, callable Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) AddToolSubmenuItem(name String, submenu PopupMenu, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) RemoveToolMenuItem(name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) GetExportAsMenu() { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) AddCustomType(type_ String, base String, script Script, icon Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) RemoveCustomType(type_ String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) AddAutoloadSingleton(name String, path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) RemoveAutoloadSingleton(name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) UpdateOverlays() { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) MakeBottomPanelItemVisible(item Control, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) HideBottomPanel() { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) GetUndoRedo() { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) AddUndoRedoInspectorHookCallback(callable Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) RemoveUndoRedoInspectorHookCallback(callable Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) QueueSaveLayout() { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) AddTranslationParserPlugin(parser EditorTranslationParserPlugin, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) RemoveTranslationParserPlugin(parser EditorTranslationParserPlugin, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) AddImportPlugin(importer EditorImportPlugin, first_priority bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) RemoveImportPlugin(importer EditorImportPlugin, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) AddSceneFormatImporterPlugin(scene_format_importer EditorSceneFormatImporter, first_priority bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) RemoveSceneFormatImporterPlugin(scene_format_importer EditorSceneFormatImporter, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) AddScenePostImportPlugin(scene_import_plugin EditorScenePostImportPlugin, first_priority bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) RemoveScenePostImportPlugin(scene_import_plugin EditorScenePostImportPlugin, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) AddExportPlugin(plugin EditorExportPlugin, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) RemoveExportPlugin(plugin EditorExportPlugin, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) AddNode3DGizmoPlugin(plugin EditorNode3DGizmoPlugin, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) RemoveNode3DGizmoPlugin(plugin EditorNode3DGizmoPlugin, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) AddInspectorPlugin(plugin EditorInspectorPlugin, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) RemoveInspectorPlugin(plugin EditorInspectorPlugin, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) AddResourceConversionPlugin(plugin EditorResourceConversionPlugin, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) RemoveResourceConversionPlugin(plugin EditorResourceConversionPlugin, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) SetInputEventForwardingAlwaysEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) SetForceDrawOverForwardingEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) GetEditorInterface() { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) GetScriptCreateDialog() { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) AddDebuggerPlugin(script EditorDebuggerPlugin, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorPlugin) RemoveDebuggerPlugin(script EditorDebuggerPlugin, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
