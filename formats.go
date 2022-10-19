package main

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

type Format struct {
	Decoder func(io.Reader) (image.Image, error)
	Encoder func(io.Writer, image.Image) error
}

var decoders map[string]Format = map[string]Format{
	".jpg":  {Decoder: jpeg.Decode, Encoder: jpegEncode},
	".jpeg": {Decoder: jpeg.Decode, Encoder: jpegEncode},
	".png":  {Decoder: png.Decode, Encoder: jpegEncode},
}

func jpegEncode(w io.Writer, i image.Image) error {
	return jpeg.Encode(w, i, &jpeg.Options{Quality: 50})
}
