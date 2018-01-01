package application

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/g3n/engine/audio/al"
	"github.com/g3n/engine/audio/ov"
	"github.com/g3n/engine/audio/vorbis"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/camera/control"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/util/logger"
	"github.com/g3n/engine/window"
)

// Application is a basic standard application object which can be extended.
// It creates a Window, OpenGL state, default cameras, default scene and Gui and
// has a method to run the render loop.
type Application struct {
	core.Dispatcher                         // Embedded event dispatcher
	core.TimerManager                       // Embedded timer manager
	win               window.IWindow        // Application window
	gl                *gls.GLS              // OpenGL state
	log               *logger.Logger        // Default application logger
	renderer          *renderer.Renderer    // Renderer object
	camPersp          *camera.Perspective   // Perspective camera
	camOrtho          *camera.Orthographic  // Orthographic camera
	camera            camera.ICamera        // Current camera
	orbit             *control.OrbitControl // Camera orbit controller
	audio             bool                  // Audio available
	vorbis            bool                  // Vorbis decoder available
	audioEFX          bool                  // Audio effect extension support available
	audioDev          *al.Device            // Audio player device
	captureDev        *al.Device            // Audio capture device
	frameRater        *FrameRater           // Render loop frame rater
	scene             *core.Node            // Node container for 3D tests
	guiroot           *gui.Root             // Gui root panel
	frameCount        uint64                // Frame counter
	frameTime         time.Time             // Time at the start of the frame
	frameDelta        time.Duration         // Time delta from previous frame
	startTime         time.Time             // Time at the start of the render loop
	fullScreen        *bool                 // Full screen option
	cpuProfile        *string               // File to write cpu profile to
	swapInterval      *int                  // Swap interval option
	targetFPS         *uint                 // Target FPS option
	noglErrors        *bool                 // No OpenGL check errors options

}

// Options defines initial options passed to application creation function
type Options struct {
	Height      int  // Initial window height (default is screen width)
	Width       int  // Initial window width (default is screen height)
	Fullscreen  bool // Window full screen flag (default = false)
	LogLevel    int  // Initial log level (default = DEBUG)
	EnableFlags bool // Enable command line flags (default = false)
	TargetFPS   uint // Desired frames per second rate (default = 60)
}

// appInstance contains the pointer to the single Application instance
var appInstance *Application

const (
	OnBeforeRender = "util.application.OnBeforeRender"
	OnAfterRender  = "util.application.OnAfterRender"
)

// Creates creates and returns the application object using the specified name for
// the window title and log messages
// This function must be called only once.
func Create(name string, ops Options) (*Application, error) {

	if appInstance != nil {
		return nil, fmt.Errorf("Application already created")
	}
	app := new(Application)
	appInstance = app
	app.Dispatcher.Initialize()
	app.TimerManager.Initialize()

	// Initialize options defaults
	app.fullScreen = new(bool)
	app.cpuProfile = new(string)
	app.swapInterval = new(int)
	*app.swapInterval = -1
	app.targetFPS = new(uint)
	*app.targetFPS = 60
	app.noglErrors = new(bool)

	// Options parameter overrides some options
	if ops.TargetFPS != 0 {
		*app.fullScreen = ops.Fullscreen
		*app.targetFPS = ops.TargetFPS
	}

	// Creates flags if requested (override options defaults)
	if ops.EnableFlags {
		app.fullScreen = flag.Bool("fullscreen", false, "Stars application with full screen")
		app.cpuProfile = flag.String("cpuprofile", "", "Activate cpu profiling writing profile to the specified file")
		app.swapInterval = flag.Int("swapinterval", -1, "Sets the swap buffers interval to this value")
		app.targetFPS = flag.Uint("targetfps", 60, "Sets the frame rate in frames per second")
		app.noglErrors = flag.Bool("noglerrors", false, "Do not check OpenGL errors at each call (may increase FPS)")
	}
	flag.Parse()

	// Creates application logger
	app.log = logger.New(name, nil)
	app.log.AddWriter(logger.NewConsole(false))
	app.log.SetFormat(logger.FTIME | logger.FMICROS)
	app.log.SetLevel(ops.LogLevel)

	// Window event handling must run on the main OS thread
	runtime.LockOSThread()

	// Creates window
	win, err := window.New("glfw", 801, 600, name, *app.fullScreen)
	if err != nil {
		return nil, err
	}
	app.win = win
	if !*app.fullScreen {
		// Calculates window size and position
		swidth, sheight := win.GetScreenResolution(nil)
		var posx, posy int
		if ops.Width != 0 {
			posx = (swidth - ops.Width) / 2
			if posx < 0 {
				posx = 0
			}
			swidth = ops.Width
		}
		if ops.Height != 0 {
			posy = (sheight - ops.Height) / 2
			if posy < 0 {
				posy = 0
			}
			sheight = ops.Height
		}
		// Sets the window size and position
		win.SetSize(swidth, sheight)
		win.SetPos(posx, posy)
	}
	// Create OpenGL state
	gl, err := gls.New()
	if err != nil {
		return nil, err
	}
	app.gl = gl
	// Checks OpenGL errors
	app.gl.SetCheckErrors(!*app.noglErrors)

	cc := math32.NewColor("gray")
	app.gl.ClearColor(cc.R, cc.G, cc.B, 1)
	app.gl.Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)

	// Creates perspective camera
	width, height := app.win.GetSize()
	aspect := float32(width) / float32(height)
	app.camPersp = camera.NewPerspective(65, aspect, 0.01, 1000)

	// Creates orthographic camera
	app.camOrtho = camera.NewOrthographic(-2, 2, 2, -2, 0.01, 100)
	app.camOrtho.SetPosition(0, 0, 3)
	app.camOrtho.LookAt(&math32.Vector3{0, 0, 0})
	app.camOrtho.SetZoom(1.0)

	// Default camera is perspective
	app.camera = app.camPersp

	// Creates orbit camera control
	// It is important to do this after the root panel subscription
	// to avoid GUI events being propagated to the orbit control.
	app.orbit = control.NewOrbitControl(app.camera, app.win)

	// Creates scene for 3D objects
	app.scene = core.NewNode()

	// Creates gui root panel
	app.guiroot = gui.NewRoot(app.gl, app.win)
	app.guiroot.SetColor(math32.NewColor("silver"))

	// Creates renderer
	app.renderer = renderer.NewRenderer(gl)
	err = app.renderer.AddDefaultShaders()
	if err != nil {
		return nil, fmt.Errorf("Error from AddDefaulShaders:%v", err)
	}
	app.renderer.SetScene(app.scene)
	app.renderer.SetGui(app.guiroot)

	// Create frame rater
	app.frameRater = NewFrameRater(*app.targetFPS)

	return app, nil
}

