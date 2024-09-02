//@ts-nocheck
import * as _96 from "./act/module/module.js";
import * as _97 from "./act/v1beta1/action.js";
import * as _98 from "./act/v1beta1/events.js";
import * as _99 from "./act/v1beta1/genesis.js";
import * as _100 from "./act/v1beta1/params.js";
import * as _101 from "./act/v1beta1/query.js";
import * as _102 from "./act/v1beta1/rule.js";
import * as _103 from "./act/v1beta1/tx.js";
import * as _104 from "./gmp/genesis.js";
import * as _105 from "./gmp/gmp.js";
import * as _106 from "./gmp/query.js";
import * as _107 from "./gmp/tx.js";
import * as _108 from "./warden/module/module.js";
import * as _109 from "./warden/v1beta3/events.js";
import * as _110 from "./warden/v1beta3/genesis.js";
import * as _111 from "./warden/v1beta3/key.js";
import * as _112 from "./warden/v1beta3/keychain.js";
import * as _113 from "./warden/v1beta3/params.js";
import * as _114 from "./warden/v1beta3/query.js";
import * as _115 from "./warden/v1beta3/signature.js";
import * as _116 from "./warden/v1beta3/space.js";
import * as _117 from "./warden/v1beta3/tx.js";
import * as _186 from "./act/v1beta1/tx.amino.js";
import * as _187 from "./gmp/tx.amino.js";
import * as _188 from "./warden/v1beta3/tx.amino.js";
import * as _189 from "./act/v1beta1/tx.registry.js";
import * as _190 from "./gmp/tx.registry.js";
import * as _191 from "./warden/v1beta3/tx.registry.js";
import * as _192 from "./act/v1beta1/query.lcd.js";
import * as _193 from "./gmp/query.lcd.js";
import * as _194 from "./warden/v1beta3/query.lcd.js";
import * as _195 from "./act/v1beta1/query.rpc.Query.js";
import * as _196 from "./gmp/query.rpc.Query.js";
import * as _197 from "./warden/v1beta3/query.rpc.Query.js";
import * as _198 from "./act/v1beta1/tx.rpc.msg.js";
import * as _199 from "./gmp/tx.rpc.msg.js";
import * as _200 from "./warden/v1beta3/tx.rpc.msg.js";
import * as _207 from "./lcd.js";
import * as _208 from "./rpc.query.js";
import * as _209 from "./rpc.tx.js";
export namespace warden {
  export namespace act {
    export const module = {
      ..._96
    };
    export const v1beta1 = {
      ..._97,
      ..._98,
      ..._99,
      ..._100,
      ..._101,
      ..._102,
      ..._103,
      ..._186,
      ..._189,
      ..._192,
      ..._195,
      ..._198
    };
  }
  export const gmp = {
    ..._104,
    ..._105,
    ..._106,
    ..._107,
    ..._187,
    ..._190,
    ..._193,
    ..._196,
    ..._199
  };
  export namespace warden {
    export const module = {
      ..._108
    };
    export const v1beta3 = {
      ..._109,
      ..._110,
      ..._111,
      ..._112,
      ..._113,
      ..._114,
      ..._115,
      ..._116,
      ..._117,
      ..._188,
      ..._191,
      ..._194,
      ..._197,
      ..._200
    };
  }
  export const ClientFactory = {
    ..._207,
    ..._208,
    ..._209
  };
}