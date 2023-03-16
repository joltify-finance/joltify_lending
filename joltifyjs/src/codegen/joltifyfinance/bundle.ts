import * as _166 from "../joltify/kyc/genesis";
import * as _167 from "../joltify/kyc/investor";
import * as _168 from "../joltify/kyc/params";
import * as _169 from "../joltify/kyc/query";
import * as _170 from "../joltify/kyc/tx";
import * as _318 from "../joltify/kyc/tx.amino";
import * as _319 from "../joltify/kyc/tx.registry";
import * as _320 from "../joltify/kyc/query.lcd";
import * as _321 from "../joltify/kyc/query.rpc.Query";
import * as _322 from "../joltify/kyc/tx.rpc.msg";
import * as _366 from "./lcd";
import * as _367 from "./rpc.query";
import * as _368 from "./rpc.tx";
export namespace joltifyfinance {
  export namespace joltify_lending {
    export const kyc = { ..._166,
      ..._167,
      ..._168,
      ..._169,
      ..._170,
      ..._318,
      ..._319,
      ..._320,
      ..._321,
      ..._322
    };
  }
  export const ClientFactory = { ..._366,
    ..._367,
    ..._368
  };
}