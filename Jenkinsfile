pipeline {
    agent any
    stages {
        stage('Checkout Code') {
            steps {
                echo "Cloning source code..."
                 git branch: 'nodejs1', credentialsId: 'ssh_global_sep15_v_ID', url: 'git@github.com:agilitydelivered/ad-jenkins-qa-nodejs.git'   
            }
        }
        
         stage('Install Node.js') {
            steps {
                sh '''
                    curl -fsSL https://deb.nodesource.com/setup_20.x | sudo -E bash -
                    sudo apt-get install -y nodejs
                    node -v
                    npm -v
                   
                '''
            }
        }
        stage('Install npm') {
            steps {
                echo "Installing npm dependencies..."
                sh 'npm ci' 
            }
        }
        
        stage('run ') {
            steps {
               
                sh 'npm run build' 
            }
        }


        
        
        
    }
}
