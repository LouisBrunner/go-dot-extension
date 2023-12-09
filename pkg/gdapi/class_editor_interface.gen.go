// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorInterface struct {
  obj gdc.ObjectPtr
}

func (me *EditorInterface) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorInterface) BaseClass() string {
  return "EditorInterface"
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
  panic("TODO: implement")
}

func  (me *EditorInterface) GetCommandPalette()  {
  panic("TODO: implement")
}

func  (me *EditorInterface) GetResourceFilesystem()  {
  panic("TODO: implement")
}

func  (me *EditorInterface) GetEditorPaths()  {
  panic("TODO: implement")
}

func  (me *EditorInterface) GetResourcePreviewer()  {
  panic("TODO: implement")
}

func  (me *EditorInterface) GetSelection()  {
  panic("TODO: implement")
}

func  (me *EditorInterface) GetEditorSettings()  {
  panic("TODO: implement")
}

func  (me *EditorInterface) MakeMeshPreviews(meshes Mesh, preview_size int, )  {
  panic("TODO: implement")
}

func  (me *EditorInterface) SetPluginEnabled(plugin String, enabled bool, )  {
  panic("TODO: implement")
}

func  (me *EditorInterface) IsPluginEnabled(plugin String, )  {
  panic("TODO: implement")
}

func  (me *EditorInterface) GetBaseControl()  {
  panic("TODO: implement")
}

func  (me *EditorInterface) GetEditorMainScreen()  {
  panic("TODO: implement")
}

func  (me *EditorInterface) GetScriptEditor()  {
  panic("TODO: implement")
}

func  (me *EditorInterface) SetMainScreenEditor(name String, )  {
  panic("TODO: implement")
}

func  (me *EditorInterface) SetDistractionFreeMode(enter bool, )  {
  panic("TODO: implement")
}

func  (me *EditorInterface) IsDistractionFreeModeEnabled()  {
  panic("TODO: implement")
}

func  (me *EditorInterface) GetEditorScale()  {
  panic("TODO: implement")
}

func  (me *EditorInterface) PopupDialog(dialog Window, rect Rect2i, )  {
  panic("TODO: implement")
}

func  (me *EditorInterface) PopupDialogCentered(dialog Window, minsize Vector2i, )  {
  panic("TODO: implement")
}

func  (me *EditorInterface) PopupDialogCenteredRatio(dialog Window, ratio float32, )  {
  panic("TODO: implement")
}

func  (me *EditorInterface) PopupDialogCenteredClamped(dialog Window, minsize Vector2i, fallback_ratio float32, )  {
  panic("TODO: implement")
}

func  (me *EditorInterface) GetFileSystemDock()  {
  panic("TODO: implement")
}

func  (me *EditorInterface) SelectFile(file String, )  {
  panic("TODO: implement")
}

func  (me *EditorInterface) GetSelectedPaths()  {
  panic("TODO: implement")
}

func  (me *EditorInterface) GetCurrentPath()  {
  panic("TODO: implement")
}

func  (me *EditorInterface) GetCurrentDirectory()  {
  panic("TODO: implement")
}

func  (me *EditorInterface) GetInspector()  {
  panic("TODO: implement")
}

func  (me *EditorInterface) InspectObject(object Object, for_property String, inspector_only bool, )  {
  panic("TODO: implement")
}

func  (me *EditorInterface) EditResource(resource Resource, )  {
  panic("TODO: implement")
}

func  (me *EditorInterface) EditNode(node Node, )  {
  panic("TODO: implement")
}

func  (me *EditorInterface) EditScript(script Script, line int, column int, grab_focus bool, )  {
  panic("TODO: implement")
}

func  (me *EditorInterface) OpenSceneFromPath(scene_filepath String, )  {
  panic("TODO: implement")
}

func  (me *EditorInterface) ReloadSceneFromPath(scene_filepath String, )  {
  panic("TODO: implement")
}

func  (me *EditorInterface) GetOpenScenes()  {
  panic("TODO: implement")
}

func  (me *EditorInterface) GetEditedSceneRoot()  {
  panic("TODO: implement")
}

func  (me *EditorInterface) SaveScene()  {
  panic("TODO: implement")
}

func  (me *EditorInterface) SaveSceneAs(path String, with_preview bool, )  {
  panic("TODO: implement")
}

func  (me *EditorInterface) MarkSceneAsUnsaved()  {
  panic("TODO: implement")
}

func  (me *EditorInterface) PlayMainScene()  {
  panic("TODO: implement")
}

func  (me *EditorInterface) PlayCurrentScene()  {
  panic("TODO: implement")
}

func  (me *EditorInterface) PlayCustomScene(scene_filepath String, )  {
  panic("TODO: implement")
}

func  (me *EditorInterface) StopPlayingScene()  {
  panic("TODO: implement")
}

func  (me *EditorInterface) IsPlayingScene()  {
  panic("TODO: implement")
}

func  (me *EditorInterface) GetPlayingScene()  {
  panic("TODO: implement")
}

func  (me *EditorInterface) SetMovieMakerEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *EditorInterface) IsMovieMakerEnabled()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
