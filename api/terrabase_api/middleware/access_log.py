from typing import Awaitable, Callable

from starlette.middleware.base import BaseHTTPMiddleware
from starlette.requests import Request
from starlette.responses import Response

from terrabase_api.logging import terrabase_api_access_logger


class AccessLogMiddleware(BaseHTTPMiddleware):
    async def dispatch(
        self, request: Request, call_next: Callable[[Request], Awaitable[Response]]
    ) -> Response:
        response = await call_next(request)

        client = request.client.host if request.client else "-"
        method = request.method
        path = request.url.path
        http_version = request.scope.get("http_version", "1.1")
        status_code = response.status_code

        if status_code >= 200 and status_code < 400:
            logger_fn = getattr(terrabase_api_access_logger, "info")
        else:
            logger_fn = getattr(terrabase_api_access_logger, "error")

        logger_fn(f'{client} - "{method} {path} HTTP/{http_version}" {status_code}')

        return response
