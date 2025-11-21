import os

from typing import Annotated, Dict

import grpc

from fastapi import Depends
from fastapi.exceptions import HTTPException
from google.protobuf.json_format import MessageToDict, ParseDict, ParseError
from starlette.status import (
    HTTP_400_BAD_REQUEST,
    HTTP_404_NOT_FOUND,
    HTTP_500_INTERNAL_SERVER_ERROR,
)

from terrabase_api.specs.terrabase.organization.v1 import organization_pb2_grpc

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
    try:
        response = await rpc(request)
    except grpc.aio.AioRpcError as err:
        raise HTTPException(
            status_code=_grpc_status_to_http(err.code()),
            detail=err.details(),
        ) from err

    return MessageToDict(response, preserving_proto_field_name=True)


async def get_organization_stub() -> organization_pb2_grpc.OrganizationServiceStub:
    channel = grpc.aio.insecure_channel(RPC_ADDR)

    return organization_pb2_grpc.OrganizationServiceStub(channel)


OrganizationService = Annotated[
    organization_pb2_grpc.OrganizationServiceStub, Depends(get_organization_stub)
]
