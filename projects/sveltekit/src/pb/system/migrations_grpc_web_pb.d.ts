import * as grpcWeb from 'grpc-web';

import * as system_migrations_pb from '../system/migrations_pb'; // proto import: "system/migrations.proto"


export class MigrationServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  run(
    request: system_migrations_pb.MigrationRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: system_migrations_pb.MigrationResponse) => void
  ): grpcWeb.ClientReadableStream<system_migrations_pb.MigrationResponse>;

}

export class MigrationServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  run(
    request: system_migrations_pb.MigrationRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<system_migrations_pb.MigrationResponse>;

}

