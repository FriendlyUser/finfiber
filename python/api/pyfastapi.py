from typing import Optional
from fastapi import FastAPI, Query, HTTPException
from ytube_dl import get_video
import uvicorn
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
        video_id = q
        try:
            video_data = get_video(video_id)
            return {"q": q, "video_data": video_data}
        except Exception as e:
            print(e)
            raise HTTPException(status_code=404, detail=str(e))
    return {"error": "no id in query parameter"}

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)