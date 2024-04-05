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

type ptrsForZIPReaderList struct {
  fnOpen gdc.MethodBindPtr
  fnClose gdc.MethodBindPtr
  fnGetFiles gdc.MethodBindPtr
  fnReadFile gdc.MethodBindPtr
  fnFileExists gdc.MethodBindPtr
}

var ptrsForZIPReader ptrsForZIPReaderList

func initZIPReaderPtrs(iface gdc.Interface) {

  className := StringNameFromStr("ZIPReader")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("open")
    defer methodName.Destroy()
    ptrsForZIPReader.fnOpen = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
  }
  {
    methodName := StringNameFromStr("close")
    defer methodName.Destroy()
    ptrsForZIPReader.fnClose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166280745))
  }
  {
    methodName := StringNameFromStr("get_files")
    defer methodName.Destroy()
    ptrsForZIPReader.fnGetFiles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2981934095))
  }
  {
    methodName := StringNameFromStr("read_file")
    defer methodName.Destroy()
    ptrsForZIPReader.fnReadFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 740857591))
  }
  {
    methodName := StringNameFromStr("file_exists")
    defer methodName.Destroy()
    ptrsForZIPReader.fnFileExists = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 35364943))
  }
}

type ZIPReader struct {
  RefCounted
}

func (me *ZIPReader) BaseClass() string {
  return "ZIPReader"
}

func NewZIPReader() *ZIPReader {
  str := StringNameFromStr("ZIPReader") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ZIPReader{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ZIPReader) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ZIPReader) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ZIPReader) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ZIPReader) Open(path String, ) Error {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForZIPReader.fnOpen), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ZIPReader) Close() Error {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForZIPReader.fnClose), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ZIPReader) GetFiles() PackedStringArray {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForZIPReader.fnGetFiles), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ZIPReader) ReadFile(path String, case_sensitive bool, ) PackedByteArray {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&case_sensitive) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedByteArray()
  pinner.Pin(&case_sensitive)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForZIPReader.fnReadFile), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ZIPReader) FileExists(path String, case_sensitive bool, ) bool {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&case_sensitive) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&case_sensitive)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForZIPReader.fnFileExists), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

// Signals
