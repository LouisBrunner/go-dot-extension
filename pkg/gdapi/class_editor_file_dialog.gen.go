// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorFileDialog struct {
  obj gdc.ObjectPtr
}

func (me *EditorFileDialog) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorFileDialog) BaseClass() string {
  return "EditorFileDialog"
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
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorFileDialog) AddFilter(filter String, description String, )  {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 233059325) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(filter.AsCTypePtr()), gdc.ConstTypePtr(description.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorFileDialog) SetFilters(filters PackedStringArray, )  {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_filters")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4015028928) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(filters.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorFileDialog) GetFilters() PackedStringArray {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_filters")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileDialog) GetCurrentDir() String {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_dir")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileDialog) GetCurrentFile() String {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileDialog) GetCurrentPath() String {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileDialog) SetCurrentDir(dir String, )  {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_current_dir")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(dir.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorFileDialog) SetCurrentFile(file String, )  {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_current_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(file.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorFileDialog) SetCurrentPath(path String, )  {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_current_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorFileDialog) SetFileMode(mode EditorFileDialogFileMode, )  {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_file_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 274150415) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorFileDialog) GetFileMode() EditorFileDialogFileMode {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_file_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2681044145) // FIXME: should cache?
  var ret EditorFileDialogFileMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileDialog) GetVbox() VBoxContainer {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vbox")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 915758477) // FIXME: should cache?
  var ret VBoxContainer
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileDialog) GetLineEdit() LineEdit {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_edit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4071694264) // FIXME: should cache?
  var ret LineEdit
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileDialog) SetAccess(access EditorFileDialogAccess, )  {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_access")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3882893764) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&access), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorFileDialog) GetAccess() EditorFileDialogAccess {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_access")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 778734016) // FIXME: should cache?
  var ret EditorFileDialogAccess
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileDialog) SetShowHiddenFiles(show bool, )  {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_show_hidden_files")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&show), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorFileDialog) IsShowingHiddenFiles() bool {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_showing_hidden_files")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileDialog) SetDisplayMode(mode EditorFileDialogDisplayMode, )  {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_display_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3049004050) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorFileDialog) GetDisplayMode() EditorFileDialogDisplayMode {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_display_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3517174669) // FIXME: should cache?
  var ret EditorFileDialogDisplayMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileDialog) SetDisableOverwriteWarning(disable bool, )  {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_disable_overwrite_warning")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&disable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorFileDialog) IsOverwriteWarningDisabled() bool {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_overwrite_warning_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileDialog) Invalidate()  {
  classNameV := StringNameFromStr("EditorFileDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("invalidate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties

func (me *EditorFileDialog) GetPropAccess() int {
  panic("TODO: implement")
}

func (me *EditorFileDialog) SetPropAccess(value int) {
  panic("TODO: implement")
}

func (me *EditorFileDialog) GetPropDisplayMode() int {
  panic("TODO: implement")
}

func (me *EditorFileDialog) SetPropDisplayMode(value int) {
  panic("TODO: implement")
}

func (me *EditorFileDialog) GetPropFileMode() int {
  panic("TODO: implement")
}

func (me *EditorFileDialog) SetPropFileMode(value int) {
  panic("TODO: implement")
}

func (me *EditorFileDialog) GetPropCurrentDir() String {
  panic("TODO: implement")
}

func (me *EditorFileDialog) SetPropCurrentDir(value String) {
  panic("TODO: implement")
}

func (me *EditorFileDialog) GetPropCurrentFile() String {
  panic("TODO: implement")
}

func (me *EditorFileDialog) SetPropCurrentFile(value String) {
  panic("TODO: implement")
}

func (me *EditorFileDialog) GetPropCurrentPath() String {
  panic("TODO: implement")
}

func (me *EditorFileDialog) SetPropCurrentPath(value String) {
  panic("TODO: implement")
}

func (me *EditorFileDialog) GetPropFilters() PackedStringArray {
  panic("TODO: implement")
}

func (me *EditorFileDialog) SetPropFilters(value PackedStringArray) {
  panic("TODO: implement")
}

func (me *EditorFileDialog) GetPropShowHiddenFiles() bool {
  panic("TODO: implement")
}

func (me *EditorFileDialog) SetPropShowHiddenFiles(value bool) {
  panic("TODO: implement")
}

func (me *EditorFileDialog) GetPropDisableOverwriteWarning() bool {
  panic("TODO: implement")
}

func (me *EditorFileDialog) SetPropDisableOverwriteWarning(value bool) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API