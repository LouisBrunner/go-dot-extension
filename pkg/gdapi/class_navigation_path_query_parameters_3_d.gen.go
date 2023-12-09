// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationPathQueryParameters3D struct {
  obj gdc.ObjectPtr
}

func (me *NavigationPathQueryParameters3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
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

func (me *NavigationPathQueryParameters3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NavigationPathQueryParameters3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *NavigationPathQueryParameters3D) SetPathfindingAlgorithm(pathfinding_algorithm NavigationPathQueryParameters3DPathfindingAlgorithm, )  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryParameters3D) GetPathfindingAlgorithm()  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryParameters3D) SetPathPostprocessing(path_postprocessing NavigationPathQueryParameters3DPathPostProcessing, )  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryParameters3D) GetPathPostprocessing()  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryParameters3D) SetMap(map_ RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryParameters3D) GetMap()  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryParameters3D) SetStartPosition(start_position Vector3, )  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryParameters3D) GetStartPosition()  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryParameters3D) SetTargetPosition(target_position Vector3, )  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryParameters3D) GetTargetPosition()  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryParameters3D) SetNavigationLayers(navigation_layers int, )  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryParameters3D) GetNavigationLayers()  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryParameters3D) SetMetadataFlags(flags NavigationPathQueryParameters3DPathMetadataFlags, )  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryParameters3D) GetMetadataFlags()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
