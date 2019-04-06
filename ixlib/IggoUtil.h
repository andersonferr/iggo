#ifndef _IGGO_H_
#define _IGGO_H_

#include <stdio.h>
#include <stdlib.h>
#include <X11/Xlib.h>
#include <X11/Xutil.h>

void IggoDrawPath(Display *display, Drawable d, GC gc, int *path);
int IggoReadEvents(Display *display, XEvent *buffer, int size);
void IggoDrawImage(Display *display, Drawable drawable, GC gc, unsigned char *image32, int width, int height, int bytes_per_line);

typedef struct
{
    Window window;
    GC gc;
    int depth;
} IggoCreateWindowStruct;

IggoCreateWindowStruct IggoCreateWindow(Display *display, int width, int height);

#endif