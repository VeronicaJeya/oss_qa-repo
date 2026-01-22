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
                sh ' clean package -DskipTests'
            }
        }

        stage('Run (long running)') {
            steps {
                sh 'java -jar target/long-running-springboot-0.0.1-SNAPSHOT.jar'
            }
        }
    }
}
