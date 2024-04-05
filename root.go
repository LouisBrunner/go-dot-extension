package gde

import (
	"github.com/LouisBrunner/go-dot-extension/pkg/entry"
	"github.com/LouisBrunner/go-dot-extension/pkg/gdextension"
)

var Register = entry.Register

var Logf = entry.Logf

var MarshalDict = gdextension.MarshalDict
var UnmarshalDict = gdextension.UnmarshalDict

type Class = gdextension.Class
type Signal = gdextension.Signal
