package ruby_integration_test

import (
	"path/filepath"
	"testing"

	"github.com/paketo-buildpacks/occam"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testRuby(t *testing.T, when spec.G, it spec.S) {
	var (
		Expect     = NewWithT(t).Expect
		Eventually = NewWithT(t).Eventually

		pack   occam.Pack
		docker occam.Docker

		name   string
		source string
	)

	it.Before(func() {
		pack = occam.NewPack()
		docker = occam.NewDocker()

		var err error
		name, err = occam.RandomName()
		Expect(err).To(BeNil())
		source, err = occam.Source(filepath.Join("testData", "default"))
		Expect(err).NotTo(HaveOccurred())
	})

	it("should build an image containing ruby", func() {
		image, _, err := pack.WithNoColor().Build.
			WithBuildpacks(rubyBuildpack).
			Execute(name, source)
		Expect(err).To(BeNil())

		container, err := docker.Container.Run.
			WithCommand("ruby --version").
			Execute(image.ID)
		Expect(err).NotTo(HaveOccurred())

		Eventually(func() string {
			cLogs, err := docker.Container.Logs.Execute(container.ID)
			Expect(err).NotTo(HaveOccurred())
			return cLogs.String()
		}, "20s").Should(ContainSubstring("ruby 2.5"))
	})
}
