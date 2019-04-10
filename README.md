# cloud-run-go-echo
Go + Echo Framework -> Google Cloud Run

1. Repo Clone -> Docker Build -> GCR Push
```
$ git clone https://github.com/onori/cloud-run-go-echo.git
$ docker build -t  gcr.io/[YOUR-PROJECT-ID]/cloud-run-go-echo:tag1
$ docker push gcr.io/[YOUR-PROJECT-ID]/cloud-run-go-echo:tag1
```

2. Deploy `Cloud Run`

https://cloud.google.com/run/docs/quickstarts/build-and-deploy

```
$ gcloud beta run deploy --image gcr.io/[YOUR-PROJECT-ID]/cloud-run-go-echo:tag1
Please specify a region:
 [1] us-central1
 [2] cancel
Please enter your numeric choice:  1   

To make this the default region, run `gcloud config set run/region us-central1`.

Service name: (cloud-run-go-echo):  
Deploying container to Cloud Run service [cloud-run-go-echo] in project [YOUR-PROJECT-ID] region [us-central1]
Allow unauthenticated invocations to new service [cloud-run-go-echo]? 
(y/N)?  y

✓ Deploying new service... Done.                                                                                                                                                                                                                             
  ✓ Creating Revision... Deploying Revision.                                                                                                                                                                                                                 
  - Routing traffic...                                                                                                                                                                                                                                       
Done.                                                                                                                                                                                                                                                        
Service [cloud-run-go-echo] revision [cloud-run-go-echo-00001] has been deployed and is serving traffic at https://xxxx.app
```
