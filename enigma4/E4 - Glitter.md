# `E4 - Glitter`
### `Alex Petz, Ignite Laboratories, June 2025`

---

### Resonant Rendering
Oh yay, another rendering framework!  Right?

Well, not exactly - I want _you_ to bring your own rendering pipeline!  Personally, I love OpenGL.  However,
the concept of _resonant rendering_ is completely separate of how you actually render anything.  You are provided
with a _renderable_ interface which you can implement, which an engine impulses to emit a frame.  The details
of when potential frames are emitted are up to you, allowing you to fully render whatever you like at any
desired pace and then composite the frame data together at a separate time interval.  Throttling the framerate
with a global variable is possible using a neural impulse engine, for example.  

Everything can hold a buffer of what the compositor should render when ready - but that buffer can intelligently
only perform render updates when its neural potential returns _true_.  By allowing these entities to activate
intelligently, a stable system can be constructed to maximize efficient operation.

OpenGL provides us with a means of passing frame buffer handles around, but the idea is often shared amongst
other libraries as well - thus, if you wish to integrate your own pipeline, you'll only have to build the
initialization logic necessary to instantiate our structures.  