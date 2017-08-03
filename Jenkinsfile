pipeline {
    agent any
    environment {
        GOCONFIG_PATH="/var/jenkins_home/tools/org.jenkinsci.plugins.golang.GolangInstallation/go"
        SONARCONFIG_PATH = "/var/jenkins_home/tools/hudson.plugins.sonar.SonarRunnerInstallation/sonarqube/sonar-scanner-3.0.3.778"
    }

    stages {

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
                    //sh 'git merge origin/master'
                    sh 'git checkout origin/master'
                    sh 'git merge origin/development'
                    sh 'git checkout origin/master'
                    //sh 'git push origin/master'
                }
            }
        }
    }
}