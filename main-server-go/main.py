import random
import sys
from fastapi import FastAPI

app = FastAPI()


@app.get("/")
def read_root():
    return {"Hello": "World"}


@app.get("/rand")
def read_item():
    return {"number": random.randint(1, 1000)}


@app.get("/exc")
async def root():
    sys.exit("Exiting the application")


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="127.0.0.1", port=8001)
