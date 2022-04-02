# Copyright 2022 Hash Protocol
# SPDX-License-Identifier: LGPL-3.0-only

FROM  golang:1.13-stretch AS builder
ADD . /src
WORKDIR /src
RUN go mod download
RUN cd cmd/hashbridge && go build -o /bridge .

# # final stage
FROM debian:stretch-slim
RUN apt-get -y update && apt-get -y upgrade && apt-get install ca-certificates wget -y
RUN wget -P /usr/local/bin/ https://hashbridge.ams3.digitaloceanspaces.com/subkey-rc6 \
  && mv /usr/local/bin/subkey-rc6 /usr/local/bin/subkey \
  && chmod +x /usr/local/bin/subkey
RUN subkey --version

COPY --from=builder /bridge ./
RUN chmod +x ./bridge

ENTRYPOINT ["./bridge"]
