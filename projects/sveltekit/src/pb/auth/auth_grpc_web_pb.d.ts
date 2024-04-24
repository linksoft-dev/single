import * as grpcWeb from 'grpc-web';

import * as auth_auth_pb from '../auth/auth_pb'; // proto import: "auth/auth.proto"


export class AuthServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  login(
    request: auth_auth_pb.LoginRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: auth_auth_pb.LoginResponse) => void
  ): grpcWeb.ClientReadableStream<auth_auth_pb.LoginResponse>;

}

export class AuthServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  login(
    request: auth_auth_pb.LoginRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<auth_auth_pb.LoginResponse>;

}

