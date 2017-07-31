pipeline {
    agent any
    environment {
        GOCONFIG_PATH="/home/vagrant/go"
    }

    stages {
        stage('Code Analysis') {
            steps {
                tool(name: 'Go 1.8.3', type: 'go')
                withEnv(["GOROOT=$GOCONFIG_PATH", "GOPATH=$GOCONFIG_PATH"]) {
                    sh 'echo Sonarqube'
                    sh 'printenv'
                }
            }
        }
        stage('Test') {
            steps {
                tool(name: 'Go 1.8.3', type: 'go')
                withEnv(["GOROOT=$GOCONFIG_PATH", "PATH+GO=$GOCONFIG_PATH/bin"]) {
                    sh 'go test'
                }
            }
        }
        stage('Build') {
            steps {
                tool(name: 'Go 1.8.3', type: 'go')
                withEnv(["GOROOT=$GOCONFIG_PATH", "GOPATH=$GOCONFIG_PATH"]) {
                    sh 'go build main.go'
                    sh 'echo Build complete'
                }
            }
        }
        stage('Deploy') {
            when {
                tool(name: 'Go 1.8.3', type: 'go')
                expression {
                    currentBuild.result == 'SUCCESS'
                }
            }
            steps {
                tool(name: 'Go 1.8.3', type: 'go')
                withEnv(["GOROOT=$GOCONFIG_PATH", "PATH+GO=$GOCONFIG_PATH/bin"]) {
                    sh 'git push origin master'
                }
            }
        }
    }
}