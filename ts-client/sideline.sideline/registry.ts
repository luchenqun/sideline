import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRegistEmployer } from "./types/sideline/sideline/tx";
import { MsgSubmitTask } from "./types/sideline/sideline/tx";
import { MsgDoTask } from "./types/sideline/sideline/tx";
import { MsgCreateTask } from "./types/sideline/sideline/tx";
import { MsgRegistDeveloper } from "./types/sideline/sideline/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/sideline.sideline.MsgRegistEmployer", MsgRegistEmployer],
    ["/sideline.sideline.MsgSubmitTask", MsgSubmitTask],
    ["/sideline.sideline.MsgDoTask", MsgDoTask],
    ["/sideline.sideline.MsgCreateTask", MsgCreateTask],
    ["/sideline.sideline.MsgRegistDeveloper", MsgRegistDeveloper],
    
];

export { msgTypes }