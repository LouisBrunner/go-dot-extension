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

type ptrsForRDUniformList struct {
  fnSetUniformType gdc.MethodBindPtr
  fnGetUniformType gdc.MethodBindPtr
  fnSetBinding gdc.MethodBindPtr
  fnGetBinding gdc.MethodBindPtr
  fnAddId gdc.MethodBindPtr
  fnClearIds gdc.MethodBindPtr
  fnGetIds gdc.MethodBindPtr
}

var ptrsForRDUniform ptrsForRDUniformList

func initRDUniformPtrs(iface gdc.Interface) {

  className := StringNameFromStr("RDUniform")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_uniform_type")
    defer methodName.Destroy()
    ptrsForRDUniform.fnSetUniformType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1664894931))
  }
  {
    methodName := StringNameFromStr("get_uniform_type")
    defer methodName.Destroy()
    ptrsForRDUniform.fnGetUniformType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 475470040))
  }
  {
    methodName := StringNameFromStr("set_binding")
    defer methodName.Destroy()
    ptrsForRDUniform.fnSetBinding = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_binding")
    defer methodName.Destroy()
    ptrsForRDUniform.fnGetBinding = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("add_id")
    defer methodName.Destroy()
    ptrsForRDUniform.fnAddId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
  }
  {
    methodName := StringNameFromStr("clear_ids")
    defer methodName.Destroy()
    ptrsForRDUniform.fnClearIds = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("get_ids")
    defer methodName.Destroy()
    ptrsForRDUniform.fnGetIds = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
}

type RDUniform struct {
  RefCounted
}

func (me *RDUniform) BaseClass() string {
  return "RDUniform"
}

func NewRDUniform() *RDUniform {
  str := StringNameFromStr("RDUniform") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &RDUniform{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *RDUniform) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RDUniform) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RDUniform) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RDUniform) SetUniformType(p_member RenderingDeviceUniformType, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDUniform.fnSetUniformType), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDUniform) GetUniformType() RenderingDeviceUniformType {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceUniformType

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDUniform.fnGetUniformType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDUniform) SetBinding(p_member int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDUniform.fnSetBinding), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDUniform) GetBinding() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDUniform.fnGetBinding), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDUniform) AddId(id RID, )  {
  cargs := []gdc.ConstTypePtr{id.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDUniform.fnAddId), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDUniform) ClearIds()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDUniform.fnClearIds), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDUniform) GetIds() []RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDUniform.fnGetIds), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[RID](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
