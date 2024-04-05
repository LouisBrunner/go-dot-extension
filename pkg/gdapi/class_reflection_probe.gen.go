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

type ptrsForReflectionProbeList struct {
  fnSetIntensity gdc.MethodBindPtr
  fnGetIntensity gdc.MethodBindPtr
  fnSetAmbientMode gdc.MethodBindPtr
  fnGetAmbientMode gdc.MethodBindPtr
  fnSetAmbientColor gdc.MethodBindPtr
  fnGetAmbientColor gdc.MethodBindPtr
  fnSetAmbientColorEnergy gdc.MethodBindPtr
  fnGetAmbientColorEnergy gdc.MethodBindPtr
  fnSetMaxDistance gdc.MethodBindPtr
  fnGetMaxDistance gdc.MethodBindPtr
  fnSetMeshLodThreshold gdc.MethodBindPtr
  fnGetMeshLodThreshold gdc.MethodBindPtr
  fnSetSize gdc.MethodBindPtr
  fnGetSize gdc.MethodBindPtr
  fnSetOriginOffset gdc.MethodBindPtr
  fnGetOriginOffset gdc.MethodBindPtr
  fnSetAsInterior gdc.MethodBindPtr
  fnIsSetAsInterior gdc.MethodBindPtr
  fnSetEnableBoxProjection gdc.MethodBindPtr
  fnIsBoxProjectionEnabled gdc.MethodBindPtr
  fnSetEnableShadows gdc.MethodBindPtr
  fnAreShadowsEnabled gdc.MethodBindPtr
  fnSetCullMask gdc.MethodBindPtr
  fnGetCullMask gdc.MethodBindPtr
  fnSetUpdateMode gdc.MethodBindPtr
  fnGetUpdateMode gdc.MethodBindPtr
}

var ptrsForReflectionProbe ptrsForReflectionProbeList

func initReflectionProbePtrs(iface gdc.Interface) {

  className := StringNameFromStr("ReflectionProbe")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_intensity")
    defer methodName.Destroy()
    ptrsForReflectionProbe.fnSetIntensity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_intensity")
    defer methodName.Destroy()
    ptrsForReflectionProbe.fnGetIntensity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_ambient_mode")
    defer methodName.Destroy()
    ptrsForReflectionProbe.fnSetAmbientMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1748981278))
  }
  {
    methodName := StringNameFromStr("get_ambient_mode")
    defer methodName.Destroy()
    ptrsForReflectionProbe.fnGetAmbientMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1014607621))
  }
  {
    methodName := StringNameFromStr("set_ambient_color")
    defer methodName.Destroy()
    ptrsForReflectionProbe.fnSetAmbientColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_ambient_color")
    defer methodName.Destroy()
    ptrsForReflectionProbe.fnGetAmbientColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_ambient_color_energy")
    defer methodName.Destroy()
    ptrsForReflectionProbe.fnSetAmbientColorEnergy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_ambient_color_energy")
    defer methodName.Destroy()
    ptrsForReflectionProbe.fnGetAmbientColorEnergy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_max_distance")
    defer methodName.Destroy()
    ptrsForReflectionProbe.fnSetMaxDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_max_distance")
    defer methodName.Destroy()
    ptrsForReflectionProbe.fnGetMaxDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_mesh_lod_threshold")
    defer methodName.Destroy()
    ptrsForReflectionProbe.fnSetMeshLodThreshold = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_mesh_lod_threshold")
    defer methodName.Destroy()
    ptrsForReflectionProbe.fnGetMeshLodThreshold = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_size")
    defer methodName.Destroy()
    ptrsForReflectionProbe.fnSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_size")
    defer methodName.Destroy()
    ptrsForReflectionProbe.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("set_origin_offset")
    defer methodName.Destroy()
    ptrsForReflectionProbe.fnSetOriginOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_origin_offset")
    defer methodName.Destroy()
    ptrsForReflectionProbe.fnGetOriginOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("set_as_interior")
    defer methodName.Destroy()
    ptrsForReflectionProbe.fnSetAsInterior = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_set_as_interior")
    defer methodName.Destroy()
    ptrsForReflectionProbe.fnIsSetAsInterior = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_enable_box_projection")
    defer methodName.Destroy()
    ptrsForReflectionProbe.fnSetEnableBoxProjection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_box_projection_enabled")
    defer methodName.Destroy()
    ptrsForReflectionProbe.fnIsBoxProjectionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_enable_shadows")
    defer methodName.Destroy()
    ptrsForReflectionProbe.fnSetEnableShadows = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("are_shadows_enabled")
    defer methodName.Destroy()
    ptrsForReflectionProbe.fnAreShadowsEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_cull_mask")
    defer methodName.Destroy()
    ptrsForReflectionProbe.fnSetCullMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_cull_mask")
    defer methodName.Destroy()
    ptrsForReflectionProbe.fnGetCullMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_update_mode")
    defer methodName.Destroy()
    ptrsForReflectionProbe.fnSetUpdateMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4090221187))
  }
  {
    methodName := StringNameFromStr("get_update_mode")
    defer methodName.Destroy()
    ptrsForReflectionProbe.fnGetUpdateMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2367550552))
  }
}

