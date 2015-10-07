#/bin/bash
echo "Checking if brew is installed...";
if hash brew 2>/dev/null; then
        echo "Brew is installed, proceeding with installation"
    else
        echo "Brew is not installed, installing..."
        ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
    fi
echo "Installing Go Programming Language..."
brew install go
echo "Specify the full path of the Go Workspace: "
read go_path
echo "Settings commands..."
echo -e "hm_update(){echo Updating...; cd $go_path/src/github.com; sudo rm -rf padnezz; go get github.com/padnezz/kahoot-hack-mod;}" >> $HOME/.bash_profile
echo -e "hm_flood(){cd $go_path/src/github.com/padnezz/kahoot-hack-mod; go run flood.go;}" >> $HOME/.bash_profile
echo -e "hm_crash(){cd $go_path/src/github.com/padnezz/kahoot-hack-mod; go run crash.go;}" >> $HOME/.bash_profile
echo -e "hm_regular(){cd $go_path/src/github.com/padnezz/kahoot-hack-mod; go run regular.go;}" >> $HOME/.bash_profile
echo -e "export GOPATH=$go_path" >> $HOME/.bash_profile
cd $go_path
echo "Installing the hack and necessary components..."
go get github.com/padnezz/kahoot-hack-mod
go get github.com/gorilla/websocket