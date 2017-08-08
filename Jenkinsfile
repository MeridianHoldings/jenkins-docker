pipeline {
    agent any
    environment {
        GOCONFIG_PATH="/var/jenkins_home/tools/org.jenkinsci.plugins.golang.GolangInstallation/go"
        SONARCONFIG_PATH = "/var/jenkins_home/tools/hudson.plugins.sonar.SonarRunnerInstallation/sonarqube"
    }

    stages {
        stage('SonarQube Analysis') {
            steps {
                tool(name: 'go', type: 'go')
                withEnv(["GOROOT=$GOCONFIG_PATH", "PATH+GO=$GOCONFIG_PATH/bin"]) {
                    sh "pwd"
                    //sh "/var/jenkins_home/workspace/go/bin/./gometalinter --checkstyle > report.xml"
                    //sh "go test -coverprofile=covert.out"
                    //sh "./go/bin/gocov convert cover.out | ./go/bin/gocov-xml > coverage.xml"
                }
                script {
                    // requires SonarQube Scanner 2.8+
                    scannerHome = tool 'sonarqube'
                }
                withEnv(["GOROOT=$SONARCONFIG_PATH", "PATH+GO=SONARCONFIG_PATH/bin"]) {
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
            steps {
                tool(name: 'go', type: 'go')
                withEnv(["GOROOT=$GOCONFIG_PATH", "PATH+GO=$GOCONFIG_PATH/bin"]) {
//                    sh 'git checkout origin/master'
//                    sh 'git merge origin/development'
//                    sh 'git checkout origin/master'
                    sh "git fetch --tags --progress https://github.com/Luwade/jenkins-docker.git +refs/heads/development:refs/remotes/origin/development"
                    sh "git push https://github.com/Luwade/jenkins-docker.git HEAD:master -f"
                    //sh 'git push origin HEAD:master'
                }
            }
        }
    }
}