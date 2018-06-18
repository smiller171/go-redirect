go-redirect [![Docker Build Status](https://img.shields.io/docker/build/millergeek/ssl-redirect.svg)]()
===========

[![Gitter](https://img.shields.io/gitter/room/smiller171/go-redirect.js.svg)](https://gitter.im/smiller171/go-redirect)


Redirect HTTP to HTTPS with Docker

[![](https://images.microbadger.com/badges/version/millergeek/ssl-redirect.svg)](https://microbadger.com/images/millergeek/ssl-redirect "Get your own version badge on microbadger.com")  
[Docker Hub](https://hub.docker.com/r/millergeek/ssl-redirect/)
# Supported Tags and `Dockerfile` links
* `latest` [(Dockerfile)](https://github.com/smiller171/go-redirect/blob/master/Dockerfile) - [![](https://images.microbadger.com/badges/image/millergeek/ssl-redirect.svg)](https://microbadger.com/images/scottmiller171/ssl-redirect "Get your own image badge on microbadger.com")
* `1.3` [(Dockerfile)](https://github.com/smiller171/go-redirect/blob/1.3/Dockerfile) - [![](https://images.microbadger.com/badges/image/millergeek/ssl-redirect:1.3.svg)](https://microbadger.com/images/millergeek/ssl-redirect:1.3 "Get your own image badge on microbadger.com")

# What is this?
This image accepts any http request and redirects to the https version of the same page.  
It does not serve any pages.

# What is this for?
The primary reason for this image is for doing SSL termination on a load balancer (eg: AWS ELB) in the simplest possible way.

# How do I use it?
* Run this image on any TCP port, and your own web server image on another.
```bash
docker run -d -p 8080:80 millergeek/ssl-redirect
docker run -d -p 80:80 my-web-server-image
```
* Forward HTTP 80 on your load balancer to whatever port this image is running on.
* Forward HTTPS 443 on your load balancer to whatever port your own web server is running on.

All HTTP requests will hit this image and be redirected to make HTTPS request. This is abstracted from the user.

# How do I build it myself?

```sh
docker build -t ssl-redirect .
docker run ssl-redirect
```
