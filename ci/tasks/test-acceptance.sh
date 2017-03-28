#!/usr/bin/env bash

set -exu

ROOT_DIR=$PWD

start-bosh
source /tmp/local-bosh/director/env

bosh int /tmp/local-bosh/director/creds.yml --path /jumpbox_ssh/private_key > /tmp/jumpbox_ssh_key.pem
chmod 400 /tmp/jumpbox_ssh_key.pem

export BOSH_SSH_PRIVATE_KEY_PATH="/tmp/jumpbox_ssh_key.pem"
export BOSH_DIRECTOR_IP="10.245.0.3"
export BOSH_BINARY_PATH=$(which bosh)
export BOSH_DEPLOYMENT="bosh-dns"

bosh upload-stemcell https://bosh.io/d/stemcells/bosh-warden-boshlite-ubuntu-trusty-go_agent
bosh -n deploy -v dns_release_path=$ROOT_DIR/dns-release $ROOT_DIR/dns-release/ci/assets/manifest.yml

export GOPATH=${ROOT_DIR}/go
export PATH="${GOPATH}/bin":$PATH

mkdir -p go/src/github.com/cloudfoundry
mkdir -p go/src/github.com/onsi
ln -s $PWD/dns-release $PWD/go/src/github.com/cloudfoundry/dns-release
ln -s $PWD/dns-release/src/vendor/github.com/onsi/ginkgo $PWD/go/src/github.com/onsi/ginkgo

go install github.com/onsi/ginkgo/ginkgo

pushd $GOPATH/src/github.com/cloudfoundry/dns-release/src/acceptance_tests
    ginkgo -r -randomizeAllSpecs -randomizeSuites -race .
popd