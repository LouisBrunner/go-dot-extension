// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MeshConvexDecompositionSettings struct {
  obj gdc.ObjectPtr
}

func (me *MeshConvexDecompositionSettings) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MeshConvexDecompositionSettings) BaseClass() string {
  return "MeshConvexDecompositionSettings"
}



// Enums

type MeshConvexDecompositionSettingsMode int
const (
  MeshConvexDecompositionSettingsModeConvexDecompositionModeVoxel MeshConvexDecompositionSettingsMode = 0
  MeshConvexDecompositionSettingsModeConvexDecompositionModeTetrahedron MeshConvexDecompositionSettingsMode = 1
)

func (me *MeshConvexDecompositionSettings) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MeshConvexDecompositionSettings) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *MeshConvexDecompositionSettings) SetMaxConcavity(max_concavity float32, )  {
  panic("TODO: implement")
}

func  (me *MeshConvexDecompositionSettings) GetMaxConcavity()  {
  panic("TODO: implement")
}

func  (me *MeshConvexDecompositionSettings) SetSymmetryPlanesClippingBias(symmetry_planes_clipping_bias float32, )  {
  panic("TODO: implement")
}

func  (me *MeshConvexDecompositionSettings) GetSymmetryPlanesClippingBias()  {
  panic("TODO: implement")
}

func  (me *MeshConvexDecompositionSettings) SetRevolutionAxesClippingBias(revolution_axes_clipping_bias float32, )  {
  panic("TODO: implement")
}

func  (me *MeshConvexDecompositionSettings) GetRevolutionAxesClippingBias()  {
  panic("TODO: implement")
}

func  (me *MeshConvexDecompositionSettings) SetMinVolumePerConvexHull(min_volume_per_convex_hull float32, )  {
  panic("TODO: implement")
}

func  (me *MeshConvexDecompositionSettings) GetMinVolumePerConvexHull()  {
  panic("TODO: implement")
}

func  (me *MeshConvexDecompositionSettings) SetResolution(min_volume_per_convex_hull int, )  {
  panic("TODO: implement")
}

func  (me *MeshConvexDecompositionSettings) GetResolution()  {
  panic("TODO: implement")
}

func  (me *MeshConvexDecompositionSettings) SetMaxNumVerticesPerConvexHull(max_num_vertices_per_convex_hull int, )  {
  panic("TODO: implement")
}

func  (me *MeshConvexDecompositionSettings) GetMaxNumVerticesPerConvexHull()  {
  panic("TODO: implement")
}

func  (me *MeshConvexDecompositionSettings) SetPlaneDownsampling(plane_downsampling int, )  {
  panic("TODO: implement")
}

func  (me *MeshConvexDecompositionSettings) GetPlaneDownsampling()  {
  panic("TODO: implement")
}

func  (me *MeshConvexDecompositionSettings) SetConvexHullDownsampling(convex_hull_downsampling int, )  {
  panic("TODO: implement")
}

func  (me *MeshConvexDecompositionSettings) GetConvexHullDownsampling()  {
  panic("TODO: implement")
}

func  (me *MeshConvexDecompositionSettings) SetNormalizeMesh(normalize_mesh bool, )  {
  panic("TODO: implement")
}

func  (me *MeshConvexDecompositionSettings) GetNormalizeMesh()  {
  panic("TODO: implement")
}

func  (me *MeshConvexDecompositionSettings) SetMode(mode MeshConvexDecompositionSettingsMode, )  {
  panic("TODO: implement")
}

func  (me *MeshConvexDecompositionSettings) GetMode()  {
  panic("TODO: implement")
}

func  (me *MeshConvexDecompositionSettings) SetConvexHullApproximation(convex_hull_approximation bool, )  {
  panic("TODO: implement")
}

func  (me *MeshConvexDecompositionSettings) GetConvexHullApproximation()  {
  panic("TODO: implement")
}

func  (me *MeshConvexDecompositionSettings) SetMaxConvexHulls(max_convex_hulls int, )  {
  panic("TODO: implement")
}

func  (me *MeshConvexDecompositionSettings) GetMaxConvexHulls()  {
  panic("TODO: implement")
}

func  (me *MeshConvexDecompositionSettings) SetProjectHullVertices(project_hull_vertices bool, )  {
  panic("TODO: implement")
}

func  (me *MeshConvexDecompositionSettings) GetProjectHullVertices()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
