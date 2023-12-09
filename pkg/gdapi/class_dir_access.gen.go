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



// Enums

func (me *DirAccess) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *DirAccess) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  DirAccessOpen(path String, )  {
  panic("TODO: implement")
}

func  DirAccessGetOpenError()  {
  panic("TODO: implement")
}

func  (me *DirAccess) ListDirBegin()  {
  panic("TODO: implement")
}

func  (me *DirAccess) GetNext()  {
  panic("TODO: implement")
}

func  (me *DirAccess) CurrentIsDir()  {
  panic("TODO: implement")
}

func  (me *DirAccess) ListDirEnd()  {
  panic("TODO: implement")
}

func  (me *DirAccess) GetFiles()  {
  panic("TODO: implement")
}

func  DirAccessGetFilesAt(path String, )  {
  panic("TODO: implement")
}

func  (me *DirAccess) GetDirectories()  {
  panic("TODO: implement")
}

func  DirAccessGetDirectoriesAt(path String, )  {
  panic("TODO: implement")
}

func  DirAccessGetDriveCount()  {
  panic("TODO: implement")
}

func  DirAccessGetDriveName(idx int, )  {
  panic("TODO: implement")
}

func  (me *DirAccess) GetCurrentDrive()  {
  panic("TODO: implement")
}

func  (me *DirAccess) ChangeDir(to_dir String, )  {
  panic("TODO: implement")
}

func  (me *DirAccess) GetCurrentDir(include_drive bool, )  {
  panic("TODO: implement")
}

func  (me *DirAccess) MakeDir(path String, )  {
  panic("TODO: implement")
}

func  DirAccessMakeDirAbsolute(path String, )  {
  panic("TODO: implement")
}

func  (me *DirAccess) MakeDirRecursive(path String, )  {
  panic("TODO: implement")
}

func  DirAccessMakeDirRecursiveAbsolute(path String, )  {
  panic("TODO: implement")
}

func  (me *DirAccess) FileExists(path String, )  {
  panic("TODO: implement")
}

func  (me *DirAccess) DirExists(path String, )  {
  panic("TODO: implement")
}

func  DirAccessDirExistsAbsolute(path String, )  {
  panic("TODO: implement")
}

func  (me *DirAccess) GetSpaceLeft()  {
  panic("TODO: implement")
}

func  (me *DirAccess) Copy(from String, to String, chmod_flags int, )  {
  panic("TODO: implement")
}

func  DirAccessCopyAbsolute(from String, to String, chmod_flags int, )  {
  panic("TODO: implement")
}

func  (me *DirAccess) Rename(from String, to String, )  {
  panic("TODO: implement")
}

func  DirAccessRenameAbsolute(from String, to String, )  {
  panic("TODO: implement")
}

func  (me *DirAccess) Remove(path String, )  {
  panic("TODO: implement")
}

func  DirAccessRemoveAbsolute(path String, )  {
  panic("TODO: implement")
}

func  (me *DirAccess) SetIncludeNavigational(enable bool, )  {
  panic("TODO: implement")
}

func  (me *DirAccess) GetIncludeNavigational()  {
  panic("TODO: implement")
}

func  (me *DirAccess) SetIncludeHidden(enable bool, )  {
  panic("TODO: implement")
}

func  (me *DirAccess) GetIncludeHidden()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
