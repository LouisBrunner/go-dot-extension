// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func  (me *MeshConvexDecompositionSettings) SetMaxConcavity(max_concavity float32, )  {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_concavity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_concavity), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MeshConvexDecompositionSettings) GetMaxConcavity() float32 {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_concavity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshConvexDecompositionSettings) SetSymmetryPlanesClippingBias(symmetry_planes_clipping_bias float32, )  {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_symmetry_planes_clipping_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&symmetry_planes_clipping_bias), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MeshConvexDecompositionSettings) GetSymmetryPlanesClippingBias() float32 {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_symmetry_planes_clipping_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshConvexDecompositionSettings) SetRevolutionAxesClippingBias(revolution_axes_clipping_bias float32, )  {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_revolution_axes_clipping_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&revolution_axes_clipping_bias), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MeshConvexDecompositionSettings) GetRevolutionAxesClippingBias() float32 {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_revolution_axes_clipping_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshConvexDecompositionSettings) SetMinVolumePerConvexHull(min_volume_per_convex_hull float32, )  {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_min_volume_per_convex_hull")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&min_volume_per_convex_hull), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MeshConvexDecompositionSettings) GetMinVolumePerConvexHull() float32 {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_min_volume_per_convex_hull")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshConvexDecompositionSettings) SetResolution(min_volume_per_convex_hull int, )  {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_resolution")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&min_volume_per_convex_hull), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MeshConvexDecompositionSettings) GetResolution() int {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_resolution")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshConvexDecompositionSettings) SetMaxNumVerticesPerConvexHull(max_num_vertices_per_convex_hull int, )  {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_num_vertices_per_convex_hull")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_num_vertices_per_convex_hull), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MeshConvexDecompositionSettings) GetMaxNumVerticesPerConvexHull() int {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_num_vertices_per_convex_hull")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshConvexDecompositionSettings) SetPlaneDownsampling(plane_downsampling int, )  {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_plane_downsampling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&plane_downsampling), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MeshConvexDecompositionSettings) GetPlaneDownsampling() int {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_plane_downsampling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshConvexDecompositionSettings) SetConvexHullDownsampling(convex_hull_downsampling int, )  {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_convex_hull_downsampling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&convex_hull_downsampling), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MeshConvexDecompositionSettings) GetConvexHullDownsampling() int {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_convex_hull_downsampling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshConvexDecompositionSettings) SetNormalizeMesh(normalize_mesh bool, )  {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_normalize_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&normalize_mesh), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MeshConvexDecompositionSettings) GetNormalizeMesh() bool {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_normalize_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshConvexDecompositionSettings) SetMode(mode MeshConvexDecompositionSettingsMode, )  {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1668072869) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MeshConvexDecompositionSettings) GetMode() MeshConvexDecompositionSettingsMode {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 23479454) // FIXME: should cache?
  var ret MeshConvexDecompositionSettingsMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshConvexDecompositionSettings) SetConvexHullApproximation(convex_hull_approximation bool, )  {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_convex_hull_approximation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&convex_hull_approximation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MeshConvexDecompositionSettings) GetConvexHullApproximation() bool {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_convex_hull_approximation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshConvexDecompositionSettings) SetMaxConvexHulls(max_convex_hulls int, )  {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_convex_hulls")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_convex_hulls), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MeshConvexDecompositionSettings) GetMaxConvexHulls() int {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_convex_hulls")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshConvexDecompositionSettings) SetProjectHullVertices(project_hull_vertices bool, )  {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_project_hull_vertices")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&project_hull_vertices), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MeshConvexDecompositionSettings) GetProjectHullVertices() bool {
  classNameV := StringNameFromStr("MeshConvexDecompositionSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_project_hull_vertices")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *MeshConvexDecompositionSettings) GetPropMaxConcavity() float32 {
  panic("TODO: implement")
}

func (me *MeshConvexDecompositionSettings) SetPropMaxConcavity(value float32) {
  panic("TODO: implement")
}

func (me *MeshConvexDecompositionSettings) GetPropSymmetryPlanesClippingBias() float32 {
  panic("TODO: implement")
}

func (me *MeshConvexDecompositionSettings) SetPropSymmetryPlanesClippingBias(value float32) {
  panic("TODO: implement")
}

func (me *MeshConvexDecompositionSettings) GetPropRevolutionAxesClippingBias() float32 {
  panic("TODO: implement")
}

func (me *MeshConvexDecompositionSettings) SetPropRevolutionAxesClippingBias(value float32) {
  panic("TODO: implement")
}

func (me *MeshConvexDecompositionSettings) GetPropMinVolumePerConvexHull() float32 {
  panic("TODO: implement")
}

func (me *MeshConvexDecompositionSettings) SetPropMinVolumePerConvexHull(value float32) {
  panic("TODO: implement")
}

func (me *MeshConvexDecompositionSettings) GetPropResolution() int {
  panic("TODO: implement")
}

func (me *MeshConvexDecompositionSettings) SetPropResolution(value int) {
  panic("TODO: implement")
}

func (me *MeshConvexDecompositionSettings) GetPropMaxNumVerticesPerConvexHull() int {
  panic("TODO: implement")
}

func (me *MeshConvexDecompositionSettings) SetPropMaxNumVerticesPerConvexHull(value int) {
  panic("TODO: implement")
}

func (me *MeshConvexDecompositionSettings) GetPropPlaneDownsampling() int {
  panic("TODO: implement")
}

func (me *MeshConvexDecompositionSettings) SetPropPlaneDownsampling(value int) {
  panic("TODO: implement")
}

func (me *MeshConvexDecompositionSettings) GetPropConvexHullDownsampling() int {
  panic("TODO: implement")
}

func (me *MeshConvexDecompositionSettings) SetPropConvexHullDownsampling(value int) {
  panic("TODO: implement")
}

func (me *MeshConvexDecompositionSettings) GetPropNormalizeMesh() bool {
  panic("TODO: implement")
}

func (me *MeshConvexDecompositionSettings) SetPropNormalizeMesh(value bool) {
  panic("TODO: implement")
}

func (me *MeshConvexDecompositionSettings) GetPropMode() int {
  panic("TODO: implement")
}

func (me *MeshConvexDecompositionSettings) SetPropMode(value int) {
  panic("TODO: implement")
}

func (me *MeshConvexDecompositionSettings) GetPropConvexHullApproximation() bool {
  panic("TODO: implement")
}

func (me *MeshConvexDecompositionSettings) SetPropConvexHullApproximation(value bool) {
  panic("TODO: implement")
}

func (me *MeshConvexDecompositionSettings) GetPropMaxConvexHulls() int {
  panic("TODO: implement")
}

func (me *MeshConvexDecompositionSettings) SetPropMaxConvexHulls(value int) {
  panic("TODO: implement")
}

func (me *MeshConvexDecompositionSettings) GetPropProjectHullVertices() bool {
  panic("TODO: implement")
}

func (me *MeshConvexDecompositionSettings) SetPropProjectHullVertices(value bool) {
  panic("TODO: implement")
}