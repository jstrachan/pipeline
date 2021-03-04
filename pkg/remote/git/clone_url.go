package git

import "strings"

// GitCloneURL creates a git clone URL for a server
func GitCloneURL(server, owner, repo string) string {
	answer := server
	if !strings.Contains(answer, ":") {
		answer = "https://" + answer
	}
	if !strings.HasSuffix(answer, "/") {
		answer += "/"
	}
	return answer + owner + "/" + repo + ".git"
}
