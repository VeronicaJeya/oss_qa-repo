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
