from __future__ import annotations

import hashlib
import hmac
import secrets

from dataclasses import dataclass


@dataclass
class APIKeyMaterial:
    prefix: str
    secret: str
    token: str
    secret_hash: str


def generate_api_key(
    pepper: str = "", prefix_bytes: int = 6, secret_bytes: int = 32
) -> APIKeyMaterial:
    prefix = secrets.token_hex(prefix_bytes)
    secret = secrets.token_urlsafe(secret_bytes)

    token = f"{prefix}.{secret}"

    secret_hash = hash_api_key_secret(secret, pepper)

    return APIKeyMaterial(
        prefix=prefix, secret=secret, token=token, secret_hash=secret_hash
    )


def hash_api_key_secret(secret: str, pepper: str = "") -> str:
    key = pepper.encode()

    return hmac.new(key, secret.encode(), hashlib.sha256).hexdigest()
