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

type ptrsForMeshConvexDecompositionSettingsList struct {
  fnSetMaxConcavity gdc.MethodBindPtr
  fnGetMaxConcavity gdc.MethodBindPtr
  fnSetSymmetryPlanesClippingBias gdc.MethodBindPtr
  fnGetSymmetryPlanesClippingBias gdc.MethodBindPtr
  fnSetRevolutionAxesClippingBias gdc.MethodBindPtr
  fnGetRevolutionAxesClippingBias gdc.MethodBindPtr
  fnSetMinVolumePerConvexHull gdc.MethodBindPtr
  fnGetMinVolumePerConvexHull gdc.MethodBindPtr
  fnSetResolution gdc.MethodBindPtr
  fnGetResolution gdc.MethodBindPtr
  fnSetMaxNumVerticesPerConvexHull gdc.MethodBindPtr
  fnGetMaxNumVerticesPerConvexHull gdc.MethodBindPtr
  fnSetPlaneDownsampling gdc.MethodBindPtr
  fnGetPlaneDownsampling gdc.MethodBindPtr
  fnSetConvexHullDownsampling gdc.MethodBindPtr
  fnGetConvexHullDownsampling gdc.MethodBindPtr
  fnSetNormalizeMesh gdc.MethodBindPtr
  fnGetNormalizeMesh gdc.MethodBindPtr
  fnSetMode gdc.MethodBindPtr
  fnGetMode gdc.MethodBindPtr
  fnSetConvexHullApproximation gdc.MethodBindPtr
  fnGetConvexHullApproximation gdc.MethodBindPtr
  fnSetMaxConvexHulls gdc.MethodBindPtr
  fnGetMaxConvexHulls gdc.MethodBindPtr
  fnSetProjectHullVertices gdc.MethodBindPtr
  fnGetProjectHullVertices gdc.MethodBindPtr
}

var ptrsForMeshConvexDecompositionSettings ptrsForMeshConvexDecompositionSettingsList

