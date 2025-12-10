pipeline {          
    agent any         
    
    stages {        
        
        stage('Git Checkout') {     
            steps {
                git branch: 'build_invalid_cred', credentialsId: 'ssh_dec5_v_I..', url: 'https://github.com/VeronicaJeya/oss_qa-repo.git'    
                
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
