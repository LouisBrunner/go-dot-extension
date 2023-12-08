// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MultiplayerAPI struct {
  obj gdc.ObjectPtr
}

func (me *MultiplayerAPI) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MultiplayerAPI) BaseClass() string {
  return "MultiplayerAPI"
}

type MultiplayerAPIRPCMode int
const (
  MultiplayerAPIRpcModeDisabled MultiplayerAPIRPCMode = 0
  MultiplayerAPIRpcModeAnyPeer MultiplayerAPIRPCMode = 1
  MultiplayerAPIRpcModeAuthority MultiplayerAPIRPCMode = 2
)
