// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFState struct {
  obj gdc.ObjectPtr
}

func (me *GLTFState) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFState) BaseClass() string {
  return "GLTFState"
}

const (
  GLTFStateHANDLE_BINARY_DISCARD_TEXTURES = 0
  GLTFStateHANDLE_BINARY_EXTRACT_TEXTURES = 1
  GLTFStateHANDLE_BINARY_EMBED_AS_BASISU = 2
  GLTFStateHANDLE_BINARY_EMBED_AS_UNCOMPRESSED = 3
)
