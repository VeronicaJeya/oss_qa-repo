pipeline {
    agent any
    tools {
        maven 'maven-3'
        
    }


    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build') {
            steps {
                sh 'maven-3 clean package -DskipTests'
            }
        }

        stage('Run (long running)') {
            steps {
                sh 'java -jar target/long-running-springboot-0.0.1-SNAPSHOT.jar'
            }
        }
    }
}
