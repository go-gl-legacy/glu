package glu

// #cgo pkg-config: glu
//
// #include <GL/glu.h>
//
import "C"
import "github.com/banthar/gl"
import "unsafe"

func Build2DMipmaps(target gl.GLenum, internalFormat int, width, height int, format gl.GLenum, data interface{}) int {
	t, p := gl.GetGLenumType(data)
	return int(C.gluBuild2DMipmaps(
		C.GLenum(target),
		C.GLint(internalFormat),
		C.GLsizei(width),
		C.GLsizei(height),
		C.GLenum(format),
		C.GLenum(t),
		p,
	))
}

func Perspective(fovy, aspect, zNear, zFar float64) {
	C.gluPerspective(
		C.GLdouble(fovy),
		C.GLdouble(aspect),
		C.GLdouble(zNear),
		C.GLdouble(zFar),
	)
}

func LookAt(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ float64) {
	C.gluLookAt(
		C.GLdouble(eyeX),
		C.GLdouble(eyeY),
		C.GLdouble(eyeZ),
		C.GLdouble(centerX),
		C.GLdouble(centerY),
		C.GLdouble(centerZ),
		C.GLdouble(upX),
		C.GLdouble(upY),
		C.GLdouble(upZ),
	)
}

func UnProject(winX, winY, winZ float64, model, proj, view unsafe.Pointer) (float64, float64, float64) {
	var ox, oy, oz C.GLdouble

	m := (*C.GLdouble)(model)
	p := (*C.GLdouble)(proj)
	v := (*C.GLint)(view)

	C.gluUnProject(
		C.GLdouble(winX),
		C.GLdouble(winY),
		C.GLdouble(winZ),
		m,
		p,
		v,
		&ox,
		&oy,
		&oz,
	)

	return float64(ox), float64(oy), float64(oz)
}

func NewQuadric() unsafe.Pointer {
	return unsafe.Pointer(C.gluNewQuadric())
}

func Sphere(q unsafe.Pointer, radius float32, slices, stacks int) {
	C.gluSphere((*[0]byte)(q), C.GLdouble(radius), C.GLint(slices), C.GLint(stacks))
}