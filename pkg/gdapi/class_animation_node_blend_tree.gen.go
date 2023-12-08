// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNodeBlendTree struct {
  obj gdc.ObjectPtr
}

func (me *AnimationNodeBlendTree) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationNodeBlendTree) BaseClass() string {
  return "AnimationNodeBlendTree"
}

const (
  AnimationNodeBlendTreeCONNECTION_OK = 0
  AnimationNodeBlendTreeCONNECTION_ERROR_NO_INPUT = 1
  AnimationNodeBlendTreeCONNECTION_ERROR_NO_INPUT_INDEX = 2
  AnimationNodeBlendTreeCONNECTION_ERROR_NO_OUTPUT = 3
  AnimationNodeBlendTreeCONNECTION_ERROR_SAME_NODE = 4
  AnimationNodeBlendTreeCONNECTION_ERROR_CONNECTION_EXISTS = 5
)
