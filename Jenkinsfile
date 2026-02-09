
pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                git 'REPLACE_WITH_YOUR_GITHUB_REPO_URL'
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
