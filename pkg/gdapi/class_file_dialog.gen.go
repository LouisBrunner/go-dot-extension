// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type FileDialog struct {
  obj gdc.ObjectPtr
}

func (me *FileDialog) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *FileDialog) BaseClass() string {
  return "FileDialog"
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
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FileDialog) AddFilter(filter String, description String, )  {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 233059325) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(filter.AsCTypePtr()), gdc.ConstTypePtr(description.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FileDialog) SetFilters(filters PackedStringArray, )  {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_filters")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4015028928) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(filters.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FileDialog) GetFilters() PackedStringArray {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_filters")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FileDialog) GetCurrentDir() String {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_dir")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FileDialog) GetCurrentFile() String {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FileDialog) GetCurrentPath() String {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FileDialog) SetCurrentDir(dir String, )  {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_current_dir")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(dir.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FileDialog) SetCurrentFile(file String, )  {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_current_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(file.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FileDialog) SetCurrentPath(path String, )  {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_current_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FileDialog) SetModeOverridesTitle(override bool, )  {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mode_overrides_title")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&override), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FileDialog) IsModeOverridingTitle() bool {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_mode_overriding_title")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FileDialog) SetFileMode(mode FileDialogFileMode, )  {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_file_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3654936397) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FileDialog) GetFileMode() FileDialogFileMode {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_file_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4074825319) // FIXME: should cache?
  var ret FileDialogFileMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FileDialog) GetVbox() VBoxContainer {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vbox")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 915758477) // FIXME: should cache?
  var ret VBoxContainer
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FileDialog) GetLineEdit() LineEdit {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_edit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4071694264) // FIXME: should cache?
  var ret LineEdit
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FileDialog) SetAccess(access FileDialogAccess, )  {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_access")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4104413466) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&access), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FileDialog) GetAccess() FileDialogAccess {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_access")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3344081076) // FIXME: should cache?
  var ret FileDialogAccess
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FileDialog) SetRootSubfolder(dir String, )  {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_root_subfolder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(dir.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FileDialog) GetRootSubfolder() String {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root_subfolder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FileDialog) SetShowHiddenFiles(show bool, )  {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_show_hidden_files")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&show), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FileDialog) IsShowingHiddenFiles() bool {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_showing_hidden_files")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *FileDialog) DeselectAll()  {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("deselect_all")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FileDialog) Invalidate()  {
  classNameV := StringNameFromStr("FileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("invalidate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties

func (me *FileDialog) GetPropModeOverridesTitle() bool {
  panic("TODO: implement")
}

func (me *FileDialog) SetPropModeOverridesTitle(value bool) {
  panic("TODO: implement")
}

func (me *FileDialog) GetPropFileMode() int {
  panic("TODO: implement")
}

func (me *FileDialog) SetPropFileMode(value int) {
  panic("TODO: implement")
}

func (me *FileDialog) GetPropAccess() int {
  panic("TODO: implement")
}

func (me *FileDialog) SetPropAccess(value int) {
  panic("TODO: implement")
}

func (me *FileDialog) GetPropRootSubfolder() String {
  panic("TODO: implement")
}

func (me *FileDialog) SetPropRootSubfolder(value String) {
  panic("TODO: implement")
}

func (me *FileDialog) GetPropFilters() PackedStringArray {
  panic("TODO: implement")
}

func (me *FileDialog) SetPropFilters(value PackedStringArray) {
  panic("TODO: implement")
}

func (me *FileDialog) GetPropShowHiddenFiles() bool {
  panic("TODO: implement")
}

func (me *FileDialog) SetPropShowHiddenFiles(value bool) {
  panic("TODO: implement")
}

func (me *FileDialog) GetPropCurrentDir() String {
  panic("TODO: implement")
}

func (me *FileDialog) SetPropCurrentDir(value String) {
  panic("TODO: implement")
}

func (me *FileDialog) GetPropCurrentFile() String {
  panic("TODO: implement")
}

func (me *FileDialog) SetPropCurrentFile(value String) {
  panic("TODO: implement")
}

func (me *FileDialog) GetPropCurrentPath() String {
  panic("TODO: implement")
}

func (me *FileDialog) SetPropCurrentPath(value String) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API