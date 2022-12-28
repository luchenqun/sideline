/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "sideline.sideline";

export interface Developer {
  index: string;
  name: string;
  introduce: string;
  email: string;
  avatar: string;
  address: string;
  education: string;
  major: string;
  skills: string[];
  taskSuccess: number;
  taskFail: number;
  feedbacks: string[];
}

function createBaseDeveloper(): Developer {
  return {
    index: "",
    name: "",
    introduce: "",
    email: "",
    avatar: "",
    address: "",
    education: "",
    major: "",
    skills: [],
    taskSuccess: 0,
    taskFail: 0,
    feedbacks: [],
  };
}

export const Developer = {
  encode(message: Developer, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.introduce !== "") {
      writer.uint32(26).string(message.introduce);
    }
    if (message.email !== "") {
      writer.uint32(34).string(message.email);
    }
    if (message.avatar !== "") {
      writer.uint32(42).string(message.avatar);
    }
    if (message.address !== "") {
      writer.uint32(50).string(message.address);
    }
    if (message.education !== "") {
      writer.uint32(58).string(message.education);
    }
    if (message.major !== "") {
      writer.uint32(66).string(message.major);
    }
    for (const v of message.skills) {
      writer.uint32(74).string(v!);
    }
    if (message.taskSuccess !== 0) {
      writer.uint32(80).uint64(message.taskSuccess);
    }
    if (message.taskFail !== 0) {
      writer.uint32(88).uint64(message.taskFail);
    }
    for (const v of message.feedbacks) {
      writer.uint32(98).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Developer {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeveloper();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.introduce = reader.string();
          break;
        case 4:
          message.email = reader.string();
          break;
        case 5:
          message.avatar = reader.string();
          break;
        case 6:
          message.address = reader.string();
          break;
        case 7:
          message.education = reader.string();
          break;
        case 8:
          message.major = reader.string();
          break;
        case 9:
          message.skills.push(reader.string());
          break;
        case 10:
          message.taskSuccess = longToNumber(reader.uint64() as Long);
          break;
        case 11:
          message.taskFail = longToNumber(reader.uint64() as Long);
          break;
        case 12:
          message.feedbacks.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Developer {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      name: isSet(object.name) ? String(object.name) : "",
      introduce: isSet(object.introduce) ? String(object.introduce) : "",
      email: isSet(object.email) ? String(object.email) : "",
      avatar: isSet(object.avatar) ? String(object.avatar) : "",
      address: isSet(object.address) ? String(object.address) : "",
      education: isSet(object.education) ? String(object.education) : "",
      major: isSet(object.major) ? String(object.major) : "",
      skills: Array.isArray(object?.skills) ? object.skills.map((e: any) => String(e)) : [],
      taskSuccess: isSet(object.taskSuccess) ? Number(object.taskSuccess) : 0,
      taskFail: isSet(object.taskFail) ? Number(object.taskFail) : 0,
      feedbacks: Array.isArray(object?.feedbacks) ? object.feedbacks.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: Developer): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.name !== undefined && (obj.name = message.name);
    message.introduce !== undefined && (obj.introduce = message.introduce);
    message.email !== undefined && (obj.email = message.email);
    message.avatar !== undefined && (obj.avatar = message.avatar);
    message.address !== undefined && (obj.address = message.address);
    message.education !== undefined && (obj.education = message.education);
    message.major !== undefined && (obj.major = message.major);
    if (message.skills) {
      obj.skills = message.skills.map((e) => e);
    } else {
      obj.skills = [];
    }
    message.taskSuccess !== undefined && (obj.taskSuccess = Math.round(message.taskSuccess));
    message.taskFail !== undefined && (obj.taskFail = Math.round(message.taskFail));
    if (message.feedbacks) {
      obj.feedbacks = message.feedbacks.map((e) => e);
    } else {
      obj.feedbacks = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Developer>, I>>(object: I): Developer {
    const message = createBaseDeveloper();
    message.index = object.index ?? "";
    message.name = object.name ?? "";
    message.introduce = object.introduce ?? "";
    message.email = object.email ?? "";
    message.avatar = object.avatar ?? "";
    message.address = object.address ?? "";
    message.education = object.education ?? "";
    message.major = object.major ?? "";
    message.skills = object.skills?.map((e) => e) || [];
    message.taskSuccess = object.taskSuccess ?? 0;
    message.taskFail = object.taskFail ?? 0;
    message.feedbacks = object.feedbacks?.map((e) => e) || [];
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
