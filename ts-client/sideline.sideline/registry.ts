import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRegistEmployer } from "./types/sideline/sideline/tx";
import { MsgRegistDeveloper } from "./types/sideline/sideline/tx";
import { MsgCreateTask } from "./types/sideline/sideline/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/sideline.sideline.MsgRegistEmployer", MsgRegistEmployer],
    ["/sideline.sideline.MsgRegistDeveloper", MsgRegistDeveloper],
    ["/sideline.sideline.MsgCreateTask", MsgCreateTask],
    
];

export { msgTypes }