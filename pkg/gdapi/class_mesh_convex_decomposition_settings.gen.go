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
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_concavity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_concavity) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshConvexDecompositionSettings) GetMaxConcavity() float64 {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_concavity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshConvexDecompositionSettings) SetSymmetryPlanesClippingBias(symmetry_planes_clipping_bias float64, )  {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_symmetry_planes_clipping_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&symmetry_planes_clipping_bias) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshConvexDecompositionSettings) GetSymmetryPlanesClippingBias() float64 {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_symmetry_planes_clipping_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshConvexDecompositionSettings) SetRevolutionAxesClippingBias(revolution_axes_clipping_bias float64, )  {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_revolution_axes_clipping_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&revolution_axes_clipping_bias) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshConvexDecompositionSettings) GetRevolutionAxesClippingBias() float64 {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_revolution_axes_clipping_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshConvexDecompositionSettings) SetMinVolumePerConvexHull(min_volume_per_convex_hull float64, )  {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_min_volume_per_convex_hull")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&min_volume_per_convex_hull) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshConvexDecompositionSettings) GetMinVolumePerConvexHull() float64 {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_min_volume_per_convex_hull")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshConvexDecompositionSettings) SetResolution(min_volume_per_convex_hull int64, )  {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_resolution")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&min_volume_per_convex_hull) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshConvexDecompositionSettings) GetResolution() int64 {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_resolution")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshConvexDecompositionSettings) SetMaxNumVerticesPerConvexHull(max_num_vertices_per_convex_hull int64, )  {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_num_vertices_per_convex_hull")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_num_vertices_per_convex_hull) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshConvexDecompositionSettings) GetMaxNumVerticesPerConvexHull() int64 {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_num_vertices_per_convex_hull")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshConvexDecompositionSettings) SetPlaneDownsampling(plane_downsampling int64, )  {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_plane_downsampling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&plane_downsampling) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshConvexDecompositionSettings) GetPlaneDownsampling() int64 {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_plane_downsampling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshConvexDecompositionSettings) SetConvexHullDownsampling(convex_hull_downsampling int64, )  {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_convex_hull_downsampling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&convex_hull_downsampling) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshConvexDecompositionSettings) GetConvexHullDownsampling() int64 {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_convex_hull_downsampling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshConvexDecompositionSettings) SetNormalizeMesh(normalize_mesh bool, )  {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_normalize_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&normalize_mesh) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshConvexDecompositionSettings) GetNormalizeMesh() bool {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_normalize_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshConvexDecompositionSettings) SetMode(mode MeshConvexDecompositionSettingsMode, )  {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1668072869) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshConvexDecompositionSettings) GetMode() MeshConvexDecompositionSettingsMode {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 23479454) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret MeshConvexDecompositionSettingsMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *MeshConvexDecompositionSettings) SetConvexHullApproximation(convex_hull_approximation bool, )  {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_convex_hull_approximation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&convex_hull_approximation) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshConvexDecompositionSettings) GetConvexHullApproximation() bool {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_convex_hull_approximation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshConvexDecompositionSettings) SetMaxConvexHulls(max_convex_hulls int64, )  {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_convex_hulls")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_convex_hulls) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshConvexDecompositionSettings) GetMaxConvexHulls() int64 {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_convex_hulls")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshConvexDecompositionSettings) SetProjectHullVertices(project_hull_vertices bool, )  {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_project_hull_vertices")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&project_hull_vertices) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshConvexDecompositionSettings) GetProjectHullVertices() bool {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_project_hull_vertices")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
