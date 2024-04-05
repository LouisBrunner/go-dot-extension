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

type ptrsForEditorInterfaceList struct {
  fnRestartEditor gdc.MethodBindPtr
  fnGetCommandPalette gdc.MethodBindPtr
  fnGetResourceFilesystem gdc.MethodBindPtr
  fnGetEditorPaths gdc.MethodBindPtr
  fnGetResourcePreviewer gdc.MethodBindPtr
  fnGetSelection gdc.MethodBindPtr
  fnGetEditorSettings gdc.MethodBindPtr
  fnMakeMeshPreviews gdc.MethodBindPtr
  fnSetPluginEnabled gdc.MethodBindPtr
  fnIsPluginEnabled gdc.MethodBindPtr
  fnGetEditorTheme gdc.MethodBindPtr
  fnGetBaseControl gdc.MethodBindPtr
  fnGetEditorMainScreen gdc.MethodBindPtr
  fnGetScriptEditor gdc.MethodBindPtr
  fnGetEditorViewport2D gdc.MethodBindPtr
  fnGetEditorViewport3D gdc.MethodBindPtr
  fnSetMainScreenEditor gdc.MethodBindPtr
  fnSetDistractionFreeMode gdc.MethodBindPtr
  fnIsDistractionFreeModeEnabled gdc.MethodBindPtr
  fnGetEditorScale gdc.MethodBindPtr
  fnPopupDialog gdc.MethodBindPtr
  fnPopupDialogCentered gdc.MethodBindPtr
  fnPopupDialogCenteredRatio gdc.MethodBindPtr
  fnPopupDialogCenteredClamped gdc.MethodBindPtr
  fnGetCurrentFeatureProfile gdc.MethodBindPtr
  fnSetCurrentFeatureProfile gdc.MethodBindPtr
  fnGetFileSystemDock gdc.MethodBindPtr
  fnSelectFile gdc.MethodBindPtr
  fnGetSelectedPaths gdc.MethodBindPtr
  fnGetCurrentPath gdc.MethodBindPtr
  fnGetCurrentDirectory gdc.MethodBindPtr
  fnGetInspector gdc.MethodBindPtr
  fnInspectObject gdc.MethodBindPtr
  fnEditResource gdc.MethodBindPtr
  fnEditNode gdc.MethodBindPtr
  fnEditScript gdc.MethodBindPtr
  fnOpenSceneFromPath gdc.MethodBindPtr
  fnReloadSceneFromPath gdc.MethodBindPtr
  fnGetOpenScenes gdc.MethodBindPtr
  fnGetEditedSceneRoot gdc.MethodBindPtr
  fnSaveScene gdc.MethodBindPtr
  fnSaveSceneAs gdc.MethodBindPtr
  fnSaveAllScenes gdc.MethodBindPtr
  fnMarkSceneAsUnsaved gdc.MethodBindPtr
  fnPlayMainScene gdc.MethodBindPtr
  fnPlayCurrentScene gdc.MethodBindPtr
  fnPlayCustomScene gdc.MethodBindPtr
  fnStopPlayingScene gdc.MethodBindPtr
  fnIsPlayingScene gdc.MethodBindPtr
  fnGetPlayingScene gdc.MethodBindPtr
  fnSetMovieMakerEnabled gdc.MethodBindPtr
  fnIsMovieMakerEnabled gdc.MethodBindPtr
}

var ptrsForEditorInterface ptrsForEditorInterfaceList

