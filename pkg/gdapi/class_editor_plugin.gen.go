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



// Enums

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

func (me *EditorPlugin) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorPlugin) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *EditorPlugin) XForwardCanvasGuiInput(event InputEvent, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) XForwardCanvasDrawOverViewport(viewport_control Control, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) XForwardCanvasForceDrawOverViewport(viewport_control Control, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) XForward3DGuiInput(viewport_camera Camera3D, event InputEvent, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) XForward3DDrawOverViewport(viewport_control Control, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) XForward3DForceDrawOverViewport(viewport_control Control, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) XGetPluginName()  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) XGetPluginIcon()  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) XHasMainScreen()  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) XMakeVisible(visible bool, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) XEdit(object Object, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) XHandles(object Object, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) XGetState()  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) XSetState(state Dictionary, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) XClear()  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) XSaveExternalData()  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) XApplyChanges()  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) XGetBreakpoints()  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) XSetWindowLayout(configuration ConfigFile, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) XGetWindowLayout(configuration ConfigFile, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) XBuild()  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) XEnablePlugin()  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) XDisablePlugin()  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) AddControlToContainer(container EditorPluginCustomControlContainer, control Control, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) AddControlToBottomPanel(control Control, title String, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) AddControlToDock(slot EditorPluginDockSlot, control Control, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) RemoveControlFromDocks(control Control, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) RemoveControlFromBottomPanel(control Control, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) RemoveControlFromContainer(container EditorPluginCustomControlContainer, control Control, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) AddToolMenuItem(name String, callable Callable, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) AddToolSubmenuItem(name String, submenu PopupMenu, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) RemoveToolMenuItem(name String, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) GetExportAsMenu()  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) AddCustomType(type_ String, base String, script Script, icon Texture2D, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) RemoveCustomType(type_ String, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) AddAutoloadSingleton(name String, path String, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) RemoveAutoloadSingleton(name String, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) UpdateOverlays()  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) MakeBottomPanelItemVisible(item Control, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) HideBottomPanel()  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) GetUndoRedo()  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) AddUndoRedoInspectorHookCallback(callable Callable, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) RemoveUndoRedoInspectorHookCallback(callable Callable, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) QueueSaveLayout()  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) AddTranslationParserPlugin(parser EditorTranslationParserPlugin, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) RemoveTranslationParserPlugin(parser EditorTranslationParserPlugin, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) AddImportPlugin(importer EditorImportPlugin, first_priority bool, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) RemoveImportPlugin(importer EditorImportPlugin, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) AddSceneFormatImporterPlugin(scene_format_importer EditorSceneFormatImporter, first_priority bool, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) RemoveSceneFormatImporterPlugin(scene_format_importer EditorSceneFormatImporter, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) AddScenePostImportPlugin(scene_import_plugin EditorScenePostImportPlugin, first_priority bool, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) RemoveScenePostImportPlugin(scene_import_plugin EditorScenePostImportPlugin, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) AddExportPlugin(plugin EditorExportPlugin, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) RemoveExportPlugin(plugin EditorExportPlugin, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) AddNode3DGizmoPlugin(plugin EditorNode3DGizmoPlugin, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) RemoveNode3DGizmoPlugin(plugin EditorNode3DGizmoPlugin, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) AddInspectorPlugin(plugin EditorInspectorPlugin, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) RemoveInspectorPlugin(plugin EditorInspectorPlugin, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) AddResourceConversionPlugin(plugin EditorResourceConversionPlugin, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) RemoveResourceConversionPlugin(plugin EditorResourceConversionPlugin, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) SetInputEventForwardingAlwaysEnabled()  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) SetForceDrawOverForwardingEnabled()  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) GetEditorInterface()  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) GetScriptCreateDialog()  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) AddDebuggerPlugin(script EditorDebuggerPlugin, )  {
  panic("TODO: implement")
}

func  (me *EditorPlugin) RemoveDebuggerPlugin(script EditorDebuggerPlugin, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
