#!/bin/bash

set -e

TERRABASE_DOMAIN="terrabase.local"

SCRIPT_DIR="$(dirname "$(readlink -f "$0")")"
PARENT_DIR="$(dirname "$SCRIPT_DIR")"
CERTS_DIR="$PARENT_DIR/certs"

if ! command -v mkcert >/dev/null 2>&1; then
    echo "mkcert not found. Install it from https://github.com/FiloSottile/mkcert"
    exit 1
fi

mkcert -install

mkdir -p "$CERTS_DIR"

mkcert -cert-file "$CERTS_DIR/terrabase.crt" -key-file "$CERTS_DIR/terrabase.key" "$TERRABASE_DOMAIN"
chmod 644 "$CERTS_DIR/terrabase.crt"
chmod 644 "$CERTS_DIR/terrabase.key"

if ! grep -q "$TERRABASE_DOMAIN" /etc/hosts; then
    echo "Adding $TERRABASE_DOMAIN to /etc/hosts"
    echo "127.0.0.1 $TERRABASE_DOMAIN" >> /etc/hosts
fi
