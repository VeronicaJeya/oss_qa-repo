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

        stage('Dependency Resolution') {
            steps {
                sh 'mvn dependency:go-offline'
            }
        }

        stage('Static Analysis / Verify') {
            steps {
                sh 'mvn verify -DskipTests'
            }
        }

        stage('Rebuild Matrix') {
            steps {
                sh '''
                  for i in 1 2 3; do
                    echo "Rebuild iteration $i"
                    mvn clean package \
                      -DskipTests \
                      -Dcheckstyle.skip=true \
                      -Dspring-javaformat.skip=true \
                      -Dspring-boot.repackage.skip=true
                  done
                '''
            }
        }

        stage('Run Application (Long)') {
            steps {
                timeout(time: 20, unit: 'MINUTES') {
                    sh 'mvn spring-boot:run'
                }
            }
        }
    }
}
