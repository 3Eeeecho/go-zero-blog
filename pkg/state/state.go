package state

type ArticleState int32

const (
	Draft ArticleState = iota
	Pending
	Approved
	Rejectd
)
