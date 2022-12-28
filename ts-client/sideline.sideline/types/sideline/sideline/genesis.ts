/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Developer } from "./developer";
import { Employer } from "./employer";
import { Params } from "./params";
import { Task } from "./task";

export const protobufPackage = "sideline.sideline";

/** GenesisState defines the sideline module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  employerList: Employer[];
  developerList: Developer[];
  taskList: Task[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  taskCount: number;
}

function createBaseGenesisState(): GenesisState {
  return { params: undefined, employerList: [], developerList: [], taskList: [], taskCount: 0 };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.employerList) {
      Employer.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.developerList) {
      Developer.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.taskList) {
      Task.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.taskCount !== 0) {
      writer.uint32(40).uint64(message.taskCount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.employerList.push(Employer.decode(reader, reader.uint32()));
          break;
        case 3:
          message.developerList.push(Developer.decode(reader, reader.uint32()));
          break;
        case 4:
          message.taskList.push(Task.decode(reader, reader.uint32()));
          break;
        case 5:
          message.taskCount = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      employerList: Array.isArray(object?.employerList)
        ? object.employerList.map((e: any) => Employer.fromJSON(e))
        : [],
      developerList: Array.isArray(object?.developerList)
        ? object.developerList.map((e: any) => Developer.fromJSON(e))
        : [],
      taskList: Array.isArray(object?.taskList) ? object.taskList.map((e: any) => Task.fromJSON(e)) : [],
      taskCount: isSet(object.taskCount) ? Number(object.taskCount) : 0,
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.employerList) {
      obj.employerList = message.employerList.map((e) => e ? Employer.toJSON(e) : undefined);
    } else {
      obj.employerList = [];
    }
    if (message.developerList) {
      obj.developerList = message.developerList.map((e) => e ? Developer.toJSON(e) : undefined);
    } else {
      obj.developerList = [];
    }
    if (message.taskList) {
      obj.taskList = message.taskList.map((e) => e ? Task.toJSON(e) : undefined);
    } else {
      obj.taskList = [];
    }
    message.taskCount !== undefined && (obj.taskCount = Math.round(message.taskCount));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.employerList = object.employerList?.map((e) => Employer.fromPartial(e)) || [];
    message.developerList = object.developerList?.map((e) => Developer.fromPartial(e)) || [];
    message.taskList = object.taskList?.map((e) => Task.fromPartial(e)) || [];
    message.taskCount = object.taskCount ?? 0;
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
