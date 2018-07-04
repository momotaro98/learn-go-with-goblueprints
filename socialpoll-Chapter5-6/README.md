# How to run app

We can run the apps with following commands

## Mongo DB

Run MongoDB

```
$ mongod --dbpath ./data/db
```

## NSQ

Run NSQ

```
$ nsqd --lookupd-tcp-address=localhost:4160
```

# Using Docker

## Mongo DB

Pull mongo image

```
$ docker pull mongo
```

Create docker volume for persistency of data in MongoDB

```
$ docker volume create goblue_mongo_volue
```

docker run from mongo image

```
$ docker run -d --rm --name mongodb --mount source=goblue_mongo_volume,target=/data/db -p 27017:27017 mongo
```

Create data of application into mongoDB from another container using linking function

```
$ docker run -it --rm --link mongodb:mongodb mongo:latest mongo mongodb://mongodb
> use ballots
> db.polls.insert({"options" : ["happy", "sad", "fail", "win"],"results" : {"fail" : 159, "win" : 711, "happy" : 233, "sad" : 166}, "title" : "How do you feel like?" })
> db.polls.find().pretty()
```
