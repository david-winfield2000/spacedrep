package storage

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func runGit(dir string, args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git failed: %v\n%s", err, string(out))
	}
	return nil
}

func SyncRepo(repoURL string) error {
	dir, err := gitRepoPath()
	if err != nil {
		return err
	}

	// check if repo exists
	if err := runGit(dir, "rev-parse", "--is-inside-work-tree"); err == nil {
		return runGit(dir, "pull")
	}

	return runGit(filepath.Dir(dir), "clone", repoURL, dir)
}

func CommitPush() error {
	dir, err := gitRepoPath()
	if err != nil {
		return err
	}

	if err := runGit(dir, "add", "-A"); err != nil {
		return err
	}
	if err := runGit(dir, "commit", "-m", "synced problems"); err != nil {
		return err
	}
	return runGit(dir, "push")
}
