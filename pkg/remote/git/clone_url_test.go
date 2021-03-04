package git_test

import (
	"github.com/tektoncd/pipeline/pkg/remote/git"
	"testing"
)

func TestCloneURL(t *testing.T) {
	testCases := []struct {
		server     string
		owner      string
		repository string
		want       string
	}{
		{
			server:     "github.com",
			owner:      "myorg",
			repository: "myrepo",
			want:       "https://github.com/myorg/myrepo.git",
		},
		{
			server:     "https://github.com",
			owner:      "myorg",
			repository: "myrepo",
			want:       "https://github.com/myorg/myrepo.git",
		},
		{
			server:     "https://github.com/",
			owner:      "myorg",
			repository: "myrepo",
			want:       "https://github.com/myorg/myrepo.git",
		},
	}

	for _, tc := range testCases {
		got := git.GitCloneURL(tc.server, tc.owner, tc.repository)

		if tc.want != got {
			t.Fatalf("GitCloneURL(%s, %s, %s) got %s ; want %s", tc.server, tc.owner, tc.repository, got, tc.want)
		}
	}
}
