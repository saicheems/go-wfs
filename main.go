///////////////////////////////////////////////////////////////////////////////
//
// The MIT License (MIT)
// Copyright (c) 2018 Jivan Amara
// Copyright (c) 2018 Tom Kralidis
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
// OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE
// USE OR OTHER DEALINGS IN THE SOFTWARE.
//
///////////////////////////////////////////////////////////////////////////////

// go-wfs project main.go
package main

import (
	"github.com/go-spatial/go-wfs/provider"
	"github.com/go-spatial/go-wfs/server"
	"github.com/go-spatial/tegola/provider/gpkg"
)

func main() {
	// Instantiate a gpkg provider to send to the server.
	gpkgPath := "test-data/athens-osm-20170921.gpkg"
	gpkgConfig, err := gpkg.AutoConfig(gpkgPath)
	if err != nil {
		panic(err)
	}
	gpkgProvider, err := gpkg.NewTileProvider(gpkgConfig)
	if err != nil {
		panic(err)
	}

	p := provider.Provider{Tiler: gpkgProvider}

	server.StartServer(":9000", p)
}
