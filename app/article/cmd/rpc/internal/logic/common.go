package logic

const (
	StateDraft     int32 = 0 // 草稿
	StatePending   int32 = 1 // 待审核
	StatePublished int32 = 2 // 审核成功
	StateRejected  int32 = 3 // 审核失败
)
