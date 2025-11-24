import os

from fastapi import FastAPI
from fastapi.responses import Response
# from fastapi.staticfiles import StaticFiles

from terrabase_api import __version__
from terrabase_api.auth.tokens import JWTManager
from terrabase_api.docs import configure_openapi, get_openapi_yaml
from terrabase_api.middleware import AccessLogMiddleware, AuthContextMiddleware
from terrabase_api.routers import auth_router, organization_router, user_router
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

JWT_SECRET = os.getenv("AUTH_JWT_SECRET")
if not JWT_SECRET:
    raise RuntimeError("AUTH_JWT_SECRET is required")

jwt_manager = JWTManager(
    secret=JWT_SECRET,
    issuer=os.getenv("AUTH_JWT_ISSUER", ""),
    audience=os.getenv("AUTH_JWT_AUDIENCE"),
)

terrabase_api.add_middleware(AuthContextMiddleware, jwt_manager=jwt_manager)
terrabase_api.add_middleware(AccessLogMiddleware)

terrabase_api.include_router(auth_router)
terrabase_api.include_router(organization_router)
terrabase_api.include_router(user_router)


@terrabase_api.get("/openapi.yaml", include_in_schema=False)
def openapi_yaml() -> Response:
    openapi_yaml = get_openapi_yaml(terrabase_api)

    return Response(openapi_yaml, media_type="text/yaml")
