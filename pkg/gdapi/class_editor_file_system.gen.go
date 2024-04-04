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

type EditorFileSystem struct {
  Node
}

func (me *EditorFileSystem) BaseClass() string {
  return "EditorFileSystem"
}

func NewEditorFileSystem() *EditorFileSystem {
  str := StringNameFromStr("EditorFileSystem") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorFileSystem{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewEditorFileSystemDirectory()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorFileSystem) IsScanning() bool {
  classNameV := StringNameFromStr("EditorFileSystem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_scanning")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorFileSystem) GetScanningProgress() float64 {
  classNameV := StringNameFromStr("EditorFileSystem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scanning_progress")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorFileSystem) Scan()  {
  classNameV := StringNameFromStr("EditorFileSystem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("scan")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorFileSystem) ScanSources()  {
  classNameV := StringNameFromStr("EditorFileSystem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("scan_sources")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorFileSystem) UpdateFile(path String, )  {
  classNameV := StringNameFromStr("EditorFileSystem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("update_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorFileSystem) GetFilesystemPath(path String, ) EditorFileSystemDirectory {
  classNameV := StringNameFromStr("EditorFileSystem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_filesystem_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3188521125) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewEditorFileSystemDirectory()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorFileSystem) GetFileType(path String, ) String {
  classNameV := StringNameFromStr("EditorFileSystem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_file_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3135753539) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorFileSystem) ReimportFiles(files PackedStringArray, )  {
  classNameV := StringNameFromStr("EditorFileSystem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("reimport_files")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4015028928) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{files.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals

type EditorFileSystemFilesystemChangedSignalFn func()

func (me *EditorFileSystem) ConnectFilesystemChanged(subs SignalSubscribers, fn EditorFileSystemFilesystemChangedSignalFn) {
  sig := StringNameFromStr("filesystem_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorFileSystem) DisconnectFilesystemChanged(subs SignalSubscribers, fn EditorFileSystemFilesystemChangedSignalFn) {
  sig := StringNameFromStr("filesystem_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorFileSystemScriptClassesUpdatedSignalFn func()

func (me *EditorFileSystem) ConnectScriptClassesUpdated(subs SignalSubscribers, fn EditorFileSystemScriptClassesUpdatedSignalFn) {
  sig := StringNameFromStr("script_classes_updated")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorFileSystem) DisconnectScriptClassesUpdated(subs SignalSubscribers, fn EditorFileSystemScriptClassesUpdatedSignalFn) {
  sig := StringNameFromStr("script_classes_updated")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorFileSystemSourcesChangedSignalFn func(exist bool, )

func (me *EditorFileSystem) ConnectSourcesChanged(subs SignalSubscribers, fn EditorFileSystemSourcesChangedSignalFn) {
  sig := StringNameFromStr("sources_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorFileSystem) DisconnectSourcesChanged(subs SignalSubscribers, fn EditorFileSystemSourcesChangedSignalFn) {
  sig := StringNameFromStr("sources_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorFileSystemResourcesReimportedSignalFn func(resources PackedStringArray, )

func (me *EditorFileSystem) ConnectResourcesReimported(subs SignalSubscribers, fn EditorFileSystemResourcesReimportedSignalFn) {
  sig := StringNameFromStr("resources_reimported")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorFileSystem) DisconnectResourcesReimported(subs SignalSubscribers, fn EditorFileSystemResourcesReimportedSignalFn) {
  sig := StringNameFromStr("resources_reimported")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorFileSystemResourcesReloadSignalFn func(resources PackedStringArray, )

func (me *EditorFileSystem) ConnectResourcesReload(subs SignalSubscribers, fn EditorFileSystemResourcesReloadSignalFn) {
  sig := StringNameFromStr("resources_reload")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorFileSystem) DisconnectResourcesReload(subs SignalSubscribers, fn EditorFileSystemResourcesReloadSignalFn) {
  sig := StringNameFromStr("resources_reload")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
