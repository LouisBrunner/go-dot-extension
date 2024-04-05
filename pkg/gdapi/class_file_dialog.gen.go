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

type ptrsForFileDialogList struct {
  fnClearFilters gdc.MethodBindPtr
  fnAddFilter gdc.MethodBindPtr
  fnSetFilters gdc.MethodBindPtr
  fnGetFilters gdc.MethodBindPtr
  fnGetCurrentDir gdc.MethodBindPtr
  fnGetCurrentFile gdc.MethodBindPtr
  fnGetCurrentPath gdc.MethodBindPtr
  fnSetCurrentDir gdc.MethodBindPtr
  fnSetCurrentFile gdc.MethodBindPtr
  fnSetCurrentPath gdc.MethodBindPtr
  fnSetModeOverridesTitle gdc.MethodBindPtr
  fnIsModeOverridingTitle gdc.MethodBindPtr
  fnSetFileMode gdc.MethodBindPtr
  fnGetFileMode gdc.MethodBindPtr
  fnGetVbox gdc.MethodBindPtr
  fnGetLineEdit gdc.MethodBindPtr
  fnSetAccess gdc.MethodBindPtr
  fnGetAccess gdc.MethodBindPtr
  fnSetRootSubfolder gdc.MethodBindPtr
  fnGetRootSubfolder gdc.MethodBindPtr
  fnSetShowHiddenFiles gdc.MethodBindPtr
  fnIsShowingHiddenFiles gdc.MethodBindPtr
  fnSetUseNativeDialog gdc.MethodBindPtr
  fnGetUseNativeDialog gdc.MethodBindPtr
  fnDeselectAll gdc.MethodBindPtr
  fnInvalidate gdc.MethodBindPtr
}

var ptrsForFileDialog ptrsForFileDialogList

