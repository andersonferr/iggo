#include "IggoUtil.h"

#include <stdio.h>
#include <string.h>

#include <sys/time.h>
#include <sys/types.h>
#include <sys/select.h>

int CanReadFromDisplay(Display *display)
{
    int x11_fd = ConnectionNumber(display);
    fd_set in_fds;
    FD_ZERO(&in_fds);
    FD_SET(x11_fd, &in_fds);
    struct timeval tv;
    tv.tv_sec = 0;
    tv.tv_usec = 20;

    int num_ready_fds = select(x11_fd + 1, &in_fds, NULL, NULL, &tv);

    if (num_ready_fds > 0)
    {
        return 1;
    }

    return 0;
}

int IggoReadEvents(Display *display, XEvent *buffer, int size)
{
    int i = 0;
    do
    {
        XNextEvent(display, &buffer[i]);
        i++;
    } while (i < size && XPending(display));

    return i;
}

void IggoDrawImage(
    Display *display,
    Drawable drawable,
    GC gc,
    unsigned char *image32,
    int width, int height, int stride)
{
    int depth = 24;      // works fine with depth = 24
    int bitmap_pad = 32; // 32 for 24 and 32 bpp, 16, for 15&16

    XImage im = (XImage){
        .width = width,
        .height = height,
        .xoffset = 0,
        .format = ZPixmap,
        .data = image32,
        .byte_order = MSBFirst,
        .bitmap_pad = 32,
        .depth = depth,
        .bits_per_pixel = 32,
        .bytes_per_line = stride,
        .depth = depth,
    };

    if (!XInitImage(&im))
    {
        printf("imagem inconsistente\n");
    }

    XPutImage(display, drawable, gc, &im, 0, 0, 0, 0, width, height); // 0, 0, 0, 0 are src x,y and dst x,y
    XSync(display, False);
}


static int compare(const void *a, const void *b)
{
  return strcmp(*(char **) a, *(char **) b);
}

/*
IggoCreateWindow create a window and return informations about it.
*/
IggoCreateWindowStruct IggoCreateWindow(Display *display, int width, int height)
{
    IggoCreateWindowStruct r;
    memset(&r, 0, sizeof r);

    int screen_number = DefaultScreen(display);
    Window root = RootWindow(display, screen_number);
    r.depth = DefaultDepth(display, screen_number);
    Visual *visual = DefaultVisual(display, screen_number);

    printf("screen number: %d\n", screen_number);
    printf("depth: %d\n", r.depth);
    //printf("");

    r.window = XCreateWindow(
        display,
        root,
        0, 0, width, height,
        0, //border width
        r.depth,
        InputOutput,
        visual,
        0, NULL);

    XSelectInput(
        display, r.window,
        ButtonPressMask |
            ExposureMask |
            KeyPressMask |
            StructureNotifyMask);

    Atom wm_delete_window = XInternAtom(display, "WM_DELETE_WINDOW", True);
    XSetWMProtocols(display, r.window, &wm_delete_window, 1);

    r.gc = XCreateGC(display, r.window, 0, NULL);
    XSetBackground(display, r.gc, WhitePixel(display, screen_number));
    XSetForeground(display, r.gc, BlackPixel(display, screen_number));

    return r;
}
