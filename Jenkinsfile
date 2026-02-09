
pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                git branch: 'golang2', credentialsId: 'ssh_dec5_v_ID', url: 'https://github.com/VeronicaJeya/oss_qa-repo.git' 
            }
        }       

        stage('go version') {
            steps {
                    go clean -cache
                    wget https://go.dev/dl/go1.24.0.linux-amd64.tar.gz
                    sudo rm -rf /usr/local/go
                    sudo tar -C /usr/local -xzf go1.24.0.linux-amd64.tar.gz
                    export PATH=/usr/local/go/bin:$PATH                    
                    which go
                    go version
            }
        }


        stage('Build') {
            steps {
                sh 'go build .'
            }
        }

        stage('Run App') {
            steps {
                sh 'go run main.go'
            }
        }

        stage('Run Tests') {
            steps {
                sh 'go test ./... -v'
            }
        }
    }
}
