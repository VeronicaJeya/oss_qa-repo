pipeline {
    agent any
   
    stages {       
       
        stage('Build Docker Image') {
            steps {
                script {                   
                    sh 'docker buildx build --load -t python:3.13.0rc3-bookworm .'
                }
            }
        }
    }
}
