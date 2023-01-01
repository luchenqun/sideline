import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgSuccessTask } from "./types/sideline/sideline/tx";
import { MsgSubmitTask } from "./types/sideline/sideline/tx";
import { MsgStartJudgeTask } from "./types/sideline/sideline/tx";
import { MsgCancelTask } from "./types/sideline/sideline/tx";
import { MsgDoTask } from "./types/sideline/sideline/tx";
import { MsgFailTask } from "./types/sideline/sideline/tx";
import { MsgRegistEmployer } from "./types/sideline/sideline/tx";
import { MsgCreateTask } from "./types/sideline/sideline/tx";
import { MsgUndoneTask } from "./types/sideline/sideline/tx";
import { MsgVoteTask } from "./types/sideline/sideline/tx";
import { MsgFeedbackTask } from "./types/sideline/sideline/tx";
import { MsgRegistDeveloper } from "./types/sideline/sideline/tx";
import { MsgJudgeTask } from "./types/sideline/sideline/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/sideline.sideline.MsgSuccessTask", MsgSuccessTask],
    ["/sideline.sideline.MsgSubmitTask", MsgSubmitTask],
    ["/sideline.sideline.MsgStartJudgeTask", MsgStartJudgeTask],
    ["/sideline.sideline.MsgCancelTask", MsgCancelTask],
    ["/sideline.sideline.MsgDoTask", MsgDoTask],
    ["/sideline.sideline.MsgFailTask", MsgFailTask],
    ["/sideline.sideline.MsgRegistEmployer", MsgRegistEmployer],
    ["/sideline.sideline.MsgCreateTask", MsgCreateTask],
    ["/sideline.sideline.MsgUndoneTask", MsgUndoneTask],
    ["/sideline.sideline.MsgVoteTask", MsgVoteTask],
    ["/sideline.sideline.MsgFeedbackTask", MsgFeedbackTask],
    ["/sideline.sideline.MsgRegistDeveloper", MsgRegistDeveloper],
    ["/sideline.sideline.MsgJudgeTask", MsgJudgeTask],
    
];

export { msgTypes }