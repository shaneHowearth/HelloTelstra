FROM iron/go:dev
WORKDIR /app
# Set the source directory.
ENV SRC_DIR=`pwd`
# Add the source code:
ADD . $SRC_DIR
# Build it:
RUN cd $SRC_DIR; go build -o helloapp; cp helloapp /app/
ENTRYPOINT ["./helloapp"]
