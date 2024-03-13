// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type DirAccess struct {
  obj gdc.ObjectPtr
}

func (me *DirAccess) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *DirAccess) BaseClass() string {
  return "DirAccess"
}



// Enums

func (me *DirAccess) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *DirAccess) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *DirAccess) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  DirAccessOpen(path String, ) DirAccess {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("open")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1923528528) // FIXME: should cache?
  var ret DirAccess
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  DirAccessGetOpenError() Error {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_open_error")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166280745) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DirAccess) ListDirBegin() Error {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("list_dir_begin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2610976713) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DirAccess) GetNext() String {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_next")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2841200299) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DirAccess) CurrentIsDir() bool {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("current_is_dir")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DirAccess) ListDirEnd()  {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("list_dir_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DirAccess) GetFiles() PackedStringArray {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_files")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2981934095) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  DirAccessGetFilesAt(path String, ) PackedStringArray {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_files_at")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3538744774) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DirAccess) GetDirectories() PackedStringArray {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_directories")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2981934095) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  DirAccessGetDirectoriesAt(path String, ) PackedStringArray {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_directories_at")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3538744774) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  DirAccessGetDriveCount() int {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_drive_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  DirAccessGetDriveName(idx int, ) String {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_drive_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 990163283) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DirAccess) GetCurrentDrive() int {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_drive")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DirAccess) ChangeDir(to_dir String, ) Error {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("change_dir")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(to_dir.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DirAccess) GetCurrentDir(include_drive bool, ) String {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_dir")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1287308131) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&include_drive), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DirAccess) MakeDir(path String, ) Error {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_dir")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  DirAccessMakeDirAbsolute(path String, ) Error {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_dir_absolute")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DirAccess) MakeDirRecursive(path String, ) Error {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_dir_recursive")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  DirAccessMakeDirRecursiveAbsolute(path String, ) Error {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_dir_recursive_absolute")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DirAccess) FileExists(path String, ) bool {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("file_exists")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2323990056) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DirAccess) DirExists(path String, ) bool {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("dir_exists")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2323990056) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  DirAccessDirExistsAbsolute(path String, ) bool {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("dir_exists_absolute")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2323990056) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DirAccess) GetSpaceLeft() int {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_space_left")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DirAccess) Copy(from String, to String, chmod_flags int, ) Error {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("copy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1063198817) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from.AsCTypePtr()), gdc.ConstTypePtr(to.AsCTypePtr()), gdc.ConstTypePtr(&chmod_flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  DirAccessCopyAbsolute(from String, to String, chmod_flags int, ) Error {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("copy_absolute")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1063198817) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from.AsCTypePtr()), gdc.ConstTypePtr(to.AsCTypePtr()), gdc.ConstTypePtr(&chmod_flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DirAccess) Rename(from String, to String, ) Error {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rename")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 852856452) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from.AsCTypePtr()), gdc.ConstTypePtr(to.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  DirAccessRenameAbsolute(from String, to String, ) Error {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rename_absolute")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 852856452) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from.AsCTypePtr()), gdc.ConstTypePtr(to.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DirAccess) Remove(path String, ) Error {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  DirAccessRemoveAbsolute(path String, ) Error {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_absolute")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DirAccess) SetIncludeNavigational(enable bool, )  {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_include_navigational")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DirAccess) GetIncludeNavigational() bool {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_include_navigational")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DirAccess) SetIncludeHidden(enable bool, )  {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_include_hidden")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DirAccess) GetIncludeHidden() bool {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_include_hidden")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DirAccess) IsCaseSensitive(path String, ) bool {
  classNameV := StringNameFromStr("DirAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_case_sensitive")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
