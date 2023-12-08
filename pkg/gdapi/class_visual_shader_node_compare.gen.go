// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeCompare struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeCompare) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeCompare) BaseClass() string {
  return "VisualShaderNodeCompare"
}

type VisualShaderNodeCompareComparisonType int
const (
  VisualShaderNodeCompareCtypeScalar VisualShaderNodeCompareComparisonType = 0
  VisualShaderNodeCompareCtypeScalarInt VisualShaderNodeCompareComparisonType = 1
  VisualShaderNodeCompareCtypeScalarUint VisualShaderNodeCompareComparisonType = 2
  VisualShaderNodeCompareCtypeVector2D VisualShaderNodeCompareComparisonType = 3
  VisualShaderNodeCompareCtypeVector3D VisualShaderNodeCompareComparisonType = 4
  VisualShaderNodeCompareCtypeVector4D VisualShaderNodeCompareComparisonType = 5
  VisualShaderNodeCompareCtypeBoolean VisualShaderNodeCompareComparisonType = 6
  VisualShaderNodeCompareCtypeTransform VisualShaderNodeCompareComparisonType = 7
  VisualShaderNodeCompareCtypeMax VisualShaderNodeCompareComparisonType = 8
)

type VisualShaderNodeCompareFunction int
const (
  VisualShaderNodeCompareFuncEqual VisualShaderNodeCompareFunction = 0
  VisualShaderNodeCompareFuncNotEqual VisualShaderNodeCompareFunction = 1
  VisualShaderNodeCompareFuncGreaterThan VisualShaderNodeCompareFunction = 2
  VisualShaderNodeCompareFuncGreaterThanEqual VisualShaderNodeCompareFunction = 3
  VisualShaderNodeCompareFuncLessThan VisualShaderNodeCompareFunction = 4
  VisualShaderNodeCompareFuncLessThanEqual VisualShaderNodeCompareFunction = 5
  VisualShaderNodeCompareFuncMax VisualShaderNodeCompareFunction = 6
)

type VisualShaderNodeCompareCondition int
const (
  VisualShaderNodeCompareCondAll VisualShaderNodeCompareCondition = 0
  VisualShaderNodeCompareCondAny VisualShaderNodeCompareCondition = 1
  VisualShaderNodeCompareCondMax VisualShaderNodeCompareCondition = 2
)
