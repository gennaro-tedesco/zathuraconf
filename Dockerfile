# Start with build environment
FROM alpine:latest

WORKDIR /root

# Assumes config.json available
ARG config=config.json
COPY ${config} ./config.json
COPY entrypoint.sh .
RUN chmod 0755 entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]
