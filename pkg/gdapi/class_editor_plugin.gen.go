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

type ptrsForEditorPluginList struct {
	fnXForwardCanvasGuiInput               gdc.MethodBindPtr
	fnXForwardCanvasDrawOverViewport       gdc.MethodBindPtr
	fnXForwardCanvasForceDrawOverViewport  gdc.MethodBindPtr
	fnXForward3DGuiInput                   gdc.MethodBindPtr
	fnXForward3DDrawOverViewport           gdc.MethodBindPtr
	fnXForward3DForceDrawOverViewport      gdc.MethodBindPtr
	fnXGetPluginName                       gdc.MethodBindPtr
	fnXGetPluginIcon                       gdc.MethodBindPtr
	fnXHasMainScreen                       gdc.MethodBindPtr
	fnXMakeVisible                         gdc.MethodBindPtr
	fnXEdit                                gdc.MethodBindPtr
	fnXHandles                             gdc.MethodBindPtr
	fnXGetState                            gdc.MethodBindPtr
	fnXSetState                            gdc.MethodBindPtr
	fnXClear                               gdc.MethodBindPtr
	fnXGetUnsavedStatus                    gdc.MethodBindPtr
	fnXSaveExternalData                    gdc.MethodBindPtr
	fnXApplyChanges                        gdc.MethodBindPtr
	fnXGetBreakpoints                      gdc.MethodBindPtr
	fnXSetWindowLayout                     gdc.MethodBindPtr
	fnXGetWindowLayout                     gdc.MethodBindPtr
	fnXBuild                               gdc.MethodBindPtr
	fnXEnablePlugin                        gdc.MethodBindPtr
	fnXDisablePlugin                       gdc.MethodBindPtr
	fnAddControlToContainer                gdc.MethodBindPtr
	fnAddControlToBottomPanel              gdc.MethodBindPtr
	fnAddControlToDock                     gdc.MethodBindPtr
	fnRemoveControlFromDocks               gdc.MethodBindPtr
	fnRemoveControlFromBottomPanel         gdc.MethodBindPtr
	fnRemoveControlFromContainer           gdc.MethodBindPtr
	fnAddToolMenuItem                      gdc.MethodBindPtr
	fnAddToolSubmenuItem                   gdc.MethodBindPtr
	fnRemoveToolMenuItem                   gdc.MethodBindPtr
	fnGetExportAsMenu                      gdc.MethodBindPtr
	fnAddCustomType                        gdc.MethodBindPtr
	fnRemoveCustomType                     gdc.MethodBindPtr
	fnAddAutoloadSingleton                 gdc.MethodBindPtr
	fnRemoveAutoloadSingleton              gdc.MethodBindPtr
	fnUpdateOverlays                       gdc.MethodBindPtr
	fnMakeBottomPanelItemVisible           gdc.MethodBindPtr
	fnHideBottomPanel                      gdc.MethodBindPtr
	fnGetUndoRedo                          gdc.MethodBindPtr
	fnAddUndoRedoInspectorHookCallback     gdc.MethodBindPtr
	fnRemoveUndoRedoInspectorHookCallback  gdc.MethodBindPtr
	fnQueueSaveLayout                      gdc.MethodBindPtr
	fnAddTranslationParserPlugin           gdc.MethodBindPtr
	fnRemoveTranslationParserPlugin        gdc.MethodBindPtr
	fnAddImportPlugin                      gdc.MethodBindPtr
	fnRemoveImportPlugin                   gdc.MethodBindPtr
	fnAddSceneFormatImporterPlugin         gdc.MethodBindPtr
	fnRemoveSceneFormatImporterPlugin      gdc.MethodBindPtr
	fnAddScenePostImportPlugin             gdc.MethodBindPtr
	fnRemoveScenePostImportPlugin          gdc.MethodBindPtr
	fnAddExportPlugin                      gdc.MethodBindPtr
	fnRemoveExportPlugin                   gdc.MethodBindPtr
	fnAddNode3DGizmoPlugin                 gdc.MethodBindPtr
	fnRemoveNode3DGizmoPlugin              gdc.MethodBindPtr
	fnAddInspectorPlugin                   gdc.MethodBindPtr
	fnRemoveInspectorPlugin                gdc.MethodBindPtr
	fnAddResourceConversionPlugin          gdc.MethodBindPtr
	fnRemoveResourceConversionPlugin       gdc.MethodBindPtr
	fnSetInputEventForwardingAlwaysEnabled gdc.MethodBindPtr
	fnSetForceDrawOverForwardingEnabled    gdc.MethodBindPtr
	fnGetEditorInterface                   gdc.MethodBindPtr
	fnGetScriptCreateDialog                gdc.MethodBindPtr
	fnAddDebuggerPlugin                    gdc.MethodBindPtr
	fnRemoveDebuggerPlugin                 gdc.MethodBindPtr
	fnGetPluginVersion                     gdc.MethodBindPtr
}

