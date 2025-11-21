from fastapi import FastAPI
from fastapi.responses import Response
# from fastapi.staticfiles import StaticFiles

from terrabase_api import __version__
from terrabase_api.docs import configure_openapi, get_openapi_yaml
from terrabase_api.routers import organization_router
from terrabase_api.tags import openapi_tags

terrabase_api = FastAPI(
    title="Terrabase API",
    version=__version__,
    root_path="/api",
    openapi_tags=openapi_tags,
    root_path_in_servers=False,
)


terrabase_api.openapi = lambda: configure_openapi(terrabase_api)

# TODO: mount assets
# terrabase_api.mount("/assets", StaticFiles(directory="assets"), name="assets")

# TODO: configure exception handlers

# TODO: add middleware

terrabase_api.include_router(organization_router)


@terrabase_api.get("/openapi.yaml", include_in_schema=False)
def openapi_yaml() -> Response:
    openapi_yaml = get_openapi_yaml(terrabase_api)

    return Response(openapi_yaml, media_type="text/yaml")
