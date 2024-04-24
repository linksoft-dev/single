import * as grpcWeb from 'grpc-web';

import * as sales_order_v1_order_pb from '../../../sales/order/v1/order_pb'; // proto import: "sales/order/v1/order.proto"


export class OrderServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  create(
    request: sales_order_v1_order_pb.CreateOrderRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: sales_order_v1_order_pb.CreateOrderResponse) => void
  ): grpcWeb.ClientReadableStream<sales_order_v1_order_pb.CreateOrderResponse>;

  update(
    request: sales_order_v1_order_pb.UpdateOrderRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: sales_order_v1_order_pb.UpdateOrderResponse) => void
  ): grpcWeb.ClientReadableStream<sales_order_v1_order_pb.UpdateOrderResponse>;

  delete(
    request: sales_order_v1_order_pb.DeleteOrderRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: sales_order_v1_order_pb.DeleteOrderResponse) => void
  ): grpcWeb.ClientReadableStream<sales_order_v1_order_pb.DeleteOrderResponse>;

  list(
    request: sales_order_v1_order_pb.ListOrderRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: sales_order_v1_order_pb.ListOrderResponse) => void
  ): grpcWeb.ClientReadableStream<sales_order_v1_order_pb.ListOrderResponse>;

  checkout(
    request: sales_order_v1_order_pb.CheckoutRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: sales_order_v1_order_pb.CheckoutResponse) => void
  ): grpcWeb.ClientReadableStream<sales_order_v1_order_pb.CheckoutResponse>;

}

export class OrderServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  create(
    request: sales_order_v1_order_pb.CreateOrderRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<sales_order_v1_order_pb.CreateOrderResponse>;

  update(
    request: sales_order_v1_order_pb.UpdateOrderRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<sales_order_v1_order_pb.UpdateOrderResponse>;

  delete(
    request: sales_order_v1_order_pb.DeleteOrderRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<sales_order_v1_order_pb.DeleteOrderResponse>;

  list(
    request: sales_order_v1_order_pb.ListOrderRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<sales_order_v1_order_pb.ListOrderResponse>;

  checkout(
    request: sales_order_v1_order_pb.CheckoutRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<sales_order_v1_order_pb.CheckoutResponse>;

}

