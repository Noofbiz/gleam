# gleam

Gleam is a platform native window manager written in go. Inspired by glfw and
shiny, gleam aims to be a library for creating platform native image.Image, opengl, and vulkan
windows as well as getting input for your applications.

It is currently in development and probably won't be ready for production use
for quite some time, if ever. It's mostly just for fun right now.

Target platforms are:

[X] Cocoa
  [X] Window Creation
    [X] Width, Height
    [X] Title
      [X] Start
      [X] Update
    [X] Position
      [X] Start
      [X] Update
    [X] NotResizable
      [X] Start
      [X] Update
    [ ] Menu
      [ ] Start
      [ ] Update
    [X] Multiple Windows
      [X] Can open multiple windows
      [X] Can update / change different windows
    [X] Full Screen
      [X] Start
      [X] Update
    [X] Closing / Cleanup
    [ ] Contexts
      [ ] OpenGL
      [ ] Vulkan / Molten
      [ ] Metal
      [ ] Image
  [ ] Keyboard
  [ ] Mouse
  [ ] Joystick
[ ] Windows
  [ ] Window Creation
  [ ] Can print an image.Image to the screen
  [ ] OpenGL surface
  [ ] Vulkan support
  [ ] Keyboard
  [ ] Mouse
  [ ] Joystick
[ ] X11
  [ ] Window Creation
  [ ] Can print an image.Image to the screen
  [ ] OpenGL surface
  [ ] Vulkan support
  [ ] Keyboard
  [ ] Mouse
  [ ] Joystick
[ ] Wayland
  [ ] Window Creation
  [ ] OpenGL surface
  [ ] Vulkan support
  [ ] Keyboard
  [ ] Mouse
  [ ] Joystick
