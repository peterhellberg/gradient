# gradient

[![Build Status](https://travis-ci.org/peterhellberg/gradient.svg?branch=master)](https://travis-ci.org/peterhellberg/gradient)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://pkg.go.dev/github.com/peterhellberg/gradient)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/gradient#license-mit)

## Installation

    go get -u github.com/peterhellberg/gradient

## Usage

```go
package main

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/peterhellberg/gradient"
)

func main() {
	v := gradient.NewVertical(200, 100, []gradient.Stop{
		{0.0, color.NRGBA{0, 0, 255, 255}},
		{0.5, color.NRGBA{128, 0, 128, 255}},
		{1.0, color.NRGBA{255, 0, 0, 255}},
	})

	h := gradient.NewHorizontal(200, 100, []gradient.Stop{
		{0.0, color.NRGBA{0, 0, 0, 255}},
		{1.0, color.NRGBA{0, 0, 0, 0}},
	})

	saveImage(v, "/tmp/vertical-gradient.png")
	saveImage(h, "/tmp/horizontal-gradient.png")
}

func saveImage(m image.Image, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, m)
}
```

## Results

![vertical-gradient.png](https://assets.c7.se/viz/vertical-gradient.png)
![horizontal-gradient.png](https://assets.c7.se/viz/horizontal-gradient.png)

## License (MIT)

Copyright (c) 2015-2023 [Peter Hellberg](https://c7.se)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:

> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
