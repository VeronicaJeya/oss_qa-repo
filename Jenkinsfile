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
        stage('parameterisedJob') {
            steps {
echo "${booleann}"
echo "${choice}"

echo "${multiline}"
echo "${password}"
echo "${string}"
echo "${runparameter}"   
echo "${runparameter_JOBNAME}"
echo "${runparameter_NUMBER}"
echo "${runparameter_NAME}"
echo "${runparameter_RESULT}"
echo "${username}"
echo "${ssh}"
echo "${secrettext}"
echo "${secrettext}"

cat File
    }
    
        }
    }
            
}
