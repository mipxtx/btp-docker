#!/bin/bash -x

exec /btpd/build/bin/btpd -C /btpd/btpd.conf > /var/log/btp.output.log 2>&1