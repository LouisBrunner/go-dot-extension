// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorFileSystem struct {
  obj gdc.ObjectPtr
}

func (me *EditorFileSystem) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorFileSystem) BaseClass() string {
  return "EditorFileSystem"
}



// Enums

func (me *EditorFileSystem) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorFileSystem) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorFileSystem) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EditorFileSystem) GetFilesystem() EditorFileSystemDirectory {
  classNameV := StringNameFromStr("EditorFileSystem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_filesystem")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 842323275) // FIXME: should cache?
  var ret EditorFileSystemDirectory
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileSystem) IsScanning() bool {
  classNameV := StringNameFromStr("EditorFileSystem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_scanning")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileSystem) GetScanningProgress() float32 {
  classNameV := StringNameFromStr("EditorFileSystem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scanning_progress")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileSystem) Scan()  {
  classNameV := StringNameFromStr("EditorFileSystem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("scan")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorFileSystem) ScanSources()  {
  classNameV := StringNameFromStr("EditorFileSystem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("scan_sources")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorFileSystem) UpdateFile(path String, )  {
  classNameV := StringNameFromStr("EditorFileSystem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("update_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorFileSystem) GetFilesystemPath(path String, ) EditorFileSystemDirectory {
  classNameV := StringNameFromStr("EditorFileSystem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_filesystem_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3188521125) // FIXME: should cache?
  var ret EditorFileSystemDirectory
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileSystem) GetFileType(path String, ) String {
  classNameV := StringNameFromStr("EditorFileSystem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_file_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3135753539) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorFileSystem) ReimportFiles(files PackedStringArray, )  {
  classNameV := StringNameFromStr("EditorFileSystem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("reimport_files")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4015028928) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(files.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties
// Signals
// FIXME: can't seem to be able to connect them from this side of the API