func initEditorInterfacePtrs(iface gdc.Interface) {

  className := StringNameFromStr("EditorInterface")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("restart_editor")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnRestartEditor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3216645846))
  }
  {
    methodName := StringNameFromStr("get_command_palette")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnGetCommandPalette = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2471163807))
  }
  {
    methodName := StringNameFromStr("get_resource_filesystem")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnGetResourceFilesystem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 780151678))
  }
  {
    methodName := StringNameFromStr("get_editor_paths")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnGetEditorPaths = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1595760068))
  }
  {
    methodName := StringNameFromStr("get_resource_previewer")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnGetResourcePreviewer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 943486957))
  }
  {
    methodName := StringNameFromStr("get_selection")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnGetSelection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2690272531))
  }
  {
    methodName := StringNameFromStr("get_editor_settings")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnGetEditorSettings = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4086932459))
  }
  {
    methodName := StringNameFromStr("make_mesh_previews")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnMakeMeshPreviews = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 878078554))
  }
  {
    methodName := StringNameFromStr("set_plugin_enabled")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnSetPluginEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2678287736))
  }
  {
    methodName := StringNameFromStr("is_plugin_enabled")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnIsPluginEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
  }
  {
    methodName := StringNameFromStr("get_editor_theme")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnGetEditorTheme = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3846893731))
  }
  {
    methodName := StringNameFromStr("get_base_control")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnGetBaseControl = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2783021301))
  }
  {
    methodName := StringNameFromStr("get_editor_main_screen")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnGetEditorMainScreen = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1706218421))
  }
  {
    methodName := StringNameFromStr("get_script_editor")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnGetScriptEditor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 90868003))
  }
  {
    methodName := StringNameFromStr("get_editor_viewport_2d")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnGetEditorViewport2D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3750751911))
  }
  {
    methodName := StringNameFromStr("get_editor_viewport_3d")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnGetEditorViewport3D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1970834490))
  }
  {
    methodName := StringNameFromStr("set_main_screen_editor")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnSetMainScreenEditor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("set_distraction_free_mode")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnSetDistractionFreeMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_distraction_free_mode_enabled")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnIsDistractionFreeModeEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_editor_scale")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnGetEditorScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("popup_dialog")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnPopupDialog = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2015770942))
  }
  {
    methodName := StringNameFromStr("popup_dialog_centered")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnPopupDialogCentered = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 346557367))
  }
  {
    methodName := StringNameFromStr("popup_dialog_centered_ratio")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnPopupDialogCenteredRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2093669136))
  }
  {
    methodName := StringNameFromStr("popup_dialog_centered_clamped")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnPopupDialogCenteredClamped = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3763385571))
  }
  {
    methodName := StringNameFromStr("get_current_feature_profile")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnGetCurrentFeatureProfile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("set_current_feature_profile")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnSetCurrentFeatureProfile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_file_system_dock")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnGetFileSystemDock = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3751012327))
  }
  {
    methodName := StringNameFromStr("select_file")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnSelectFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_selected_paths")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnGetSelectedPaths = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
  }
  {
    methodName := StringNameFromStr("get_current_path")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnGetCurrentPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("get_current_directory")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnGetCurrentDirectory = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("get_inspector")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnGetInspector = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3517113938))
  }
  {
    methodName := StringNameFromStr("inspect_object")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnInspectObject = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 127962172))
  }
  {
    methodName := StringNameFromStr("edit_resource")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnEditResource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 968641751))
  }
  {
    methodName := StringNameFromStr("edit_node")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnEditNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1078189570))
  }
  {
    methodName := StringNameFromStr("edit_script")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnEditScript = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 219829402))
  }
  {
    methodName := StringNameFromStr("open_scene_from_path")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnOpenSceneFromPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("reload_scene_from_path")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnReloadSceneFromPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_open_scenes")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnGetOpenScenes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
  }
  {
    methodName := StringNameFromStr("get_edited_scene_root")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnGetEditedSceneRoot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3160264692))
  }
  {
    methodName := StringNameFromStr("save_scene")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnSaveScene = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166280745))
  }
  {
    methodName := StringNameFromStr("save_scene_as")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnSaveSceneAs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3647332257))
  }
  {
    methodName := StringNameFromStr("save_all_scenes")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnSaveAllScenes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("mark_scene_as_unsaved")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnMarkSceneAsUnsaved = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("play_main_scene")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnPlayMainScene = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("play_current_scene")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnPlayCurrentScene = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("play_custom_scene")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnPlayCustomScene = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("stop_playing_scene")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnStopPlayingScene = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("is_playing_scene")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnIsPlayingScene = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_playing_scene")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnGetPlayingScene = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("set_movie_maker_enabled")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnSetMovieMakerEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_movie_maker_enabled")
    defer methodName.Destroy()
    ptrsForEditorInterface.fnIsMovieMakerEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

type EditorInterface struct {
  Object
}

func (me *EditorInterface) BaseClass() string {
  return "EditorInterface"
}

