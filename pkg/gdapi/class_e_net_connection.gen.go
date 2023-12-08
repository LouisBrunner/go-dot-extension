// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ENetConnection struct {
  obj gdc.ObjectPtr
}

func (me *ENetConnection) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ENetConnection) BaseClass() string {
  return "ENetConnection"
}

type ENetConnectionCompressionMode int
const (
  ENetConnectionCompressNone ENetConnectionCompressionMode = 0
  ENetConnectionCompressRangeCoder ENetConnectionCompressionMode = 1
  ENetConnectionCompressFastlz ENetConnectionCompressionMode = 2
  ENetConnectionCompressZlib ENetConnectionCompressionMode = 3
  ENetConnectionCompressZstd ENetConnectionCompressionMode = 4
)

type ENetConnectionEventType int
const (
  ENetConnectionEventError ENetConnectionEventType = -1
  ENetConnectionEventNone ENetConnectionEventType = 0
  ENetConnectionEventConnect ENetConnectionEventType = 1
  ENetConnectionEventDisconnect ENetConnectionEventType = 2
  ENetConnectionEventReceive ENetConnectionEventType = 3
)

type ENetConnectionHostStatistic int
const (
  ENetConnectionHostTotalSentData ENetConnectionHostStatistic = 0
  ENetConnectionHostTotalSentPackets ENetConnectionHostStatistic = 1
  ENetConnectionHostTotalReceivedData ENetConnectionHostStatistic = 2
  ENetConnectionHostTotalReceivedPackets ENetConnectionHostStatistic = 3
)
