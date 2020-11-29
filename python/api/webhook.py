import os
import requests
import json
# if this works make a tutorial of this
def send_webhook(file_path):
    # discord webhook
    webhook = os.getenv('DISCORD_WEBHOOK')
    if webhook != None:
        with open(file_path, 'r') as file_data:
          headers = {'Content-Type': 'application/json'}
          file_contents = file_data.readlines()
          print(file_contents)
          payload = {
            "file": file_contents
          }
          r = requests.post(webhook,
                    data=json.dumps(payload), verify=False, headers=headers)
          # print(r)
          # data = r.json()
          # print(data)
    else:
      print("EXPECTED WEBHOOK TO BE AVAILABLE")
        
if __name__ == "__main__":
  send_webhook("ReadMe.md")
