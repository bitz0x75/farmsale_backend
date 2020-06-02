# Hello world example with go
# We specify the base image we need for our
# go application
FROM golang:1.10

# We create an /app directory within our
# image that will hold our application source
# files
RUN mkdir /app

# We copy everything in the root directory
# into our /app directory
ADD . /app
RUN touch /app/.env
# We specify that we now wish to execute 
# any further commands inside our /app
# directory

WORKDIR /app

# #Import local packages
RUN go get "github.com/maxwellgithinji/farmsale_backend/config/mdb"
RUN go get "github.com/maxwellgithinji/farmsale_backend/controllers/productscontroller"
RUN go get "github.com/maxwellgithinji/farmsale_backend/models/productsmodel"

# we run go build to compile the binary
# executable of our Go program
RUN go build -o main

# Our start command which kicks off
# our newly created binary executable
CMD ["/app/main"]