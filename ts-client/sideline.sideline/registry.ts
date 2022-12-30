import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgSubmitTask } from "./types/sideline/sideline/tx";
import { MsgRegistDeveloper } from "./types/sideline/sideline/tx";
import { MsgCreateTask } from "./types/sideline/sideline/tx";
import { MsgStartJudgeTask } from "./types/sideline/sideline/tx";
import { MsgRegistEmployer } from "./types/sideline/sideline/tx";
import { MsgSuccessTask } from "./types/sideline/sideline/tx";
import { MsgFailTask } from "./types/sideline/sideline/tx";
import { MsgDoTask } from "./types/sideline/sideline/tx";
import { MsgCancelTask } from "./types/sideline/sideline/tx";
import { MsgJudgeTask } from "./types/sideline/sideline/tx";
import { MsgUndoneTask } from "./types/sideline/sideline/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/sideline.sideline.MsgSubmitTask", MsgSubmitTask],
    ["/sideline.sideline.MsgRegistDeveloper", MsgRegistDeveloper],
    ["/sideline.sideline.MsgCreateTask", MsgCreateTask],
    ["/sideline.sideline.MsgStartJudgeTask", MsgStartJudgeTask],
    ["/sideline.sideline.MsgRegistEmployer", MsgRegistEmployer],
    ["/sideline.sideline.MsgSuccessTask", MsgSuccessTask],
    ["/sideline.sideline.MsgFailTask", MsgFailTask],
    ["/sideline.sideline.MsgDoTask", MsgDoTask],
    ["/sideline.sideline.MsgCancelTask", MsgCancelTask],
    ["/sideline.sideline.MsgJudgeTask", MsgJudgeTask],
    ["/sideline.sideline.MsgUndoneTask", MsgUndoneTask],
    
];

export { msgTypes }