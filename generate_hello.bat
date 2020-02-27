@rem Generate the Go code for .proto files

setlocal

@rem enter this directory
cd /d %~dp0

set TOOLS_PATH=C:\Users\hemin\.nuget\packages\grpc.tools\2.25.0\tools\windows_x64
%TOOLS_PATH%\protoc -I../../../ -I=. --go_out=../../../ ./messages/messages.proto
%TOOLS_PATH%\protoc -I../../../ -I=. --go_out=plugins=grpc:. ./helloworld/helloworld.proto

endlocal