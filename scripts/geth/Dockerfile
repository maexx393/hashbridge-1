# Copyright 2022 Hash Protocol
# SPDX-License-Identifier: LGPL-3.0-only

FROM ethereum/client-go:stable

WORKDIR /root

COPY ./genesis.json .
COPY password.txt .
COPY ./keystore /root/keystore
COPY entrypoint.sh .

RUN chmod +x entrypoint.sh

ENTRYPOINT ["/root/entrypoint.sh"]