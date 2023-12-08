// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  FastNoiseLiteNoiseTypeTypeValue FastNoiseLiteNoiseType = 5
  FastNoiseLiteNoiseTypeTypeValueCubic FastNoiseLiteNoiseType = 4
  FastNoiseLiteNoiseTypeTypePerlin FastNoiseLiteNoiseType = 3
  FastNoiseLiteNoiseTypeTypeCellular FastNoiseLiteNoiseType = 2
  FastNoiseLiteNoiseTypeTypeSimplex FastNoiseLiteNoiseType = 0
  FastNoiseLiteNoiseTypeTypeSimplexSmooth FastNoiseLiteNoiseType = 1
)

type FastNoiseLiteFractalType int
const (
  FastNoiseLiteFractalTypeFractalNone FastNoiseLiteFractalType = 0
  FastNoiseLiteFractalTypeFractalFbm FastNoiseLiteFractalType = 1
  FastNoiseLiteFractalTypeFractalRidged FastNoiseLiteFractalType = 2
  FastNoiseLiteFractalTypeFractalPingPong FastNoiseLiteFractalType = 3
)

type FastNoiseLiteCellularDistanceFunction int
const (
  FastNoiseLiteCellularDistanceFunctionDistanceEuclidean FastNoiseLiteCellularDistanceFunction = 0
  FastNoiseLiteCellularDistanceFunctionDistanceEuclideanSquared FastNoiseLiteCellularDistanceFunction = 1
  FastNoiseLiteCellularDistanceFunctionDistanceManhattan FastNoiseLiteCellularDistanceFunction = 2
  FastNoiseLiteCellularDistanceFunctionDistanceHybrid FastNoiseLiteCellularDistanceFunction = 3
)

type FastNoiseLiteCellularReturnType int
const (
  FastNoiseLiteCellularReturnTypeReturnCellValue FastNoiseLiteCellularReturnType = 0
  FastNoiseLiteCellularReturnTypeReturnDistance FastNoiseLiteCellularReturnType = 1
  FastNoiseLiteCellularReturnTypeReturnDistance2 FastNoiseLiteCellularReturnType = 2
  FastNoiseLiteCellularReturnTypeReturnDistance2Add FastNoiseLiteCellularReturnType = 3
  FastNoiseLiteCellularReturnTypeReturnDistance2Sub FastNoiseLiteCellularReturnType = 4
  FastNoiseLiteCellularReturnTypeReturnDistance2Mul FastNoiseLiteCellularReturnType = 5
  FastNoiseLiteCellularReturnTypeReturnDistance2Div FastNoiseLiteCellularReturnType = 6
)

type FastNoiseLiteDomainWarpType int
const (
  FastNoiseLiteDomainWarpTypeDomainWarpSimplex FastNoiseLiteDomainWarpType = 0
  FastNoiseLiteDomainWarpTypeDomainWarpSimplexReduced FastNoiseLiteDomainWarpType = 1
  FastNoiseLiteDomainWarpTypeDomainWarpBasicGrid FastNoiseLiteDomainWarpType = 2
)

type FastNoiseLiteDomainWarpFractalType int
const (
  FastNoiseLiteDomainWarpFractalTypeDomainWarpFractalNone FastNoiseLiteDomainWarpFractalType = 0
  FastNoiseLiteDomainWarpFractalTypeDomainWarpFractalProgressive FastNoiseLiteDomainWarpFractalType = 1
  FastNoiseLiteDomainWarpFractalTypeDomainWarpFractalIndependent FastNoiseLiteDomainWarpFractalType = 2
)

func  (me *FastNoiseLite) SetNoiseType(type_ FastNoiseLiteNoiseType, ) { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) GetNoiseType() { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) SetSeed(seed int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) GetSeed() { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) SetFrequency(freq float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) GetFrequency() { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) SetOffset(offset Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) GetOffset() { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) SetFractalType(type_ FastNoiseLiteFractalType, ) { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) GetFractalType() { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) SetFractalOctaves(octave_count int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) GetFractalOctaves() { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) SetFractalLacunarity(lacunarity float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) GetFractalLacunarity() { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) SetFractalGain(gain float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) GetFractalGain() { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) SetFractalWeightedStrength(weighted_strength float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) GetFractalWeightedStrength() { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) SetFractalPingPongStrength(ping_pong_strength float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) GetFractalPingPongStrength() { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) SetCellularDistanceFunction(func_ FastNoiseLiteCellularDistanceFunction, ) { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) GetCellularDistanceFunction() { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) SetCellularJitter(jitter float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) GetCellularJitter() { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) SetCellularReturnType(ret FastNoiseLiteCellularReturnType, ) { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) GetCellularReturnType() { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) SetDomainWarpEnabled(domain_warp_enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) IsDomainWarpEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) SetDomainWarpType(domain_warp_type FastNoiseLiteDomainWarpType, ) { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) GetDomainWarpType() { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) SetDomainWarpAmplitude(domain_warp_amplitude float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) GetDomainWarpAmplitude() { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) SetDomainWarpFrequency(domain_warp_frequency float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) GetDomainWarpFrequency() { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) SetDomainWarpFractalType(domain_warp_fractal_type FastNoiseLiteDomainWarpFractalType, ) { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) GetDomainWarpFractalType() { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) SetDomainWarpFractalOctaves(domain_warp_octave_count int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) GetDomainWarpFractalOctaves() { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) SetDomainWarpFractalLacunarity(domain_warp_lacunarity float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) GetDomainWarpFractalLacunarity() { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) SetDomainWarpFractalGain(domain_warp_gain float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *FastNoiseLite) GetDomainWarpFractalGain() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
