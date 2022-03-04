package ruby_integration_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

const rubyBuildpack = "../"

func TestRubyIntegration(t *testing.T) {
	suite := spec.New("Ruby Integration", spec.Sequential(), spec.Report(report.Terminal{}))
	suite("ruby", testRuby)
	suite.Run(t)
}
