// package: 
// file: staff.proto

import * as staff_pb from "./staff_pb";
export class members {
  static serviceName = "members";
}
export namespace members {
  export class GetUser {
    static readonly methodName = "GetUser";
    static readonly service = members;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = staff_pb.GetUserMessage;
    static readonly responseType = staff_pb.User;
  }
}
