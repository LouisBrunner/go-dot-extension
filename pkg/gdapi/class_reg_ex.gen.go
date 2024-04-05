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

type ptrsForRegExList struct {
  fnCreateFromString gdc.MethodBindPtr
  fnClear gdc.MethodBindPtr
  fnCompile gdc.MethodBindPtr
  fnSearch gdc.MethodBindPtr
  fnSearchAll gdc.MethodBindPtr
  fnSub gdc.MethodBindPtr
  fnIsValid gdc.MethodBindPtr
  fnGetPattern gdc.MethodBindPtr
  fnGetGroupCount gdc.MethodBindPtr
  fnGetNames gdc.MethodBindPtr
}

var ptrsForRegEx ptrsForRegExList

func initRegExPtrs(iface gdc.Interface) {

  className := StringNameFromStr("RegEx")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("create_from_string")
    defer methodName.Destroy()
    ptrsForRegEx.fnCreateFromString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2150300909))
  }
  {
    methodName := StringNameFromStr("clear")
    defer methodName.Destroy()
    ptrsForRegEx.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("compile")
    defer methodName.Destroy()
    ptrsForRegEx.fnCompile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
  }
  {
    methodName := StringNameFromStr("search")
    defer methodName.Destroy()
    ptrsForRegEx.fnSearch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3365977994))
  }
  {
    methodName := StringNameFromStr("search_all")
    defer methodName.Destroy()
    ptrsForRegEx.fnSearchAll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 849021363))
  }
  {
    methodName := StringNameFromStr("sub")
    defer methodName.Destroy()
    ptrsForRegEx.fnSub = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 54019702))
  }
  {
    methodName := StringNameFromStr("is_valid")
    defer methodName.Destroy()
    ptrsForRegEx.fnIsValid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_pattern")
    defer methodName.Destroy()
    ptrsForRegEx.fnGetPattern = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("get_group_count")
    defer methodName.Destroy()
    ptrsForRegEx.fnGetGroupCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_names")
    defer methodName.Destroy()
    ptrsForRegEx.fnGetNames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
  }
}

type RegEx struct {
  RefCounted
}

func (me *RegEx) BaseClass() string {
  return "RegEx"
}

func NewRegEx() *RegEx {
  str := StringNameFromStr("RegEx") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &RegEx{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *RegEx) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RegEx) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RegEx) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  RegExCreateFromString(pattern String, ) RegEx {
  cargs := []gdc.ConstTypePtr{pattern.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRegEx()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRegEx.fnCreateFromString), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RegEx) Clear()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRegEx.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RegEx) Compile(pattern String, ) Error {
  cargs := []gdc.ConstTypePtr{pattern.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRegEx.fnCompile), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RegEx) Search(subject String, offset int64, end int64, ) RegExMatch {
  cargs := []gdc.ConstTypePtr{subject.AsCTypePtr(), gdc.ConstTypePtr(&offset) , gdc.ConstTypePtr(&end) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRegExMatch()
  pinner.Pin(&offset)
  pinner.Pin(&end)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRegEx.fnSearch), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RegEx) SearchAll(subject String, offset int64, end int64, ) []RegExMatch {
  cargs := []gdc.ConstTypePtr{subject.AsCTypePtr(), gdc.ConstTypePtr(&offset) , gdc.ConstTypePtr(&end) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()
  pinner.Pin(&offset)
  pinner.Pin(&end)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRegEx.fnSearchAll), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[RegExMatch](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *RegEx) Sub(subject String, replacement String, all bool, offset int64, end int64, ) String {
  cargs := []gdc.ConstTypePtr{subject.AsCTypePtr(), replacement.AsCTypePtr(), gdc.ConstTypePtr(&all) , gdc.ConstTypePtr(&offset) , gdc.ConstTypePtr(&end) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&all)
  pinner.Pin(&offset)
  pinner.Pin(&end)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRegEx.fnSub), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RegEx) IsValid() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRegEx.fnIsValid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RegEx) GetPattern() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRegEx.fnGetPattern), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RegEx) GetGroupCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRegEx.fnGetGroupCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RegEx) GetNames() PackedStringArray {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRegEx.fnGetNames), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
