// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type NavigationPathQueryParameters3D struct {
  RefCounted
}

func (me *NavigationPathQueryParameters3D) BaseClass() string {
  return "NavigationPathQueryParameters3D"
}



// Enums

type NavigationPathQueryParameters3DPathfindingAlgorithm int
const (
  NavigationPathQueryParameters3DPathfindingAlgorithmPathfindingAlgorithmAstar NavigationPathQueryParameters3DPathfindingAlgorithm = 0
)

type NavigationPathQueryParameters3DPathPostProcessing int
const (
  NavigationPathQueryParameters3DPathPostProcessingPathPostprocessingCorridorfunnel NavigationPathQueryParameters3DPathPostProcessing = 0
  NavigationPathQueryParameters3DPathPostProcessingPathPostprocessingEdgecentered NavigationPathQueryParameters3DPathPostProcessing = 1
)

type NavigationPathQueryParameters3DPathMetadataFlags int
const (
  NavigationPathQueryParameters3DPathMetadataFlagsPathMetadataIncludeNone NavigationPathQueryParameters3DPathMetadataFlags = 0
  NavigationPathQueryParameters3DPathMetadataFlagsPathMetadataIncludeTypes NavigationPathQueryParameters3DPathMetadataFlags = 1
  NavigationPathQueryParameters3DPathMetadataFlagsPathMetadataIncludeRids NavigationPathQueryParameters3DPathMetadataFlags = 2
  NavigationPathQueryParameters3DPathMetadataFlagsPathMetadataIncludeOwners NavigationPathQueryParameters3DPathMetadataFlags = 4
  NavigationPathQueryParameters3DPathMetadataFlagsPathMetadataIncludeAll NavigationPathQueryParameters3DPathMetadataFlags = 7
)

func (me *NavigationPathQueryParameters3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *NavigationPathQueryParameters3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NavigationPathQueryParameters3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *NavigationPathQueryParameters3D) SetPathfindingAlgorithm(pathfinding_algorithm NavigationPathQueryParameters3DPathfindingAlgorithm, )  {
  classNameV := StringNameFromStr("NavigationPathQueryParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pathfinding_algorithm")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 394560454) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pathfinding_algorithm), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPathQueryParameters3D) GetPathfindingAlgorithm() NavigationPathQueryParameters3DPathfindingAlgorithm {
  classNameV := StringNameFromStr("NavigationPathQueryParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pathfinding_algorithm")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3398491350) // FIXME: should cache?
  var ret NavigationPathQueryParameters3DPathfindingAlgorithm
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationPathQueryParameters3D) SetPathPostprocessing(path_postprocessing NavigationPathQueryParameters3DPathPostProcessing, )  {
  classNameV := StringNameFromStr("NavigationPathQueryParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path_postprocessing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2267362344) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&path_postprocessing), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPathQueryParameters3D) GetPathPostprocessing() NavigationPathQueryParameters3DPathPostProcessing {
  classNameV := StringNameFromStr("NavigationPathQueryParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path_postprocessing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3883858360) // FIXME: should cache?
  var ret NavigationPathQueryParameters3DPathPostProcessing
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationPathQueryParameters3D) SetMap(map_ RID, )  {
  classNameV := StringNameFromStr("NavigationPathQueryParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(map_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPathQueryParameters3D) GetMap() RID {
  classNameV := StringNameFromStr("NavigationPathQueryParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationPathQueryParameters3D) SetStartPosition(start_position Vector3, )  {
  classNameV := StringNameFromStr("NavigationPathQueryParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_start_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(start_position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPathQueryParameters3D) GetStartPosition() Vector3 {
  classNameV := StringNameFromStr("NavigationPathQueryParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_start_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationPathQueryParameters3D) SetTargetPosition(target_position Vector3, )  {
  classNameV := StringNameFromStr("NavigationPathQueryParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_target_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(target_position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPathQueryParameters3D) GetTargetPosition() Vector3 {
  classNameV := StringNameFromStr("NavigationPathQueryParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_target_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationPathQueryParameters3D) SetNavigationLayers(navigation_layers int, )  {
  classNameV := StringNameFromStr("NavigationPathQueryParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_navigation_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&navigation_layers), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPathQueryParameters3D) GetNavigationLayers() int {
  classNameV := StringNameFromStr("NavigationPathQueryParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationPathQueryParameters3D) SetMetadataFlags(flags NavigationPathQueryParameters3DPathMetadataFlags, )  {
  classNameV := StringNameFromStr("NavigationPathQueryParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_metadata_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2713846708) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPathQueryParameters3D) GetMetadataFlags() NavigationPathQueryParameters3DPathMetadataFlags {
  classNameV := StringNameFromStr("NavigationPathQueryParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_metadata_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1582332802) // FIXME: should cache?
  var ret NavigationPathQueryParameters3DPathMetadataFlags
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
