from typing import Annotated

import grpc

from fastapi import Depends

from terrabase_api.specs.terrabase.organization.v1 import organization_pb2_grpc
from terrabase_api.utils import RPC_ADDR


async def get_organization_stub() -> organization_pb2_grpc.OrganizationServiceStub:
    channel = grpc.aio.insecure_channel(RPC_ADDR)

    return organization_pb2_grpc.OrganizationServiceStub(channel)


OrganizationService = Annotated[
    organization_pb2_grpc.OrganizationServiceStub, Depends(get_organization_stub)
]
