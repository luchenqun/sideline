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

const MinConfirmSubmitHeight = 300 // 最小确认提交高度
const MinConfirmJudgeHeight = 600  // 最小确认仲裁高度
