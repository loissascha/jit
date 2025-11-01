package commits

type Commit struct {
	Parent string `json:"parent"`
	Author string `json:"author"`
}

// TODO: find a way to store a reference (hash or whatever) to each file that is currently in this commit. so you can find all the files

// filepath: ./.jit/commits/[commithash] -> holds the commit data (parent, author, ...)