var ptrsForEditorPlugin ptrsForEditorPluginList

func initEditorPluginPtrs(iface gdc.Interface) {

	className := StringNameFromStr("EditorPlugin")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("add_control_to_container")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnAddControlToContainer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3092750152))
	}
	{
		methodName := StringNameFromStr("add_control_to_bottom_panel")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnAddControlToBottomPanel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3526039376))
	}
	{
		methodName := StringNameFromStr("add_control_to_dock")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnAddControlToDock = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3354871258))
	}
	{
		methodName := StringNameFromStr("remove_control_from_docks")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnRemoveControlFromDocks = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1496901182))
	}
	{
		methodName := StringNameFromStr("remove_control_from_bottom_panel")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnRemoveControlFromBottomPanel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1496901182))
	}
	{
		methodName := StringNameFromStr("remove_control_from_container")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnRemoveControlFromContainer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3092750152))
	}
	{
		methodName := StringNameFromStr("add_tool_menu_item")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnAddToolMenuItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2137474292))
	}
	{
		methodName := StringNameFromStr("add_tool_submenu_item")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnAddToolSubmenuItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1019428915))
	}
	{
		methodName := StringNameFromStr("remove_tool_menu_item")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnRemoveToolMenuItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_export_as_menu")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnGetExportAsMenu = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1775878644))
	}
	{
		methodName := StringNameFromStr("add_custom_type")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnAddCustomType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1986814599))
	}
	{
		methodName := StringNameFromStr("remove_custom_type")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnRemoveCustomType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("add_autoload_singleton")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnAddAutoloadSingleton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3186203200))
	}
	{
		methodName := StringNameFromStr("remove_autoload_singleton")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnRemoveAutoloadSingleton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("update_overlays")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnUpdateOverlays = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("make_bottom_panel_item_visible")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnMakeBottomPanelItemVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1496901182))
	}
	{
		methodName := StringNameFromStr("hide_bottom_panel")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnHideBottomPanel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_undo_redo")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnGetUndoRedo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 773492341))
	}
	{
		methodName := StringNameFromStr("add_undo_redo_inspector_hook_callback")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnAddUndoRedoInspectorHookCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1611583062))
	}
	{
		methodName := StringNameFromStr("remove_undo_redo_inspector_hook_callback")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnRemoveUndoRedoInspectorHookCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1611583062))
	}
	{
		methodName := StringNameFromStr("queue_save_layout")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnQueueSaveLayout = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("add_translation_parser_plugin")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnAddTranslationParserPlugin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3116463128))
	}
	{
		methodName := StringNameFromStr("remove_translation_parser_plugin")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnRemoveTranslationParserPlugin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3116463128))
	}
	{
		methodName := StringNameFromStr("add_import_plugin")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnAddImportPlugin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3113975762))
	}
	{
		methodName := StringNameFromStr("remove_import_plugin")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnRemoveImportPlugin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2312482773))
	}
	{
		methodName := StringNameFromStr("add_scene_format_importer_plugin")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnAddSceneFormatImporterPlugin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2764104752))
	}
	{
		methodName := StringNameFromStr("remove_scene_format_importer_plugin")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnRemoveSceneFormatImporterPlugin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2637776123))
	}
	{
		methodName := StringNameFromStr("add_scene_post_import_plugin")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnAddScenePostImportPlugin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3492436322))
	}
	{
		methodName := StringNameFromStr("remove_scene_post_import_plugin")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnRemoveScenePostImportPlugin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3045178206))
	}
	{
		methodName := StringNameFromStr("add_export_plugin")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnAddExportPlugin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4095952207))
	}
	{
		methodName := StringNameFromStr("remove_export_plugin")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnRemoveExportPlugin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4095952207))
	}
	{
		methodName := StringNameFromStr("add_node_3d_gizmo_plugin")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnAddNode3DGizmoPlugin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1541015022))
	}
	{
		methodName := StringNameFromStr("remove_node_3d_gizmo_plugin")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnRemoveNode3DGizmoPlugin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1541015022))
	}
	{
		methodName := StringNameFromStr("add_inspector_plugin")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnAddInspectorPlugin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 546395733))
	}
	{
		methodName := StringNameFromStr("remove_inspector_plugin")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnRemoveInspectorPlugin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 546395733))
	}
	{
		methodName := StringNameFromStr("add_resource_conversion_plugin")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnAddResourceConversionPlugin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2124849111))
	}
	{
		methodName := StringNameFromStr("remove_resource_conversion_plugin")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnRemoveResourceConversionPlugin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2124849111))
	}
	{
		methodName := StringNameFromStr("set_input_event_forwarding_always_enabled")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnSetInputEventForwardingAlwaysEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_force_draw_over_forwarding_enabled")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnSetForceDrawOverForwardingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_editor_interface")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnGetEditorInterface = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4223731786))
	}
	{
		methodName := StringNameFromStr("get_script_create_dialog")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnGetScriptCreateDialog = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3121871482))
	}
	{
		methodName := StringNameFromStr("add_debugger_plugin")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnAddDebuggerPlugin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3749880309))
	}
	{
		methodName := StringNameFromStr("remove_debugger_plugin")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnRemoveDebuggerPlugin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3749880309))
	}
	{
		methodName := StringNameFromStr("get_plugin_version")
		defer methodName.Destroy()
		ptrsForEditorPlugin.fnGetPluginVersion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}

}

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
	EditorPluginCustomControlContainerContainerToolbar                EditorPluginCustomControlContainer = 0
	EditorPluginCustomControlContainerContainerSpatialEditorMenu      EditorPluginCustomControlContainer = 1
	EditorPluginCustomControlContainerContainerSpatialEditorSideLeft  EditorPluginCustomControlContainer = 2
	EditorPluginCustomControlContainerContainerSpatialEditorSideRight EditorPluginCustomControlContainer = 3
	EditorPluginCustomControlContainerContainerSpatialEditorBottom    EditorPluginCustomControlContainer = 4
	EditorPluginCustomControlContainerContainerCanvasEditorMenu       EditorPluginCustomControlContainer = 5
	EditorPluginCustomControlContainerContainerCanvasEditorSideLeft   EditorPluginCustomControlContainer = 6
	EditorPluginCustomControlContainerContainerCanvasEditorSideRight  EditorPluginCustomControlContainer = 7
	EditorPluginCustomControlContainerContainerCanvasEditorBottom     EditorPluginCustomControlContainer = 8
	EditorPluginCustomControlContainerContainerInspectorBottom        EditorPluginCustomControlContainer = 9
	EditorPluginCustomControlContainerContainerProjectSettingTabLeft  EditorPluginCustomControlContainer = 10
	EditorPluginCustomControlContainerContainerProjectSettingTabRight EditorPluginCustomControlContainer = 11
)

