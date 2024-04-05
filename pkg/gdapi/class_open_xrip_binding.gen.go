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

type ptrsForOpenXRIPBindingList struct {
  fnSetAction gdc.MethodBindPtr
  fnGetAction gdc.MethodBindPtr
  fnGetPathCount gdc.MethodBindPtr
  fnSetPaths gdc.MethodBindPtr
  fnGetPaths gdc.MethodBindPtr
  fnHasPath gdc.MethodBindPtr
  fnAddPath gdc.MethodBindPtr
  fnRemovePath gdc.MethodBindPtr
}

var ptrsForOpenXRIPBinding ptrsForOpenXRIPBindingList

func initOpenXRIPBindingPtrs(iface gdc.Interface) {

  className := StringNameFromStr("OpenXRIPBinding")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_action")
    defer methodName.Destroy()
    ptrsForOpenXRIPBinding.fnSetAction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 349361333))
  }
  {
    methodName := StringNameFromStr("get_action")
    defer methodName.Destroy()
    ptrsForOpenXRIPBinding.fnGetAction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4072409085))
  }
  {
    methodName := StringNameFromStr("get_path_count")
    defer methodName.Destroy()
    ptrsForOpenXRIPBinding.fnGetPathCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_paths")
    defer methodName.Destroy()
    ptrsForOpenXRIPBinding.fnSetPaths = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4015028928))
  }
  {
    methodName := StringNameFromStr("get_paths")
    defer methodName.Destroy()
    ptrsForOpenXRIPBinding.fnGetPaths = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
  }
  {
    methodName := StringNameFromStr("has_path")
    defer methodName.Destroy()
    ptrsForOpenXRIPBinding.fnHasPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
  }
  {
    methodName := StringNameFromStr("add_path")
    defer methodName.Destroy()
    ptrsForOpenXRIPBinding.fnAddPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("remove_path")
    defer methodName.Destroy()
    ptrsForOpenXRIPBinding.fnRemovePath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
}

type OpenXRIPBinding struct {
  Resource
}

func (me *OpenXRIPBinding) BaseClass() string {
  return "OpenXRIPBinding"
}

func NewOpenXRIPBinding() *OpenXRIPBinding {
  str := StringNameFromStr("OpenXRIPBinding") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &OpenXRIPBinding{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *OpenXRIPBinding) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *OpenXRIPBinding) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *OpenXRIPBinding) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *OpenXRIPBinding) SetAction(action OpenXRAction, )  {
  cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRIPBinding.fnSetAction), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OpenXRIPBinding) GetAction() OpenXRAction {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewOpenXRAction()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRIPBinding.fnGetAction), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *OpenXRIPBinding) GetPathCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRIPBinding.fnGetPathCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OpenXRIPBinding) SetPaths(paths PackedStringArray, )  {
  cargs := []gdc.ConstTypePtr{paths.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRIPBinding.fnSetPaths), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OpenXRIPBinding) GetPaths() PackedStringArray {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRIPBinding.fnGetPaths), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *OpenXRIPBinding) HasPath(path String, ) bool {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRIPBinding.fnHasPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OpenXRIPBinding) AddPath(path String, )  {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRIPBinding.fnAddPath), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OpenXRIPBinding) RemovePath(path String, )  {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRIPBinding.fnRemovePath), me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
