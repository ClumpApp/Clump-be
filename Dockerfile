FROM scratch

WORKDIR /
COPY /clump /clump

EXPOSE 8080

ENTRYPOINT ["/clump"]
