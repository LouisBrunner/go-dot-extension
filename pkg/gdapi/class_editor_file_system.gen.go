// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func  (me *EditorFileSystem) GetFilesystem()  {
  panic("TODO: implement")
}

func  (me *EditorFileSystem) IsScanning()  {
  panic("TODO: implement")
}

func  (me *EditorFileSystem) GetScanningProgress()  {
  panic("TODO: implement")
}

func  (me *EditorFileSystem) Scan()  {
  panic("TODO: implement")
}

func  (me *EditorFileSystem) ScanSources()  {
  panic("TODO: implement")
}

func  (me *EditorFileSystem) UpdateFile(path String, )  {
  panic("TODO: implement")
}

func  (me *EditorFileSystem) GetFilesystemPath(path String, )  {
  panic("TODO: implement")
}

func  (me *EditorFileSystem) GetFileType(path String, )  {
  panic("TODO: implement")
}

func  (me *EditorFileSystem) ReimportFiles(files PackedStringArray, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
