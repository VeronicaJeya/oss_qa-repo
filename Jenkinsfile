pipeline {          
    agent any         
    
    stages {        
        
        stage('Git Checkout') {     
            steps {
                git branch: 'build1', credentialsId: 'ssh_dec5_v_ID', url: 'https://github.com/VeronicaJeya/oss_qa-repo.git' 
                
            }
        }
        
        stage('Build') {
            steps {
                sh '''
                             
                    python3 list_akanksha.py                                         
                '''
                
            }
        }
    }
    
   
            
}
