// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorFileSystemDirectory struct {
  obj gdc.ObjectPtr
}

func (me *EditorFileSystemDirectory) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorFileSystemDirectory) BaseClass() string {
  return "EditorFileSystemDirectory"
}



// Enums

func (me *EditorFileSystemDirectory) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorFileSystemDirectory) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorFileSystemDirectory) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EditorFileSystemDirectory) GetSubdirCount() int {
  classNameV := StringNameFromStr("EditorFileSystemDirectory")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_subdir_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileSystemDirectory) GetSubdir(idx int, ) EditorFileSystemDirectory {
  classNameV := StringNameFromStr("EditorFileSystemDirectory")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_subdir")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2330964164) // FIXME: should cache?
  var ret EditorFileSystemDirectory
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileSystemDirectory) GetFileCount() int {
  classNameV := StringNameFromStr("EditorFileSystemDirectory")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_file_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileSystemDirectory) GetFile(idx int, ) String {
  classNameV := StringNameFromStr("EditorFileSystemDirectory")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileSystemDirectory) GetFilePath(idx int, ) String {
  classNameV := StringNameFromStr("EditorFileSystemDirectory")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_file_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileSystemDirectory) GetFileType(idx int, ) StringName {
  classNameV := StringNameFromStr("EditorFileSystemDirectory")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_file_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 659327637) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileSystemDirectory) GetFileScriptClassName(idx int, ) String {
  classNameV := StringNameFromStr("EditorFileSystemDirectory")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_file_script_class_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileSystemDirectory) GetFileScriptClassExtends(idx int, ) String {
  classNameV := StringNameFromStr("EditorFileSystemDirectory")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_file_script_class_extends")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileSystemDirectory) GetFileImportIsValid(idx int, ) bool {
  classNameV := StringNameFromStr("EditorFileSystemDirectory")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_file_import_is_valid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileSystemDirectory) GetName() String {
  classNameV := StringNameFromStr("EditorFileSystemDirectory")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2841200299) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileSystemDirectory) GetPath() String {
  classNameV := StringNameFromStr("EditorFileSystemDirectory")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileSystemDirectory) GetParent() EditorFileSystemDirectory {
  classNameV := StringNameFromStr("EditorFileSystemDirectory")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_parent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 842323275) // FIXME: should cache?
  var ret EditorFileSystemDirectory
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileSystemDirectory) FindFileIndex(name String, ) int {
  classNameV := StringNameFromStr("EditorFileSystemDirectory")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_file_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1321353865) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileSystemDirectory) FindDirIndex(name String, ) int {
  classNameV := StringNameFromStr("EditorFileSystemDirectory")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_dir_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1321353865) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties