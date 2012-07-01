package main

// This demonstrates basic usage of the tessalator.

import (
   "github.com/banthar/Go-SDL/sdl"
   "github.com/banthar/gl"
   "github.com/0xfaded/glu"
   //"github.com/banthar/glu" restore this after merge
)


// Test shape is a square with a square hole inside.
var OuterContour [4][3]float64 = [4][3]float64{[3]float64{-2,  2, 0},
                                               [3]float64{-2, -2, 0},
                                               [3]float64{ 2, -2, 0},
                                               [3]float64{ 2,  2, 0}}

var InnerContour [4][3]float64 = [4][3]float64{[3]float64{-1,  1, 0},
                                               [3]float64{ 1,  1, 0},
                                               [3]float64{ 1, -1, 0},
                                               [3]float64{-1, -1, 0}}

// Pentagram with crossing edges. Invokes the combine callback.
var StarContour [5][3]float64 = [5][3]float64{[3]float64{ 0,  2, 0},
                                              [3]float64{-2, -2, 0},
                                              [3]float64{ 2,  0, 0},
                                              [3]float64{-2,  0, 0},
                                              [3]float64{ 2, -2, 0}}
func main() {

	sdl.Init(sdl.INIT_VIDEO)

	var screen = sdl.SetVideoMode(640, 480, 32, sdl.OPENGL)

	if screen == nil {
		panic("sdl error")
	}

	if gl.Init() != 0 {
		panic("gl error")
	}

	gl.MatrixMode(gl.PROJECTION)

	gl.Viewport(0, 0, int(screen.W), int(screen.H))
	gl.LoadIdentity()
	gl.Ortho(-5, 5, -2.5, 2.5, -1.0, 1.0)

	gl.ClearColor(0, 0, 0, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)

	tess := glu.NewTess()


	tess.SetBeginCallback   (tessBeginHandler   )
	tess.SetVertexCallback  (tessVertexHandler  )
	tess.SetEndCallback     (tessEndHandler     )
	tess.SetErrorCallback   (tessErrorHandler   )
	tess.SetEdgeFlagCallback(tessEdgeFlagHandler)
	tess.SetCombineCallback (tessCombineHandler )

	tess.Normal(0, 0, 1)

	var running = true

	for running {

		gl.MatrixMode(gl.MODELVIEW)
		gl.LoadIdentity()
		gl.Translated(-2.5, 0, 0)

		tess.BeginPolygon(nil)
		tess.BeginContour()

		for v := range OuterContour {
			tess.Vertex(OuterContour[v], &OuterContour[v])
		}

		tess.EndContour()
		tess.BeginContour()

		for v := range InnerContour {
			tess.Vertex(InnerContour[v], &InnerContour[v])
		}

		tess.EndContour()
		tess.EndPolygon()

		gl.Translated(5, 0, 0)

		tess.BeginPolygon(nil)
		tess.BeginContour()

		for v := range StarContour {
			tess.Vertex(StarContour[v], &StarContour[v])
		}

		tess.EndContour()
		tess.EndPolygon()

		sdl.GL_SwapBuffers()
		sdl.Delay(25)
	}

	tess.Delete()
	sdl.Quit()

}

func tessBeginHandler(tessType gl.GLenum, polygonData interface{}) {
	gl.Begin(tessType)
}

func tessVertexHandler(vertexData interface{}, polygonData interface{}) {
	v := vertexData.(*[3]float64)
	gl.Vertex3d((*v)[0], (*v)[1], (*v)[2])
}

func tessEndHandler(polygonData interface{}) {
	gl.End()
}

func tessErrorHandler(errno gl.GLenum, polygonData interface{}) {
}

func tessEdgeFlagHandler(flag bool, polygonData interface{}) {
}

func tessCombineHandler(coords      [3]float64,
                        vertexData  [4]interface{},
                        weight      [4]float32,
                        polygonData interface{}) (outData interface{}) {
	return &[3]float64{coords[0], coords[1], coords[2]}
}


