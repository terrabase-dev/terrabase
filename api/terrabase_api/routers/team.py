from typing import Dict

from fastapi import APIRouter, Body, Path
from starlette.status import (
    HTTP_400_BAD_REQUEST,
    HTTP_404_NOT_FOUND,
    HTTP_500_INTERNAL_SERVER_ERROR,
)

from terrabase_api.services import TeamService
from terrabase_api.specs.terrabase.team.v1 import team_pb2
from terrabase_api.types import ErrorResponse
from terrabase_api.utils import call, get_paginated_request_params, parse

team_router = APIRouter(
    prefix="/team",
    responses={
        HTTP_400_BAD_REQUEST: {"model": ErrorResponse},
        HTTP_404_NOT_FOUND: {"model": ErrorResponse},
        HTTP_500_INTERNAL_SERVER_ERROR: {"model": ErrorResponse},
    },
    tags=["team"],
)

page_size_param, page_token_param = get_paginated_request_params()


@team_router.get("/", response_model=None)
async def list_teams(
    service: TeamService,
    page_size: int = page_size_param,
    page_token: str | None = page_token_param,
):
    return await call(
        service.ListTeams,
        team_pb2.ListTeamsRequest(page_size=page_size, page_token=page_token),
    )


@team_router.post("/", response_model=None)
async def create_team(service: TeamService, req: Dict = Body(...)):
    return await call(service.CreateTeam, parse(team_pb2.CreateTeamRequest, req))


@team_router.get("/{id}", response_model=None)
async def get_team(
    service: TeamService, id: str = Path(description="The ID of the team to retrieve")
):
    return await call(service.GetTeam, team_pb2.GetTeamRequest(id=id))


@team_router.post("/{id}", response_model=None)
async def update_team(
    service: TeamService,
    req: Dict = Body(...),
    id: str = Path(description="The ID of the team to update"),
):
    update_req = parse(team_pb2.UpdateTeamRequest, req)
    update_req.id = id

    return await call(service.UpdateTeam, update_req)


@team_router.delete("/{id}", response_model=None)
async def delete_team(
    service: TeamService, id: str = Path(description="The ID of the team to delete")
):
    return await call(service.DeleteTeam, team_pb2.DeleteTeamRequest(id=id))
