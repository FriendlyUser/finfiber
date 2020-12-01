from typing import Optional
from fastapi import FastAPI, Query, HTTPException
from ytube_dl import get_video, handle_wbts
from webhook import send_content
import uvicorn

app = FastAPI()

# domain where this api is hosted for example : localhost:5000/docs to see swagger documentation automagically generated.


@app.get("/")
def home():
    try:
        return {"message": "Hello TutLinks.com"}
    except Exception as e:
        print(e)


# see from starlette.requests import Request
@app.get("/video")
def video(
    q: Optional[str] = Query(
        None,
        max_length=50,
        title="Youtube Video Query",
        description="Query Params - should have youtube id",
    ),
):
    if q:
        video_id = q
        try:
            video_data = get_video(video_id)
            # send data to webhook
            # send_content(video_data)
            return {"q": q, "video_data": video_data}
        except Exception as e:
            print(e)
            raise HTTPException(status_code=404, detail=str(e))
    return {"error": "no id in query parameter"}


@app.post("/video/webhook")
def video_transcript(
    transcript_id: Optional[str] = Query(
        None,
        max_length=50,
        title="Assembly AI transcript ID",
        description="Query Params - should have youtube id",
    ),
    status: Optional[str] = Query(
        None,
        max_length=50,
        title="Application Status",
        description="Query Params - should have youtube id",
    ),
):
    try:
        report_data = handle_wbts(transcript_id)
        return {"report_data": report_data}
    except Exception as e:
        print(e)


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)