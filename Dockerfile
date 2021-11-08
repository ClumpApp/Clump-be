FROM scratch

WORKDIR /
COPY /home/runner/work/clump-be/clump-be/clump /clump

EXPOSE 8080

ENTRYPOINT ["/clump"]
