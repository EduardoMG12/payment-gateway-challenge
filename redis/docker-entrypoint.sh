#!/bin/sh
set -e

: "${WRITER_REDIS_USER:=writer}"
: "${WRITER_REDIS_PASSWORD:=writerpass}"
: "${READER_REDIS_USER:=reader}"
: "${READER_REDIS_PASSWORD:=readerpass}"

mkdir -p /data

cat >/data/users.acl <<EOF
user default off
user ${WRITER_REDIS_USER} on >${WRITER_REDIS_PASSWORD} ~* &* +@all
user ${READER_REDIS_USER} on >${READER_REDIS_PASSWORD} ~* &* +@read +@connection
EOF

exec redis-server /usr/local/etc/redis/redis.conf
