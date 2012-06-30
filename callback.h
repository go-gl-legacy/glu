// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

#ifndef _CALLBACK_H_
#define _CALLBACK_H_

#include <GL/glu.h>

extern void goTessBeginData(GLenum type, void *polygon_data);
extern void goTessVertexData(void *vertex_data, void *polygon_data);
extern void goTessEndData(void *polygon_data);
extern void goTessErrorData(GLenum errno, void *polygon_data);
extern void goTessEdgeFlagData(GLboolean flag, void *polygon_data);
extern void goTessCombineData(void *coords, void *vertex_data,
                              void *weight, void *outData,
                              void *polygon_data);

void setGluTessCallback(GLUtesselator *tess, GLenum which);

#endif // _CALLBACK_H_
