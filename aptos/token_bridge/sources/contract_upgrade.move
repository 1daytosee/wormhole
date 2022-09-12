/// This module implements upgradeability for the token bridge contract.
///
/// Contract upgrades are authorised by governance, which means that performing
/// an upgrade requires a governance VAA signed by a supermajority of the
/// wormhole guardians.
///
/// Upgrades are performed in a commit-reveal scheme, where submitting the VAA
/// authorises a particular contract hash. Then in a subsequent transaction, the
/// bytecode is uploaded, and if the hash of the bytecode matches the committed
/// hash, then the upgrade proceeds.
///
/// This two-phase process has the advantage that even if the bytecode can't be
/// upgraded to for whatever reason, the governance VAA won't be possible to
/// replay in the future, since the commit transaction replay protects it.
module token_bridge::contract_upgrade {
    use std::aptos_hash;
    use std::vector;
    use aptos_framework::code;
    use wormhole::deserialize;
    use wormhole::cursor;
    use wormhole::vaa;
    use wormhole::state as core;

    use token_bridge::vaa as token_bridge_vaa;
    use token_bridge::state;

    /// "TokenBridge" (left padded)
    const TOKEN_BRIDGE: vector<u8> = x"000000000000000000000000000000000000000000546f6b656e427269646765";

    const E_UPGRADE_UNAUTHORIZED: u64 = 0;
    const E_UNEXPECTED_HASH: u64 = 1;
    const E_INVALID_MODULE: u64 = 2;
    const E_INVALID_ACTION: u64 = 3;
    const E_INVALID_TARGET: u64 = 4;

    /// The `UpgradeAuthorized` type in the global storage represents the fact
    /// there is an ongoing approved upgrade.
    /// When the upgrade is finalised in `upgrade`, this object is deleted.
    struct UpgradeAuthorized has key {
        hash: vector<u8>
    }

    struct Hash {
        hash: vector<u8>
    }

    fun parse_payload(payload: vector<u8>): Hash {
        let cur = cursor::init(payload);
        let target_module = deserialize::deserialize_vector(&mut cur, 32);

        assert!(target_module == TOKEN_BRIDGE, E_INVALID_MODULE);

        let action = deserialize::deserialize_u8(&mut cur);
        assert!(action == 0x02, E_INVALID_ACTION);

        let chain = deserialize::deserialize_u16(&mut cur);
        assert!(chain == core::get_chain_id(), E_INVALID_TARGET);

        let hash = deserialize::deserialize_vector(&mut cur, 32);

        cursor::destroy_empty(cur);

        Hash { hash }
    }

// -----------------------------------------------------------------------------
// Commit

    public entry fun submit_vaa(vaa: vector<u8>) acquires UpgradeAuthorized {
        let vaa = vaa::parse_and_verify(vaa);
        vaa::assert_governance(&vaa);
        token_bridge_vaa::replay_protect(&vaa);

        authorize_upgrade(parse_payload(vaa::destroy(vaa)));
    }

    fun authorize_upgrade(hash: Hash) acquires UpgradeAuthorized {
        let Hash { hash } = hash;
        let token_bridge = state::token_bridge_signer();
        if (exists<UpgradeAuthorized>(@token_bridge)) {
            // TODO(csongor): here we're dropping the upgrade hash, in case an
            // upgrade fails for some reason. Should we emit a log or something?
            let UpgradeAuthorized { hash: _ } = move_from<UpgradeAuthorized>(@token_bridge);
        };
        move_to(&token_bridge, UpgradeAuthorized { hash });
    }

    #[test_only]
    public fun authorized_hash(): vector<u8> acquires UpgradeAuthorized {
        let u = borrow_global<UpgradeAuthorized>(@token_bridge);
        u.hash
    }

// -----------------------------------------------------------------------------
// Reveal

    public entry fun upgrade(
        metadata_serialized: vector<u8>,
        code: vector<vector<u8>>
    ) acquires UpgradeAuthorized {
        assert!(exists<UpgradeAuthorized>(@token_bridge), E_UPGRADE_UNAUTHORIZED);
        let UpgradeAuthorized { hash } = move_from<UpgradeAuthorized>(@token_bridge);

        let c = copy code;
        vector::reverse(&mut c);
        let a = vector::empty<u8>();
        while (!vector::is_empty(&c)) vector::append(&mut a, vector::pop_back(&mut c));
        assert!(aptos_hash::keccak256(a) == hash, E_UNEXPECTED_HASH);

        let token_bridge = state::token_bridge_signer();
        code::publish_package_txn(&token_bridge, metadata_serialized, code);
    }
}

#[test_only]
module token_bridge::contract_upgrade_test {
    use wormhole::wormhole;

    use token_bridge::contract_upgrade;
    use token_bridge::token_bridge;

    /// A token bridge upgrade VAA that upgrades to 0x10263f154c466b139fda0bf2caa08fd9819d8ded3810446274a99399f886fc76
    const UPGRADE_VAA: vector<u8> = x"01000000000100b5ebfcccb84d740684429622f2fbc16638fb01222e4a580a6d2049227f37a31a7162d32770f72398fe10d160a968c94256eae9225a3da9c69ab7a41d7b307ede010000000100000001000100000000000000000000000000000000000000000000000000000000000000040000000001f96c9900000000000000000000000000000000000000000000546f6b656e42726964676502001610263f154c466b139fda0bf2caa08fd9819d8ded3810446274a99399f886fc76";

    /// A token bridge upgrade VAA that targets ethereum
    const ETH_UPGRADE: vector<u8> = x"0100000000010090014add41120b33eb4a03c5dce613815071d18b69a185bf322f327cc79cc52d7d133a59515d13ccfb030f9cc26a86b2bcd4dbe34d8ca6c4cc83299efb3e9b430100000001000000010001000000000000000000000000000000000000000000000000000000000000000400000000030a9ea600000000000000000000000000000000000000000000546f6b656e42726964676502000210263f154c466b139fda0bf2caa08fd9819d8ded3810446274a99399f886fc76";

    fun setup(deployer: &signer) {
        let aptos_framework = std::account::create_account_for_test(@aptos_framework);
        std::timestamp::set_time_has_started_for_testing(&aptos_framework);
        wormhole::init_test(
            22,
            1,
            x"0000000000000000000000000000000000000000000000000000000000000004",
            x"beFA429d57cD18b7F8A4d91A2da9AB4AF05d0FBe",
            0
        );
        token_bridge::init_test(deployer);
    }

    #[test(deployer = @deployer)]
    public fun test_contract_upgrade_authorize(deployer: &signer) {
        setup(deployer);

        contract_upgrade::submit_vaa(UPGRADE_VAA);
        let expected_hash = x"10263f154c466b139fda0bf2caa08fd9819d8ded3810446274a99399f886fc76";

        assert!(contract_upgrade::authorized_hash() == expected_hash, 0);
    }

    #[test(deployer = @deployer)]
    #[expected_failure(abort_code = 0x6407)]
    public fun test_contract_upgrade_double(deployer: &signer) {
        setup(deployer);

        // make sure we can't replay a VAA
        contract_upgrade::submit_vaa(UPGRADE_VAA);
        contract_upgrade::submit_vaa(UPGRADE_VAA);
    }

    #[test(deployer = @deployer)]
    #[expected_failure(abort_code = 4)]
    public fun test_contract_upgrade_wrong_chain(deployer: &signer) {
        setup(deployer);

        contract_upgrade::submit_vaa(ETH_UPGRADE);
    }

}
