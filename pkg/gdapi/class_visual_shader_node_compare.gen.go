// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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



// Enums

type VisualShaderNodeCompareComparisonType int
const (
  VisualShaderNodeCompareComparisonTypeCtypeScalar VisualShaderNodeCompareComparisonType = 0
  VisualShaderNodeCompareComparisonTypeCtypeScalarInt VisualShaderNodeCompareComparisonType = 1
  VisualShaderNodeCompareComparisonTypeCtypeScalarUint VisualShaderNodeCompareComparisonType = 2
  VisualShaderNodeCompareComparisonTypeCtypeVector2D VisualShaderNodeCompareComparisonType = 3
  VisualShaderNodeCompareComparisonTypeCtypeVector3D VisualShaderNodeCompareComparisonType = 4
  VisualShaderNodeCompareComparisonTypeCtypeVector4D VisualShaderNodeCompareComparisonType = 5
  VisualShaderNodeCompareComparisonTypeCtypeBoolean VisualShaderNodeCompareComparisonType = 6
  VisualShaderNodeCompareComparisonTypeCtypeTransform VisualShaderNodeCompareComparisonType = 7
  VisualShaderNodeCompareComparisonTypeCtypeMax VisualShaderNodeCompareComparisonType = 8
)

type VisualShaderNodeCompareFunction int
const (
  VisualShaderNodeCompareFunctionFuncEqual VisualShaderNodeCompareFunction = 0
  VisualShaderNodeCompareFunctionFuncNotEqual VisualShaderNodeCompareFunction = 1
  VisualShaderNodeCompareFunctionFuncGreaterThan VisualShaderNodeCompareFunction = 2
  VisualShaderNodeCompareFunctionFuncGreaterThanEqual VisualShaderNodeCompareFunction = 3
  VisualShaderNodeCompareFunctionFuncLessThan VisualShaderNodeCompareFunction = 4
  VisualShaderNodeCompareFunctionFuncLessThanEqual VisualShaderNodeCompareFunction = 5
  VisualShaderNodeCompareFunctionFuncMax VisualShaderNodeCompareFunction = 6
)

type VisualShaderNodeCompareCondition int
const (
  VisualShaderNodeCompareConditionCondAll VisualShaderNodeCompareCondition = 0
  VisualShaderNodeCompareConditionCondAny VisualShaderNodeCompareCondition = 1
  VisualShaderNodeCompareConditionCondMax VisualShaderNodeCompareCondition = 2
)

func (me *VisualShaderNodeCompare) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeCompare) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeCompare) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeCompare) SetComparisonType(type_ VisualShaderNodeCompareComparisonType, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeCompare) GetComparisonType()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeCompare) SetFunction(func_ VisualShaderNodeCompareFunction, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeCompare) GetFunction()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeCompare) SetCondition(condition VisualShaderNodeCompareCondition, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeCompare) GetCondition()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
