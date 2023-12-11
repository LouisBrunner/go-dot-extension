// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type NavigationPathQueryParameters2D struct {
  obj gdc.ObjectPtr
}

func (me *NavigationPathQueryParameters2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationPathQueryParameters2D) BaseClass() string {
  return "NavigationPathQueryParameters2D"
}



// Enums

type NavigationPathQueryParameters2DPathfindingAlgorithm int
const (
  NavigationPathQueryParameters2DPathfindingAlgorithmPathfindingAlgorithmAstar NavigationPathQueryParameters2DPathfindingAlgorithm = 0
)

type NavigationPathQueryParameters2DPathPostProcessing int
const (
  NavigationPathQueryParameters2DPathPostProcessingPathPostprocessingCorridorfunnel NavigationPathQueryParameters2DPathPostProcessing = 0
  NavigationPathQueryParameters2DPathPostProcessingPathPostprocessingEdgecentered NavigationPathQueryParameters2DPathPostProcessing = 1
)

type NavigationPathQueryParameters2DPathMetadataFlags int
const (
  NavigationPathQueryParameters2DPathMetadataFlagsPathMetadataIncludeNone NavigationPathQueryParameters2DPathMetadataFlags = 0
  NavigationPathQueryParameters2DPathMetadataFlagsPathMetadataIncludeTypes NavigationPathQueryParameters2DPathMetadataFlags = 1
  NavigationPathQueryParameters2DPathMetadataFlagsPathMetadataIncludeRids NavigationPathQueryParameters2DPathMetadataFlags = 2
  NavigationPathQueryParameters2DPathMetadataFlagsPathMetadataIncludeOwners NavigationPathQueryParameters2DPathMetadataFlags = 4
  NavigationPathQueryParameters2DPathMetadataFlagsPathMetadataIncludeAll NavigationPathQueryParameters2DPathMetadataFlags = 7
)

func (me *NavigationPathQueryParameters2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *NavigationPathQueryParameters2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NavigationPathQueryParameters2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *NavigationPathQueryParameters2D) SetPathfindingAlgorithm(pathfinding_algorithm NavigationPathQueryParameters2DPathfindingAlgorithm, )  {
  classNameV := StringNameFromStr("NavigationPathQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pathfinding_algorithm")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2783519915) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pathfinding_algorithm), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPathQueryParameters2D) GetPathfindingAlgorithm() NavigationPathQueryParameters2DPathfindingAlgorithm {
  classNameV := StringNameFromStr("NavigationPathQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pathfinding_algorithm")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3000421146) // FIXME: should cache?
  var ret NavigationPathQueryParameters2DPathfindingAlgorithm
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationPathQueryParameters2D) SetPathPostprocessing(path_postprocessing NavigationPathQueryParameters2DPathPostProcessing, )  {
  classNameV := StringNameFromStr("NavigationPathQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path_postprocessing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2864409082) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&path_postprocessing), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPathQueryParameters2D) GetPathPostprocessing() NavigationPathQueryParameters2DPathPostProcessing {
  classNameV := StringNameFromStr("NavigationPathQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path_postprocessing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3798118993) // FIXME: should cache?
  var ret NavigationPathQueryParameters2DPathPostProcessing
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationPathQueryParameters2D) SetMap(map_ RID, )  {
  classNameV := StringNameFromStr("NavigationPathQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(map_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPathQueryParameters2D) GetMap() RID {
  classNameV := StringNameFromStr("NavigationPathQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationPathQueryParameters2D) SetStartPosition(start_position Vector2, )  {
  classNameV := StringNameFromStr("NavigationPathQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_start_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(start_position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPathQueryParameters2D) GetStartPosition() Vector2 {
  classNameV := StringNameFromStr("NavigationPathQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_start_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationPathQueryParameters2D) SetTargetPosition(target_position Vector2, )  {
  classNameV := StringNameFromStr("NavigationPathQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_target_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(target_position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPathQueryParameters2D) GetTargetPosition() Vector2 {
  classNameV := StringNameFromStr("NavigationPathQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_target_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationPathQueryParameters2D) SetNavigationLayers(navigation_layers int, )  {
  classNameV := StringNameFromStr("NavigationPathQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_navigation_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&navigation_layers), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPathQueryParameters2D) GetNavigationLayers() int {
  classNameV := StringNameFromStr("NavigationPathQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationPathQueryParameters2D) SetMetadataFlags(flags NavigationPathQueryParameters2DPathMetadataFlags, )  {
  classNameV := StringNameFromStr("NavigationPathQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_metadata_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 24274129) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPathQueryParameters2D) GetMetadataFlags() NavigationPathQueryParameters2DPathMetadataFlags {
  classNameV := StringNameFromStr("NavigationPathQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_metadata_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 488152976) // FIXME: should cache?
  var ret NavigationPathQueryParameters2DPathMetadataFlags
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
