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

type ptrsForInstancePlaceholderList struct {
  fnGetStoredValues gdc.MethodBindPtr
  fnCreateInstance gdc.MethodBindPtr
  fnGetInstancePath gdc.MethodBindPtr
}

var ptrsForInstancePlaceholder ptrsForInstancePlaceholderList

func initInstancePlaceholderPtrs(iface gdc.Interface) {

  className := StringNameFromStr("InstancePlaceholder")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_stored_values")
    defer methodName.Destroy()
    ptrsForInstancePlaceholder.fnGetStoredValues = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2230153369))
  }
  {
    methodName := StringNameFromStr("create_instance")
    defer methodName.Destroy()
    ptrsForInstancePlaceholder.fnCreateInstance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3794612210))
  }
  {
    methodName := StringNameFromStr("get_instance_path")
    defer methodName.Destroy()
    ptrsForInstancePlaceholder.fnGetInstancePath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
}

type InstancePlaceholder struct {
  Node
}

func (me *InstancePlaceholder) BaseClass() string {
  return "InstancePlaceholder"
}

func NewInstancePlaceholder() *InstancePlaceholder {
  str := StringNameFromStr("InstancePlaceholder") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &InstancePlaceholder{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *InstancePlaceholder) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *InstancePlaceholder) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InstancePlaceholder) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *InstancePlaceholder) GetStoredValues(with_order bool, ) Dictionary {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&with_order) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()
  pinner.Pin(&with_order)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInstancePlaceholder.fnGetStoredValues), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *InstancePlaceholder) CreateInstance(replace bool, custom_scene PackedScene, ) Node {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&replace) , custom_scene.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNode()
  pinner.Pin(&replace)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInstancePlaceholder.fnCreateInstance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *InstancePlaceholder) GetInstancePath() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInstancePlaceholder.fnGetInstancePath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
