from fastapi import FastAPI
import uvicorn
from fastapi.middleware.cors import CORSMiddleware
from decouple import config

from app.routes import encoder, decoder

app = FastAPI(
    title="AI Service",
    description="A  User-Service for User Related MetaData.",
    version="0.0.1",
)
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

app.include_router(encoder.router)
app.include_router(decoder.router)

if __name__ == "__main__":
    workers = 1
    uvicorn.run(
        "main:app",
        port=3081,
        host="0.0.0.0",
        reload=False,
        workers=workers,
        lifespan="on",
    )
