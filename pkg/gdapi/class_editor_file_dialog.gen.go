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

type EditorFileDialog struct {
  ConfirmationDialog
}

func (me *EditorFileDialog) BaseClass() string {
  return "EditorFileDialog"
}

func NewEditorFileDialog() *EditorFileDialog {
  str := StringNameFromStr("EditorFileDialog") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorFileDialog{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type EditorFileDialogFileMode int
const (
  EditorFileDialogFileModeFileModeOpenFile EditorFileDialogFileMode = 0
  EditorFileDialogFileModeFileModeOpenFiles EditorFileDialogFileMode = 1
  EditorFileDialogFileModeFileModeOpenDir EditorFileDialogFileMode = 2
  EditorFileDialogFileModeFileModeOpenAny EditorFileDialogFileMode = 3
  EditorFileDialogFileModeFileModeSaveFile EditorFileDialogFileMode = 4
)

type EditorFileDialogAccess int
const (
  EditorFileDialogAccessAccessResources EditorFileDialogAccess = 0
  EditorFileDialogAccessAccessUserdata EditorFileDialogAccess = 1
  EditorFileDialogAccessAccessFilesystem EditorFileDialogAccess = 2
)

type EditorFileDialogDisplayMode int
const (
  EditorFileDialogDisplayModeDisplayThumbnails EditorFileDialogDisplayMode = 0
  EditorFileDialogDisplayModeDisplayList EditorFileDialogDisplayMode = 1
)

func (me *EditorFileDialog) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorFileDialog) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorFileDialog) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EditorFileDialog) ClearFilters()  {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_filters")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorFileDialog) AddFilter(filter String, description String, )  {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3388804757) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{filter.AsCTypePtr(), description.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorFileDialog) SetFilters(filters PackedStringArray, )  {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_filters")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4015028928) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{filters.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorFileDialog) GetFilters() PackedStringArray {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_filters")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorFileDialog) GetCurrentDir() String {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_dir")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorFileDialog) GetCurrentFile() String {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorFileDialog) GetCurrentPath() String {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorFileDialog) SetCurrentDir(dir String, )  {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_current_dir")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{dir.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorFileDialog) SetCurrentFile(file String, )  {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_current_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{file.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorFileDialog) SetCurrentPath(path String, )  {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_current_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorFileDialog) SetFileMode(mode EditorFileDialogFileMode, )  {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_file_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 274150415) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorFileDialog) GetFileMode() EditorFileDialogFileMode {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_file_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2681044145) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret EditorFileDialogFileMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *EditorFileDialog) GetVbox() VBoxContainer {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vbox")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 915758477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVBoxContainer()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorFileDialog) GetLineEdit() LineEdit {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_edit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4071694264) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewLineEdit()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorFileDialog) SetAccess(access EditorFileDialogAccess, )  {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_access")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3882893764) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&access) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorFileDialog) GetAccess() EditorFileDialogAccess {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_access")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 778734016) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret EditorFileDialogAccess

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *EditorFileDialog) SetShowHiddenFiles(show bool, )  {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_show_hidden_files")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&show) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorFileDialog) IsShowingHiddenFiles() bool {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_showing_hidden_files")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorFileDialog) SetDisplayMode(mode EditorFileDialogDisplayMode, )  {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_display_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3049004050) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorFileDialog) GetDisplayMode() EditorFileDialogDisplayMode {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_display_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3517174669) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret EditorFileDialogDisplayMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *EditorFileDialog) SetDisableOverwriteWarning(disable bool, )  {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_disable_overwrite_warning")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&disable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorFileDialog) IsOverwriteWarningDisabled() bool {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_overwrite_warning_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorFileDialog) AddSideMenu(menu Control, title String, )  {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_side_menu")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 402368861) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{menu.AsCTypePtr(), title.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorFileDialog) Invalidate()  {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("invalidate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type EditorFileDialogFileSelectedSignalFn func(path String, )

func (me *EditorFileDialog) ConnectFileSelected(subs SignalSubscribers, fn EditorFileDialogFileSelectedSignalFn) {
  sig := StringNameFromStr("file_selected")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorFileDialog) DisconnectFileSelected(subs SignalSubscribers, fn EditorFileDialogFileSelectedSignalFn) {
  sig := StringNameFromStr("file_selected")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorFileDialogFilesSelectedSignalFn func(paths PackedStringArray, )

func (me *EditorFileDialog) ConnectFilesSelected(subs SignalSubscribers, fn EditorFileDialogFilesSelectedSignalFn) {
  sig := StringNameFromStr("files_selected")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorFileDialog) DisconnectFilesSelected(subs SignalSubscribers, fn EditorFileDialogFilesSelectedSignalFn) {
  sig := StringNameFromStr("files_selected")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorFileDialogDirSelectedSignalFn func(dir String, )

func (me *EditorFileDialog) ConnectDirSelected(subs SignalSubscribers, fn EditorFileDialogDirSelectedSignalFn) {
  sig := StringNameFromStr("dir_selected")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorFileDialog) DisconnectDirSelected(subs SignalSubscribers, fn EditorFileDialogDirSelectedSignalFn) {
  sig := StringNameFromStr("dir_selected")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
