#!/bin/bash
pushd /node

export CHTD_HOME="${PWD?}"


# This fails immediately, but creates the node keys
chtd init "${CHTD_MONIKER:-chronicNode}"  --home=.cht

set -xe

# export bech addresses on http.
#
# note: should be unnecessary. rpc/status has:
#
# - node-id in `.node_info.id`
# - validator address in `.validator_info.address`,
#   but it is in hex and `chtd keys parse` is broken (again).

if test -n "$ENABLE_ID_SERVER" ; then
  mkdir web
  chtd tendermint show-node-id   > web/node-id.txt  --home=.cht
  chtd tendermint show-validator > web/validator-pubkey.txt  --home=.cht
  pushd web
  # Run a web server so that the file can be retrieved
  python3 -m http.server 8080 &
  popd
fi

curl -s "${GENESIS_URL?}" > .cht/config/genesis.json

cat config.toml | python3 -u ./patch_config_toml.py > .cht/config/config.toml

# Copy over all the other filesthat the node needs
cp -v app.toml .cht/config/

# Run the node for real now
exec chtd start --home=.cht