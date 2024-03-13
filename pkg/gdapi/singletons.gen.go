// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type singletons struct {
  Performance Performance
  TextServerManager TextServerManager
  PhysicsServer2DManager PhysicsServer2DManager
  PhysicsServer3DManager PhysicsServer3DManager
  NavigationMeshGenerator NavigationMeshGenerator
  ProjectSettings ProjectSettings
  IP IP
  Geometry2D Geometry2D
  Geometry3D Geometry3D
  ResourceLoader ResourceLoader
  ResourceSaver ResourceSaver
  OS OS
  Engine Engine
  ClassDB ClassDB
  Marshalls Marshalls
  TranslationServer TranslationServer
  Input Input
  InputMap InputMap
  EngineDebugger EngineDebugger
  Time Time
  GDExtensionManager GDExtensionManager
  ResourceUID ResourceUID
  WorkerThreadPool WorkerThreadPool
  ThemeDB ThemeDB
  EditorInterface EditorInterface
  JavaClassWrapper JavaClassWrapper
  JavaScriptBridge JavaScriptBridge
  DisplayServer DisplayServer
  RenderingServer RenderingServer
  AudioServer AudioServer
  PhysicsServer2D PhysicsServer2D
  PhysicsServer3D PhysicsServer3D
  NavigationServer2D NavigationServer2D
  NavigationServer3D NavigationServer3D
  XRServer XRServer
  CameraServer CameraServer
}

