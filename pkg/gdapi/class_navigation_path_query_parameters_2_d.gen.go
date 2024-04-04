// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type NavigationPathQueryParameters2D struct {
  RefCounted
}

func (me *NavigationPathQueryParameters2D) BaseClass() string {
  return "NavigationPathQueryParameters2D"
}

func NewNavigationPathQueryParameters2D() *NavigationPathQueryParameters2D {
  str := StringNameFromStr("NavigationPathQueryParameters2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &NavigationPathQueryParameters2D{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pathfinding_algorithm) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPathQueryParameters2D) GetPathfindingAlgorithm() NavigationPathQueryParameters2DPathfindingAlgorithm {
  classNameV := StringNameFromStr("NavigationPathQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pathfinding_algorithm")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3000421146) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret NavigationPathQueryParameters2DPathfindingAlgorithm

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *NavigationPathQueryParameters2D) SetPathPostprocessing(path_postprocessing NavigationPathQueryParameters2DPathPostProcessing, )  {
  classNameV := StringNameFromStr("NavigationPathQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path_postprocessing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2864409082) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&path_postprocessing) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPathQueryParameters2D) GetPathPostprocessing() NavigationPathQueryParameters2DPathPostProcessing {
  classNameV := StringNameFromStr("NavigationPathQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path_postprocessing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3798118993) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret NavigationPathQueryParameters2DPathPostProcessing

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *NavigationPathQueryParameters2D) SetMap(map_ RID, )  {
  classNameV := StringNameFromStr("NavigationPathQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPathQueryParameters2D) GetMap() RID {
  classNameV := StringNameFromStr("NavigationPathQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationPathQueryParameters2D) SetStartPosition(start_position Vector2, )  {
  classNameV := StringNameFromStr("NavigationPathQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_start_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{start_position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPathQueryParameters2D) GetStartPosition() Vector2 {
  classNameV := StringNameFromStr("NavigationPathQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_start_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationPathQueryParameters2D) SetTargetPosition(target_position Vector2, )  {
  classNameV := StringNameFromStr("NavigationPathQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_target_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{target_position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPathQueryParameters2D) GetTargetPosition() Vector2 {
  classNameV := StringNameFromStr("NavigationPathQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_target_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationPathQueryParameters2D) SetNavigationLayers(navigation_layers int64, )  {
  classNameV := StringNameFromStr("NavigationPathQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_navigation_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&navigation_layers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPathQueryParameters2D) GetNavigationLayers() int64 {
  classNameV := StringNameFromStr("NavigationPathQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationPathQueryParameters2D) SetMetadataFlags(flags NavigationPathQueryParameters2DPathMetadataFlags, )  {
  classNameV := StringNameFromStr("NavigationPathQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_metadata_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 24274129) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPathQueryParameters2D) GetMetadataFlags() NavigationPathQueryParameters2DPathMetadataFlags {
  classNameV := StringNameFromStr("NavigationPathQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_metadata_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 488152976) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret NavigationPathQueryParameters2DPathMetadataFlags

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