type ReflectionProbe struct {
  VisualInstance3D
}

func (me *ReflectionProbe) BaseClass() string {
  return "ReflectionProbe"
}

func NewReflectionProbe() *ReflectionProbe {
  str := StringNameFromStr("ReflectionProbe") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ReflectionProbe{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type ReflectionProbeUpdateMode int
const (
  ReflectionProbeUpdateModeUpdateOnce ReflectionProbeUpdateMode = 0
  ReflectionProbeUpdateModeUpdateAlways ReflectionProbeUpdateMode = 1
)

type ReflectionProbeAmbientMode int
const (
  ReflectionProbeAmbientModeAmbientDisabled ReflectionProbeAmbientMode = 0
  ReflectionProbeAmbientModeAmbientEnvironment ReflectionProbeAmbientMode = 1
  ReflectionProbeAmbientModeAmbientColor ReflectionProbeAmbientMode = 2
)

func (me *ReflectionProbe) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ReflectionProbe) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ReflectionProbe) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ReflectionProbe) SetIntensity(intensity float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&intensity) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReflectionProbe.fnSetIntensity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ReflectionProbe) GetIntensity() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReflectionProbe.fnGetIntensity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ReflectionProbe) SetAmbientMode(ambient ReflectionProbeAmbientMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ambient) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReflectionProbe.fnSetAmbientMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ReflectionProbe) GetAmbientMode() ReflectionProbeAmbientMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ReflectionProbeAmbientMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReflectionProbe.fnGetAmbientMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ReflectionProbe) SetAmbientColor(ambient Color, )  {
  cargs := []gdc.ConstTypePtr{ambient.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReflectionProbe.fnSetAmbientColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ReflectionProbe) GetAmbientColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReflectionProbe.fnGetAmbientColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ReflectionProbe) SetAmbientColorEnergy(ambient_energy float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ambient_energy) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReflectionProbe.fnSetAmbientColorEnergy), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ReflectionProbe) GetAmbientColorEnergy() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReflectionProbe.fnGetAmbientColorEnergy), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ReflectionProbe) SetMaxDistance(max_distance float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReflectionProbe.fnSetMaxDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ReflectionProbe) GetMaxDistance() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReflectionProbe.fnGetMaxDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ReflectionProbe) SetMeshLodThreshold(ratio float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReflectionProbe.fnSetMeshLodThreshold), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ReflectionProbe) GetMeshLodThreshold() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReflectionProbe.fnGetMeshLodThreshold), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ReflectionProbe) SetSize(size Vector3, )  {
  cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReflectionProbe.fnSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ReflectionProbe) GetSize() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReflectionProbe.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ReflectionProbe) SetOriginOffset(origin_offset Vector3, )  {
  cargs := []gdc.ConstTypePtr{origin_offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReflectionProbe.fnSetOriginOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ReflectionProbe) GetOriginOffset() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReflectionProbe.fnGetOriginOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ReflectionProbe) SetAsInterior(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReflectionProbe.fnSetAsInterior), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ReflectionProbe) IsSetAsInterior() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReflectionProbe.fnIsSetAsInterior), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ReflectionProbe) SetEnableBoxProjection(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReflectionProbe.fnSetEnableBoxProjection), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ReflectionProbe) IsBoxProjectionEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReflectionProbe.fnIsBoxProjectionEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ReflectionProbe) SetEnableShadows(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReflectionProbe.fnSetEnableShadows), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ReflectionProbe) AreShadowsEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReflectionProbe.fnAreShadowsEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ReflectionProbe) SetCullMask(layers int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReflectionProbe.fnSetCullMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ReflectionProbe) GetCullMask() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReflectionProbe.fnGetCullMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ReflectionProbe) SetUpdateMode(mode ReflectionProbeUpdateMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReflectionProbe.fnSetUpdateMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ReflectionProbe) GetUpdateMode() ReflectionProbeUpdateMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ReflectionProbeUpdateMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReflectionProbe.fnGetUpdateMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
