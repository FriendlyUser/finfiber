from typing import Optional
from fastapi import FastAPI, Query, HTTPException, Request
from ytube_dl import get_video
from webhook import send_content
import uvicorn

app = FastAPI()

# domain where this api is hosted for example : localhost:5000/docs to see swagger documentation automagically generated.


@app.get("/path")
async def my_route(request: Request) -> None:
    print(request.url.path)


@app.get("/")
def home():
    return {"message": "Hello TutLinks.com"}


# see from starlette.requests import Request
@app.get("/video")
def video(
    request: Request,
    q: Optional[str] = Query(
        None,
        max_length=50,
        title="Youtube Video Query",
        description="Query Params - should have youtube id",
    ),
):
    print("Visited video endpoint")
    print(request.url.path)
    path = request.url.path
    if q:
        video_id = q
        try:
            video_data = get_video(video_id, path)
            # send data to webhook
            # send_content(video_data)
            return {"q": q, "video_data": video_data}
        except Exception as e:
            print(e)
            raise HTTPException(status_code=404, detail=str(e))
    return {"error": "no id in query parameter"}


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)