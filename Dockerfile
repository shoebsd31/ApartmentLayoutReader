FROM golang:1.15-alpine as dev

WORKDIR /work

# FROM golang:1.15-alpine as build

# WORKDIR /apartmentLayout
# COPY ./apartmentLayout/* /apartmentLayout/
# RUN go build -o apartmentLayout


# FROM alpine as runtime 
# COPY --from=build /apartmentLayout/apartmentLayout /usr/local/bin/apartmentLayout
# COPY ./apartmentLayout/apartmentLayout.json /
# COPY run.sh /
# RUN chmod +x /run.sh
# ENTRYPOINT [ "./run.sh" ]