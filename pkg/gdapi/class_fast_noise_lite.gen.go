// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type FastNoiseLite struct {
  obj gdc.ObjectPtr
}

func (me *FastNoiseLite) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *FastNoiseLite) BaseClass() string {
  return "FastNoiseLite"
}

type FastNoiseLiteNoiseType int
const (
  FastNoiseLiteTypeValue FastNoiseLiteNoiseType = 5
  FastNoiseLiteTypeValueCubic FastNoiseLiteNoiseType = 4
  FastNoiseLiteTypePerlin FastNoiseLiteNoiseType = 3
  FastNoiseLiteTypeCellular FastNoiseLiteNoiseType = 2
  FastNoiseLiteTypeSimplex FastNoiseLiteNoiseType = 0
  FastNoiseLiteTypeSimplexSmooth FastNoiseLiteNoiseType = 1
)

type FastNoiseLiteFractalType int
const (
  FastNoiseLiteFractalNone FastNoiseLiteFractalType = 0
  FastNoiseLiteFractalFbm FastNoiseLiteFractalType = 1
  FastNoiseLiteFractalRidged FastNoiseLiteFractalType = 2
  FastNoiseLiteFractalPingPong FastNoiseLiteFractalType = 3
)

type FastNoiseLiteCellularDistanceFunction int
const (
  FastNoiseLiteDistanceEuclidean FastNoiseLiteCellularDistanceFunction = 0
  FastNoiseLiteDistanceEuclideanSquared FastNoiseLiteCellularDistanceFunction = 1
  FastNoiseLiteDistanceManhattan FastNoiseLiteCellularDistanceFunction = 2
  FastNoiseLiteDistanceHybrid FastNoiseLiteCellularDistanceFunction = 3
)

type FastNoiseLiteCellularReturnType int
const (
  FastNoiseLiteReturnCellValue FastNoiseLiteCellularReturnType = 0
  FastNoiseLiteReturnDistance FastNoiseLiteCellularReturnType = 1
  FastNoiseLiteReturnDistance2 FastNoiseLiteCellularReturnType = 2
  FastNoiseLiteReturnDistance2Add FastNoiseLiteCellularReturnType = 3
  FastNoiseLiteReturnDistance2Sub FastNoiseLiteCellularReturnType = 4
  FastNoiseLiteReturnDistance2Mul FastNoiseLiteCellularReturnType = 5
  FastNoiseLiteReturnDistance2Div FastNoiseLiteCellularReturnType = 6
)

type FastNoiseLiteDomainWarpType int
const (
  FastNoiseLiteDomainWarpSimplex FastNoiseLiteDomainWarpType = 0
  FastNoiseLiteDomainWarpSimplexReduced FastNoiseLiteDomainWarpType = 1
  FastNoiseLiteDomainWarpBasicGrid FastNoiseLiteDomainWarpType = 2
)

type FastNoiseLiteDomainWarpFractalType int
const (
  FastNoiseLiteDomainWarpFractalNone FastNoiseLiteDomainWarpFractalType = 0
  FastNoiseLiteDomainWarpFractalProgressive FastNoiseLiteDomainWarpFractalType = 1
  FastNoiseLiteDomainWarpFractalIndependent FastNoiseLiteDomainWarpFractalType = 2
)
