pipeline {
    agent any  
       stages {      

         stage('Clean workspace') {  
            steps {
                cleanWs()
            }
        }   

        stage('Build') {
            steps {                
                echo "Build stage"                
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
