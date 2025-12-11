package core

import (
	"log"
	"os"

	"github.com/go-gl/gl/v4.6-core/gl"
)

func CreateCustomShader(vertexFile string, fragmentFile string) uint32 {
	vertex := gl.CreateShader(gl.VERTEX_SHADER)
	vertexTxt, err := os.ReadFile(vertexFile)
	if err != nil {
		log.Fatal(PrefixErr + "Couldn't read Vertex Shader file: " + vertexFile)
	}
	vertexC, freeV := gl.Strs(string(vertexTxt))
	gl.ShaderSource(vertex, 1, vertexC, nil)
	freeV()
	gl.CompileShader(vertex)

	fragment := gl.CreateShader(gl.FRAGMENT_SHADER)
	fragmentTxt, err := os.ReadFile(fragmentFile)
	if err != nil {
		log.Fatal(PrefixErr + "Couldn't read Fragment Shader file: " + fragmentFile)
	}
	fragmentC, freeF := gl.Strs(string(fragmentTxt))
	gl.ShaderSource(fragment, 1, fragmentC, nil)
	freeF()
	gl.CompileShader(fragment)

	program := gl.CreateProgram()
	gl.AttachShader(program, vertex)
	gl.AttachShader(program, fragment)
	gl.LinkProgram(program)

	gl.DeleteShader(vertex)
	gl.DeleteShader(fragment)

	return program
}

// Private functions for render backend to call
func defaultShader() {

}
