from typing import Optional
from fastapi import FastAPI, Query
from ytube_dl import get_video

app = FastAPI()

# domain where this api is hosted for example : localhost:5000/docs to see swagger documentation automagically generated.


@app.get("/")
def home():
    return {"message": "Hello TutLinks.com"}


@app.get("/video")
def video(
    q: Optional[str] = Query(
        None,
        max_length=50,
        title="Youtube Video Query",
        description="Query Params - should have youtube id",
    )
):
    if q:
        print(q)
        video_id = q.get("id")
        video_data = get_video(video_id)
        return {"q": q, "video_data": video_data}
    return {"error": "no id in query parameter"}