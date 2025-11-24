import os

from typing import Dict, List, Tuple

import grpc

from fastapi.exceptions import HTTPException
from google.protobuf.json_format import MessageToDict, ParseDict, ParseError
from starlette.status import (
    HTTP_400_BAD_REQUEST,
    HTTP_404_NOT_FOUND,
    HTTP_500_INTERNAL_SERVER_ERROR,
)

from terrabase_api.auth import get_auth_context, get_client_info

RPC_ADDR = os.getenv("TERRABASE_RPC_ADDR", "localhost:8080")


def _grpc_status_to_http(status: grpc.StatusCode) -> int:
    return {
        grpc.StatusCode.INVALID_ARGUMENT: HTTP_400_BAD_REQUEST,
        grpc.StatusCode.NOT_FOUND: HTTP_404_NOT_FOUND,
    }.get(status, HTTP_500_INTERNAL_SERVER_ERROR)


def parse(message_cls, data: Dict):
    message = message_cls()
    try:
        ParseDict(data, message)
    except ParseError as err:
        raise HTTPException(status_code=HTTP_400_BAD_REQUEST, detail=str(err))
    return message


async def call(rpc, request) -> Dict:
    metadata: List[Tuple[str, str]] = []
    auth_ctx = get_auth_context()

    if auth_ctx and auth_ctx.raw_authorization:
        header_key = (
            "x-api-key" if auth_ctx.auth_scheme == "api_key" else "authorization"
        )

        metadata.append((header_key, auth_ctx.raw_authorization))

    client_info = get_client_info()
    if client_info:
        if client_info.user_agent:
            metadata.append(("user-agent", client_info.user_agent))
        if client_info.ip:
            metadata.append(("x-forwarded-for", client_info.ip))

    try:
        response = await rpc(request, metadata=metadata or None)
    except grpc.aio.AioRpcError as err:
        raise HTTPException(
            status_code=_grpc_status_to_http(err.code()),
            detail=err.details(),
        ) from err

    return MessageToDict(response, preserving_proto_field_name=True)
