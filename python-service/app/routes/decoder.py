from fastapi import APIRouter, Depends
from fastapi.responses import JSONResponse
from app.controllers.decoder import DecoderController
from app.interfaces.decoder import LLMResponseContext_PM
from app.settings.enums import ServicePaths
from zoldics_service_utils.middlewares import ContextSetter


router = APIRouter(
    prefix=ServicePaths.CONTEXT_PATH.value + "/decoder",
    tags=["HealthCheck"],
    responses={"404": {"description": "Not found"}},
    dependencies=[Depends(ContextSetter())],
)


@router.post("/get-llmresponse")
def get_embeddings(payload: LLMResponseContext_PM):
    llm_response = DecoderController.get_llm_response(payload=payload)
    return JSONResponse(status_code=200, content=llm_response)
