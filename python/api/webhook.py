import os
import requests
import json
from discord import Webhook, RequestsWebhookAdapter, File

def send_content(content):
    webhook = os.getenv('DISCORD_WEBHOOK')
    if webhook != None:
        webhook = Webhook.from_url(webhook,  adapter=RequestsWebhookAdapter())
        content = content[0:1990]
        webhook.send(content, username='Youtube NLP')
    else:
      print("EXPECTED WEBHOOK TO BE AVAILABLE")

def send_file(file_path, content="Youtube Video"):
    webhook = os.getenv('DISCORD_WEBHOOK')
    if webhook != None:
      with open(file=file_path, mode='rb') as f:
        my_file = File(f)
        webhook = Webhook.from_url(webhook,  adapter=RequestsWebhookAdapter())
        webhook.send(content, username='Youtube NLP', file=my_file)
    else:
      print("EXPECTED WEBHOOK TO BE AVAILABLE")
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
  send_file("ReadMe.md")
