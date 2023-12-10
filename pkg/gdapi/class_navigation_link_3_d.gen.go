// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type NavigationLink3D struct {
  obj gdc.ObjectPtr
}

func (me *NavigationLink3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationLink3D) BaseClass() string {
  return "NavigationLink3D"
}



// Enums

func (me *NavigationLink3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *NavigationLink3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NavigationLink3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *NavigationLink3D) SetEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("NavigationLink3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationLink3D) IsEnabled() bool {
  classNameV := StringNameFromStr("NavigationLink3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationLink3D) SetBidirectional(bidirectional bool, )  {
  classNameV := StringNameFromStr("NavigationLink3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bidirectional")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bidirectional), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationLink3D) IsBidirectional() bool {
  classNameV := StringNameFromStr("NavigationLink3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_bidirectional")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationLink3D) SetNavigationLayers(navigation_layers int, )  {
  classNameV := StringNameFromStr("NavigationLink3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_navigation_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&navigation_layers), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationLink3D) GetNavigationLayers() int {
  classNameV := StringNameFromStr("NavigationLink3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationLink3D) SetNavigationLayerValue(layer_number int, value bool, )  {
  classNameV := StringNameFromStr("NavigationLink3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_navigation_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationLink3D) GetNavigationLayerValue(layer_number int, ) bool {
  classNameV := StringNameFromStr("NavigationLink3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationLink3D) SetStartPosition(position Vector3, )  {
  classNameV := StringNameFromStr("NavigationLink3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_start_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationLink3D) GetStartPosition() Vector3 {
  classNameV := StringNameFromStr("NavigationLink3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_start_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationLink3D) SetEndPosition(position Vector3, )  {
  classNameV := StringNameFromStr("NavigationLink3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_end_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationLink3D) GetEndPosition() Vector3 {
  classNameV := StringNameFromStr("NavigationLink3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_end_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationLink3D) SetGlobalStartPosition(position Vector3, )  {
  classNameV := StringNameFromStr("NavigationLink3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_global_start_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationLink3D) GetGlobalStartPosition() Vector3 {
  classNameV := StringNameFromStr("NavigationLink3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_global_start_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationLink3D) SetGlobalEndPosition(position Vector3, )  {
  classNameV := StringNameFromStr("NavigationLink3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_global_end_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationLink3D) GetGlobalEndPosition() Vector3 {
  classNameV := StringNameFromStr("NavigationLink3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_global_end_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationLink3D) SetEnterCost(enter_cost float32, )  {
  classNameV := StringNameFromStr("NavigationLink3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enter_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enter_cost), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationLink3D) GetEnterCost() float32 {
  classNameV := StringNameFromStr("NavigationLink3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_enter_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationLink3D) SetTravelCost(travel_cost float32, )  {
  classNameV := StringNameFromStr("NavigationLink3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_travel_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&travel_cost), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationLink3D) GetTravelCost() float32 {
  classNameV := StringNameFromStr("NavigationLink3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_travel_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *NavigationLink3D) GetPropEnabled() bool {
  panic("TODO: implement")
}

func (me *NavigationLink3D) SetPropEnabled(value bool) {
  panic("TODO: implement")
}

func (me *NavigationLink3D) GetPropBidirectional() bool {
  panic("TODO: implement")
}

func (me *NavigationLink3D) SetPropBidirectional(value bool) {
  panic("TODO: implement")
}

func (me *NavigationLink3D) GetPropNavigationLayers() int {
  panic("TODO: implement")
}

func (me *NavigationLink3D) SetPropNavigationLayers(value int) {
  panic("TODO: implement")
}

func (me *NavigationLink3D) GetPropStartPosition() Vector3 {
  panic("TODO: implement")
}

func (me *NavigationLink3D) SetPropStartPosition(value Vector3) {
  panic("TODO: implement")
}

func (me *NavigationLink3D) GetPropEndPosition() Vector3 {
  panic("TODO: implement")
}

func (me *NavigationLink3D) SetPropEndPosition(value Vector3) {
  panic("TODO: implement")
}

func (me *NavigationLink3D) GetPropEnterCost() float32 {
  panic("TODO: implement")
}

func (me *NavigationLink3D) SetPropEnterCost(value float32) {
  panic("TODO: implement")
}

func (me *NavigationLink3D) GetPropTravelCost() float32 {
  panic("TODO: implement")
}

func (me *NavigationLink3D) SetPropTravelCost(value float32) {
  panic("TODO: implement")
}