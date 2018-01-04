// package: staff
// file: staff.proto

import * as staff_pb from "./staff_pb";
export class staff {
  static serviceName = "staff.staff";
}
export namespace staff {
  export class GetUser {
    static readonly methodName = "GetUser";
    static readonly service = staff;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = staff_pb.GetUserRequest;
    static readonly responseType = staff_pb.User;
  }
  export class UpdateUser {
    static readonly methodName = "UpdateUser";
    static readonly service = staff;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = staff_pb.UpdateUserRequest;
    static readonly responseType = staff_pb.NilResult;
  }
  export class AuthStatus {
    static readonly methodName = "AuthStatus";
    static readonly service = staff;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = staff_pb.GetAuthStatusRequest;
    static readonly responseType = staff_pb.GetAuthStatusResult;
  }
}
