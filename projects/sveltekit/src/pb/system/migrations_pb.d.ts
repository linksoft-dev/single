import * as jspb from 'google-protobuf'

import * as plugins_validate_validate_pb from '../plugins/validate/validate_pb'; // proto import: "plugins/validate/validate.proto"
import * as plugins_service_service_pb from '../plugins/service/service_pb'; // proto import: "plugins/service/service.proto"
import * as filter_filter_pb from '../filter/filter_pb'; // proto import: "filter/filter.proto"
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb'; // proto import: "google/protobuf/timestamp.proto"
import * as google_protobuf_field_mask_pb from 'google-protobuf/google/protobuf/field_mask_pb'; // proto import: "google/protobuf/field_mask.proto"


export class MigrationRequest extends jspb.Message {
  getId(): string;
  setId(value: string): MigrationRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MigrationRequest.AsObject;
  static toObject(includeInstance: boolean, msg: MigrationRequest): MigrationRequest.AsObject;
  static serializeBinaryToWriter(message: MigrationRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MigrationRequest;
  static deserializeBinaryFromReader(message: MigrationRequest, reader: jspb.BinaryReader): MigrationRequest;
}

export namespace MigrationRequest {
  export type AsObject = {
    id: string,
  }
}

export class MigrationResponse extends jspb.Message {
  getMigration(): Migration | undefined;
  setMigration(value?: Migration): MigrationResponse;
  hasMigration(): boolean;
  clearMigration(): MigrationResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MigrationResponse.AsObject;
  static toObject(includeInstance: boolean, msg: MigrationResponse): MigrationResponse.AsObject;
  static serializeBinaryToWriter(message: MigrationResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MigrationResponse;
  static deserializeBinaryFromReader(message: MigrationResponse, reader: jspb.BinaryReader): MigrationResponse;
}

export namespace MigrationResponse {
  export type AsObject = {
    migration?: Migration.AsObject,
  }
}

export class Migration extends jspb.Message {
  getId(): string;
  setId(value: string): Migration;

  getDescription(): string;
  setDescription(value: string): Migration;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Migration;
  hasCreatedAt(): boolean;
  clearCreatedAt(): Migration;

  getAppliedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setAppliedAt(value?: google_protobuf_timestamp_pb.Timestamp): Migration;
  hasAppliedAt(): boolean;
  clearAppliedAt(): Migration;

  getResult(): string;
  setResult(value: string): Migration;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Migration.AsObject;
  static toObject(includeInstance: boolean, msg: Migration): Migration.AsObject;
  static serializeBinaryToWriter(message: Migration, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Migration;
  static deserializeBinaryFromReader(message: Migration, reader: jspb.BinaryReader): Migration;
}

export namespace Migration {
  export type AsObject = {
    id: string,
    description: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    appliedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    result: string,
  }
}

export class CreateMigrationRequest extends jspb.Message {
  getMigration(): Migration | undefined;
  setMigration(value?: Migration): CreateMigrationRequest;
  hasMigration(): boolean;
  clearMigration(): CreateMigrationRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateMigrationRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateMigrationRequest): CreateMigrationRequest.AsObject;
  static serializeBinaryToWriter(message: CreateMigrationRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateMigrationRequest;
  static deserializeBinaryFromReader(message: CreateMigrationRequest, reader: jspb.BinaryReader): CreateMigrationRequest;
}

export namespace CreateMigrationRequest {
  export type AsObject = {
    migration?: Migration.AsObject,
  }
}

export class CreateMigrationResponse extends jspb.Message {
  getMigration(): Migration | undefined;
  setMigration(value?: Migration): CreateMigrationResponse;
  hasMigration(): boolean;
  clearMigration(): CreateMigrationResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateMigrationResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateMigrationResponse): CreateMigrationResponse.AsObject;
  static serializeBinaryToWriter(message: CreateMigrationResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateMigrationResponse;
  static deserializeBinaryFromReader(message: CreateMigrationResponse, reader: jspb.BinaryReader): CreateMigrationResponse;
}

export namespace CreateMigrationResponse {
  export type AsObject = {
    migration?: Migration.AsObject,
  }
}

export class UpdateMigrationRequest extends jspb.Message {
  getId(): string;
  setId(value: string): UpdateMigrationRequest;

  getMigration(): Migration | undefined;
  setMigration(value?: Migration): UpdateMigrationRequest;
  hasMigration(): boolean;
  clearMigration(): UpdateMigrationRequest;

