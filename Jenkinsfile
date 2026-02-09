
pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                checkout scm
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
