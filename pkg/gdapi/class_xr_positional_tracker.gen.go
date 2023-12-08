// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type XRPositionalTracker struct {
  obj gdc.ObjectPtr
}

func (me *XRPositionalTracker) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *XRPositionalTracker) BaseClass() string {
  return "XRPositionalTracker"
}

type XRPositionalTrackerTrackerHand int
const (
  XRPositionalTrackerTrackerHandUnknown XRPositionalTrackerTrackerHand = 0
  XRPositionalTrackerTrackerHandLeft XRPositionalTrackerTrackerHand = 1
  XRPositionalTrackerTrackerHandRight XRPositionalTrackerTrackerHand = 2
)
