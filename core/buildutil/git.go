package buildutil

import (
	"errors"
	"fmt"

	"github.com/go-git/go-git/v5"
)

// GitClient has opinionated methods for working with Git during the BUILD step
//
// This client will print necessary information to the console
type GitClient struct {
	repo *git.Repository
}

// NewGitClient creates and returns a new buildutil.GitClient
func NewGitClient() (*GitClient, error) {
	repo, err := git.PlainOpen("../")
	if err != nil {
		return nil, err
	}

	client := &GitClient{repo}

	return client, nil
}

// IsClean checks that the Git workingtree is clean
func (g GitClient) IsClean() (bool, error) {
	fmt.Println("checking git status")

	worktree, err := g.repo.Worktree()
	if err != nil {
		return false, err
	}

	status, err := worktree.Status()
	if err != nil {
		return false, err
	}

	if !status.IsClean() {
		fmt.Printf(`Git worktree is not clean, cannot deploy with uncommitted changes.
			
%s`, status.String())

		return false, nil
	}

	return true, nil
}

// Pull attempts to Pull the latest changes. If there are no changes or if the changes are pulled
// successfully, true is returned. If there are any errors, false is returned.
func (g GitClient) Pull() (bool, error) {
	fmt.Println("pulling remote changes")
	worktree, err := g.repo.Worktree()
	if err != nil {
		return false, err
	}

	err = worktree.Pull(&git.PullOptions{})
	if err != nil {
		if errors.Is(err, git.NoErrAlreadyUpToDate) {
			fmt.Println("No pull: already up to date.")
		} else {
			return false, err
		}
	}

	return true, nil
}
