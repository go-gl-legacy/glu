package glu

// #cgo pkg-config: glu
//
// #include <GL/glu.h>
//
import "C"
import "github.com/banthar/gl"

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

func UnProject(winX, winY, winZ float64, model, proj [16]float64, view [4]float64) (objX, objY, objZ float64) {
	ox := (*C.GLdouble)(&objX)
	oy := (*C.GLdouble)(&objY)
	oz := (*C.GLdouble)(&objX)

	m := [16]C.GLdouble{C.GLdouble(model[0])}
	p := [16]C.GLdouble{C.GLdouble(proj[0])}
	v := [4]C.GLint{C.GLint(view[0])}
	for i := 0; i < 16; i++ {
		m[i] = C.GLdouble(model[i])
		p[i] = C.GLdouble(proj[i])

		if i < 4 {
			v[i] = C.GLint(view[i])
		}
	}

	C.gluUnProject(
		C.GLdouble(winX),
		C.GLdouble(winY),
		C.GLdouble(winZ),
		&m[0],
		&p[0],
		&v[0],
		ox,
		oy,
		oz,
	)

	return float64(*ox), float64(*oy), float64(*oz)
}