FROM public.ecr.aws/lambda/provided:al2 as build
# install compiler
RUN yum install -y golang
RUN go env -w GOPROXY=direct
# cache dependencies
ADD go.mod go.sum ./
RUN go mod download
# build
ADD . .
RUN go build -o /main
# copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2
COPY --from=build /main /main
ENTRYPOINT [ "/main" ] 


# FROM golang:1.20.1-alpine3.17 as build

# WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./

# RUN go mod download

# COPY ./ ./

# RUN go build -o /finance-api

# FROM alpine:3.17 as publish

# COPY --from=build /finance-api ./

# CMD [ "/finance-api","-r" ]