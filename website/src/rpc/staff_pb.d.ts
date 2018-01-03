// package: staff
// file: staff.proto

import * as jspb from "google-protobuf";

export class GetAuthStatusRequest extends jspb.Message {
  getToken(): string;
  setToken(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAuthStatusRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetAuthStatusRequest): GetAuthStatusRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetAuthStatusRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAuthStatusRequest;
  static deserializeBinaryFromReader(message: GetAuthStatusRequest, reader: jspb.BinaryReader): GetAuthStatusRequest;
}

export namespace GetAuthStatusRequest {
  export type AsObject = {
    token: string,
  }
}

export class GetAuthStatusResult extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAuthStatusResult.AsObject;
  static toObject(includeInstance: boolean, msg: GetAuthStatusResult): GetAuthStatusResult.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetAuthStatusResult, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAuthStatusResult;
  static deserializeBinaryFromReader(message: GetAuthStatusResult, reader: jspb.BinaryReader): GetAuthStatusResult;
}

export namespace GetAuthStatusResult {
  export type AsObject = {
  }
}

export class GetUserRequest extends jspb.Message {
  getSearch(): string;
  setSearch(value: string): void;

  getType(): GetUserRequest.searchType;
  setType(value: GetUserRequest.searchType): void;

  getToken(): string;
  setToken(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetUserRequest): GetUserRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUserRequest;
  static deserializeBinaryFromReader(message: GetUserRequest, reader: jspb.BinaryReader): GetUserRequest;
}

export namespace GetUserRequest {
  export type AsObject = {
    search: string,
    type: GetUserRequest.searchType,
    token: string,
  }

  export enum searchType {
    ID = 0,
    TSDBID = 1,
    TSUUID = 2,
    EMAIL = 3,
    STEAMID = 4,
    TOKEN = 5,
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

  getSteamid(): string;
  setSteamid(value: string): void;

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
    steamid: string,
  }
}

