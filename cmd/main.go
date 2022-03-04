package main

import (
	"github.com/buildpacks/libcnb"

	rubyBuildpack "github.com/AidanDelaney/ruby/pkg"
)

func main() {
	libcnb.Main(
		rubyBuildpack.Detect{},
		rubyBuildpack.Build{},
	)
}