func initFileDialogPtrs(iface gdc.Interface) {

  className := StringNameFromStr("FileDialog")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("clear_filters")
    defer methodName.Destroy()
    ptrsForFileDialog.fnClearFilters = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("add_filter")
    defer methodName.Destroy()
    ptrsForFileDialog.fnAddFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3388804757))
  }
  {
    methodName := StringNameFromStr("set_filters")
    defer methodName.Destroy()
    ptrsForFileDialog.fnSetFilters = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4015028928))
  }
  {
    methodName := StringNameFromStr("get_filters")
    defer methodName.Destroy()
    ptrsForFileDialog.fnGetFilters = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
  }
  {
    methodName := StringNameFromStr("get_current_dir")
    defer methodName.Destroy()
    ptrsForFileDialog.fnGetCurrentDir = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("get_current_file")
    defer methodName.Destroy()
    ptrsForFileDialog.fnGetCurrentFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("get_current_path")
    defer methodName.Destroy()
    ptrsForFileDialog.fnGetCurrentPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("set_current_dir")
    defer methodName.Destroy()
    ptrsForFileDialog.fnSetCurrentDir = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("set_current_file")
    defer methodName.Destroy()
    ptrsForFileDialog.fnSetCurrentFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("set_current_path")
    defer methodName.Destroy()
    ptrsForFileDialog.fnSetCurrentPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("set_mode_overrides_title")
    defer methodName.Destroy()
    ptrsForFileDialog.fnSetModeOverridesTitle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_mode_overriding_title")
    defer methodName.Destroy()
    ptrsForFileDialog.fnIsModeOverridingTitle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_file_mode")
    defer methodName.Destroy()
    ptrsForFileDialog.fnSetFileMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3654936397))
  }
  {
    methodName := StringNameFromStr("get_file_mode")
    defer methodName.Destroy()
    ptrsForFileDialog.fnGetFileMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4074825319))
  }
  {
    methodName := StringNameFromStr("get_vbox")
    defer methodName.Destroy()
    ptrsForFileDialog.fnGetVbox = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 915758477))
  }
  {
    methodName := StringNameFromStr("get_line_edit")
    defer methodName.Destroy()
    ptrsForFileDialog.fnGetLineEdit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4071694264))
  }
  {
    methodName := StringNameFromStr("set_access")
    defer methodName.Destroy()
    ptrsForFileDialog.fnSetAccess = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4104413466))
  }
  {
    methodName := StringNameFromStr("get_access")
    defer methodName.Destroy()
    ptrsForFileDialog.fnGetAccess = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3344081076))
  }
  {
    methodName := StringNameFromStr("set_root_subfolder")
    defer methodName.Destroy()
    ptrsForFileDialog.fnSetRootSubfolder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_root_subfolder")
    defer methodName.Destroy()
    ptrsForFileDialog.fnGetRootSubfolder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("set_show_hidden_files")
    defer methodName.Destroy()
    ptrsForFileDialog.fnSetShowHiddenFiles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_showing_hidden_files")
    defer methodName.Destroy()
    ptrsForFileDialog.fnIsShowingHiddenFiles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_use_native_dialog")
    defer methodName.Destroy()
    ptrsForFileDialog.fnSetUseNativeDialog = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_use_native_dialog")
    defer methodName.Destroy()
    ptrsForFileDialog.fnGetUseNativeDialog = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("deselect_all")
    defer methodName.Destroy()
    ptrsForFileDialog.fnDeselectAll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("invalidate")
    defer methodName.Destroy()
    ptrsForFileDialog.fnInvalidate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
}

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
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileDialog.fnClearFilters), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileDialog) AddFilter(filter String, description String, )  {
  cargs := []gdc.ConstTypePtr{filter.AsCTypePtr(), description.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileDialog.fnAddFilter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileDialog) SetFilters(filters PackedStringArray, )  {
  cargs := []gdc.ConstTypePtr{filters.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileDialog.fnSetFilters), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileDialog) GetFilters() PackedStringArray {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileDialog.fnGetFilters), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FileDialog) GetCurrentDir() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileDialog.fnGetCurrentDir), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FileDialog) GetCurrentFile() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileDialog.fnGetCurrentFile), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FileDialog) GetCurrentPath() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileDialog.fnGetCurrentPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FileDialog) SetCurrentDir(dir String, )  {
  cargs := []gdc.ConstTypePtr{dir.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileDialog.fnSetCurrentDir), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileDialog) SetCurrentFile(file String, )  {
  cargs := []gdc.ConstTypePtr{file.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileDialog.fnSetCurrentFile), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileDialog) SetCurrentPath(path String, )  {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileDialog.fnSetCurrentPath), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileDialog) SetModeOverridesTitle(override bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&override) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileDialog.fnSetModeOverridesTitle), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileDialog) IsModeOverridingTitle() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileDialog.fnIsModeOverridingTitle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileDialog) SetFileMode(mode FileDialogFileMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileDialog.fnSetFileMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileDialog) GetFileMode() FileDialogFileMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret FileDialogFileMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileDialog.fnGetFileMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *FileDialog) GetVbox() VBoxContainer {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVBoxContainer()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileDialog.fnGetVbox), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FileDialog) GetLineEdit() LineEdit {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewLineEdit()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileDialog.fnGetLineEdit), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FileDialog) SetAccess(access FileDialogAccess, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&access) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileDialog.fnSetAccess), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileDialog) GetAccess() FileDialogAccess {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret FileDialogAccess

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileDialog.fnGetAccess), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *FileDialog) SetRootSubfolder(dir String, )  {
  cargs := []gdc.ConstTypePtr{dir.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileDialog.fnSetRootSubfolder), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileDialog) GetRootSubfolder() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileDialog.fnGetRootSubfolder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FileDialog) SetShowHiddenFiles(show bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&show) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileDialog.fnSetShowHiddenFiles), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileDialog) IsShowingHiddenFiles() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileDialog.fnIsShowingHiddenFiles), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileDialog) SetUseNativeDialog(native bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&native) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileDialog.fnSetUseNativeDialog), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileDialog) GetUseNativeDialog() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileDialog.fnGetUseNativeDialog), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileDialog) DeselectAll()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileDialog.fnDeselectAll), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileDialog) Invalidate()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileDialog.fnInvalidate), me.obj, unsafe.SliceData(cargs), nil)

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
