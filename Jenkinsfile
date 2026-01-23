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
                sh '''
                  mvn dependency:go-offline \
                    -Dcheckstyle.skip=true \
                    -Dspring-javaformat.skip=true
                '''
            }
        }

        stage('Static Analysis / Verify') {
            steps {
                sh '''
                  mvn verify -DskipTests \
                    -Dcheckstyle.skip=true \
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
                      -Dcheckstyle.skip=true \
                      -Dspring-javaformat.skip=true \
                      -Dspring-boot.repackage.skip=true
                  done
                '''
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
