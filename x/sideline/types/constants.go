package types

const (
	TaskStatusCreated = iota
	TaskStatusDoing
	TaskStatusSubmited
	TaskStatusUndone
	TaskStatusSuccess
	TaskStatusFail
	TaskStatusJudging
	TaskStatusEmployerWin
	TaskStatusDeveloperWin
	TaskStatusCanced
)

const MinConfirmSubmitHeight = 100 // 最小确认提交高度
