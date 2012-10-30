// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifndef _CALLBACK_H_
#define _CALLBACK_H_

#include <GL/glu.h>

//Apple's glu.h defines its function pointers a little differently, so this tries to fix things up:
#ifdef __APPLE__
#ifndef GLAPIENTRY
typedef GLvoid (*_GLUfuncptr)(void);
#endif
#endif

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
