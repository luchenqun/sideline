package types

// sideline module event types
const (
	EventTypeCancelTask      = "cancel_task"
	EventTypeCreateTask      = "create_task"
	EventTypeDoTask          = "do_task"
	EventTypeFailTask        = "fail_task"
	EventTypeJudgeTask       = "judge_task"
	EventTypeStartJudgeTask  = "start_judge_task"
	EventTypeRegistDeveloper = "regist_developer"
	EventTypeRegistEmployer  = "regist_employer"
	EventTypeSubmitTask      = "submit_task"
	EventTypeSuccessTask     = "success_task"
	EventTypeUndoneTask      = "undone_task"
	EventTypeVoteTask        = "vote_task"
	EventTypeFeedbackTask    = "feedback_task"

	AttributeKeyDeveloper              = "developer"
	AttributeKeyEmployer               = "employer"
	AttributeKeyRegistrationFee        = "registration_fee"
	AttributeKeyRegistrationAverageFee = "registration_average_fee"
	AttributeKeyTaskId                 = "task_id"
)