  getUpdateMask(): google_protobuf_field_mask_pb.FieldMask | undefined;
  setUpdateMask(value?: google_protobuf_field_mask_pb.FieldMask): UpdateMigrationRequest;
  hasUpdateMask(): boolean;
  clearUpdateMask(): UpdateMigrationRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateMigrationRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateMigrationRequest): UpdateMigrationRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateMigrationRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateMigrationRequest;
  static deserializeBinaryFromReader(message: UpdateMigrationRequest, reader: jspb.BinaryReader): UpdateMigrationRequest;
}

export namespace UpdateMigrationRequest {
  export type AsObject = {
    id: string,
    migration?: Migration.AsObject,
    updateMask?: google_protobuf_field_mask_pb.FieldMask.AsObject,
  }
}

export class UpdateMigrationResponse extends jspb.Message {
  getMigration(): Migration | undefined;
  setMigration(value?: Migration): UpdateMigrationResponse;
  hasMigration(): boolean;
  clearMigration(): UpdateMigrationResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateMigrationResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateMigrationResponse): UpdateMigrationResponse.AsObject;
  static serializeBinaryToWriter(message: UpdateMigrationResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateMigrationResponse;
  static deserializeBinaryFromReader(message: UpdateMigrationResponse, reader: jspb.BinaryReader): UpdateMigrationResponse;
}

export namespace UpdateMigrationResponse {
  export type AsObject = {
    migration?: Migration.AsObject,
  }
}

export class DeleteMigrationRequest extends jspb.Message {
  getId(): string;
  setId(value: string): DeleteMigrationRequest;

  getHard(): boolean;
  setHard(value: boolean): DeleteMigrationRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteMigrationRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteMigrationRequest): DeleteMigrationRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteMigrationRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteMigrationRequest;
  static deserializeBinaryFromReader(message: DeleteMigrationRequest, reader: jspb.BinaryReader): DeleteMigrationRequest;
}

export namespace DeleteMigrationRequest {
  export type AsObject = {
    id: string,
    hard: boolean,
  }
}

export class DeleteMigrationResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteMigrationResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteMigrationResponse): DeleteMigrationResponse.AsObject;
  static serializeBinaryToWriter(message: DeleteMigrationResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteMigrationResponse;
  static deserializeBinaryFromReader(message: DeleteMigrationResponse, reader: jspb.BinaryReader): DeleteMigrationResponse;
}

export namespace DeleteMigrationResponse {
  export type AsObject = {
  }
}

export class ListMigrationRequest extends jspb.Message {
  getId(): string;
  setId(value: string): ListMigrationRequest;

  getNome(): string;
  setNome(value: string): ListMigrationRequest;

  getFilter(): filter_filter_pb.Filter | undefined;
  setFilter(value?: filter_filter_pb.Filter): ListMigrationRequest;
  hasFilter(): boolean;
  clearFilter(): ListMigrationRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListMigrationRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListMigrationRequest): ListMigrationRequest.AsObject;
  static serializeBinaryToWriter(message: ListMigrationRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListMigrationRequest;
  static deserializeBinaryFromReader(message: ListMigrationRequest, reader: jspb.BinaryReader): ListMigrationRequest;
}

export namespace ListMigrationRequest {
  export type AsObject = {
    id: string,
    nome: string,
    filter?: filter_filter_pb.Filter.AsObject,
  }
}

export class ListMigrationResponse extends jspb.Message {
  getMigrationsList(): Array<Migration>;
  setMigrationsList(value: Array<Migration>): ListMigrationResponse;
  clearMigrationsList(): ListMigrationResponse;
  addMigrations(value?: Migration, index?: number): Migration;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListMigrationResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListMigrationResponse): ListMigrationResponse.AsObject;
  static serializeBinaryToWriter(message: ListMigrationResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListMigrationResponse;
  static deserializeBinaryFromReader(message: ListMigrationResponse, reader: jspb.BinaryReader): ListMigrationResponse;
}

export namespace ListMigrationResponse {
  export type AsObject = {
    migrationsList: Array<Migration.AsObject>,
  }
}

export class GetMigrationRequest extends jspb.Message {
  getId(): string;
  setId(value: string): GetMigrationRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetMigrationRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetMigrationRequest): GetMigrationRequest.AsObject;
  static serializeBinaryToWriter(message: GetMigrationRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetMigrationRequest;
  static deserializeBinaryFromReader(message: GetMigrationRequest, reader: jspb.BinaryReader): GetMigrationRequest;
}

export namespace GetMigrationRequest {
  export type AsObject = {
    id: string,
  }
}

export class GetMigrationResponse extends jspb.Message {
  getMigration(): Migration | undefined;
  setMigration(value?: Migration): GetMigrationResponse;
  hasMigration(): boolean;
  clearMigration(): GetMigrationResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetMigrationResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetMigrationResponse): GetMigrationResponse.AsObject;
  static serializeBinaryToWriter(message: GetMigrationResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetMigrationResponse;
  static deserializeBinaryFromReader(message: GetMigrationResponse, reader: jspb.BinaryReader): GetMigrationResponse;
}

export namespace GetMigrationResponse {
  export type AsObject = {
    migration?: Migration.AsObject,
  }
}

