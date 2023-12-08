// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type DirAccess struct {
  obj gdc.ObjectPtr
}

func (me *DirAccess) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *DirAccess) BaseClass() string {
  return "DirAccess"
}

func  DirAccessOpen(path String, ) { // TODO: return value
  // TODO: implement
}

func  DirAccessGetOpenError() { // TODO: return value
  // TODO: implement
}

func  (me *DirAccess) ListDirBegin() { // TODO: return value
  // TODO: implement
}

func  (me *DirAccess) GetNext() { // TODO: return value
  // TODO: implement
}

func  (me *DirAccess) CurrentIsDir() { // TODO: return value
  // TODO: implement
}

func  (me *DirAccess) ListDirEnd() { // TODO: return value
  // TODO: implement
}

func  (me *DirAccess) GetFiles() { // TODO: return value
  // TODO: implement
}

func  DirAccessGetFilesAt(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *DirAccess) GetDirectories() { // TODO: return value
  // TODO: implement
}

func  DirAccessGetDirectoriesAt(path String, ) { // TODO: return value
  // TODO: implement
}

func  DirAccessGetDriveCount() { // TODO: return value
  // TODO: implement
}

func  DirAccessGetDriveName(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DirAccess) GetCurrentDrive() { // TODO: return value
  // TODO: implement
}

func  (me *DirAccess) ChangeDir(to_dir String, ) { // TODO: return value
  // TODO: implement
}

func  (me *DirAccess) GetCurrentDir(include_drive bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *DirAccess) MakeDir(path String, ) { // TODO: return value
  // TODO: implement
}

func  DirAccessMakeDirAbsolute(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *DirAccess) MakeDirRecursive(path String, ) { // TODO: return value
  // TODO: implement
}

func  DirAccessMakeDirRecursiveAbsolute(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *DirAccess) FileExists(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *DirAccess) DirExists(path String, ) { // TODO: return value
  // TODO: implement
}

func  DirAccessDirExistsAbsolute(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *DirAccess) GetSpaceLeft() { // TODO: return value
  // TODO: implement
}

func  (me *DirAccess) Copy(from String, to String, chmod_flags int, ) { // TODO: return value
  // TODO: implement
}

func  DirAccessCopyAbsolute(from String, to String, chmod_flags int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DirAccess) Rename(from String, to String, ) { // TODO: return value
  // TODO: implement
}

func  DirAccessRenameAbsolute(from String, to String, ) { // TODO: return value
  // TODO: implement
}

func  (me *DirAccess) Remove(path String, ) { // TODO: return value
  // TODO: implement
}

func  DirAccessRemoveAbsolute(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *DirAccess) SetIncludeNavigational(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *DirAccess) GetIncludeNavigational() { // TODO: return value
  // TODO: implement
}

func  (me *DirAccess) SetIncludeHidden(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *DirAccess) GetIncludeHidden() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
