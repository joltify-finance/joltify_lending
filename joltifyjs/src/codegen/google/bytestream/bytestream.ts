import { Long, DeepPartial } from "../../helpers";
import * as _m0 from "protobufjs/minimal";
/** Request object for ByteStream.Read. */

export interface ReadRequest {
  /** The name of the resource to read. */
  resourceName: string;
  /**
   * The offset for the first byte to return in the read, relative to the start
   * of the resource.
   * 
   * A `read_offset` that is negative or greater than the size of the resource
   * will cause an `OUT_OF_RANGE` error.
   */

  readOffset: Long;
  /**
   * The maximum number of `data` bytes the server is allowed to return in the
   * sum of all `ReadResponse` messages. A `read_limit` of zero indicates that
   * there is no limit, and a negative `read_limit` will cause an error.
   * 
   * If the stream returns fewer bytes than allowed by the `read_limit` and no
   * error occurred, the stream includes all data from the `read_offset` to the
   * end of the resource.
   */

  readLimit: Long;
}
/** Request object for ByteStream.Read. */

export interface ReadRequestSDKType {
  resource_name: string;
  read_offset: Long;
  read_limit: Long;
}
/** Response object for ByteStream.Read. */

export interface ReadResponse {
  /**
   * A portion of the data for the resource. The service **may** leave `data`
   * empty for any given `ReadResponse`. This enables the service to inform the
   * client that the request is still live while it is running an operation to
   * generate more data.
   */
  data: Uint8Array;
}
/** Response object for ByteStream.Read. */

export interface ReadResponseSDKType {
  data: Uint8Array;
}
/** Request object for ByteStream.Write. */

export interface WriteRequest {
  /**
   * The name of the resource to write. This **must** be set on the first
   * `WriteRequest` of each `Write()` action. If it is set on subsequent calls,
   * it **must** match the value of the first request.
   */
  resourceName: string;
  /**
   * The offset from the beginning of the resource at which the data should be
   * written. It is required on all `WriteRequest`s.
   * 
   * In the first `WriteRequest` of a `Write()` action, it indicates
   * the initial offset for the `Write()` call. The value **must** be equal to
   * the `committed_size` that a call to `QueryWriteStatus()` would return.
   * 
   * On subsequent calls, this value **must** be set and **must** be equal to
   * the sum of the first `write_offset` and the sizes of all `data` bundles
   * sent previously on this stream.
   * 
   * An incorrect value will cause an error.
   */

  writeOffset: Long;
  /**
   * If `true`, this indicates that the write is complete. Sending any
   * `WriteRequest`s subsequent to one in which `finish_write` is `true` will
   * cause an error.
   */

  finishWrite: boolean;
  /**
   * A portion of the data for the resource. The client **may** leave `data`
   * empty for any given `WriteRequest`. This enables the client to inform the
   * service that the request is still live while it is running an operation to
   * generate more data.
   */

  data: Uint8Array;
}
/** Request object for ByteStream.Write. */

export interface WriteRequestSDKType {
  resource_name: string;
  write_offset: Long;
  finish_write: boolean;
  data: Uint8Array;
}
/** Response object for ByteStream.Write. */

export interface WriteResponse {
  /** The number of bytes that have been processed for the given resource. */
  committedSize: Long;
}
/** Response object for ByteStream.Write. */

export interface WriteResponseSDKType {
  committed_size: Long;
}
/** Request object for ByteStream.QueryWriteStatus. */

export interface QueryWriteStatusRequest {
  /** The name of the resource whose write status is being requested. */
  resourceName: string;
}
/** Request object for ByteStream.QueryWriteStatus. */

export interface QueryWriteStatusRequestSDKType {
  resource_name: string;
}
/** Response object for ByteStream.QueryWriteStatus. */

export interface QueryWriteStatusResponse {
  /** The number of bytes that have been processed for the given resource. */
  committedSize: Long;
  /**
   * `complete` is `true` only if the client has sent a `WriteRequest` with
   * `finish_write` set to true, and the server has processed that request.
   */

  complete: boolean;
}
/** Response object for ByteStream.QueryWriteStatus. */

export interface QueryWriteStatusResponseSDKType {
  committed_size: Long;
  complete: boolean;
}

function createBaseReadRequest(): ReadRequest {
  return {
    resourceName: "",
    readOffset: Long.ZERO,
    readLimit: Long.ZERO
  };
}

