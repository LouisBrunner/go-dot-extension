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

type ptrsForEditorFileSystemList struct {
  fnGetFilesystem gdc.MethodBindPtr
  fnIsScanning gdc.MethodBindPtr
  fnGetScanningProgress gdc.MethodBindPtr
  fnScan gdc.MethodBindPtr
  fnScanSources gdc.MethodBindPtr
  fnUpdateFile gdc.MethodBindPtr
  fnGetFilesystemPath gdc.MethodBindPtr
  fnGetFileType gdc.MethodBindPtr
  fnReimportFiles gdc.MethodBindPtr
}

var ptrsForEditorFileSystem ptrsForEditorFileSystemList

func initEditorFileSystemPtrs(iface gdc.Interface) {

  className := StringNameFromStr("EditorFileSystem")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_filesystem")
    defer methodName.Destroy()
    ptrsForEditorFileSystem.fnGetFilesystem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 842323275))
  }
  {
    methodName := StringNameFromStr("is_scanning")
    defer methodName.Destroy()
    ptrsForEditorFileSystem.fnIsScanning = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_scanning_progress")
    defer methodName.Destroy()
    ptrsForEditorFileSystem.fnGetScanningProgress = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("scan")
    defer methodName.Destroy()
    ptrsForEditorFileSystem.fnScan = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("scan_sources")
    defer methodName.Destroy()
    ptrsForEditorFileSystem.fnScanSources = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("update_file")
    defer methodName.Destroy()
    ptrsForEditorFileSystem.fnUpdateFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_filesystem_path")
    defer methodName.Destroy()
    ptrsForEditorFileSystem.fnGetFilesystemPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3188521125))
  }
  {
    methodName := StringNameFromStr("get_file_type")
    defer methodName.Destroy()
    ptrsForEditorFileSystem.fnGetFileType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3135753539))
  }
  {
    methodName := StringNameFromStr("reimport_files")
    defer methodName.Destroy()
    ptrsForEditorFileSystem.fnReimportFiles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4015028928))
  }
}

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
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewEditorFileSystemDirectory()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileSystem.fnGetFilesystem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorFileSystem) IsScanning() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileSystem.fnIsScanning), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorFileSystem) GetScanningProgress() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileSystem.fnGetScanningProgress), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorFileSystem) Scan()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileSystem.fnScan), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorFileSystem) ScanSources()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileSystem.fnScanSources), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorFileSystem) UpdateFile(path String, )  {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileSystem.fnUpdateFile), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorFileSystem) GetFilesystemPath(path String, ) EditorFileSystemDirectory {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewEditorFileSystemDirectory()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileSystem.fnGetFilesystemPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorFileSystem) GetFileType(path String, ) String {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileSystem.fnGetFileType), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorFileSystem) ReimportFiles(files PackedStringArray, )  {
  cargs := []gdc.ConstTypePtr{files.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileSystem.fnReimportFiles), me.obj, unsafe.SliceData(cargs), nil)

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
