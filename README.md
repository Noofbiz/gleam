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
    [ ] Title
      [X] Start
      [ ] Update
    [ ] Position
      [X] Start
      [ ] Update
    [ ] NotResizable
      [X] Start
      [ ] Update
    [ ] Menu
      [ ] Start
      [ ] Update
    [ ] Multiple Windows
      [X] Can open multiple windows
      [ ] Can update / change different windows
    [ ] Full Screen
      [X] Start
      [ ] Update
    [ ] Closing / Cleanup
  [ ] Can print an image.Image to the screen
  [ ] OpenGL surface
  [ ] Vulkan support
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
