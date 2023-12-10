// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type RegEx struct {
  obj gdc.ObjectPtr
}

func (me *RegEx) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RegEx) BaseClass() string {
  return "RegEx"
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
  classNameV := StringNameFromStr("RegEx")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_from_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2150300909) // FIXME: should cache?
  var ret RegEx
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(pattern.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RegEx) Clear()  {
  classNameV := StringNameFromStr("RegEx")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RegEx) Compile(pattern String, ) Error {
  classNameV := StringNameFromStr("RegEx")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("compile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(pattern.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RegEx) Search(subject String, offset int, end int, ) RegExMatch {
  classNameV := StringNameFromStr("RegEx")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("search")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4087180739) // FIXME: should cache?
  var ret RegExMatch
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(subject.AsCTypePtr()), gdc.ConstTypePtr(&offset), gdc.ConstTypePtr(&end), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RegEx) SearchAll(subject String, offset int, end int, ) RegExMatch {
  classNameV := StringNameFromStr("RegEx")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("search_all")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3354100289) // FIXME: should cache?
  var ret RegExMatch
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(subject.AsCTypePtr()), gdc.ConstTypePtr(&offset), gdc.ConstTypePtr(&end), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RegEx) Sub(subject String, replacement String, all bool, offset int, end int, ) String {
  classNameV := StringNameFromStr("RegEx")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("sub")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 758293621) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(subject.AsCTypePtr()), gdc.ConstTypePtr(replacement.AsCTypePtr()), gdc.ConstTypePtr(&all), gdc.ConstTypePtr(&offset), gdc.ConstTypePtr(&end), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RegEx) IsValid() bool {
  classNameV := StringNameFromStr("RegEx")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_valid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RegEx) GetPattern() String {
  classNameV := StringNameFromStr("RegEx")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pattern")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RegEx) GetGroupCount() int {
  classNameV := StringNameFromStr("RegEx")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_group_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RegEx) GetNames() PackedStringArray {
  classNameV := StringNameFromStr("RegEx")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_names")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties