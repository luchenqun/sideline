package types

const (
	TaskStatusCreated = iota
	TaskStatusDoing
	TaskStatusSubmitted
	TaskStatusUndone
	TaskStatusSuccess
	TaskStatusFail
	TaskStatusJudging
	TaskStatusEmployerWin
	TaskStatusDeveloperWin
	TaskStatusCanceled
)

const SidelineDenom = "wrmb"
