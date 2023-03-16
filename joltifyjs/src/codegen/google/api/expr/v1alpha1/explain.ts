import { Value, ValueSDKType } from "./value";
import * as _m0 from "protobufjs/minimal";
import { DeepPartial, Long } from "../../../../helpers";
/**
 * Values of intermediate expressions produced when evaluating expression.
 * Deprecated, use `EvalState` instead.
 */

/** @deprecated */

export interface Explain {
  /**
   * All of the observed values.
   * 
   * The field value_index is an index in the values list.
   * Separating values from steps is needed to remove redundant values.
   */
  values: Value[];
  /**
   * List of steps.
   * 
   * Repeated evaluations of the same expression generate new ExprStep
   * instances. The order of such ExprStep instances matches the order of
   * elements returned by Comprehension.iter_range.
   */

  exprSteps: Explain_ExprStep[];
}
/**
 * Values of intermediate expressions produced when evaluating expression.
 * Deprecated, use `EvalState` instead.
 */

/** @deprecated */

export interface ExplainSDKType {
  values: ValueSDKType[];
  expr_steps: Explain_ExprStepSDKType[];
}
/** ID and value index of one step. */

export interface Explain_ExprStep {
  /** ID of corresponding Expr node. */
  id: Long;
  /** Index of the value in the values list. */

  valueIndex: number;
}
/** ID and value index of one step. */

export interface Explain_ExprStepSDKType {
  id: Long;
  value_index: number;
}

function createBaseExplain(): Explain {
  return {
    values: [],
    exprSteps: []
  };
}

export const Explain = {
  encode(message: Explain, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.values) {
      Value.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    for (const v of message.exprSteps) {
      Explain_ExprStep.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Explain {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseExplain();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.values.push(Value.decode(reader, reader.uint32()));
          break;

        case 2:
          message.exprSteps.push(Explain_ExprStep.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<Explain>): Explain {
    const message = createBaseExplain();
    message.values = object.values?.map(e => Value.fromPartial(e)) || [];
    message.exprSteps = object.exprSteps?.map(e => Explain_ExprStep.fromPartial(e)) || [];
    return message;
  }

};

function createBaseExplain_ExprStep(): Explain_ExprStep {
  return {
    id: Long.ZERO,
    valueIndex: 0
  };
}

export const Explain_ExprStep = {
  encode(message: Explain_ExprStep, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (!message.id.isZero()) {
      writer.uint32(8).int64(message.id);
    }

    if (message.valueIndex !== 0) {
      writer.uint32(16).int32(message.valueIndex);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Explain_ExprStep {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseExplain_ExprStep();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.id = (reader.int64() as Long);
          break;

        case 2:
          message.valueIndex = reader.int32();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<Explain_ExprStep>): Explain_ExprStep {
    const message = createBaseExplain_ExprStep();
    message.id = object.id !== undefined && object.id !== null ? Long.fromValue(object.id) : Long.ZERO;
    message.valueIndex = object.valueIndex ?? 0;
    return message;
  }

};