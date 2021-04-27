#!/bin/bash -x
cd ./btp-front && docker build . --tag=mamba:btp-front && cd .. && \
cd ./btp-php && docker build . --tag=mamba:btp-php && cd .. && \
cd ./btpd && docker build . --tag=mamba:btpd && cd .. && \
echo "you should run docker compose up to start the system"

