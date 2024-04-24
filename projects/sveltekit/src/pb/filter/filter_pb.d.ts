import * as jspb from 'google-protobuf'



export class Filter extends jspb.Message {
  getMainFilter(): string;
  setMainFilter(value: string): Filter;

  getSelectFieldsList(): Array<string>;
  setSelectFieldsList(value: Array<string>): Filter;
  clearSelectFieldsList(): Filter;
  addSelectFields(value: string, index?: number): Filter;

  getIdsList(): Array<string>;
  setIdsList(value: Array<string>): Filter;
  clearIdsList(): Filter;
  addIds(value: string, index?: number): Filter;

  getConditionsList(): Array<condition>;
  setConditionsList(value: Array<condition>): Filter;
  clearConditionsList(): Filter;
  addConditions(value?: condition, index?: number): condition;

  getOrconditionsList(): Array<condition>;
  setOrconditionsList(value: Array<condition>): Filter;
  clearOrconditionsList(): Filter;
  addOrconditions(value?: condition, index?: number): condition;

  getOrderbyList(): Array<OrderBy>;
  setOrderbyList(value: Array<OrderBy>): Filter;
  clearOrderbyList(): Filter;
  addOrderby(value?: OrderBy, index?: number): OrderBy;

  getLimit(): number;
  setLimit(value: number): Filter;

  getSkip(): number;
  setSkip(value: number): Filter;

  getFirst(): number;
  setFirst(value: number): Filter;

  getLast(): number;
  setLast(value: number): Filter;

  getRawfilter(): string;
  setRawfilter(value: string): Filter;

  getIgnoresoftdelete(): boolean;
  setIgnoresoftdelete(value: boolean): Filter;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Filter.AsObject;
  static toObject(includeInstance: boolean, msg: Filter): Filter.AsObject;
  static serializeBinaryToWriter(message: Filter, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Filter;
  static deserializeBinaryFromReader(message: Filter, reader: jspb.BinaryReader): Filter;
}

export namespace Filter {
  export type AsObject = {
    mainFilter: string,
    selectFieldsList: Array<string>,
    idsList: Array<string>,
    conditionsList: Array<condition.AsObject>,
    orconditionsList: Array<condition.AsObject>,
    orderbyList: Array<OrderBy.AsObject>,
    limit: number,
    skip: number,
    first: number,
    last: number,
    rawfilter: string,
    ignoresoftdelete: boolean,
  }
}

export class condition extends jspb.Message {
  getFieldName(): string;
  setFieldName(value: string): condition;

  getOperator(): Operator;
  setOperator(value: Operator): condition;

  getValue(): string;
  setValue(value: string): condition;

  getNot(): boolean;
  setNot(value: boolean): condition;

  getFilterOperator(): string;
  setFilterOperator(value: string): condition;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): condition.AsObject;
  static toObject(includeInstance: boolean, msg: condition): condition.AsObject;
  static serializeBinaryToWriter(message: condition, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): condition;
  static deserializeBinaryFromReader(message: condition, reader: jspb.BinaryReader): condition;
}

export namespace condition {
  export type AsObject = {
    fieldName: string,
    operator: Operator,
    value: string,
    not: boolean,
    filterOperator: string,
  }
}

export class OrderBy extends jspb.Message {
  getFieldName(): string;
  setFieldName(value: string): OrderBy;

  getDirection(): Direction;
  setDirection(value: Direction): OrderBy;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): OrderBy.AsObject;
  static toObject(includeInstance: boolean, msg: OrderBy): OrderBy.AsObject;
  static serializeBinaryToWriter(message: OrderBy, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): OrderBy;
  static deserializeBinaryFromReader(message: OrderBy, reader: jspb.BinaryReader): OrderBy;
}

export namespace OrderBy {
  export type AsObject = {
    fieldName: string,
    direction: Direction,
  }
}

export enum Direction { 
  ASC = 0,
  DESC = 1,
}
export enum Operator { 
  EQUALS = 0,
  CONTAINS = 1,
  STARTS = 2,
  IN = 3,
  GT = 4,
  GTE = 5,
  LT = 6,
  LTE = 7,
}
