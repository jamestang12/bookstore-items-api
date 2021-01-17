# Start from base image 1.15.16
FROM golang:1.15.6

# Configure the repo url so it can configure the work directory
ENV REPO_URL=/Users/jamestang/Desktop/Project/bookstore-items-api

# ENV GOPATH=/Users/jamestang/go

ENV APP_PATH=${REPO_URL}

# Cpoy the entire source code from the current directory to $WORKPATH
ENV WORKPATH=${APP_PATH}/src
COPY src $WORKPATH
WORKDIR ${WORKPATH}

RUN go build -o items-api .

# Expose port 8090 to the world:
EXPOSE 8090

CMD ["./items-api"]


# Step
# docker build -t main .
# docker build -t items-api .
# docker run -p 8090:8090 -p 9200:9200 items-api:latest or docker run -p 8090:8090 items-api:latest