import * as _215 from "./abci/types";
import * as _216 from "./crypto/keys";
import * as _217 from "./crypto/proof";
import * as _218 from "./libs/bits/types";
import * as _219 from "./p2p/types";
import * as _220 from "./types/block";
import * as _221 from "./types/evidence";
import * as _222 from "./types/params";
import * as _223 from "./types/types";
import * as _224 from "./types/validator";
import * as _225 from "./version/types";
export namespace tendermint {
  export const abci = { ..._215
  };
  export const crypto = { ..._216,
    ..._217
  };
  export namespace libs {
    export const bits = { ..._218
    };
  }
  export const p2p = { ..._219
  };
  export const types = { ..._220,
    ..._221,
    ..._222,
    ..._223,
    ..._224
  };
  export const version = { ..._225
  };
}