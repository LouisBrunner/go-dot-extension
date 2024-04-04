// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type EditorPlugin struct {
  Node
}

func (me *EditorPlugin) BaseClass() string {
  return "EditorPlugin"
}

func NewEditorPlugin() *EditorPlugin {
  str := StringNameFromStr("EditorPlugin") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorPlugin{}
  obj.SetBaseObject(objPtr)
  return obj
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

func (me *EditorPlugin) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorPlugin) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorPlugin) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EditorPlugin) AddControlToContainer(container EditorPluginCustomControlContainer, control Control, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_control_to_container")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3092750152) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&container) , control.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) AddControlToBottomPanel(control Control, title String, ) Button {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_control_to_bottom_panel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3526039376) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{control.AsCTypePtr(), title.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewButton()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorPlugin) AddControlToDock(slot EditorPluginDockSlot, control Control, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_control_to_dock")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3354871258) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot) , control.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) RemoveControlFromDocks(control Control, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_control_from_docks")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1496901182) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{control.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) RemoveControlFromBottomPanel(control Control, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_control_from_bottom_panel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1496901182) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{control.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) RemoveControlFromContainer(container EditorPluginCustomControlContainer, control Control, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_control_from_container")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3092750152) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&container) , control.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) AddToolMenuItem(name String, callable Callable, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_tool_menu_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2137474292) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), callable.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) AddToolSubmenuItem(name String, submenu PopupMenu, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_tool_submenu_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1019428915) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), submenu.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) RemoveToolMenuItem(name String, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_tool_menu_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) GetExportAsMenu() PopupMenu {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_export_as_menu")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1775878644) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPopupMenu()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorPlugin) AddCustomType(type_ String, base String, script Script, icon Texture2D, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_custom_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1986814599) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{type_.AsCTypePtr(), base.AsCTypePtr(), script.AsCTypePtr(), icon.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) RemoveCustomType(type_ String, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_custom_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{type_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) AddAutoloadSingleton(name String, path String, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_autoload_singleton")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3186203200) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) RemoveAutoloadSingleton(name String, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_autoload_singleton")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) UpdateOverlays() int64 {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("update_overlays")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorPlugin) MakeBottomPanelItemVisible(item Control, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_bottom_panel_item_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1496901182) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{item.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) HideBottomPanel()  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("hide_bottom_panel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) GetUndoRedo() EditorUndoRedoManager {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_undo_redo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 773492341) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewEditorUndoRedoManager()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorPlugin) AddUndoRedoInspectorHookCallback(callable Callable, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_undo_redo_inspector_hook_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1611583062) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{callable.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) RemoveUndoRedoInspectorHookCallback(callable Callable, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_undo_redo_inspector_hook_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1611583062) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{callable.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) QueueSaveLayout()  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("queue_save_layout")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) AddTranslationParserPlugin(parser EditorTranslationParserPlugin, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_translation_parser_plugin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3116463128) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{parser.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) RemoveTranslationParserPlugin(parser EditorTranslationParserPlugin, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_translation_parser_plugin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3116463128) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{parser.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) AddImportPlugin(importer EditorImportPlugin, first_priority bool, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_import_plugin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3113975762) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{importer.AsCTypePtr(), gdc.ConstTypePtr(&first_priority) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) RemoveImportPlugin(importer EditorImportPlugin, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_import_plugin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2312482773) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{importer.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) AddSceneFormatImporterPlugin(scene_format_importer EditorSceneFormatImporter, first_priority bool, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_scene_format_importer_plugin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2764104752) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{scene_format_importer.AsCTypePtr(), gdc.ConstTypePtr(&first_priority) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) RemoveSceneFormatImporterPlugin(scene_format_importer EditorSceneFormatImporter, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_scene_format_importer_plugin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2637776123) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{scene_format_importer.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) AddScenePostImportPlugin(scene_import_plugin EditorScenePostImportPlugin, first_priority bool, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_scene_post_import_plugin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3492436322) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{scene_import_plugin.AsCTypePtr(), gdc.ConstTypePtr(&first_priority) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) RemoveScenePostImportPlugin(scene_import_plugin EditorScenePostImportPlugin, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_scene_post_import_plugin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3045178206) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{scene_import_plugin.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) AddExportPlugin(plugin EditorExportPlugin, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_export_plugin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4095952207) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{plugin.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) RemoveExportPlugin(plugin EditorExportPlugin, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_export_plugin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4095952207) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{plugin.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) AddNode3DGizmoPlugin(plugin EditorNode3DGizmoPlugin, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_node_3d_gizmo_plugin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1541015022) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{plugin.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) RemoveNode3DGizmoPlugin(plugin EditorNode3DGizmoPlugin, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_node_3d_gizmo_plugin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1541015022) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{plugin.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) AddInspectorPlugin(plugin EditorInspectorPlugin, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_inspector_plugin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 546395733) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{plugin.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) RemoveInspectorPlugin(plugin EditorInspectorPlugin, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_inspector_plugin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 546395733) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{plugin.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) AddResourceConversionPlugin(plugin EditorResourceConversionPlugin, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_resource_conversion_plugin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2124849111) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{plugin.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) RemoveResourceConversionPlugin(plugin EditorResourceConversionPlugin, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_resource_conversion_plugin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2124849111) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{plugin.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) SetInputEventForwardingAlwaysEnabled()  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_input_event_forwarding_always_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) SetForceDrawOverForwardingEnabled()  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_force_draw_over_forwarding_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) GetEditorInterface() EditorInterface {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_editor_interface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4223731786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewEditorInterface()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorPlugin) GetScriptCreateDialog() ScriptCreateDialog {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_script_create_dialog")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3121871482) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewScriptCreateDialog()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorPlugin) AddDebuggerPlugin(script EditorDebuggerPlugin, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_debugger_plugin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3749880309) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{script.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) RemoveDebuggerPlugin(script EditorDebuggerPlugin, )  {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_debugger_plugin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3749880309) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{script.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorPlugin) GetPluginVersion() String {
  classNameV := StringNameFromStr("EditorPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_plugin_version")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals

type EditorPluginSceneChangedSignalFn func(scene_root Node, )

func (me *EditorPlugin) ConnectSceneChanged(subs SignalSubscribers, fn EditorPluginSceneChangedSignalFn) {
  sig := StringNameFromStr("scene_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorPlugin) DisconnectSceneChanged(subs SignalSubscribers, fn EditorPluginSceneChangedSignalFn) {
  sig := StringNameFromStr("scene_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorPluginSceneClosedSignalFn func(filepath String, )

func (me *EditorPlugin) ConnectSceneClosed(subs SignalSubscribers, fn EditorPluginSceneClosedSignalFn) {
  sig := StringNameFromStr("scene_closed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorPlugin) DisconnectSceneClosed(subs SignalSubscribers, fn EditorPluginSceneClosedSignalFn) {
  sig := StringNameFromStr("scene_closed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorPluginMainScreenChangedSignalFn func(screen_name String, )

func (me *EditorPlugin) ConnectMainScreenChanged(subs SignalSubscribers, fn EditorPluginMainScreenChangedSignalFn) {
  sig := StringNameFromStr("main_screen_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorPlugin) DisconnectMainScreenChanged(subs SignalSubscribers, fn EditorPluginMainScreenChangedSignalFn) {
  sig := StringNameFromStr("main_screen_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorPluginResourceSavedSignalFn func(resource Resource, )

func (me *EditorPlugin) ConnectResourceSaved(subs SignalSubscribers, fn EditorPluginResourceSavedSignalFn) {
  sig := StringNameFromStr("resource_saved")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorPlugin) DisconnectResourceSaved(subs SignalSubscribers, fn EditorPluginResourceSavedSignalFn) {
  sig := StringNameFromStr("resource_saved")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorPluginProjectSettingsChangedSignalFn func()

func (me *EditorPlugin) ConnectProjectSettingsChanged(subs SignalSubscribers, fn EditorPluginProjectSettingsChangedSignalFn) {
  sig := StringNameFromStr("project_settings_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorPlugin) DisconnectProjectSettingsChanged(subs SignalSubscribers, fn EditorPluginProjectSettingsChangedSignalFn) {
  sig := StringNameFromStr("project_settings_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
