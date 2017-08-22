#!groovy​

pipeline {
    agent {docker 'golang:1.8.3'}
    environment {
        GOROOT = "${JENKINS_HOME}/tools/org.jenkinsci.plugins.golang.GolangInstallation/go"
        // GOCONFIG_PATH="/var/jenkins_home/tools/org.jenkinsci.plugins.golang.GolangInstallation/go"
        SONARCONFIG_PATH = "/var/jenkins_home/tools/hudson.plugins.sonar.SonarRunnerInstallation/sonarqube"
    }
    String goPath = "/home/tomcat/go"
    stages {
        // Start a docker container using the golang:1.8.0-alpine image, mount the current directory to the goPath we specified earlier
        stage("Create binaries") {
            agent {docker 'golang:1.8.3'}
            steps {
                script {
                    docker.image("golang:1.8.3").inside("-v ${pwd()}:${goPath}") {
                        for (command in binaryBuildCommands) {
                            sh "go version"
                        }
                    }
                }
            }
        }
        stage('SonarQube Analysis') {
            steps {
                script {
                    goHome = tool 'go'
                    dockerHome = tool 'docker'
                }
                withEnv(["PATH+GO=${GOROOT}/bin", "PATH+GIT=${GIT_EXEC_PATH}"]) {
                    sh "pwd"
                    sh "go env"
                    sh "${dockerHome}/bin/docker ps"
                    // sh "whoami && go get -u github.com/alecthomas/gometalinter"
                    // sh "cd /home/tomcat/go/src/github.com && ls"
                    sh "${goHome}/bin/gometalinter"
                    //sh "/var/jenkins_home/workspace/go/bin/./gometalinter --checkstyle > report.xml"
                    //sh "go test -coverprofile=covert.out"
                    //sh "/var/jenkins_home/workspace/go/bin/./gocov convert cover.out | /var/jenkins_home/workspace/go/bin/./gocov-xml > coverage.xml"
                    //sh "go test -v ./... | /var/jenkins_home/workspace/go/bin/./go-junit-report > test.xml"
                }
                script {
                    // requires SonarQube Scanner 2.8+
                    scannerHome = tool 'sonar'
                }
                withEnv(["GOROOT=$SONARCONFIG_PATH", "PATH+GO=SONARCONFIG_PATH/bin"]) {
                    sh "${scannerHome}/bin/sonar-scanner"
                }
            }
        }
        stage('Test') {
            steps {
                tool(name: 'go', type: 'go')
                withEnv(["PATH+GO=${GOROOT}/bin", "PATH+GIT=${GIT_EXEC_PATH}"]) {
                    sh 'go test'
                }
            }
        }
        stage('Build') {
            steps {
                tool(name: 'go', type: 'go')
                withEnv(["PATH+GO=${GOROOT}/bin", "PATH+GIT=${GIT_EXEC_PATH}"]) {
                    sh 'go build main.go'
                    sh 'echo Build complete'
                }
            }
        }
        stage('Deploy') {
            steps {
                tool(name: 'go', type: 'go')
                withEnv(["PATH+GO=${GOROOT}/bin", "PATH+GIT=${GIT_EXEC_PATH}"]) {
                    sh "git fetch --tags --progress https://${env.username}:${env.password}@github.com/MeridianHoldings/jenkins-docker.git +refs/heads/development:refs/remotes/origin/development"
                    sh "git push https://${env.username}:${env.password}@github.com/MeridianHoldings/jenkins-docker.git HEAD:master -f"
                    //sh 'git push origin HEAD:master'
                }
            }
        }
        stage('Build Image') {
            steps {
                script {
                    docker_home = tool 'docker'
                }
                sh "${docker_home}/bin/docker build -t my-app ."
                sh "${docker_home}/bin/docker run my-app"
            }
        }
    }
}