// package: 
// file: staff.proto

import * as jspb from "google-protobuf";

export class GetUserMessage extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUserMessage.AsObject;
  static toObject(includeInstance: boolean, msg: GetUserMessage): GetUserMessage.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetUserMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUserMessage;
  static deserializeBinaryFromReader(message: GetUserMessage, reader: jspb.BinaryReader): GetUserMessage;
}

export namespace GetUserMessage {
  export type AsObject = {
    id: string,
  }
}

export class User extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getTsname(): string;
  setTsname(value: string): void;

  getTsdbid(): string;
  setTsdbid(value: string): void;

  getTsuuid(): string;
  setTsuuid(value: string): void;

  getTscreated(): string;
  setTscreated(value: string): void;

  getTslastconnected(): string;
  setTslastconnected(value: string): void;

  getEmail(): string;
  setEmail(value: string): void;

  getJoindate(): string;
  setJoindate(value: string): void;

  getDob(): string;
  setDob(value: string): void;

  getGender(): string;
  setGender(value: string): void;

  getActive(): boolean;
  setActive(value: boolean): void;

  getAdmin(): number;
  setAdmin(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    id: string,
    tsname: string,
    tsdbid: string,
    tsuuid: string,
    tscreated: string,
    tslastconnected: string,
    email: string,
    joindate: string,
    dob: string,
    gender: string,
    active: boolean,
    admin: number,
  }
}

