pipeline {          
    agent any         
    
    stages {  
                
        stage('Git Checkout') {  
            steps {
                git branch: 'build1', credentialsId: 'pat-may15-v-ID', url: 'https://github.com/VeronicaJeya/oss_qa-repo.git' 
                
            }
        }
        
        stage('Build') {
            steps {
                sh '''
                    apt update
                    apt install -y python3  
                    python3 list_akanksha.py                                         
                '''
                
            }
        }
    }
    
   
            
}