// Get returns the application single instance or nil
// if the application was not created yet
func Get() *Application {

	return appInstance
}

// Log returns the application logger
func (app *Application) Log() *logger.Logger {

	return app.log
}

// Window returns the application window
func (app *Application) Window() window.IWindow {

	return app.win
}

// Gl returns the OpenGL state
func (app *Application) Gl() *gls.GLS {

	return app.gl
}

// Gui returns the current application Gui root panel
func (app *Application) Gui() *gui.Root {

	return app.guiroot
}

// Scene returns the current application 3D scene
func (app *Application) Scene() *core.Node {

	return app.scene
}

// SetScene sets the 3D scene to be rendered
func (app *Application) SetScene(scene *core.Node) {

	app.renderer.SetScene(scene)
}

// SetGui sets the root panel of the gui to be rendered
func (app *Application) SetGui(root *gui.Root) {

	app.guiroot = root
	app.renderer.SetGui(app.guiroot)
}

// SetPanel3D sets the gui panel inside which the 3D scene is shown.
func (app *Application) SetPanel3D(panel3D gui.IPanel) {

	app.renderer.SetGuiPanel3D(panel3D)
}

// Panel3D returns the current gui panel where the 3D scene is shown.
func (app *Application) Panel3D() gui.IPanel {

	return app.renderer.Panel3D()
}

// CameraPersp returns the application perspective camera
func (app *Application) CameraPersp() *camera.Perspective {

	return app.camPersp
}

// CameraOrtho returns the application orthographic camera
func (app *Application) CameraOrtho() *camera.Orthographic {

	return app.camOrtho
}

// Camera returns the current application camera
func (app *Application) Camera() camera.ICamera {

	return app.camera
}

// SetCamera sets the current application camera
func (app *Application) SetCamera(cam camera.ICamera) {

	app.camera = cam
}

// Orbit returns the current camera orbit control
func (app *Application) Orbit() *control.OrbitControl {

	return app.orbit
}

// SetOrbit sets the camera orbit control
func (app *Application) SetOrbit(oc *control.OrbitControl) {

	app.orbit = oc
}

// FrameRater returns the FrameRater object
func (app *Application) FrameRater() *FrameRater {

	return app.frameRater
}

// FrameCount returns the total number of frames since the call to Run()
func (app *Application) FrameCount() uint64 {

	return app.frameCount
}

// FrameDelta returns the duration of the previous frame
func (app *Application) FrameDelta() time.Duration {

	return app.frameDelta
}

// FrameDeltaSeconds returns the duration of the previous frame
// in float32 seconds
func (app *Application) FrameDeltaSeconds() float32 {

	return float32(app.frameDelta.Seconds())
}

