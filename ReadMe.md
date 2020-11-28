# finfiber

Finance apps microservices hosted on google cloud platform app engine.

This contains various finance apis done using golang and microservices for various machine learning dash apps.

#### To run locally

Need google creds to run the nlp locally, the app engine service has permissions for nlp.
```
set GOOGLE_APPLICATION_CREDENTIALS=stocks-272500-1b406ea9557e.json
```

### Golang Features
In order for the golang paths to be resolved for future projects, I decided to have the main go file in the root.

- [x] Using fiber v2
- [ ] Sending email notifications
- [x] Sending discord to webhook
- [x] NLP via google - the google app engine instance has permission to call nlp functions



### Python Features
#### Other Notes

~~DO NOT OPEN SOURCE UNDER ANY CIRCUMSTANCES~~ changed env vars

As for the google cloud platform build CI script, USE WITH CAUTION. I was very unhappy getting charged 1 cent each month to pay for storage I never needed. Remove `gcrgc.sh` in a production environment.

Eventually redo code structure following [project-layout](https://github.com/golang-standards/project-layout)