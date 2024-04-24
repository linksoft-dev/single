import * as jspb from 'google-protobuf'

import * as google_api_annotations_pb from '../google/api/annotations_pb'; // proto import: "google/api/annotations.proto"
import * as google_api_resource_pb from '../google/api/resource_pb'; // proto import: "google/api/resource.proto"
import * as plugins_validate_validate_pb from '../plugins/validate/validate_pb'; // proto import: "plugins/validate/validate.proto"
import * as plugins_service_service_pb from '../plugins/service/service_pb'; // proto import: "plugins/service/service.proto"
import * as filter_filter_pb from '../filter/filter_pb'; // proto import: "filter/filter.proto"
import * as google_protobuf_field_mask_pb from 'google-protobuf/google/protobuf/field_mask_pb'; // proto import: "google/protobuf/field_mask.proto"


export class CreateCategoriaRequest extends jspb.Message {
  getCategoria(): Categoria | undefined;
  setCategoria(value?: Categoria): CreateCategoriaRequest;
  hasCategoria(): boolean;
  clearCategoria(): CreateCategoriaRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateCategoriaRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateCategoriaRequest): CreateCategoriaRequest.AsObject;
  static serializeBinaryToWriter(message: CreateCategoriaRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateCategoriaRequest;
  static deserializeBinaryFromReader(message: CreateCategoriaRequest, reader: jspb.BinaryReader): CreateCategoriaRequest;
}

export namespace CreateCategoriaRequest {
  export type AsObject = {
    categoria?: Categoria.AsObject,
  }
}

export class CreateCategoriaResponse extends jspb.Message {
  getCategoria(): Categoria | undefined;
  setCategoria(value?: Categoria): CreateCategoriaResponse;
  hasCategoria(): boolean;
  clearCategoria(): CreateCategoriaResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateCategoriaResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateCategoriaResponse): CreateCategoriaResponse.AsObject;
  static serializeBinaryToWriter(message: CreateCategoriaResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateCategoriaResponse;
  static deserializeBinaryFromReader(message: CreateCategoriaResponse, reader: jspb.BinaryReader): CreateCategoriaResponse;
}

export namespace CreateCategoriaResponse {
  export type AsObject = {
    categoria?: Categoria.AsObject,
  }
}

export class UpdateCategoriaRequest extends jspb.Message {
  getId(): string;
  setId(value: string): UpdateCategoriaRequest;

  getCategoria(): Categoria | undefined;
  setCategoria(value?: Categoria): UpdateCategoriaRequest;
  hasCategoria(): boolean;
  clearCategoria(): UpdateCategoriaRequest;

  getUpdateMask(): google_protobuf_field_mask_pb.FieldMask | undefined;
  setUpdateMask(value?: google_protobuf_field_mask_pb.FieldMask): UpdateCategoriaRequest;
  hasUpdateMask(): boolean;
  clearUpdateMask(): UpdateCategoriaRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateCategoriaRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateCategoriaRequest): UpdateCategoriaRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateCategoriaRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateCategoriaRequest;
  static deserializeBinaryFromReader(message: UpdateCategoriaRequest, reader: jspb.BinaryReader): UpdateCategoriaRequest;
}

export namespace UpdateCategoriaRequest {
  export type AsObject = {
    id: string,
    categoria?: Categoria.AsObject,
    updateMask?: google_protobuf_field_mask_pb.FieldMask.AsObject,
  }
}

export class UpdateCategoriaResponse extends jspb.Message {
  getCategoria(): Categoria | undefined;
  setCategoria(value?: Categoria): UpdateCategoriaResponse;
  hasCategoria(): boolean;
  clearCategoria(): UpdateCategoriaResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateCategoriaResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateCategoriaResponse): UpdateCategoriaResponse.AsObject;
  static serializeBinaryToWriter(message: UpdateCategoriaResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateCategoriaResponse;
  static deserializeBinaryFromReader(message: UpdateCategoriaResponse, reader: jspb.BinaryReader): UpdateCategoriaResponse;
}

export namespace UpdateCategoriaResponse {
  export type AsObject = {
    categoria?: Categoria.AsObject,
  }
}

export class DeleteCategoriaRequest extends jspb.Message {
  getId(): string;
  setId(value: string): DeleteCategoriaRequest;

  getHard(): boolean;
  setHard(value: boolean): DeleteCategoriaRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteCategoriaRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteCategoriaRequest): DeleteCategoriaRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteCategoriaRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteCategoriaRequest;
  static deserializeBinaryFromReader(message: DeleteCategoriaRequest, reader: jspb.BinaryReader): DeleteCategoriaRequest;
}

