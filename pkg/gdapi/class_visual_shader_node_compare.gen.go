// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  classNameV := StringNameFromStr("VisualShaderNodeCompare")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_comparison_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 516558320) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeCompare) GetComparisonType() VisualShaderNodeCompareComparisonType {
  classNameV := StringNameFromStr("VisualShaderNodeCompare")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_comparison_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3495315961) // FIXME: should cache?
  var ret VisualShaderNodeCompareComparisonType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeCompare) SetFunction(func_ VisualShaderNodeCompareFunction, )  {
  classNameV := StringNameFromStr("VisualShaderNodeCompare")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_function")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2370951349) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&func_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeCompare) GetFunction() VisualShaderNodeCompareFunction {
  classNameV := StringNameFromStr("VisualShaderNodeCompare")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_function")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4089164265) // FIXME: should cache?
  var ret VisualShaderNodeCompareFunction
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeCompare) SetCondition(condition VisualShaderNodeCompareCondition, )  {
  classNameV := StringNameFromStr("VisualShaderNodeCompare")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_condition")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 918742392) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&condition), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeCompare) GetCondition() VisualShaderNodeCompareCondition {
  classNameV := StringNameFromStr("VisualShaderNodeCompare")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_condition")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3281078941) // FIXME: should cache?
  var ret VisualShaderNodeCompareCondition
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
