#!/usr/bin/env sh

set -x

cd /root

./habor-sync

/root/image-syncer --proc=10 --auth=/root/auth.json --images=/root/image.json