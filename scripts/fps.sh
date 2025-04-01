
mkdir -p tmp

go build -o tmp/fps-go src/fps/main.go
valk build -o tmp/fps-valk src/fps/main.valk

# /usr/bin/time -v tmp/fps-go
# /usr/bin/time -v tmp/fps-valk
tmp/fps-go
tmp/fps-valk
