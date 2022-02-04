package githubrelease

import (
	"path"

	"github.com/updatecli/updatecli/pkg/core/tmp"
	"github.com/updatecli/updatecli/pkg/plugins/github"
	"github.com/updatecli/updatecli/pkg/plugins/version"
)

// GitHubRelease defines a resource of kind "githubrelease"
type GitHubRelease struct {
	ghHandler     github.GithubHandler
	versionFilter version.Filter
}

// New returns a new valid GitHubRelease object.
func New(newSpec github.Spec) (GitHubRelease, error) {
	if newSpec.Directory == "" {
		newSpec.Directory = path.Join(tmp.Directory, newSpec.Owner, newSpec.Repository)
	}

	if newSpec.URL == "" {
		newSpec.URL = "github.com"
	}

	newHandler, err := github.New(newSpec)
	if err != nil {
		return GitHubRelease{}, err
	}

	return GitHubRelease{
		ghHandler:     &newHandler,
		versionFilter: newHandler.Spec.VersionFilter,
	}, nil
}