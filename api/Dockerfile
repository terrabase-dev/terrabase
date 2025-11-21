FROM python:3.12-slim AS builder

WORKDIR /build

COPY --from=ghcr.io/astral-sh/uv:latest /uv /bin/uv

ENV \
    UV_CACHE_DIR=/tmp/.cache \
    UV_PYTHON_DOWNLOADS=never \
    UV_LOCKED=1

COPY api/pyproject.toml .
RUN uv pip install --group base --target .

ADD api/terrabase_api terrabase_api/

FROM python:3.12-slim

WORKDIR /app

ENV PYTHONPATH="/app/terrabase_api/specs:/app"

COPY --from=builder /build .

CMD ["python", "-m", "fastapi", "run", "terrabase_api/main.py", "--proxy-headers", "--port", "8000"]
