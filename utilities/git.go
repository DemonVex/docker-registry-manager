package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os/exec"
)

// Ref contains the response from the github ref request
type Ref struct {
	Object struct {
		Sha  string `json:"sha"`
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"object"`
	Ref string `json:"ref"`
	URL string `json:"url"`
}

/* Github oriented functions
 *
 *
 *
 */

// GetRemoteBranchSHA returns the commit SHA of the remote branch from github
func GetRemoteBranchSHA(branch string) (string, error) {

	// Get the branches current SHA
	url := "https://api.github.com/repos/stefannaglee/docker-registry-manager/git/refs/heads/" + branch
	res, err := http.Get(url)
	if err != nil {
		Log.Error(err)
		Log.Error("Failed to query githubs API for the remote branch with the url '" + url + "''")
		return "", err
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		Log.Error(err)
		return "", err
	}

	r := Ref{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		Log.Error(err)
		return string(body), err
	}

	return r.Object.Sha, nil
}

/* Local git oriented functions
 *
 *
 *
 */

// GetAppBranch returns the current branch name
func GetAppBranch() (string, error) {
	cmdArgs := []string{"rev-parse", "--abbrev-ref", "HEAD"}
	branchName, err := exec.Command("/usr/bin/git", cmdArgs...).Output()
	return string(branchName), err
}

// GetAppSHA returns the current SHA location
func GetAppSHA() (string, error) {
	cmdArgs := []string{"rev-parse", "HEAD"}
	branchSHA, err := exec.Command("/usr/bin/git", cmdArgs...).Output()
	return string(branchSHA), err
}

// GetBaseSHA returns the current base SHA location
// http://stackoverflow.com/questions/3258243/check-if-pull-needed-in-git
func GetBaseSHA() (string, error) {
	cmdArgs := []string{"merge-base", "@", "@{u}"}
	baseSHA, err := exec.Command("/usr/bin/git", cmdArgs...).Output()
	return string(baseSHA), err
}

// RemoteUpdate updates with origin
func RemoteUpdate() error {
	cmdArgs := []string{"remote", "update"}
	err := exec.Command("/usr/bin/git", cmdArgs...).Run()
	return err
}
