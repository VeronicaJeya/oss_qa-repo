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
                sh '''
                mvn clean package \
  -DskipTests \
  -Dcheckstyle.skip=true \
  -Dspring-javaformat.skip=true \
  -Dspring-boot.repackage.skip=true
  '''


            }
        }

        stage('Run (long running)') {
            steps {
               sh 'java -jar target/*.jar'

            }
        }
    }
}
