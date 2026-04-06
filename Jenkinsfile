pipeline {
    agent any

    stages {
              
        stage('Checkout') {
            steps {
                git branch: 'restartStage', credentialsId: 'oss-ad-pat-Jan30', url: 'https://github.com/VeronicaJeya/oss_qa-repo.git' 
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
