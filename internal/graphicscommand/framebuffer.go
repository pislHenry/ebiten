// Copyright 2014 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graphicscommand

import (
	"github.com/hajimehoshi/ebiten/internal/opengl"
)

// framebuffer is a wrapper of OpenGL's framebuffer.
type framebuffer struct {
	native    opengl.Framebuffer
	proMatrix []float32
	width     int
	height    int
}

// newFramebufferFromTexture creates a framebuffer from the given texture.
func newFramebufferFromTexture(texture opengl.Texture, width, height int) (*framebuffer, error) {
	native, err := opengl.GetContext().NewFramebuffer(texture)
	if err != nil {
		return nil, err
	}
	return &framebuffer{
		native: native,
		width:  width,
		height: height,
	}, nil
}

// newScreenFramebuffer creates a framebuffer for the screen.
func newScreenFramebuffer(width, height int) *framebuffer {
	return &framebuffer{
		native: opengl.GetContext().ScreenFramebuffer(),
		width:  width,
		height: height,
	}
}

// projectionMatrix returns a projection matrix of the framebuffer.
//
// A projection matrix converts the coodinates on the framebuffer
// (0, 0) - (viewport width, viewport height)
// to the normalized device coodinates (-1, -1) - (1, 1) with adjustment.
func (f *framebuffer) projectionMatrix() []float32 {
	if f.proMatrix != nil {
		return f.proMatrix
	}
	f.proMatrix = opengl.OrthoProjectionMatrix(0, f.width, 0, f.height)
	return f.proMatrix
}