from typing import Dict

from fastapi import APIRouter, Body, Path
from starlette.status import (
    HTTP_400_BAD_REQUEST,
    HTTP_404_NOT_FOUND,
    HTTP_500_INTERNAL_SERVER_ERROR,
)

from terrabase_api.services import OrganizationService
from terrabase_api.specs.terrabase.organization.v1 import organization_pb2
from terrabase_api.types import ErrorResponse
from terrabase_api.utils import (
    call,
    get_paginated_request_params,
    parse,
)

organization_router = APIRouter(
    prefix="/organization",
    responses={
        HTTP_400_BAD_REQUEST: {"model": ErrorResponse},
        HTTP_404_NOT_FOUND: {"model": ErrorResponse},
        HTTP_500_INTERNAL_SERVER_ERROR: {"model": ErrorResponse},
    },
    tags=["organization"],
)

page_size_param, page_token_param = get_paginated_request_params()


@organization_router.get("/", response_model=None)
async def list_organizations(
    service: OrganizationService,
    page_size: int = page_size_param,
    page_token: str | None = page_token_param,
):
    return await call(
        service.ListOrganizations,
        organization_pb2.ListOrganizationsRequest(
            page_size=page_size, page_token=page_token
        ),
    )


@organization_router.post("/", response_model=None)
async def create_organization(
    service: OrganizationService,
    req: Dict = Body(...),
):
    return await call(
        service.CreateOrganization,
        parse(organization_pb2.CreateOrganizationRequest, req),
    )


@organization_router.get("/{id}", response_model=None)
async def get_organization(
    service: OrganizationService,
    id: str = Path(description="The ID of the organization to retrieve"),
):
    return await call(
        service.GetOrganization, organization_pb2.GetOrganizationRequest(id=id)
    )


@organization_router.post("/{id}", response_model=None)
async def update_organization(
    service: OrganizationService,
    req: Dict = Body(...),
    id: str = Path(description="The ID of the organization to update"),
):
    update_req = parse(organization_pb2.UpdateOrganizationRequest, req)
    update_req.id = id

    return await call(service.UpdateOrganization, update_req)


@organization_router.delete("/{id}", response_model=None)
async def delete_organization(
    service: OrganizationService,
    id: str = Path(description="The ID of the organization to delete"),
):
    return await call(
        service.DeleteOrganization,
        organization_pb2.DeleteOrganizationRequest(id=id),
    )
