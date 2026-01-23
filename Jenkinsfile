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
                   
                    -Dspring-javaformat.skip=true \
                    -Dspring-boot.repackage.skip=true
                '''
            }
        }

       
        stage('Dependency Resolution') {
            steps {
                sh '''
                  mvn dependency:go-offline \
                   
                    -Dspring-javaformat.skip=true
                '''
            }
        }

        stage('Static Analysis / Verify') {
            steps {
                sh '''
                  mvn verify -DskipTests \
                   
                    -Dspring-javaformat.skip=true \
                    -Dspring-boot.repackage.skip=true
                '''
            }
        }

        stage('Rebuild Matrix') {
            steps {
                sh '''
                  for i in 1 2 3; do
                    echo "Rebuild iteration $i"
                    mvn clean package \
                      -DskipTests \
                     
                      -Dspring-javaformat.skip=true \
                      -Dspring-boot.repackage.skip=true
                  done
                '''
            }
        }

         stage('UNIT TEST'){
            steps {
                sh 'mvn test'
            }
        }

	    stage('INTEGRATION TEST'){
            steps {
                sh 'mvn verify -DskipUnitTests'
            }
        }

         stage('Artifact'){
            steps {
               echo 'Now Archiving...'
               archiveArtifacts artifacts: '**/target/*.war'

            }
        }

        
    }
}
