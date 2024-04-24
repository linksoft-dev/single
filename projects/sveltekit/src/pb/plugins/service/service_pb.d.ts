import * as jspb from 'google-protobuf'

import * as google_protobuf_descriptor_pb from 'google-protobuf/google/protobuf/descriptor_pb'; // proto import: "google/protobuf/descriptor.proto"


export class Field extends jspb.Message {
  getUppernospacenoaccent(): Boolean;
  setUppernospacenoaccent(value: Boolean): Field;

  getUppercase(): Boolean;
  setUppercase(value: Boolean): Field;

  getTrimspace(): Boolean;
  setTrimspace(value: Boolean): Field;

  getRemoveaccent(): Boolean;
  setRemoveaccent(value: Boolean): Field;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Field.AsObject;
  static toObject(includeInstance: boolean, msg: Field): Field.AsObject;
  static serializeBinaryToWriter(message: Field, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Field;
  static deserializeBinaryFromReader(message: Field, reader: jspb.BinaryReader): Field;
}

export namespace Field {
  export type AsObject = {
    uppernospacenoaccent: Boolean,
    uppercase: Boolean,
    trimspace: Boolean,
    removeaccent: Boolean,
  }
}

export enum Boolean { 
  UNSPECIFIED = 0,
  YES = 1,
  NO = 2,
}
