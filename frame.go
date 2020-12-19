package main

import "image/color"

type frame [][]color.Color

func rotateX(f frame) frame {
	for i, row := range f {
		f[i] = append([]color.Color{row[len(row)-1]}, row[:len(row)-1]...)
	}

	return f
}

func rotateY(f frame) frame {
	f = append(frame{f[len(f)-1]}, f[:len(f)-1]...)
	return f
}
