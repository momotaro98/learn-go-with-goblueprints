# How to deploy

## Set the project

```
$ gcloud config set project answersapp-211814
```

## Run app on local

```
$ cd api
$ dev_appserver.app app.yaml
```

## Deploy app to GAE

```
$ gcloud app deploy app.yaml
```
