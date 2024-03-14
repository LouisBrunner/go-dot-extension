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
      Object: Object{
        obj: iface.GlobalGetSingleton(strPerformance.AsCPtr()),
      },
    },
    TextServerManager: TextServerManager{
      Object: Object{
        obj: iface.GlobalGetSingleton(strTextServerManager.AsCPtr()),
      },
    },
    PhysicsServer2DManager: PhysicsServer2DManager{
      Object: Object{
        obj: iface.GlobalGetSingleton(strPhysicsServer2DManager.AsCPtr()),
      },
    },
    PhysicsServer3DManager: PhysicsServer3DManager{
      Object: Object{
        obj: iface.GlobalGetSingleton(strPhysicsServer3DManager.AsCPtr()),
      },
    },
    NavigationMeshGenerator: NavigationMeshGenerator{
      Object: Object{
        obj: iface.GlobalGetSingleton(strNavigationMeshGenerator.AsCPtr()),
      },
    },
    ProjectSettings: ProjectSettings{
      Object: Object{
        obj: iface.GlobalGetSingleton(strProjectSettings.AsCPtr()),
      },
    },
    IP: IP{
      Object: Object{
        obj: iface.GlobalGetSingleton(strIP.AsCPtr()),
      },
    },
    Geometry2D: Geometry2D{
      Object: Object{
        obj: iface.GlobalGetSingleton(strGeometry2D.AsCPtr()),
      },
    },
    Geometry3D: Geometry3D{
      Object: Object{
        obj: iface.GlobalGetSingleton(strGeometry3D.AsCPtr()),
      },
    },
    ResourceLoader: ResourceLoader{
      Object: Object{
        obj: iface.GlobalGetSingleton(strResourceLoader.AsCPtr()),
      },
    },
    ResourceSaver: ResourceSaver{
      Object: Object{
        obj: iface.GlobalGetSingleton(strResourceSaver.AsCPtr()),
      },
    },
    OS: OS{
      Object: Object{
        obj: iface.GlobalGetSingleton(strOS.AsCPtr()),
      },
    },
    Engine: Engine{
      Object: Object{
        obj: iface.GlobalGetSingleton(strEngine.AsCPtr()),
      },
    },
    ClassDB: ClassDB{
      Object: Object{
        obj: iface.GlobalGetSingleton(strClassDB.AsCPtr()),
      },
    },
    Marshalls: Marshalls{
      Object: Object{
        obj: iface.GlobalGetSingleton(strMarshalls.AsCPtr()),
      },
    },
    TranslationServer: TranslationServer{
      Object: Object{
        obj: iface.GlobalGetSingleton(strTranslationServer.AsCPtr()),
      },
    },
    Input: Input{
      Object: Object{
        obj: iface.GlobalGetSingleton(strInput.AsCPtr()),
      },
    },
    InputMap: InputMap{
      Object: Object{
        obj: iface.GlobalGetSingleton(strInputMap.AsCPtr()),
      },
    },
    EngineDebugger: EngineDebugger{
      Object: Object{
        obj: iface.GlobalGetSingleton(strEngineDebugger.AsCPtr()),
      },
    },
    Time: Time{
      Object: Object{
        obj: iface.GlobalGetSingleton(strTime.AsCPtr()),
      },
    },
    GDExtensionManager: GDExtensionManager{
      Object: Object{
        obj: iface.GlobalGetSingleton(strGDExtensionManager.AsCPtr()),
      },
    },
    ResourceUID: ResourceUID{
      Object: Object{
        obj: iface.GlobalGetSingleton(strResourceUID.AsCPtr()),
      },
    },
    WorkerThreadPool: WorkerThreadPool{
      Object: Object{
        obj: iface.GlobalGetSingleton(strWorkerThreadPool.AsCPtr()),
      },
    },
    ThemeDB: ThemeDB{
      Object: Object{
        obj: iface.GlobalGetSingleton(strThemeDB.AsCPtr()),
      },
    },
    EditorInterface: EditorInterface{
      Object: Object{
        obj: iface.GlobalGetSingleton(strEditorInterface.AsCPtr()),
      },
    },
    JavaClassWrapper: JavaClassWrapper{
      Object: Object{
        obj: iface.GlobalGetSingleton(strJavaClassWrapper.AsCPtr()),
      },
    },
    JavaScriptBridge: JavaScriptBridge{
      Object: Object{
        obj: iface.GlobalGetSingleton(strJavaScriptBridge.AsCPtr()),
      },
    },
    DisplayServer: DisplayServer{
      Object: Object{
        obj: iface.GlobalGetSingleton(strDisplayServer.AsCPtr()),
      },
    },
    RenderingServer: RenderingServer{
      Object: Object{
        obj: iface.GlobalGetSingleton(strRenderingServer.AsCPtr()),
      },
    },
    AudioServer: AudioServer{
      Object: Object{
        obj: iface.GlobalGetSingleton(strAudioServer.AsCPtr()),
      },
    },
    PhysicsServer2D: PhysicsServer2D{
      Object: Object{
        obj: iface.GlobalGetSingleton(strPhysicsServer2D.AsCPtr()),
      },
    },
    PhysicsServer3D: PhysicsServer3D{
      Object: Object{
        obj: iface.GlobalGetSingleton(strPhysicsServer3D.AsCPtr()),
      },
    },
    NavigationServer2D: NavigationServer2D{
      Object: Object{
        obj: iface.GlobalGetSingleton(strNavigationServer2D.AsCPtr()),
      },
    },
    NavigationServer3D: NavigationServer3D{
      Object: Object{
        obj: iface.GlobalGetSingleton(strNavigationServer3D.AsCPtr()),
      },
    },
    XRServer: XRServer{
      Object: Object{
        obj: iface.GlobalGetSingleton(strXRServer.AsCPtr()),
      },
    },
    CameraServer: CameraServer{
      Object: Object{
        obj: iface.GlobalGetSingleton(strCameraServer.AsCPtr()),
      },
    },
}
}
