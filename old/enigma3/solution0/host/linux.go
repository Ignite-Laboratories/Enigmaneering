//go:build linux

package host

/*
   #cgo LDFLAGS: -lX11
   #include <X11/Xlib.h>

   static void GetMouseCoordinates(int *x, int *y) {
       // Open a connection to the X window server
       Display *display = XOpenDisplay(NULL);
       if (display == NULL) return;

       // Get the root window of the default screen
       Window root_window = RootWindow(display, DefaultScreen(display));

       // Setup some output variables
       Window root, child;
       int root_x, root_y;
       int win_x, win_y;
       unsigned int mask;

       // Query the mouse position relative to the root window found above
       XQueryPointer(display, root_window, &root, &child, &root_x, &root_y, &win_x, &win_y, &mask);

       // Assign values to our input parameters
       *x = root_x;
       *y = root_y;

        // Be a good samaritan and close the connection to the server
       XCloseDisplay(display);
   }
*/
import "C"
import "fmt"

func GetCoordinates() (x int, y int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("failed to get mouse position: %v", r)
		}
	}()

	var cX, cY C.int
	C.GetMouseCoordinates(&cX, &cY)
	return int(cX), int(cY), nil
}
