import { Expr, ExprSDKType } from "./expr";
import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../../../helpers";
/** A declaration. */

export interface Decl {
  /** The id of the declaration. */
  id: number;
  /** The name of the declaration. */

  name: string;
  /** The documentation string for the declaration. */

  doc: string;
  /** An identifier declaration. */

  ident?: IdentDecl;
  /** A function declaration. */

  function?: FunctionDecl;
}
/** A declaration. */

export interface DeclSDKType {
  id: number;
  name: string;
  doc: string;
  ident?: IdentDeclSDKType;
  function?: FunctionDeclSDKType;
}
/**
 * The declared type of a variable.
 * 
 * Extends runtime type values with extra information used for type checking
 * and dispatching.
 */

export interface DeclType {
  /** The expression id of the declared type, if applicable. */
  id: number;
  /** The type name, e.g. 'int', 'my.type.Type' or 'T' */

  type: string;
  /**
   * An ordered list of type parameters, e.g. `<string, int>`.
   * Only applies to a subset of types, e.g. `map`, `list`.
   */

  typeParams: DeclType[];
}
/**
 * The declared type of a variable.
 * 
 * Extends runtime type values with extra information used for type checking
 * and dispatching.
 */

export interface DeclTypeSDKType {
  id: number;
  type: string;
  type_params: DeclTypeSDKType[];
}
/** An identifier declaration. */

export interface IdentDecl {
  /** Optional type of the identifier. */
  type?: DeclType;
  /** Optional value of the identifier. */

  value?: Expr;
}
/** An identifier declaration. */

export interface IdentDeclSDKType {
  type?: DeclTypeSDKType;
  value?: ExprSDKType;
}
/** A function declaration. */

export interface FunctionDecl {
  /** The function arguments. */
  args: IdentDecl[];
  /** Optional declared return type. */

  returnType?: DeclType;
  /** If the first argument of the function is the receiver. */

  receiverFunction: boolean;
}
/** A function declaration. */

export interface FunctionDeclSDKType {
  args: IdentDeclSDKType[];
  return_type?: DeclTypeSDKType;
  receiver_function: boolean;
}

function createBaseDecl(): Decl {
  return {
    id: 0,
    name: "",
    doc: "",
    ident: undefined,
    function: undefined
  };
}

export const Decl = {
  encode(message: Decl, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).int32(message.id);
    }

    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }

    if (message.doc !== "") {
      writer.uint32(26).string(message.doc);
    }

    if (message.ident !== undefined) {
      IdentDecl.encode(message.ident, writer.uint32(34).fork()).ldelim();
    }

    if (message.function !== undefined) {
      FunctionDecl.encode(message.function, writer.uint32(42).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Decl {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDecl();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.id = reader.int32();
          break;

        case 2:
          message.name = reader.string();
          break;

        case 3:
          message.doc = reader.string();
          break;

        case 4:
          message.ident = IdentDecl.decode(reader, reader.uint32());
          break;

        case 5:
          message.function = FunctionDecl.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<Decl>): Decl {
    const message = createBaseDecl();
    message.id = object.id ?? 0;
    message.name = object.name ?? "";
    message.doc = object.doc ?? "";
    message.ident = object.ident !== undefined && object.ident !== null ? IdentDecl.fromPartial(object.ident) : undefined;
    message.function = object.function !== undefined && object.function !== null ? FunctionDecl.fromPartial(object.function) : undefined;
    return message;
  }

};

function createBaseDeclType(): DeclType {
  return {
    id: 0,
    type: "",
    typeParams: []
  };
}

export const DeclType = {
  encode(message: DeclType, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).int32(message.id);
    }

    if (message.type !== "") {
      writer.uint32(18).string(message.type);
    }

    for (const v of message.typeParams) {
      DeclType.encode(v!, writer.uint32(34).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeclType {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeclType();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.id = reader.int32();
          break;

        case 2:
          message.type = reader.string();
          break;

        case 4:
          message.typeParams.push(DeclType.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<DeclType>): DeclType {
    const message = createBaseDeclType();
    message.id = object.id ?? 0;
    message.type = object.type ?? "";
    message.typeParams = object.typeParams?.map(e => DeclType.fromPartial(e)) || [];
    return message;
  }

};

function createBaseIdentDecl(): IdentDecl {
  return {
    type: undefined,
    value: undefined
  };
}

export const IdentDecl = {
  encode(message: IdentDecl, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.type !== undefined) {
      DeclType.encode(message.type, writer.uint32(26).fork()).ldelim();
    }

    if (message.value !== undefined) {
      Expr.encode(message.value, writer.uint32(34).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IdentDecl {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIdentDecl();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 3:
          message.type = DeclType.decode(reader, reader.uint32());
          break;

        case 4:
          message.value = Expr.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<IdentDecl>): IdentDecl {
    const message = createBaseIdentDecl();
    message.type = object.type !== undefined && object.type !== null ? DeclType.fromPartial(object.type) : undefined;
    message.value = object.value !== undefined && object.value !== null ? Expr.fromPartial(object.value) : undefined;
    return message;
  }

};

function createBaseFunctionDecl(): FunctionDecl {
  return {
    args: [],
    returnType: undefined,
    receiverFunction: false
  };
}

export const FunctionDecl = {
  encode(message: FunctionDecl, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.args) {
      IdentDecl.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    if (message.returnType !== undefined) {
      DeclType.encode(message.returnType, writer.uint32(18).fork()).ldelim();
    }

    if (message.receiverFunction === true) {
      writer.uint32(24).bool(message.receiverFunction);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FunctionDecl {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFunctionDecl();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.args.push(IdentDecl.decode(reader, reader.uint32()));
          break;

        case 2:
          message.returnType = DeclType.decode(reader, reader.uint32());
          break;

        case 3:
          message.receiverFunction = reader.bool();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<FunctionDecl>): FunctionDecl {
    const message = createBaseFunctionDecl();
    message.args = object.args?.map(e => IdentDecl.fromPartial(e)) || [];
    message.returnType = object.returnType !== undefined && object.returnType !== null ? DeclType.fromPartial(object.returnType) : undefined;
    message.receiverFunction = object.receiverFunction ?? false;
    return message;
  }

};