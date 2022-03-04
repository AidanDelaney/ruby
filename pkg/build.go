package ruby_buildpack

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/buildpacks/libcnb"
	"github.com/paketo-buildpacks/libpak/crush"
)

const (
	RUBY_URL = "https://s3-external-1.amazonaws.com/heroku-buildpack-ruby/heroku-18/ruby-2.5.1.tgz"
)

type Build struct {
}

func (b Build) Build(bc libcnb.BuildContext) (libcnb.BuildResult, error) {
	contributor := Contributor{}

	return libcnb.BuildResult{
		Layers: []libcnb.LayerContributor{contributor},
	}, nil
}

func fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	tgz, err := ioutil.ReadAll(resp.Body)
	return tgz, err
}

type Contributor struct {
}

func (c Contributor) Contribute(layer libcnb.Layer) (libcnb.Layer, error) {
	layer.LayerTypes.Launch = true

	archive, err := fetch(RUBY_URL)
	if err != nil {
		return layer, errors.New("cannot fetch ruby runtime")
	}

	if err := crush.ExtractTarGz(bytes.NewReader(archive), layer.Path, 0); err != nil {
		return layer, errors.New("cannot extract ruby runtime")
	}

	return layer, nil
}

func (c Contributor) Name() string {
	return "ruby"
}
