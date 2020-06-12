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
RUN go get -u "github.com/maxwellgithinji/farmsale_backend/config/mdb"
RUN go get -u "github.com/maxwellgithinji/farmsale_backend/controllers/productscontroller"
RUN go get -u "github.com/maxwellgithinji/farmsale_backend/models/productsmodel"
RUN go get -u "github.com/maxwellgithinji/farmsale_backend/models/usersmodel"
RUN go get -u "github.com/maxwellgithinji/farmsale_backend/controllers/userscontroller"
RUN go get -u "github.com/maxwellgithinji/farmsale_backend/utils"
RUN go get -u "github.com/maxwellgithinji/farmsale_backend/middleware/auth"
RUN go get -u "github.com/maxwellgithinji/farmsale_backend/routes"
RUN go get -u "github.com/maxwellgithinji/farmsale_backend/models/jwtmodel"

#Import experimental packages
RUN go get "golang.org/x/crypto/bcrypt"
RUN go get "go.mongodb.org/mongo-driver/x/bsonx"
# we run go build to compile the binary
# executable of our Go program
RUN go build -o main

# Our start command which kicks off
# our newly created binary executable
CMD ["/app/main"]