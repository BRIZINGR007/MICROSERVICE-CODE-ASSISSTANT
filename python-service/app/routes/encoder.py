from fastapi import APIRouter, Depends, Query
from fastapi.responses import JSONResponse
from app.encoders.EncoderHandlers import EmbeddingGenerator
from app.settings.enums import ServicePaths
from zoldics_service_utils.middlewares import ContextSetter


router = APIRouter(
    prefix=ServicePaths.CONTEXT_PATH.value + "/encoder",
    tags=["HealthCheck"],
    responses={"404": {"description": "Not found"}},
    dependencies=[Depends(ContextSetter())],
)


@router.get("/get-embedding")
def get_embeddings(query: str = Query(...)):
    embedding = EmbeddingGenerator().generate_embedding(query)
    return JSONResponse(status_code=200, content=embedding)
