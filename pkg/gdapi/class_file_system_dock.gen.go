// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type FileSystemDock struct {
  VBoxContainer
}

func (me *FileSystemDock) BaseClass() string {
  return "FileSystemDock"
}



// Enums

func (me *FileSystemDock) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *FileSystemDock) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *FileSystemDock) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *FileSystemDock) NavigateToPath(path String, )  {
  classNameV := StringNameFromStr("FileSystemDock")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("navigate_to_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FileSystemDock) AddResourceTooltipPlugin(plugin EditorResourceTooltipPlugin, )  {
  classNameV := StringNameFromStr("FileSystemDock")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_resource_tooltip_plugin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2258356838) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(plugin.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FileSystemDock) RemoveResourceTooltipPlugin(plugin EditorResourceTooltipPlugin, )  {
  classNameV := StringNameFromStr("FileSystemDock")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_resource_tooltip_plugin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2258356838) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(plugin.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Signals

type FileSystemDockInheritSignalFn func(file String, )

func (me *FileSystemDock) ConnectInherit(subs SignalSubscribers, fn FileSystemDockInheritSignalFn) {
  sig := StringNameFromStr("inherit")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *FileSystemDock) DisconnectInherit(subs SignalSubscribers, fn FileSystemDockInheritSignalFn) {
  sig := StringNameFromStr("inherit")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type FileSystemDockInstantiateSignalFn func(files PackedStringArray, )

func (me *FileSystemDock) ConnectInstantiate(subs SignalSubscribers, fn FileSystemDockInstantiateSignalFn) {
  sig := StringNameFromStr("instantiate")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *FileSystemDock) DisconnectInstantiate(subs SignalSubscribers, fn FileSystemDockInstantiateSignalFn) {
  sig := StringNameFromStr("instantiate")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type FileSystemDockResourceRemovedSignalFn func(resource Resource, )

func (me *FileSystemDock) ConnectResourceRemoved(subs SignalSubscribers, fn FileSystemDockResourceRemovedSignalFn) {
  sig := StringNameFromStr("resource_removed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *FileSystemDock) DisconnectResourceRemoved(subs SignalSubscribers, fn FileSystemDockResourceRemovedSignalFn) {
  sig := StringNameFromStr("resource_removed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type FileSystemDockFileRemovedSignalFn func(file String, )

func (me *FileSystemDock) ConnectFileRemoved(subs SignalSubscribers, fn FileSystemDockFileRemovedSignalFn) {
  sig := StringNameFromStr("file_removed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *FileSystemDock) DisconnectFileRemoved(subs SignalSubscribers, fn FileSystemDockFileRemovedSignalFn) {
  sig := StringNameFromStr("file_removed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type FileSystemDockFolderRemovedSignalFn func(folder String, )

func (me *FileSystemDock) ConnectFolderRemoved(subs SignalSubscribers, fn FileSystemDockFolderRemovedSignalFn) {
  sig := StringNameFromStr("folder_removed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *FileSystemDock) DisconnectFolderRemoved(subs SignalSubscribers, fn FileSystemDockFolderRemovedSignalFn) {
  sig := StringNameFromStr("folder_removed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type FileSystemDockFilesMovedSignalFn func(old_file String, new_file String, )

func (me *FileSystemDock) ConnectFilesMoved(subs SignalSubscribers, fn FileSystemDockFilesMovedSignalFn) {
  sig := StringNameFromStr("files_moved")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *FileSystemDock) DisconnectFilesMoved(subs SignalSubscribers, fn FileSystemDockFilesMovedSignalFn) {
  sig := StringNameFromStr("files_moved")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type FileSystemDockFolderMovedSignalFn func(old_folder String, new_folder String, )

func (me *FileSystemDock) ConnectFolderMoved(subs SignalSubscribers, fn FileSystemDockFolderMovedSignalFn) {
  sig := StringNameFromStr("folder_moved")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *FileSystemDock) DisconnectFolderMoved(subs SignalSubscribers, fn FileSystemDockFolderMovedSignalFn) {
  sig := StringNameFromStr("folder_moved")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type FileSystemDockDisplayModeChangedSignalFn func()

func (me *FileSystemDock) ConnectDisplayModeChanged(subs SignalSubscribers, fn FileSystemDockDisplayModeChangedSignalFn) {
  sig := StringNameFromStr("display_mode_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *FileSystemDock) DisconnectDisplayModeChanged(subs SignalSubscribers, fn FileSystemDockDisplayModeChangedSignalFn) {
  sig := StringNameFromStr("display_mode_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
