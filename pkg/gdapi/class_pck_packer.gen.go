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

type ptrsForPCKPackerList struct {
  fnPckStart gdc.MethodBindPtr
  fnAddFile gdc.MethodBindPtr
  fnFlush gdc.MethodBindPtr
}

var ptrsForPCKPacker ptrsForPCKPackerList

func initPCKPackerPtrs(iface gdc.Interface) {

  className := StringNameFromStr("PCKPacker")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("pck_start")
    defer methodName.Destroy()
    ptrsForPCKPacker.fnPckStart = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 508410629))
  }
  {
    methodName := StringNameFromStr("add_file")
    defer methodName.Destroy()
    ptrsForPCKPacker.fnAddFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2215643711))
  }
  {
    methodName := StringNameFromStr("flush")
    defer methodName.Destroy()
    ptrsForPCKPacker.fnFlush = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1633102583))
  }
}

type PCKPacker struct {
  RefCounted
}

func (me *PCKPacker) BaseClass() string {
  return "PCKPacker"
}

func NewPCKPacker() *PCKPacker {
  str := StringNameFromStr("PCKPacker") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PCKPacker{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *PCKPacker) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PCKPacker) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PCKPacker) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PCKPacker) PckStart(pck_name String, alignment int64, key String, encrypt_directory bool, ) Error {
  cargs := []gdc.ConstTypePtr{pck_name.AsCTypePtr(), gdc.ConstTypePtr(&alignment) , key.AsCTypePtr(), gdc.ConstTypePtr(&encrypt_directory) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&alignment)
  pinner.Pin(&encrypt_directory)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPCKPacker.fnPckStart), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *PCKPacker) AddFile(pck_path String, source_path String, encrypt bool, ) Error {
  cargs := []gdc.ConstTypePtr{pck_path.AsCTypePtr(), source_path.AsCTypePtr(), gdc.ConstTypePtr(&encrypt) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&encrypt)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPCKPacker.fnAddFile), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *PCKPacker) Flush(verbose bool, ) Error {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&verbose) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&verbose)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPCKPacker.fnFlush), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

// Signals
