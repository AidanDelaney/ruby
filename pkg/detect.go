package ruby_buildpack

import (
	"github.com/buildpacks/libcnb"
)

type Detect struct {
}

func (d Detect) Detect(dc libcnb.DetectContext) (libcnb.DetectResult, error) {
	return libcnb.DetectResult{
		Pass: true,
	}, nil
}
