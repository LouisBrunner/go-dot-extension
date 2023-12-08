// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Geometry2D struct {
  obj gdc.ObjectPtr
}

func (me *Geometry2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Geometry2D) BaseClass() string {
  return "Geometry2D"
}

type Geometry2DPolyBooleanOperation int
const (
  Geometry2DOperationUnion Geometry2DPolyBooleanOperation = 0
  Geometry2DOperationDifference Geometry2DPolyBooleanOperation = 1
  Geometry2DOperationIntersection Geometry2DPolyBooleanOperation = 2
  Geometry2DOperationXor Geometry2DPolyBooleanOperation = 3
)

type Geometry2DPolyJoinType int
const (
  Geometry2DJoinSquare Geometry2DPolyJoinType = 0
  Geometry2DJoinRound Geometry2DPolyJoinType = 1
  Geometry2DJoinMiter Geometry2DPolyJoinType = 2
)

type Geometry2DPolyEndType int
const (
  Geometry2DEndPolygon Geometry2DPolyEndType = 0
  Geometry2DEndJoined Geometry2DPolyEndType = 1
  Geometry2DEndButt Geometry2DPolyEndType = 2
  Geometry2DEndSquare Geometry2DPolyEndType = 3
  Geometry2DEndRound Geometry2DPolyEndType = 4
)
