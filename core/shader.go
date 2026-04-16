package core

import (
	"log"
	"os"

	"github.com/go-gl/gl/v4.6-core/gl"
)

var DefaultShaderPr uint32

func CreateCustomShader(vertexFile string, fragmentFile string) {
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
}

// Loads default shader
func defaultShader() {
	vertex := gl.CreateShader(gl.VERTEX_SHADER)
	vertexC, freeV := gl.Strs(defVertexSh)
	gl.ShaderSource(vertex, 1, vertexC, nil)
	freeV()
	gl.CompileShader(vertex)

	fragment := gl.CreateShader(gl.FRAGMENT_SHADER)
	fragmentC, freeF := gl.Strs(defFragmentSh)
	gl.ShaderSource(fragment, 1, fragmentC, nil)
	freeF()
	gl.CompileShader(fragment)

	program := gl.CreateProgram()
	gl.AttachShader(program, vertex)
	gl.AttachShader(program, fragment)
	gl.LinkProgram(program)

	gl.DeleteShader(vertex)
	gl.DeleteShader(fragment)

	DefaultShaderPr = program
}

//Private function for material manipulation in render runtime will make some of them public if they are useful for user

var defVertexSh = `
#version 330 core
layout (location = 0) in vec3 aPos;
layout (location = 1) in vec2 aTexCoord;

out vec2 TexCoord;

uniform mat4 model;
uniform mat4 view;
uniform mat4 projection;

void main()
{
    gl_Position = projection * view * model * vec4(aPos, 1.0);
    TexCoord = aTexCoord;
}
` + "\x00"

var defFragmentSh = `
#version 330 core
out vec4 FragColor;
in vec2 TexCoord;

uniform sampler2D Texture;

void main()
{
    FragColor = texture(Texture, TexCoord);
} 
` + "\x00"