export namespace DeleteCategoriaRequest {
  export type AsObject = {
    id: string,
    hard: boolean,
  }
}

export class DeleteCategoriaResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteCategoriaResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteCategoriaResponse): DeleteCategoriaResponse.AsObject;
  static serializeBinaryToWriter(message: DeleteCategoriaResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteCategoriaResponse;
  static deserializeBinaryFromReader(message: DeleteCategoriaResponse, reader: jspb.BinaryReader): DeleteCategoriaResponse;
}

export namespace DeleteCategoriaResponse {
  export type AsObject = {
  }
}

export class ListCategoriaRequest extends jspb.Message {
  getId(): string;
  setId(value: string): ListCategoriaRequest;

  getNome(): string;
  setNome(value: string): ListCategoriaRequest;

  getFilter(): filter_filter_pb.Filter | undefined;
  setFilter(value?: filter_filter_pb.Filter): ListCategoriaRequest;
  hasFilter(): boolean;
  clearFilter(): ListCategoriaRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListCategoriaRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListCategoriaRequest): ListCategoriaRequest.AsObject;
  static serializeBinaryToWriter(message: ListCategoriaRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListCategoriaRequest;
  static deserializeBinaryFromReader(message: ListCategoriaRequest, reader: jspb.BinaryReader): ListCategoriaRequest;
}

export namespace ListCategoriaRequest {
  export type AsObject = {
    id: string,
    nome: string,
    filter?: filter_filter_pb.Filter.AsObject,
  }
}

export class ListCategoriaResponse extends jspb.Message {
  getCategoriasList(): Array<Categoria>;
  setCategoriasList(value: Array<Categoria>): ListCategoriaResponse;
  clearCategoriasList(): ListCategoriaResponse;
  addCategorias(value?: Categoria, index?: number): Categoria;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListCategoriaResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListCategoriaResponse): ListCategoriaResponse.AsObject;
  static serializeBinaryToWriter(message: ListCategoriaResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListCategoriaResponse;
  static deserializeBinaryFromReader(message: ListCategoriaResponse, reader: jspb.BinaryReader): ListCategoriaResponse;
}

export namespace ListCategoriaResponse {
  export type AsObject = {
    categoriasList: Array<Categoria.AsObject>,
  }
}

export class GetCategoriaRequest extends jspb.Message {
  getId(): string;
  setId(value: string): GetCategoriaRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetCategoriaRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetCategoriaRequest): GetCategoriaRequest.AsObject;
  static serializeBinaryToWriter(message: GetCategoriaRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetCategoriaRequest;
  static deserializeBinaryFromReader(message: GetCategoriaRequest, reader: jspb.BinaryReader): GetCategoriaRequest;
}

export namespace GetCategoriaRequest {
  export type AsObject = {
    id: string,
  }
}

export class GetCategoriaResponse extends jspb.Message {
  getCategoria(): Categoria | undefined;
  setCategoria(value?: Categoria): GetCategoriaResponse;
  hasCategoria(): boolean;
  clearCategoria(): GetCategoriaResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetCategoriaResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetCategoriaResponse): GetCategoriaResponse.AsObject;
  static serializeBinaryToWriter(message: GetCategoriaResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetCategoriaResponse;
  static deserializeBinaryFromReader(message: GetCategoriaResponse, reader: jspb.BinaryReader): GetCategoriaResponse;
}

export namespace GetCategoriaResponse {
  export type AsObject = {
    categoria?: Categoria.AsObject,
  }
}

export class Categoria extends jspb.Message {
  getId(): string;
  setId(value: string): Categoria;

  getNome(): string;
  setNome(value: string): Categoria;

  getTributacaoId(): string;
  setTributacaoId(value: string): Categoria;

  getTributacaoNome(): string;
  setTributacaoNome(value: string): Categoria;

  getNcm(): string;
  setNcm(value: string): Categoria;

  getQuantidadeMinima(): number;
  setQuantidadeMinima(value: number): Categoria;

  getQuantidadeMaxima(): number;
  setQuantidadeMaxima(value: number): Categoria;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Categoria.AsObject;
  static toObject(includeInstance: boolean, msg: Categoria): Categoria.AsObject;
  static serializeBinaryToWriter(message: Categoria, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Categoria;
  static deserializeBinaryFromReader(message: Categoria, reader: jspb.BinaryReader): Categoria;
}

export namespace Categoria {
  export type AsObject = {
    id: string,
    nome: string,
    tributacaoId: string,
    tributacaoNome: string,
    ncm: string,
    quantidadeMinima: number,
    quantidadeMaxima: number,
  }
}

