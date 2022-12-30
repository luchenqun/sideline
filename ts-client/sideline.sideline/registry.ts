import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgSubmitTask } from "./types/sideline/sideline/tx";
import { MsgUndoneTask } from "./types/sideline/sideline/tx";
import { MsgRegistDeveloper } from "./types/sideline/sideline/tx";
import { MsgCreateTask } from "./types/sideline/sideline/tx";
import { MsgFailTask } from "./types/sideline/sideline/tx";
import { MsgRegistEmployer } from "./types/sideline/sideline/tx";
import { MsgDoTask } from "./types/sideline/sideline/tx";
import { MsgSuccessTask } from "./types/sideline/sideline/tx";
import { MsgVoteTask } from "./types/sideline/sideline/tx";
import { MsgCancelTask } from "./types/sideline/sideline/tx";
import { MsgStartJudgeTask } from "./types/sideline/sideline/tx";
import { MsgJudgeTask } from "./types/sideline/sideline/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/sideline.sideline.MsgSubmitTask", MsgSubmitTask],
    ["/sideline.sideline.MsgUndoneTask", MsgUndoneTask],
    ["/sideline.sideline.MsgRegistDeveloper", MsgRegistDeveloper],
    ["/sideline.sideline.MsgCreateTask", MsgCreateTask],
    ["/sideline.sideline.MsgFailTask", MsgFailTask],
    ["/sideline.sideline.MsgRegistEmployer", MsgRegistEmployer],
    ["/sideline.sideline.MsgDoTask", MsgDoTask],
    ["/sideline.sideline.MsgSuccessTask", MsgSuccessTask],
    ["/sideline.sideline.MsgVoteTask", MsgVoteTask],
    ["/sideline.sideline.MsgCancelTask", MsgCancelTask],
    ["/sideline.sideline.MsgStartJudgeTask", MsgStartJudgeTask],
    ["/sideline.sideline.MsgJudgeTask", MsgJudgeTask],
    
];

export { msgTypes }