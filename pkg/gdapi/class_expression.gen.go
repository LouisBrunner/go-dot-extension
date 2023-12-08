// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Expression struct {
  obj gdc.ObjectPtr
}

func createExpression(obj gdc.ObjectPtr) *Expression {
  return &Expression{
    obj: obj,
  }
}

func (me *Expression) BaseClass() string {
  return "Expression"
}
