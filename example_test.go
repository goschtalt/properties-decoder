// SPDX-FileCopyrightText: 2022 Weston Schmidt <weston_schmidt@alumni.purdue.edu>
// SPDX-License-Identifier: Apache-2.0
//go:build !windows

package properties_test

import (
	"fmt"
	"strings"
	"testing/fstest"

	"github.com/goschtalt/goschtalt"
	_ "github.com/goschtalt/properties-decoder"
)

const text = `# example file
Example.Version = 1
Example.Colors.0 = red
Example.Colors.1 = green
Example.Colors.2 = blue`

func Example() {
	fs := fstest.MapFS{
		"example.properties": &fstest.MapFile{
			Data: []byte(text),
			Mode: 0644,
		},
	}
	// Normally, you use something like os.DirFS("/etc/program")
	g, err := goschtalt.New(goschtalt.AddDir(fs, "."))
	if err != nil {
		panic(err)
	}

	err = g.Compile()
	if err != nil {
		panic(err)
	}

	var cfg struct {
		Example struct {
			Version int
			Colors  []string
		}
	}

	err = g.Unmarshal("", &cfg)
	if err != nil {
		panic(err)
	}

	fmt.Println("example")
	fmt.Printf("    version = %d\n", cfg.Example.Version)
	fmt.Printf("    colors  = [ %s ]\n", strings.Join(cfg.Example.Colors, ", "))

	// Output:
	// example
	//     version = 1
	//     colors  = [ red, green, blue ]
}
