package git

import (
	"github.com/jenkins-x/go-scm/scm"
	"github.com/pkg/errors"
	"strings"
)

// GitURI for parsing the git URIs in uses
type GitURI struct {
	Owner      string
	Repository string
	Path       string
	SHA        string
}

// ParseGitURI parses git source URIs or returns nil if its not a valid URI
//
// handles strings of the form "owner/repository(/path)@sha"
func ParseGitURI(text string) (*GitURI, error) {
	idx := strings.Index(text, "@")
	if idx < 0 {
		return nil, nil
	}

	sha := text[idx+1:]
	if sha == "" {
		return nil, errors.Errorf("missing version, branch or sha after the '@' character in the git URI %s", text)
	}

	names := text[0:idx]
	parts := strings.SplitN(names, "/", 3)

	path := ""
	switch len(parts) {
	case 0, 1:
		return nil, errors.Errorf("expecting format 'owner/repository(/path)@sha' but got git URI %s", text)
	case 3:
		path = parts[2]
	}
	owner := parts[0]

	return &GitURI{
		Owner:      owner,
		Repository: parts[1],
		Path:       path,
		SHA:        sha,
	}, nil
}

// String returns the string representation of the git URI
func (u *GitURI) String() string {
	path := scm.Join(u.Owner, u.Repository)
	if u.Path != "" {
		if !strings.HasPrefix(u.Path, "/") {
			path += "/"
		}
		path += u.Path
	}
	path = strings.TrimSuffix(path, "/")
	sha := u.SHA
	if sha == "" {
		sha = "head"
	}
	prefix := ""
	return prefix + path + "@" + sha
}
