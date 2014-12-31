FROM allomov/golang

MAINTAINER allomov

# add project
ADD ./src /gopath/src/github.com/allomov/iv-simulator
RUN cd /gopath/src/github.com/allomov/iv-simulator && godep restore

RUN go install github.com/allomov/iv-simulator



# default is with gmond for seeing something
CMD ["--with-gmond"]
