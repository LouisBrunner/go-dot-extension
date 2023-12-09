// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationRegion3D struct {
  obj gdc.ObjectPtr
}

func (me *NavigationRegion3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationRegion3D) BaseClass() string {
  return "NavigationRegion3D"
}



// Enums

func (me *NavigationRegion3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *NavigationRegion3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NavigationRegion3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *NavigationRegion3D) SetNavigationMesh(navigation_mesh NavigationMesh, )  {
  panic("TODO: implement")
}

func  (me *NavigationRegion3D) GetNavigationMesh()  {
  panic("TODO: implement")
}

func  (me *NavigationRegion3D) SetEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationRegion3D) IsEnabled()  {
  panic("TODO: implement")
}

func  (me *NavigationRegion3D) SetUseEdgeConnections(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationRegion3D) GetUseEdgeConnections()  {
  panic("TODO: implement")
}

func  (me *NavigationRegion3D) SetNavigationLayers(navigation_layers int, )  {
  panic("TODO: implement")
}

func  (me *NavigationRegion3D) GetNavigationLayers()  {
  panic("TODO: implement")
}

func  (me *NavigationRegion3D) SetNavigationLayerValue(layer_number int, value bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationRegion3D) GetNavigationLayerValue(layer_number int, )  {
  panic("TODO: implement")
}

func  (me *NavigationRegion3D) GetRegionRid()  {
  panic("TODO: implement")
}

func  (me *NavigationRegion3D) SetEnterCost(enter_cost float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationRegion3D) GetEnterCost()  {
  panic("TODO: implement")
}

func  (me *NavigationRegion3D) SetTravelCost(travel_cost float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationRegion3D) GetTravelCost()  {
  panic("TODO: implement")
}

func  (me *NavigationRegion3D) BakeNavigationMesh(on_thread bool, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
