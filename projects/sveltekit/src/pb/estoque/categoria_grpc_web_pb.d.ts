import * as grpcWeb from 'grpc-web';

import * as estoque_categoria_pb from '../estoque/categoria_pb'; // proto import: "estoque/categoria.proto"


export class CategoriaServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  create(
    request: estoque_categoria_pb.CreateCategoriaRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: estoque_categoria_pb.CreateCategoriaResponse) => void
  ): grpcWeb.ClientReadableStream<estoque_categoria_pb.CreateCategoriaResponse>;

  update(
    request: estoque_categoria_pb.UpdateCategoriaRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: estoque_categoria_pb.UpdateCategoriaResponse) => void
  ): grpcWeb.ClientReadableStream<estoque_categoria_pb.UpdateCategoriaResponse>;

  delete(
    request: estoque_categoria_pb.DeleteCategoriaRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: estoque_categoria_pb.DeleteCategoriaResponse) => void
  ): grpcWeb.ClientReadableStream<estoque_categoria_pb.DeleteCategoriaResponse>;

  list(
    request: estoque_categoria_pb.ListCategoriaRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: estoque_categoria_pb.ListCategoriaResponse) => void
  ): grpcWeb.ClientReadableStream<estoque_categoria_pb.ListCategoriaResponse>;

  get(
    request: estoque_categoria_pb.GetCategoriaRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: estoque_categoria_pb.GetCategoriaResponse) => void
  ): grpcWeb.ClientReadableStream<estoque_categoria_pb.GetCategoriaResponse>;

}

export class CategoriaServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  create(
    request: estoque_categoria_pb.CreateCategoriaRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<estoque_categoria_pb.CreateCategoriaResponse>;

  update(
    request: estoque_categoria_pb.UpdateCategoriaRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<estoque_categoria_pb.UpdateCategoriaResponse>;

  delete(
    request: estoque_categoria_pb.DeleteCategoriaRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<estoque_categoria_pb.DeleteCategoriaResponse>;

  list(
    request: estoque_categoria_pb.ListCategoriaRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<estoque_categoria_pb.ListCategoriaResponse>;

  get(
    request: estoque_categoria_pb.GetCategoriaRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<estoque_categoria_pb.GetCategoriaResponse>;

}

