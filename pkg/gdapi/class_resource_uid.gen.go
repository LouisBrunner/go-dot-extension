// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ResourceUID struct {
  obj gdc.ObjectPtr
}

func (me *ResourceUID) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ResourceUID) BaseClass() string {
  return "ResourceUID"
}



// Constants

var (
  ResourceUIDInvalidId = "-1" // TODO: construct correctly
)

// Enums

func (me *ResourceUID) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceUID) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceUID) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ResourceUID) IdToText(id int, ) String {
  classNameV := StringNameFromStr("ResourceUID")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("id_to_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ResourceUID) TextToId(text_id String, ) int {
  classNameV := StringNameFromStr("ResourceUID")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("text_to_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1321353865) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(text_id.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ResourceUID) CreateId() int {
  classNameV := StringNameFromStr("ResourceUID")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ResourceUID) HasId(id int, ) bool {
  classNameV := StringNameFromStr("ResourceUID")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ResourceUID) AddId(id int, path String, )  {
  classNameV := StringNameFromStr("ResourceUID")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ResourceUID) SetId(id int, path String, )  {
  classNameV := StringNameFromStr("ResourceUID")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ResourceUID) GetIdPath(id int, ) String {
  classNameV := StringNameFromStr("ResourceUID")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_id_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ResourceUID) RemoveId(id int, )  {
  classNameV := StringNameFromStr("ResourceUID")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties