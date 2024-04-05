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

type ptrsForLight2DList struct {
  fnSetEnabled gdc.MethodBindPtr
  fnIsEnabled gdc.MethodBindPtr
  fnSetEditorOnly gdc.MethodBindPtr
  fnIsEditorOnly gdc.MethodBindPtr
  fnSetColor gdc.MethodBindPtr
  fnGetColor gdc.MethodBindPtr
  fnSetEnergy gdc.MethodBindPtr
  fnGetEnergy gdc.MethodBindPtr
  fnSetZRangeMin gdc.MethodBindPtr
  fnGetZRangeMin gdc.MethodBindPtr
  fnSetZRangeMax gdc.MethodBindPtr
  fnGetZRangeMax gdc.MethodBindPtr
  fnSetLayerRangeMin gdc.MethodBindPtr
  fnGetLayerRangeMin gdc.MethodBindPtr
  fnSetLayerRangeMax gdc.MethodBindPtr
  fnGetLayerRangeMax gdc.MethodBindPtr
  fnSetItemCullMask gdc.MethodBindPtr
  fnGetItemCullMask gdc.MethodBindPtr
  fnSetItemShadowCullMask gdc.MethodBindPtr
  fnGetItemShadowCullMask gdc.MethodBindPtr
  fnSetShadowEnabled gdc.MethodBindPtr
  fnIsShadowEnabled gdc.MethodBindPtr
  fnSetShadowSmooth gdc.MethodBindPtr
  fnGetShadowSmooth gdc.MethodBindPtr
  fnSetShadowFilter gdc.MethodBindPtr
  fnGetShadowFilter gdc.MethodBindPtr
  fnSetShadowColor gdc.MethodBindPtr
  fnGetShadowColor gdc.MethodBindPtr
  fnSetBlendMode gdc.MethodBindPtr
  fnGetBlendMode gdc.MethodBindPtr
  fnSetHeight gdc.MethodBindPtr
  fnGetHeight gdc.MethodBindPtr
}

var ptrsForLight2D ptrsForLight2DList

func initLight2DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("Light2D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_enabled")
    defer methodName.Destroy()
    ptrsForLight2D.fnSetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_enabled")
    defer methodName.Destroy()
    ptrsForLight2D.fnIsEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_editor_only")
    defer methodName.Destroy()
    ptrsForLight2D.fnSetEditorOnly = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_editor_only")
    defer methodName.Destroy()
    ptrsForLight2D.fnIsEditorOnly = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_color")
    defer methodName.Destroy()
    ptrsForLight2D.fnSetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_color")
    defer methodName.Destroy()
    ptrsForLight2D.fnGetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_energy")
    defer methodName.Destroy()
    ptrsForLight2D.fnSetEnergy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_energy")
    defer methodName.Destroy()
    ptrsForLight2D.fnGetEnergy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_z_range_min")
    defer methodName.Destroy()
    ptrsForLight2D.fnSetZRangeMin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_z_range_min")
    defer methodName.Destroy()
    ptrsForLight2D.fnGetZRangeMin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_z_range_max")
    defer methodName.Destroy()
    ptrsForLight2D.fnSetZRangeMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_z_range_max")
    defer methodName.Destroy()
    ptrsForLight2D.fnGetZRangeMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_layer_range_min")
    defer methodName.Destroy()
    ptrsForLight2D.fnSetLayerRangeMin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_layer_range_min")
    defer methodName.Destroy()
    ptrsForLight2D.fnGetLayerRangeMin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_layer_range_max")
    defer methodName.Destroy()
    ptrsForLight2D.fnSetLayerRangeMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_layer_range_max")
    defer methodName.Destroy()
    ptrsForLight2D.fnGetLayerRangeMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_item_cull_mask")
    defer methodName.Destroy()
    ptrsForLight2D.fnSetItemCullMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_item_cull_mask")
    defer methodName.Destroy()
    ptrsForLight2D.fnGetItemCullMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_item_shadow_cull_mask")
    defer methodName.Destroy()
    ptrsForLight2D.fnSetItemShadowCullMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_item_shadow_cull_mask")
    defer methodName.Destroy()
    ptrsForLight2D.fnGetItemShadowCullMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_shadow_enabled")
    defer methodName.Destroy()
    ptrsForLight2D.fnSetShadowEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_shadow_enabled")
    defer methodName.Destroy()
    ptrsForLight2D.fnIsShadowEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_shadow_smooth")
    defer methodName.Destroy()
    ptrsForLight2D.fnSetShadowSmooth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_shadow_smooth")
    defer methodName.Destroy()
    ptrsForLight2D.fnGetShadowSmooth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_shadow_filter")
    defer methodName.Destroy()
    ptrsForLight2D.fnSetShadowFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3209356555))
  }
  {
    methodName := StringNameFromStr("get_shadow_filter")
    defer methodName.Destroy()
    ptrsForLight2D.fnGetShadowFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1973619177))
  }
  {
    methodName := StringNameFromStr("set_shadow_color")
    defer methodName.Destroy()
    ptrsForLight2D.fnSetShadowColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_shadow_color")
    defer methodName.Destroy()
    ptrsForLight2D.fnGetShadowColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_blend_mode")
    defer methodName.Destroy()
    ptrsForLight2D.fnSetBlendMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2916638796))
  }
  {
    methodName := StringNameFromStr("get_blend_mode")
    defer methodName.Destroy()
    ptrsForLight2D.fnGetBlendMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 936255250))
  }
  {
    methodName := StringNameFromStr("set_height")
    defer methodName.Destroy()
    ptrsForLight2D.fnSetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_height")
    defer methodName.Destroy()
    ptrsForLight2D.fnGetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
}

type Light2D struct {
  Node2D
}

func (me *Light2D) BaseClass() string {
  return "Light2D"
}

func NewLight2D() *Light2D {
  str := StringNameFromStr("Light2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Light2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type Light2DShadowFilter int
const (
  Light2DShadowFilterShadowFilterNone Light2DShadowFilter = 0
  Light2DShadowFilterShadowFilterPcf5 Light2DShadowFilter = 1
  Light2DShadowFilterShadowFilterPcf13 Light2DShadowFilter = 2
)

type Light2DBlendMode int
const (
  Light2DBlendModeBlendModeAdd Light2DBlendMode = 0
  Light2DBlendModeBlendModeSub Light2DBlendMode = 1
  Light2DBlendModeBlendModeMix Light2DBlendMode = 2
)

func (me *Light2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Light2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Light2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Light2D) SetEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnSetEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light2D) IsEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnIsEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light2D) SetEditorOnly(editor_only bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&editor_only) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnSetEditorOnly), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light2D) IsEditorOnly() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnIsEditorOnly), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light2D) SetColor(color Color, )  {
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnSetColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light2D) GetColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnGetColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Light2D) SetEnergy(energy float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&energy) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnSetEnergy), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light2D) GetEnergy() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnGetEnergy), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light2D) SetZRangeMin(z int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&z) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnSetZRangeMin), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light2D) GetZRangeMin() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnGetZRangeMin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light2D) SetZRangeMax(z int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&z) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnSetZRangeMax), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light2D) GetZRangeMax() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnGetZRangeMax), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light2D) SetLayerRangeMin(layer int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnSetLayerRangeMin), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light2D) GetLayerRangeMin() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnGetLayerRangeMin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light2D) SetLayerRangeMax(layer int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnSetLayerRangeMax), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light2D) GetLayerRangeMax() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnGetLayerRangeMax), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light2D) SetItemCullMask(item_cull_mask int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&item_cull_mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnSetItemCullMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light2D) GetItemCullMask() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnGetItemCullMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light2D) SetItemShadowCullMask(item_shadow_cull_mask int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&item_shadow_cull_mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnSetItemShadowCullMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light2D) GetItemShadowCullMask() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnGetItemShadowCullMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light2D) SetShadowEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnSetShadowEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light2D) IsShadowEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnIsShadowEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light2D) SetShadowSmooth(smooth float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&smooth) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnSetShadowSmooth), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light2D) GetShadowSmooth() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnGetShadowSmooth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Light2D) SetShadowFilter(filter Light2DShadowFilter, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&filter) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnSetShadowFilter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light2D) GetShadowFilter() Light2DShadowFilter {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Light2DShadowFilter

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnGetShadowFilter), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Light2D) SetShadowColor(shadow_color Color, )  {
  cargs := []gdc.ConstTypePtr{shadow_color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnSetShadowColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light2D) GetShadowColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnGetShadowColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Light2D) SetBlendMode(mode Light2DBlendMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnSetBlendMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light2D) GetBlendMode() Light2DBlendMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Light2DBlendMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnGetBlendMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Light2D) SetHeight(height float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnSetHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Light2D) GetHeight() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLight2D.fnGetHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