// RunTime returns the duration since the call to Run()
func (app *Application) RunTime() time.Duration {

	return time.Now().Sub(app.startTime)
}

// RunSeconds returns the elapsed time in seconds since the call to Run()
func (app *Application) RunSeconds() float32 {

	return float32(time.Now().Sub(app.startTime).Seconds())
}

// Renderer returns the application renderer
func (app *Application) Renderer() *renderer.Renderer {

	return app.renderer
}

// AudioSupport returns if the audio library was loaded OK
func (app *Application) AudioSupport() bool {

	return app.audio
}

// VorbisSupport returns if the Ogg Vorbis decoder library was loaded OK
func (app *Application) VorbisSupport() bool {

	return app.vorbis
}

// SetCpuProfile must be called before Run() and sets the file name for cpu profiling.
// If set the cpu profiling starts before running the render loop and continues
// till the end of the application.
func (app *Application) SetCpuProfile(fname string) {

	*app.cpuProfile = fname
}

// Runs runs the application render loop
func (app *Application) Run() error {

	// Set swap interval
	if *app.swapInterval >= 0 {
		app.win.SwapInterval(*app.swapInterval)
		app.log.Debug("Swap interval set to:%v", *app.swapInterval)
	}

	// Start profiling if requested
	if *app.cpuProfile != "" {
		f, err := os.Create(*app.cpuProfile)
		if err != nil {
			return err
		}
		err = pprof.StartCPUProfile(f)
		if err != nil {
			return err
		}
		app.log.Info("Started writing CPU profile to:%s", app.cpuProfile)
		defer pprof.StopCPUProfile()
	}

	app.startTime = time.Now()
	app.frameTime = time.Now()

	// Render loop
	for !app.win.ShouldClose() {
		// Starts measuring this frame
		app.frameRater.Start()

		// Updates frame start and time delta in context
		now := time.Now()
		app.frameDelta = now.Sub(app.frameTime)
		app.frameTime = now

		// Process root panel timers
		if app.Gui() != nil {
			app.Gui().TimerManager.ProcessTimers()
		}

		// Process application timers
		app.ProcessTimers()

		// Dispatch before render event
		app.Dispatch(OnBeforeRender, nil)

		// Renders the current scene and/or gui
		rendered, err := app.renderer.Render(app.camera)
		if err != nil {
			return err
		}

		// Poll input events and process them
		app.win.PollEvents()

		if rendered {
			app.win.SwapBuffers()
		}

		// Dispatch after render event
		app.Dispatch(OnAfterRender, nil)

		// Controls the frame rate and updates the FPS for the user
		app.frameRater.Wait()
		app.frameCount++
	}
	return nil
}

// Quit ends the application
func (app *Application) Quit() {

	app.win.SetShouldClose(true)
}

// OnWindowResize is default handler for window resize events.
func (app *Application) OnWindowResize(evname string, ev interface{}) {

	// Get window size and sets the viewport to the same size
	width, height := app.win.GetSize()
	app.gl.Viewport(0, 0, int32(width), int32(height))

	// Sets perspective camera aspect ratio
	aspect := float32(width) / float32(height)
	app.camPersp.SetAspect(aspect)

	// Sets the GUI root panel size to the size of the screen
	if app.guiroot != nil {
		app.guiroot.SetSize(float32(width), float32(height))
	}
}

// LoadAudioLibs try to load audio libraries
func (app *Application) LoadAudioLibs() error {

	// Try to load OpenAL
	err := al.Load()
	if err != nil {
		return err
	}

	// Opens default audio device
	app.audioDev, err = al.OpenDevice("")
	if app.audioDev == nil {
		return fmt.Errorf("Error: %s opening OpenAL default device", err)
	}

	// Checks for OpenAL effects extension support
	if al.IsExtensionPresent("ALC_EXT_EFX") {
		app.audioEFX = true
	}

	// Creates audio context with auxiliary sends
	var attribs []int
	if app.audioEFX {
		attribs = []int{al.MAX_AUXILIARY_SENDS, 4}
	}
	acx, err := al.CreateContext(app.audioDev, attribs)
	if err != nil {
		return fmt.Errorf("Error creating audio context:%s", err)
	}

	// Makes the context the current one
	err = al.MakeContextCurrent(acx)
	if err != nil {
		return fmt.Errorf("Error setting audio context current:%s", err)
	}
	app.audio = true
	app.log.Info("%s version: %s", al.GetString(al.Vendor), al.GetString(al.Version))

	// Ogg Vorbis support
	err = ov.Load()
	if err == nil {
		app.vorbis = true
		vorbis.Load()
		app.log.Info("%s", vorbis.VersionString())
	}
	return nil
}
