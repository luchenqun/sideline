/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "sideline.sideline";

export interface Task {
  id: number;
  title: string;
  description: string;
  remuneration: string;
  depositEmployer: string;
  depositDeveloper: string;
  employer: string;
  developer: string;
  deadline: number;
  status: number;
}

function createBaseTask(): Task {
  return {
    id: 0,
    title: "",
    description: "",
    remuneration: "",
    depositEmployer: "",
    depositDeveloper: "",
    employer: "",
    developer: "",
    deadline: 0,
    status: 0,
  };
}

export const Task = {
  encode(message: Task, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.title !== "") {
      writer.uint32(18).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.remuneration !== "") {
      writer.uint32(34).string(message.remuneration);
    }
    if (message.depositEmployer !== "") {
      writer.uint32(42).string(message.depositEmployer);
    }
    if (message.depositDeveloper !== "") {
      writer.uint32(50).string(message.depositDeveloper);
    }
    if (message.employer !== "") {
      writer.uint32(58).string(message.employer);
    }
    if (message.developer !== "") {
      writer.uint32(66).string(message.developer);
    }
    if (message.deadline !== 0) {
      writer.uint32(72).uint64(message.deadline);
    }
    if (message.status !== 0) {
      writer.uint32(80).uint64(message.status);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Task {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTask();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.title = reader.string();
          break;
        case 3:
          message.description = reader.string();
          break;
        case 4:
          message.remuneration = reader.string();
          break;
        case 5:
          message.depositEmployer = reader.string();
          break;
        case 6:
          message.depositDeveloper = reader.string();
          break;
        case 7:
          message.employer = reader.string();
          break;
        case 8:
          message.developer = reader.string();
          break;
        case 9:
          message.deadline = longToNumber(reader.uint64() as Long);
          break;
        case 10:
          message.status = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Task {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      title: isSet(object.title) ? String(object.title) : "",
      description: isSet(object.description) ? String(object.description) : "",
      remuneration: isSet(object.remuneration) ? String(object.remuneration) : "",
      depositEmployer: isSet(object.depositEmployer) ? String(object.depositEmployer) : "",
      depositDeveloper: isSet(object.depositDeveloper) ? String(object.depositDeveloper) : "",
      employer: isSet(object.employer) ? String(object.employer) : "",
      developer: isSet(object.developer) ? String(object.developer) : "",
      deadline: isSet(object.deadline) ? Number(object.deadline) : 0,
      status: isSet(object.status) ? Number(object.status) : 0,
    };
  },

  toJSON(message: Task): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.title !== undefined && (obj.title = message.title);
    message.description !== undefined && (obj.description = message.description);
    message.remuneration !== undefined && (obj.remuneration = message.remuneration);
    message.depositEmployer !== undefined && (obj.depositEmployer = message.depositEmployer);
    message.depositDeveloper !== undefined && (obj.depositDeveloper = message.depositDeveloper);
    message.employer !== undefined && (obj.employer = message.employer);
    message.developer !== undefined && (obj.developer = message.developer);
    message.deadline !== undefined && (obj.deadline = Math.round(message.deadline));
    message.status !== undefined && (obj.status = Math.round(message.status));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Task>, I>>(object: I): Task {
    const message = createBaseTask();
    message.id = object.id ?? 0;
    message.title = object.title ?? "";
    message.description = object.description ?? "";
    message.remuneration = object.remuneration ?? "";
    message.depositEmployer = object.depositEmployer ?? "";
    message.depositDeveloper = object.depositDeveloper ?? "";
    message.employer = object.employer ?? "";
    message.developer = object.developer ?? "";
    message.deadline = object.deadline ?? 0;
    message.status = object.status ?? 0;
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
