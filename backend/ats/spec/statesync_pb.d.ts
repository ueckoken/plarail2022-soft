// package: 
// file: statesync.proto

import * as jspb from "google-protobuf";

export class Station extends jspb.Message {
  getStationid(): StationIdMap[keyof StationIdMap];
  setStationid(value: StationIdMap[keyof StationIdMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Station.AsObject;
  static toObject(includeInstance: boolean, msg: Station): Station.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Station, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Station;
  static deserializeBinaryFromReader(message: Station, reader: jspb.BinaryReader): Station;
}

export namespace Station {
  export type AsObject = {
    stationid: StationIdMap[keyof StationIdMap],
  }
}

export class PointAndState extends jspb.Message {
  hasStation(): boolean;
  clearStation(): void;
  getStation(): Station | undefined;
  setStation(value?: Station): void;

  getState(): PointStateEnumMap[keyof PointStateEnumMap];
  setState(value: PointStateEnumMap[keyof PointStateEnumMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PointAndState.AsObject;
  static toObject(includeInstance: boolean, msg: PointAndState): PointAndState.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PointAndState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PointAndState;
  static deserializeBinaryFromReader(message: PointAndState, reader: jspb.BinaryReader): PointAndState;
}

export namespace PointAndState {
  export type AsObject = {
    station?: Station.AsObject,
    state: PointStateEnumMap[keyof PointStateEnumMap],
  }
}

export class UpdatePointStateRequest extends jspb.Message {
  hasState(): boolean;
  clearState(): void;
  getState(): PointAndState | undefined;
  setState(value?: PointAndState): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdatePointStateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdatePointStateRequest): UpdatePointStateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdatePointStateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdatePointStateRequest;
  static deserializeBinaryFromReader(message: UpdatePointStateRequest, reader: jspb.BinaryReader): UpdatePointStateRequest;
}

export namespace UpdatePointStateRequest {
  export type AsObject = {
    state?: PointAndState.AsObject,
  }
}

export class UpdatePointStateResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdatePointStateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdatePointStateResponse): UpdatePointStateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdatePointStateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdatePointStateResponse;
  static deserializeBinaryFromReader(message: UpdatePointStateResponse, reader: jspb.BinaryReader): UpdatePointStateResponse;
}

export namespace UpdatePointStateResponse {
  export type AsObject = {
  }
}

export class NotifyPointStateRequest extends jspb.Message {
  hasState(): boolean;
  clearState(): void;
  getState(): PointAndState | undefined;
  setState(value?: PointAndState): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NotifyPointStateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: NotifyPointStateRequest): NotifyPointStateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NotifyPointStateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NotifyPointStateRequest;
  static deserializeBinaryFromReader(message: NotifyPointStateRequest, reader: jspb.BinaryReader): NotifyPointStateRequest;
}

export namespace NotifyPointStateRequest {
  export type AsObject = {
    state?: PointAndState.AsObject,
  }
}

export class NotifyPointStateResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NotifyPointStateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: NotifyPointStateResponse): NotifyPointStateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NotifyPointStateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NotifyPointStateResponse;
  static deserializeBinaryFromReader(message: NotifyPointStateResponse, reader: jspb.BinaryReader): NotifyPointStateResponse;
}

export namespace NotifyPointStateResponse {
  export type AsObject = {
  }
}

export interface StationIdMap {
  STATIONID_UNKNOWN: 0;
  SHINJUKU_S1: 2;
  SHINJUKU_S2: 3;
  SAKURAJOSUI_P1: 11;
  SAKURAJOSUI_P2: 12;
  SAKURAJOSUI_S0: 13;
  SAKURAJOSUI_S1: 14;
  SAKURAJOSUI_S2: 15;
  SAKURAJOSUI_S3: 16;
  SAKURAJOSUI_S4: 17;
  SAKURAJOSUI_S5: 18;
  CHOFU_P1: 21;
  CHOFU_S0: 22;
  CHOFU_S1: 23;
  CHOFU_S2: 24;
  CHOFU_S3: 25;
  CHOFU_S4: 26;
  HASHIMOTO_S1: 31;
  HASHIMOTO_S2: 32;
  HACHIOJI_S1: 41;
  HACHIOJI_S2: 42;
}

export const StationId: StationIdMap;

export interface PointStateEnumMap {
  POINTSTATE_UNKNOWN: 0;
  POINTSTATE_ON: 1;
  POINTSTATE_OFF: 2;
}

export const PointStateEnum: PointStateEnumMap;

