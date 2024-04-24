import * as grpcWeb from 'grpc-web';

import * as test_test_pb from '../test/test_pb'; // proto import: "test/test.proto"


export class TestServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  test(
    request: test_test_pb.TestRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: test_test_pb.TestResponse) => void
  ): grpcWeb.ClientReadableStream<test_test_pb.TestResponse>;

}

export class TestServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  test(
    request: test_test_pb.TestRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<test_test_pb.TestResponse>;

}

