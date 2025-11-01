package branches

type Branch struct {
	CurrentCommit string `json:"current_commit"`
}

// filepath: ./.jit/branches/[branchname] -> holds the branch data