func NewEditorInterface() *EditorInterface {
  str := StringNameFromStr("EditorInterface") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorInterface{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *EditorInterface) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorInterface) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorInterface) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EditorInterface) RestartEditor(save bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&save) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnRestartEditor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) GetCommandPalette() EditorCommandPalette {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewEditorCommandPalette()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnGetCommandPalette), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetResourceFilesystem() EditorFileSystem {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewEditorFileSystem()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnGetResourceFilesystem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetEditorPaths() EditorPaths {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewEditorPaths()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnGetEditorPaths), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetResourcePreviewer() EditorResourcePreview {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewEditorResourcePreview()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnGetResourcePreviewer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetSelection() EditorSelection {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewEditorSelection()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnGetSelection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetEditorSettings() EditorSettings {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewEditorSettings()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnGetEditorSettings), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) MakeMeshPreviews(meshes []Mesh, preview_size int64, ) []Texture2D {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&meshes) , gdc.ConstTypePtr(&preview_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()
  pinner.Pin(&meshes)
  pinner.Pin(&preview_size)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnMakeMeshPreviews), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Texture2D](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *EditorInterface) SetPluginEnabled(plugin String, enabled bool, )  {
  cargs := []gdc.ConstTypePtr{plugin.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnSetPluginEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) IsPluginEnabled(plugin String, ) bool {
  cargs := []gdc.ConstTypePtr{plugin.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnIsPluginEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorInterface) GetEditorTheme() Theme {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTheme()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnGetEditorTheme), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetBaseControl() Control {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewControl()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnGetBaseControl), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetEditorMainScreen() VBoxContainer {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVBoxContainer()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnGetEditorMainScreen), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetScriptEditor() ScriptEditor {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewScriptEditor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnGetScriptEditor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetEditorViewport2D() SubViewport {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewSubViewport()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnGetEditorViewport2D), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetEditorViewport3D(idx int64, ) SubViewport {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewSubViewport()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnGetEditorViewport3D), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) SetMainScreenEditor(name String, )  {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnSetMainScreenEditor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) SetDistractionFreeMode(enter bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enter) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnSetDistractionFreeMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) IsDistractionFreeModeEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnIsDistractionFreeModeEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorInterface) GetEditorScale() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnGetEditorScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorInterface) PopupDialog(dialog Window, rect Rect2i, )  {
  cargs := []gdc.ConstTypePtr{dialog.AsCTypePtr(), rect.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnPopupDialog), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) PopupDialogCentered(dialog Window, minsize Vector2i, )  {
  cargs := []gdc.ConstTypePtr{dialog.AsCTypePtr(), minsize.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnPopupDialogCentered), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) PopupDialogCenteredRatio(dialog Window, ratio float64, )  {
  cargs := []gdc.ConstTypePtr{dialog.AsCTypePtr(), gdc.ConstTypePtr(&ratio) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnPopupDialogCenteredRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) PopupDialogCenteredClamped(dialog Window, minsize Vector2i, fallback_ratio float64, )  {
  cargs := []gdc.ConstTypePtr{dialog.AsCTypePtr(), minsize.AsCTypePtr(), gdc.ConstTypePtr(&fallback_ratio) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnPopupDialogCenteredClamped), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) GetCurrentFeatureProfile() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnGetCurrentFeatureProfile), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) SetCurrentFeatureProfile(profile_name String, )  {
  cargs := []gdc.ConstTypePtr{profile_name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnSetCurrentFeatureProfile), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) GetFileSystemDock() FileSystemDock {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFileSystemDock()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnGetFileSystemDock), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) SelectFile(file String, )  {
  cargs := []gdc.ConstTypePtr{file.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnSelectFile), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) GetSelectedPaths() PackedStringArray {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnGetSelectedPaths), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetCurrentPath() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnGetCurrentPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetCurrentDirectory() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnGetCurrentDirectory), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetInspector() EditorInspector {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewEditorInspector()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnGetInspector), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) InspectObject(object Object, for_property String, inspector_only bool, )  {
  cargs := []gdc.ConstTypePtr{object.AsCTypePtr(), for_property.AsCTypePtr(), gdc.ConstTypePtr(&inspector_only) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnInspectObject), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) EditResource(resource Resource, )  {
  cargs := []gdc.ConstTypePtr{resource.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnEditResource), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) EditNode(node Node, )  {
  cargs := []gdc.ConstTypePtr{node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnEditNode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) EditScript(script Script, line int64, column int64, grab_focus bool, )  {
  cargs := []gdc.ConstTypePtr{script.AsCTypePtr(), gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&column) , gdc.ConstTypePtr(&grab_focus) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnEditScript), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) OpenSceneFromPath(scene_filepath String, )  {
  cargs := []gdc.ConstTypePtr{scene_filepath.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnOpenSceneFromPath), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) ReloadSceneFromPath(scene_filepath String, )  {
  cargs := []gdc.ConstTypePtr{scene_filepath.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnReloadSceneFromPath), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) GetOpenScenes() PackedStringArray {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnGetOpenScenes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) GetEditedSceneRoot() Node {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNode()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnGetEditedSceneRoot), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) SaveScene() Error {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnSaveScene), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *EditorInterface) SaveSceneAs(path String, with_preview bool, )  {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&with_preview) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnSaveSceneAs), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) SaveAllScenes()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnSaveAllScenes), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) MarkSceneAsUnsaved()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnMarkSceneAsUnsaved), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) PlayMainScene()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnPlayMainScene), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) PlayCurrentScene()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnPlayCurrentScene), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) PlayCustomScene(scene_filepath String, )  {
  cargs := []gdc.ConstTypePtr{scene_filepath.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnPlayCustomScene), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) StopPlayingScene()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnStopPlayingScene), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) IsPlayingScene() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnIsPlayingScene), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorInterface) GetPlayingScene() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnGetPlayingScene), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInterface) SetMovieMakerEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnSetMovieMakerEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorInterface) IsMovieMakerEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorInterface.fnIsMovieMakerEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