func newSingletons(iface gdc.Interface) *singletons {
  strPerformance := StringNameFromStr("Performance")
  defer strPerformance.Destroy()
  strTextServerManager := StringNameFromStr("TextServerManager")
  defer strTextServerManager.Destroy()
  strPhysicsServer2DManager := StringNameFromStr("PhysicsServer2DManager")
  defer strPhysicsServer2DManager.Destroy()
  strPhysicsServer3DManager := StringNameFromStr("PhysicsServer3DManager")
  defer strPhysicsServer3DManager.Destroy()
  strNavigationMeshGenerator := StringNameFromStr("NavigationMeshGenerator")
  defer strNavigationMeshGenerator.Destroy()
  strProjectSettings := StringNameFromStr("ProjectSettings")
  defer strProjectSettings.Destroy()
  strIP := StringNameFromStr("IP")
  defer strIP.Destroy()
  strGeometry2D := StringNameFromStr("Geometry2D")
  defer strGeometry2D.Destroy()
  strGeometry3D := StringNameFromStr("Geometry3D")
  defer strGeometry3D.Destroy()
  strResourceLoader := StringNameFromStr("ResourceLoader")
  defer strResourceLoader.Destroy()
  strResourceSaver := StringNameFromStr("ResourceSaver")
  defer strResourceSaver.Destroy()
  strOS := StringNameFromStr("OS")
  defer strOS.Destroy()
  strEngine := StringNameFromStr("Engine")
  defer strEngine.Destroy()
  strClassDB := StringNameFromStr("ClassDB")
  defer strClassDB.Destroy()
  strMarshalls := StringNameFromStr("Marshalls")
  defer strMarshalls.Destroy()
  strTranslationServer := StringNameFromStr("TranslationServer")
  defer strTranslationServer.Destroy()
  strInput := StringNameFromStr("Input")
  defer strInput.Destroy()
  strInputMap := StringNameFromStr("InputMap")
  defer strInputMap.Destroy()
  strEngineDebugger := StringNameFromStr("EngineDebugger")
  defer strEngineDebugger.Destroy()
  strTime := StringNameFromStr("Time")
  defer strTime.Destroy()
  strGDExtensionManager := StringNameFromStr("GDExtensionManager")
  defer strGDExtensionManager.Destroy()
  strResourceUID := StringNameFromStr("ResourceUID")
  defer strResourceUID.Destroy()
  strWorkerThreadPool := StringNameFromStr("WorkerThreadPool")
  defer strWorkerThreadPool.Destroy()
  strThemeDB := StringNameFromStr("ThemeDB")
  defer strThemeDB.Destroy()
  strEditorInterface := StringNameFromStr("EditorInterface")
  defer strEditorInterface.Destroy()
  strJavaClassWrapper := StringNameFromStr("JavaClassWrapper")
  defer strJavaClassWrapper.Destroy()
  strJavaScriptBridge := StringNameFromStr("JavaScriptBridge")
  defer strJavaScriptBridge.Destroy()
  strDisplayServer := StringNameFromStr("DisplayServer")
  defer strDisplayServer.Destroy()
  strRenderingServer := StringNameFromStr("RenderingServer")
  defer strRenderingServer.Destroy()
  strAudioServer := StringNameFromStr("AudioServer")
  defer strAudioServer.Destroy()
  strPhysicsServer2D := StringNameFromStr("PhysicsServer2D")
  defer strPhysicsServer2D.Destroy()
  strPhysicsServer3D := StringNameFromStr("PhysicsServer3D")
  defer strPhysicsServer3D.Destroy()
  strNavigationServer2D := StringNameFromStr("NavigationServer2D")
  defer strNavigationServer2D.Destroy()
  strNavigationServer3D := StringNameFromStr("NavigationServer3D")
  defer strNavigationServer3D.Destroy()
  strXRServer := StringNameFromStr("XRServer")
  defer strXRServer.Destroy()
  strCameraServer := StringNameFromStr("CameraServer")
  defer strCameraServer.Destroy()
  return &singletons{
    Performance: Performance{
      obj: iface.GlobalGetSingleton(strPerformance.AsCPtr()),
    },
    TextServerManager: TextServerManager{
      obj: iface.GlobalGetSingleton(strTextServerManager.AsCPtr()),
    },
    PhysicsServer2DManager: PhysicsServer2DManager{
      obj: iface.GlobalGetSingleton(strPhysicsServer2DManager.AsCPtr()),
    },
    PhysicsServer3DManager: PhysicsServer3DManager{
      obj: iface.GlobalGetSingleton(strPhysicsServer3DManager.AsCPtr()),
    },
    NavigationMeshGenerator: NavigationMeshGenerator{
      obj: iface.GlobalGetSingleton(strNavigationMeshGenerator.AsCPtr()),
    },
    ProjectSettings: ProjectSettings{
      obj: iface.GlobalGetSingleton(strProjectSettings.AsCPtr()),
    },
    IP: IP{
      obj: iface.GlobalGetSingleton(strIP.AsCPtr()),
    },
    Geometry2D: Geometry2D{
      obj: iface.GlobalGetSingleton(strGeometry2D.AsCPtr()),
    },
    Geometry3D: Geometry3D{
      obj: iface.GlobalGetSingleton(strGeometry3D.AsCPtr()),
    },
    ResourceLoader: ResourceLoader{
      obj: iface.GlobalGetSingleton(strResourceLoader.AsCPtr()),
    },
    ResourceSaver: ResourceSaver{
      obj: iface.GlobalGetSingleton(strResourceSaver.AsCPtr()),
    },
    OS: OS{
      obj: iface.GlobalGetSingleton(strOS.AsCPtr()),
    },
    Engine: Engine{
      obj: iface.GlobalGetSingleton(strEngine.AsCPtr()),
    },
    ClassDB: ClassDB{
      obj: iface.GlobalGetSingleton(strClassDB.AsCPtr()),
    },
    Marshalls: Marshalls{
      obj: iface.GlobalGetSingleton(strMarshalls.AsCPtr()),
    },
    TranslationServer: TranslationServer{
      obj: iface.GlobalGetSingleton(strTranslationServer.AsCPtr()),
    },
    Input: Input{
      obj: iface.GlobalGetSingleton(strInput.AsCPtr()),
    },
    InputMap: InputMap{
      obj: iface.GlobalGetSingleton(strInputMap.AsCPtr()),
    },
    EngineDebugger: EngineDebugger{
      obj: iface.GlobalGetSingleton(strEngineDebugger.AsCPtr()),
    },
    Time: Time{
      obj: iface.GlobalGetSingleton(strTime.AsCPtr()),
    },
    GDExtensionManager: GDExtensionManager{
      obj: iface.GlobalGetSingleton(strGDExtensionManager.AsCPtr()),
    },
    ResourceUID: ResourceUID{
      obj: iface.GlobalGetSingleton(strResourceUID.AsCPtr()),
    },
    WorkerThreadPool: WorkerThreadPool{
      obj: iface.GlobalGetSingleton(strWorkerThreadPool.AsCPtr()),
    },
    ThemeDB: ThemeDB{
      obj: iface.GlobalGetSingleton(strThemeDB.AsCPtr()),
    },
    EditorInterface: EditorInterface{
      obj: iface.GlobalGetSingleton(strEditorInterface.AsCPtr()),
    },
    JavaClassWrapper: JavaClassWrapper{
      obj: iface.GlobalGetSingleton(strJavaClassWrapper.AsCPtr()),
    },
    JavaScriptBridge: JavaScriptBridge{
      obj: iface.GlobalGetSingleton(strJavaScriptBridge.AsCPtr()),
    },
    DisplayServer: DisplayServer{
      obj: iface.GlobalGetSingleton(strDisplayServer.AsCPtr()),
    },
    RenderingServer: RenderingServer{
      obj: iface.GlobalGetSingleton(strRenderingServer.AsCPtr()),
    },
    AudioServer: AudioServer{
      obj: iface.GlobalGetSingleton(strAudioServer.AsCPtr()),
    },
    PhysicsServer2D: PhysicsServer2D{
      obj: iface.GlobalGetSingleton(strPhysicsServer2D.AsCPtr()),
    },
    PhysicsServer3D: PhysicsServer3D{
      obj: iface.GlobalGetSingleton(strPhysicsServer3D.AsCPtr()),
    },
    NavigationServer2D: NavigationServer2D{
      obj: iface.GlobalGetSingleton(strNavigationServer2D.AsCPtr()),
    },
    NavigationServer3D: NavigationServer3D{
      obj: iface.GlobalGetSingleton(strNavigationServer3D.AsCPtr()),
    },
    XRServer: XRServer{
      obj: iface.GlobalGetSingleton(strXRServer.AsCPtr()),
    },
    CameraServer: CameraServer{
      obj: iface.GlobalGetSingleton(strCameraServer.AsCPtr()),
    },
}
}
