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

type ptrsForMaterialList struct {
  fnXGetShaderRid gdc.MethodBindPtr
  fnXGetShaderMode gdc.MethodBindPtr
  fnXCanDoNextPass gdc.MethodBindPtr
  fnXCanUseRenderPriority gdc.MethodBindPtr
  fnSetNextPass gdc.MethodBindPtr
  fnGetNextPass gdc.MethodBindPtr
  fnSetRenderPriority gdc.MethodBindPtr
  fnGetRenderPriority gdc.MethodBindPtr
  fnInspectNativeShaderCode gdc.MethodBindPtr
  fnCreatePlaceholder gdc.MethodBindPtr
}

var ptrsForMaterial ptrsForMaterialList

func initMaterialPtrs(iface gdc.Interface) {

  className := StringNameFromStr("Material")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_next_pass")
    defer methodName.Destroy()
    ptrsForMaterial.fnSetNextPass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2757459619))
  }
  {
    methodName := StringNameFromStr("get_next_pass")
    defer methodName.Destroy()
    ptrsForMaterial.fnGetNextPass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 5934680))
  }
  {
    methodName := StringNameFromStr("set_render_priority")
    defer methodName.Destroy()
    ptrsForMaterial.fnSetRenderPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_render_priority")
    defer methodName.Destroy()
    ptrsForMaterial.fnGetRenderPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("inspect_native_shader_code")
    defer methodName.Destroy()
    ptrsForMaterial.fnInspectNativeShaderCode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("create_placeholder")
    defer methodName.Destroy()
    ptrsForMaterial.fnCreatePlaceholder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 121922552))
  }
}

type Material struct {
  Resource
}

func (me *Material) BaseClass() string {
  return "Material"
}

func NewMaterial() *Material {
  str := StringNameFromStr("Material") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Material{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Constants

var (
  MaterialRenderPriorityMax = "127" // TODO: construct correctly
  MaterialRenderPriorityMin = "-128" // TODO: construct correctly
)

// Enums

func (me *Material) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Material) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Material) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Material) SetNextPass(next_pass Material, )  {
  cargs := []gdc.ConstTypePtr{next_pass.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMaterial.fnSetNextPass), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Material) GetNextPass() Material {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMaterial()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMaterial.fnGetNextPass), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Material) SetRenderPriority(priority int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMaterial.fnSetRenderPriority), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Material) GetRenderPriority() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMaterial.fnGetRenderPriority), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Material) InspectNativeShaderCode()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMaterial.fnInspectNativeShaderCode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Material) CreatePlaceholder() Resource {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewResource()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMaterial.fnCreatePlaceholder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
