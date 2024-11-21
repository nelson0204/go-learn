# GO LEARNING
This repository showcases my journey through the fundamentals of Go.

## Installation (WSL2 Ubuntu 24.04)

Based on the tutorial (https://www.brettfullam.com/create-a-containerized-go-development-environment-in-docker)

- Dowload go [https://go.dev/dl/go1.23.3.linux-arm64.tar.gz](https://go.dev/dl/)

  ```
  wget https://go.dev/dl/go1.23.3.linux-arm64.tar.gz
  ```

- Unzip the tar.gz package. "-C /usr/local" changes the directory where the tar command will put the extracted output to the /usr/local directory which is where local binaries are located.
  ```
  tar -C /usr/local -xzf go1.17.5.linux-amd64.tar.gz
  ```
-  Add the $PATH and $GOPATH environment variables in the .bashrc
    ```
    nano ~/.bashrc
    export PATH=$PATH:/usr/local/go/bin
    export GOPATH=$HOME/gocode
    ```
- After set variables need reload shell
    ```
    source ~/.bashrc
    ```
- Verify go version
    ```
    go version
    ```
