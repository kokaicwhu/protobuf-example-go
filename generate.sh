protoc -I proto/  --go_out=proto/ proto/simple/simple.proto

protoc -I proto/  --go_out=proto/ proto/enumpb/enum.proto

protoc -I proto/  --go_out=proto/ proto/complexpb/complex.proto
