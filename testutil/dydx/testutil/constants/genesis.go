package constants

// This is a copy of the localnet genesis.json. This can be retrieved from the localnet docker container path:
// /dydxprotocol/chain/.alice/config/genesis.json
// Disable linter for exchange config.
//
//nolint:all
const GenesisState = `
{ 
  "genesis_time": "2023-01-01T00:00:00Z",
  "chain_id": "localdydxprotocol",
  "initial_height": "1",
  "app_hash": null,
  "app_state": {
    "assets": {
      "assets": [
        {
          "id": 0,
          "symbol": "USDC",
          "denom": "ibc/8E27BA2D5493AF5636760E354E46004562C46AB7EC0CC4C1CA14E9E20E2545B5",
          "denom_exponent": "-6",
          "has_market": false,
          "market_id": 0,
          "atomic_resolution": -6
        }
      ]
    },
    "auth": {
      "params": {
        "max_memo_characters": "256",
        "tx_sig_limit": "7",
        "tx_size_cost_per_byte": "10",
        "sig_verify_cost_ed25519": "590",
        "sig_verify_cost_secp256k1": "1000"
      },
      "accounts": [
        {
          "@type": "/cosmos.auth.v1beta1.BaseAccount",
          "address": "jolt1x4p6mruqfctlprtde8hg9jz2xq5m798f50yuuu",
          "pub_key": null,
          "account_number": "0",
          "sequence": "0"
        },
        {
          "@type": "/cosmos.auth.v1beta1.BaseAccount",
          "address": "jolt199tqg4wdlnu4qjlxchpd7seg454937hjq0q20t",
          "pub_key": null,
          "account_number": "1",
          "sequence": "0"
        },
        {
          "@type": "/cosmos.auth.v1beta1.BaseAccount",
          "address": "jolt10fx7sy6ywd5senxae9dwytf8jxek3t2gmqqjlw",
          "pub_key": null,
          "account_number": "2",
          "sequence": "0"
        },
        {
          "@type": "/cosmos.auth.v1beta1.BaseAccount",
          "address": "jolt1fjg6zp6vv8t9wvy4lps03r5l4g7tkjw9d4g0d3",
          "pub_key": null,
          "account_number": "3",
          "sequence": "0"
        },
        {
          "@type": "/cosmos.auth.v1beta1.BaseAccount",
          "address": "jolt1wau5mja7j7zdavtfq9lu7ejef05hm6ff62vqrd",
          "pub_key": null,
          "account_number": "4",
          "sequence": "0"
        },
        {
          "@type": "/cosmos.auth.v1beta1.BaseAccount",
          "address": "jolt13gcfjapx049mhh52w7kucqcu0vva8vxnkwdqqq",
          "pub_key": null,
          "account_number": "5",
          "sequence": "0"
        }
      ]
    },
    "authz": {
      "authorization": []
    },
    "bank": {
      "params": {
        "send_enabled": [],
        "default_send_enabled": true
      },
      "balances": [
        {
          "address": "jolt199tqg4wdlnu4qjlxchpd7seg454937hjq0q20t",
          "coins": [
            {
              "denom": "adv4tnt",
              "amount": "1000000000000000000000000"
            },
            {
              "denom": "ibc/8E27BA2D5493AF5636760E354E46004562C46AB7EC0CC4C1CA14E9E20E2545B5",
              "amount": "100000000000000000"
            }
          ]
        },
        {
          "address": "jolt1x4p6mruqfctlprtde8hg9jz2xq5m798f50yuuu",
          "coins": [
            {
              "denom": "adv4tnt",
              "amount": "1000000000000000000000000"
            },
            {
              "denom": "ibc/8E27BA2D5493AF5636760E354E46004562C46AB7EC0CC4C1CA14E9E20E2545B5",
              "amount": "100000000000000000"
            }
          ]
        },
        {
          "address": "jolt1xmzu58dpmvwatv0fqfkhtytq6tzdc9xc7nzmmm",
          "coins": [
            {
              "denom": "adv4tnt",
              "amount": "200000000000000000000000000"
            }
          ]
        },
        {
          "address": "jolt1gjggsthlt0usfuwp8fuu3vrh3v9276hu0hk000",
          "coins": [
            {
              "denom": "ibc/8E27BA2D5493AF5636760E354E46004562C46AB7EC0CC4C1CA14E9E20E2545B5",
              "amount": "900000001000000000"
            }
          ]
        },
        {
          "address": "jolt1fjg6zp6vv8t9wvy4lps03r5l4g7tkjw9d4g0d3",
          "coins": [
            {
              "denom": "adv4tnt",
              "amount": "1000000000000000000000000"
            },
            {
              "denom": "ibc/8E27BA2D5493AF5636760E354E46004562C46AB7EC0CC4C1CA14E9E20E2545B5",
              "amount": "100000000000000000"
            }
          ]
        },
        {
          "address": "jolt1wau5mja7j7zdavtfq9lu7ejef05hm6ff62vqrd",
          "coins": [
            {
              "denom": "adv4tnt",
              "amount": "1000000000000000000000000"
            },
            {
              "denom": "ibc/8E27BA2D5493AF5636760E354E46004562C46AB7EC0CC4C1CA14E9E20E2545B5",
              "amount": "100000000000000000"
            }
          ]
        },
        {
          "address": "jolt10fx7sy6ywd5senxae9dwytf8jxek3t2gmqqjlw",
          "coins": [
            {
              "denom": "adv4tnt",
              "amount": "1000000000000000000000000"
            },
            {
              "denom": "ibc/8E27BA2D5493AF5636760E354E46004562C46AB7EC0CC4C1CA14E9E20E2545B5",
              "amount": "100000000000000000"
            }
          ]
        },
        {
          "address": "jolt13gcfjapx049mhh52w7kucqcu0vva8vxnkwdqqq",
          "coins": [
            {
              "denom": "adv4tnt",
              "amount": "1000000000000000000000000"
            },
            {
              "denom": "ibc/8E27BA2D5493AF5636760E354E46004562C46AB7EC0CC4C1CA14E9E20E2545B5",
              "amount": "900000000000000000"
            }
          ]
        },
        {
          "address": "jolt1mc6teylz69jdvj9w9tm9v2tmztxs4q8ce85ggg",
          "coins": [
            {
              "denom": "adv4tnt",
              "amount": "799000000000000000000000000"
            }
          ]
        }
      ],
      "supply": [],
      "denom_metadata": [
        {
          "description": "The native token of the network",
          "denom_units": [
            {
              "denom": "adv4tnt",
              "exponent": 0,
              "aliases": []
            },
            {
              "denom": "dv4tnt",
              "exponent": 18,
              "aliases": []
            }
          ],
          "base": "adv4tnt",
          "display": "dv4tnt",
          "name": "dYdX V4 Testnet Token",
          "symbol": "dv4tnt",
          "uri": "",
          "uri_hash": ""
        }
      ],
      "send_enabled": []
    },
    "blocktime": {
      "params": {
        "durations": [
          "300s",
          "1800s"
        ]
      }
    },
    "bridge": {
      "event_params": {
        "denom": "adv4tnt",
        "eth_chain_id": 11155111,
        "eth_address": "0xf75012c350e4ad55be2048bd67ce6e03b20de82d"
      },
      "propose_params": {
        "max_bridges_per_block": 10,
        "propose_delay_duration": "60s",
        "skip_rate_ppm": 800000,
        "skip_if_block_delayed_by_duration": "5s"
      },
      "safety_params": {
        "is_disabled": false,
        "delay_blocks": 30
      },
      "acknowledged_event_info": {
        "next_id": 5,
        "eth_block_height": 4322136
      }
    },
    "capability": {
      "index": "1",
      "owners": []
    },
    "clob": {
      "clob_pairs": [
        {
          "id": 0,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 0
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 100000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 1,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 1
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 100000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 2,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 2
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 3,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 3
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 4,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 4
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 5,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 5
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 6,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 6
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 7,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 7
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 8,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 8
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 9,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 9
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 10,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 10
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 11,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 11
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 12,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 12
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 13,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 13
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 14,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 14
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 15,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 15
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 16,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 16
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 17,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 17
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 18,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 18
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 19,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 19
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 20,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 20
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 21,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 21
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 22,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 22
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 23,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 23
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 24,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 24
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 25,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 25
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 26,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 26
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 27,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 27
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 28,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 28
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 29,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 29
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 30,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 30
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 31,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 31
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 32,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 32
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 1000000,
          "quantum_conversion_exponent": -9
        },
        {
          "id": 33,
          "status": "STATUS_ACTIVE",
          "perpetual_clob_metadata": {
            "perpetual_id": 33
          },
          "step_base_quantums": 1000000,
          "subticks_per_tick": 100,
          "quantum_conversion_exponent": -8
        }
      ],
      "liquidations_config": {
        "max_liquidation_fee_ppm": 15000,
        "position_block_limits": {
          "min_position_notional_liquidated": 1000000000,
          "max_position_portion_liquidated_ppm": 100000
        },
        "subaccount_block_limits": {
          "max_notional_liquidated": 100000000000,
          "max_quantums_insurance_lost": 1000000000000
        },
        "fillable_price_config": {
          "bankruptcy_adjustment_ppm": 1000000,
          "spread_to_maintenance_margin_ratio_ppm": 1500000
        }
      },
      "block_rate_limit_config": {
        "max_short_term_orders_per_n_blocks": [],
        "max_stateful_orders_per_n_blocks": [
          {
            "limit": 2,
            "num_blocks": 1
          },
          {
            "limit": 20,
            "num_blocks": 100
          }
        ],
        "max_short_term_order_cancellations_per_n_blocks": [],
        "max_short_term_orders_and_cancels_per_n_blocks": [
          {
            "limit": 400,
            "num_blocks": 1
          }
        ]
      },
      "equity_tier_limit_config": {
        "short_term_order_equity_tiers": [
          {
            "limit": 0,
            "usd_tnc_required": "0"
          },
          {
            "limit": 1,
            "usd_tnc_required": "20000000"
          },
          {
            "limit": 5,
            "usd_tnc_required": "100000000"
          },
          {
            "limit": 10,
            "usd_tnc_required": "1000000000"
          },
          {
            "limit": 100,
            "usd_tnc_required": "10000000000"
          },
          {
            "limit": 1000,
            "usd_tnc_required": "100000000000"
          }
        ],
        "stateful_order_equity_tiers": [
          {
            "limit": 0,
            "usd_tnc_required": "0"
          },
          {
            "limit": 1,
            "usd_tnc_required": "20000000"
          },
          {
            "limit": 5,
            "usd_tnc_required": "100000000"
          },
          {
            "limit": 10,
            "usd_tnc_required": "1000000000"
          },
          {
            "limit": 100,
            "usd_tnc_required": "10000000000"
          },
          {
            "limit": 200,
            "usd_tnc_required": "100000000000"
          }
        ]
      }
    },
    "crisis": {
      "constant_fee": {
        "denom": "adv4tnt",
        "amount": "1000"
      }
    },
    "delaymsg": {
      "delayed_messages": [
        {
          "id": 0,
          "msg": {
            "@type": "/dydxprotocol.feetiers.MsgUpdatePerpetualFeeParams",
            "authority": "dydx1mkkvp26dngu6n8rmalaxyp3gwkjuzztq5zx6tr",
            "params": {
              "tiers": [
                {
                  "name": "1",
                  "absolute_volume_requirement": "0",
                  "total_volume_share_requirement_ppm": 0,
                  "maker_volume_share_requirement_ppm": 0,
                  "maker_fee_ppm": 100,
                  "taker_fee_ppm": 500
                },
                {
                  "name": "2",
                  "absolute_volume_requirement": "1000000000000",
                  "total_volume_share_requirement_ppm": 0,
                  "maker_volume_share_requirement_ppm": 0,
                  "maker_fee_ppm": 100,
                  "taker_fee_ppm": 450
                },
                {
                  "name": "3",
                  "absolute_volume_requirement": "5000000000000",
                  "total_volume_share_requirement_ppm": 0,
                  "maker_volume_share_requirement_ppm": 0,
                  "maker_fee_ppm": 50,
                  "taker_fee_ppm": 400
                },
                {
                  "name": "4",
                  "absolute_volume_requirement": "25000000000000",
                  "total_volume_share_requirement_ppm": 0,
                  "maker_volume_share_requirement_ppm": 0,
                  "maker_fee_ppm": 0,
                  "taker_fee_ppm": 350
                },
                {
                  "name": "5",
                  "absolute_volume_requirement": "125000000000000",
                  "total_volume_share_requirement_ppm": 0,
                  "maker_volume_share_requirement_ppm": 0,
                  "maker_fee_ppm": 0,
                  "taker_fee_ppm": 300
                },
                {
                  "name": "6",
                  "absolute_volume_requirement": "125000000000000",
                  "total_volume_share_requirement_ppm": 5000,
                  "maker_volume_share_requirement_ppm": 0,
                  "maker_fee_ppm": -50,
                  "taker_fee_ppm": 250
                },
                {
                  "name": "7",
                  "absolute_volume_requirement": "125000000000000",
                  "total_volume_share_requirement_ppm": 5000,
                  "maker_volume_share_requirement_ppm": 10000,
                  "maker_fee_ppm": -70,
                  "taker_fee_ppm": 250
                },
                {
                  "name": "8",
                  "absolute_volume_requirement": "125000000000000",
                  "total_volume_share_requirement_ppm": 5000,
                  "maker_volume_share_requirement_ppm": 20000,
                  "maker_fee_ppm": -90,
                  "taker_fee_ppm": 250
                },
                {
                  "name": "9",
                  "absolute_volume_requirement": "125000000000000",
                  "total_volume_share_requirement_ppm": 5000,
                  "maker_volume_share_requirement_ppm": 40000,
                  "maker_fee_ppm": -110,
                  "taker_fee_ppm": 250
                }
              ]
            }
          },
          "block_height": 378000
        }
      ],
      "next_delayed_message_id": 1
    },
    "distribution": {
      "params": {
        "community_tax": "0.020000000000000000",
        "base_proposer_reward": "0.000000000000000000",
        "bonus_proposer_reward": "0.000000000000000000",
        "withdraw_addr_enabled": true
      },
      "fee_pool": {
        "community_pool": []
      },
      "delegator_withdraw_infos": [],
      "previous_proposer": "",
      "outstanding_rewards": [],
      "validator_accumulated_commissions": [],
      "validator_historical_rewards": [],
      "validator_current_rewards": [],
      "delegator_starting_infos": [],
      "validator_slash_events": []
    },
    "dydxaccountplus": {
      "accounts": []
    },
    "epochs": {
      "epoch_info_list": [
        {
          "name": "funding-sample",
          "next_tick": 30,
          "duration": 60,
          "current_epoch": 0,
          "current_epoch_start_block": 0,
          "is_initialized": false,
          "fast_forward_next_tick": true
        },
        {
          "name": "funding-tick",
          "next_tick": 0,
          "duration": 3600,
          "current_epoch": 0,
          "current_epoch_start_block": 0,
          "is_initialized": false,
          "fast_forward_next_tick": true
        },
        {
          "name": "stats-epoch",
          "next_tick": 0,
          "duration": 3600,
          "current_epoch": 0,
          "current_epoch_start_block": 0,
          "is_initialized": false,
          "fast_forward_next_tick": true
        }
      ]
    },
    "evidence": {
      "evidence": []
    },
    "feegrant": {
      "allowances": []
    },
    "feetiers": {
      "params": {
        "tiers": [
          {
            "name": "1",
            "absolute_volume_requirement": "0",
            "total_volume_share_requirement_ppm": 0,
            "maker_volume_share_requirement_ppm": 0,
            "maker_fee_ppm": -110,
            "taker_fee_ppm": 500
          },
          {
            "name": "2",
            "absolute_volume_requirement": "1000000000000",
            "total_volume_share_requirement_ppm": 0,
            "maker_volume_share_requirement_ppm": 0,
            "maker_fee_ppm": -110,
            "taker_fee_ppm": 450
          },
          {
            "name": "3",
            "absolute_volume_requirement": "5000000000000",
            "total_volume_share_requirement_ppm": 0,
            "maker_volume_share_requirement_ppm": 0,
            "maker_fee_ppm": -110,
            "taker_fee_ppm": 400
          },
          {
            "name": "4",
            "absolute_volume_requirement": "25000000000000",
            "total_volume_share_requirement_ppm": 0,
            "maker_volume_share_requirement_ppm": 0,
            "maker_fee_ppm": -110,
            "taker_fee_ppm": 350
          },
          {
            "name": "5",
            "absolute_volume_requirement": "125000000000000",
            "total_volume_share_requirement_ppm": 0,
            "maker_volume_share_requirement_ppm": 0,
            "maker_fee_ppm": -110,
            "taker_fee_ppm": 300
          },
          {
            "name": "6",
            "absolute_volume_requirement": "125000000000000",
            "total_volume_share_requirement_ppm": 5000,
            "maker_volume_share_requirement_ppm": 0,
            "maker_fee_ppm": -110,
            "taker_fee_ppm": 250
          },
          {
            "name": "7",
            "absolute_volume_requirement": "125000000000000",
            "total_volume_share_requirement_ppm": 5000,
            "maker_volume_share_requirement_ppm": 10000,
            "maker_fee_ppm": -110,
            "taker_fee_ppm": 250
          },
          {
            "name": "8",
            "absolute_volume_requirement": "125000000000000",
            "total_volume_share_requirement_ppm": 5000,
            "maker_volume_share_requirement_ppm": 20000,
            "maker_fee_ppm": -110,
            "taker_fee_ppm": 250
          },
          {
            "name": "9",
            "absolute_volume_requirement": "125000000000000",
            "total_volume_share_requirement_ppm": 5000,
            "maker_volume_share_requirement_ppm": 40000,
            "maker_fee_ppm": -110,
            "taker_fee_ppm": 250
          }
        ]
      }
    },
    "genutil": {
      "gen_txs": [
        {
          "body": {
            "messages": [
              {
                "@type": "/cosmos.staking.v1beta1.MsgCreateValidator",
                "description": {
                  "moniker": "alice",
                  "identity": "",
                  "website": "",
                  "security_contact": "",
                  "details": ""
                },
                "commission": {
                  "rate": "1.000000000000000000",
                  "max_rate": "1.000000000000000000",
                  "max_change_rate": "0.010000000000000000"
                },
                "min_self_delegation": "1",
                "delegator_address": "",
                "validator_address": "joltvaloper199tqg4wdlnu4qjlxchpd7seg454937hjh86377",
                "pubkey": {
                  "@type": "/cosmos.crypto.ed25519.PubKey",
                  "key": "YiARx8259Z+fGFUxQLrz/5FU2RYRT6f5yzvt7D7CrQM="
                },
                "value": {
                  "denom": "adv4tnt",
                  "amount": "500000000000000000000000"
                }
              }
            ],
            "memo": "17e5e45691f0d01449c84fd4ae87279578cdd7ec@172.17.0.2:26656",
            "timeout_height": "0",
            "extension_options": [],
            "non_critical_extension_options": []
          },
          "auth_info": {
            "signer_infos": [
              {
                "public_key": {
                  "@type": "/cosmos.crypto.secp256k1.PubKey",
                  "key": "A0iQ+HpUfJGcgcH7iiEzY9VwCYWCTwg5LsTjc/q1XwSc"
                },
                "mode_info": {
                  "single": {
                    "mode": "SIGN_MODE_DIRECT"
                  }
                },
                "sequence": "0"
              }
            ],
            "fee": {
              "amount": [],
              "gas_limit": "200000",
              "payer": "",
              "granter": ""
            },
            "tip": null
          },
          "signatures": [
            "LPiDilA3vVxoWaLRUQOF63j0eqxvGu3gqxOhjo6OfW0+YTkjaczYdIcUWmdDRC+ge+mS4OIwkTJaUAKq8aZm+Q=="
          ]
        },
        {
          "body": {
            "messages": [
              {
                "@type": "/cosmos.staking.v1beta1.MsgCreateValidator",
                "description": {
                  "moniker": "carl",
                  "identity": "",
                  "website": "",
                  "security_contact": "",
                  "details": ""
                },
                "commission": {
                  "rate": "1.000000000000000000",
                  "max_rate": "1.000000000000000000",
                  "max_change_rate": "0.010000000000000000"
                },
                "min_self_delegation": "1",
                "delegator_address": "",
                "validator_address": "joltvaloper1fjg6zp6vv8t9wvy4lps03r5l4g7tkjw96aj5uy",
                "pubkey": {
                  "@type": "/cosmos.crypto.ed25519.PubKey",
                  "key": "ytLfs1W6E2I41iteKC/YwjyZ/51+CAYCHYxmRHiBeY4="
                },
                "value": {
                  "denom": "adv4tnt",
                  "amount": "500000000000000000000000"
                }
              }
            ],
            "memo": "47539956aaa8e624e0f1d926040e54908ad0eb44@172.17.0.2:26656",
            "timeout_height": "0",
            "extension_options": [],
            "non_critical_extension_options": []
          },
          "auth_info": {
            "signer_infos": [
              {
                "public_key": {
                  "@type": "/cosmos.crypto.secp256k1.PubKey",
                  "key": "AkA1fsLUhCSWbnemBIAR9CPkK1Ra1LlYZcrAKm/Ymvqn"
                },
                "mode_info": {
                  "single": {
                    "mode": "SIGN_MODE_DIRECT"
                  }
                },
                "sequence": "0"
              }
            ],
            "fee": {
              "amount": [],
              "gas_limit": "200000",
              "payer": "",
              "granter": ""
            },
            "tip": null
          },
          "signatures": [
            "8Ct5briHcvxvhVNdZJJ4o9DlJ7YbXIea6uiRxGESVIkSSfPeZJB7raFxPvKbXQeQjVzun0S3BJWmhtqRLzMxlA=="
          ]
        },
        {
          "body": {
            "messages": [
              {
                "@type": "/cosmos.staking.v1beta1.MsgCreateValidator",
                "description": {
                  "moniker": "dave",
                  "identity": "",
                  "website": "",
                  "security_contact": "",
                  "details": ""
                },
                "commission": {
                  "rate": "1.000000000000000000",
                  "max_rate": "1.000000000000000000",
                  "max_change_rate": "0.010000000000000000"
                },
                "min_self_delegation": "1",
                "delegator_address": "",
                "validator_address": "joltvaloper1wau5mja7j7zdavtfq9lu7ejef05hm6ffdzkmjc",
                "pubkey": {
                  "@type": "/cosmos.crypto.ed25519.PubKey",
                  "key": "yG29kRfZ/hgAE1I7uWjbKQJJL4/gX/05XBnfB+m196A="
                },
                "value": {
                  "denom": "adv4tnt",
                  "amount": "500000000000000000000000"
                }
              }
            ],
            "memo": "5882428984d83b03d0c907c1f0af343534987052@172.17.0.2:26656",
            "timeout_height": "0",
            "extension_options": [],
            "non_critical_extension_options": []
          },
          "auth_info": {
            "signer_infos": [
              {
                "public_key": {
                  "@type": "/cosmos.crypto.secp256k1.PubKey",
                  "key": "A87MchHGMj7i1xBwUfECtXzXJIgli/JVFoSaxUqIN86R"
                },
                "mode_info": {
                  "single": {
                    "mode": "SIGN_MODE_DIRECT"
                  }
                },
                "sequence": "0"
              }
            ],
            "fee": {
              "amount": [],
              "gas_limit": "200000",
              "payer": "",
              "granter": ""
            },
            "tip": null
          },
          "signatures": [
            "Owy37CpJt2XyFNAulUQ34Bpa9iIRMEQRgxD4dbsI9htg/xi3n9TShiVezHqRRsKUa89hpctiVqY33W3hEniLiA=="
          ]
        },
        {
          "body": {
            "messages": [
              {
                "@type": "/cosmos.staking.v1beta1.MsgCreateValidator",
                "description": {
                  "moniker": "bob",
                  "identity": "",
                  "website": "",
                  "security_contact": "",
                  "details": ""
                },
                "commission": {
                  "rate": "1.000000000000000000",
                  "max_rate": "1.000000000000000000",
                  "max_change_rate": "0.010000000000000000"
                },
                "min_self_delegation": "1",
                "delegator_address": "",
                "validator_address": "joltvaloper10fx7sy6ywd5senxae9dwytf8jxek3t2gvg6fwm",
                "pubkey": {
                  "@type": "/cosmos.crypto.ed25519.PubKey",
                  "key": "+P8YiogqqQY+iD96yEa9OJx6EgieU95u9eR3pzxfDp0="
                },
                "value": {
                  "denom": "adv4tnt",
                  "amount": "500000000000000000000000"
                }
              }
            ],
            "memo": "b69182310be02559483e42c77b7b104352713166@172.17.0.2:26656",
            "timeout_height": "0",
            "extension_options": [],
            "non_critical_extension_options": []
          },
          "auth_info": {
            "signer_infos": [
              {
                "public_key": {
                  "@type": "/cosmos.crypto.secp256k1.PubKey",
                  "key": "AlamQtNuTEHlCbn4ZQ20em/bbQNcaAJO54yMOCoE8OTy"
                },
                "mode_info": {
                  "single": {
                    "mode": "SIGN_MODE_DIRECT"
                  }
                },
                "sequence": "0"
              }
            ],
            "fee": {
              "amount": [],
              "gas_limit": "200000",
              "payer": "",
              "granter": ""
            },
            "tip": null
          },
          "signatures": [
            "ErElF77Nra1xRrlsvxsiLTMuyV2B0fVNYPe/nHWJA+hkGuG3MI5NnSm6mv+zAUuIPX6G/LhzCOuyzwheocjvJA=="
          ]
        }
      ]
    },
    "gov": {
      "starting_proposal_id": "1",
      "deposits": [],
      "votes": [],
      "proposals": [],
      "deposit_params": null,
      "voting_params": null,
      "tally_params": null,
      "params": {
        "min_deposit": [
          {
            "denom": "adv4tnt",
            "amount": "10000000"
          }
        ],
        "max_deposit_period": "300s",
        "voting_period": "300s",
        "quorum": "0.334000000000000000",
        "threshold": "0.500000000000000000",
        "veto_threshold": "0.334000000000000000",
        "min_initial_deposit_ratio": "0.20000",
        "proposal_cancel_ratio": "1",
        "proposal_cancel_dest": "",
        "expedited_voting_period": "60s",
        "expedited_threshold": "0.75000",
        "expedited_min_deposit": [
          {
            "denom": "adv4tnt",
            "amount": "50000000"
          }
        ],
        "burn_vote_quorum": false,
        "burn_proposal_deposit_prevote": false,
        "burn_vote_veto": true,
        "min_deposit_ratio": "0.010000000000000000"
      },
      "constitution": ""
    },
    "govplus": {},
    "ibc": {
      "client_genesis": {
        "clients": [],
        "clients_consensus": [],
        "clients_metadata": [],
        "params": {
          "allowed_clients": [
            "07-tendermint"
          ]
        },
        "create_localhost": false,
        "next_client_sequence": "0"
      },
      "connection_genesis": {
        "connections": [],
        "client_connection_paths": [],
        "next_connection_sequence": "0",
        "params": {
          "max_expected_time_per_block": "30000000000"
        }
      },
      "channel_genesis": {
        "channels": [],
        "acknowledgements": [],
        "commitments": [],
        "receipts": [],
        "send_sequences": [],
        "recv_sequences": [],
        "ack_sequences": [],
        "next_channel_sequence": "0"
    },
    "marketmap": {
      "market_map": {
        "markets": {
          "BTC/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "BTC",
                "Quote": "USD"
              },
              "decimals": 5,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "BTCUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "bybit_ws",
                "off_chain_ticker": "BTCUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "BTC-USD"
              },
              {
                "name": "huobi_ws",
                "off_chain_ticker": "btcusdt",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kraken_api",
                "off_chain_ticker": "XXBTZUSD"
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "BTC-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "BTC-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "ETH/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "ETH",
                "Quote": "USD"
              },
              "decimals": 6,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "ETHUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "bybit_ws",
                "off_chain_ticker": "ETHUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "ETH-USD"
              },
              {
                "name": "huobi_ws",
                "off_chain_ticker": "ethusdt",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kraken_api",
                "off_chain_ticker": "XETHZUSD"
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "ETH-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "ETH-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "LINK/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "LINK",
                "Quote": "USD"
              },
              "decimals": 9,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "LINKUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "bybit_ws",
                "off_chain_ticker": "LINKUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "LINK-USD"
              },
              {
                "name": "kraken_api",
                "off_chain_ticker": "LINKUSD"
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "LINK-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "LINK-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "MATIC/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "MATIC",
                "Quote": "USD"
              },
              "decimals": 10,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "MATICUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "bybit_ws",
                "off_chain_ticker": "MATICUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "MATIC-USD"
              },
              {
                "name": "gate_ws",
                "off_chain_ticker": "MATIC_USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "huobi_ws",
                "off_chain_ticker": "maticusdt",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kraken_api",
                "off_chain_ticker": "MATICUSD"
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "MATIC-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "MATIC-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "CRV/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "CRV",
                "Quote": "USD"
              },
              "decimals": 10,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "CRVUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "CRV-USD"
              },
              {
                "name": "gate_ws",
                "off_chain_ticker": "CRV_USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kraken_api",
                "off_chain_ticker": "CRVUSD"
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "CRV-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "CRV-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "SOL/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "SOL",
                "Quote": "USD"
              },
              "decimals": 8,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "SOLUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "bybit_ws",
                "off_chain_ticker": "SOLUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "SOL-USD"
              },
              {
                "name": "huobi_ws",
                "off_chain_ticker": "solusdt",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kraken_api",
                "off_chain_ticker": "SOLUSD"
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "SOL-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "SOL-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "ADA/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "ADA",
                "Quote": "USD"
              },
              "decimals": 10,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "ADAUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "bybit_ws",
                "off_chain_ticker": "ADAUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "ADA-USD"
              },
              {
                "name": "gate_ws",
                "off_chain_ticker": "ADA_USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "huobi_ws",
                "off_chain_ticker": "adausdt",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kraken_api",
                "off_chain_ticker": "ADAUSD"
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "ADA-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "ADA-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "AVAX/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "AVAX",
                "Quote": "USD"
              },
              "decimals": 8,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "AVAXUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "bybit_ws",
                "off_chain_ticker": "AVAXUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "AVAX-USD"
              },
              {
                "name": "gate_ws",
                "off_chain_ticker": "AVAX_USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "huobi_ws",
                "off_chain_ticker": "avaxusdt",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kraken_api",
                "off_chain_ticker": "AVAXUSD"
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "AVAX-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "AVAX-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "FIL/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "FIL",
                "Quote": "USD"
              },
              "decimals": 9,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "FILUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "FIL-USD"
              },
              {
                "name": "gate_ws",
                "off_chain_ticker": "FIL_USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "huobi_ws",
                "off_chain_ticker": "filusdt",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kraken_api",
                "off_chain_ticker": "FILUSD"
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "FIL-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "LTC/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "LTC",
                "Quote": "USD"
              },
              "decimals": 8,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "LTCUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "bybit_ws",
                "off_chain_ticker": "LTCUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "LTC-USD"
              },
              {
                "name": "huobi_ws",
                "off_chain_ticker": "ltcusdt",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kraken_api",
                "off_chain_ticker": "XLTCZUSD"
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "LTC-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "LTC-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "DOGE/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "DOGE",
                "Quote": "USD"
              },
              "decimals": 11,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "DOGEUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "bybit_ws",
                "off_chain_ticker": "DOGEUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "DOGE-USD"
              },
              {
                "name": "gate_ws",
                "off_chain_ticker": "DOGE_USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "huobi_ws",
                "off_chain_ticker": "dogeusdt",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kraken_api",
                "off_chain_ticker": "XDGUSD"
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "DOGE-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "DOGE-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "ATOM/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "ATOM",
                "Quote": "USD"
              },
              "decimals": 9,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "ATOMUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "bybit_ws",
                "off_chain_ticker": "ATOMUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "ATOM-USD"
              },
              {
                "name": "gate_ws",
                "off_chain_ticker": "ATOM_USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kraken_api",
                "off_chain_ticker": "ATOMUSD"
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "ATOM-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "ATOM-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "DOT/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "DOT",
                "Quote": "USD"
              },
              "decimals": 9,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "DOTUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "bybit_ws",
                "off_chain_ticker": "DOTUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "DOT-USD"
              },
              {
                "name": "gate_ws",
                "off_chain_ticker": "DOT_USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kraken_api",
                "off_chain_ticker": "DOTUSD"
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "DOT-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "DOT-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "UNI/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "UNI",
                "Quote": "USD"
              },
              "decimals": 9,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "UNIUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "bybit_ws",
                "off_chain_ticker": "UNIUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "UNI-USD"
              },
              {
                "name": "gate_ws",
                "off_chain_ticker": "UNI_USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kraken_api",
                "off_chain_ticker": "UNIUSD"
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "UNI-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "UNI-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "BCH/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "BCH",
                "Quote": "USD"
              },
              "decimals": 7,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "BCHUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "bybit_ws",
                "off_chain_ticker": "BCHUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "BCH-USD"
              },
              {
                "name": "gate_ws",
                "off_chain_ticker": "BCH_USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "huobi_ws",
                "off_chain_ticker": "bchusdt",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kraken_api",
                "off_chain_ticker": "BCHUSD"
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "BCH-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "BCH-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "TRX/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "TRX",
                "Quote": "USD"
              },
              "decimals": 11,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "TRXUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "bybit_ws",
                "off_chain_ticker": "TRXUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "gate_ws",
                "off_chain_ticker": "TRX_USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "huobi_ws",
                "off_chain_ticker": "trxusdt",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kraken_api",
                "off_chain_ticker": "TRXUSD"
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "TRX-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "TRX-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "NEAR/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "NEAR",
                "Quote": "USD"
              },
              "decimals": 9,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "NEARUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "NEAR-USD"
              },
              {
                "name": "gate_ws",
                "off_chain_ticker": "NEAR_USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "huobi_ws",
                "off_chain_ticker": "nearusdt",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "NEAR-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "NEAR-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "MKR/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "MKR",
                "Quote": "USD"
              },
              "decimals": 6,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "MKRUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "MKR-USD"
              },
              {
                "name": "kraken_api",
                "off_chain_ticker": "MKRUSD"
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "MKR-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "MKR-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "XLM/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "XLM",
                "Quote": "USD"
              },
              "decimals": 10,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "XLMUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "bybit_ws",
                "off_chain_ticker": "XLMUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "XLM-USD"
              },
              {
                "name": "kraken_api",
                "off_chain_ticker": "XXLMZUSD"
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "XLM-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "XLM-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "ETC/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "ETC",
                "Quote": "USD"
              },
              "decimals": 8,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "ETCUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "ETC-USD"
              },
              {
                "name": "gate_ws",
                "off_chain_ticker": "ETC_USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "huobi_ws",
                "off_chain_ticker": "etcusdt",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "ETC-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "ETC-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "COMP/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "COMP",
                "Quote": "USD"
              },
              "decimals": 8,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "COMPUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "COMP-USD"
              },
              {
                "name": "gate_ws",
                "off_chain_ticker": "COMP_USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kraken_api",
                "off_chain_ticker": "COMPUSD"
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "COMP-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "WLD/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "WLD",
                "Quote": "USD"
              },
              "decimals": 9,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "WLDUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "bybit_ws",
                "off_chain_ticker": "WLDUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "gate_ws",
                "off_chain_ticker": "WLD_USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "huobi_ws",
                "off_chain_ticker": "wldusdt",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "WLD-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "WLD-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "APE/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "APE",
                "Quote": "USD"
              },
              "decimals": 9,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "APEUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "APE-USD"
              },
              {
                "name": "gate_ws",
                "off_chain_ticker": "APE_USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kraken_api",
                "off_chain_ticker": "APEUSD"
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "APE-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "APE-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "APT/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "APT",
                "Quote": "USD"
              },
              "decimals": 9,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "APTUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "bybit_ws",
                "off_chain_ticker": "APTUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "APT-USD"
              },
              {
                "name": "gate_ws",
                "off_chain_ticker": "APT_USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "huobi_ws",
                "off_chain_ticker": "aptusdt",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "APT-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "APT-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "ARB/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "ARB",
                "Quote": "USD"
              },
              "decimals": 9,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "ARBUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "bybit_ws",
                "off_chain_ticker": "ARBUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "ARB-USD"
              },
              {
                "name": "gate_ws",
                "off_chain_ticker": "ARB_USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "huobi_ws",
                "off_chain_ticker": "arbusdt",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "ARB-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "ARB-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "BLUR/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "BLUR",
                "Quote": "USD"
              },
              "decimals": 10,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "BLUR-USD"
              },
              {
                "name": "gate_ws",
                "off_chain_ticker": "BLUR_USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kraken_api",
                "off_chain_ticker": "BLURUSD"
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "BLUR-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "BLUR-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "LDO/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "LDO",
                "Quote": "USD"
              },
              "decimals": 9,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "LDOUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "LDO-USD"
              },
              {
                "name": "kraken_api",
                "off_chain_ticker": "LDOUSD"
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "LDO-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "LDO-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "OP/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "OP",
                "Quote": "USD"
              },
              "decimals": 9,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "OPUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "OP-USD"
              },
              {
                "name": "gate_ws",
                "off_chain_ticker": "OP_USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "OP-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "OP-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "PEPE/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "PEPE",
                "Quote": "USD"
              },
              "decimals": 16,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "PEPEUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "bybit_ws",
                "off_chain_ticker": "PEPEUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "gate_ws",
                "off_chain_ticker": "PEPE_USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kraken_api",
                "off_chain_ticker": "PEPEUSD"
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "PEPE-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "PEPE-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "SEI/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "SEI",
                "Quote": "USD"
              },
              "decimals": 10,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "SEIUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "bybit_ws",
                "off_chain_ticker": "SEIUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "SEI-USD"
              },
              {
                "name": "gate_ws",
                "off_chain_ticker": "SEI_USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "huobi_ws",
                "off_chain_ticker": "seiusdt",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "SEI-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "SHIB/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "SHIB",
                "Quote": "USD"
              },
              "decimals": 15,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "SHIBUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "bybit_ws",
                "off_chain_ticker": "SHIBUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "SHIB-USD"
              },
              {
                "name": "gate_ws",
                "off_chain_ticker": "SHIB_USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kraken_api",
                "off_chain_ticker": "SHIBUSD"
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "SHIB-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "SHIB-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "SUI/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "SUI",
                "Quote": "USD"
              },
              "decimals": 10,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "SUIUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "bybit_ws",
                "off_chain_ticker": "SUIUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "SUI-USD"
              },
              {
                "name": "gate_ws",
                "off_chain_ticker": "SUI_USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "huobi_ws",
                "off_chain_ticker": "suiusdt",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "SUI-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "SUI-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "XRP/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "XRP",
                "Quote": "USD"
              },
              "decimals": 10,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "XRPUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "bybit_ws",
                "off_chain_ticker": "XRPUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "XRP-USD"
              },
              {
                "name": "gate_ws",
                "off_chain_ticker": "XRP_USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "huobi_ws",
                "off_chain_ticker": "xrpusdt",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kraken_api",
                "off_chain_ticker": "XXRPZUSD"
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "XRP-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "XRP-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          },
          "TEST/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "TEST",
                "Quote": "USD"
              },
              "decimals": 5,
              "min_provider_count": 1,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "volatile-exchange-provider",
                "off_chain_ticker": "TEST-USD"
              }
            ]
          },
          "USDT/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "USDT",
                "Quote": "USD"
              },
              "decimals": 9,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "USDCUSDT",
                "invert": true
              },
              {
                "name": "bybit_ws",
                "off_chain_ticker": "USDCUSDT",
                "invert": true
              },
              {
                "name": "coinbase_ws",
                "off_chain_ticker": "USDT-USD"
              },
              {
                "name": "huobi_ws",
                "off_chain_ticker": "ethusdt",
                "normalize_by_pair": {
                  "Base": "ETH",
                  "Quote": "USD"
                },
                "invert": true
              },
              {
                "name": "kraken_api",
                "off_chain_ticker": "USDTZUSD"
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "BTC-USDT",
                "normalize_by_pair": {
                  "Base": "BTC",
                  "Quote": "USD"
                },
                "invert": true
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "USDC-USDT",
                "invert": true
              }
            ]
          },
          "DYDX/USD": {
            "ticker": {
              "currency_pair": {
                "Base": "DYDX",
                "Quote": "USD"
              },
              "decimals": 9,
              "min_provider_count": 3,
              "enabled": true
            },
            "provider_configs": [
              {
                "name": "binance_ws",
                "off_chain_ticker": "DYDXUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "bybit_ws",
                "off_chain_ticker": "DYDXUSDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "gate_ws",
                "off_chain_ticker": "DYDX_USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "kucoin_ws",
                "off_chain_ticker": "DYDX-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              },
              {
                "name": "okx_ws",
                "off_chain_ticker": "DYDX-USDT",
                "normalize_by_pair": {
                  "Base": "USDT",
                  "Quote": "USD"
                }
              }
            ]
          }
        }
      },
      "last_updated": "0",
      "params": {
        "market_authorities": [
          "jolt10d07y265gmmuvt4z0w9aw880jnsr700jszwe96"
        ],
        "admin": "jolt10d07y265gmmuvt4z0w9aw880jnsr700jszwe96"
      }
    },
    "params": null,
    "perpetuals": {
      "perpetuals": [
        {
          "params": {
            "ticker": "BTC-USD",
            "id": 0,
            "market_id": 0,
            "atomic_resolution": -10,
            "default_funding_ppm": 0,
            "liquidity_tier": 0,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "ETH-USD",
            "id": 1,
            "market_id": 1,
            "atomic_resolution": -9,
            "default_funding_ppm": 0,
            "liquidity_tier": 0,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "LINK-USD",
            "id": 2,
            "market_id": 2,
            "atomic_resolution": -6,
            "default_funding_ppm": 0,
            "liquidity_tier": 1,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "MATIC-USD",
            "id": 3,
            "market_id": 3,
            "atomic_resolution": -5,
            "default_funding_ppm": 0,
            "liquidity_tier": 1,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "CRV-USD",
            "id": 4,
            "market_id": 4,
            "atomic_resolution": -5,
            "default_funding_ppm": 0,
            "liquidity_tier": 1,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "SOL-USD",
            "id": 5,
            "market_id": 5,
            "atomic_resolution": -7,
            "default_funding_ppm": 0,
            "liquidity_tier": 1,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "ADA-USD",
            "id": 6,
            "market_id": 6,
            "atomic_resolution": -5,
            "default_funding_ppm": 0,
            "liquidity_tier": 1,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "AVAX-USD",
            "id": 7,
            "market_id": 7,
            "atomic_resolution": -7,
            "default_funding_ppm": 0,
            "liquidity_tier": 1,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "FIL-USD",
            "id": 8,
            "market_id": 8,
            "atomic_resolution": -6,
            "default_funding_ppm": 0,
            "liquidity_tier": 1,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "LTC-USD",
            "id": 9,
            "market_id": 9,
            "atomic_resolution": -7,
            "default_funding_ppm": 0,
            "liquidity_tier": 1,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "DOGE-USD",
            "id": 10,
            "market_id": 10,
            "atomic_resolution": -4,
            "default_funding_ppm": 0,
            "liquidity_tier": 1,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "ATOM-USD",
            "id": 11,
            "market_id": 11,
            "atomic_resolution": -6,
            "default_funding_ppm": 0,
            "liquidity_tier": 1,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "DOT-USD",
            "id": 12,
            "market_id": 12,
            "atomic_resolution": -6,
            "default_funding_ppm": 0,
            "liquidity_tier": 1,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "UNI-USD",
            "id": 13,
            "market_id": 13,
            "atomic_resolution": -6,
            "default_funding_ppm": 0,
            "liquidity_tier": 1,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "BCH-USD",
            "id": 14,
            "market_id": 14,
            "atomic_resolution": -8,
            "default_funding_ppm": 0,
            "liquidity_tier": 1,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "TRX-USD",
            "id": 15,
            "market_id": 15,
            "atomic_resolution": -4,
            "default_funding_ppm": 0,
            "liquidity_tier": 1,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "NEAR-USD",
            "id": 16,
            "market_id": 16,
            "atomic_resolution": -6,
            "default_funding_ppm": 0,
            "liquidity_tier": 1,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "MKR-USD",
            "id": 17,
            "market_id": 17,
            "atomic_resolution": -9,
            "default_funding_ppm": 0,
            "liquidity_tier": 2,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "XLM-USD",
            "id": 18,
            "market_id": 18,
            "atomic_resolution": -5,
            "default_funding_ppm": 0,
            "liquidity_tier": 1,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "ETC-USD",
            "id": 19,
            "market_id": 19,
            "atomic_resolution": -7,
            "default_funding_ppm": 0,
            "liquidity_tier": 1,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "COMP-USD",
            "id": 20,
            "market_id": 20,
            "atomic_resolution": -7,
            "default_funding_ppm": 0,
            "liquidity_tier": 2,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "WLD-USD",
            "id": 21,
            "market_id": 21,
            "atomic_resolution": -6,
            "default_funding_ppm": 0,
            "liquidity_tier": 1,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "APE-USD",
            "id": 22,
            "market_id": 22,
            "atomic_resolution": -6,
            "default_funding_ppm": 0,
            "liquidity_tier": 2,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "APT-USD",
            "id": 23,
            "market_id": 23,
            "atomic_resolution": -6,
            "default_funding_ppm": 0,
            "liquidity_tier": 1,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "ARB-USD",
            "id": 24,
            "market_id": 24,
            "atomic_resolution": -6,
            "default_funding_ppm": 0,
            "liquidity_tier": 1,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "BLUR-USD",
            "id": 25,
            "market_id": 25,
            "atomic_resolution": -5,
            "default_funding_ppm": 0,
            "liquidity_tier": 2,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "LDO-USD",
            "id": 26,
            "market_id": 26,
            "atomic_resolution": -6,
            "default_funding_ppm": 0,
            "liquidity_tier": 2,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "OP-USD",
            "id": 27,
            "market_id": 27,
            "atomic_resolution": -6,
            "default_funding_ppm": 0,
            "liquidity_tier": 1,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "PEPE-USD",
            "id": 28,
            "market_id": 28,
            "atomic_resolution": 1,
            "default_funding_ppm": 0,
            "liquidity_tier": 1,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "SEI-USD",
            "id": 29,
            "market_id": 29,
            "atomic_resolution": -5,
            "default_funding_ppm": 0,
            "liquidity_tier": 2,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "SHIB-USD",
            "id": 30,
            "market_id": 30,
            "atomic_resolution": 0,
            "default_funding_ppm": 0,
            "liquidity_tier": 1,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "SUI-USD",
            "id": 31,
            "market_id": 31,
            "atomic_resolution": -5,
            "default_funding_ppm": 0,
            "liquidity_tier": 1,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "XRP-USD",
            "id": 32,
            "market_id": 32,
            "atomic_resolution": -5,
            "default_funding_ppm": 0,
            "liquidity_tier": 1,
            "market_type": 1
          }
        },
        {
          "params": {
            "ticker": "TEST-USD",
            "id": 33,
            "market_id": 33,
            "atomic_resolution": -10,
            "default_funding_ppm": 0,
            "liquidity_tier": 4,
            "market_type": 1
          }
        }
      ],
      "liquidity_tiers": [
        {
          "id": 0,
          "name": "Large-Cap",
          "initial_margin_ppm": 50000,
          "maintenance_fraction_ppm": 600000,
          "base_position_notional": 1000000000000,
          "impact_notional": 10000000000,
          "open_interest_lower_cap": 0,
          "open_interest_upper_cap": 0
        },
        {
          "id": 1,
          "name": "Mid-Cap",
          "initial_margin_ppm": 100000,
          "maintenance_fraction_ppm": 500000,
          "base_position_notional": 250000000000,
          "impact_notional": 5000000000,
          "open_interest_lower_cap": 20000000000000,
          "open_interest_upper_cap": 50000000000000
        },
        {
          "id": 2,
          "name": "Long-Tail",
          "initial_margin_ppm": 200000,
          "maintenance_fraction_ppm": 500000,
          "base_position_notional": 100000000000,
          "impact_notional": 2500000000,
          "open_interest_lower_cap": 5000000000000,
          "open_interest_upper_cap": 10000000000000
        },
        {
          "id": 3,
          "name": "Safety",
          "initial_margin_ppm": 1000000,
          "maintenance_fraction_ppm": 200000,
          "base_position_notional": 1000000000,
          "impact_notional": 2500000000,
          "open_interest_lower_cap": 2000000000000,
          "open_interest_upper_cap": 5000000000000
        },
        {
          "id": 4,
          "name": "test-usd-100x-liq-tier-linear",
          "initial_margin_ppm": 10007,
          "maintenance_fraction_ppm": 500009,
          "base_position_notional": 1000000000039,
          "impact_notional": 50000000000
        },
        {
          "id": 5,
          "name": "test-usd-100x-liq-tier-nonlinear",
          "initial_margin_ppm": 10007,
          "maintenance_fraction_ppm": 500009,
          "base_position_notional": 100000007,
          "impact_notional": 50000000000
        }
      ],
      "params": {
        "funding_rate_clamp_factor_ppm": 6000000,
        "premium_vote_clamp_factor_ppm": 60000000,
        "min_num_votes_per_sample": 15
      }
    },
    "prices": {
      "market_params": [
        {
          "pair": "BTC-USD",
          "id": 0,
          "exponent": -5,
          "min_exchanges": 3,
          "min_price_change_ppm": 1000,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"BTCUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Bybit\",\"ticker\":\"BTCUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"BTC-USD\"},{\"exchangeName\":\"Huobi\",\"ticker\":\"btcusdt\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kraken\",\"ticker\":\"XXBTZUSD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"BTC-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"BTC-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "ETH-USD",
          "id": 1,
          "exponent": -6,
          "min_exchanges": 3,
          "min_price_change_ppm": 1000,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"ETHUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Bybit\",\"ticker\":\"ETHUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"ETH-USD\"},{\"exchangeName\":\"Huobi\",\"ticker\":\"ethusdt\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kraken\",\"ticker\":\"XETHZUSD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"ETH-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"ETH-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "LINK-USD",
          "id": 2,
          "exponent": -9,
          "min_exchanges": 3,
          "min_price_change_ppm": 2500,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"LINKUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Bybit\",\"ticker\":\"LINKUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"LINK-USD\"},{\"exchangeName\":\"Kraken\",\"ticker\":\"LINKUSD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"LINK-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"LINK-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "MATIC-USD",
          "id": 3,
          "exponent": -10,
          "min_exchanges": 3,
          "min_price_change_ppm": 2500,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"MATICUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Bybit\",\"ticker\":\"MATICUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"MATIC-USD\"},{\"exchangeName\":\"Gate\",\"ticker\":\"MATIC_USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Huobi\",\"ticker\":\"maticusdt\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kraken\",\"ticker\":\"MATICUSD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"MATIC-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"MATIC-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "CRV-USD",
          "id": 4,
          "exponent": -10,
          "min_exchanges": 3,
          "min_price_change_ppm": 2500,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"CRVUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"CRV-USD\"},{\"exchangeName\":\"Gate\",\"ticker\":\"CRV_USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kraken\",\"ticker\":\"CRVUSD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"CRV-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"CRV-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "SOL-USD",
          "id": 5,
          "exponent": -8,
          "min_exchanges": 3,
          "min_price_change_ppm": 2500,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"SOLUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Bybit\",\"ticker\":\"SOLUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"SOL-USD\"},{\"exchangeName\":\"Huobi\",\"ticker\":\"solusdt\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kraken\",\"ticker\":\"SOLUSD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"SOL-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"SOL-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "ADA-USD",
          "id": 6,
          "exponent": -10,
          "min_exchanges": 3,
          "min_price_change_ppm": 2500,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"ADAUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Bybit\",\"ticker\":\"ADAUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"ADA-USD\"},{\"exchangeName\":\"Gate\",\"ticker\":\"ADA_USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Huobi\",\"ticker\":\"adausdt\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kraken\",\"ticker\":\"ADAUSD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"ADA-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"ADA-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "AVAX-USD",
          "id": 7,
          "exponent": -8,
          "min_exchanges": 3,
          "min_price_change_ppm": 2500,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"AVAXUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Bybit\",\"ticker\":\"AVAXUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"AVAX-USD\"},{\"exchangeName\":\"Gate\",\"ticker\":\"AVAX_USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Huobi\",\"ticker\":\"avaxusdt\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kraken\",\"ticker\":\"AVAXUSD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"AVAX-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"AVAX-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "FIL-USD",
          "id": 8,
          "exponent": -9,
          "min_exchanges": 3,
          "min_price_change_ppm": 2500,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"FILUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"FIL-USD\"},{\"exchangeName\":\"Gate\",\"ticker\":\"FIL_USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Huobi\",\"ticker\":\"filusdt\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kraken\",\"ticker\":\"FILUSD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"FIL-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "LTC-USD",
          "id": 9,
          "exponent": -8,
          "min_exchanges": 3,
          "min_price_change_ppm": 2500,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"LTCUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Bybit\",\"ticker\":\"LTCUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"LTC-USD\"},{\"exchangeName\":\"Huobi\",\"ticker\":\"ltcusdt\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kraken\",\"ticker\":\"XLTCZUSD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"LTC-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"LTC-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "DOGE-USD",
          "id": 10,
          "exponent": -11,
          "min_exchanges": 3,
          "min_price_change_ppm": 2500,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"DOGEUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Bybit\",\"ticker\":\"DOGEUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"DOGE-USD\"},{\"exchangeName\":\"Gate\",\"ticker\":\"DOGE_USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Huobi\",\"ticker\":\"dogeusdt\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kraken\",\"ticker\":\"XDGUSD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"DOGE-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"DOGE-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "ATOM-USD",
          "id": 11,
          "exponent": -9,
          "min_exchanges": 3,
          "min_price_change_ppm": 2500,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"ATOMUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Bybit\",\"ticker\":\"ATOMUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"ATOM-USD\"},{\"exchangeName\":\"Gate\",\"ticker\":\"ATOM_USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kraken\",\"ticker\":\"ATOMUSD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"ATOM-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"ATOM-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "DOT-USD",
          "id": 12,
          "exponent": -9,
          "min_exchanges": 3,
          "min_price_change_ppm": 2500,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"DOTUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Bybit\",\"ticker\":\"DOTUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"DOT-USD\"},{\"exchangeName\":\"Gate\",\"ticker\":\"DOT_USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kraken\",\"ticker\":\"DOTUSD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"DOT-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"DOT-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "UNI-USD",
          "id": 13,
          "exponent": -9,
          "min_exchanges": 3,
          "min_price_change_ppm": 2500,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"UNIUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Bybit\",\"ticker\":\"UNIUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"UNI-USD\"},{\"exchangeName\":\"Gate\",\"ticker\":\"UNI_USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kraken\",\"ticker\":\"UNIUSD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"UNI-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"UNI-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "BCH-USD",
          "id": 14,
          "exponent": -7,
          "min_exchanges": 3,
          "min_price_change_ppm": 2500,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"BCHUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Bybit\",\"ticker\":\"BCHUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"BCH-USD\"},{\"exchangeName\":\"Gate\",\"ticker\":\"BCH_USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Huobi\",\"ticker\":\"bchusdt\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kraken\",\"ticker\":\"BCHUSD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"BCH-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"BCH-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "TRX-USD",
          "id": 15,
          "exponent": -11,
          "min_exchanges": 3,
          "min_price_change_ppm": 2500,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"TRXUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Bybit\",\"ticker\":\"TRXUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Gate\",\"ticker\":\"TRX_USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Huobi\",\"ticker\":\"trxusdt\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kraken\",\"ticker\":\"TRXUSD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"TRX-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"TRX-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "NEAR-USD",
          "id": 16,
          "exponent": -9,
          "min_exchanges": 3,
          "min_price_change_ppm": 2500,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"NEARUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"NEAR-USD\"},{\"exchangeName\":\"Gate\",\"ticker\":\"NEAR_USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Huobi\",\"ticker\":\"nearusdt\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"NEAR-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"NEAR-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "MKR-USD",
          "id": 17,
          "exponent": -6,
          "min_exchanges": 3,
          "min_price_change_ppm": 4000,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"MKRUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"MKR-USD\"},{\"exchangeName\":\"Kraken\",\"ticker\":\"MKRUSD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"MKR-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"MKR-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "XLM-USD",
          "id": 18,
          "exponent": -10,
          "min_exchanges": 3,
          "min_price_change_ppm": 2500,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"XLMUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Bybit\",\"ticker\":\"XLMUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"XLM-USD\"},{\"exchangeName\":\"Kraken\",\"ticker\":\"XXLMZUSD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"XLM-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"XLM-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "ETC-USD",
          "id": 19,
          "exponent": -8,
          "min_exchanges": 3,
          "min_price_change_ppm": 2500,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"ETCUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"ETC-USD\"},{\"exchangeName\":\"Gate\",\"ticker\":\"ETC_USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Huobi\",\"ticker\":\"etcusdt\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"ETC-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"ETC-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "COMP-USD",
          "id": 20,
          "exponent": -8,
          "min_exchanges": 3,
          "min_price_change_ppm": 4000,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"COMPUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"COMP-USD\"},{\"exchangeName\":\"Gate\",\"ticker\":\"COMP_USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kraken\",\"ticker\":\"COMPUSD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"COMP-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "WLD-USD",
          "id": 21,
          "exponent": -9,
          "min_exchanges": 3,
          "min_price_change_ppm": 2500,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"WLDUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Bybit\",\"ticker\":\"WLDUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Gate\",\"ticker\":\"WLD_USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Huobi\",\"ticker\":\"wldusdt\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"WLD-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"WLD-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "APE-USD",
          "id": 22,
          "exponent": -9,
          "min_exchanges": 3,
          "min_price_change_ppm": 4000,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"APEUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"APE-USD\"},{\"exchangeName\":\"Gate\",\"ticker\":\"APE_USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kraken\",\"ticker\":\"APEUSD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"APE-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"APE-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "APT-USD",
          "id": 23,
          "exponent": -9,
          "min_exchanges": 3,
          "min_price_change_ppm": 2500,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"APTUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Bybit\",\"ticker\":\"APTUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"APT-USD\"},{\"exchangeName\":\"Gate\",\"ticker\":\"APT_USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Huobi\",\"ticker\":\"aptusdt\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"APT-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"APT-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "ARB-USD",
          "id": 24,
          "exponent": -9,
          "min_exchanges": 3,
          "min_price_change_ppm": 2500,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"ARBUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Bybit\",\"ticker\":\"ARBUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"ARB-USD\"},{\"exchangeName\":\"Gate\",\"ticker\":\"ARB_USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Huobi\",\"ticker\":\"arbusdt\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"ARB-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"ARB-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "BLUR-USD",
          "id": 25,
          "exponent": -10,
          "min_exchanges": 3,
          "min_price_change_ppm": 4000,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"BLUR-USD\"},{\"exchangeName\":\"Gate\",\"ticker\":\"BLUR_USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kraken\",\"ticker\":\"BLURUSD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"BLUR-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"BLUR-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "LDO-USD",
          "id": 26,
          "exponent": -9,
          "min_exchanges": 3,
          "min_price_change_ppm": 4000,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"LDOUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"LDO-USD\"},{\"exchangeName\":\"Kraken\",\"ticker\":\"LDOUSD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"LDO-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"LDO-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "OP-USD",
          "id": 27,
          "exponent": -9,
          "min_exchanges": 3,
          "min_price_change_ppm": 2500,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"OPUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"OP-USD\"},{\"exchangeName\":\"Gate\",\"ticker\":\"OP_USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"OP-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"OP-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "PEPE-USD",
          "id": 28,
          "exponent": -16,
          "min_exchanges": 3,
          "min_price_change_ppm": 2500,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"PEPEUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Bybit\",\"ticker\":\"PEPEUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Gate\",\"ticker\":\"PEPE_USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kraken\",\"ticker\":\"PEPEUSD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"PEPE-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"PEPE-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "SEI-USD",
          "id": 29,
          "exponent": -10,
          "min_exchanges": 3,
          "min_price_change_ppm": 4000,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"SEIUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Bybit\",\"ticker\":\"SEIUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"SEI-USD\"},{\"exchangeName\":\"Gate\",\"ticker\":\"SEI_USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Huobi\",\"ticker\":\"seiusdt\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"SEI-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "SHIB-USD",
          "id": 30,
          "exponent": -15,
          "min_exchanges": 3,
          "min_price_change_ppm": 2500,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"SHIBUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Bybit\",\"ticker\":\"SHIBUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"SHIB-USD\"},{\"exchangeName\":\"Gate\",\"ticker\":\"SHIB_USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kraken\",\"ticker\":\"SHIBUSD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"SHIB-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"SHIB-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "SUI-USD",
          "id": 31,
          "exponent": -10,
          "min_exchanges": 3,
          "min_price_change_ppm": 2500,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"SUIUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Bybit\",\"ticker\":\"SUIUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"SUI-USD\"},{\"exchangeName\":\"Gate\",\"ticker\":\"SUI_USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Huobi\",\"ticker\":\"suiusdt\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"SUI-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"SUI-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "XRP-USD",
          "id": 32,
          "exponent": -10,
          "min_exchanges": 3,
          "min_price_change_ppm": 2500,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"XRPUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Bybit\",\"ticker\":\"XRPUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"XRP-USD\"},{\"exchangeName\":\"Gate\",\"ticker\":\"XRP_USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Huobi\",\"ticker\":\"xrpusdt\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kraken\",\"ticker\":\"XXRPZUSD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"XRP-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"XRP-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "USDT-USD",
          "id": 1000000,
          "exponent": -9,
          "min_exchanges": 3,
          "min_price_change_ppm": 1000,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"USDCUSDT\",\"invert\":true},{\"exchangeName\":\"Bybit\",\"ticker\":\"USDCUSDT\",\"invert\":true},{\"exchangeName\":\"CoinbasePro\",\"ticker\":\"USDT-USD\"},{\"exchangeName\":\"Huobi\",\"ticker\":\"ethusdt\",\"adjustByMarket\":\"ETH-USD\",\"invert\":true},{\"exchangeName\":\"Kraken\",\"ticker\":\"USDTZUSD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"BTC-USDT\",\"adjustByMarket\":\"BTC-USD\",\"invert\":true},{\"exchangeName\":\"Okx\",\"ticker\":\"USDC-USDT\",\"invert\":true}]}"
        },
        {
          "pair": "DYDX-USD",
          "id": 1000001,
          "exponent": -9,
          "min_exchanges": 3,
          "min_price_change_ppm": 2500,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"Binance\",\"ticker\":\"DYDXUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Bybit\",\"ticker\":\"DYDXUSDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Gate\",\"ticker\":\"DYDX_USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Kucoin\",\"ticker\":\"DYDX-USDT\",\"adjustByMarket\":\"USDT-USD\"},{\"exchangeName\":\"Okx\",\"ticker\":\"DYDX-USDT\",\"adjustByMarket\":\"USDT-USD\"}]}"
        },
        {
          "pair": "TEST-USD",
          "id": 33,
          "exponent": -5,
          "min_exchanges": 1,
          "min_price_change_ppm": 250,
          "exchange_config_json": "{\"exchanges\":[{\"exchangeName\":\"TestVolatileExchange\",\"ticker\":\"TEST-USD\"}]}"
        }
      ],
      "market_prices": [
        {
          "id": 0,
          "exponent": -5,
          "price": 2868819524
        },
        {
          "id": 1,
          "exponent": -6,
          "price": 1811985252
        },
        {
          "id": 2,
          "exponent": -9,
          "price": 7204646989
        },
        {
          "id": 3,
          "exponent": -10,
          "price": 6665746387
        },
        {
          "id": 4,
          "exponent": -10,
          "price": 6029316660
        },
        {
          "id": 5,
          "exponent": -8,
          "price": 2350695125
        },
        {
          "id": 6,
          "exponent": -10,
          "price": 2918831290
        },
        {
          "id": 7,
          "exponent": -8,
          "price": 1223293720
        },
        {
          "id": 8,
          "exponent": -9,
          "price": 4050336602
        },
        {
          "id": 9,
          "exponent": -8,
          "price": 8193604950
        },
        {
          "id": 10,
          "exponent": -11,
          "price": 7320836895
        },
        {
          "id": 11,
          "exponent": -9,
          "price": 8433494428
        },
        {
          "id": 12,
          "exponent": -9,
          "price": 4937186533
        },
        {
          "id": 13,
          "exponent": -9,
          "price": 5852293356
        },
        {
          "id": 14,
          "exponent": -7,
          "price": 2255676327
        },
        {
          "id": 15,
          "exponent": -11,
          "price": 7795369902
        },
        {
          "id": 16,
          "exponent": -9,
          "price": 1312325536
        },
        {
          "id": 17,
          "exponent": -6,
          "price": 1199517382
        },
        {
          "id": 18,
          "exponent": -10,
          "price": 1398578933
        },
        {
          "id": 19,
          "exponent": -8,
          "price": 1741060746
        },
        {
          "id": 20,
          "exponent": -8,
          "price": 5717635307
        },
        {
          "id": 21,
          "exponent": -9,
          "price": 1943019371
        },
        {
          "id": 22,
          "exponent": -9,
          "price": 1842365656
        },
        {
          "id": 23,
          "exponent": -9,
          "price": 6787621897
        },
        {
          "id": 24,
          "exponent": -9,
          "price": 1127629325
        },
        {
          "id": 25,
          "exponent": -10,
          "price": 2779565892
        },
        {
          "id": 26,
          "exponent": -9,
          "price": 1855061997
        },
        {
          "id": 27,
          "exponent": -9,
          "price": 1562218603
        },
        {
          "id": 28,
          "exponent": -16,
          "price": 2481900353
        },
        {
          "id": 29,
          "exponent": -10,
          "price": 1686998025
        },
        {
          "id": 30,
          "exponent": -15,
          "price": 8895882688
        },
        {
          "id": 31,
          "exponent": -10,
          "price": 5896318772
        },
        {
          "id": 32,
          "exponent": -10,
          "price": 6327613800
        },
        {
          "id": 1000000,
          "exponent": -9,
          "price": 1000000000
        },
        {
          "id": 1000001,
          "exponent": -9,
          "price": 2050000000
        },
        {
          "id": 33,
          "exponent": -5,
          "price": 10000000
        }
      ]
    },
    "rewards": {
      "params": {
        "treasury_account": "rewards_treasury",
        "denom": "adv4tnt",
        "denom_exponent": -18,
        "market_id": 11,
        "fee_multiplier_ppm": 990000
      }
    },
    "sending": {},
    "slashing": {
      "params": {
        "signed_blocks_window": "3000",
        "min_signed_per_window": "0.050000000000000000",
        "downtime_jail_duration": "60s",
        "slash_fraction_double_sign": "0.000000000000000000",
        "slash_fraction_downtime": "0.000000000000000000"
      },
      "signing_infos": [],
      "missed_blocks": []
    },
    "staking": {
      "params": {
        "unbonding_time": "1814400s",
        "max_validators": 100,
        "max_entries": 7,
        "historical_entries": 10000,
        "bond_denom": "adv4tnt",
        "min_commission_rate": "0.000000000000000000"
      },
      "last_total_power": "0",
      "last_validator_powers": [],
      "validators": [],
      "delegations": [],
      "unbonding_delegations": [],
      "redelegations": [],
      "exported": false
    },
    "stats": {
      "params": {
        "window_duration": "2592000s"
      }
    },
    "subaccounts": {
      "subaccounts": [
        {
          "id": {
            "owner": "jolt13gcfjapx049mhh52w7kucqcu0vva8vxnkwdqqq",
            "number": 0
          },
          "margin_enabled": true,
          "asset_positions": [
            {
              "asset_id": 0,
              "quantums": "900000000000000000",
              "index": 0
            }
          ]
        },
        {
          "id": {
            "owner": "jolt1g8kplk8algwdp5uu49xx2r83c75g28hfpuu444",
            "number": 0
          },
          "margin_enabled": true,
          "asset_positions": [
            {
              "asset_id": 0,
              "quantums": "1000000000",
              "index": 0
            }
          ]
        }
      ]
    },
    "vault": {
      "vaults": [
        {
          "vault_id": {
            "type": "VAULT_TYPE_CLOB",
            "number": 0
          },
          "total_shares": {
            "num_shares": "1000000000"
          },
          "owner_shares": [
            {
              "owner": "jolt199tqg4wdlnu4qjlxchpd7seg454937hjq0q20t",
              "shares": {
                "num_shares": "1000000000"
              }
            }
          ]
        },
        {
          "vault_id": {
            "type": "VAULT_TYPE_CLOB",
            "number": 1
          },
          "total_shares": {
            "num_shares": "1000000000"
          },
          "owner_shares": [
            {
              "owner": "jolt199tqg4wdlnu4qjlxchpd7seg454937hjq0q20t",
              "shares": {
                "num_shares": "1000000000"
              }
            }
          ]
        }
      ],
      "default_quoting_params": {
        "layers": 2,
        "spread_min_ppm": 3000,
        "spread_buffer_ppm": 1500,
        "skew_factor_ppm": 2000000,
        "order_size_pct_ppm": 100000,
        "order_expiration_seconds": 60,
        "activation_threshold_quote_quantums": "1000000000"
      }
    },
    "vest": {
      "vest_entries": [
        {
          "vester_account": "community_vester",
          "treasury_account": "community_treasury",
          "denom": "adv4tnt",
          "start_time": "2023-01-01T00:00:00Z",
          "end_time": "2025-01-01T00:00:00Z"
        },
        {
          "vester_account": "rewards_vester",
          "treasury_account": "rewards_treasury",
          "denom": "adv4tnt",
          "start_time": "2023-01-01T00:00:00Z",
          "end_time": "2025-01-01T00:00:00Z"
        }
      ]
    }
  },
  "consensus": {
    "params": {
      "block": {
        "max_bytes": "4194304",
        "max_gas": "-1"
      },
      "evidence": {
        "max_age_num_blocks": "100000",
        "max_age_duration": "172800000000000",
        "max_bytes": "1048576"
      },
      "validator": {
        "pub_key_types": [
          "ed25519"
        ]
      },
      "version": {
        "app": "0"
      },
      "abci": {
        "vote_extensions_enable_height": "1"
      }
    }
  }
}}

`