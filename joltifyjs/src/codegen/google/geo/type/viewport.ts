import { LatLng, LatLngSDKType } from "../../type/latlng";
import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../../helpers";
/**
 * A latitude-longitude viewport, represented as two diagonally opposite `low`
 * and `high` points. A viewport is considered a closed region, i.e. it includes
 * its boundary. The latitude bounds must range between -90 to 90 degrees
 * inclusive, and the longitude bounds must range between -180 to 180 degrees
 * inclusive. Various cases include:
 * 
 *  - If `low` = `high`, the viewport consists of that single point.
 * 
 *  - If `low.longitude` > `high.longitude`, the longitude range is inverted
 *    (the viewport crosses the 180 degree longitude line).
 * 
 *  - If `low.longitude` = -180 degrees and `high.longitude` = 180 degrees,
 *    the viewport includes all longitudes.
 * 
 *  - If `low.longitude` = 180 degrees and `high.longitude` = -180 degrees,
 *    the longitude range is empty.
 * 
 *  - If `low.latitude` > `high.latitude`, the latitude range is empty.
 * 
 * Both `low` and `high` must be populated, and the represented box cannot be
 * empty (as specified by the definitions above). An empty viewport will result
 * in an error.
 * 
 * For example, this viewport fully encloses New York City:
 * 
 * {
 *     "low": {
 *         "latitude": 40.477398,
 *         "longitude": -74.259087
 *     },
 *     "high": {
 *         "latitude": 40.91618,
 *         "longitude": -73.70018
 *     }
 * }
 */

export interface Viewport {
  /** Required. The low point of the viewport. */
  low?: LatLng;
  /** Required. The high point of the viewport. */

  high?: LatLng;
}
/**
 * A latitude-longitude viewport, represented as two diagonally opposite `low`
 * and `high` points. A viewport is considered a closed region, i.e. it includes
 * its boundary. The latitude bounds must range between -90 to 90 degrees
 * inclusive, and the longitude bounds must range between -180 to 180 degrees
 * inclusive. Various cases include:
 * 
 *  - If `low` = `high`, the viewport consists of that single point.
 * 
 *  - If `low.longitude` > `high.longitude`, the longitude range is inverted
 *    (the viewport crosses the 180 degree longitude line).
 * 
 *  - If `low.longitude` = -180 degrees and `high.longitude` = 180 degrees,
 *    the viewport includes all longitudes.
 * 
 *  - If `low.longitude` = 180 degrees and `high.longitude` = -180 degrees,
 *    the longitude range is empty.
 * 
 *  - If `low.latitude` > `high.latitude`, the latitude range is empty.
 * 
 * Both `low` and `high` must be populated, and the represented box cannot be
 * empty (as specified by the definitions above). An empty viewport will result
 * in an error.
 * 
 * For example, this viewport fully encloses New York City:
 * 
 * {
 *     "low": {
 *         "latitude": 40.477398,
 *         "longitude": -74.259087
 *     },
 *     "high": {
 *         "latitude": 40.91618,
 *         "longitude": -73.70018
 *     }
 * }
 */

export interface ViewportSDKType {
  low?: LatLngSDKType;
  high?: LatLngSDKType;
}

function createBaseViewport(): Viewport {
  return {
    low: undefined,
    high: undefined
  };
}

export const Viewport = {
  encode(message: Viewport, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.low !== undefined) {
      LatLng.encode(message.low, writer.uint32(10).fork()).ldelim();
    }

    if (message.high !== undefined) {
      LatLng.encode(message.high, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Viewport {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseViewport();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.low = LatLng.decode(reader, reader.uint32());
          break;

        case 2:
          message.high = LatLng.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<Viewport>): Viewport {
    const message = createBaseViewport();
    message.low = object.low !== undefined && object.low !== null ? LatLng.fromPartial(object.low) : undefined;
    message.high = object.high !== undefined && object.high !== null ? LatLng.fromPartial(object.high) : undefined;
    return message;
  }

};