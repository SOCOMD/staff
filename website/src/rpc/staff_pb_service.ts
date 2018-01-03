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
    static readonly requestType = staff_pb.GetUserMessage;
    static readonly responseType = staff_pb.User;
  }
}