func initMeshConvexDecompositionSettingsPtrs(iface gdc.Interface) {

  className := StringNameFromStr("MeshConvexDecompositionSettings")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_max_concavity")
    defer methodName.Destroy()
    ptrsForMeshConvexDecompositionSettings.fnSetMaxConcavity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_max_concavity")
    defer methodName.Destroy()
    ptrsForMeshConvexDecompositionSettings.fnGetMaxConcavity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_symmetry_planes_clipping_bias")
    defer methodName.Destroy()
    ptrsForMeshConvexDecompositionSettings.fnSetSymmetryPlanesClippingBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_symmetry_planes_clipping_bias")
    defer methodName.Destroy()
    ptrsForMeshConvexDecompositionSettings.fnGetSymmetryPlanesClippingBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_revolution_axes_clipping_bias")
    defer methodName.Destroy()
    ptrsForMeshConvexDecompositionSettings.fnSetRevolutionAxesClippingBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_revolution_axes_clipping_bias")
    defer methodName.Destroy()
    ptrsForMeshConvexDecompositionSettings.fnGetRevolutionAxesClippingBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_min_volume_per_convex_hull")
    defer methodName.Destroy()
    ptrsForMeshConvexDecompositionSettings.fnSetMinVolumePerConvexHull = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_min_volume_per_convex_hull")
    defer methodName.Destroy()
    ptrsForMeshConvexDecompositionSettings.fnGetMinVolumePerConvexHull = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_resolution")
    defer methodName.Destroy()
    ptrsForMeshConvexDecompositionSettings.fnSetResolution = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_resolution")
    defer methodName.Destroy()
    ptrsForMeshConvexDecompositionSettings.fnGetResolution = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_max_num_vertices_per_convex_hull")
    defer methodName.Destroy()
    ptrsForMeshConvexDecompositionSettings.fnSetMaxNumVerticesPerConvexHull = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_max_num_vertices_per_convex_hull")
    defer methodName.Destroy()
    ptrsForMeshConvexDecompositionSettings.fnGetMaxNumVerticesPerConvexHull = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_plane_downsampling")
    defer methodName.Destroy()
    ptrsForMeshConvexDecompositionSettings.fnSetPlaneDownsampling = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_plane_downsampling")
    defer methodName.Destroy()
    ptrsForMeshConvexDecompositionSettings.fnGetPlaneDownsampling = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_convex_hull_downsampling")
    defer methodName.Destroy()
    ptrsForMeshConvexDecompositionSettings.fnSetConvexHullDownsampling = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_convex_hull_downsampling")
    defer methodName.Destroy()
    ptrsForMeshConvexDecompositionSettings.fnGetConvexHullDownsampling = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_normalize_mesh")
    defer methodName.Destroy()
    ptrsForMeshConvexDecompositionSettings.fnSetNormalizeMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_normalize_mesh")
    defer methodName.Destroy()
    ptrsForMeshConvexDecompositionSettings.fnGetNormalizeMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_mode")
    defer methodName.Destroy()
    ptrsForMeshConvexDecompositionSettings.fnSetMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1668072869))
  }
  {
    methodName := StringNameFromStr("get_mode")
    defer methodName.Destroy()
    ptrsForMeshConvexDecompositionSettings.fnGetMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 23479454))
  }
  {
    methodName := StringNameFromStr("set_convex_hull_approximation")
    defer methodName.Destroy()
    ptrsForMeshConvexDecompositionSettings.fnSetConvexHullApproximation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_convex_hull_approximation")
    defer methodName.Destroy()
    ptrsForMeshConvexDecompositionSettings.fnGetConvexHullApproximation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_max_convex_hulls")
    defer methodName.Destroy()
    ptrsForMeshConvexDecompositionSettings.fnSetMaxConvexHulls = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_max_convex_hulls")
    defer methodName.Destroy()
    ptrsForMeshConvexDecompositionSettings.fnGetMaxConvexHulls = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_project_hull_vertices")
    defer methodName.Destroy()
    ptrsForMeshConvexDecompositionSettings.fnSetProjectHullVertices = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_project_hull_vertices")
    defer methodName.Destroy()
    ptrsForMeshConvexDecompositionSettings.fnGetProjectHullVertices = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

type MeshConvexDecompositionSettings struct {
  RefCounted
}

func (me *MeshConvexDecompositionSettings) BaseClass() string {
  return "MeshConvexDecompositionSettings"
}

func NewMeshConvexDecompositionSettings() *MeshConvexDecompositionSettings {
  str := StringNameFromStr("MeshConvexDecompositionSettings") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &MeshConvexDecompositionSettings{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type MeshConvexDecompositionSettingsMode int
const (
  MeshConvexDecompositionSettingsModeConvexDecompositionModeVoxel MeshConvexDecompositionSettingsMode = 0
  MeshConvexDecompositionSettingsModeConvexDecompositionModeTetrahedron MeshConvexDecompositionSettingsMode = 1
)

func (me *MeshConvexDecompositionSettings) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MeshConvexDecompositionSettings) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MeshConvexDecompositionSettings) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *MeshConvexDecompositionSettings) SetMaxConcavity(max_concavity float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_concavity) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshConvexDecompositionSettings.fnSetMaxConcavity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshConvexDecompositionSettings) GetMaxConcavity() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshConvexDecompositionSettings.fnGetMaxConcavity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshConvexDecompositionSettings) SetSymmetryPlanesClippingBias(symmetry_planes_clipping_bias float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&symmetry_planes_clipping_bias) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshConvexDecompositionSettings.fnSetSymmetryPlanesClippingBias), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshConvexDecompositionSettings) GetSymmetryPlanesClippingBias() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshConvexDecompositionSettings.fnGetSymmetryPlanesClippingBias), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshConvexDecompositionSettings) SetRevolutionAxesClippingBias(revolution_axes_clipping_bias float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&revolution_axes_clipping_bias) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshConvexDecompositionSettings.fnSetRevolutionAxesClippingBias), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshConvexDecompositionSettings) GetRevolutionAxesClippingBias() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshConvexDecompositionSettings.fnGetRevolutionAxesClippingBias), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshConvexDecompositionSettings) SetMinVolumePerConvexHull(min_volume_per_convex_hull float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&min_volume_per_convex_hull) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshConvexDecompositionSettings.fnSetMinVolumePerConvexHull), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshConvexDecompositionSettings) GetMinVolumePerConvexHull() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshConvexDecompositionSettings.fnGetMinVolumePerConvexHull), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshConvexDecompositionSettings) SetResolution(min_volume_per_convex_hull int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&min_volume_per_convex_hull) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshConvexDecompositionSettings.fnSetResolution), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshConvexDecompositionSettings) GetResolution() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshConvexDecompositionSettings.fnGetResolution), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshConvexDecompositionSettings) SetMaxNumVerticesPerConvexHull(max_num_vertices_per_convex_hull int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_num_vertices_per_convex_hull) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshConvexDecompositionSettings.fnSetMaxNumVerticesPerConvexHull), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshConvexDecompositionSettings) GetMaxNumVerticesPerConvexHull() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshConvexDecompositionSettings.fnGetMaxNumVerticesPerConvexHull), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshConvexDecompositionSettings) SetPlaneDownsampling(plane_downsampling int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&plane_downsampling) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshConvexDecompositionSettings.fnSetPlaneDownsampling), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshConvexDecompositionSettings) GetPlaneDownsampling() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshConvexDecompositionSettings.fnGetPlaneDownsampling), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshConvexDecompositionSettings) SetConvexHullDownsampling(convex_hull_downsampling int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&convex_hull_downsampling) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshConvexDecompositionSettings.fnSetConvexHullDownsampling), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshConvexDecompositionSettings) GetConvexHullDownsampling() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshConvexDecompositionSettings.fnGetConvexHullDownsampling), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshConvexDecompositionSettings) SetNormalizeMesh(normalize_mesh bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&normalize_mesh) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshConvexDecompositionSettings.fnSetNormalizeMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshConvexDecompositionSettings) GetNormalizeMesh() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshConvexDecompositionSettings.fnGetNormalizeMesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshConvexDecompositionSettings) SetMode(mode MeshConvexDecompositionSettingsMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshConvexDecompositionSettings.fnSetMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshConvexDecompositionSettings) GetMode() MeshConvexDecompositionSettingsMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret MeshConvexDecompositionSettingsMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshConvexDecompositionSettings.fnGetMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *MeshConvexDecompositionSettings) SetConvexHullApproximation(convex_hull_approximation bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&convex_hull_approximation) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshConvexDecompositionSettings.fnSetConvexHullApproximation), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshConvexDecompositionSettings) GetConvexHullApproximation() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshConvexDecompositionSettings.fnGetConvexHullApproximation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshConvexDecompositionSettings) SetMaxConvexHulls(max_convex_hulls int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_convex_hulls) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshConvexDecompositionSettings.fnSetMaxConvexHulls), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshConvexDecompositionSettings) GetMaxConvexHulls() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshConvexDecompositionSettings.fnGetMaxConvexHulls), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshConvexDecompositionSettings) SetProjectHullVertices(project_hull_vertices bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&project_hull_vertices) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshConvexDecompositionSettings.fnSetProjectHullVertices), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshConvexDecompositionSettings) GetProjectHullVertices() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshConvexDecompositionSettings.fnGetProjectHullVertices), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
