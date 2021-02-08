---
layout: page
title: Google Cloud Run
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Google Cloud Run

To get started in a hosted environment you can deploy this project to the Google Cloud Platform.

From your [Google Cloud dashboard](https://console.cloud.google.com/home/dashboard) create a new project and call it:
```
moov-wire-demo
```

Enable the [Container Registry](https://cloud.google.com/container-registry) API for your project and associate a [billing account](https://cloud.google.com/billing/docs/how-to/manage-billing-account) if needed. Then, open the Cloud Shell terminal and run the following Docker commands, substituting your unique project ID:

```
docker pull moov/wire
docker tag moov/wire gcr.io/<PROJECT-ID>/wire
docker push gcr.io/<PROJECT-ID>/wire
```

Deploy the container to Cloud Run:
```
gcloud run deploy --image gcr.io/<PROJECT-ID>/wire --port 8088
```

Select your target platform to `1`, service name to `wire`, and region to the one closest to you (enable Google API service if a prompt appears). Upon a successful build you will be given a URL where the API has been deployed:

```
https://YOUR-WIRE-APP-URL.a.run.app
```

Now you can list files stored in-memory:
```
curl https://YOUR-WIRE-APP-URL.a.run.app/files
```
You should get this response:
```
null
```