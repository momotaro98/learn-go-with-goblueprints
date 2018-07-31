# How to deploy

## Set the project

```
$ gcloud config set project answersapp-211814
```

## Run a service app on local

```
$ cd api
$ dev_appserver.app app.yaml
```

## Deploy services to GAE

```
$ gcloud app deploy dispatch.yaml default/app.yaml api/app.yaml web/app.yaml
```
