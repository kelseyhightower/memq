FROM scratch
MAINTAINER Kelsey Hightower <kelsey.hightower@gmail.com>
ADD memq /memq
ENTRYPOINT ["/memq"]
