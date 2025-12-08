pipeline {          
    agent any         
    
    stages {  
        stage('Install Python') {     
            steps {
                sh '''  
                    sudo apt update
                    sudo apt install -y python3 python3-venv python3-pip                              
                     echo "Python installation completed."          
                '''    
            }   
        }  

        stage('Python Version') {        
            steps {
                sh 'python3 --version' 
                echo "Python version displayed"
            }
        } 
  
        stage('Git Checkout') {  
            steps {
                git build1: 'build1', credentialsId: 'ssh_dec5_v_ID', url: 'https://github.com/VeronicaJeya/oss_qa-repo.git'   
            }
        }
        
        stage('Build') {
            steps {
                sh '''
                    python3 -m venv venv
                    . venv/bin/activate                    
                    python3 list_akanksha.py                                         
                '''
                
            }
        }
    }
    
    post {
        always {
            script {
                emailext(
                    subject: "pipeline-jenkinsfile-build-email",
                    body: '${JELLY_SCRIPT, template="ad_default_template"}', 
                    mimeType: 'text/html',
                    to: 'jeya.veronica@agilitydelivered.com',           
                                    
                )
            }
        }
    }
}
