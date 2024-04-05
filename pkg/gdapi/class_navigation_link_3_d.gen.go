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

type ptrsForNavigationLink3DList struct {
  fnSetEnabled gdc.MethodBindPtr
  fnIsEnabled gdc.MethodBindPtr
  fnSetBidirectional gdc.MethodBindPtr
  fnIsBidirectional gdc.MethodBindPtr
  fnSetNavigationLayers gdc.MethodBindPtr
  fnGetNavigationLayers gdc.MethodBindPtr
  fnSetNavigationLayerValue gdc.MethodBindPtr
  fnGetNavigationLayerValue gdc.MethodBindPtr
  fnSetStartPosition gdc.MethodBindPtr
  fnGetStartPosition gdc.MethodBindPtr
  fnSetEndPosition gdc.MethodBindPtr
  fnGetEndPosition gdc.MethodBindPtr
  fnSetGlobalStartPosition gdc.MethodBindPtr
  fnGetGlobalStartPosition gdc.MethodBindPtr
  fnSetGlobalEndPosition gdc.MethodBindPtr
  fnGetGlobalEndPosition gdc.MethodBindPtr
  fnSetEnterCost gdc.MethodBindPtr
  fnGetEnterCost gdc.MethodBindPtr
  fnSetTravelCost gdc.MethodBindPtr
  fnGetTravelCost gdc.MethodBindPtr
}

var ptrsForNavigationLink3D ptrsForNavigationLink3DList

func initNavigationLink3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("NavigationLink3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_enabled")
    defer methodName.Destroy()
    ptrsForNavigationLink3D.fnSetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_enabled")
    defer methodName.Destroy()
    ptrsForNavigationLink3D.fnIsEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_bidirectional")
    defer methodName.Destroy()
    ptrsForNavigationLink3D.fnSetBidirectional = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_bidirectional")
    defer methodName.Destroy()
    ptrsForNavigationLink3D.fnIsBidirectional = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_navigation_layers")
    defer methodName.Destroy()
    ptrsForNavigationLink3D.fnSetNavigationLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_navigation_layers")
    defer methodName.Destroy()
    ptrsForNavigationLink3D.fnGetNavigationLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_navigation_layer_value")
    defer methodName.Destroy()
    ptrsForNavigationLink3D.fnSetNavigationLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("get_navigation_layer_value")
    defer methodName.Destroy()
    ptrsForNavigationLink3D.fnGetNavigationLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("set_start_position")
    defer methodName.Destroy()
    ptrsForNavigationLink3D.fnSetStartPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_start_position")
    defer methodName.Destroy()
    ptrsForNavigationLink3D.fnGetStartPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("set_end_position")
    defer methodName.Destroy()
    ptrsForNavigationLink3D.fnSetEndPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_end_position")
    defer methodName.Destroy()
    ptrsForNavigationLink3D.fnGetEndPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("set_global_start_position")
    defer methodName.Destroy()
    ptrsForNavigationLink3D.fnSetGlobalStartPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_global_start_position")
    defer methodName.Destroy()
    ptrsForNavigationLink3D.fnGetGlobalStartPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("set_global_end_position")
    defer methodName.Destroy()
    ptrsForNavigationLink3D.fnSetGlobalEndPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_global_end_position")
    defer methodName.Destroy()
    ptrsForNavigationLink3D.fnGetGlobalEndPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("set_enter_cost")
    defer methodName.Destroy()
    ptrsForNavigationLink3D.fnSetEnterCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_enter_cost")
    defer methodName.Destroy()
    ptrsForNavigationLink3D.fnGetEnterCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_travel_cost")
    defer methodName.Destroy()
    ptrsForNavigationLink3D.fnSetTravelCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_travel_cost")
    defer methodName.Destroy()
    ptrsForNavigationLink3D.fnGetTravelCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
}

type NavigationLink3D struct {
  Node3D
}

func (me *NavigationLink3D) BaseClass() string {
  return "NavigationLink3D"
}

func NewNavigationLink3D() *NavigationLink3D {
  str := StringNameFromStr("NavigationLink3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &NavigationLink3D{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink3D.fnSetEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationLink3D) IsEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink3D.fnIsEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationLink3D) SetBidirectional(bidirectional bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bidirectional) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink3D.fnSetBidirectional), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationLink3D) IsBidirectional() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink3D.fnIsBidirectional), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationLink3D) SetNavigationLayers(navigation_layers int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&navigation_layers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink3D.fnSetNavigationLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationLink3D) GetNavigationLayers() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink3D.fnGetNavigationLayers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationLink3D) SetNavigationLayerValue(layer_number int64, value bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink3D.fnSetNavigationLayerValue), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationLink3D) GetNavigationLayerValue(layer_number int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&layer_number)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink3D.fnGetNavigationLayerValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationLink3D) SetStartPosition(position Vector3, )  {
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink3D.fnSetStartPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationLink3D) GetStartPosition() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink3D.fnGetStartPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationLink3D) SetEndPosition(position Vector3, )  {
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink3D.fnSetEndPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationLink3D) GetEndPosition() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink3D.fnGetEndPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationLink3D) SetGlobalStartPosition(position Vector3, )  {
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink3D.fnSetGlobalStartPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationLink3D) GetGlobalStartPosition() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink3D.fnGetGlobalStartPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationLink3D) SetGlobalEndPosition(position Vector3, )  {
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink3D.fnSetGlobalEndPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationLink3D) GetGlobalEndPosition() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink3D.fnGetGlobalEndPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationLink3D) SetEnterCost(enter_cost float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enter_cost) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink3D.fnSetEnterCost), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationLink3D) GetEnterCost() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink3D.fnGetEnterCost), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationLink3D) SetTravelCost(travel_cost float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&travel_cost) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink3D.fnSetTravelCost), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationLink3D) GetTravelCost() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink3D.fnGetTravelCost), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
