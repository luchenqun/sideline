import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRegistEmployer } from "./types/sideline/sideline/tx";
import { MsgDoTask } from "./types/sideline/sideline/tx";
import { MsgSubmitTask } from "./types/sideline/sideline/tx";
import { MsgFailTask } from "./types/sideline/sideline/tx";
import { MsgSuccessTask } from "./types/sideline/sideline/tx";
import { MsgCreateTask } from "./types/sideline/sideline/tx";
import { MsgRegistDeveloper } from "./types/sideline/sideline/tx";
import { MsgUndoneTask } from "./types/sideline/sideline/tx";
import { MsgStartJudgeTask } from "./types/sideline/sideline/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/sideline.sideline.MsgRegistEmployer", MsgRegistEmployer],
    ["/sideline.sideline.MsgDoTask", MsgDoTask],
    ["/sideline.sideline.MsgSubmitTask", MsgSubmitTask],
    ["/sideline.sideline.MsgFailTask", MsgFailTask],
    ["/sideline.sideline.MsgSuccessTask", MsgSuccessTask],
    ["/sideline.sideline.MsgCreateTask", MsgCreateTask],
    ["/sideline.sideline.MsgRegistDeveloper", MsgRegistDeveloper],
    ["/sideline.sideline.MsgUndoneTask", MsgUndoneTask],
    ["/sideline.sideline.MsgStartJudgeTask", MsgStartJudgeTask],
    
];

export { msgTypes }