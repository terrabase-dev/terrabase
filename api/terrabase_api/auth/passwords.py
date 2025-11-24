from __future__ import annotations

from passlib.context import CryptContext

# Argon2id tuned for ~50-150ms on typical servers; adjust via env/config if needed.
_pwd_context = CryptContext(
    schemes=["argon2"],
    deprecated="auto",
    argon2__memory_cost=65536,  # 64MB
    argon2__time_cost=3,
    argon2__parallelism=1,
    argon2__hash_len=32,
    argon2__salt_size=16,
)


class PasswordHasher:
    def hash(self, password: str) -> str:
        return _pwd_context.hash(password)

    def verify(self, password: str, hashed: str) -> bool:
        return _pwd_context.verify(password, hashed)

    def needs_rehash(self, hashed: str) -> bool:
        return _pwd_context.needs_update(hashed)