export const ReadRequest = {
  encode(message: ReadRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.resourceName !== "") {
      writer.uint32(10).string(message.resourceName);
    }

    if (!message.readOffset.isZero()) {
      writer.uint32(16).int64(message.readOffset);
    }

    if (!message.readLimit.isZero()) {
      writer.uint32(24).int64(message.readLimit);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ReadRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReadRequest();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.resourceName = reader.string();
          break;

        case 2:
          message.readOffset = (reader.int64() as Long);
          break;

        case 3:
          message.readLimit = (reader.int64() as Long);
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<ReadRequest>): ReadRequest {
    const message = createBaseReadRequest();
    message.resourceName = object.resourceName ?? "";
    message.readOffset = object.readOffset !== undefined && object.readOffset !== null ? Long.fromValue(object.readOffset) : Long.ZERO;
    message.readLimit = object.readLimit !== undefined && object.readLimit !== null ? Long.fromValue(object.readLimit) : Long.ZERO;
    return message;
  }

};

function createBaseReadResponse(): ReadResponse {
  return {
    data: new Uint8Array()
  };
}

export const ReadResponse = {
  encode(message: ReadResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.data.length !== 0) {
      writer.uint32(82).bytes(message.data);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ReadResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReadResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 10:
          message.data = reader.bytes();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<ReadResponse>): ReadResponse {
    const message = createBaseReadResponse();
    message.data = object.data ?? new Uint8Array();
    return message;
  }

};

function createBaseWriteRequest(): WriteRequest {
  return {
    resourceName: "",
    writeOffset: Long.ZERO,
    finishWrite: false,
    data: new Uint8Array()
  };
}

export const WriteRequest = {
  encode(message: WriteRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.resourceName !== "") {
      writer.uint32(10).string(message.resourceName);
    }

    if (!message.writeOffset.isZero()) {
      writer.uint32(16).int64(message.writeOffset);
    }

    if (message.finishWrite === true) {
      writer.uint32(24).bool(message.finishWrite);
    }

    if (message.data.length !== 0) {
      writer.uint32(82).bytes(message.data);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WriteRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWriteRequest();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.resourceName = reader.string();
          break;

        case 2:
          message.writeOffset = (reader.int64() as Long);
          break;

        case 3:
          message.finishWrite = reader.bool();
          break;

        case 10:
          message.data = reader.bytes();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<WriteRequest>): WriteRequest {
    const message = createBaseWriteRequest();
    message.resourceName = object.resourceName ?? "";
    message.writeOffset = object.writeOffset !== undefined && object.writeOffset !== null ? Long.fromValue(object.writeOffset) : Long.ZERO;
    message.finishWrite = object.finishWrite ?? false;
    message.data = object.data ?? new Uint8Array();
    return message;
  }

};

function createBaseWriteResponse(): WriteResponse {
  return {
    committedSize: Long.ZERO
  };
}

export const WriteResponse = {
  encode(message: WriteResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (!message.committedSize.isZero()) {
      writer.uint32(8).int64(message.committedSize);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WriteResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWriteResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.committedSize = (reader.int64() as Long);
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<WriteResponse>): WriteResponse {
    const message = createBaseWriteResponse();
    message.committedSize = object.committedSize !== undefined && object.committedSize !== null ? Long.fromValue(object.committedSize) : Long.ZERO;
    return message;
  }

};

function createBaseQueryWriteStatusRequest(): QueryWriteStatusRequest {
  return {
    resourceName: ""
  };
}

export const QueryWriteStatusRequest = {
  encode(message: QueryWriteStatusRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.resourceName !== "") {
      writer.uint32(10).string(message.resourceName);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryWriteStatusRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryWriteStatusRequest();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.resourceName = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryWriteStatusRequest>): QueryWriteStatusRequest {
    const message = createBaseQueryWriteStatusRequest();
    message.resourceName = object.resourceName ?? "";
    return message;
  }

};

function createBaseQueryWriteStatusResponse(): QueryWriteStatusResponse {
  return {
    committedSize: Long.ZERO,
    complete: false
  };
}

export const QueryWriteStatusResponse = {
  encode(message: QueryWriteStatusResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (!message.committedSize.isZero()) {
      writer.uint32(8).int64(message.committedSize);
    }

    if (message.complete === true) {
      writer.uint32(16).bool(message.complete);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryWriteStatusResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryWriteStatusResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.committedSize = (reader.int64() as Long);
          break;

        case 2:
          message.complete = reader.bool();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<QueryWriteStatusResponse>): QueryWriteStatusResponse {
    const message = createBaseQueryWriteStatusResponse();
    message.committedSize = object.committedSize !== undefined && object.committedSize !== null ? Long.fromValue(object.committedSize) : Long.ZERO;
    message.complete = object.complete ?? false;
    return message;
  }

};