from __future__ import annotations

from datetime import datetime, timedelta, timezone
from typing import Any, Dict, Optional, Tuple

from jose import JWTError, jwt

from terrabase_api.auth.context import AuthContext


class JWTManager:
    def __init__(
        self,
        secret: str,
        issuer: str = "",
        audience: Optional[str] = None,
        algorithm: str = "HS256",
        access_ttl: timedelta = timedelta(minutes=15),
    ):
        if not secret:
            raise ValueError("JWT secret is required")

        self.secret = secret
        self.issuer = issuer
        self.audience = audience
        self.algorithm = algorithm
        self.access_ttl = access_ttl

    def issue_access_token(
        self,
        subject_id: str,
        subject_type: str = "user",
        *,
        scopes: Optional[list[str]] = None,
        entitlements: Optional[Dict[str, list[str]]] = None,
        metadata: Optional[Dict[str, Any]] = None,
        token_id: Optional[str] = None,
        expires_at: Optional[datetime] = None,
    ) -> str:
        now = datetime.now(timezone.utc)
        exp = expires_at or now + self.access_ttl

        payload: Dict[str, Any] = {
            "sub": subject_id,
            "sub_type": subject_type,
            "scopes": scopes or [],
            "entitlements": entitlements or {},
            "metadata": metadata or {},
            "iat": int(now.timestamp()),
            "exp": int(exp.timestamp()),
            "nbf": int(now.timestamp()),
        }

        if token_id:
            payload["jti"] = token_id

        if self.issuer:
            payload["iss"] = self.issuer

        if self.audience:
            payload["aud"] = self.audience

        return jwt.encode(payload, self.secret, algorithm=self.algorithm)

    def decode(self, token: str) -> Tuple[AuthContext, Dict[str, Any]]:
        options = {"verify_aud": self.audience is not None}

        decoded = jwt.decode(
            token,
            self.secret,
            algorithms=[self.algorithm],
            audience=self.audience,
            issuer=self.issuer or None,
            options=options,
        )

        ctx = AuthContext(
            subject_id=decoded["sub"],
            subject_type=decoded.get("sub_type", "user"),
            name=decoded.get("name"),
            email=decoded.get("email"),
            default_role=decoded.get("default_role"),
            scopes=decoded.get("scopes", []),
            entitlements=decoded.get("entitlements", {}),
            metadata=decoded.get("metadata", {}),
            token_id=decoded.get("jti"),
            auth_scheme="bearer",
            raw_authorization=f"Bearer {token}",
        )

        return ctx, decoded


__all__ = ["JWTManager", "JWTError"]
