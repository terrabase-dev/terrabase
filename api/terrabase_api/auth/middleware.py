from __future__ import annotations

from typing import Awaitable, Callable

from fastapi import Request
from jose import JWTError
from starlette.middleware.base import BaseHTTPMiddleware
from starlette.responses import JSONResponse, Response

from terrabase_api.auth.context import (
    AuthContext,
    ClientInfo,
    set_auth_context,
    set_client_info,
)
from terrabase_api.auth.tokens import JWTManager


class AuthContextMiddleware(BaseHTTPMiddleware):
    """
    Parses Authorization/X-API-Key headers, attaches AuthContext for downstream handlers,
    and keeps the raw credential so we can forward to the RPC layer.
    """

    def __init__(
        self,
        app,
        jwt_manager: JWTManager,
        allow_api_key_passthrough: bool = True,
    ):
        super().__init__(app)

        self.jwt_manager = jwt_manager
        self.allow_api_key_passthrough = allow_api_key_passthrough

    async def dispatch(
        self, request: Request, call_next: Callable[[Request], Awaitable[Response]]
    ) -> Response:
        set_auth_context(None)

        set_client_info(
            ClientInfo(
                ip=_extract_ip(request),
                user_agent=request.headers.get("user-agent"),
            )
        )

        auth_header = request.headers.get("authorization")
        api_key_header = request.headers.get("x-api-key")

        if auth_header:
            scheme, _, credential = auth_header.partition(" ")
            scheme = scheme.lower()

            if scheme == "bearer" and credential:
                try:
                    ctx, _ = self.jwt_manager.decode(credential)
                    ctx.raw_authorization = auth_header

                    set_auth_context(ctx)
                except JWTError:
                    return JSONResponse(
                        {"detail": "invalid or expired token"}, status_code=401
                    )
            elif scheme == "apikey" and credential:
                if not self.allow_api_key_passthrough:
                    return JSONResponse(
                        {"detail": "API key auth is not enabled"}, status_code=401
                    )

                set_auth_context(
                    AuthContext(
                        subject_id="",
                        subject_type="api_key",
                        scopes=[],
                        entitlements={},
                        metadata={},
                        auth_scheme="api_key",
                        raw_authorization=auth_header,
                    )
                )
            else:
                return JSONResponse(
                    {"detail": "unsupported auth scheme"}, status_code=401
                )
        elif api_key_header:
            if not self.allow_api_key_passthrough:
                return JSONResponse(
                    {"detail": "API key auth is not enabled"}, status_code=401
                )

            set_auth_context(
                AuthContext(
                    subject_id="",
                    subject_type="api_key",
                    scopes=[],
                    entitlements={},
                    metadata={},
                    auth_scheme="api_key",
                    raw_authorization=api_key_header,
                )
            )

        response = await call_next(request)

        return response


def _extract_ip(request: Request) -> str | None:
    xff = request.headers.get("x-forwarded-for")

    if xff:
        return xff.split(",")[0].strip()

    if request.client:
        return request.client.host

    return None
