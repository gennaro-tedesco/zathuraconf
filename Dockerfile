# Start with build environment
FROM alpine:latest

WORKDIR /root

# Assumes config.json available
COPY config.json .
COPY entrypoint.sh .
RUN chmod 0755 entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]
