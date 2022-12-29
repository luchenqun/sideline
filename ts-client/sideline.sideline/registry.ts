import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateTask } from "./types/sideline/sideline/tx";
import { MsgSubmitTask } from "./types/sideline/sideline/tx";
import { MsgRegistEmployer } from "./types/sideline/sideline/tx";
import { MsgUndoneTask } from "./types/sideline/sideline/tx";
import { MsgDoTask } from "./types/sideline/sideline/tx";
import { MsgRegistDeveloper } from "./types/sideline/sideline/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/sideline.sideline.MsgCreateTask", MsgCreateTask],
    ["/sideline.sideline.MsgSubmitTask", MsgSubmitTask],
    ["/sideline.sideline.MsgRegistEmployer", MsgRegistEmployer],
    ["/sideline.sideline.MsgUndoneTask", MsgUndoneTask],
    ["/sideline.sideline.MsgDoTask", MsgDoTask],
    ["/sideline.sideline.MsgRegistDeveloper", MsgRegistDeveloper],
    
];

export { msgTypes }