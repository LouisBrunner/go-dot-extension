// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type OS struct {
  obj gdc.ObjectPtr
}

func (me *OS) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *OS) BaseClass() string {
  return "OS"
}

type OSRenderingDriver int
const (
  OSRenderingDriverVulkan OSRenderingDriver = 0
  OSRenderingDriverOpengl3 OSRenderingDriver = 1
)

type OSSystemDir int
const (
  OSSystemDirDesktop OSSystemDir = 0
  OSSystemDirDcim OSSystemDir = 1
  OSSystemDirDocuments OSSystemDir = 2
  OSSystemDirDownloads OSSystemDir = 3
  OSSystemDirMovies OSSystemDir = 4
  OSSystemDirMusic OSSystemDir = 5
  OSSystemDirPictures OSSystemDir = 6
  OSSystemDirRingtones OSSystemDir = 7
)
