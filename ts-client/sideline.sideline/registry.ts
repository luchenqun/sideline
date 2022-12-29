import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateTask } from "./types/sideline/sideline/tx";
import { MsgRegistEmployer } from "./types/sideline/sideline/tx";
import { MsgRegistDeveloper } from "./types/sideline/sideline/tx";
import { MsgDoTask } from "./types/sideline/sideline/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/sideline.sideline.MsgCreateTask", MsgCreateTask],
    ["/sideline.sideline.MsgRegistEmployer", MsgRegistEmployer],
    ["/sideline.sideline.MsgRegistDeveloper", MsgRegistDeveloper],
    ["/sideline.sideline.MsgDoTask", MsgDoTask],
    
];

export { msgTypes }