type EditorPluginDockSlot int

const (
	EditorPluginDockSlotDockSlotLeftUl  EditorPluginDockSlot = 0
	EditorPluginDockSlotDockSlotLeftBl  EditorPluginDockSlot = 1
	EditorPluginDockSlotDockSlotLeftUr  EditorPluginDockSlot = 2
	EditorPluginDockSlotDockSlotLeftBr  EditorPluginDockSlot = 3
	EditorPluginDockSlotDockSlotRightUl EditorPluginDockSlot = 4
	EditorPluginDockSlotDockSlotRightBl EditorPluginDockSlot = 5
	EditorPluginDockSlotDockSlotRightUr EditorPluginDockSlot = 6
	EditorPluginDockSlotDockSlotRightBr EditorPluginDockSlot = 7
	EditorPluginDockSlotDockSlotMax     EditorPluginDockSlot = 8
)

type EditorPluginAfterGUIInput int

const (
	EditorPluginAfterGUIInputAfterGuiInputPass   EditorPluginAfterGUIInput = 0
	EditorPluginAfterGUIInputAfterGuiInputStop   EditorPluginAfterGUIInput = 1
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

func (me *EditorPlugin) AddControlToContainer(container EditorPluginCustomControlContainer, control Control) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&container), control.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnAddControlToContainer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) AddControlToBottomPanel(control Control, title String) Button {
	cargs := []gdc.ConstTypePtr{control.AsCTypePtr(), title.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewButton()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnAddControlToBottomPanel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorPlugin) AddControlToDock(slot EditorPluginDockSlot, control Control) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot), control.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnAddControlToDock), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) RemoveControlFromDocks(control Control) {
	cargs := []gdc.ConstTypePtr{control.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnRemoveControlFromDocks), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) RemoveControlFromBottomPanel(control Control) {
	cargs := []gdc.ConstTypePtr{control.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnRemoveControlFromBottomPanel), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) RemoveControlFromContainer(container EditorPluginCustomControlContainer, control Control) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&container), control.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnRemoveControlFromContainer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) AddToolMenuItem(name String, callable Callable) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), callable.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnAddToolMenuItem), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) AddToolSubmenuItem(name String, submenu PopupMenu) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), submenu.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnAddToolSubmenuItem), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) RemoveToolMenuItem(name String) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnRemoveToolMenuItem), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) GetExportAsMenu() PopupMenu {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPopupMenu()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnGetExportAsMenu), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorPlugin) AddCustomType(type_ String, base String, script Script, icon Texture2D) {
	cargs := []gdc.ConstTypePtr{type_.AsCTypePtr(), base.AsCTypePtr(), script.AsCTypePtr(), icon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnAddCustomType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) RemoveCustomType(type_ String) {
	cargs := []gdc.ConstTypePtr{type_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnRemoveCustomType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) AddAutoloadSingleton(name String, path String) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnAddAutoloadSingleton), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) RemoveAutoloadSingleton(name String) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnRemoveAutoloadSingleton), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) UpdateOverlays() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnUpdateOverlays), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *EditorPlugin) MakeBottomPanelItemVisible(item Control) {
	cargs := []gdc.ConstTypePtr{item.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnMakeBottomPanelItemVisible), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) HideBottomPanel() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnHideBottomPanel), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) GetUndoRedo() EditorUndoRedoManager {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewEditorUndoRedoManager()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnGetUndoRedo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorPlugin) AddUndoRedoInspectorHookCallback(callable Callable) {
	cargs := []gdc.ConstTypePtr{callable.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnAddUndoRedoInspectorHookCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) RemoveUndoRedoInspectorHookCallback(callable Callable) {
	cargs := []gdc.ConstTypePtr{callable.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnRemoveUndoRedoInspectorHookCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) QueueSaveLayout() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnQueueSaveLayout), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) AddTranslationParserPlugin(parser EditorTranslationParserPlugin) {
	cargs := []gdc.ConstTypePtr{parser.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnAddTranslationParserPlugin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) RemoveTranslationParserPlugin(parser EditorTranslationParserPlugin) {
	cargs := []gdc.ConstTypePtr{parser.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnRemoveTranslationParserPlugin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) AddImportPlugin(importer EditorImportPlugin, first_priority bool) {
	cargs := []gdc.ConstTypePtr{importer.AsCTypePtr(), gdc.ConstTypePtr(&first_priority)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnAddImportPlugin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) RemoveImportPlugin(importer EditorImportPlugin) {
	cargs := []gdc.ConstTypePtr{importer.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnRemoveImportPlugin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) AddSceneFormatImporterPlugin(scene_format_importer EditorSceneFormatImporter, first_priority bool) {
	cargs := []gdc.ConstTypePtr{scene_format_importer.AsCTypePtr(), gdc.ConstTypePtr(&first_priority)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnAddSceneFormatImporterPlugin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) RemoveSceneFormatImporterPlugin(scene_format_importer EditorSceneFormatImporter) {
	cargs := []gdc.ConstTypePtr{scene_format_importer.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnRemoveSceneFormatImporterPlugin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) AddScenePostImportPlugin(scene_import_plugin EditorScenePostImportPlugin, first_priority bool) {
	cargs := []gdc.ConstTypePtr{scene_import_plugin.AsCTypePtr(), gdc.ConstTypePtr(&first_priority)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnAddScenePostImportPlugin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) RemoveScenePostImportPlugin(scene_import_plugin EditorScenePostImportPlugin) {
	cargs := []gdc.ConstTypePtr{scene_import_plugin.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnRemoveScenePostImportPlugin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) AddExportPlugin(plugin EditorExportPlugin) {
	cargs := []gdc.ConstTypePtr{plugin.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnAddExportPlugin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) RemoveExportPlugin(plugin EditorExportPlugin) {
	cargs := []gdc.ConstTypePtr{plugin.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnRemoveExportPlugin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) AddNode3DGizmoPlugin(plugin EditorNode3DGizmoPlugin) {
	cargs := []gdc.ConstTypePtr{plugin.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnAddNode3DGizmoPlugin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) RemoveNode3DGizmoPlugin(plugin EditorNode3DGizmoPlugin) {
	cargs := []gdc.ConstTypePtr{plugin.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnRemoveNode3DGizmoPlugin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) AddInspectorPlugin(plugin EditorInspectorPlugin) {
	cargs := []gdc.ConstTypePtr{plugin.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnAddInspectorPlugin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) RemoveInspectorPlugin(plugin EditorInspectorPlugin) {
	cargs := []gdc.ConstTypePtr{plugin.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnRemoveInspectorPlugin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) AddResourceConversionPlugin(plugin EditorResourceConversionPlugin) {
	cargs := []gdc.ConstTypePtr{plugin.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnAddResourceConversionPlugin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) RemoveResourceConversionPlugin(plugin EditorResourceConversionPlugin) {
	cargs := []gdc.ConstTypePtr{plugin.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnRemoveResourceConversionPlugin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) SetInputEventForwardingAlwaysEnabled() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnSetInputEventForwardingAlwaysEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) SetForceDrawOverForwardingEnabled() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnSetForceDrawOverForwardingEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) GetEditorInterface() EditorInterface {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewEditorInterface()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnGetEditorInterface), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorPlugin) GetScriptCreateDialog() ScriptCreateDialog {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewScriptCreateDialog()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnGetScriptCreateDialog), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorPlugin) AddDebuggerPlugin(script EditorDebuggerPlugin) {
	cargs := []gdc.ConstTypePtr{script.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnAddDebuggerPlugin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) RemoveDebuggerPlugin(script EditorDebuggerPlugin) {
	cargs := []gdc.ConstTypePtr{script.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnRemoveDebuggerPlugin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorPlugin) GetPluginVersion() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorPlugin.fnGetPluginVersion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals

type EditorPluginSceneChangedSignalFn func(scene_root Node)

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

type EditorPluginSceneClosedSignalFn func(filepath String)

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

type EditorPluginMainScreenChangedSignalFn func(screen_name String)

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

type EditorPluginResourceSavedSignalFn func(resource Resource)

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
