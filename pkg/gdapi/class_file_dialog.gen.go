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

type FileDialog struct {
  ConfirmationDialog
}

func (me *FileDialog) BaseClass() string {
  return "FileDialog"
}

func NewFileDialog() *FileDialog {
  str := StringNameFromStr("FileDialog") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &FileDialog{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type FileDialogFileMode int
const (
  FileDialogFileModeFileModeOpenFile FileDialogFileMode = 0
  FileDialogFileModeFileModeOpenFiles FileDialogFileMode = 1
  FileDialogFileModeFileModeOpenDir FileDialogFileMode = 2
  FileDialogFileModeFileModeOpenAny FileDialogFileMode = 3
  FileDialogFileModeFileModeSaveFile FileDialogFileMode = 4
)

type FileDialogAccess int
const (
  FileDialogAccessAccessResources FileDialogAccess = 0
  FileDialogAccessAccessUserdata FileDialogAccess = 1
  FileDialogAccessAccessFilesystem FileDialogAccess = 2
)

func (me *FileDialog) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *FileDialog) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *FileDialog) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *FileDialog) ClearFilters()  {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_filters")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileDialog) AddFilter(filter String, description String, )  {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3388804757) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{filter.AsCTypePtr(), description.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileDialog) SetFilters(filters PackedStringArray, )  {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_filters")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4015028928) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{filters.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileDialog) GetFilters() PackedStringArray {
  classNameV := StringNameFromStr("FileDialog")
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

func  (me *FileDialog) GetCurrentDir() String {
  classNameV := StringNameFromStr("FileDialog")
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

func  (me *FileDialog) GetCurrentFile() String {
  classNameV := StringNameFromStr("FileDialog")
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

func  (me *FileDialog) GetCurrentPath() String {
  classNameV := StringNameFromStr("FileDialog")
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

func  (me *FileDialog) SetCurrentDir(dir String, )  {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_current_dir")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{dir.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileDialog) SetCurrentFile(file String, )  {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_current_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{file.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileDialog) SetCurrentPath(path String, )  {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_current_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileDialog) SetModeOverridesTitle(override bool, )  {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mode_overrides_title")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&override) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileDialog) IsModeOverridingTitle() bool {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_mode_overriding_title")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileDialog) SetFileMode(mode FileDialogFileMode, )  {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_file_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3654936397) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileDialog) GetFileMode() FileDialogFileMode {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_file_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4074825319) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret FileDialogFileMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *FileDialog) GetVbox() VBoxContainer {
  classNameV := StringNameFromStr("FileDialog")
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

func  (me *FileDialog) GetLineEdit() LineEdit {
  classNameV := StringNameFromStr("FileDialog")
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

func  (me *FileDialog) SetAccess(access FileDialogAccess, )  {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_access")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4104413466) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&access) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileDialog) GetAccess() FileDialogAccess {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_access")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3344081076) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret FileDialogAccess

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *FileDialog) SetRootSubfolder(dir String, )  {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_root_subfolder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{dir.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileDialog) GetRootSubfolder() String {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root_subfolder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FileDialog) SetShowHiddenFiles(show bool, )  {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_show_hidden_files")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&show) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileDialog) IsShowingHiddenFiles() bool {
  classNameV := StringNameFromStr("FileDialog")
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

func  (me *FileDialog) SetUseNativeDialog(native bool, )  {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_native_dialog")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&native) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileDialog) GetUseNativeDialog() bool {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_use_native_dialog")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileDialog) DeselectAll()  {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("deselect_all")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileDialog) Invalidate()  {
  classNameV := StringNameFromStr("FileDialog")
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

type FileDialogFileSelectedSignalFn func(path String, )

func (me *FileDialog) ConnectFileSelected(subs SignalSubscribers, fn FileDialogFileSelectedSignalFn) {
  sig := StringNameFromStr("file_selected")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *FileDialog) DisconnectFileSelected(subs SignalSubscribers, fn FileDialogFileSelectedSignalFn) {
  sig := StringNameFromStr("file_selected")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type FileDialogFilesSelectedSignalFn func(paths PackedStringArray, )

func (me *FileDialog) ConnectFilesSelected(subs SignalSubscribers, fn FileDialogFilesSelectedSignalFn) {
  sig := StringNameFromStr("files_selected")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *FileDialog) DisconnectFilesSelected(subs SignalSubscribers, fn FileDialogFilesSelectedSignalFn) {
  sig := StringNameFromStr("files_selected")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type FileDialogDirSelectedSignalFn func(dir String, )

func (me *FileDialog) ConnectDirSelected(subs SignalSubscribers, fn FileDialogDirSelectedSignalFn) {
  sig := StringNameFromStr("dir_selected")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *FileDialog) DisconnectDirSelected(subs SignalSubscribers, fn FileDialogDirSelectedSignalFn) {
  sig := StringNameFromStr("dir_selected")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
