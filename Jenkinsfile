pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                checkout scm
                sh 'git rev-parse HEAD > commit.txt'
                sh 'cat commit.txt'
            }
        }

        stage('Build') {
            steps {
                echo "Build stage1"
            }
        }
        
        stage('Test') {
            steps {
                echo "Test stage"
            }
        }
        
        stage('Deploy') {
            steps {
                echo "Deploy stage"
            }
        }
    }
}
