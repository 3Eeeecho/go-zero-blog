package state

type ArticleState int

const (
	Draft ArticleState = iota
	Pending
	Approved
	Rejectd
)
