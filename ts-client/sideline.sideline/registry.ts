import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRegistDeveloper } from "./types/sideline/sideline/tx";
import { MsgCreateTask } from "./types/sideline/sideline/tx";
import { MsgRegistEmployer } from "./types/sideline/sideline/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/sideline.sideline.MsgRegistDeveloper", MsgRegistDeveloper],
    ["/sideline.sideline.MsgCreateTask", MsgCreateTask],
    ["/sideline.sideline.MsgRegistEmployer", MsgRegistEmployer],
    
];

export { msgTypes }