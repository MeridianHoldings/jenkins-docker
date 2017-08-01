pipeline {
    agent any
    environment {
        GOCONFIG_PATH="/var/jenkins_home/tools/org.jenkinsci.plugins.golang.GolangInstallation/go"
    }

    stages {
        stage('Code Analysis') {
            steps {
                tool(name: 'go', type: 'go')
                withEnv(["GOROOT=$GOCONFIG_PATH", "PATH+GO=$GOCONFIG_PATH/bin"]) {
                    sh 'echo Sonarqube'
                    sh 'printenv'
                }
            }
        }
        stage('SonarQube analysis') {
            steps {
                // requires SonarQube Scanner 2.8+
                def scannerHome = tool 'SonarQube Scanner 2.8';
                withSonarQubeEnv('sonarqube') {
                  sh "${scannerHome}/bin/sonar-scanner"
                }
            }
        }
        stage('Test') {
            steps {
                tool(name: 'go', type: 'go')
                withEnv(["GOROOT=$GOCONFIG_PATH", "PATH+GO=$GOCONFIG_PATH/bin"]) {
                    sh 'go test'
                }
            }
        }
        stage('Build') {
            steps {
                tool(name: 'go', type: 'go')
                withEnv(["GOROOT=$GOCONFIG_PATH", "PATH+GO=$GOCONFIG_PATH/bin"]) {
                    sh 'go build main.go'
                    sh 'echo Build complete'
                }
            }
        }
        stage('Deploy') {
//            when {
//                expression {
//                    currentBuild.result == 'SUCCESS'
//                }
//            }
            steps {
                tool(name: 'go', type: 'go')
                withEnv(["GOROOT=$GOCONFIG_PATH", "PATH+GO=$GOCONFIG_PATH/bin"]) {
                    //sh 'git config --global user.email "luwade.pillay@meridianholdings.co.za"'
                    //sh 'git config --global user.name "Luwade"'
                    //sh 'git checkout origin/master'
                    //sh 'git merge origin/development'
                    //sh 'git push origin master --force'
                    sh 'git merge origin/master'
                    sh 'git checkout origin/master'
                    sh 'git merge origin/development'
                    sh 'git checkout origin/master'
                    sh 'git push origin/master'
                }
            }
        }
    }
}