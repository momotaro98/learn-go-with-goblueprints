# How to deploy

## Set the project

```
$ gcloud config set project answersapp-211814
```

## Run a service app on local

```
$ cd answersapp
$ dev_appserver.py dispatch.yaml default/app.yaml api/app.yaml web/app.yaml
```

## Update/Create Indexes

```
$ cd answersapp
$ gcloud app deploy index.yaml
```

## Deploy services to GAE

```
$ cd answersapp
$ gcloud app deploy dispatch.yaml default/app.yaml api/app.yaml web/app.yaml
```

# Deployed app link

https://answersapp-211814.appspot.com/

