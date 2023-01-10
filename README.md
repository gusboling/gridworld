# grid_simulator
_A 2D Simulation/Game Engine Written In Golang_

## Overview
The purpose of this project is to practice my golang coding while also providing
a generalized framework for simulating a 2 dimensional environment.

## Architecture
This project should follow a model-view-controller architecture (mostly because
its the first relevant thing I found on Wikipedia; fuck it).

A random search of the interwebs finds us [this post](https://www.reddit.com/r/golang/comments/a3lojm/comment/eb7usaf/?utm_source=share&utm_medium=web2x&context=3) on the "/r/golang" subreddit, which gives us the following comment:

>MVC in Go as I would do it.. would be a front end UI library like React, the VIEW, making API calls to a back end Go API the controller, and using Models as the request/response payloads, implemented in GO on the API side to convert JSON to/from the Go model classes.

So that's what we're doing I guess.

### Model
Is responsible for both the definition of data structures and the source of
truth at runtime. Common structs/serialization, and separetly, persistence 
(database) support.

## View
Is responsible for rendering current state.

## Controller
Is responsible for updating the model.
